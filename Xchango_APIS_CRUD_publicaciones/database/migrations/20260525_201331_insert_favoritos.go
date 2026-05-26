package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertFavoritos_20260525_201331 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertFavoritos_20260525_201331{}
	m.Created = "20260525_201331"

	migration.Register("InsertFavoritos_20260525_201331", m)
}

// Run the migrations
func (m *InsertFavoritos_20260525_201331) Up() {
	 m.SQL("INSERT INTO Publicaciones.favorito (id_usuario, id_publicacion) VALUES (1, 1)")

}

// Reverse the migrations
func (m *InsertFavoritos_20260525_201331) Down() {
	 m.SQL("DELETE FROM Publicaciones.favorito WHERE id_usuario='1'") 

}
