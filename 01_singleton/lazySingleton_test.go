package singleton

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_GetLazyInstance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	convey.Convey("TestUnit_GetLazyInstance", t, func() {
		convey.Convey("test", func() {
			GetLazyInstance()
		})
	})
}
