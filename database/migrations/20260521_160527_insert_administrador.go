package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertAdministrador_20260521_160527 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertAdministrador_20260521_160527{}
	m.Created = "20260521_160527"

	migration.Register("InsertAdministrador_20260521_160527", m)
}

// Run the migrations
func (m *InsertAdministrador_20260521_160527) Up() {
    m.SQL("INSERT INTO sistema.administrador (id_usuario, id_nivelacceso) VALUES (1, 1), (2, 2), (3, 3), (4, 4), (5, 5);")
}

// Reverse the migrations
func (m *InsertAdministrador_20260521_160527) Down() {
    m.SQL("DELETE FROM sistema.administrador WHERE id_usuario IN (1, 2, 3, 4, 5);")
}
