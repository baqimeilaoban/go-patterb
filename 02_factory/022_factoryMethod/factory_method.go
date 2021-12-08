package factory

// IRuleConfigParser 解析config的接口
type IRuleConfigParser interface {
	Parse(data []byte)
}

// jsonRuleConfigParser 解析JSON的结构体
type jsonRuleConfigParser struct {
}

// Parse 解析
func (j jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// yamlRuleConfigParser 解析yaml的结构体
type yamlRuleConfigParser struct {
}

// Parse 解析
func (y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// IRuleConfigParserFactory 工厂方法接口
type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

// yamlRuleConfigParseFactory yamlRuleConfigParseFactory的工厂类
type yamlRuleConfigParseFactory struct {
}

// CreateParser 创建解析
func (y yamlRuleConfigParseFactory) CreateParser() IRuleConfigParser {
	return yamlRuleConfigParser{}
}

// jsonRuleConfigParseFactory jsonRuleConfigParseFactory的工厂类
type jsonRuleConfigParseFactory struct {
}

// CreateParser 创建解析
func (j jsonRuleConfigParseFactory) CreateParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

// NewIRuleConfigParserFactory 用一个简单方法封装工厂方法
func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return jsonRuleConfigParseFactory{}
	case "yaml":
		return yamlRuleConfigParseFactory{}
	}
	return nil
}
