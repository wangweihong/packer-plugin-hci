package iso

import (
	"context"
	"fmt"

	"github.com/wangweihong/packer-plugin-hci/hci/example"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

// This is a definition of a builder step and should implement multistep.Step
type StepLoadISO struct {
	Repository string
	ISO        string
}

// TODO: 后续支持从外部下载iso, 并推送到指定的镜像仓库。
// Run should execute the purpose of this step
func (s *StepLoadISO) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)
	c := state.Get("c").(*example.Client)
	tenant := state.Get("tenant").(string)
	cluster := state.Get("cluster").(string)

	repo, err := getISORepository(c, cluster, tenant, s.Repository)
	if err != nil {
		err = fmt.Errorf("the specified target repository %s is not exist or available, err:%v", s.Repository, err)
		state.Put("error", err)
		return multistep.ActionHalt
	}
	state.Put("repository", repo)
	ui.Message(fmt.Sprintf("the specified repository %s is available,uuid:%s", s.Repository, repo))

	iso, err := getISO(c, cluster, tenant, repo, s.ISO)
	if err != nil {
		err = fmt.Errorf("the specified target iso %s is not exist or available: %v", s.ISO, err)
		state.Put("error", err)
		return multistep.ActionHalt
	}
	state.Put("iso", iso)
	state.Put("iso_name", s.ISO)

	ui.Message(fmt.Sprintf("the specified iso %s is available, uuid:%s", s.ISO, iso))

	return multistep.ActionContinue
}

// Cleanup can be used to clean up any artifact created by the step.
// A step's clean up always run at the end of a build, regardless of whether provisioning succeeds or fails.
func (s *StepLoadISO) Cleanup(_ multistep.StateBag) {
	// Nothing to clean
}

func getISORepository(client *example.Client, cluster string, tenant string, name string) (string, error) {
	response, err := example.NewRepository(client).RepositoryList(context.Background(), &iexample.RepositoryListRequest{
		Cluster:     cluster,
		Tenant:      tenant,
		FilterUsage: "cdrom",
		FilterName:  name,
	})
	if err != nil {
		return "", fmt.Errorf("error getting repository, err=%s", err)
	}
	for _, v := range response.Data.List {
		if v.Name == name {
			return v.UUID, nil
		}
	}

	return "", fmt.Errorf("repository %v not found", name)
}

func getISO(client *example.Client, cluster string, tenant string, repository, name string) (string, error) {
	response, err := example.NewRepository(client).RepositoryImageGet(context.Background(), &iexample.RepositoryImageGetRequest{
		Cluster:    cluster,
		Tenant:     tenant,
		Name:       name,
		Repository: repository,
	})
	if err != nil {
		return "", fmt.Errorf("error getting iso, err=%s", err)
	}

	return response.Data.Image.UUID, nil
}
