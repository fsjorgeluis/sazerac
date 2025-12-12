package commands

import (
	"fmt"
	"path/filepath"

	"github.com/fsjorgeluis/sazerac/internal"
	"github.com/fsjorgeluis/sazerac/internal/templates"
	"github.com/spf13/cobra"
)

func NewMakeUseCaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "usecase <Name> <Entity>",
		Short: "Generate a usecase",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			entity := args[1]
			namePascal := internal.ToPascalCase(name)
			entityPascal := internal.ToPascalCase(entity)
			projectType := detectProjectType()
			features := getFeatureConfig()

			out := filepath.Join(
				"internal/usecases",
				internal.ToSnake(name)+"_usecase.go",
			)

			data := map[string]any{
				"Name":     namePascal,
				"Entity":   entityPascal,
				"Module":   internal.GetModuleName(),
				"Features": features,
			}

			// Use project-type specific template
			templatePath := fmt.Sprintf("project_types/%s/usecase/usecase.go.tpl", projectType)
			err := internal.WriteTemplate(
				templates.FS,
				templatePath,
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
