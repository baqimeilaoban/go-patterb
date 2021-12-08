package factory

// IRuleConfigParser 解析config的接口
type IRuleConfigParser interface {
	Parse(data []byte)
}

// jsonRuleConfigParser 解析JSON
type jsonRuleConfigParser struct {
}

// yamlRuleConfigParser 解析yaml
type yamlRuleConfigParser struct {
}

// Parse 解析
func (j jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// Parse 解析
func (y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// NewIRuleConfigParser 创建对象
func NewIRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return jsonRuleConfigParser{}
	case "yaml":
		return yamlRuleConfigParser{}
	}
	return nil
}
