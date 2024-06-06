// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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
type StepLoadResourcePool struct {
}

// Run should execute the purpose of this step
func (s *StepLoadResourcePool) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)
	c := state.Get("c").(*example.Client)
	cluster := state.Get("cluster").(string)

	resourcePool, err := getDefaultResourcePool(c, cluster)
	if err != nil {
		err = fmt.Errorf(" default resource pool is not exist or available")
		state.Put("error", err)
		return multistep.ActionHalt
	}
	state.Put("resourcePool", resourcePool)
	ui.Message(fmt.Sprintf("the default pool is available"))

	return multistep.ActionContinue
}

// Cleanup can be used to clean up any artifact created by the step.
// A step's clean up always run at the end of a build, regardless of whether provisioning succeeds or fails.
func (s *StepLoadResourcePool) Cleanup(_ multistep.StateBag) {
	// Nothing to clean
}

func getDefaultResourcePool(client *example.Client, cluster string) (string, error) {
	response, err := example.NewResourcePool(client).ResourcePoolList(context.Background(), &iexample.ResourcePoolListRequest{
		Cluster:    cluster,
		FilterName: "default",
	})
	if err != nil {
		return "", fmt.Errorf("error getting tenant, err=%s", err)
	}

	for _, v := range response.Data.List {
		if v.Name == "default" {
			return v.UUID, nil
		}
	}

	return "", fmt.Errorf("cannot find default resource pool")
}
