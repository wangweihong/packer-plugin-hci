//go:generate packer-sdc struct-markdown

package common

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
	"github.com/wangweihong/packer-plugin-hci/hci/example"
)

// AccessConfig is for common configuration related to HCI access
type AccessConfig struct {
	// The access key of the HCI to use.
	// If omitted, the HCI_ENDPOINT environment variable is used.
	Endpoint string `mapstructure:"endpoint"  required:"true" `
	// The name of tenant in which to launch the server to create the image.
	// If omitted, the HCI_TENANT environment variable is used.
	Tenant string `mapstructure:"tenant" required:"true"`
	// The name of tenant in which to launch the server to create the image.
	// If omitted, the HCI_USER environment variable is used.
	User string `mapstructure:"user" required:"true"`
	// The name of tenant in which to launch the server to create the image.
	// If omitted, the HCI_PASSWORD environment variable is used.
	Password string `mapstructure:"password" required:"true"`
}

func (c *AccessConfig) Prepare(ctx *interpolate.Context) []error {
	if c.Endpoint == "" {
		c.Endpoint = os.Getenv("HCI_ENDPOINT")
	}
	if c.Tenant == "" {
		c.Tenant = os.Getenv("HCI_TENANT")
	}
	if c.User == "" {
		c.User = os.Getenv("HCI_USER")
	}
	if c.Password == "" {
		c.Password = os.Getenv("HCI_PASSWORD")
	}
	// access parameters validation
	if c.Endpoint == "" || c.User == "" || c.Password == "" || c.Tenant == "" {
		paraErr := fmt.Errorf("endpoint, tennat, user and password  must be set")
		return []error{paraErr}
	}

	return nil
}

// HCIClient is the hci common client
func (c *AccessConfig) HCIClient() (*example.Client, error) {
	cc, err := example.NewClient(c.Endpoint, c.Tenant, c.User, c.Password)
	if err != nil {
		return nil, err
	}

	return cc, nil
}
