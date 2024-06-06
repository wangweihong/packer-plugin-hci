package common

import (
	"context"
	"fmt"
	"log"
	"time"

	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/wangweihong/gotoolbox/pkg/wait"

	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/wangweihong/packer-plugin-hci/hci/example"
)

type StepCreateImage struct {
	Config ImageConfig
}

// Run should execute the purpose of this step
func (s *StepCreateImage) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)
	c := state.Get("c").(*example.Client)
	tenant := state.Get("tenant").(string)
	cluster := state.Get("cluster").(string)
	server := state.Get("server").(string)

	if err := ShutoffServer(c, cluster, tenant, server); err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	repo, err := GetImageRepository(c, cluster, tenant, s.Config.ExportRepository)
	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}
	state.Put("export_repository", repo)

	if err := exportDisk(c, cluster, tenant, server, s.Config.ImageType, s.Config.ImageName, repo); err != nil {
		err = fmt.Errorf("exportDisk err:%v", err)
		state.Put("error", err)
		return multistep.ActionHalt
	}
	log.Printf("[DEBUG] Waiting for disk export\n")

	if err := wait.Poll(10*time.Second, 10*time.Minute, func() (done bool, err error) {
		progress, err := exportDiskProgress(c, cluster, tenant, server)
		if err != nil {
			return false, err
		}

		log.Printf("[Trace]current progress:%v,Waiting 10s for next try\n", progress)
		if progress == 100 {
			return true, nil
		}
		return false, nil
	}); err != nil {
		state.Put("error", fmt.Errorf("wait for disk export eror:%v", err))
		return multistep.ActionHalt
	}

	if err := SyncImage(c, cluster, tenant, repo); err != nil {
		state.Put("error", fmt.Errorf("sync image eror:%v", err))
		return multistep.ActionHalt
	}

	image, err := GetImage(c, cluster, tenant, repo, s.Config.ImageName, s.Config.ImageType)
	if err != nil {
		state.Put("error", fmt.Errorf("get image eror:%v", err))
		return multistep.ActionHalt
	}

	state.Put("image", image)

	ui.Say("Export Image success")
	return multistep.ActionContinue
}

func (s *StepCreateImage) Cleanup(_ multistep.StateBag) {
	// Nothing to clean
}

func exportDisk(client *example.Client, cluster string, tenant string, ecs string, imageType string, imageName string, repo string) error {
	_, err := example.NewCompute(client).EcsDiskExport(context.Background(), &iexample.EcsDiskExportRequest{
		Cluster:  cluster,
		Tenant:   tenant,
		ECS:      ecs,
		Compress: false,
		Disk: []iexample.EcsDiskExportInfo{{
			Device:         "sda",
			SnapshotUUID:   "",
			ImageName:      imageName,
			ImageType:      imageType,
			RepositoryUUID: repo,
		}},
	})
	return err
}

func exportDiskProgress(client *example.Client, cluster string, tenant string, ecs string) (int, error) {
	resp, err := example.NewCompute(client).EcsDiskExportProgress(context.Background(), &iexample.EcsDiskExportProgressRequest{
		Cluster: cluster,
		Tenant:  tenant,
		ECS:     ecs,
	})
	if err != nil {
		return 0, err
	}
	return resp.Data.Progress, nil
}
