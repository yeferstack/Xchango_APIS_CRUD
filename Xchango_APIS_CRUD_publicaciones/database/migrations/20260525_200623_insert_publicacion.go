package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertPublicacion_20260525_201441 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertPublicacion_20260525_201441{}
	m.Created = "20260525_200623"

	migration.Register("InsertPublicacion_20260525_200623", m)
}

// Run the migrations
func (m *InsertPublicacion_20260525_201441) Up() {
	 m.SQL("INSERT INTO Publicaciones.publicacion (id_usuario,titulo, descripcion, tipo, departamento, municipio) VALUES ('1', 'Titulo prueba', 'Descripcion prueba', 'producto', 'Cundinamarca', 'Bogota')") 
}

// Reverse the migrations
func (m *InsertPublicacion_20260525_201441) Down() {
	 m.SQL("DELETE FROM Publicaciones.publicacion WHERE id_usuario='1'") 

}
