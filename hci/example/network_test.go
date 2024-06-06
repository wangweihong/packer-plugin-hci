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

func TestNetworkSwitchGet(t *testing.T) {
	// change log default logger level to debug
	Convey("", t, func() {
		Convey("tenant", func() {
			os.Setenv("HTTPCLI_DEBUG", "1")
			os.Setenv("HTTPCLI_DEBUG_HUGE", "1")
			resp, err := example.NewNetwork(c).VirtualSwitchGet(context.Background(), &iexample.VirtualSwitchGetRequest{
				Cluster: cluster,
				Name:    "manage",
				Tenant:  container,
			})
			So(err, ShouldBeNil)
			json.PrintStructObject(resp)
		})

	})
}

func TestNetworkSwitchPortGroupGet(t *testing.T) {
	// change log default logger level to debug
	Convey("", t, func() {
		Convey("tenant", func() {
			//	os.Setenv("HTTPCLI_DEBUG", "1")
			//	os.Setenv("HTTPCLI_DEBUG_HUGE", "1")
			resp, err := example.NewNetwork(c).VirtualSwitchPortGroupGet(context.Background(), &iexample.VirtualSwitchPortGroupGetRequest{
				Cluster:     cluster,
				Name:        "default",
				Tenant:      container,
				VSwitchName: "manage",
			})
			So(err, ShouldBeNil)
			json.PrintStructObject(resp)
		})

	})
}
