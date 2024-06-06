package iso

import (
	"context"
	"fmt"
	"time"

	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
	vnc "github.com/mitchellh/go-vnc"

	"github.com/hashicorp/packer-plugin-sdk/bootcommand"
	"github.com/hashicorp/packer-plugin-sdk/communicator"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
)

// This is a definition of a builder step and should implement multistep.Step
type StepVncBootCommand struct {
	Config bootcommand.VNCConfig
	Comm   *communicator.Config
	VMName string
	Ctx    interpolate.Context
}

type VNCBootCommandTemplateData struct {
	HTTPIP       string
	HTTPPort     int
	Name         string
	SSHPublicKey string
}

// Run should execute the purpose of this step
func (s *StepVncBootCommand) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	debug := state.Get("debug").(bool)
	httpPort := state.Get("http_port").(int)
	hostIP := state.Get("http_ip").(string)

	ui := state.Get("ui").(packersdk.Ui)
	conn := state.Get("vnc_conn").(*vnc.ClientConn)
	defer conn.Close()

	// Wait the for the vm to boot.
	if int64(s.Config.BootWait) > 0 {
		ui.Say(fmt.Sprintf("Waiting %s for boot...", s.Config.BootWait.String()))
		select {
		case <-time.After(s.Config.BootWait):
			break
		case <-ctx.Done():
			return multistep.ActionHalt
		}
	}

	var pauseFn multistep.DebugPauseFn
	if debug {
		pauseFn = state.Get("pauseFn").(multistep.DebugPauseFn)
	}

	s.Ctx.Data = &VNCBootCommandTemplateData{
		HTTPIP:       hostIP,
		HTTPPort:     httpPort,
		Name:         s.VMName,
		SSHPublicKey: string(s.Comm.SSHPublicKey),
	}
	d := bootcommand.NewVNCDriver(conn, s.Config.BootKeyInterval)

	ui.Say("Typing the boot command over VNC...")
	flatBootCommand := s.Config.FlatBootCommand()
	// 渲染模板
	command, err := interpolate.Render(flatBootCommand, &s.Ctx)
	if err != nil {
		err := fmt.Errorf("Error preparing boot command: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	ui.Say(command)

	seq, err := bootcommand.GenerateExpressionSequence(command)
	if err != nil {
		err := fmt.Errorf("error generating boot command: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	if err := seq.Do(ctx, d); err != nil {
		err := fmt.Errorf("error running boot command: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	if pauseFn != nil {
		pauseFn(multistep.DebugLocationAfterRun,
			fmt.Sprintf("boot_command: %s", command), state)
	}

	return multistep.ActionContinue
}

// Cleanup can be used to clean up any artifact created by the step.
// A step's clean up always run at the end of a build, regardless of whether provisioning succeeds or fails.
func (s *StepVncBootCommand) Cleanup(_ multistep.StateBag) {
	// Nothing to clean
}
