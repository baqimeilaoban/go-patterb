package build

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_NewResourcePoolConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	convey.Convey("TestUnit_NewResourcePoolConfig", t, func() {
		convey.Convey("name is empty", func() {
			opts := []ResourcePoolConfigOptionFunc{
				func(option *ResourcePoolConfigOption) {
				},
			}
			config, err := NewResourcePoolConfig("", opts...)
			convey.So(config, convey.ShouldBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("max < 0", func() {
			opts := []ResourcePoolConfigOptionFunc{
				func(option *ResourcePoolConfigOption) {
					option.maxTotal = -1
				},
			}
			config, err := NewResourcePoolConfig("test", opts...)
			convey.So(config, convey.ShouldBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("max < min", func() {
			opts := []ResourcePoolConfigOptionFunc{
				func(option *ResourcePoolConfigOption) {
					option.maxIdle = 1
					option.minIdle = 2
				},
			}
			config, err := NewResourcePoolConfig("test", opts...)
			convey.So(config, convey.ShouldBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("suc", func() {
			opts := []ResourcePoolConfigOptionFunc{
				func(option *ResourcePoolConfigOption) {
					option.maxTotal = 20
				},
			}
			poolConfig := &ResourcePoolConfig{
				name:     "test",
				maxTotal: 20,
				maxIdle:  9,
				minIdle:  1,
			}
			config, err := NewResourcePoolConfig("test", opts...)
			convey.So(reflect.DeepEqual(config, poolConfig), convey.ShouldBeTrue)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}
