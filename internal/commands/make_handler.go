package commands

import (
	"fmt"
	"path/filepath"

	"github.com/fsjorgeluis/sazerac/internal"
	"github.com/fsjorgeluis/sazerac/internal/templates"
	"github.com/spf13/cobra"
)

func NewMakeHandlerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "handler <Name> <UseCase>",
		Short: "Generate the handler for a use case",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			usecase := args[1]
			namePascal := internal.ToPascalCase(name)
			useCasePascal := internal.ToPascalCase(usecase)
			projectType := detectProjectType()
			features := getFeatureConfig()

			out := filepath.Join(
				"internal/handlers",
				internal.ToSnake(name)+"_handler.go",
			)

			data := map[string]any{
				"Name":     namePascal,
				"UseCase":  useCasePascal,
				"Module":   internal.GetModuleName(),
				"Features": features,
			}

			// Use project-type specific template
			var templatePath string
			if projectType == "lambda" {
				templatePath = "project_types/lambda/handler/lambda_handler.go.tpl"
			} else {
				templatePath = "project_types/cli/handler/handler.go.tpl"
			}

			err := internal.WriteTemplate(templates.FS, templatePath, out, data)
			if err != nil {
				return err
			}

			fmt.Println("Handler served 🥃:", out)
			return nil
		},
	}

	return cmd
}
