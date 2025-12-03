package commands

import (
	"fmt"
	"path/filepath"

	"github.com/fsjorgeluis/sazerac/internal"
	"github.com/fsjorgeluis/sazerac/internal/templates"
	"github.com/spf13/cobra"
)

func NewMakeMapperCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mapper <Entity>",
		Short: "Generate a entity mapper <-> DTO",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			entity := args[0]
			entityPascal := internal.ToPascalCase(entity)

			out := filepath.Join(
				"internal/domain/mappers",
				internal.ToSnake(entity)+"_mapper.go",
			)

			data := map[string]any{
				"Entity": entityPascal,
				"Module": internal.GetModuleName(),
			}

			err := internal.WriteTemplate(templates.FS, "mapper/mapper.go.tpl", out, data)
			if err != nil {
				return err
			}

			fmt.Println("Mapper served 🥃:", out)
			return nil
		},
	}

	return cmd
}
