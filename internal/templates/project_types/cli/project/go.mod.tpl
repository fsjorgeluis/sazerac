module {{ .Module }}

go 1.21

// For local development, tell Go to use the current directory
replace {{ .Module }} => ./