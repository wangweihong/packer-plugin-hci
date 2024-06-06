package common

import (
	"context"
	"fmt"

	"github.com/wangweihong/packer-plugin-hci/hci/example"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
)

type StepLoadNetwork struct {
}

// Run should execute the purpose of this step
func (s *StepLoadNetwork) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	//ui := state.Get("ui").(packersdk.Ui)
	c := state.Get("c").(*example.Client)
	tenant := state.Get("tenant").(string)
	cluster := state.Get("cluster").(string)

	portGroup, err := getDefaultVirtualSwitchAndPort(c, cluster, tenant)
	if err != nil {
		err = fmt.Errorf("the specified target default port group is not exist or available")
		state.Put("error", err)
		return multistep.ActionHalt
	}
	state.Put("port_group", portGroup)
	return multistep.ActionContinue
}

func getDefaultVirtualSwitchAndPort(client *example.Client, cluster string, tenant string) (string, error) {
	vSwitch, err := example.NewNetwork(client).VirtualSwitchGet(context.Background(), &iexample.VirtualSwitchGetRequest{
		Tenant: tenant,
		Name:   "manage",
	})
	if err != nil {
		return "", fmt.Errorf("error getting manage virtual switch, err=%s", err)
	}

	for _, port := range vSwitch.Data.PortGroupList {
		if port.PortGroupName == "default" {
			return port.PortGroupUUID, nil
		}
	}

	return "", fmt.Errorf("error getting default port group in manage virtual switch")

}
