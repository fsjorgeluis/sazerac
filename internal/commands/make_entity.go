package commands

import (
	"embed"
	"fmt"
	"path/filepath"

	"github.com/fsjorgeluis/sazerac/internal"
	"github.com/spf13/cobra"
)

//go:embed ../templates/entity/*
var entityFS embed.FS

func NewMakeEntityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "make entity <Name>",
		Short: "Generates a domain entity",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			out := filepath.Join("internal/domain/entities", internal.ToSnake(name)+".go")

			data := map[string]any{
				"Name": name,
			}

			if err := internal.WriteTemplate(entityFS, "../templates/entity/entity.go.tpl", out, data); err != nil {
				return err
			}

			fmt.Println("Entity ready:", out)
			return nil
		},
	}

	return cmd
}
