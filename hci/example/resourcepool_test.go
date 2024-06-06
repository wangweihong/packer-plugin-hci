package example_test

import (
	"context"
	"testing"

	"github.com/wangweihong/packer-plugin-hci/hci/example"

	"github.com/wangweihong/gotoolbox/pkg/json"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"

	. "github.com/smartystreets/goconvey/convey"
)

func TestResourcePoolGet(t *testing.T) {
	// change log default logger level to debug
	Convey("", t, func() {
		Convey("tenant", func() {
			resp, err := example.NewResourcePool(c).ResourcePoolGet(context.Background(), &iexample.ResourcePoolGetRequest{
				Cluster: cluster,
				Name:    "container",
			})
			So(err, ShouldBeNil)
			json.PrintStructObject(resp)
		})

	})
}
