package internal

import (
	"embed"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

func WriteTemplate(baseFS embed.FS, tplPath, outPath string, data any) error {
	content, err := baseFS.ReadFile(tplPath)
	if err != nil {
		return err
	}

	// Create template with custom functions
	tpl, err := template.New(filepath.Base(tplPath)).Funcs(template.FuncMap{
		"ToSnake":      ToSnake,
		"ToPascalCase": ToPascalCase,
		"ToLower":      strings.ToLower,
	}).Parse(string(content))
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return err
	}

	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	return tpl.Execute(f, data)
}

func ToSnake(name string) string {
	var out []rune
	for i, r := range name {
		if i > 0 && r >= 'A' && r <= 'Z' {
			out = append(out, '_')
		}
		out = append(out, r)
	}
	return strings.ToLower(string(out))
}

// ToPascalCase converts a string to PascalCase (first letter uppercase)
// If the string already starts with uppercase, it returns it as is
func ToPascalCase(name string) string {
	if name == "" {
		return name
	}
	// If already starts with uppercase, return as is
	runes := []rune(name)
	if len(runes) > 0 && runes[0] >= 'A' && runes[0] <= 'Z' {
		return name
	}
	// Convert first letter to uppercase
	if len(runes) > 0 {
		runes[0] = []rune(strings.ToUpper(string(runes[0])))[0]
	}
	return string(runes)
}

func GetModuleName() string {
	content, err := os.ReadFile("go.mod")
	if err != nil {
		return ""
	}

	lines := strings.Split(string(content), "\n")
	if len(lines) > 0 {
		fields := strings.Fields(lines[0])
		if len(fields) == 2 {
			return fields[1]
		}
	}

	return ""
}

// GetProjectName extracts the project name from .sazerac.yaml or go.mod
func GetProjectName() string {
	// First try to read from .sazerac.yaml
	if data, err := os.ReadFile(".sazerac.yaml"); err == nil {
		var config struct {
			Project struct {
				Name string `yaml:"name"`
			} `yaml:"project"`
		}
		if err := yaml.Unmarshal(data, &config); err == nil && config.Project.Name != "" {
			return config.Project.Name
		}
	}

	// Fallback to extracting from go.mod
	data, err := os.ReadFile("go.mod")
	if err != nil {
		return ""
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "module ") {
			modulePath := strings.TrimPrefix(line, "module ")
			modulePath = strings.TrimSpace(modulePath)
			parts := strings.Split(modulePath, "/")
			return parts[len(parts)-1]
		}
	}

	return ""
}
