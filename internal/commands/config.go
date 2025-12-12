package commands

import (
	"fmt"
	"os"

	"github.com/fsjorgeluis/sazerac/internal"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// ProjectConfig represents the .sazerac.yaml configuration
type ProjectConfig struct {
	Project struct {
		Name    string `yaml:"name"`
		Type    string `yaml:"type"`
		Module  string `yaml:"module"`
		Version string `yaml:"version"`
	} `yaml:"project"`
	Features struct {
		Database      string `yaml:"database"`
		Tests         bool   `yaml:"tests"`
		ErrorHandling bool   `yaml:"error_handling"`
		Docker        bool   `yaml:"docker,omitempty"`
		SAMTemplate   bool   `yaml:"sam_template,omitempty"`
		APIGateway    bool   `yaml:"api_gateway,omitempty"`
	} `yaml:"features"`
}

func NewConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage project configuration",
		Long:  "View and manage the project's .sazerac.yaml configuration file",
	}

	cmd.AddCommand(newConfigShowCmd())

	return cmd
}

func newConfigShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Show current project configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Check if .sazerac.yaml exists
			if _, err := os.Stat(".sazerac.yaml"); os.IsNotExist(err) {
				// Try to infer from go.mod
				return showInferredConfig()
			}

			// Load configuration
			data, err := os.ReadFile(".sazerac.yaml")
			if err != nil {
				return fmt.Errorf("failed to read .sazerac.yaml: %w", err)
			}

			var cfg ProjectConfig
			if err := yaml.Unmarshal(data, &cfg); err != nil {
				return fmt.Errorf("failed to parse .sazerac.yaml: %w", err)
			}

			// Display configuration
			fmt.Println("📋 Project Configuration")
			fmt.Println()
			fmt.Printf("Project:\n")
			fmt.Printf("  Name:    %s\n", cfg.Project.Name)
			fmt.Printf("  Type:    %s\n", cfg.Project.Type)
			fmt.Printf("  Module:  %s\n", cfg.Project.Module)
			fmt.Printf("  Version: %s\n", cfg.Project.Version)
			fmt.Println()
			fmt.Printf("Features:\n")
			fmt.Printf("  Database:        %s\n", cfg.Features.Database)
			fmt.Printf("  Tests:           %v\n", cfg.Features.Tests)
			fmt.Printf("  Error Handling:  %v\n", cfg.Features.ErrorHandling)

			if cfg.Project.Type == "lambda" {
				fmt.Printf("  Docker:          %v\n", cfg.Features.Docker)
				fmt.Printf("  SAM Template:    %v\n", cfg.Features.SAMTemplate)
				fmt.Printf("  API Gateway:     %v\n", cfg.Features.APIGateway)
			}

			return nil
		},
	}
}

func showInferredConfig() error {
	// Try to infer from go.mod
	moduleName := internal.GetModuleName()
	projectName := internal.GetProjectName()

	if moduleName == "" {
		return fmt.Errorf("no .sazerac.yaml found and could not determine project from go.mod")
	}

	fmt.Println("📋 Inferred Project Configuration (no .sazerac.yaml found)")
	fmt.Println()
	fmt.Printf("Project:\n")
	fmt.Printf("  Name:    %s\n", projectName)
	fmt.Printf("  Module:  %s\n", moduleName)
	fmt.Println()
	fmt.Println("Note: Run 'sazerac init' to create a project configuration file")

	return nil
}
