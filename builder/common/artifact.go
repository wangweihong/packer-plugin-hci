// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package common

import (
	"log"
	"strings"

	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/wangweihong/packer-plugin-hci/hci/example"
)

const (
	BuilderId = "wangweihong.hci"
)

// packersdk.Artifact implementation
type Artifact struct {
	Cluster string
	Tenant  string
	Repo    string
	// ImageId of built image
	ImageId string

	// BuilderIdValue is the unique ID for the builder that created this image
	BuilderIdValue string

	//  client for performing API stuff.
	Client *example.Client
}

func (a *Artifact) BuilderId() string {
	return a.BuilderIdValue
}

func (a *Artifact) Files() []string {
	return []string{}
}

func (*Artifact) Id() string {
	return ""
}

func (a *Artifact) String() string {
	return ""
}

func (a *Artifact) State(name string) interface{} {
	return nil
}

func (a *Artifact) Destroy() error {
	errors := make([]error, 0)
	log.Printf("Destroying image: %s", a.ImageId)

	for _, id := range strings.Split(a.ImageId, ";") {
		if id == "" {
			continue
		}

		errors := make([]error, 0)
		log.Printf("Destroying image: %s", a.ImageId)

		for _, id := range strings.Split(a.ImageId, ";") {
			if id == "" {
				continue
			}
			if err := DeleteImage(a.Client, a.Cluster, a.Tenant, a.Repo, a.ImageId); err != nil {
				errors = append(errors, err)
				continue
			}
		}

		if len(errors) > 0 {
			if len(errors) == 1 {
				return errors[0]
			} else {
				return &packer.MultiError{Errors: errors}
			}
		}

		return nil
	}

	if len(errors) > 0 {
		if len(errors) == 1 {
			return errors[0]
		} else {
			return &packer.MultiError{Errors: errors}
		}
	}

	return nil
}
