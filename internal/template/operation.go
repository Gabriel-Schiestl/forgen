package template

import (
	"github.com/Gabriel-Schiestl/forgen/internal/generate"
)

func execute(option TemplateOption, operation TemplateOperation, params generate.TemplateParams) error {
	paths := getPaths(option)

	var err error
	switch operation {
	case GenerateOperation:
		err = generate.Execute(paths, params)
	}

	return err
}

func getPaths(option TemplateOption) []string {
	paths := make([]string, 1, 2)
	switch option {
	case ScriptTemplate:
		paths[0] = "script/"
	case TestTemplate:
		paths[0] = "test/"
	case AllTemplate:
		paths[0] = "script/"
		paths = append(paths, "test/")
	}

	return paths
}
