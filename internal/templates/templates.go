package templates

import "embed"

//go:embed project/* entity/* usecase/* repository/* handler/* validator/* mapper/*
var FS embed.FS
