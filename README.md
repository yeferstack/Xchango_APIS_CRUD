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


# Xchango_APIS_CRUD_publicaciones
Xchango es una API CRUD en desarrollo creada para gestionar la información de la plataforma, conectando la base de datos SQL con el sistema mediante operaciones de creación, consulta, actualización y eliminación de datos.

# Xchango API CRUD Publicaciones

API REST desarrollada en **Go** con el framework **Beego**, conectada a **PostgreSQL**. Gestiona el módulo de publicaciones de la plataforma Xchango, incluyendo publicaciones, imágenes, categorías, bienes físicos, bienes digitales, servicios y favoritos.

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
- Base de datos `Xchango_db` creada con el schema `Publicaciones`

---

## Instalación y ejecución

```bash
# Clonar el repositorio
git clone <url-del-repo>
cd Xchango_APIS_CRUD_publicaciones

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

| Método   | Ruta                    | Descripción           |
|----------|-------------------------|-----------------------|
| `POST`   | `/v1/{recurso}/`        | Crear registro        |
| `GET`    | `/v1/{recurso}/`        | Listar todos          |
| `GET`    | `/v1/{recurso}/:id`     | Obtener uno por ID    |
| `PUT`    | `/v1/{recurso}/:id`     | Actualizar por ID     |
| `DELETE` | `/v1/{recurso}/:id`     | Eliminar por ID       |

### Recursos disponibles

| Recurso                  | Ruta base                    |
|--------------------------|------------------------------|
| Publicación              | `/v1/publicacion`            |
| Imagen Publicación       | `/v1/imagenpublicacion`      |
| Publicación Categoría    | `/v1/publicacioncategoria`   |
| Sub Categoría            | `/v1/sub_categoria`          |
| Bien Físico              | `/v1/bienfisico`             |
| Servicio                 | `/v1/servicio`               |
| Bien Digital             | `/v1/biendigital`            |
| Favorito                 | `/v1/favorito`               |

### Parámetros de consulta para GetAll

| Parámetro | Tipo   | Descripción                                   |
|-----------|--------|-----------------------------------------------|
| `limit`   | int    | Cantidad de registros (default: 10)           |
| `offset`  | int    | Desplazamiento (default: 0)                   |
| `sortby`  | string | Campo(s) para ordenar. Ej: `titulo`           |
| `order`   | string | `asc` o `desc`                                |
| `fields`  | string | Campos a retornar. Ej: `titulo,activo`        |
| `query`   | string | Filtros. Ej: `activo:true`                    |

---

## Estructura del proyecto

```
Xchango_APIS_CRUD_publicaciones/
├── conf/
│   └── app.conf           # Configuración de Beego
├── controllers/           # Lógica de los endpoints
├── models/                # Modelos y acceso a base de datos
├── routers/               # Definición de rutas
├── database/
│   ├── migrations/        # Migraciones en Go
│   └── scripts/           # Scripts SQL (up/down)
├── swagger/               # Archivos estáticos de Swagger UI
├── .env                   # Variables de entorno (no subir al repo)
├── go.mod
└── main.go
```

---

## Base de datos

El schema `Publicaciones` contiene las siguientes tablas:

- `publicacion` — Publicaciones de los usuarios con información de ubicación, tipo, visibilidad y estado
- `imagenpublicacion` — Imágenes asociadas a cada publicación
- `sub_categoria` — Subcategorías disponibles (electrónico, vehículos, ropa, libros, muebles, juguetes)
- `publicacioncategoria` — Relación entre publicaciones y subcategorías
- `bienFisico` — Detalle de bienes físicos vinculados a una publicación (marca, modelo, color, peso, dimensiones)
- `servicio` — Detalle de servicios vinculados a una publicación (duración, modalidad, disponibilidad, requisitos)
- `biendigital` — Detalle de bienes digitales vinculados a una publicación (tipo de archivo, tamaño, licencia)
- `favorito` — Registro de publicaciones marcadas como favoritas por los usuarios

---

## CORS

La API tiene CORS habilitado para todos los orígenes, aceptando los métodos `GET`, `POST`, `PUT`, `PATCH`, `DELETE` y `OPTIONS`.

# Xchango_APIS_CRUD_Transacciones_comunicacion
Xchango es una API CRUD en desarrollo creada para gestionar la información de la plataforma, conectando la base de datos SQL con el sistema mediante operaciones de creación, consulta, actualización y eliminación de datos.

# Xchango API CRUD Transacciones y Comunicación

API REST desarrollada en **Go** con el framework **Beego**, conectada a **PostgreSQL**. Gestiona el módulo de transacciones y comunicación de la plataforma Xchango, incluyendo intercambios entre usuarios, contacto por WhatsApp, calificaciones y reportes de publicaciones.

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
- Base de datos `Xchango_db` creada con el schema `Transacciones_comunicacion`

---

## Instalación y ejecución

```bash
# Clonar el repositorio
git clone <url-del-repo>
cd Xchango_APIS_CRUD_Transacciones_comunicacion

