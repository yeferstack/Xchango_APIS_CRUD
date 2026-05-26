package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertBiendigital_20260525_201305 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertBiendigital_20260525_201305{}
	m.Created = "20260525_201305"

	migration.Register("InsertBiendigital_20260525_201305", m)
}

// Run the migrations
func (m *InsertBiendigital_20260525_201305) Up() {
	m.SQL("INSERT INTO Publicaciones.biendigital (id_publicacion, tipo_archivo) VALUES (1,'PDF')")

}

// Reverse the migrations
func (m *InsertBiendigital_20260525_201305) Down() {
	m.SQL("DELETE FROM Publicaciones.biendigital  WHERE id_publicacion='1'")

}
