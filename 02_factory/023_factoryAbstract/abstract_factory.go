package factory

// IRuleConfigParser 解析规则接口
type IRuleConfigParser interface {
	Parse(data []byte)
}

// jsonRuleConfigParser
type jsonRuleConfigParser struct {
}

func (j jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// ISystemConfigParser 抽象接口
type ISystemConfigParser interface {
	ParserSystem(data []byte)
}

type jsonSystemConfigParser struct {
}

func (j jsonSystemConfigParser) ParserSystem(data []byte) {
	panic("implement me")
}

// IConfigParserFactory 工厂方法接口
type IConfigParserFactory interface {
	CreateRuleParser() IRuleConfigParser
	CreateSystemParser() ISystemConfigParser
}

type jsonConfigParserFactory struct {
}

func (j jsonConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

func (j jsonConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return jsonSystemConfigParser{}
}
