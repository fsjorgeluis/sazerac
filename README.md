# Sazerac 🥃

The CLI for clean architecture enthusiasts

Sazerac es una herramienta de línea de comandos que facilita la creación de proyectos Go siguiendo los principios de Clean Architecture. Genera automáticamente la estructura y los archivos necesarios para entidades, casos de uso, repositorios, handlers, mappers y validadores.

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

## Uso

### Inicializar un nuevo proyecto

Crea un nuevo proyecto con la estructura de Clean Architecture:

```bash
sazerac init mi-proyecto
```

Este comando creará:
- La estructura de directorios básica
- Archivos `main.go`, `go.mod` y `README.md`
- Directorios para entidades, mappers, validadores, casos de uso, repositorios, handlers e infraestructura MySQL

**Nota:** El módulo en `go.mod` se generará como `example.com/<project-name>`. Deberás editarlo para usar tu propio módulo (por ejemplo, `github.com/tu-usuario/mi-proyecto`).

### Generar componentes individuales

#### Entidad (Entity)

Genera una entidad de dominio:

```bash
sazerac make entity User
```

Esto creará `internal/domain/entities/user.go` con una estructura básica.

#### Repositorio (Repository)

Genera la interfaz del repositorio y su implementación MySQL:

```bash
sazerac make repo User
```

Esto generará:
- `internal/repository/user_repository.go` (interfaz)
- `infrastructure/database/mysql/user_mysql.go` (implementación MySQL)

#### Caso de Uso (UseCase)

Genera un caso de uso:

```bash
sazerac make usecase CreateUser User
```

El primer argumento es el nombre del caso de uso y el segundo es la entidad relacionada. Esto creará `internal/usecases/create_user_usecase.go`.

#### Handler

Genera un handler para ejecutar un caso de uso:

```bash
sazerac make handler CreateUser CreateUser
```

El primer argumento es el nombre del handler y el segundo es el nombre del caso de uso. Esto creará `internal/handlers/create_user_handler.go` con un método `Run()` que ejecuta el caso de uso y muestra el resultado.

#### Mapper

Genera un mapper para convertir entre entidades y DTOs:

```bash
sazerac make mapper User
```

Esto creará `internal/domain/mappers/user_mapper.go`.

#### Validator

Genera un validador para una entidad:

```bash
sazerac make validator User
```

Esto creará `internal/domain/validators/user_validator.go`.

### Generar todo de una vez

Para generar todos los componentes relacionados (entidad, repositorio, caso de uso y handler) en un solo comando:

```bash
sazerac make all User CreateUser
```

El primer argumento es el nombre de la entidad y el segundo es el nombre del caso de uso. Este comando ejecutará automáticamente:
1. `make entity` para la entidad
2. `make repo` para el repositorio
3. `make usecase` para el caso de uso (genera entidades con nombres aleatorios)
4. `make handler` para el handler
5. `make di` para el contenedor de dependency injection
6. Actualización de `main.go` que ejecuta el handler directamente

**Nota:** Después de generar los componentes, puedes ejecutar el proyecto con `go run cmd/<project-name>/main.go` y verás un mensaje con la entidad creada.

## Convenciones de nombres

Sazerac convierte automáticamente los nombres a formato snake_case para los archivos:
- `CreateUser` → `create_user`
- `UserProfile` → `user_profile`
- `OrderItem` → `order_item`

Los nombres de las estructuras y tipos se mantienen en PascalCase como los proporcionaste.

## Ejemplo completo

Aquí tienes un ejemplo de cómo crear un módulo completo para gestionar usuarios:

```bash
# 1. Inicializar el proyecto
sazerac init mi-api

# 2. Navegar al proyecto
cd mi-api

# 3. Generar todos los componentes para el módulo de usuarios
sazerac make all User CreateUser

# 4. Ejecutar el proyecto para verificar que funciona
go run cmd/mi-api/main.go
# Salida esperada:
# Have a good drink! 🥃
# Entity created: ID=1234567890, Name=Alice

# 5. Generar componentes adicionales si es necesario
sazerac make mapper User
sazerac make validator User
```

## Comandos disponibles

| Comando | Descripción | Argumentos |
|---------|-------------|-------------|
| `init <nombre>` | Inicializa un nuevo proyecto | Nombre del proyecto |
| `make entity <Nombre>` | Genera una entidad | Nombre de la entidad |
| `make repo <Entity>` | Genera repositorio e implementación MySQL | Nombre de la entidad |
| `make usecase <Name> <Entity>` | Genera un caso de uso | Nombre del caso de uso, Entidad |
| `make handler <Name> <UseCase>` | Genera un handler con método Run() | Nombre del handler, Caso de uso |
| `make mapper <Entity>` | Genera un mapper | Nombre de la entidad |
| `make validator <Entity>` | Genera un validador | Nombre de la entidad |
| `make di <UseCase> <Entity>` | Genera el contenedor de dependency injection | Caso de uso, Entidad |
| `make all <Entity> <UseCase>` | Genera todos los componentes básicos | Entidad, Caso de uso |

## Requisitos

- Go 1.16 o superior (para soporte de `embed.FS`)

## Contribuir

Las contribuciones son bienvenidas. Por favor, abre un issue o envía un pull request.

## Licencia

Ver el archivo LICENSE para más detalles.
