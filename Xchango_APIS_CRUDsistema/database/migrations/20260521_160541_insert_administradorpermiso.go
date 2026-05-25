package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertAdministradorpermiso_20260521_160541 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertAdministradorpermiso_20260521_160541{}
	m.Created = "20260521_160541"

	migration.Register("InsertAdministradorpermiso_20260521_160541", m)
}

// Run the migrations

func (m *InsertAdministradorpermiso_20260521_160541) Up() {
    m.SQL("INSERT INTO sistema.administradorpermiso (id_admin, id_permiso) VALUES (1, 1), (1, 2), (1, 3), (1, 4), (1, 5), (1, 6), (1, 7), (1, 8), (2, 2), (2, 3), (2, 4), (2, 5), (3, 4), (3, 7), (4, 7), (5, 4);")
}

// Reverse the migrations
func (m *InsertAdministradorpermiso_20260521_160541) Down() {
    m.SQL("DELETE FROM sistema.administradorpermiso WHERE id_admin IN (1, 2, 3, 4, 5);")
}
