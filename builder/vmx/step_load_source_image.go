package vmx

import (
	"context"
	"fmt"

	"github.com/wangweihong/packer-plugin-hci/builder/common"

	"github.com/wangweihong/packer-plugin-hci/hci/example"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

// This is a definition of a builder step and should implement multistep.Step
type StepLoadSourceImage struct {
	Repository  string
	SourceImage string
}

// Run should execute the purpose of this step
func (s *StepLoadSourceImage) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)
	c := state.Get("c").(*example.Client)
	tenant := state.Get("tenant").(string)
	cluster := state.Get("cluster").(string)

	repo, err := common.GetImageRepository(c, cluster, tenant, s.Repository)
	if err != nil {
		err = fmt.Errorf("the specified target repository %s is not exist or available, err:%v", s.Repository, err)
		state.Put("error", err)
		return multistep.ActionHalt
	}
	state.Put("repository", repo)
	ui.Message(fmt.Sprintf("the specified repository %s is available,uuid:%s", s.Repository, repo))

	sourceImage, err := getSourceImage(c, cluster, tenant, repo, s.SourceImage)
	if err != nil {
		err = fmt.Errorf("the specified target iso %s is not exist or available: %v", s.SourceImage, err)
		state.Put("error", err)
		return multistep.ActionHalt
	}
	state.Put("source_image", sourceImage)
	state.Put("source_image_name", s.SourceImage)

	ui.Message(fmt.Sprintf("the specified source_image %s is available, uuid:%s", s.SourceImage, sourceImage))

	return multistep.ActionContinue
}

// Cleanup can be used to clean up any artifact created by the step.
// A step's clean up always run at the end of a build, regardless of whether provisioning succeeds or fails.
func (s *StepLoadSourceImage) Cleanup(_ multistep.StateBag) {
	// Nothing to clean
}

func getSourceImage(client *example.Client, cluster string, tenant string, repository, name string) (string, error) {
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
