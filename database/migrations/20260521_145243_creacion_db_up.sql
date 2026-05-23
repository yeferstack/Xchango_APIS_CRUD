CREATE SCHEMA IF NOT EXISTS usuario_seguridad;

CREATE TABLE usuario_seguridad.usuario (
    id_usuario         SERIAL PRIMARY KEY,
    correo             VARCHAR(100) UNIQUE NOT NULL,
    correo_verificado  BOOLEAN NOT NULL DEFAULT FALSE,
    activo             BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT now(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE usuario_seguridad.credenciales (
    id_usuario         INT PRIMARY KEY,
    contrasena_hash    TEXT NOT NULL,
    intentos_fallidos  INT NOT NULL DEFAULT 0,
    bloqueado          BOOLEAN NOT NULL DEFAULT FALSE,
    ultimo_login       TIMESTAMP,
    activo             BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT now(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT now(),
    -- FK interna 
    CONSTRAINT fk_credenciales_usuario
        FOREIGN KEY (id_usuario) REFERENCES usuario_seguridad.usuario(id_usuario)
        ON DELETE CASCADE
);

CREATE TABLE usuario_seguridad.crear_contrasena (
    id_usuario           INT PRIMARY KEY,
    contrasena           TEXT NOT NULL,
    confirmar_contrasena TEXT NOT NULL,
    activo               BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion       TIMESTAMP NOT NULL DEFAULT now(),
    fecha_modificacion   TIMESTAMP NOT NULL DEFAULT now(),
    -- FK interna 
    CONSTRAINT fk_crear_contrasena_usuario
        FOREIGN KEY (id_usuario) REFERENCES usuario_seguridad.usuario(id_usuario)
        ON DELETE CASCADE
);

CREATE TABLE usuario_seguridad.historial_contrasena (
    id_historial       SERIAL PRIMARY KEY,
    id_usuario         INT NOT NULL,
    contrasena_hash    TEXT NOT NULL,
    fecha_cambio       TIMESTAMP NOT NULL DEFAULT now(),
    activo             BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT now(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT now(),
    -- FK interna 
    CONSTRAINT fk_historialcontrasena_usuario
        FOREIGN KEY (id_usuario) REFERENCES usuario_seguridad.usuario(id_usuario)
        ON DELETE CASCADE
);

CREATE TABLE usuario_seguridad.recuperacion_contrasena (
    id_recuperacion    SERIAL PRIMARY KEY,
    id_usuario         INT NOT NULL,
    token              TEXT NOT NULL,
    codigo             VARCHAR(10),
    usado              BOOLEAN NOT NULL DEFAULT FALSE,
    fecha_expiracion   TIMESTAMP NOT NULL,
    activo             BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT now(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT now(),
    -- FK interna 
    CONSTRAINT fk_recuperacioncontrasena_usuario
        FOREIGN KEY (id_usuario) REFERENCES usuario_seguridad.usuario(id_usuario)
        ON DELETE CASCADE
);

CREATE TABLE usuario_seguridad.sesion (
    id_sesion          SERIAL PRIMARY KEY,
    id_usuario         INT NOT NULL,
    token_sesion       TEXT NOT NULL,
    ip_origen          VARCHAR(45),
    user_agent         TEXT,
    fecha_inicio       TIMESTAMP NOT NULL DEFAULT now(),
    fecha_expiracion   TIMESTAMP NOT NULL,
    revocada           BOOLEAN NOT NULL DEFAULT FALSE,
    activo             BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT now(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT now(),
    -- FK interna 
    CONSTRAINT fk_sesion_usuario
        FOREIGN KEY (id_usuario) REFERENCES usuario_seguridad.usuario(id_usuario)
        ON DELETE CASCADE
);

CREATE TABLE usuario_seguridad.intento_login (
    id_intento         SERIAL PRIMARY KEY,
    id_usuario         INT,
    email_ingresado    VARCHAR(100),
    exitoso            BOOLEAN NOT NULL DEFAULT FALSE,
    motivo_fallo       VARCHAR(100),
    ip_origen          VARCHAR(45),
    fecha              TIMESTAMP NOT NULL DEFAULT now(),
    activo             BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT now(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT now(),
    -- FK interna 
    CONSTRAINT fk_intentologin_usuario
        FOREIGN KEY (id_usuario) REFERENCES usuario_seguridad.usuario(id_usuario)
        ON DELETE CASCADE
);

CREATE TABLE usuario_seguridad.perfil (
    id_perfil          SERIAL PRIMARY KEY,
    id_usuario         INT NOT NULL UNIQUE,
    nombre             VARCHAR(100) NOT NULL,
    apellido           VARCHAR(100) NOT NULL,
    telefono           VARCHAR(20),
    whatsapp           VARCHAR(20) NOT NULL,
    municipio          VARCHAR(50) NOT NULL,
    barrio             VARCHAR(50),
    genero             VARCHAR(20),
    fecha_nacimiento   DATE,
    biografia          TEXT,
    foto_perfil        TEXT,
    activo             BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion     TIMESTAMP NOT NULL DEFAULT now(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT now(),
    -- FK interna 
    CONSTRAINT fk_perfil_usuario
        FOREIGN KEY (id_usuario) REFERENCES usuario_seguridad.usuario(id_usuario)
        ON DELETE CASCADE
);