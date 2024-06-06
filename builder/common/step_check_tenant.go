// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package common

import (
	"context"
	"fmt"

	"github.com/wangweihong/gotoolbox/pkg/typeutil"
	"github.com/wangweihong/packer-plugin-hci/hci/example"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

// This is a definition of a builder step and should implement multistep.Step
type StepLoadTenant struct {
	TargetTenant string
}

// Run should execute the purpose of this step
func (s *StepLoadTenant) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)
	c := state.Get("c").(*example.Client)

	tenant, err := GetTenant(c, s.TargetTenant)
	if err != nil {
		err = fmt.Errorf("the specified target tenant %s is not exist or available", s.TargetTenant)
		state.Put("error", err)
		return multistep.ActionHalt
	}
	state.Put("tenant", tenant)
	ui.Message(fmt.Sprintf("the specified target tenant %s is available", s.TargetTenant))

	return multistep.ActionContinue
}

// Cleanup can be used to clean up any artifact created by the step.
// A step's clean up always run at the end of a build, regardless of whether provisioning succeeds or fails.
func (s *StepLoadTenant) Cleanup(_ multistep.StateBag) {
	// Nothing to clean
}

func GetTenant(client *example.Client, name string) (string, error) {
	response, err := example.NewOperationCenter(client).TenantGet(context.Background(), &iexample.TenantGetRequest{
		TenantName: typeutil.String(name),
	})
	if err != nil {
		return "", fmt.Errorf("error getting tenant, err=%s", err)
	}

	return response.Data.UUID, nil
}
