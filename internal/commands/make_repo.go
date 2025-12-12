package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsjorgeluis/sazerac/internal"
	"github.com/fsjorgeluis/sazerac/internal/config"
	"github.com/fsjorgeluis/sazerac/internal/templates"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// detectProjectType tries to detect project type from .sazerac.yaml or defaults to CLI
func detectProjectType() string {
	// Try to read .sazerac.yaml
	if data, err := os.ReadFile(".sazerac.yaml"); err == nil {
		var cfg ProjectConfig
		if err := yaml.Unmarshal(data, &cfg); err == nil {
			return cfg.Project.Type
		}
	}

	// Default to CLI
	return "cli"
}

// getFeatureConfig loads feature config or returns defaults
func getFeatureConfig() *config.FeatureConfig {
	// Try to read .sazerac.yaml
	if data, err := os.ReadFile(".sazerac.yaml"); err == nil {
		var cfg ProjectConfig
		if err := yaml.Unmarshal(data, &cfg); err == nil {
			return &config.FeatureConfig{
				Database:      cfg.Features.Database,
				Tests:         cfg.Features.Tests,
				ErrorHandling: cfg.Features.ErrorHandling,
				Docker:        cfg.Features.Docker,
				SAMTemplate:   cfg.Features.SAMTemplate,
				APIGateway:    cfg.Features.APIGateway,
			}
		}
	}

	// Return defaults
	return &config.FeatureConfig{
		ErrorHandling: true,
		Tests:         false,
		Database:      "none",
	}
}

func NewMakeRepoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repo <Entity>",
		Short: "Generate repository interface and implementation",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			entity := args[0]
			entityPascal := internal.ToPascalCase(entity)
			projectType := detectProjectType()
			features := getFeatureConfig()

			data := map[string]any{
				"Entity":   entityPascal,
				"Module":   internal.GetModuleName(),
				"Features": features,
			}

			// Repository interface
			outInterface := filepath.Join(
				"internal/repository",
				internal.ToSnake(entity)+"_repository.go",
			)

			// Generate interface
			interfacePath := fmt.Sprintf("project_types/%s/repository/repository.go.tpl", projectType)
			if err := internal.WriteTemplate(templates.FS, interfacePath, outInterface, data); err != nil {
				return err
			}

			fmt.Println("Repository interface served 🥃:", outInterface)

			// Determine implementation based on database
			var implPath, implTemplate string
			switch features.Database {
			case "dynamodb":
				implPath = filepath.Join("infrastructure", "database", "dynamodb", fmt.Sprintf("%s_dynamodb.go", internal.ToSnake(entity)))
				implTemplate = "infrastructure/dynamodb/repo_dynamodb.go.tpl"
				fmt.Printf("dynamodb implementation served 🥃: %s\n", implPath)
			case "mysql", "mysql-rds":
				implPath = filepath.Join("infrastructure", "database", "mysql", fmt.Sprintf("%s_mysql.go", internal.ToSnake(entity)))
				implTemplate = "infrastructure/mysql/repo_mysql.go.tpl"
				fmt.Printf("mysql implementation served 🥃: %s\n", implPath)
			case "none":
				implPath = filepath.Join("infrastructure", "database", "inmemory", fmt.Sprintf("%s_inmemory.go", internal.ToSnake(entity)))
				implTemplate = "infrastructure/inmemory/repo_inmemory.go.tpl"
				fmt.Printf("inmemory implementation served 🥃: %s\n", implPath)
			default:
				return fmt.Errorf("unsupported database type: %s", features.Database)
			}

			implData := map[string]any{
				"Entity":   entityPascal,
				"Module":   internal.GetModuleName(),
				"Features": features,
			}

			if err := internal.WriteTemplate(templates.FS, implTemplate, implPath, implData); err != nil {
				return fmt.Errorf("failed to generate implementation: %w", err)
			}

			return nil
		},
	}

	return cmd
}
