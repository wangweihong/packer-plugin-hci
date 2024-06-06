//go:generate packer-sdc struct-markdown
//go:generate packer-sdc mapstructure-to-hcl2 -type Config
package vmx

import (
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	commondriver "github.com/wangweihong/packer-plugin-hci/builder/common"

	"github.com/hashicorp/packer-plugin-sdk/bootcommand"
	"github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/multistep/commonsteps"
	"github.com/hashicorp/packer-plugin-sdk/shutdowncommand"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	// packer将根据该配置来启动一个http服务, 用于提供boot command自定义配置
	commonsteps.HTTPConfig         `mapstructure:",squash"`
	commonsteps.ISOConfig          `mapstructure:",squash"`
	commonsteps.FloppyConfig       `mapstructure:",squash"`
	commonsteps.CDConfig           `mapstructure:",squash"`
	bootcommand.VNCConfig          `mapstructure:",squash"`
	shutdowncommand.ShutdownConfig `mapstructure:",squash"`
	commondriver.SSHConfig         `mapstructure:",squash"`
	commondriver.AccessConfig      `mapstructure:",squash"`
	commondriver.ImageConfig       `mapstructure:",squash"`
	RunConfig                      `mapstructure:",squash"`

	ctx interpolate.Context
}

func (c *Config) Prepare(raws ...interface{}) ([]string, error) {
	if err := config.Decode(c, &config.DecodeOpts{
		Interpolate:        true,
		InterpolateContext: &c.ctx,
		InterpolateFilter: &interpolate.RenderFilter{
			Exclude: []string{
				"boot_command",
			},
		},
	}, raws...); err != nil {
		return nil, err
	}

	// Accumulate any errors
	var warnings []string
	var errs *packer.MultiError
	//TODO: 后续需要支持。
	//isoWarnings, isoErrs := c.ISOConfig.Prepare(&c.ctx)
	//warnings = append(warnings, isoWarnings...)
	//errs = packer.MultiErrorAppend(errs, isoErrs...)
	errs = packer.MultiErrorAppend(errs, c.HTTPConfig.Prepare(&c.ctx)...)
	//errs = packer.MultiErrorAppend(errs, c.HTTPConfig.Prepare(&c.ctx)...)
	errs = packer.MultiErrorAppend(errs, c.AccessConfig.Prepare(&c.ctx)...)
	errs = packer.MultiErrorAppend(errs, c.ImageConfig.Prepare(&c.ctx)...)
	errs = packer.MultiErrorAppend(errs, c.VNCConfig.Prepare(&c.ctx)...)
	errs = packer.MultiErrorAppend(errs, c.RunConfig.Prepare(&c.ctx)...)
	errs = packer.MultiErrorAppend(errs, c.SSHConfig.Prepare(&c.ctx)...)

	if errs != nil && len(errs.Errors) > 0 {
		return nil, errs
	}

	if c.InstanceName == "" {
		c.InstanceName = fmt.Sprintf("packer-%s", c.PackerBuildName)
	}

	if c.DiskSize == 0 {
		c.DiskSize = 107374182400
	}

	if c.ImageType == "" {
		c.ImageType = "qcow2"
	}

	if c.ExportRepository == "" {
		c.ExportRepository = c.RepositoryName
	}

	packer.LogSecretFilter.Set(c.User, c.Password)

	// Warnings
	if c.ShutdownCommand == "" {
		warnings = append(warnings,
			"A shutdown_command was not specified. Without a shutdown command, Packer\n"+
				"will forcibly halt the virtual machine, which may result in data loss.")
	}

	if errs != nil && len(errs.Errors) > 0 {
		return warnings, errs
	}
	return nil, nil
}
