package iso

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
)

// Step to discover the current host ip
// To make sure the IP is set before boot command and http server steps
type StepHTTPIPDiscover struct{}

func (s *StepHTTPIPDiscover) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)

	// Determine the host IP
	hostIPs, err := GetIPAddrs(false)
	if err != nil {
		err := fmt.Errorf("Error detecting host IP: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	if len(hostIPs) == 0 {
		err := fmt.Errorf("Error find host IP")
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	log.Printf("Host IP for current node: %s", hostIPs[0])
	state.Put("http_ip", hostIPs[0])

	return multistep.ActionContinue
}

func (*StepHTTPIPDiscover) Cleanup(multistep.StateBag) {}

func GetIPAddrs(wantIpv6 bool) ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	ips := make([]string, 0)
	for _, iface := range ifaces {
		if strings.HasPrefix(iface.Name, "e") ||
			strings.HasPrefix(iface.Name, "br") {
			addrs, err := iface.Addrs()
			if err != nil {
				continue
			}

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok {
					if p4 := ipnet.IP.To4(); len(p4) == net.IPv4len {
						ips = append(ips, ipnet.IP.String())
					} else if len(ipnet.IP) == net.IPv6len && wantIpv6 {
						ips = append(ips, ipnet.IP.String())
					}
				}
			}
		}
	}

	return ips, nil
}
