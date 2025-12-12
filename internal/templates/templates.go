package templates

import "embed"

//go:embed common/* project_types/*/* infrastructure/mysql/* infrastructure/dynamodb/* infrastructure/inmemory/*
var FS embed.FS
