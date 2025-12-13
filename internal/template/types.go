package template

type TemplateOption string
type TemplateOperation string

const (
	ScriptTemplate TemplateOption = "script"
	TestTemplate   TemplateOption = "test"
	AllTemplate    TemplateOption = "all"
)

const (
	GenerateOperation TemplateOperation = "generate"
)
