package commands

import (
	"fmt"
	"path/filepath"

	"github.com/fsjorgeluis/sazerac/internal"
	"github.com/fsjorgeluis/sazerac/internal/templates"
	"github.com/spf13/cobra"
)

func NewMakeAllCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all <Entity> <UseCase>",
		Short: "Generate all resources in a single shot",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			entity := args[0]
			usecase := args[1]

			fmt.Println(">> Serving entity 🥃:", entity)
			entityCmd := NewMakeEntityCmd()
			if err := entityCmd.RunE(cmd, []string{entity}); err != nil {
				return err
			}

			fmt.Println(">> Serving repo 🥃:", entity)
			repoCmd := NewMakeRepoCmd()
			if err := repoCmd.RunE(cmd, []string{entity}); err != nil {
				return err
			}

			fmt.Println(">> Serving usecase 🥃:", usecase)
			usecaseCmd := NewMakeUseCaseCmd()
			if err := usecaseCmd.RunE(cmd, []string{usecase, entity}); err != nil {
				return err
			}

			fmt.Println(">> Serving handler 🥃:", usecase)
			handlerCmd := NewMakeHandlerCmd()
			if err := handlerCmd.RunE(cmd, []string{usecase, usecase}); err != nil {
				return err
			}

			fmt.Println(">> Serving dependency injection 🥃")
			projectName := internal.GetProjectName()
			if projectName == "" {
				fmt.Println("⚠️  Warning: Could not determine project name. Skipping DI generation.")
			} else {
				// Generate DI using the dedicated command
				diCmd := NewMakeDiCmd()
				if err := diCmd.RunE(cmd, []string{usecase, entity}); err != nil {
					fmt.Printf("⚠️  Warning: Failed to generate DI: %v\n", err)
				}

				// Generate main.go
				projectType := detectProjectType()
				features := getFeatureConfig()
				useCasePascal := internal.ToPascalCase(usecase)
				entityPascal := internal.ToPascalCase(entity)

				var mainPath string
				if projectType == "lambda" {
					mainPath = filepath.Join("cmd", "lambda", "main.go")
				} else {
					mainPath = filepath.Join("cmd", projectName, "main.go")
				}

				data := map[string]any{
					"UseCase":     useCasePascal,
					"Entity":      entityPascal,
					"Module":      internal.GetModuleName(),
					"ProjectName": projectName,
					"Features":    features,
				}

				templatePath := fmt.Sprintf("project_types/%s/project/main.go.tpl", projectType)
				if err := internal.WriteTemplate(templates.FS, templatePath, mainPath, data); err != nil {
					fmt.Printf("⚠️  Warning: Failed to update main.go: %v\n", err)
				} else {
					fmt.Println("Main.go updated 🥃:", mainPath)
				}
			}

			fmt.Println("✔️  Everything served successfully 🥃")

			return nil
		},
	}

	return cmd
}
