# Documentación del Servicio de Usuarios (@user-service)

## Descripción General

Este proyecto es un microservicio de gestión de usuarios desarrollado en Go, utilizando el framework Echo para la creación de APIs RESTful. El servicio proporciona funcionalidades básicas para el registro, autenticación y gestión de usuarios.

## Estructura del Proyecto

El proyecto sigue una estructura modular típica de aplicaciones Go:

- `config/`: Contiene la configuración de la aplicación.
- `db/`: Maneja la conexión y operaciones con la base de datos.
- `handlers/`: Contiene los manejadores de las rutas HTTP.
- `middleware/`: Incluye middlewares personalizados.
- `models/`: Define las estructuras de datos utilizadas en la aplicación.
- `main.go`: Punto de entrada de la aplicación.

## Configuración

La configuración se maneja a través de variables de entorno y un archivo `.env` para entornos de desarrollo. Las principales configuraciones incluyen:

- `DB_CONNECTION_STRING`: Cadena de conexión a la base de datos PostgreSQL.
- `APP_ENV`: Entorno de la aplicación (development, production, etc.).

## Rutas Principales

1. Registro de Usuario:
   - Ruta: `POST /register`
   - Funcionalidad: Permite registrar un nuevo usuario en el sistema.

2. Inicio de Sesión:
   - Ruta: `POST /login`
   - Funcionalidad: Autentica a un usuario y devuelve un token de acceso.

3. Listar Usuarios:
   - Ruta: `GET /users`
   - Funcionalidad: Obtiene una lista de todos los usuarios registrados.

4. Validar Token:
   - Ruta: `GET /validate`
   - Funcionalidad: Valida el token de un usuario y devuelve la información del usuario.

## Modelos de Datos

El modelo principal es `User`, que incluye campos como:


```5:17:user-service/models/user.go
type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password,omitempty"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	UserToken string `json:"user_token,omitempty"`
}
```


## Seguridad

- Las contraseñas se almacenan cifradas utilizando bcrypt.
- Se implementa un sistema de tokens para la autenticación.
- Se utiliza un middleware de autenticación para proteger rutas específicas.

## Despliegue

El proyecto incluye configuraciones para despliegue utilizando Docker y Docker Compose. Los archivos relevantes son:

- `Dockerfile`: Para la construcción de la imagen Docker del servicio.
- `docker-compose.yml`: Para orquestar el servicio junto con la base de datos PostgreSQL.

## Desarrollo y Pruebas

Para el desarrollo local:

1. Clonar el repositorio.
2. Copiar el archivo `.env.example` a `.env` y configurar las variables de entorno.
3. Ejecutar `go mod download` para instalar las dependencias.
4. Usar `go run main.go` para iniciar el servidor de desarrollo.

Para pruebas, se incluye un flujo de CI/CD en GitHub Actions que ejecuta pruebas automáticas en cada pull request.

## Notas Adicionales

Este proyecto es para fines educativos y demuestra conceptos como:
- Desarrollo de microservicios en Go
- Uso de frameworks web como Echo
- Implementación de autenticación y autorización
- Manejo de bases de datos con GORM
- Configuración de CI/CD con GitHub Actions

Se recomienda revisar y mejorar las prácticas de seguridad antes de utilizar en un entorno de producción real.