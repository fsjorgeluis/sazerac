package commands

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsjorgeluis/sazerac/internal"
	"github.com/spf13/cobra"
)

//go:embed ../templates/project/*
var projectFS embed.FS

func NewInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init <project-name>",
		Short: "Start a project with Clean Architecture",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			module := fmt.Sprintf("github.com/<UserName>/%s", name)

			paths := map[string]string{
				"../templates/project/main.go.tpl":   filepath.Join(name, "cmd", name, "main.go"),
				"../templates/project/go.mod.tpl":    filepath.Join(name, "go.mod"),
				"../templates/project/readme.md.tpl": filepath.Join(name, "README.md"),
			}

			data := map[string]any{
				"ProjectName": name,
				"Module":      module,
			}

			for tpl, out := range paths {
				if err := internal.WriteTemplate(projectFS, tpl, out, data); err != nil {
					return err
				}
			}

			// create an empty structure
			dirs := []string{
				"internal/domain/entities",
				"internal/domain/usecase",
				"internal/domain/interfaces",
				"internal/infrastructure/repositories",
				"internal/infrastructure/http",
			}
			for _, d := range dirs {
				if err := os.MkdirAll(filepath.Join(name, d), 0755); err != nil {
					return err
				}
			}

			fmt.Println("Project ready:", name)
			return nil
		},
	}
	return cmd
}
