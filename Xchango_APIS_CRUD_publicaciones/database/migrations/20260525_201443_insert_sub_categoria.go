package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertSubCategoria_20260525_201527 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertSubCategoria_20260525_201527{}
	m.Created = "20260525_201443"

	migration.Register("InsertSubCategoria_20260525_201443", m)
}

// Run the migrations
func (m *InsertSubCategoria_20260525_201527) Up() {
	 m.SQL("INSERT INTO Publicaciones.sub_categoria (id_categoria, electronico) VALUES (1, 'Celulares')") 

}

// Reverse the migrations
func (m *InsertSubCategoria_20260525_201527) Down() {
	 m.SQL("DELETE FROM Publicaciones.sub_categoria WHERE id_categoria='1'")

}
