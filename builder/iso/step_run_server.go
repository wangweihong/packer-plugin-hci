package iso

import (
	"context"
	"fmt"

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
	//spec := state.Get("spec").(*iexample.SpecEntry)
	//
	interfaces := s.buildNetwork(state, s.InstanceCIDR)
	cdrom := s.buildCDROM(state)
	cpu := s.buildServerCPU(state)
	mem := s.buildServerMem(state)
	disk := s.buildRootVolume(state, s.Name, s.RootVolumeSize)

	if s.err != nil {
		state.Put("error", s.err)
		return multistep.ActionHalt
	}
	req := &iexample.EcsCreateRequest{
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
		Type:        typeutil.String("GUS"),
		Interfaces:  interfaces,
		Disks:       disk,
		CDRom:       cdrom,
		ClusterUUID: cluster,
		Tenant:      tenant,
	}
	ui.Message(fmt.Sprintf(" create ecs:%v with ip :%s", s.Name, req.Interfaces[0].BoundIP))

	resp, err := example.NewCompute(c).EcsCreate(context.Background(), req)
	if err != nil {
		state.Put("error", fmt.Errorf("careat server error:%v", err.Error()))
		return multistep.ActionHalt
	}

	if resp.Data.Fail != 0 {
		state.Put("error", fmt.Errorf("careat server error:%v", resp.Data.Results[0].Error()))

		return multistep.ActionHalt
	}
	ui.Message(fmt.Sprintf("success to create ecs,name:%v,uuid:%v", resp.Data.Results[0].Data.Name, resp.Data.Results[0].Data.UUID))
	state.Put("server", resp.Data.Results[0].Data.UUID)

	if err := s.startServer(state); err != nil {
		state.Put("error", fmt.Errorf("start server error:%v", err.Error()))
		return multistep.ActionHalt
	}
	ui.Message(fmt.Sprintf("run ecs:%v", s.Name))

	ecsInspect, err := s.inspectServer(state)
	if err != nil {
		state.Put("error", fmt.Errorf("inspect server error:%v", err.Error()))
		return multistep.ActionHalt
	}
	state.Put("host_ip", ecsInspect.Data.MachineManagerIP)
	state.Put("vnc_ip", ecsInspect.Data.MachineManagerIP)
	state.Put("vnc_port", ecsInspect.Data.VncPort)
	state.Put("vnc_password", ecsInspect.Data.VncPasswordPlaintext)

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

func (s *StepRunServer) buildCDROM(state multistep.StateBag) *iexample.EcsDiskSpecial {
	if s.err != nil {
		return nil
	}
	repository := state.Get("repository").(string)
	if repository == "" {
		s.err = fmt.Errorf("repository is missing")
		return nil
	}

	iso := state.Get("iso").(string)
	isoName := state.Get("iso_name").(string)
	if repository == "" {
		s.err = fmt.Errorf("repository image is missing")
		return nil
	}

	return &iexample.EcsDiskSpecial{
		RepositoryUUID: repository,
		ImageUUID:      iso,
		Name:           isoName,
	}
}

func (s *StepRunServer) buildNetwork(state multistep.StateBag, instanceCIDR string) []iexample.ECSInterface {
	if s.err != nil {
		return nil
	}
	c := state.Get("c").(*example.Client)
	tenant := state.Get("tenant").(string)
	cluster := state.Get("cluster").(string)
	boundIP := state.Get("instance_ip").(string)

	//portGroup := state.Get("port_group").(string)
	//if portGroup == "" {
	//	s.err = fmt.Errorf("port group is missing")
	//	return nil
	//}
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

	return []iexample.ECSInterface{{
		InterfaceModelType: "virtio",
		InterfaceType:      "manage",
		PortGroup:          resp.Data.PortGroupUUID,
		BoundIP:            boundIP,
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

func (s *StepRunServer) startServer(state multistep.StateBag) error {
	c := state.Get("c").(*example.Client)
	tenant := state.Get("tenant").(string)
	cluster := state.Get("cluster").(string)
	server := state.Get("server").(string)
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

func (s *StepRunServer) inspectServer(state multistep.StateBag) (*iexample.EcsInspectResponse, error) {
	c := state.Get("c").(*example.Client)
	tenant := state.Get("tenant").(string)
	cluster := state.Get("cluster").(string)
	server := state.Get("server").(string)
	resp, err := example.NewCompute(c).EcsInspect(context.Background(), &iexample.EcsInspectRequest{
		Cluster:       cluster,
		Tenant:        tenant,
		ECS:           server,
		ShowPlaintext: true,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
