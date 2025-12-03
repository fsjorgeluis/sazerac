package commands

import (
	"embed"
	"fmt"
	"path/filepath"

	"github.com/fsjorgeluis/sazerac/internal"
	"github.com/spf13/cobra"
)

//go:embed ../templates/mapper/*
var mapperFS embed.FS

func NewMakeMapperCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "make mapper <Entity>",
		Short: "Generate a entity mapper <-> DTO",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			entity := args[0]

			out := filepath.Join(
				"internal/domain/mappers",
				internal.ToSnake(entity)+"_mapper.go",
			)

			data := map[string]any{
				"Entity": entity,
				"Module": internal.GetModuleName(),
			}

			err := internal.WriteTemplate(mapperFS, "../templates/mapper/mapper.go.tpl", out, data)
			if err != nil {
				return err
			}

			fmt.Println("Mapper served 🥃:", out)
			return nil
		},
	}

	return cmd
}
