package vmx

import (
	"context"
	"fmt"
	"log"
	"time"

	commondriver "github.com/wangweihong/packer-plugin-hci/builder/common"

	"github.com/wangweihong/gotoolbox/pkg/wait"

	"github.com/wangweihong/gotoolbox/pkg/async"

	"github.com/hashicorp/packer-plugin-sdk/packer"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/wangweihong/gotoolbox/pkg/typeutil"
	"github.com/wangweihong/packer-plugin-hci/hci/example"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
)

type StepRunServer struct {
	err            error
	Name           string
	RootVolumeSize uint64
	InstanceCIDR   string
}

// Run should execute the purpose of this step
func (s *StepRunServer) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)
	c := state.Get("c").(*example.Client)
	tenant := state.Get("tenant").(string)
	cluster := state.Get("cluster").(string)

	ui.Message(fmt.Sprintf(" create ecs:%v 	", s.Name))

	server, err := s.createServer(state)
	if err != nil {
		state.Put("error", fmt.Errorf("start server error:%v", err.Error()))
		return multistep.ActionHalt
	}
	ui.Message(fmt.Sprintf("success to create ecs,name:%v,uuid:%v", s.Name, server))
	state.Put("server", server)

	if err := s.startServer(c, cluster, tenant, server); err != nil {
		state.Put("error", fmt.Errorf("start server error:%v", err.Error()))
		return multistep.ActionHalt
	}
	ui.Message(fmt.Sprintf("run ecs:%v", s.Name))
	log.Printf("[DEBUG] Waiting for network interface ip\n")

	if err := wait.Poll(10*time.Second, 10*time.Minute, func() (done bool, err error) {
		log.Printf("[Trace] Waiting 10s for next try\n")

		if _, ok := state.GetOk(multistep.StateCancelled); ok {
			return false, fmt.Errorf("force cancel")
		}

		ecsInspect, err := commondriver.InspectServer(c, cluster, tenant, server)
		if err != nil {
			return false, fmt.Errorf("inspect server error:%v", err.Error())
		}

		if len(ecsInspect.Data.InterfaceIps) != 0 {
			state.Put("vm_ip", ecsInspect.Data.InterfaceIps[0].IP)
			ui.Message(fmt.Sprintf("discover ecs name:%v ip:%v", s.Name, ecsInspect.Data.InterfaceIps[0].IP))

			return true, nil
		}
		return false, nil

	}); err != nil {
		state.Put("error", fmt.Errorf("get server interface ips error:%v", err.Error()))
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

// Cleanup can be used to clean up any artifact created by the step.
// A step's clean up always run at the end of a build, regardless of whether provisioning succeeds or fails.
func (s *StepRunServer) Cleanup(state multistep.StateBag) {
	// Nothing to clean
	c := state.Get("c").(*example.Client)
	tenant := state.Get("tenant").(string)
	cluster := state.Get("cluster").(string)
	serverI := state.Get("server")
	server := ""
	if serverI != nil {
		server = serverI.(string)
	}
	if server != "" {
		_, err := example.NewCompute(c).EcsDelete(context.Background(), &iexample.EcsDeleteRequest{
			ClusterUUID:     cluster,
			Tenant:          tenant,
			ECS:             server,
			RealDelete:      true,
			IsDeleteVolumes: true,
		})
		if err != nil {
			state.Put("error", fmt.Errorf("delete server error:%v", err))
		}
	}
}

func (s *StepRunServer) buildNetwork(state multistep.StateBag, instanceCIDR string) []iexample.ECSInterface {
	if s.err != nil {
		return nil
	}
	c := state.Get("c").(*example.Client)
	tenant := state.Get("tenant").(string)
	cluster := state.Get("cluster").(string)
	//  如果拟机交换机没有开启DHCP, 则静态IP不生效

	resp, err := example.NewNetwork(c).VirtualSwitchPortGroupGet(context.Background(), &iexample.VirtualSwitchPortGroupGetRequest{
		Cluster:     cluster,
		Name:        "default",
		Tenant:      tenant,
		VSwitchName: "manage",
	})
	if err != nil {
		s.err = err
		return nil
	}

	if resp.Data.EnableDhcp {
		boundIP := state.Get("instance_ip").(string)
		return []iexample.ECSInterface{{
			InterfaceModelType: "virtio",
			InterfaceType:      "manage",
			PortGroup:          resp.Data.PortGroupUUID,
			BoundIP:            boundIP,
		}}
	}
	// 虚拟交换机端口组如果不支持dhcp, 则静态ip不生效
	return []iexample.ECSInterface{{
		InterfaceModelType: "virtio",
		InterfaceType:      "manage",
		PortGroup:          resp.Data.PortGroupUUID,
	}}
}

func (s *StepRunServer) buildRootVolume(state multistep.StateBag, vmName string, diskSize uint64) []iexample.EcsDisk {
	if s.err != nil {
		return nil
	}

	c := state.Get("c").(*example.Client)
	tenant := state.Get("tenant").(string)
	cluster := state.Get("cluster").(string)
	resourePool := state.Get("resourcePool").(string)
	repo := state.Get("repository").(string)
	image := state.Get("source_image").(string)

	// find first pool match desired size
	poolResp, err := example.NewVolume(c).PoolList(context.Background(), &iexample.PoolListRequest{
		PagingParam: iexample.PagingParam{
			PageNum:  0,
			PageSize: 1,
		},
		Cluster:             cluster,
		Tenant:              tenant,
		FilterAvailMoreThan: diskSize,
	})
	if err != nil {
		s.err = err
		return nil
	}

	if poolResp.Data.TotalCount == 0 {
		s.err = fmt.Errorf("no pool available size match desire:%vM", diskSize/1024/1024)
		return nil
	}

	return []iexample.EcsDisk{
		{
			Disk: iexample.Disk{
				DiskTargetBus: "scsi",
			},
			RepositoryUUID: repo,
			ImageUUID:      image,
			DiskUnifiedVsd: &iexample.DiskUnifiedVsdSource{
				VolumeCreateAttribute: iexample.VolumeCreateAttribute{
					VolumeName: vmName + "-volume1",
					Pool:       poolResp.Data.List[0].UUID,
					Capacity:   diskSize,
					Attribute: map[string]string{
						"ReadIOPSLimit":     "0",
						"WriteIOPSLimit":    "0",
						"ReadBytesLimit":    "0",
						"WriteBytesLimit":   "0",
						"Encrypto":          "off",
						"ThinProvision":     "on",
						"VmCache":           "default",
						"ScheduleOption":    "aggregate",
						"VsdVendor":         "LIBTARGET",
						"ChecksumHeader":    "off",
						"ChecksumData":      "off",
						"Sector":            "512",
						"AttrComponentType": "chain",
						"MultiPath":         "off",
						"CryptoType":        "aes",
						"CryptoMode":        "ecb",
						"ComponentShift":    "30",
						"DataDevType":       "cacheGroup",
						"DriveType":         "HDD",
						"replica":           "1",
						"DevType":           "target",
						"Zone":              resourePool,
						"Safety":            "first",
					},
					SnapCapacity:    typeutil.Uint64(3 * diskSize),
					RebuildPriority: 0,
				},
			},
		}}
}

func (s *StepRunServer) createServer(state multistep.StateBag) (string, error) {
	c := state.Get("c").(*example.Client)
	tenant := state.Get("tenant").(string)
	cluster := state.Get("cluster").(string)
	ui := state.Get("ui").(packersdk.Ui)

	interfaces := s.buildNetwork(state, s.InstanceCIDR)
	cpu := s.buildServerCPU(state)
	mem := s.buildServerMem(state)
	disk := s.buildRootVolume(state, s.Name, s.RootVolumeSize)

	if s.err != nil {
		return "", s.err
	}
	req := &iexample.EcsImportRequest{
		Baseline: typeutil.String("Haswell-noTSX"),
		EcsCPU:   *cpu,
		EcsMem:   *mem,
		Meta: iexample.EcsMeta{
			DomainCpuMode:         "custom",
			PanicModel:            "isa",
			DomainMouseInputType:  "usb",
			DomainBootMenuTimeOut: "0",
			OsLoaderPath:          "BIOS",
			ComputeName:           s.Name,
			// enable vnc
			VncEnable: typeutil.Bool(true),
			// 设置后才能使用网页控制台
			Video: []*iexample.EcsVideo{{
				Heads:     1,
				ModelType: "vga",
				RamSize:   32768,
				VGAMem:    8192,
			}},
		},
		Type:       typeutil.String("GUS"),
		Interfaces: interfaces,
		Disks:      disk,
		//CDRom:       cdrom,
		ClusterUUID: cluster,
		Tenant:      tenant,
	}

	resp, err := example.NewCompute(c).EcsImport(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("create server error:%v", err.Error())
	}

	if resp.Data.Fail != 0 {
		return "", fmt.Errorf("create server error:%v", resp.Data.Results[0].Error())
	}

	_, err = waitForImageImportJobSuccess(ui, state, c, cluster, tenant, resp.Data.Results[0].Data.VmUUID)
	if err != nil {
		return "", err
	}

	return resp.Data.Results[0].Data.VmUUID, nil
}

func (s *StepRunServer) buildServerCPU(state multistep.StateBag) *iexample.EcsCPU {
	if s.err != nil {
		return nil
	}
	spec := state.Get("spec").(*iexample.SpecEntry)

	return &spec.UnifiedVmSpec.EcsCPU
}

func (s *StepRunServer) buildServerMem(state multistep.StateBag) *iexample.EcsMem {
	if s.err != nil {
		return nil
	}
	spec := state.Get("spec").(*iexample.SpecEntry)

	return &spec.UnifiedVmSpec.EcsMem
}

func (s *StepRunServer) startServer(c *example.Client, cluster, tenant, server string) error {
	_, err := example.NewCompute(c).EcsStart(context.Background(), &iexample.EcsStartRequest{
		Cluster: cluster,
		Tenant:  tenant,
		ECS:     server,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *StepRunServer) inspectServer(c *example.Client, cluster, tenant, server string) (*iexample.EcsInspectResponse, error) {
	resp, err := example.NewCompute(c).EcsInspect(context.Background(), &iexample.EcsInspectRequest{
		Cluster: cluster,
		Tenant:  tenant,
		ECS:     server,
		// get vnc password
		ShowPlaintext: true,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func waitForImageImportJobSuccess(ui packer.Ui, state multistep.StateBag, client *example.Client, cluster, tenant, server string) (*iexample.EcsInspectResponse, error) {
	ui.Message("Waiting for import image to ECS success...")
	stateChange := async.StateChangeConf{
		Pending:      []string{"importing"},
		Target:       []string{"noaction"},
		Refresh:      importJobStateRefreshFunc(client, cluster, tenant, server),
		Timeout:      10 * time.Minute,
		Delay:        10 * time.Second,
		PollInterval: 10 * time.Second,
	}
	serverJob, err := stateChange.WaitForState()
	if err != nil {
		err = fmt.Errorf("error waiting for server (%s)  import image to become ready: %s", server, err)
		ui.Error(err.Error())
		return nil, err
	}

	return serverJob.(*iexample.EcsInspectResponse), nil
}

func importJobStateRefreshFunc(client *example.Client, cluster, tenant, server string) async.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, err := example.NewCompute(client).EcsInspect(context.Background(), &iexample.EcsInspectRequest{
			Cluster: cluster,
			Tenant:  tenant,
			ECS:     server,
		})
		if err != nil {
			return nil, "", err
		}

		if resp.Data.Action == "noaction" {
			return resp, resp.Data.Action, nil
		}

		if resp.Data.Action == "importing" {
			return resp, "importing", nil
		}

		return resp, "unknown", nil
	}
}
