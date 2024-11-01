package aws

type Configurable interface {
	TemplateGeneration(name string) error // config示例需要实现这个方法
}

type TerrformCgGen struct {
	Configurable
}
