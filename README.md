# Sazerac 🥃

The CLI for clean architecture enthusiasts

Sazerac es una herramienta de línea de comandos que facilita la creación de proyectos Go siguiendo los principios de Clean Architecture. Genera automáticamente la estructura y los archivos necesarios para proyectos CLI y AWS Lambda con soporte para múltiples bases de datos.

## Características

- ✅ **Multi-Project Support**: CLI y AWS Lambda
- ✅ **Interactive CLI**: Modo interactivo con prompts inteligentes
- ✅ **Database Options**: None (in-memory), MySQL, DynamoDB
- ✅ **Feature Toggles**: Control granular de características (tests, error handling, SAM, etc.)
- ✅ **Clean Architecture**: Estructura automática siguiendo principios SOLID
- ✅ **Dependency Injection**: Contenedor DI generado automáticamente
- ✅ **Context Support**: Todos los repositorios usan `context.Context`
- ✅ **Error Management**: Sistema de errores de dominio con códigos HTTP
- ✅ **Ready-to-Run**: Proyectos ejecutables sin código adicional
- ✅ **High Test Coverage**: 60.7% coverage con tests automatizados

## Instalación

Para instalar Sazerac, ejecuta:

```bash
go install github.com/fsjorgeluis/sazerac@latest
```

O si estás trabajando con el código fuente localmente:

```bash
go install .
```

Asegúrate de que `$GOPATH/bin` o `$HOME/go/bin` esté en tu `PATH` para poder ejecutar `sazerac` desde cualquier directorio.

## Inicio Rápido

### Modo Interactivo (Recomendado)

```bash
sazerac init
```

El CLI te guiará interactivamente para configurar tu proyecto:
- Tipo de proyecto (CLI o Lambda)
- Nombre del módulo
- Base de datos (none, MySQL, DynamoDB)
- Características opcionales (tests, error handling, Docker, SAM template)

### Modo No Interactivo

```bash
# Proyecto CLI con MySQL
sazerac init my-api --type cli --module github.com/user/my-api --db mysql

# Proyecto Lambda con DynamoDB y SAM
sazerac init my-lambda --type lambda --module github.com/user/my-lambda --db dynamodb --sam --api-gateway

# Proyecto CLI sin base de datos
sazerac init my-cli --type cli --module github.com/user/my-cli --db none
```

### Generar Componentes

```bash
cd my-project

# Generar todos los componentes de una vez
sazerac make all User CreateUser

# O generarlos individualmente
sazerac make entity User
sazerac make repo User
sazerac make usecase CreateUser User
sazerac make handler CreateUser CreateUser
sazerac make di CreateUser User
```

### Ejecutar el Proyecto

**CLI Project:**
```bash
go mod tidy
go run cmd/my-cli/main.go
```

**Lambda Project:**
```bash
go mod tidy

# Build for Lambda
GOOS=linux GOARCH=amd64 go build -o bootstrap cmd/lambda/main.go

# O deploy con SAM (si se generó template.yaml)
sam build && sam deploy --guided
```

## Tipos de Proyecto

### CLI Projects

Proyectos de línea de comandos con:
- Handler con método `Run()`
- Ejecución directa sin servidor HTTP
- Soporte para MySQL o in-memory storage

**Ejemplo de uso:**
```go
// El main.go generado ejecuta directamente el handler
container, _ := di.NewContainer()
container.CreateUserHandler.Run()
```

### Lambda Projects

Proyectos AWS Lambda con:
- Handler compatible con API Gateway
- Integración con DynamoDB o MySQL-RDS
- Templates SAM opcionales
- Dockerfile opcional para despliegue

**Ejemplo de uso:**
```go
// El main.go generado usa Lambda runtime
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    return container.ProcessOrderHandler.Handle(ctx, request)
}
```

## Opciones de Base de Datos

### None (In-Memory)
- Perfecto para prototipos y demos
- Repositorio in-memory con thread-safety
- Sin dependencias externas

### MySQL
- Para proyectos CLI con MySQL
- Incluye implementación completa
- Requiere `database/sql` y driver MySQL

### MySQL-RDS
- Para proyectos Lambda con Amazon RDS
- Configuración via variables de entorno
- Driver incluido en los templates

### DynamoDB
- Para proyectos Lambda serverless
- AWS SDK v2 integrado
- Table name configurable via env vars

## Comandos

### Init

Inicializa un nuevo proyecto:

```bash
# Modo interactivo
sazerac init

# Con flags
sazerac init <nombre> --type <cli|lambda> --module <module-path> --db <none|mysql|dynamodb> [--sam] [--api-gateway] [--skip-tests]
```

**Flags disponibles:**
- `--type`: Tipo de proyecto (cli, lambda)
- `--module`: Ruta del módulo Go
- `--db`: Base de datos (none, mysql, mysql-rds, dynamodb)
- `--sam`: Incluir SAM template (solo Lambda)
- `--api-gateway`: Incluir API Gateway (solo Lambda)
- `--docker`: Incluir Dockerfile (solo Lambda)
- `--skip-tests`: No generar archivos de test

### Config

