package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertNivelAcceso_20260521_155659 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertNivelAcceso_20260521_155659{}
	m.Created = "20260521_155659"

	migration.Register("InsertNivelAcceso_20260521_155659", m)
}

// Run the migrations
func (m *InsertNivelAcceso_20260521_155659) Up() {
	m.SQL(`INSERT INTO sistema.nivel_acceso (nombre, fecha_asignacion) VALUES 
		('super admin', '2026-05-20 17:14:45.839255'),
		('admin', '2026-05-20 17:14:45.839255'),
		('moderador', '2026-05-20 17:14:45.839255'),
		('soporte', '2026-05-20 17:14:45.839255'),
		('auditor', '2026-05-20 17:14:45.839255');`)
}

// Reverse the migrations
func (m *InsertNivelAcceso_20260521_155659) Down() {
	m.SQL("DELETE FROM sistema.nivel_acceso WHERE nombre IN ('super admin', 'admin', 'moderador', 'soporte', 'auditor');")
}