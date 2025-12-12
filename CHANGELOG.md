# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added - Major MVP Release 🥃
- **Multi-Project Template System**: Support for CLI and AWS Lambda project types
- **Interactive CLI Experience**: Interactive prompts using survey/v2 library with automatic flag fallback
- **Project Configuration System**: `.sazerac.yaml` manifest for project metadata and feature configuration
- **Feature Toggles**: Granular control over database, tests, error handling, Docker, SAM template, and API Gateway
- **Database Support**:
  - None (in-memory repository)
  - MySQL/MySQL-RDS
  - DynamoDB
- **Lambda Support**:
  - AWS Lambda handler templates
  - DynamoDB repository implementation
  - SAM template generation (optional)
  - API Gateway integration (optional)
  - Dockerfile for Lambda (optional)
- **In-Memory Repository**: Automatic in-memory repository generation for projects without database
- **Error Management System**: Standardized domain error types with HTTP status codes
- **Modular Template Structure**:
  - `common/`: Shared templates (entity, mapper, validator, errors)
  - `project_types/`: Project-specific templates (cli, lambda)
  - `infrastructure/`: Database implementations (mysql, dynamodb, inmemory)
- **Config Command**: `sazerac config show` to display current project configuration
- **Context Support**: All repository methods now use `context.Context` for better cancellation and timeout control
- **Auto Replace Directive**: `go.mod` templates include `replace` directive for local development

### Changed
- **Template Organization**: Complete restructure from flat to hierarchical template directory
- **Init Command**: Now supports interactive mode with flags for non-interactive use
- **Make Repo Command**: Intelligently selects database implementation based on project configuration
- **Make Commands**: All commands now detect project type and use appropriate templates
- **DI Generation**: Conditional dependency injection based on project type and features
- **Project Initialization**: Generates `.sazerac.yaml` with project metadata

### Fixed
- Import path issues in generated code (now uses `{{ .Module }}` consistently)
- Template function availability (added `ToLower` to template func map)
- Repository interface signatures now consistent across CLI and Lambda
- Nil pointer dereferences when using in-memory repository
- Context parameter usage in repository implementations

### Testing
- All 10/10 command tests passing
- Coverage: 60.9% (internal), 60.7% (commands)
- Validated CLI projects (with and without DB)
- Validated Lambda projects (with and without DB)

## [0.0.2-beta] - 2025-12-04

### Added
- Added comprehensive test suite for utility functions (`ToSnake`, `ToPascalCase`, `GetModuleName`, `GetProjectName`)
- Added benchmark tests for performance-critical functions (`BenchmarkToSnake`, `BenchmarkToPascalCase`)
- Added comprehensive test suite for all commands (`init`, `make entity`, `make repo`, `make usecase`, `make handler`, `make mapper`, `make validator`, `make di`, `make all`)
- Tests verify command creation, execution, file generation, and argument validation
- Tests use temporary directories to avoid affecting the project structure
- Achieved 87.6% code coverage for commands package and 60.9% for internal package
- Added "Development" section to README with testing instructions
- Added "Características" (Features) section to README
- Enhanced "Contribuir" (Contributing) section with guidelines
- Added Clean Architecture diagram to README

### Changed
- Updated README.md with comprehensive documentation including features, development guide, and contributing guidelines
- Enhanced CHANGELOG.md with detailed information about all changes and improvements

## [0.0.1-beta] - 2025-12-03

### Added
- Initial project structure with Clean Architecture support
- `init` command to bootstrap new projects with proper directory structure
- `make entity` command to generate domain entities
- `make repo` command to generate repository interfaces and MySQL implementations
- `make usecase` command to generate use cases
- `make handler` command to generate HTTP handlers
- `make mapper` command to generate entity-to-DTO mappers
- `make validator` command to generate validators
- `make di` command to generate dependency injection container
- `make all` command to generate all components (entity, repo, usecase, handler, di) in a single operation
- Template system with embedded filesystem for all component types
- Automatic snake_case conversion for file names while preserving PascalCase for types
- Module name detection from `go.mod` file
- Dependency injection (DI) container generation
- `main.go` template that initializes DI container and executes handler directly (no HTTP server)
- `di.go` template that follows Clean Architecture pattern (DB → Repository → UseCase → Handler)
- `GetProjectName()` helper function to extract project name from module path
- `ToPascalCase()` helper function to ensure exported types have correct capitalization
- UseCase templates now generate entities with random names for demonstration (Alice, Bob, Charlie, etc.)
- Handler templates now include `Run()` method that executes use case and displays results
- Comprehensive documentation in README.md

### Fixed
- Fixed invalid module path in `init` command (changed from `github.com/<UserName>/...` to `github.com/user-name/...` to avoid syntax errors with invalid characters)
- Fixed command structure: reorganized all `make` commands as subcommands under a parent `make` command
- Fixed `make all` command to properly pass arguments to subcommands (usecase and handler now receive correct arguments)
- Fixed unused imports in templates (`entities` in usecase, `repository` in di, `fmt` in di)
- Fixed template indentation issues in handler template
- Fixed type export issues: all generated types now use PascalCase for proper export
- Fixed UseCase template to return entities instead of strings, properly using the repository
- Fixed handler template to display entity information correctly

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

## [0.0.1] - 2025-07-17

### Added
- Initial commit
- Basic project setup
- Go module initialization

[Unreleased]: https://github.com/fsjorgeluis/sazerac/compare/v0.0.2-beta...HEAD
[0.0.2-beta]: https://github.com/fsjorgeluis/sazerac/compare/v0.0.1-beta...v0.0.2-beta
[0.0.1-beta]: https://github.com/fsjorgeluis/sazerac/compare/v0.0.1...v0.0.1-beta
[0.0.1]: https://github.com/fsjorgeluis/sazerac/releases/tag/v0.0.1