Muestra la configuración actual del proyecto:

```bash
sazerac config show
```

Esto lee `.sazerac.yaml` o infiere la configuración desde `go.mod`.

### Make Commands

| Comando | Descripción | Argumentos |
|---------|-------------|------------|
| `make entity <Name>` | Genera una entidad de dominio | Nombre de la entidad |
| `make repo <Entity>` | Genera repositorio e implementación | Nombre de la entidad |
| `make usecase <Name> <Entity>` | Genera un caso de uso | Nombre del caso de uso, Entidad |
| `make handler <Name> <UseCase>` | Genera un handler | Nombre del handler, Caso de uso |
| `make mapper <Entity>` | Genera un mapper DTO | Nombre de la entidad |
| `make validator <Entity>` | Genera un validador | Nombre de la entidad |
| `make di <UseCase> <Entity>` | Genera contenedor DI | Caso de uso, Entidad |
| `make all <Entity> <UseCase>` | Genera todo (entity, repo, usecase, handler, DI) | Entidad, Caso de uso |

**Detección automática**: Los comandos `make` detectan automáticamente el tipo de proyecto desde `.sazerac.yaml` y usan los templates apropiados.

## Arquitectura Clean Architecture

Sazerac genera proyectos siguiendo los principios de Clean Architecture:

```
┌─────────────────────────────────────────────────────────────┐
│                        main.go                              │
│  (Punto de entrada de la aplicación)                        │
└───────────────────────┬─────────────────────────────────────┘
                        │
                        ▼
┌─────────────────────────────────────────────────────────────┐
│                    di/di.go                                 │
│  (Dependency Injection Container)                           │
│  - Inicializa database (MySQL/DynamoDB/InMemory)            │
│  - Crea repositories, use cases, handlers                   │
│  - Conecta las capas de la arquitectura                     │
└───────┬─────────────────────────────────────────────────────┘
        │
        ├─────────────────┬──────────────────┬──────────────┐
        ▼                 ▼                  ▼              ▼
┌──────────────┐  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐
│  Handlers    │  │  UseCases    │  │ Repository   │  │  Entities    │
│  (Capa de    │  │  (Lógica de  │  │  (Interfaz)  │  │  (Dominio)   │
│  aplicación) │  │  negocio)    │  │              │  │              │
└──────┬───────┘  └──────┬───────┘  └──────┬───────┘  └──────────────┘
       │                │                  │
       │                │                  │
       │                ▼                  │
       │         ┌──────────────────────┐  │
       │         │  Infrastructure      │  │
       │         │  (MySQL/DynamoDB/    │  │
       │         │   InMemory)          │  │
       │         └──────────────────────┘  │
       │                                   │
       └───────────────────────────────────┘

Flujo de ejecución:
main.go → di.NewContainer()
  ├── database.NewRepo(connection) // MySQL, DynamoDB, o InMemory
  ├── usecases.NewUseCase(repo)
  └── handlers.NewHandler(usecase)
  
handler.Run() → usecase.Execute(ctx) → repository.Save(ctx, entity)
```

### Capas de la Arquitectura

1. **Entities (Dominio)**: Objetos de negocio puros
   - Sin dependencias externas
   - Generados desde `common/entity/entity.go.tpl`

2. **Repository (Interfaz)**: Contratos para acceso a datos
   - Define `Save(ctx, entity)` y `FindByID(ctx, id)`
   - CLI: `project_types/cli/repository/repository.go.tpl`
   - Lambda: `project_types/lambda/repository/repository.go.tpl`

3. **Repository (Implementación)**: Acceso a datos real
   - MySQL: `infrastructure/mysql/repo_mysql.go.tpl`
   - DynamoDB: `infrastructure/dynamodb/repo_dynamodb.go.tpl`
   - InMemory: `infrastructure/inmemory/repo_inmemory.go.tpl`

4. **UseCases (Lógica de negocio)**: Casos de uso
   - CLI: `project_types/cli/usecase/usecase.go.tpl`
   - Lambda: `project_types/lambda/usecase/usecase.go.tpl`

5. **Handlers (Capa de aplicación)**: Orquestación
   - CLI: Handler con método `Run()`
   - Lambda: Handler compatible con API Gateway

6. **DI Container**: Inyección de dependencias
   - Gestiona todas las dependencias
   - Adapta según tipo de proyecto y DB

7. **Error Management** (Opcional): Errores de dominio
   - Códigos HTTP estandarizados
   - Generado desde `common/errors/`

### Principios Aplicados

- **Dependency Rule**: Las dependencias apuntan hacia el dominio
- **Independencia de frameworks**: No acoplamiento a librerías externas
- **Testabilidad**: Cada capa testeada independientemente
- **Independencia de UI**: Lógica de negocio desacoplada
- **Independencia de DB**: El dominio no conoce detalles de persistencia

## Estructura del Proyecto Generado

