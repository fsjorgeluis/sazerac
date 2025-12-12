package prompts

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fsjorgeluis/sazerac/internal/config"
)

// ProjectTypePrompt prompts for project type selection
func ProjectTypePrompt() (string, error) {
	var projectType string
	prompt := &survey.Select{
		Message: "Select project type:",
		Options: []string{"cli", "lambda"},
		Default: "cli",
	}
	if err := survey.AskOne(prompt, &projectType); err != nil {
		return "", err
	}
	return projectType, nil
}

// ProjectNamePrompt prompts for project name
func ProjectNamePrompt() (string, error) {
	var name string
	prompt := &survey.Input{
		Message: "Project name:",
	}
	if err := survey.AskOne(prompt, &name, survey.WithValidator(survey.Required)); err != nil {
		return "", err
	}
	return name, nil
}

// ModulePathPrompt prompts for module path
func ModulePathPrompt(defaultModule string) (string, error) {
	var module string
	prompt := &survey.Input{
		Message: "Module path:",
		Default: defaultModule,
	}
	if err := survey.AskOne(prompt, &module, survey.WithValidator(survey.Required)); err != nil {
		return "", err
	}
	return module, nil
}

// DatabasePrompt prompts for database selection based on manifest
func DatabasePrompt(manifest *config.Manifest) (string, error) {
	if manifest.Features.Database == nil {
		return "none", nil
	}

	var database string
	prompt := &survey.Select{
		Message: manifest.Features.Database.Prompt,
		Options: manifest.Features.Database.Options,
		Default: manifest.Features.Database.Default,
	}
	if err := survey.AskOne(prompt, &database); err != nil {
		return "", err
	}
	return database, nil
}

// BooleanPrompt prompts for a yes/no question
func BooleanPrompt(message string, defaultValue bool) (bool, error) {
	var result bool
	prompt := &survey.Confirm{
		Message: message,
		Default: defaultValue,
	}
	if err := survey.AskOne(prompt, &result); err != nil {
		return false, err
	}
	return result, nil
}

// FeaturePrompts prompts for all features based on manifest
func FeaturePrompts(manifest *config.Manifest) (*config.FeatureConfig, error) {
	cfg := &config.FeatureConfig{}

	// Database
	if manifest.Features.Database != nil && manifest.Features.Database.Optional {
		db, err := DatabasePrompt(manifest)
		if err != nil {
			return nil, err
		}
		cfg.Database = db
	}

	// Tests
	if manifest.Features.Tests != nil && manifest.Features.Tests.Optional {
		tests, err := BooleanPrompt(
			manifest.Features.Tests.Prompt,
			manifest.Features.Tests.Default.(bool),
		)
		if err != nil {
			return nil, err
		}
		cfg.Tests = tests
	}

	// Error Handling
	if manifest.Features.ErrorHandling != nil && manifest.Features.ErrorHandling.Optional {
		errHandling, err := BooleanPrompt(
			manifest.Features.ErrorHandling.Prompt,
			manifest.Features.ErrorHandling.Default.(bool),
		)
		if err != nil {
			return nil, err
		}
		cfg.ErrorHandling = errHandling
	}

	// Docker (Lambda only)
	if manifest.Features.Docker != nil && manifest.Features.Docker.Optional {
		docker, err := BooleanPrompt(
			manifest.Features.Docker.Prompt,
			manifest.Features.Docker.Default.(bool),
		)
		if err != nil {
			return nil, err
		}
		cfg.Docker = docker
	}

	// SAM Template (Lambda only)
	if manifest.Features.SAMTemplate != nil && manifest.Features.SAMTemplate.Optional {
		sam, err := BooleanPrompt(
			manifest.Features.SAMTemplate.Prompt,
			manifest.Features.SAMTemplate.Default.(bool),
		)
		if err != nil {
			return nil, err
		}
		cfg.SAMTemplate = sam
	}

	// API Gateway (Lambda only)
	if manifest.Features.APIGateway != nil && manifest.Features.APIGateway.Optional {
		apiGw, err := BooleanPrompt(
			manifest.Features.APIGateway.Prompt,
			manifest.Features.APIGateway.Default.(bool),
		)
		if err != nil {
			return nil, err
		}
		cfg.APIGateway = apiGw
	}

	return cfg, nil
}

// EntityNamePrompt prompts for entity name
func EntityNamePrompt() (string, error) {
	var name string
	prompt := &survey.Input{
		Message: "Entity name:",
	}
	if err := survey.AskOne(prompt, &name, survey.WithValidator(survey.Required)); err != nil {
		return "", err
	}
	return name, nil
}

// UseCaseNamePrompt prompts for use case name
func UseCaseNamePrompt() (string, error) {
	var name string
	prompt := &survey.Input{
		Message: "UseCase name:",
	}
	if err := survey.AskOne(prompt, &name, survey.WithValidator(survey.Required)); err != nil {
		return "", err
	}
	return name, nil
}

// RepositoryImplementationPrompt prompts for repository implementation
func RepositoryImplementationPrompt(options []string, defaultImpl string) (string, error) {
	if len(options) == 0 {
		return "none", nil
	}

	var impl string
	prompt := &survey.Select{
		Message: "Repository implementation:",
		Options: options,
		Default: defaultImpl,
	}
	if err := survey.AskOne(prompt, &impl); err != nil {
		return "", err
	}
	return impl, nil
}

// ConfirmationPrompt asks for confirmation
func ConfirmationPrompt(message string) (bool, error) {
	var confirmed bool
	prompt := &survey.Confirm{
		Message: message,
		Default: false,
	}
	if err := survey.AskOne(prompt, &confirmed); err != nil {
		return false, err
	}
	return confirmed, nil
}

// DisplaySummary shows a summary of selections
func DisplaySummary(projectName, projectType, module, database string, features *config.FeatureConfig) {
	fmt.Println("\n📋 Project Configuration Summary:")
	fmt.Printf("  Project Name: %s\n", projectName)
	fmt.Printf("  Project Type: %s\n", projectType)
	fmt.Printf("  Module: %s\n", module)
	fmt.Printf("  Database: %s\n", database)

	if features != nil {
		fmt.Println("\n  Features:")
		fmt.Printf("    Error Handling: %v\n", features.ErrorHandling)
		fmt.Printf("    Tests: %v\n", features.Tests)

		if projectType == "lambda" {
			fmt.Printf("    Docker: %v\n", features.Docker)
			fmt.Printf("    SAM Template: %v\n", features.SAMTemplate)
			fmt.Printf("    API Gateway: %v\n", features.APIGateway)
		}
	}
	fmt.Println()
}
