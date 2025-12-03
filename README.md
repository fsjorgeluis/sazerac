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
- Directorios para entidades, casos de uso, interfaces, repositorios e infraestructura HTTP

**Nota:** Después de inicializar el proyecto, deberás editar el `go.mod` para actualizar el módulo con tu nombre de usuario de GitHub.

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

Genera un handler HTTP para un caso de uso:

```bash
sazerac make handler CreateUser CreateUser
```

El primer argumento es el nombre del handler y el segundo es el nombre del caso de uso. Esto creará `internal/handlers/create_user_handler.go`.

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
sazerac all User CreateUser
```

El primer argumento es el nombre de la entidad y el segundo es el nombre del caso de uso. Este comando ejecutará automáticamente:
1. `make entity` para la entidad
2. `make repo` para el repositorio
3. `make usecase` para el caso de uso
4. `make handler` para el handler

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
sazerac all User CreateUser

# 4. Generar componentes adicionales si es necesario
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
| `make handler <Name> <UseCase>` | Genera un handler HTTP | Nombre del handler, Caso de uso |
| `make mapper <Entity>` | Genera un mapper | Nombre de la entidad |
| `make validator <Entity>` | Genera un validador | Nombre de la entidad |
| `all <Entity> <UseCase>` | Genera todos los componentes básicos | Entidad, Caso de uso |

## Requisitos

- Go 1.16 o superior (para soporte de `embed.FS`)

## Contribuir

Las contribuciones son bienvenidas. Por favor, abre un issue o envía un pull request.

## Licencia

Ver el archivo LICENSE para más detalles.