### CLI Project
```
my-cli/
├── .sazerac.yaml           # Configuración del proyecto
├── go.mod
├── README.md
├── cmd/
│   └── my-cli/
│       ├── main.go         # Punto de entrada
│       └── di/
│           └── di.go       # Dependency injection
├── internal/
│   ├── domain/
│   │   ├── entities/       # Entidades de dominio
│   │   ├── errors/         # Errores personalizados (opcional)
│   │   ├── mappers/        # Mappers DTO (opcional)
│   │   └── validators/     # Validadores (opcional)
│   ├── repository/         # Interfaces de repositorio
│   ├── usecases/           # Casos de uso
│   └── handlers/           # Handlers CLI
└── infrastructure/
    └── database/
        ├── mysql/          # Implementaciones MySQL
        └── inmemory/       # Implementaciones in-memory
```

### Lambda Project
```
my-lambda/
├── .sazerac.yaml
├── go.mod
├── template.yaml           # SAM template (opcional)
├── Dockerfile              # Para despliegue (opcional)
├── cmd/
│   └── lambda/
│       ├── main.go         # Lambda handler
│       └── di/
│           └── di.go
├── internal/
│   ├── domain/
│   │   ├── entities/
│   │   └── errors/
│   ├── repository/
│   ├── usecases/
│   └── handlers/           # Lambda handlers
└── infrastructure/
    └── database/
        ├── dynamodb/       # Implementaciones DynamoDB
        └── inmemory/       # Implementaciones in-memory
```

## Archivo .sazerac.yaml

El archivo `.sazerac.yaml` contiene la metadata del proyecto:

```yaml
project:
  name: "my-project"
  type: "cli"              # o "lambda"
  module: "github.com/user/my-project"
  version: "1.0.0"

features:
  database: "mysql"        # none, mysql, mysql-rds, dynamodb
  tests: true
  error_handling: true
  docker: false            # solo Lambda
  sam_template: false      # solo Lambda
  api_gateway: false       # solo Lambda
```

Este archivo permite a los comandos `make` detectar automáticamente el tipo de proyecto y generar los templates correctos.

## Ejemplo Completo

### Proyecto CLI con MySQL

```bash
# 1. Crear proyecto
sazerac init my-api --type cli --module github.com/user/my-api --db mysql

# 2. Navegar al proyecto
cd my-api

# 3. Generar componentes
sazerac make all User CreateUser

# 4. Instalar dependencias
go mod tidy

# 5. Ejecutar
go run cmd/my-api/main.go

# Salida esperada:
# Have a good drink! 🥃
# Entity created: ID=1670123456, Name=Alice
```

### Proyecto Lambda con DynamoDB

```bash
# 1. Crear proyecto
sazerac init order-service --type lambda --module github.com/user/order-service --db dynamodb --sam --api-gateway

# 2. Navegar al proyecto
cd order-service

# 3. Generar componentes
sazerac make all Order ProcessOrder

# 4. Instalar dependencias
go mod tidy

# 5. Build para Lambda
GOOS=linux GOARCH=amd64 go build -o bootstrap cmd/lambda/main.go

# 6. Deploy con SAM
sam build
sam deploy --guided
```

## Desarrollo

### Ejecutar Tests

```bash
# Todos los tests
go test ./...

# Con cobertura
go test ./... -cover

# Modo verbose
go test ./... -v

# Benchmarks
go test ./internal -bench=. -benchmem
```

### Cobertura de Código

- **internal/commands**: 60.7%
- **internal**: 60.9%

### Estructura del Proyecto Sazerac

```
sazerac/
├── cmd/
│   └── sazerac.go
├── internal/
│   ├── commands/          # Comandos CLI
│   ├── config/            # Config management
│   ├── prompts/           # Interactive prompts
│   ├── templates/         # Templates embebidos
│   │   ├── common/        # Shared templates
│   │   ├── project_types/ # CLI y Lambda
│   │   └── infrastructure/# DB implementations
│   ├── generator.go       # Utilidades
│   └── *_test.go          # Tests
├── go.mod
├── README.md
└── CHANGELOG.md
```

## Requisitos

- Go 1.21 o superior
- Para proyectos Lambda:
  - AWS CLI configurado (para deployment)
  - SAM CLI (opcional, para SAM templates)
  - Docker (opcional, para local testing)

## Contribuir

Las contribuciones son bienvenidas. Por favor:

1. Fork el repositorio
2. Crea una rama feature (`git checkout -b feature/amazing-feature`)
3. Asegúrate de que los tests pasen: `go test ./...`
4. Commit tus cambios (`git commit -m 'Add amazing feature'`)
5. Push a la rama (`git push origin feature/amazing-feature`)
6. Abre un Pull Request

### Guidelines

- Mantén la cobertura de tests arriba del 60%
- Actualiza CHANGELOG.md con tus cambios
- Sigue las convenciones de código existentes
- Agrega tests para nuevas funcionalidades

## Roadmap

- [ ] Gin HTTP project support
- [ ] PostgreSQL repository templates
- [ ] Middleware generation
- [ ] Custom user-defined templates
- [ ] CI/CD pipeline templates
- [ ] OpenAPI/Swagger generation
- [ ] Observability/monitoring templates
- [ ] GraphQL handler templates

## Licencia

Ver el archivo LICENSE para más detalles.

---

**Have a good drink! 🥃**
