package common

import (
	"context"
	"fmt"
	"net"

	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"

	vnc "github.com/mitchellh/go-vnc"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
)

// This is a definition of a builder step and should implement multistep.Step
type StepVncConnect struct {
}

// Run should execute the purpose of this step
func (s *StepVncConnect) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)

	ui.Say("Connecting to VNC...")
	c, err := s.ConnectVNC(state)
	if err != nil {
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	state.Put("vnc_conn", c)
	ui.Say("Connected VNC...")

	return multistep.ActionContinue
}

func (s *StepVncConnect) ConnectVNC(state multistep.StateBag) (*vnc.ClientConn, error) {
	vncIp := state.Get("vnc_ip").(string)
	vncPort := state.Get("vnc_port").(int)
	vncPassword := state.Get("vnc_password")

	nc, err := net.Dial("tcp", fmt.Sprintf("%s:%d", vncIp, vncPort))
	if err != nil {
		err := fmt.Errorf("Error connecting to VNC: %s", err)
		state.Put("error", err)
		return nil, err
	}

	auth := []vnc.ClientAuth{new(vnc.ClientAuthNone)}
	if vncPassword != nil && len(vncPassword.(string)) > 0 {
		auth = []vnc.ClientAuth{&vnc.PasswordAuth{Password: vncPassword.(string)}}
	}

	c, err := vnc.Client(nc, &vnc.ClientConfig{Auth: auth, Exclusive: true})
	if err != nil {
		err := fmt.Errorf("Error handshaking with VNC: %s", err)
		state.Put("error", err)
		return nil, err
	}
	return c, nil
}

// Cleanup can be used to clean up any artifact created by the step.
// A step's clean up always run at the end of a build, regardless of whether provisioning succeeds or fails.
func (s *StepVncConnect) Cleanup(_ multistep.StateBag) {
	// Nothing to clean
}
