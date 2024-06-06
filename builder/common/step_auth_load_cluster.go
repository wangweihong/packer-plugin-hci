// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package common

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

// This is a definition of a builder step and should implement multistep.Step
type StepAuth struct {
}

// Run should execute the purpose of this step
func (s *StepAuth) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)
	config := state.Get("access_config").(*AccessConfig)

	c, err := config.HCIClient()
	if err != nil {
		err = fmt.Errorf("error initializing client: %s", err)
		state.Put("error", err)
		return multistep.ActionHalt
	}
	ui.Say("initializing client success.")
	state.Put("c", c)
	state.Put("cluster", c.GetCluster())
	return multistep.ActionContinue
}

// Cleanup can be used to clean up any artifact created by the step.
// A step's clean up always run at the end of a build, regardless of whether provisioning succeeds or fails.
func (s *StepAuth) Cleanup(_ multistep.StateBag) {
	// Nothing to clean
}
