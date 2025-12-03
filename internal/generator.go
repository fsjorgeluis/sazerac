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
