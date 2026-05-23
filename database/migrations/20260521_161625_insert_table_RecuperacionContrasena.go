package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableRecuperacionContrasena_20260521_161625 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableRecuperacionContrasena_20260521_161625{}
	m.Created = "20260521_161625"

	migration.Register("InsertTableRecuperacionContrasena_20260521_161625", m)
}

// Run the migrations
func (m *InsertTableRecuperacionContrasena_20260521_161625) Up() {
	m.SQL("INSERT INTO usuario_seguridad.recuperacion_contrasena (id_usuario, token, codigo, usado, fecha_expiracion) VALUES ('1', 'tkn_valentina_abc123xyz', '847201', 'true', '2025-04-10 12:00:00')")
}

// Reverse the migrations
func (m *InsertTableRecuperacionContrasena_20260521_161625) Down() {
	m.SQL("DELETE FROM usuario_seguridad.recuperacion_contrasena WHERE token='tkn_valentina_abc123xyz'") 

}
