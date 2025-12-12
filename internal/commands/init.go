package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsjorgeluis/sazerac/internal"
	"github.com/fsjorgeluis/sazerac/internal/config"
	"github.com/fsjorgeluis/sazerac/internal/prompts"
	"github.com/fsjorgeluis/sazerac/internal/templates"
	"github.com/spf13/cobra"
)

func NewInitCmd() *cobra.Command {
	var (
		projectType string
		module      string
		database    string
		skipTests   bool
		skipErrors  bool
		docker      bool
		sam         bool
		apiGateway  bool
		listTypes   bool
	)

	cmd := &cobra.Command{
		Use:   "init [project-name]",
		Short: "Start a project with Clean Architecture",
		Long: `Initialize a new project with Clean Architecture.

If no flags are provided, interactive mode will be activated.

Examples:
  # Interactive mode
  sazerac init

  # With flags
  sazerac init my-project --type lambda --db dynamodb

  # List available project types
  sazerac init --list-types`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// List types if requested
			if listTypes {
				fmt.Println("Available project types:")
				fmt.Println("  - cli     : Command-line application")
				fmt.Println("  - lambda  : AWS Lambda function")
				return nil
			}

			var (
				projectName string
				err         error
			)

			// Determine if we're in interactive mode
			interactive := len(args) == 0 || projectType == ""

			if interactive {
				// Interactive mode
				fmt.Println("🥃 Welcome to Sazerac - Clean Architecture CLI")
				fmt.Println()

				// Get project name
				if len(args) > 0 {
					projectName = args[0]
				} else {
					projectName, err = prompts.ProjectNamePrompt()
					if err != nil {
						return err
					}
				}

				// Get project type
				if projectType == "" {
					projectType, err = prompts.ProjectTypePrompt()
					if err != nil {
						return err
					}
				}

				// Get module path
				if module == "" {
					defaultModule := fmt.Sprintf("github.com/user/%s", projectName)
					module, err = prompts.ModulePathPrompt(defaultModule)
					if err != nil {
						return err
					}
				}

				// Load manifest
				manifest, err := config.LoadManifest(templates.FS, projectType)
				if err != nil {
					return fmt.Errorf("failed to load manifest: %w", err)
				}

				// Get features
				features, err := prompts.FeaturePrompts(manifest)
				if err != nil {
					return err
				}

				// Show summary
				prompts.DisplaySummary(projectName, projectType, module, features.Database, features)

				// Confirm
				confirmed, err := prompts.ConfirmationPrompt("Proceed with this configuration?")
				if err != nil {
					return err
				}
				if !confirmed {
					fmt.Println("Cancelled.")
					return nil
				}

				// Use feature values
				database = features.Database
				skipTests = !features.Tests
				skipErrors = !features.ErrorHandling
				docker = features.Docker
				sam = features.SAMTemplate
				apiGateway = features.APIGateway

			} else {
				// Non-interactive mode with flags
				if len(args) == 0 {
					return fmt.Errorf("project name is required")
				}
				projectName = args[0]

				if module == "" {
					module = fmt.Sprintf("github.com/user/%s", projectName)
				}
			}

			// Load manifest for validation
			manifest, err := config.LoadManifest(templates.FS, projectType)
			if err != nil {
				return fmt.Errorf("failed to load manifest: %w", err)
			}

			// Build feature config
			featureConfig := &config.FeatureConfig{
				Database:      database,
				Tests:         !skipTests,
				ErrorHandling: !skipErrors,
				Docker:        docker,
				SAMTemplate:   sam,
				APIGateway:    apiGateway,
			}

			// Validate configuration
			if err := manifest.ValidateFeatureConfig(featureConfig); err != nil {
				return err
			}

			// Create project
			if err := createProject(projectName, projectType, module, manifest, featureConfig); err != nil {
				return err
			}

			fmt.Printf("\n✅ Project '%s' created successfully!\n", projectName)
			fmt.Println("\nNext steps:")
			fmt.Printf("  cd %s\n", projectName)
			fmt.Println("  go mod tidy")
			if projectType == "lambda" {
				fmt.Println("  # Build for Lambda:")
				fmt.Println("  GOOS=linux GOARCH=amd64 go build -o bootstrap cmd/lambda/main.go")
				if sam {
					fmt.Println("  # Or deploy with SAM:")
					fmt.Println("  sam build && sam deploy --guided")
				}
			} else {
				fmt.Printf("  go run cmd/%s/main.go\n", projectName)
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&projectType, "type", "", "Project type (cli, lambda)")
	cmd.Flags().StringVar(&module, "module", "", "Go module path")
	cmd.Flags().StringVar(&database, "db", "", "Database type (none, mysql, dynamodb, mysql-rds)")
	cmd.Flags().BoolVar(&skipTests, "skip-tests", false, "Skip generating tests")
	cmd.Flags().BoolVar(&skipErrors, "skip-errors", false, "Skip error management")
	cmd.Flags().BoolVar(&docker, "docker", false, "Generate Dockerfile (Lambda only)")
	cmd.Flags().BoolVar(&sam, "sam", false, "Generate SAM template (Lambda only)")
	cmd.Flags().BoolVar(&apiGateway, "api-gateway", false, "Configure API Gateway (Lambda only)")
	cmd.Flags().BoolVar(&listTypes, "list-types", false, "List available project types")

	return cmd
}

func createProject(name, projectType, module string, manifest *config.Manifest, features *config.FeatureConfig) error {
	// Create directories
	for _, dir := range manifest.Structure.Directories {
		dirPath := filepath.Join(name, strings.ReplaceAll(dir, "{project_name}", name))
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return err
		}
	}

	// Create infrastructure directories if database is configured
	if features.Database != "none" && features.Database != "" {
		infraDir := filepath.Join(name, "infrastructure", "database", features.Database)
		if err := os.MkdirAll(infraDir, 0755); err != nil {
			return err
		}
	}

	// Template data
	data := map[string]any{
		"ProjectName": name,
		"Module":      module,
		"Features":    features,
	}

	// Generate basic project files
	projectTemplatesPath := fmt.Sprintf("project_types/%s/project", projectType)
	projectTemplates := map[string]string{
		fmt.Sprintf("%s/main.go.tpl", projectTemplatesPath): filepath.Join(name, "cmd", name, "main.go"),
		fmt.Sprintf("%s/go.mod.tpl", projectTemplatesPath):  filepath.Join(name, "go.mod"),
	}

	// Add optional templates
	if features.Docker && projectType == "lambda" {
		projectTemplates[fmt.Sprintf("%s/Dockerfile.tpl", projectTemplatesPath)] = filepath.Join(name, "Dockerfile")
	}

	if features.SAMTemplate && projectType == "lambda" {
		projectTemplates[fmt.Sprintf("%s/template.yaml.tpl", projectTemplatesPath)] = filepath.Join(name, "template.yaml")
	}

	// Write templates
	for tpl, out := range projectTemplates {
		if err := internal.WriteTemplate(templates.FS, tpl, out, data); err != nil {
			// Skip if template doesn't exist (optional templates)
			continue
		}
	}

	// Generate error management if enabled
	if features.ErrorHandling {
		errorsPath := filepath.Join(name, "internal", "domain", "errors")
		if err := os.MkdirAll(errorsPath, 0755); err != nil {
			return err
		}

		if err := internal.WriteTemplate(templates.FS, "common/errors/errors.go.tpl",
			filepath.Join(errorsPath, "errors.go"), data); err != nil {
			return err
		}

		if err := internal.WriteTemplate(templates.FS, "common/errors/error_types.go.tpl",
			filepath.Join(errorsPath, "error_types.go"), data); err != nil {
			return err
		}
	}

	// Create .sazerac.yaml config file
	sazeracConfig := fmt.Sprintf(`project:
  name: "%s"
  type: "%s"  
  module: "%s"
  version: "1.0.0"

features:
  database: "%s"
  tests: %v
  error_handling: %v
`, name, projectType, module, features.Database, features.Tests, features.ErrorHandling)

	if projectType == "lambda" {
		sazeracConfig += fmt.Sprintf(`  docker: %v
  sam_template: %v
  api_gateway: %v
`, features.Docker, features.SAMTemplate, features.APIGateway)
	}

	if err := os.WriteFile(filepath.Join(name, ".sazerac.yaml"), []byte(sazeracConfig), 0644); err != nil {
		return err
	}
	// Create README
	readme := fmt.Sprintf("# %s\n\nProject created with Sazerac 🥃\n\nType: %s\n", name, projectType)
	return os.WriteFile(filepath.Join(name, "README.md"), []byte(readme), 0644)
}
