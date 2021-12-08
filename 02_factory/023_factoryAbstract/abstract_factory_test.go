package factory

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/smartystreets/goconvey/convey"
)

var j = jsonConfigParserFactory{}

func TestUnit_CreateRuleParser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	convey.Convey("TestUnit_CreateRuleParser", t, func() {
		convey.Convey("test", func() {
			iRuleParser := jsonRuleConfigParser{}
			parser := j.CreateRuleParser()
			convey.So(reflect.DeepEqual(iRuleParser, parser), convey.ShouldBeTrue)
		})
	})
}

func TestUnit_CreateSystemParser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	convey.Convey("TestUnit_CreateSystemParser", t, func() {
		convey.Convey("test", func() {
			iSystemParser := jsonSystemConfigParser{}
			parser := j.CreateSystemParser()
			convey.So(reflect.DeepEqual(iSystemParser, parser), convey.ShouldBeTrue)
		})
	})
}
