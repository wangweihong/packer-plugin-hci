package example_test

import (
	"context"
	"testing"

	"github.com/wangweihong/packer-plugin-hci/hci/example"

	"github.com/wangweihong/gotoolbox/pkg/json"
	"github.com/wangweihong/gotoolbox/pkg/typeutil"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTenantGet(t *testing.T) {
	// change log default logger level to debug
	Convey("", t, func() {
		Convey("tenant", func() {
			resp, err := example.NewOperationCenter(c).TenantGet(context.Background(), &iexample.TenantGetRequest{
				TenantName: typeutil.String("container"),
			})
			So(err, ShouldBeNil)
			json.PrintStructObject(resp)
		})
		SkipConvey("uuid", func() {
			resp, err := example.NewOperationCenter(c).TenantGet(context.Background(), &iexample.TenantGetRequest{
				Tenant: "system",
			})
			So(err, ShouldBeNil)
			json.PrintStructObject(resp)
		})
	})
}
