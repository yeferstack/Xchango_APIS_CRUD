package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableHistorialContrasena_20260521_163556 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableHistorialContrasena_20260521_163556{}
	m.Created = "20260521_163556"

	migration.Register("InsertTableHistorialContrasena_20260521_163556", m)
}

// Run the migrations
func (m *InsertTableHistorialContrasena_20260521_163556) Up() {
	m.SQL("INSERT INTO usuario_seguridad.historial_contrasena (id_usuario, contrasena_hash, fecha_cambio) VALUES ('1', '$2b$12$OLD1sofia', '2025-01-10 10:00:00')")

}

// Reverse the migrations
func (m *InsertTableHistorialContrasena_20260521_163556) Down() {
	m.SQL("DELETE FROM usuario_seguridad.historial_contrasena WHERE contrasena_hash='$2b$12$OLD1sofia'") 

}
