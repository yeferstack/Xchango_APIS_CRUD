# Xchango_APIS_CRUD
Xchango es una API CRUD en desarrollo creada para gestionar la información de la plataforma, conectando la base de datos SQL con el sistema mediante operaciones de creación, consulta, actualización y eliminación de datos.

# Xchango API CRUD Sistema

API REST desarrollada en **Go** con el framework **Beego**, conectada a **PostgreSQL**. Gestiona el módulo de administración del sistema Xchango, incluyendo administradores, permisos, niveles de acceso, notificaciones e historial de acciones.

---

## Tecnologías

- Go 1.26
- Beego v2.1.0
- PostgreSQL
- Swagger (documentación automática)

---

## Requisitos previos

- Go instalado (v1.26+)
- PostgreSQL corriendo
- Base de datos `xchango_db` creada con el schema `sistema`

---



## Instalación y ejecución

```bash
# Clonar el repositorio
git clone <url-del-repo>
cd Xchango_APIS_CRUD

# Instalar dependencias
go mod tidy

# Ejecutar migraciones (crear tablas e insertar datos iniciales)
bee migrate

# Iniciar el servidor
go run main.go
```

El servidor corre en: `http://localhost:8081`

---

## Documentación Swagger

Disponible en modo `dev` en:

```
http://localhost:8081/swagger/
```

---

## Endpoints

Todos los endpoints tienen el prefijo `/v1`.

Cada recurso expone las operaciones CRUD estándar:

| Método | Ruta | Descripción |
|--------|------|-------------|
| `POST` | `/v1/{recurso}/` | Crear registro |
| `GET` | `/v1/{recurso}/` | Listar todos |
| `GET` | `/v1/{recurso}/:id` | Obtener uno por ID |
| `PUT` | `/v1/{recurso}/:id` | Actualizar por ID |
| `DELETE` | `/v1/{recurso}/:id` | Eliminar por ID |

### Recursos disponibles

| Recurso | Ruta base |
|---------|-----------|
| Nivel de Acceso | `/v1/nivel_acceso` |
| Permiso | `/v1/permiso` |
| Administrador | `/v1/administrador` |
| Administrador Permiso | `/v1/administradorpermiso` |
| Historial Admin | `/v1/historialadmin` |
| Notificación | `/v1/notificacion` |

### Parámetros de consulta para GetAll

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| `limit` | int | Cantidad de registros (default: 10) |
| `offset` | int | Desplazamiento (default: 0) |
| `sortby` | string | Campo(s) para ordenar. Ej: `nombre` |
| `order` | string | `asc` o `desc` |
| `fields` | string | Campos a retornar. Ej: `nombre,activo` |
| `query` | string | Filtros. Ej: `activo:true` |

---

## Estructura del proyecto

```
Xchango_APIS_CRUD/
├── conf/
│   └── app.conf           # Configuración de Beego
├── controllers/           # Lógica de los endpoints
├── models/                # Modelos y acceso a base de datos
├── routers/               # Definición de rutas
├── database/
│   ├── migrations/        # Migraciones en Go
│   └── scripst/           # Scripts SQL (up/down)
├── swagger/               # Archivos estáticos de Swagger UI
├── .env                   # Variables de entorno (no subir al repo)
├── go.mod
└── main.go
```

---

## Base de datos

El schema `sistema` contiene las siguientes tablas:

- `nivel_acceso` — Niveles de acceso del sistema
- `permiso` — Permisos disponibles
- `administrador` — Administradores vinculados a un usuario y nivel de acceso
- `administradorpermiso` — Relación entre administradores y permisos
- `historialadmin` — Registro de acciones realizadas por administradores
- `notificacion` — Notificaciones asociadas a usuarios

---

## CORS

La API tiene CORS habilitado para todos los orígenes, aceptando los métodos `GET`, `POST`, `PUT`, `PATCH`, `DELETE` y `OPTIONS`.