# Instalar dependencias
go mod tidy

# Ejecutar migraciones (crear tablas e insertar datos iniciales)
bee migrate

# Iniciar el servidor
go run main.go
```

El servidor corre en: `http://localhost:8080`

---

## Documentación Swagger

Disponible en modo `dev` en:

```
http://localhost:8080/swagger/
```

---

## Endpoints

Todos los endpoints tienen el prefijo `/v1`.

Cada recurso expone las operaciones CRUD estándar:

| Método   | Ruta                    | Descripción           |
|----------|-------------------------|-----------------------|
| `POST`   | `/v1/{recurso}/`        | Crear registro        |
| `GET`    | `/v1/{recurso}/`        | Listar todos          |
| `GET`    | `/v1/{recurso}/:id`     | Obtener uno por ID    |
| `PUT`    | `/v1/{recurso}/:id`     | Actualizar por ID     |
| `DELETE` | `/v1/{recurso}/:id`     | Eliminar por ID       |

### Recursos disponibles

| Recurso             | Ruta base                  |
|---------------------|----------------------------|
| Intercambio         | `/v1/Intercambio`          |
| Contacto WhatsApp   | `/v1/ContactoWhatsApp`     |
| Calificación        | `/v1/Calificacion`         |
| Reporte             | `/v1/Reporte`              |

### Parámetros de consulta para GetAll

| Parámetro | Tipo   | Descripción                                        |
|-----------|--------|----------------------------------------------------|
| `limit`   | int    | Cantidad de registros (default: 10)                |
| `offset`  | int    | Desplazamiento (default: 0)                        |
| `sortby`  | string | Campo(s) para ordenar. Ej: `estado`                |
| `order`   | string | `asc` o `desc`                                     |
| `fields`  | string | Campos a retornar. Ej: `estado,activo`             |
| `query`   | string | Filtros. Ej: `activo:true`                         |

---

## Estructura del proyecto

```
Xchango_APIS_CRUD_Transacciones_comunicacion/
├── conf/
│   └── app.conf           # Configuración de Beego
├── controllers/           # Lógica de los endpoints
├── models/                # Modelos y acceso a base de datos
├── routers/               # Definición de rutas
├── database/
│   ├── migrations/        # Migraciones en Go
│   └── scripts/           # Scripts SQL (up/down)
├── swagger/               # Archivos estáticos de Swagger UI
├── .env                   # Variables de entorno (no subir al repo)
├── go.mod
└── main.go
```

---

## Base de datos

El schema `Transacciones_comunicacion` contiene las siguientes tablas:

- `Intercambio` — Registro de solicitudes de trueque entre un usuario interesado y una publicación, con estado del proceso (`iniciado`, etc.) y mensaje inicial
- `ContactoWhatsApp` — Registro de contactos iniciados por WhatsApp entre el usuario interesado y el propietario de una publicación
- `Calificacion` — Evaluaciones entre usuarios tras un intercambio, con puntuación y comentario opcional; vinculada a `Intercambio`
- `Reporte` — Reportes de publicaciones realizados por usuarios, con descripción y estado de revisión (`pendiente`, etc.)

---

## CORS

La API tiene CORS habilitado para todos los orígenes, aceptando los métodos `GET`, `POST`, `PUT`, `PATCH`, `DELETE` y `OPTIONS`.

# Xchango_APIS_CRUD_Usuarioseguridad
Xchango es una API CRUD en desarrollo creada para gestionar la información de la plataforma, conectando la base de datos SQL con el sistema mediante operaciones de creación, consulta, actualización y eliminación de datos.

# Xchango API CRUD Usuario y Seguridad

