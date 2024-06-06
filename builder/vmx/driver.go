package vmx

import (
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	commondriver "github.com/wangweihong/packer-plugin-hci/builder/common"
)

type Driver struct {
	BoundIP string
}

func NewDriver(runConfig RunConfig) (*Driver, error) {
	boundIP, err := commondriver.FindAvailableIP(runConfig.InstanceCIDR)
	if err != nil {
		return nil, err
	}

	return &Driver{
		BoundIP: boundIP,
	}, nil
}

func (d *Driver) HostIP(state multistep.StateBag) (string, error) {
	VmIP, ok := state.GetOk("vm_ip")
	if ok {
		return VmIP.(string), nil
	}
	return d.BoundIP, nil
}
