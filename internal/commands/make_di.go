package commands

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/fsjorgeluis/sazerac/internal"
	"github.com/fsjorgeluis/sazerac/internal/templates"
	"github.com/spf13/cobra"
)

func NewMakeDiCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "di <UseCase> <Entity>",
		Short: "Generate dependency injection container",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			usecase := args[0]
			entity := args[1]
			useCasePascal := internal.ToPascalCase(usecase)
			entityPascal := internal.ToPascalCase(entity)

			projectName := internal.GetProjectName()
			if projectName == "" {
				return fmt.Errorf("could not determine project name. Make sure you're in the project root directory")
			}

			out := filepath.Join("cmd", projectName, "di", "di.go")

			// Convert usecase name to route (e.g., CreateUser -> create-user or createuser)
			useCaseRoute := strings.ToLower(internal.ToSnake(useCasePascal))

			data := map[string]any{
				"UseCase":      useCasePascal,
				"Entity":       entityPascal,
				"Module":       internal.GetModuleName(),
				"ProjectName":  projectName,
				"UseCaseRoute": useCaseRoute,
			}

			if err := internal.WriteTemplate(templates.FS, "project/di.go.tpl", out, data); err != nil {
				return err
			}

			fmt.Println("Dependency injection container served 🥃:", out)
			return nil
		},
	}

	return cmd
}

