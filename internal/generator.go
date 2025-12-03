package internal

import (
	"embed"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func WriteTemplate(baseFS embed.FS, tplPath, outPath string, data any) error {
	content, err := baseFS.ReadFile(tplPath)
	if err != nil {
		return err
	}

	tpl, err := template.New(filepath.Base(tplPath)).Parse(string(content))
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

// GetProjectName extracts project name from go.mod module path
func GetProjectName() string {
	module := GetModuleName()
	if module == "" {
		return ""
	}
	
	// Extract project name from module path (e.g., github.com/user/project-name -> project-name)
	parts := strings.Split(module, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	
	return ""
}

//package internal
//
//import (
//	"embed"
//	"os"
//	"path/filepath"
//	"strings"
//	"text/template"
//)
//
//func WriteTemplate(baseFS embed.FS, tplPath, outPath string, data any) error {
//	content, err := baseFS.ReadFile(tplPath)
//	if err != nil {
//		return err
//	}
//
//	tpl, err := template.New(filepath.Base(tplPath)).Parse(string(content))
//	if err != nil {
//		return err
//	}
//
//	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
//		return err
//	}
//
//	f, err := os.Create(outPath)
//	if err != nil {
//		return err
//	}
//	defer func(f *os.File) {
//		err := f.Close()
//		if err != nil {
//			panic(err)
//		}
//	}(f)
//
//	return tpl.Execute(f, data)
//}
//
//func ToSnake(name string) string {
//	var out []rune
//	for i, r := range name {
//		if i > 0 && r >= 'A' && r <= 'Z' {
//			out = append(out, '_')
//		}
//		out = append(out, r)
//	}
//	return strings.ToLower(string(out))
//}
//
//func GetModuleName() string {
//	content, err := os.ReadFile("go.mod")
//	if err != nil {
//		return ""
//	}
//
//	lines := strings.Split(string(content), "\n")
//	if len(lines) > 0 {
//		fields := strings.Fields(lines[0])
//		if len(fields) == 2 {
//			return fields[1]
//		}
//	}
//
//	return ""
//}
