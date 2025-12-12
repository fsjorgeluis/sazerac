package config

import (
	"embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

// Manifest represents a project type manifest
type Manifest struct {
	Name         string       `yaml:"name"`
	Description  string       `yaml:"description"`
	Version      string       `yaml:"version"`
	Features     Features     `yaml:"features"`
	Dependencies Dependencies `yaml:"dependencies"`
	Structure    Structure    `yaml:"structure"`
}

// Features represents available feature toggles
type Features struct {
	Database      *Feature `yaml:"database,omitempty"`
	Tests         *Feature `yaml:"tests,omitempty"`
	ErrorHandling *Feature `yaml:"error_handling,omitempty"`
	Docker        *Feature `yaml:"docker,omitempty"`
	SAMTemplate   *Feature `yaml:"sam_template,omitempty"`
	APIGateway    *Feature `yaml:"api_gateway,omitempty"`
}

// Feature represents a single feature configuration
type Feature struct {
	Optional bool     `yaml:"optional"`
	Options  []string `yaml:"options,omitempty"`
	Default  any      `yaml:"default"`
	Prompt   string   `yaml:"prompt"`
}

// Dependencies represents project dependencies
type Dependencies struct {
	Required []string            `yaml:"required"`
	Optional map[string][]string `yaml:"optional"`
}

// Structure represents project directory structure
type Structure struct {
	Directories []string `yaml:"directories"`
}

// FeatureConfig holds the selected features for a project
type FeatureConfig struct {
	Database      string
	Tests         bool
	ErrorHandling bool
	Docker        bool
	SAMTemplate   bool
	APIGateway    bool
}

// LoadManifest loads a manifest file from embedded FS
func LoadManifest(fs embed.FS, projectType string) (*Manifest, error) {
	data, err := fs.ReadFile(fmt.Sprintf("project_types/%s/manifest.yaml", projectType))
	if err != nil {
		return nil, fmt.Errorf("failed to read manifest for %s: %w", projectType, err)
	}

	var manifest Manifest
	if err := yaml.Unmarshal(data, &manifest); err != nil {
		return nil, fmt.Errorf("failed to parse manifest for %s: %w", projectType, err)
	}

	return &manifest, nil
}

// ValidateFeatureConfig validates a feature configuration against the manifest
func (m *Manifest) ValidateFeatureConfig(cfg *FeatureConfig) error {
	// Validate database option
	if m.Features.Database != nil && cfg.Database != "" {
		valid := false
		for _, opt := range m.Features.Database.Options {
			if opt == cfg.Database {
				valid = true
				break
			}
		}
		if !valid {
			return fmt.Errorf("invalid database option: %s (valid options: %v)", cfg.Database, m.Features.Database.Options)
		}
	}

	return nil
}

// GetDefaultFeatureConfig returns default feature configuration from manifest
func (m *Manifest) GetDefaultFeatureConfig() *FeatureConfig {
	cfg := &FeatureConfig{}

	if m.Features.Database != nil {
		if db, ok := m.Features.Database.Default.(string); ok {
			cfg.Database = db
		}
	}

	if m.Features.Tests != nil {
		if tests, ok := m.Features.Tests.Default.(bool); ok {
			cfg.Tests = tests
		}
	}

	if m.Features.ErrorHandling != nil {
		if errHandling, ok := m.Features.ErrorHandling.Default.(bool); ok {
			cfg.ErrorHandling = errHandling
		}
	}

	if m.Features.Docker != nil {
		if docker, ok := m.Features.Docker.Default.(bool); ok {
			cfg.Docker = docker
		}
	}

	if m.Features.SAMTemplate != nil {
		if sam, ok := m.Features.SAMTemplate.Default.(bool); ok {
			cfg.SAMTemplate = sam
		}
	}

	if m.Features.APIGateway != nil {
		if apiGw, ok := m.Features.APIGateway.Default.(bool); ok {
			cfg.APIGateway = apiGw
		}
	}

	return cfg
}
