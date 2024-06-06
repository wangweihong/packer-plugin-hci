//go:generate packer-sdc struct-markdown

package vmx

import (
	"errors"
	"fmt"

	"github.com/wangweihong/gotoolbox/pkg/netutil"

	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

// RunConfig contains configuration for running an instance from a source image
// and details on how to access that launched image.
type RunConfig struct {
	// The name for the desired specfication for the server to be created.
	Specification string `mapstructure:"specification" required:"true"`
	// Name that is applied to the server instance created by Packer. If this
	// isn't specified, the default is same as image_name.
	InstanceName string `mapstructure:"instance_name" required:"false"`
	TargetTenant string `mapstructure:"target_tenant" required:"true"`

	RepositoryName string `mapstructure:"repository_name" required:"true"`
	SourceImage    string `mapstructure:"source_image" required:"true"`
	DiskSize       uint64 `mapstructure:"disk_size" required:"true"`

	InstanceCIDR string `mapstructure:"instance_cidr" required:"true"`
}

func (c *RunConfig) Prepare(ctx *interpolate.Context) []error {
	// Validation
	var errs []error
	if c.Specification == "" {
		errs = append(errs, errors.New("A Specification must be specified"))
	}

	if c.TargetTenant == "" {
		errs = append(errs, errors.New("target tenant must be specified"))
	}

	if c.RepositoryName == "" {
		errs = append(errs, fmt.Errorf("repository_name must be specified"))
	}

	if c.SourceImage == "" {
		errs = append(errs, fmt.Errorf("source_image must be specified"))
	}

	if c.DiskSize == 0 {
		errs = append(errs, fmt.Errorf("disk_size must be specified"))
	}

	if c.InstanceCIDR == "" {
		errs = append(errs, fmt.Errorf("instance_cidr must be specified"))
	}

	if _, err := netutil.ValidateCIDR(c.InstanceCIDR); err != nil {
		errs = append(errs, fmt.Errorf("instance_cidr %v validate err:%v", c.InstanceCIDR, err.Error()))
	}

	return errs
}
