package template

import (
	"github.com/Gabriel-Schiestl/forgen/internal/generate"
)

func execute(option TemplateOption, operation TemplateOperation, name string) error {
	paths := getPaths(option)

	var err error
	switch operation {
	case GenerateOperation:
		err = generate.Execute(paths, name)
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
