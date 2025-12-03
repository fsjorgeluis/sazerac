package commands

import (
	"fmt"
	"path/filepath"

	"github.com/fsjorgeluis/sazerac/internal"
	"github.com/fsjorgeluis/sazerac/internal/templates"
	"github.com/spf13/cobra"
)

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
				templates.FS,
				"validator/validator.go.tpl",
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
