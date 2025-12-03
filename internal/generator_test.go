package internal

import (
	"os"
	"testing"
)

func TestToSnake(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple PascalCase",
			input:    "User",
			expected: "user",
		},
		{
			name:     "Multiple words",
			input:    "CreateUser",
			expected: "create_user",
		},
		{
			name:     "Three words",
			input:    "UserProfile",
			expected: "user_profile",
		},
		{
			name:     "Four words",
			input:    "OrderItemDetail",
			expected: "order_item_detail",
		},
		{
			name:     "Already lowercase",
			input:    "user",
			expected: "user",
		},
		{
			name:     "Mixed case",
			input:    "XMLParser",
			expected: "x_m_l_parser",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Single uppercase",
			input:    "A",
			expected: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToSnake(tt.input)
			if result != tt.expected {
				t.Errorf("ToSnake(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToPascalCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Already PascalCase",
			input:    "User",
			expected: "User",
		},
		{
			name:     "Lowercase single word",
			input:    "user",
			expected: "User",
		},
		{
			name:     "Lowercase multiple words",
			input:    "createuser",
			expected: "Createuser",
		},
		{
			name:     "Mixed case",
			input:    "cReAtEuSeR",
			expected: "CReAtEuSeR",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Single lowercase",
			input:    "a",
			expected: "A",
		},
		{
			name:     "Single uppercase",
			input:    "A",
			expected: "A",
		},
		{
			name:     "Numbers",
			input:    "user123",
			expected: "User123",
		},
		{
			name:     "Starts with number",
			input:    "123user",
			expected: "123user",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToPascalCase(tt.input)
			if result != tt.expected {
				t.Errorf("ToPascalCase(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGetModuleName(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}

	// Change to temp directory
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}
	defer os.Chdir(originalDir)

	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:     "Standard module",
			content:  "module github.com/user/project\n\ngo 1.21\n",
			expected: "github.com/user/project",
		},
		{
			name:     "Module with version",
			content:  "module example.com/myapp\n\ngo 1.24.4\n",
			expected: "example.com/myapp",
		},
		{
			name:     "Module with extra spaces",
			content:  "module  github.com/user/project  \n\ngo 1.21\n",
			expected: "github.com/user/project",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Write go.mod file
			if err := os.WriteFile("go.mod", []byte(tt.content), 0644); err != nil {
				t.Fatalf("Failed to write go.mod: %v", err)
			}

			result := GetModuleName()
			if result != tt.expected {
				t.Errorf("GetModuleName() = %q, expected %q", result, tt.expected)
			}

			// Clean up
			os.Remove("go.mod")
		})
	}

	// Test with no go.mod file
	t.Run("No go.mod file", func(t *testing.T) {
		os.Remove("go.mod")
		result := GetModuleName()
		if result != "" {
			t.Errorf("GetModuleName() = %q, expected empty string", result)
		}
	})
}

func TestGetProjectName(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}

	// Change to temp directory
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}
	defer os.Chdir(originalDir)

	tests := []struct {
		name     string
		module   string
		expected string
	}{
		{
			name:     "Standard GitHub path",
			module:   "github.com/user/project-name",
			expected: "project-name",
		},
		{
			name:     "Simple module path",
			module:   "example.com/myapp",
			expected: "myapp",
		},
		{
			name:     "Nested path",
			module:   "github.com/org/team/project",
			expected: "project",
		},
		{
			name:     "Single segment",
			module:   "project",
			expected: "project",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Write go.mod file with the module
			content := "module " + tt.module + "\n\ngo 1.21\n"
			if err := os.WriteFile("go.mod", []byte(content), 0644); err != nil {
				t.Fatalf("Failed to write go.mod: %v", err)
			}

			result := GetProjectName()
			if result != tt.expected {
				t.Errorf("GetProjectName() = %q, expected %q", result, tt.expected)
			}

			// Clean up
			os.Remove("go.mod")
		})
	}

	// Test with no go.mod file
	t.Run("No go.mod file", func(t *testing.T) {
		os.Remove("go.mod")
		result := GetProjectName()
		if result != "" {
			t.Errorf("GetProjectName() = %q, expected empty string", result)
		}
	})
}

func TestToSnake_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Consecutive capitals",
			input:    "XMLHTTPRequest",
			expected: "x_m_l_h_t_t_p_request",
		},
		{
			name:     "Numbers",
			input:    "User123",
			expected: "user123",
		},
		{
			name:     "Special characters",
			input:    "User-Name",
			expected: "user-_name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToSnake(tt.input)
			if result != tt.expected {
				t.Errorf("ToSnake(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToPascalCase_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Unicode characters",
			input:    "ñame",
			expected: "Ñame",
		},
		{
			name:     "Special characters",
			input:    "_user",
			expected: "_user",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToPascalCase(tt.input)
			if result != tt.expected {
				t.Errorf("ToPascalCase(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkToSnake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToSnake("CreateUserProfileHandler")
	}
}

func BenchmarkToPascalCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToPascalCase("createuserprofilehandler")
	}
}

