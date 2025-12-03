package commands

import (
	"embed"
	"fmt"
	"path/filepath"

	"github.com/fsjorgeluis/sazerac/internal"
	"github.com/spf13/cobra"
)

//go:embed ../templates/usecase/*
var useCaseFS embed.FS

func NewMakeUseCaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "make usecase <Name> <Entity>",
		Short: "Generate a usecase",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			entity := args[1]

			out := filepath.Join(
				"internal/usecases",
				internal.ToSnake(name)+"_usecase.go",
			)

			data := map[string]any{
				"Name":   name,
				"Entity": entity,
				"Module": internal.GetModuleName(),
			}

			err := internal.WriteTemplate(
				useCaseFS,
				"../templates/usecase/usecase.go.tpl",
				out,
				data,
			)

			if err != nil {
				return err
			}

			fmt.Println("UseCase served 🥃:", out)
			return nil
		},
	}

	return cmd
}
