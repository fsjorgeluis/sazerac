package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func NewMakeAllCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all [Entity] [UseCase]",
		Short: "Generate all resources in a single shot",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			entity := args[0]
			usecase := args[1]

			// Basic normalization
			entityLower := strings.ToLower(entity)
			useCaseLower := strings.ToLower(usecase)

			fmt.Println(">> Serving entity 🥃:", entity)
			_ = NewMakeEntityCmd().RunE(cmd, []string{entityLower})

			fmt.Println(">> Serving repo 🥃:", entity)
			_ = NewMakeRepoCmd().RunE(cmd, []string{entityLower})

			fmt.Println(">> Serving usecase 🥃:", usecase)
			_ = NewMakeUseCaseCmd().RunE(cmd, []string{useCaseLower})

			fmt.Println(">> Serving handler 🥃:", usecase)
			_ = NewMakeHandlerCmd().RunE(cmd, []string{useCaseLower})

			fmt.Println("✔️  Everything served successfully 🥃")

			return nil
		},
	}

	return cmd
}
