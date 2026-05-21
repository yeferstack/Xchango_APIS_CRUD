package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableCredenciales_20260521_163608 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableCredenciales_20260521_163608{}
	m.Created = "20260521_163608"

	migration.Register("InsertTableCredenciales_20260521_163608", m)
}

// Run the migrations
func (m *InsertTableCredenciales_20260521_163608) Up() {
	m.SQL("INSERT INTO usuario_seguridad.Credenciales (contrasena_hash, intentos_fallidos, bloqueado, ultimo_login) VALUES ('$2b$12$KIXabc123hasheado1sofia', '0', 'false', '2025-05-18 08:30:00')")

}

// Reverse the migrations
func (m *InsertTableCredenciales_20260521_163608) Down() {
	m.SQL("DELETE FROM usuario_seguridad.Credenciales WHERE contrasena_hash='$2b$12$KIXabc123hasheado1sofia'") 

}
