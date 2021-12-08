package factory

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_NewIRuleConfigParserFactory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	convey.Convey("TestUnit_NewIRuleConfigParserFactory", t, func() {
		convey.Convey("json", func() {
			jsonParserFactory := jsonRuleConfigParseFactory{}
			jsonFactory := NewIRuleConfigParserFactory("json")
			convey.So(reflect.DeepEqual(jsonFactory, jsonParserFactory), convey.ShouldBeTrue)
		})
		convey.Convey("yaml", func() {
			yamlParserFactory := yamlRuleConfigParseFactory{}
			yamlFactory := NewIRuleConfigParserFactory("yaml")
			convey.So(reflect.DeepEqual(yamlFactory, yamlParserFactory), convey.ShouldBeTrue)
		})
		convey.Convey("test", func() {
			nilFactory := NewIRuleConfigParserFactory("test")
			convey.So(nilFactory, convey.ShouldBeNil)
		})
	})
}
