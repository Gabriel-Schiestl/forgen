package template

import (
	"fmt"

	"github.com/Gabriel-Schiestl/forgen/internal/generate"
)

func ApplyTemplateOperation(templateOption TemplateOption, operation TemplateOperation, params generate.TemplateParams) error {
	option := validateOption(templateOption)

	if !validateOperation(operation) {
		return fmt.Errorf("unknown template operation: %s", operation)
	}

	if params.Name == "" {
		params.Name = "MyContract"
	}
	return execute(option, operation, params)
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
