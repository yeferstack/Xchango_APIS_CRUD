CREATE SCHEMA IF NOT EXISTS sistema;

-- =====================================================
-- TABLA: nivel_acceso
-- =====================================================

CREATE TABLE sistema.nivel_acceso (
    id_nivelacceso     SERIAL PRIMARY KEY,
    nombre             VARCHAR(20) NOT NULL,
    activo             BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_asignacion   TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);

-- =====================================================
-- TABLA: permiso
-- =====================================================

CREATE TABLE sistema.permiso (
    id_permiso         SERIAL PRIMARY KEY,
    nombre             VARCHAR(100) NOT NULL,
    descripcion        TEXT,
    activo             BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);

-- =====================================================
-- TABLA: notificacion
-- =====================================================

CREATE TABLE sistema.notificacion (
    id_notificacion    SERIAL PRIMARY KEY,
    id_usuario         INT NOT NULL,
    titulo             VARCHAR(100) NOT NULL,
    mensaje            TEXT NOT NULL,
    tipo               VARCHAR(50) NOT NULL,
    id_referencia      INT,
    tipo_referencia    VARCHAR(50),
    leido              BOOLEAN NOT NULL DEFAULT FALSE,
    activo             BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);

-- =====================================================
-- TABLA: administrador
-- =====================================================

CREATE TABLE sistema.administrador (
    id_admin           SERIAL PRIMARY KEY,
    id_usuario         INT NOT NULL UNIQUE,
    id_nivelacceso     INT NOT NULL,
    activo             BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_asignacion   TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_administrador_nivelacceso
        FOREIGN KEY (id_nivelacceso)
        REFERENCES sistema.nivel_acceso(id_nivelacceso)
);

-- =====================================================
-- TABLA: administradorpermiso
-- =====================================================

CREATE TABLE sistema.administradorpermiso (
    id                 SERIAL PRIMARY KEY,
    id_admin           INT NOT NULL,
    id_permiso         INT NOT NULL,
    activo             BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_adminpermiso_admin
        FOREIGN KEY (id_admin)
        REFERENCES sistema.administrador(id_admin),

    CONSTRAINT fk_adminpermiso_permiso
        FOREIGN KEY (id_permiso)
        REFERENCES sistema.permiso(id_permiso)
);

-- =====================================================
-- TABLA: historialadmin
-- =====================================================

CREATE TABLE sistema.historialadmin (
    id_historial       SERIAL PRIMARY KEY,
    id_admin           INT NOT NULL,
    accion             VARCHAR(100) NOT NULL,
    descripcion        TEXT,
    tipo_objeto        VARCHAR(50),
    id_objeto          INT,
    fecha_accion       TIMESTAMP NOT NULL DEFAULT NOW(),
    activo             BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_historialadmin_admin
        FOREIGN KEY (id_admin)
        REFERENCES sistema.administrador(id_admin)
);

