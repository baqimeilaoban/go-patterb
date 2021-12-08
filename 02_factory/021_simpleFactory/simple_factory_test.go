package factory

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_NewIRuleConfigParser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	convey.Convey("TestUnit_NewIRuleConfigParser", t, func() {
		convey.Convey("json", func() {
			jsonParser := jsonRuleConfigParser{}
			parser := NewIRuleConfigParser("json")
			convey.So(reflect.DeepEqual(jsonParser, parser), convey.ShouldBeTrue)
		})
		convey.Convey("yaml", func() {
			yamlParser := yamlRuleConfigParser{}
			parser := NewIRuleConfigParser("yaml")
			convey.So(reflect.DeepEqual(yamlParser, parser), convey.ShouldBeTrue)
		})
		convey.Convey("nil", func() {
			parser := NewIRuleConfigParser("test")
			convey.So(parser, convey.ShouldBeNil)
		})
	})
}
