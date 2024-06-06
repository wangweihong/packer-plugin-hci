package example_test

import (
	"context"
	"os"
	"testing"

	"github.com/wangweihong/gotoolbox/pkg/typeutil"

	"github.com/wangweihong/packer-plugin-hci/hci/example"

	"github.com/wangweihong/gotoolbox/pkg/json"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEcsCreate(t *testing.T) {
	// change log default logger level to debug
	Convey("", t, func() {
		Convey("最小参数创建", func() {
			os.Setenv("HTTPCLI_DEBUG", "1")
			os.Setenv("HTTPCLI_DEBUG_HUGE", "1")
			resp, err := example.NewCompute(c).EcsCreate(context.Background(), &iexample.EcsCreateRequest{
				EcsCPU: iexample.EcsCPU{
					CPUnr:          1,
					MaxCPUnr:       16,
					CpuSocketCores: typeutil.Int(16),
					CpuSockets:     typeutil.Int(1),
				},
				EcsMem: iexample.EcsMem{
					MemorySize: 1073741824,
					MemoryMax:  549755813888,
				},
				Meta: iexample.EcsMeta{
					ComputeName: "test-123",
					// enable vnc
					VncEnable: typeutil.Bool(true),
				},
				//	Interfaces:  network,
				Disks: []iexample.EcsDisk{
					{
						Disk: iexample.Disk{
							DiskTargetBus: "scsi",
						},
						DiskUnifiedVsd: &iexample.DiskUnifiedVsdSource{
							VolumeCreateAttribute: iexample.VolumeCreateAttribute{
								VolumeName: "test-123",
								Pool:       pool,
								Capacity:   107374182400,
								Attribute: map[string]string{
									//"ReadIOPSLimit":     "0",
									//"WriteIOPSLimit":    "0",
									//"ReadBytesLimit":    "0",
									//"WriteBytesLimit":   "0",
									//"Encrypto":          "off",
									//"ThinProvision":     "on",
									//"VmCache":           "default",
									//"ScheduleOption":    "aggregate",
									//"VsdVendor":         "LIBTARGET",
									//"ChecksumHeader":    "off",
									//"ChecksumData":      "off",
									//"Sector":            "512",
									//"AttrComponentType": "chain",
									//"MultiPath":         "off",
									//"CryptoType":        "aes",
									//"CryptoMode":        "ecb",
									//"ComponentShift":    "30",
									"DataDevType": "cacheGroup",
									"DriveType":   "HDD",
									"replica":     "1",
									"DevType":     "target",
									"Zone":        "8dbde5bf-d9e2-472c-8c47-f58deeca3a6c",
									"Safety":      "first",
								},
								SnapCapacity:    typeutil.Uint64(322122547200),
								RebuildPriority: 0,
							},
						},
					}},
				//	CDRom:       cdrom,
				ClusterUUID: cluster,
				Tenant:      container,
				CreateNum:   typeutil.Int(1),
			})
			So(err, ShouldBeNil)
			json.PrintStructObject(resp)
		})

	})
}

func TestEcsInspect(t *testing.T) {
	// change log default logger level to debug
	Convey("", t, func() {
		Convey("最小参数创建", func() {
			os.Setenv("HTTPCLI_DEBUG", "1")
			os.Setenv("HTTPCLI_DEBUG_HUGE", "1")
			resp, err := example.NewCompute(c).EcsInspect(context.Background(), &iexample.EcsInspectRequest{
				Cluster:       cluster,
				Tenant:        container,
				ECS:           ecs,
				ShowPlaintext: true,
			})
			So(err, ShouldBeNil)
			json.PrintStructObject(resp)
		})

	})
}

func TestEcsImport(t *testing.T) {
	// change log default logger level to debug
	Convey("", t, func() {
		Convey("最小参数创建", func() {
			os.Setenv("HTTPCLI_DEBUG", "1")
			os.Setenv("HTTPCLI_DEBUG_HUGE", "1")
			resp, err := example.NewCompute(c).EcsImport(context.Background(), &iexample.EcsImportRequest{
				EcsCPU: iexample.EcsCPU{
					CPUnr:          1,
					MaxCPUnr:       16,
					CpuSocketCores: typeutil.Int(16),
					CpuSockets:     typeutil.Int(1),
				},
				EcsMem: iexample.EcsMem{
					MemorySize: 1073741824,
					MemoryMax:  549755813888,
				},
				Meta: iexample.EcsMeta{
					ComputeName: "import-test-123",
					// enable vnc
					VncEnable: typeutil.Bool(true),
				},
				//	Interfaces:  network,
				Disks: []iexample.EcsDisk{
					{
						RepositoryUUID: repo,
						ImageUUID:      image,
						Disk: iexample.Disk{
							DiskTargetBus: "scsi",
						},
						DiskUnifiedVsd: &iexample.DiskUnifiedVsdSource{
							VolumeCreateAttribute: iexample.VolumeCreateAttribute{
								VolumeName: "test-123",
								Pool:       pool,
								Capacity:   107374182400,
								Attribute: map[string]string{
									//"ReadIOPSLimit":     "0",
									//"WriteIOPSLimit":    "0",
									//"ReadBytesLimit":    "0",
									//"WriteBytesLimit":   "0",
									//"Encrypto":          "off",
									//"ThinProvision":     "on",
									//"VmCache":           "default",
									//"ScheduleOption":    "aggregate",
									//"VsdVendor":         "LIBTARGET",
									//"ChecksumHeader":    "off",
									//"ChecksumData":      "off",
									//"Sector":            "512",
									//"AttrComponentType": "chain",
									//"MultiPath":         "off",
									//"CryptoType":        "aes",
									//"CryptoMode":        "ecb",
									//"ComponentShift":    "30",
									"DataDevType": "cacheGroup",
									"DriveType":   "HDD",
									"replica":     "1",
									"DevType":     "target",
									"Zone":        "8dbde5bf-d9e2-472c-8c47-f58deeca3a6c",
									"Safety":      "first",
								},
								SnapCapacity:    typeutil.Uint64(322122547200),
								RebuildPriority: 0,
							},
						},
					}},
				//	CDRom:       cdrom,
				ClusterUUID: cluster,
				Tenant:      container,
				CreateNum:   typeutil.Int(1),
			})
			So(err, ShouldBeNil)
			json.PrintStructObject(resp)
		})

	})
}

func TestEcsImportProgress(t *testing.T) {
	// change log default logger level to debug
	Convey("", t, func() {
		Convey("最小参数创建", func() {
			os.Setenv("HTTPCLI_DEBUG", "1")
			os.Setenv("HTTPCLI_DEBUG_HUGE", "1")
			resp, err := example.NewCompute(c).EcsDiskExportProgress(context.Background(), &iexample.EcsDiskExportProgressRequest{
				Cluster: cluster,
				Tenant:  container,
				ECS:     ecs,
			})
			So(err, ShouldBeNil)
			json.PrintStructObject(resp)
		})

	})
}
