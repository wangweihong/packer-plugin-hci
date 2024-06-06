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
type StepLoadSpec struct {
	Spec string
}

// Run should execute the purpose of this step
func (s *StepLoadSpec) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)
	c := state.Get("c").(*example.Client)
	tenant := state.Get("tenant").(string)

	spec, err := getSpec(c, tenant, s.Spec)
	if err != nil {
		err = fmt.Errorf("the specified target spec %s is not exist or available,err:%v", s.Spec, err)
		state.Put("error", err)
		return multistep.ActionHalt
	}
	state.Put("spec", spec)
	ui.Message(fmt.Sprintf("the specified spec %s is available", s.Spec))

	return multistep.ActionContinue
}

// Cleanup can be used to clean up any artifact created by the step.
// A step's clean up always run at the end of a build, regardless of whether provisioning succeeds or fails.
func (s *StepLoadSpec) Cleanup(_ multistep.StateBag) {
	// Nothing to clean
}

func getSpec(client *example.Client, tenant string, name string) (*iexample.SpecEntry, error) {
	response, err := example.NewSpecification(client).SpecificationGet(context.Background(), &iexample.SpecGetRequest{
		Tenant: tenant,
		Name:   name,
	})
	if err != nil {
		return nil, fmt.Errorf("error getting tenant, err=%s", err)
	}

	return response.Data, nil
}