API REST desarrollada en **Go** con el framework **Beego**, conectada a **PostgreSQL**. Gestiona el módulo de usuarios y seguridad de la plataforma Xchango, incluyendo registro de usuarios, credenciales, sesiones, intentos de login, recuperación de contraseña y perfiles.

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
- Base de datos `Xchango_db` creada con el schema `usuario_seguridad`

---

## Instalación y ejecución

```bash
# Clonar el repositorio
git clone <url-del-repo>
cd Xchango_APIS_CRUD_Usuarioseguridad

# Instalar dependencias
go mod tidy

# Ejecutar migraciones (crear tablas e insertar datos iniciales)
bee migrate

# Iniciar el servidor
go run main.go
```

El servidor corre en: `http://localhost:8080`

---

## Documentación Swagger

Disponible en modo `dev` en:

```
http://localhost:8080/swagger/
```

---

## Endpoints

Todos los endpoints tienen el prefijo `/v1`.

Cada recurso expone las operaciones CRUD estándar:

| Método   | Ruta                    | Descripción           |
|----------|-------------------------|-----------------------|
| `POST`   | `/v1/{recurso}/`        | Crear registro        |
| `GET`    | `/v1/{recurso}/`        | Listar todos          |
| `GET`    | `/v1/{recurso}/:id`     | Obtener uno por ID    |
| `PUT`    | `/v1/{recurso}/:id`     | Actualizar por ID     |
| `DELETE` | `/v1/{recurso}/:id`     | Eliminar por ID       |

### Recursos disponibles

| Recurso                   | Ruta base                         |
|---------------------------|-----------------------------------|
| Usuario                   | `/v1/Usuario`                     |
| Credenciales              | `/v1/Credenciales`                |
| Crear Contraseña          | `/v1/Crear_contrasena`            |
| Historial Contraseña      | `/v1/Historial_contrasena`        |
| Recuperación Contraseña   | `/v1/Recuperacion_contrasena`     |
| Sesión                    | `/v1/Sesion`                      |
| Intento Login             | `/v1/Intento_login`               |
| Perfil                    | `/v1/Perfil`                      |

### Parámetros de consulta para GetAll

| Parámetro | Tipo   | Descripción                                        |
|-----------|--------|----------------------------------------------------|
| `limit`   | int    | Cantidad de registros (default: 10)                |
| `offset`  | int    | Desplazamiento (default: 0)                        |
| `sortby`  | string | Campo(s) para ordenar. Ej: `nombre`                |
| `order`   | string | `asc` o `desc`                                     |
| `fields`  | string | Campos a retornar. Ej: `correo,activo`             |
| `query`   | string | Filtros. Ej: `activo:true`                         |

---

## Estructura del proyecto

```
Xchango_APIS_CRUD_Usuarioseguridad/
├── conf/
│   └── app.conf           # Configuración de Beego
├── controllers/           # Lógica de los endpoints
├── models/                # Modelos y acceso a base de datos
├── routers/               # Definición de rutas
├── database/
│   ├── migrations/        # Migraciones en Go
│   └── scripts/           # Scripts SQL (up/down)
├── swagger/               # Archivos estáticos de Swagger UI
├── .env                   # Variables de entorno (no subir al repo)
├── go.mod
└── main.go
```

---

## Base de datos

El schema `usuario_seguridad` contiene las siguientes tablas:

- `Usuario` — Registro base del usuario con correo y estado de verificación
- `Credenciales` — Hash de contraseña, intentos fallidos y estado de bloqueo por usuario
- `Crear_contrasena` — Almacenamiento temporal de contraseña y confirmación durante el proceso de creación
- `Historial_contrasena` — Registro histórico de contraseñas anteriores por usuario
- `Recuperacion_contrasena` — Tokens y códigos de recuperación con fecha de expiración y estado de uso
- `Sesion` — Sesiones activas con token, IP de origen, user agent y fecha de expiración
- `Intento_login` — Auditoría de intentos de inicio de sesión (exitosos y fallidos) con IP y motivo de fallo
- `Perfil` — Datos personales del usuario: nombre, apellido, teléfono, WhatsApp, municipio y foto de perfil

---

## CORS

La API tiene CORS habilitado para todos los orígenes, aceptando los métodos `GET`, `POST`, `PUT`, `PATCH`, `DELETE` y `OPTIONS`.

