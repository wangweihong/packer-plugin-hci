package example_test

import (
	"context"
	"testing"

	"github.com/wangweihong/packer-plugin-hci/hci/example"

	"github.com/wangweihong/gotoolbox/pkg/json"
	"github.com/wangweihong/packer-plugin-hci/hci/example/iexample"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	c, _      = example.NewClient("http://10.30.100.25:9990", "system", "wwhvw", "wwhvw")
	cluster   = "49bb08f9-1c60-49ee-85d6-6fde276895c5"
	container = "4b538ae7-9872-46f9-a169-876c81c5b68a"
	pool      = "80ed5589-646f-4f38-8c14-a0c6750ba23c"
	repo      = "89be9434-6d2d-4df2-839c-0975aa4d68b1"
	ecs       = "5106c8e5-185a-49e3-b66a-084bfdbb68cd"
	image     = "94dfc033-8b26-479e-8e6b-f36102674efb"
)

func TestClient(t *testing.T) {
	Convey("", t, func() {
		SkipConvey("master", func() {
			resp, err := example.NewSystem(c).ServerIPList(context.Background(), &iexample.ServerIPListRequest{})
			So(err, ShouldBeNil)

			json.PrintStructObject(resp)
		})
		Convey("not master", func() {
			resp, err := example.NewSystem(c).ServerIPList(context.Background(), &iexample.ServerIPListRequest{})
			So(err, ShouldBeNil)
			json.PrintStructObject(resp)
		})
	})

}
