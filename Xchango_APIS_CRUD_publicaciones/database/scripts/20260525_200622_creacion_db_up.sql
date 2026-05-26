CREATE SCHEMA IF NOT EXISTS Publicaciones;


CREATE TABLE Publicaciones.publicacion (
  id_publicacion          SERIAL PRIMARY KEY,
  id_usuario              INT NOT NULL,       
  estadopublicacion       BOOLEAN NOT NULL DEFAULT TRUE,
  titulo                  VARCHAR(150) NOT NULL,
  descripcion             TEXT NOT NULL,
  tipo                    VARCHAR(20)  NOT NULL,
  visibilidad             VARCHAR(20)  DEFAULT 'publica',
  departamento            VARCHAR(50)  NOT NULL,
  municipio               VARCHAR(50)  NOT NULL,
  barrio                  VARCHAR(50),
  disponible_para_trueque BOOLEAN      DEFAULT TRUE,
  cantidad_disponible     INT          DEFAULT 1,
  prioridad               INT          DEFAULT 0,
  mensaje_contacto        TEXT,
  vistas                  INT          DEFAULT 0,
  favoritos               INT          DEFAULT 0,
  activo                  BOOLEAN      NOT NULL DEFAULT TRUE,
  fecha_creacion          TIMESTAMP    NOT NULL DEFAULT now(),
  fecha_modificacion      TIMESTAMP    NOT NULL DEFAULT now()
);

CREATE TABLE Publicaciones.imagenpublicacion (
  id_imagen          SERIAL PRIMARY KEY,
  id_publicacion     INT NOT NULL,
  url                TEXT NOT NULL,
  tipo               VARCHAR(20),
  orden              INT DEFAULT 1,
  activo             BOOLEAN   NOT NULL DEFAULT TRUE,
  fecha_creacion     TIMESTAMP NOT NULL DEFAULT now(),
  fecha_modificacion TIMESTAMP NOT NULL DEFAULT now(),
  CONSTRAINT fk_imagen_publicacion
    FOREIGN KEY (id_publicacion) REFERENCES Publicaciones.publicacion(id_publicacion)ON DELETE CASCADE
);

CREATE TABLE Publicaciones.sub_categoria (
  id_categoria       SERIAL PRIMARY KEY,
  electronico        VARCHAR,
  vehiculos          VARCHAR,
  ropa               VARCHAR,
  libros             VARCHAR,
  muebles            VARCHAR,
  juguetes           VARCHAR,
  activo             BOOLEAN   NOT NULL DEFAULT TRUE,
  fecha_creacion     TIMESTAMP NOT NULL DEFAULT now(),
  fecha_modificacion TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE Publicaciones.publicacioncategoria (
  id                 SERIAL PRIMARY KEY,
  id_publicacion     INT NOT NULL,
  id_categoria       INT NOT NULL,
  activo             BOOLEAN   NOT NULL DEFAULT TRUE,
  fecha_creacion     TIMESTAMP NOT NULL DEFAULT now(),
  fecha_modificacion TIMESTAMP NOT NULL DEFAULT now(),

  CONSTRAINT fk_publicacioncategoria_publicacion
    FOREIGN KEY (id_publicacion) 
    REFERENCES Publicaciones.publicacion(id_publicacion)
    ON DELETE CASCADE,

  CONSTRAINT fk_publicacioncategoria_categoria
    FOREIGN KEY (id_categoria)   
    REFERENCES Publicaciones.sub_categoria(id_categoria)
    ON DELETE CASCADE
);

CREATE TABLE Publicaciones.bienFisico (
  id_bien            SERIAL PRIMARY KEY,
  id_publicacion     INT NOT NULL UNIQUE,
  estado_producto    VARCHAR(50),
  marca              VARCHAR(100),
  modelo             VARCHAR(100),
  color              VARCHAR(50),
  peso               DECIMAL,
  dimensiones        VARCHAR(100),
  activo             BOOLEAN   NOT NULL DEFAULT TRUE,
  fecha_creacion     TIMESTAMP NOT NULL DEFAULT now(),
  fecha_modificacion TIMESTAMP NOT NULL DEFAULT now(),
  CONSTRAINT fk_bienfisico_publicacion
    FOREIGN KEY (id_publicacion) REFERENCES Publicaciones.publicacion(id_publicacion)ON DELETE CASCADE);

CREATE TABLE Publicaciones.servicio (
  id_servicio        SERIAL PRIMARY KEY,
  id_publicacion     INT NOT NULL UNIQUE,
  duracion           VARCHAR(50),
  modalidad          VARCHAR(50),
  disponibilidad     TEXT,
  requisitos         TEXT,
  activo             BOOLEAN   NOT NULL DEFAULT TRUE,
  fecha_creacion     TIMESTAMP NOT NULL DEFAULT now(),
  fecha_modificacion TIMESTAMP NOT NULL DEFAULT now(),
  CONSTRAINT fk_servicio_publicacion
    FOREIGN KEY (id_publicacion) REFERENCES Publicaciones.publicacion(id_publicacion)ON DELETE CASCADE
);

CREATE TABLE Publicaciones.biendigital (
  id_bien_digital    SERIAL PRIMARY KEY,
  id_publicacion     INT NOT NULL UNIQUE,
  tipo_archivo       VARCHAR(50),
  tamano_mb          DECIMAL,
  licencia           VARCHAR(100),
  acceso_inmediato   BOOLEAN,
  activo             BOOLEAN   NOT NULL DEFAULT TRUE,
  fecha_creacion     TIMESTAMP NOT NULL DEFAULT now(),
  fecha_modificacion TIMESTAMP NOT NULL DEFAULT now(),
  CONSTRAINT fk_biendigital_publicacion
    FOREIGN KEY (id_publicacion) REFERENCES Publicaciones.publicacion(id_publicacion)ON DELETE CASCADE
);

CREATE TABLE Publicaciones.favorito (
  id_favorito        SERIAL PRIMARY KEY,
  id_usuario         INT NOT NULL,       -- sin FK externa
  id_publicacion     INT NOT NULL,
  activo             BOOLEAN   NOT NULL DEFAULT TRUE,
  fecha_creacion     TIMESTAMP NOT NULL DEFAULT now(),
  fecha_modificacion TIMESTAMP NOT NULL DEFAULT now(),
  CONSTRAINT fk_favorito_publicacion
    FOREIGN KEY (id_publicacion) REFERENCES Publicaciones.publicacion(id_publicacion)ON DELETE CASCADE
);
