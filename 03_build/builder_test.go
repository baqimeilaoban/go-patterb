package build

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/smartystreets/goconvey/convey"
)

func TestUint_Build(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	convey.Convey("TestUint_Build", t, func() {
		convey.Convey("name is empty", func() {
			builder := &ResourcePoolConfigBuilder{
				name:     "",
				maxTotal: 10,
				maxIdle:  9,
				minIdle:  1,
			}
			config, err := builder.Build()
			convey.So(config, convey.ShouldBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("maxTotal < maxIdle", func() {
			builder := &ResourcePoolConfigBuilder{
				name:     "test",
				maxTotal: 2,
				maxIdle:  10,
				minIdle:  1,
			}
			config, err := builder.Build()
			convey.So(config, convey.ShouldBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("maxIdle < minIdle", func() {
			builder := &ResourcePoolConfigBuilder{
				name:     "test",
				maxTotal: 10,
				maxIdle:  1,
				minIdle:  9,
			}
			config, err := builder.Build()
			convey.So(config, convey.ShouldBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("suc", func() {
			builder := &ResourcePoolConfigBuilder{
				name:     "test",
				maxTotal: 0,
				maxIdle:  0,
				minIdle:  0,
			}
			poolConfig := &ResourcePoolConfig{
				name:     "test",
				maxTotal: 10,
				maxIdle:  9,
				minIdle:  1,
			}
			config, err := builder.Build()
			convey.So(reflect.DeepEqual(config, poolConfig), convey.ShouldBeTrue)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}
