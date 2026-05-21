package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableCrearContrasena_20260521_163625 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableCrearContrasena_20260521_163625{}
	m.Created = "20260521_163625"

	migration.Register("InsertTableCrearContrasena_20260521_163625", m)
}

// Run the migrations
func (m *InsertTableCrearContrasena_20260521_163625) Up() {
	m.SQL("INSERT INTO usuario_seguridad.Crear_contrasena (contrasena, confirmar_contrasena) VALUES ('Sofia@2025!', 'Sofia@2025!')")

}

// Reverse the migrations
func (m *InsertTableCrearContrasena_20260521_163625) Down() {
	m.SQL("DELETE FROM usuario_seguridad.Crear_contrasena WHERE contrasena='Sofia@2025!'") 

}
