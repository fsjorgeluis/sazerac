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
	rootCmd.AddCommand(commands.NewMakeEntityCmd())
	rootCmd.AddCommand(commands.NewMakeUseCaseCmd())
	rootCmd.AddCommand(commands.NewMakeRepoCmd())
	rootCmd.AddCommand(commands.NewMakeHandlerCmd())
	rootCmd.AddCommand(commands.NewMakeMapperCmd())
	rootCmd.AddCommand(commands.NewMakeValidatorCmd())

	rootCmd.AddCommand(commands.NewMakeAllCmd())
}
