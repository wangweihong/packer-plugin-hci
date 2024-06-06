package example_test

import (
	"context"
	"os"
	"testing"

	"github.com/wangweihong/packer-plugin-hci/hci/example"

	"github.com/wangweihong/gotoolbox/pkg/json"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRepositoryGet(t *testing.T) {
	// change log default logger level to debug
	Convey("", t, func() {
		Convey("tenant", func() {
			resp, err := example.NewRepository(c).RepositoryGet(context.Background(), &iexample.RepositoryGetRequest{
				Cluster: cluster,
				Name:    "container",
			})
			So(err, ShouldBeNil)
			json.PrintStructObject(resp)
		})

	})
}

func TestRepositoryImageGet(t *testing.T) {
	// change log default logger level to debug
	Convey("", t, func() {
		Convey("tenant", func() {
			os.Setenv("HTTPCLI_DEBUG", "1")
			os.Setenv("HTTPCLI_DEBUG_HUGE", "1")
			resp, err := example.NewRepository(c).RepositoryImageGet(context.Background(), &iexample.RepositoryImageGetRequest{
				Cluster:    cluster,
				Name:       "ubuntu-16.04.6-server-amd64.iso",
				Repository: repo,
				Tenant:     container,
			})
			So(err, ShouldBeNil)
			json.PrintStructObject(resp)
		})

	})
}
