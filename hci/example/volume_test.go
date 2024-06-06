package example_test

import (
	"context"
	"testing"

	"github.com/wangweihong/packer-plugin-hci/hci/example"

	"github.com/wangweihong/gotoolbox/pkg/json"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVolumeList(t *testing.T) {
	// change log default logger level to debug
	Convey("", t, func() {
		SkipConvey("tenant", func() {
			resp, err := example.NewVolume(c).PoolList(context.Background(), &iexample.PoolListRequest{
				Cluster: cluster,
				Tenant:  container,
			})
			So(err, ShouldBeNil)
			So(len(resp.Data.List), ShouldEqual, 2)
			json.PrintStructObject(resp)
		})

		Convey("filter", func() {
			resp, err := example.NewVolume(c).PoolList(context.Background(), &iexample.PoolListRequest{
				Cluster: cluster,
				Tenant:  container,
				//FilterAvailMoreThan: 107374182400,
			})
			So(err, ShouldBeNil)
			json.PrintStructObject(resp)
		})
	})
}
