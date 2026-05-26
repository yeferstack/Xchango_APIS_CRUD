package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertPublicacioncategoria_20260525_201445 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertPublicacioncategoria_20260525_201445{}
	m.Created = "20260525_201445"

	migration.Register("InsertPublicacioncategoria_20260525_201445", m)
}

// Run the migrations
func (m *InsertPublicacioncategoria_20260525_201445) Up() {
	m.SQL("INSERT INTO Publicaciones.publicacioncategoria (id_publicacion, id_categoria) VALUES (1,1)") 
}

// Reverse the migrations
func (m *InsertPublicacioncategoria_20260525_201445) Down() {
	 m.SQL("DELETE FROM Publicacionespublicacioncategoria WHERE id_publicacion='1'") 

}
