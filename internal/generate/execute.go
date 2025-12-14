package generate

import (
	"strings"

	"github.com/Gabriel-Schiestl/forgen/internal/base"
	"github.com/Gabriel-Schiestl/forgen/internal/utils"
)

//create a struct params(name, path to create)
func Execute(paths []string, name string) error {
	for _, p := range paths {
		switch p {
		case "/script":
			err := generate(name, "script")
			if err != nil {
				return err
			}
		case "/test":
			err := generate(name, "test")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func generate(name, operation string) error {
	nameVariable := utils.GetCamelCase(name)
	exists := utils.FindDir(operation)
	if !exists {
		if err := utils.CreateDir(operation); err != nil {
			return err
		}
	}

	path := operation + "/" + name + ".s.sol"

	baseTemplate := base.Bases[operation]

	baseTemplate = strings.ReplaceAll(baseTemplate, "[name]", name)
	baseTemplate = strings.ReplaceAll(baseTemplate, "[name_variable]", nameVariable)
	baseTemplate = strings.ReplaceAll(baseTemplate, "[version]", "0.8.30")
	
	if err := utils.CreateFile(path, baseTemplate); err != nil {
		return err
	}

	return nil
}