package command

import (
	"github.com/Gabriel-Schiestl/forgen/internal/generate"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "fg",
	Short: "Simple Foundry code templates generator",
	Long: `forgen is a powerful CLI tool for generating code templates, like scripts and tests for foundry projects.

It streamlines the development process by providing ready-to-use templates, allowing developers to focus on writing code rather than boilerplate setup.
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var params = generate.TemplateParams{}

func init() {
	RootCmd.PersistentFlags().StringVarP(&params.Name, "name", "n", "", "Specify the module name for the generated code")
	RootCmd.PersistentFlags().StringVarP(&params.Version, "version", "v", "0.8.30", "Specify the Solidity version for the generated code")
	RootCmd.AddCommand(genCommand)
}
