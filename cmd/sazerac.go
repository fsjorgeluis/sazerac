package cmd

import (
	"github.com/fsjorgeluis/sazerac/internal/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sazerac",
	Short: "CLI Clean Architecture",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(commands.NewInitCmd())
	rootCmd.AddCommand(commands.NewConfigCmd())

	// Create make command as parent
	makeCmd := &cobra.Command{
		Use:   "make",
		Short: "Generate components for Clean Architecture",
		Long:  "Generate entities, repositories, use cases, handlers, mappers, and validators",
	}

	// Add all make subcommands
	makeCmd.AddCommand(commands.NewMakeEntityCmd())
	makeCmd.AddCommand(commands.NewMakeRepoCmd())
	makeCmd.AddCommand(commands.NewMakeUseCaseCmd())
	makeCmd.AddCommand(commands.NewMakeHandlerCmd())
	makeCmd.AddCommand(commands.NewMakeMapperCmd())
	makeCmd.AddCommand(commands.NewMakeValidatorCmd())
	makeCmd.AddCommand(commands.NewMakeDiCmd())
	makeCmd.AddCommand(commands.NewMakeAllCmd())

	rootCmd.AddCommand(makeCmd)
}
