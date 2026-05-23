package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTablePerfil_20260521_162627 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablePerfil_20260521_162627{}
	m.Created = "20260521_162627"

	migration.Register("InsertTablePerfil_20260521_162627", m)
}

// Run the migrations
func (m *InsertTablePerfil_20260521_162627) Up() {
	m.SQL("INSERT INTO usuario_seguridad.perfil (id_usuario, nombre, apellido, telefono, whatsapp, municipio, barrio, genero, fecha_nacimiento, biografia, foto_perfil) VALUES ('1', 'Sofía', 'Ramírez', '3012345678', '3012345678', 'Medellín', 'Laureles', 'Femenino', '1998-06-12', 'Amante del diseño y el intercambio de libros.', 'https://cdn.xchango.co/fotos/sofia_ramirez.jpg')")

}

// Reverse the migrations
func (m *InsertTablePerfil_20260521_162627) Down() {
	m.SQL("DELETE FROM usuario_seguridad.perfil WHERE id_usuario='1'") 

}
