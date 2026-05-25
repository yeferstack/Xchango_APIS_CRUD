CREATE SCHEMA IF NOT EXISTS "Transacciones_comunicacion";

CREATE TABLE "Transacciones_comunicacion"."Intercambio" (
  id_intercambio         SERIAL PRIMARY KEY,
  id_publicacion         INT NOT NULL,
  id_usuario_interesado  INT NOT NULL,
  estado                 VARCHAR(20) DEFAULT 'iniciado',
  mensaje_inicial        TEXT,
  activo                 BOOLEAN   NOT NULL DEFAULT TRUE,
  fecha_creacion         TIMESTAMP NOT NULL DEFAULT now(),
  fecha_modificacion     TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "Transacciones_comunicacion"."ContactoWhatsApp" (
  id_contacto            SERIAL PRIMARY KEY,
  id_publicacion         INT NOT NULL,
  id_usuario_interesado  INT NOT NULL,
  id_usuario_propietario INT NOT NULL,
  activo                 BOOLEAN   NOT NULL DEFAULT TRUE,
  fecha_creacion         TIMESTAMP NOT NULL DEFAULT now(),
  fecha_modificacion     TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE "Transacciones_comunicacion"."Calificacion" (
  id_calificacion        SERIAL PRIMARY KEY,
  id_intercambio         INT NOT NULL,
  id_usuario_califica    INT NOT NULL,
  id_usuario_calificado  INT NOT NULL,
  puntuacion             INT NOT NULL,
  comentario             TEXT,
  activo                 BOOLEAN   NOT NULL DEFAULT TRUE,
  fecha_creacion         TIMESTAMP NOT NULL DEFAULT now(),
  fecha_modificacion     TIMESTAMP NOT NULL DEFAULT now(),
  CONSTRAINT fk_calificacion_intercambio
    FOREIGN KEY (id_intercambio) REFERENCES "Transacciones_comunicacion"."Intercambio"(id_intercambio)
);

CREATE TABLE "Transacciones_comunicacion"."Reporte" (
  id_reporte         SERIAL PRIMARY KEY,
  id_usuario         INT NOT NULL,
  id_publicacion     INT NOT NULL,
  descripcion        TEXT NOT NULL,
  estado             VARCHAR(20) DEFAULT 'pendiente',
  activo             BOOLEAN   NOT NULL DEFAULT TRUE,
  fecha_creacion     TIMESTAMP NOT NULL DEFAULT now(),
  fecha_modificacion TIMESTAMP NOT NULL DEFAULT now()
)
