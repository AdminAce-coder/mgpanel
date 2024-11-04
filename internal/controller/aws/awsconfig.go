package aws

// Configurable 接口，所有的aws服务都需要实现这个方法
type Configurable interface {
	TemplateGeneration() error
}

// AwsGenConfig 结构体，用于包含 Configurable
type AwsGenConfig struct {
	cg Configurable
}

// TemplateGeneration 实现 Configurable 接口的 TemplateGeneration 方法
func (a *AwsGenConfig) TemplateGeneration() error {
	return nil
}

// NewAwsGenConfig 函数，将传入的 Configurable 包装为 AwsGenConfig 并返回
func NewAwsGenConfig(aws Configurable) Configurable {
	return aws
}
