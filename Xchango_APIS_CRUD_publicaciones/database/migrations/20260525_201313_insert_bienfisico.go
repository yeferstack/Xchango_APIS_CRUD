package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertBienfisico_20260525_201313 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertBienfisico_20260525_201313{}
	m.Created = "20260525_201313"

	migration.Register("InsertBienfisico_20260525_201313", m)
}

// Run the migrations
func (m *InsertBienfisico_20260525_201313) Up() {
	m.SQL("INSERT INTO Publicaciones.bienfisico (id_publicacion, estado_producto) VALUES (1,'Nuevo')")
}

// Reverse the migrations
func (m *InsertBienfisico_20260525_201313) Down() {
	m.SQL("DELETE FROMPublicaciones.bienfisico WHERE id_usuario='1'")

}
