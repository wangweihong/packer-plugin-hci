package example_test

import (
	"context"
	"testing"

	"github.com/wangweihong/packer-plugin-hci/hci/example"

	"github.com/wangweihong/gotoolbox/pkg/json"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpecGet(t *testing.T) {
	// change log default logger level to debug
	Convey("", t, func() {
		Convey("tenant", func() {
			resp, err := example.NewSpecification(c).SpecificationGet(context.Background(), &iexample.SpecGetRequest{
				Name: "高性能虚拟机",
			})
			So(err, ShouldBeNil)
			json.PrintStructObject(resp)
		})

	})
}
