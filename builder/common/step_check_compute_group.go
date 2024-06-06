package common

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/wangweihong/packer-plugin-hci/hci/example"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"
)

// This is a definition of a builder step and should implement multistep.Step
type StepLoadComputeGroup struct {
}

// Run should execute the purpose of this step
func (s *StepLoadComputeGroup) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)
	c := state.Get("c").(*example.Client)
	cluster := state.Get("cluster").(string)
	tenant := state.Get("tenant").(string)

	computeGroup, err := getDefaultComputeGroup(c, cluster, tenant)
	if err != nil {
		err = fmt.Errorf(" default resource pool is not exist or available")
		state.Put("error", err)
		return multistep.ActionHalt
	}
	state.Put("computeGroup", computeGroup)
	ui.Message(fmt.Sprintf("find default compute group %v", computeGroup))

	return multistep.ActionContinue
}

// Cleanup can be used to clean up any artifact created by the step.
// A step's clean up always run at the end of a build, regardless of whether provisioning succeeds or fails.
func (s *StepLoadComputeGroup) Cleanup(_ multistep.StateBag) {
	// Nothing to clean
}

func getDefaultComputeGroup(client *example.Client, cluster string, tenant string) (string, error) {
	response, err := example.NewCompute(client).ComputeGroupList(context.Background(), &iexample.ComputeGroupListRequest{
		Cluster:    cluster,
		Tenant:     tenant,
		FilterName: "default",
	})
	if err != nil {
		return "", fmt.Errorf("error getting compute group, err=%s", err)
	}

	for _, v := range response.Data.List {
		if v.Name == "default" {
			return v.UUID, nil
		}
	}

	return "", fmt.Errorf("cannot find default compute group")
}
