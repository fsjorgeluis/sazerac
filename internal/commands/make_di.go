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
			projectType := detectProjectType()
			features := getFeatureConfig()

			projectName := internal.GetProjectName()
			if projectName == "" {
				return fmt.Errorf("could not determine project name")
			}

			var out string
			if projectType == "lambda" {
				out = filepath.Join("cmd", "lambda", "di", "di.go")
			} else {
				out = filepath.Join("cmd", projectName, "di", "di.go")
			}

			useCaseRoute := strings.ToLower(internal.ToSnake(useCasePascal))

			data := map[string]any{
				"UseCase":      useCasePascal,
				"Entity":       entityPascal,
				"Module":       internal.GetModuleName(),
				"ProjectName":  projectName,
				"UseCaseRoute": useCaseRoute,
				"Features":     features,
			}

			templatePath := fmt.Sprintf("project_types/%s/project/di.go.tpl", projectType)
			if err := internal.WriteTemplate(templates.FS, templatePath, out, data); err != nil {
				return err
			}

			fmt.Println("DI container served:", out)
			return nil
		},
	}

	return cmd
}
