package template

import "fmt"

func ApplyTemplateOperation(templateOption TemplateOption, operation TemplateOperation, name string) error {
	option := validateOption(templateOption)

	if !validateOperation(operation) {
		return fmt.Errorf("unknown template operation: %s", operation)
	}

	return execute(option, operation, name)
}

func validateOption(option TemplateOption) TemplateOption {
	switch option {
	case ScriptTemplate, TestTemplate, AllTemplate:
		return option
	default:
		return AllTemplate
	}
}

func validateOperation(operation TemplateOperation) bool {
	switch operation {
	case GenerateOperation:
		return true
	default:
		return false
	}
}
