package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertImagenpublicacion_20260525_201349 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertImagenpublicacion_20260525_201349{}
	m.Created = "20260525_201349"

	migration.Register("InsertImagenpublicacion_20260525_201349", m)
}

// Run the migrations
func (m *InsertImagenpublicacion_20260525_201349) Up() {
	m.SQL("INSERT INTO Publicaciones.imagenpublicacion (id_publicacion, url, tipo, orden, activo)VALUES (1, 'https://miweb.com/img1.jpg', 'jpg', 1, true)") 

}

// Reverse the migrations
func (m *InsertImagenpublicacion_20260525_201349) Down() {
	m.SQL("DELETE FROM Publicaciones.imagenpublicacion WHERE id_publicacion='1'") 

}
