package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertServicio_20260525_201526 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertServicio_20260525_201526{}
	m.Created = "20260525_201526"

	migration.Register("InsertServicio_20260525_201526", m)
}

// Run the migrations
func (m *InsertServicio_20260525_201526) Up() {
	 m.SQL("INSERT INTO Publicaciones.servicio (id_servicio, id_publicacion, duracion) VALUES (1, 1, '2 horas')")

}

// Reverse the migrations
func (m *InsertServicio_20260525_201526) Down() {
	 m.SQL("DELETE FROM Publicaciones.servicio WHERE id_publicacion='1'") 

}
