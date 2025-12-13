package command

import (
	"github.com/Gabriel-Schiestl/forgen/internal/template"
	"github.com/spf13/cobra"
)

var genCommand = &cobra.Command{
	Use:   "gen [option]",
	Short: "Generate a Foundry template",
	Long: `Generates a Foundry template for the specified option.
The script will be created in the 'script' directory of your Foundry project.
If no option is provided, it will generate templates for all options(script, test).
	`,
	Args: cobra.MaximumNArgs(1),
	Run:  runGenCommand,
}

func runGenCommand(cmd *cobra.Command, args []string) {
	scriptOption := template.AllTemplate
	if len(args) == 1 {
		scriptOption = template.TemplateOption(args[0])
	}

	err := template.ApplyTemplateOperation(scriptOption, template.GenerateOperation, Name)
	if err != nil {
		panic(err)
	}
}
