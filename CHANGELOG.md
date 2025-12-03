# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Added dependency injection (DI) container generation with `make di` command
- Added automatic DI generation in `make all` command
- Added `main.go` template that initializes DI container and executes handler directly (no HTTP server)
- Added `di.go` template that follows Clean Architecture pattern (DB -> Repository -> UseCase -> Handler)
- Added `GetProjectName()` helper function to extract project name from module path
- Added `ToPascalCase()` helper function to ensure exported types have correct capitalization
- UseCase templates now generate entities with random names for demonstration
- Handler templates now include `Run()` method that executes use case and displays results

### Fixed
- Fixed invalid module path in `init` command (changed from `github.com/<UserName>/...` to `example.com/...` to avoid syntax errors)
- Fixed command structure: reorganized all `make` commands as subcommands under a parent `make` command
- Fixed `make all` command to properly pass arguments to subcommands (usecase and handler now receive correct arguments)
- Fixed unused imports in templates (`entities` in usecase, `repository` in di, `fmt` in di)
- Fixed template indentation issues in handler template
- Fixed type export issues: all generated types now use PascalCase for proper export

### Changed
- Reorganized CLI structure: all generator commands are now under `sazerac make` (e.g., `sazerac make entity`, `sazerac make all`)
- Improved command help output: commands now display correctly with proper descriptions
- Refactored all command files to use centralized `templates.FS` from `internal/templates` package
- Improved code organization by removing duplicate `embed.FS` declarations
- Updated `init` command to create correct directory structure matching all generator commands
- UseCase templates now return entities instead of strings, demonstrating full Clean Architecture flow
- Handler templates changed from HTTP handlers to console handlers with `Run()` method
- Main.go template changed from HTTP server to direct handler execution
- Projects generated can now be executed directly without additional code

## [0.1.0] - 2025-12-03

### Added
- Initial project structure with Clean Architecture support
- `init` command to bootstrap new projects with proper directory structure
- `make entity` command to generate domain entities
- `make repo` command to generate repository interfaces and MySQL implementations
- `make usecase` command to generate use cases
- `make handler` command to generate HTTP handlers
- `make mapper` command to generate entity-to-DTO mappers
- `make validator` command to generate validators
- `all` command to generate all components (entity, repo, usecase, handler) in a single operation
- Template system with embedded filesystem for all component types
- Automatic snake_case conversion for file names while preserving PascalCase for types
- Module name detection from `go.mod` file
- Comprehensive documentation in README.md

### Changed
- Project initialization now creates a complete Clean Architecture structure

## [0.0.1] - 2025-07-03

### Added
- Initial commit
- Basic project setup
- Go module initialization

[Unreleased]: https://github.com/fsjorgeluis/sazerac/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/fsjorgeluis/sazerac/compare/v0.0.1...v0.1.0
[0.0.1]: https://github.com/fsjorgeluis/sazerac/releases/tag/v0.0.1

