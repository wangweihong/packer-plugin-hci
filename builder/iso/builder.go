package iso

import (
	"context"

	"github.com/wangweihong/packer-plugin-hci/builder/common"
	commondriver "github.com/wangweihong/packer-plugin-hci/builder/common"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/communicator"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/multistep/commonsteps"
	"github.com/hashicorp/packer-plugin-sdk/packer"
)

type Builder struct {
	config Config
	runner multistep.Runner
}

func (b *Builder) ConfigSpec() hcldec.ObjectSpec { return b.config.FlatMapstructure().HCL2Spec() }

func (b *Builder) Prepare(raws ...interface{}) ( /*generatedVars*/ []string /*warnings*/, []string /*err*/, error) {
	warnings, errs := b.config.Prepare(raws...)
	if errs != nil {
		return nil, warnings, errs
	}
	return nil, warnings, nil
}

func (b *Builder) Run(ctx context.Context, ui packer.Ui, hook packer.Hook) (packer.Artifact, error) {
	driver, err := NewDriver(b.config.RunConfig)
	if err != nil {
		return nil, err
	}
	// Setup the state bag and initial state for the steps
	state := new(multistep.BasicStateBag)
	state.Put("config", &b.config)
	state.Put("access_config", &b.config.AccessConfig)
	state.Put("hook", hook)
	state.Put("ui", ui)
	state.Put("debug", b.config.PackerDebug)
	state.Put("instance_ip", driver.BoundIP)

	// Set the value of the generated data that will become available to provisioners.
	// To share the data with post-processors, use the StateData in the artifact.
	state.Put("generated_data", map[string]interface{}{
		"GeneratedMockData": "mock-build-data",
	})

	// Build the steps
	steps := []multistep.Step{
		&commondriver.StepAuth{},
		&commondriver.StepLoadTenant{
			TargetTenant: b.config.TargetTenant,
		},
		&commondriver.StepLoadSpec{
			Spec: b.config.Specification,
		},

		&StepLoadISO{
			Repository: b.config.RepositoryName,
			ISO:        b.config.ISOName,
		},
		&commondriver.StepLoadResourcePool{},

		&StepHTTPIPDiscover{},
		commonsteps.HTTPServerFromHTTPConfig(&b.config.HTTPConfig),
		//TODO: 支持使用非默认资源池和计算组
		// 加载默认资源池
		//// 加载默认计算组
		//&StepLoadComputeGroup{},
		&StepRunServer{
			Name:           b.config.InstanceName,
			RootVolumeSize: b.config.DiskSize,
		},

		// ssh或者winrm远程连接

		multistep.If(b.config.Comm.Type == "ssh", &communicator.StepSSHKeyGen{
			CommConf:            &b.config.Comm,
			SSHTemporaryKeyPair: b.config.Comm.SSHTemporaryKeyPair,
		}),
		&common.StepVncConnect{},
		&StepVncBootCommand{
			Config: b.config.VNCConfig,
			Comm:   &b.config.Comm,
			VMName: b.config.InstanceName,
			Ctx:    b.config.ctx,
		},
		&communicator.StepConnect{
			Config:    &b.config.Comm,
			Host:      driver.HostIP,
			SSHConfig: b.config.Comm.SSHConfigFunc(),
		},
		&commonsteps.StepProvision{},
		&commonsteps.StepCleanupTempKeys{
			//	Comm: &b.config.SSHConfig.Comm,
		},

		//&StepKeyPair{
		//	Debug:        b.config.PackerDebug,
		//	Comm:         &b.config.Comm,
		//	DebugKeyPath: fmt.Sprintf("ecs_%s.pem", b.config.PackerBuildName),
		//},
		//&StepCreateNetwork{
		//	VpcID:          b.config.VpcID,
		//	Subnets:        b.config.Subnets,
		//	SecurityGroups: b.config.SecurityGroups,
		//},

		//&StepRunSourceServer{
		//	Name:             b.config.InstanceName,
		//	VpcID:            b.config.VpcID,
		//	Subnets:          b.config.Subnets,
		//	SecurityGroups:   b.config.SecurityGroups,
		//	RootVolumeType:   b.config.VolumeType,
		//	RootVolumeSize:   b.config.VolumeSize,
		//	KmsKeyID:         b.config.KmsKeyID,
		//	UserData:         b.config.UserData,
		//	UserDataFile:     b.config.UserDataFile,
		//	InstanceMetadata: b.config.InstanceMetadata,
		//},
		//&StepAttachVolume{
		//	PrefixName: b.config.InstanceName,
		//},
		//&StepGetPassword{
		//	Debug: b.config.PackerDebug,
		//	Comm:  &b.config.RunConfig.Comm,
		//},
		//&communicator.StepConnect{
		//	Config:    &b.config.RunConfig.Comm,
		//	Host:      CommHost(b.config.RunConfig.Comm.SSHHost),
		//	SSHConfig: b.config.RunConfig.Comm.SSHConfigFunc(),
		//},
		//&commonsteps.StepProvision{},
		//&commonsteps.StepCleanupTempKeys{
		//	Comm: &b.config.RunConfig.Comm,
		//},
		//&StepStopServer{},
		//&stepCreateImage{
		//	WaitTimeout: b.config.WaitImageReadyTimeout,
		//},
		//&stepAddImageMembers{},
	}

	// Run!
	b.runner = commonsteps.NewRunner(steps, b.config.PackerConfig, ui)
	b.runner.Run(ctx, state)

	// If there was an error, return that
	if err, ok := state.GetOk("error"); ok {
		return nil, err.(error)
	}

	artifact := &commondriver.Artifact{
		//ImageId:        state.Get("image").(string),
		//BuilderIdValue: BuilderId,
		//Client:         imsClient,
	}
	return artifact, nil
}
