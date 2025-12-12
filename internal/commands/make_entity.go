package commands

import (
	"fmt"
	"path/filepath"

	"github.com/fsjorgeluis/sazerac/internal"
	"github.com/fsjorgeluis/sazerac/internal/templates"
	"github.com/spf13/cobra"
)

func NewMakeEntityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "entity <Name>",
		Short: "Generates a domain entity",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			namePascal := internal.ToPascalCase(name)

			out := filepath.Join("internal/domain/entities", internal.ToSnake(name)+".go")

			data := map[string]any{
				"Name": namePascal,
			}

			if err := internal.WriteTemplate(templates.FS, "common/entity/entity.go.tpl", out, data); err != nil {
				return err
			}

			fmt.Println("Entity ready:", out)
			return nil
		},
	}

	return cmd
}
