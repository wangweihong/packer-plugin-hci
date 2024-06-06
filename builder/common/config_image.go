//go:generate packer-sdc struct-markdown

package common

import (
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

// ImageConfig is for common configuration related to creating Images.
type ImageConfig struct {
	// The name of the packer image.
	ImageName string `mapstructure:"image_name" required:"true"`
	// The description of the packer image.
	ImageDescription string `mapstructure:"image_description" required:"false"`
	// The type of the packer image. Available values include:
	// raw,qcow2, 默认为qcow2
	ImageType string `mapstructure:"image_type" required:"false"`
	// 导出的镜像仓库, 不指定使用原仓库
	ExportRepository string `mapstructure:"export_repository" required:"false"`
}

func (c *ImageConfig) Prepare(ctx *interpolate.Context) []error {
	errs := make([]error, 0)
	if c.ImageName == "" {
		errs = append(errs, fmt.Errorf("image_name must be specified"))
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
