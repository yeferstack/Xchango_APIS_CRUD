package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableUsuario_20260521_152942 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableUsuario_20260521_152942{}
	m.Created = "20260521_152942"

	migration.Register("InsertTableUsuario_20260521_152942", m)
}

// Run the migrations
func (m *InsertTableUsuario_20260521_152942) Up() {
	m.SQL("INSERT INTO usuario_seguridad.usuario (correo, correo_verificado) VALUES ('sofia.ramirez@gmail.com', 'true')")

}

// Reverse the migrations
func (m *InsertTableUsuario_20260521_152942) Down() {
	m.SQL("DELETE FROM usuario_seguridad.usuario WHERE correo='sofia.ramirez@gmail.com'") 

}
