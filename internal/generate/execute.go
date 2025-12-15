package generate

import (
	"strings"

	"github.com/Gabriel-Schiestl/forgen/internal/base"
	"github.com/Gabriel-Schiestl/forgen/internal/utils"
)

type TemplateParams struct {
	Name string
	Version string
}

//create a struct params(name, path to create)
func Execute(paths []string, params TemplateParams) error {
	for _, p := range paths {
		switch p {
		case "script/":
			err := generate(params, "script")
			if err != nil {
				return err
			}
		case "test/":
			err := generate(params, "test")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func generate(params TemplateParams, operation string) error {
	nameVariable := utils.GetCamelCase(params.Name)
	exists := utils.FindDir(operation)
	if !exists {
		if err := utils.CreateDir(operation); err != nil {
			return err
		}
	}

	endingPath := operation[0:1]

	path := operation + "/" + params.Name + "." + endingPath + ".sol"

	baseTemplate := base.Bases[operation]

	baseTemplate = strings.ReplaceAll(baseTemplate, "[name]", params.Name)
	baseTemplate = strings.ReplaceAll(baseTemplate, "[name_variable]", nameVariable)
	baseTemplate = strings.ReplaceAll(baseTemplate, "[version]", params.Version)
	
	if err := utils.CreateFile(path, baseTemplate); err != nil {
		return err
	}

	return nil
}