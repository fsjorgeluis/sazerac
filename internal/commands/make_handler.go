package commands

import (
	"embed"
	"fmt"
	"path/filepath"

	"github.com/fsjorgeluis/sazerac/internal"
	"github.com/spf13/cobra"
)

//go:embed ../templates/handler/*
var handlerFS embed.FS

func NewMakeHandlerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "make handler <Name> <UseCase>",
		Short: "Generate the handler for a use case",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			usecase := args[1]

			out := filepath.Join(
				"internal/handlers",
				internal.ToSnake(name)+"_handler.go",
			)

			data := map[string]any{
				"Name":    name,
				"UseCase": usecase,
				"Module":  internal.GetModuleName(),
			}

			err := internal.WriteTemplate(handlerFS, "../templates/handler/handler.go.tpl", out, data)
			if err != nil {
				return err
			}

			fmt.Println("Handler served 🥃:", out)
			return nil
		},
	}

	return cmd
}
