package vmx

import (
	"context"

	"github.com/wangweihong/packer-plugin-hci/hci/example"

	commondriver "github.com/wangweihong/packer-plugin-hci/builder/common"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/communicator"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/multistep/commonsteps"
	"github.com/hashicorp/packer-plugin-sdk/packer"
)

const (
	BuilderId = "wangweihong.hci"
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
	//state.Put("instance_ip", driver.BoundIP)

	// Set the value of the generated data that will become available to provisioners.
	// To share the data with post-processors, use the StateData in the artifact.
	state.Put("generated_data", map[string]interface{}{
		"GeneratedMockData": "mock-build-data",
	})

	// Build the steps
	steps := []multistep.Step{
		// 确认目标租户是否存在
		&commondriver.StepAuth{},
		&commondriver.StepLoadTenant{
			TargetTenant: b.config.TargetTenant,
		},
		// 检测规格是否存在
		&commondriver.StepLoadSpec{
			Spec: b.config.Specification,
		},
		//&StepKeyPair{
		//	Debug:        b.config.PackerDebug,
		//	Comm:         &b.config.Comm,
		//	DebugKeyPath: fmt.Sprintf("ecs_%s.pem", b.config.PackerBuildName),
		//},
		&StepLoadSourceImage{
			Repository:  b.config.RepositoryName,
			SourceImage: b.config.SourceImage,
		},
		&commondriver.StepLoadResourcePool{},

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

		&communicator.StepConnect{
			Config:    &b.config.Comm,
			Host:      driver.HostIP,
			SSHConfig: b.config.Comm.SSHConfigFunc(),
		},
		&commonsteps.StepProvision{},
		&commonsteps.StepCleanupTempKeys{
			Comm: &b.config.SSHConfig.Comm,
		},

		&commondriver.StepCreateImage{
			Config: b.config.ImageConfig,
		},
	}

	// Run!
	b.runner = commonsteps.NewRunner(steps, b.config.PackerConfig, ui)
	b.runner.Run(ctx, state)

	// If there was an error, return that
	if err, ok := state.GetOk("error"); ok {
		return nil, err.(error)
	}
	// If there are no images, then just return
	if _, ok := state.GetOk("image"); !ok {
		return nil, nil
	}
	// Build the artifact and return it
	artifact := &commondriver.Artifact{
		Cluster:        state.Get("cluster").(string),
		ImageId:        state.Get("image").(string),
		BuilderIdValue: BuilderId,
		Client:         state.Get("c").(*example.Client),
		Tenant:         state.Get("tenant").(string),
		Repo:           state.Get("export_repository").(string),
	}
	return artifact, nil
}
