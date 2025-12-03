package commands

import (
	"embed"
	"fmt"
	"path/filepath"

	"github.com/fsjorgeluis/sazerac/internal"
	"github.com/spf13/cobra"
)

//go:embed ../templates/validator/*
var validatorFS embed.FS

func NewMakeValidatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "make validator <Entity>",
		Short: "Generate a simple validator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			entity := args[0]

			out := filepath.Join(
				"internal/domain/validators",
				internal.ToSnake(entity)+"_validator.go",
			)

			data := map[string]any{
				"Entity": entity,
			}

			err := internal.WriteTemplate(
				validatorFS,
				"../templates/validator/validator.go.tpl",
				out,
				data,
			)

			if err != nil {
				return err
			}

			fmt.Println("Validator served 🥃:", out)
			return nil
		},
	}

	return cmd
}
