package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertPermiso_20260521_160430 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertPermiso_20260521_160430{}
	m.Created = "20260521_160430"

	migration.Register("InsertPermiso_20260521_160430", m)
}

// Run the migrations
func (m *InsertPermiso_20260521_160430) Up() {
    m.SQL("INSERT INTO sistema.permiso (nombre, descripcion) VALUES ('gestionar usuarios', 'permite crear, editar y eliminar usuarios'), ('gestionar productos', 'permite administrar productos del sistema'), ('gestionar ventas', 'permite registrar y controlar ventas'), ('ver reportes', 'permite visualizar reportes y estadísticas'), ('gestionar pedidos', 'permite administrar pedidos'), ('gestionar inventario', 'permite controlar el inventario'), ('enviar notificaciones', 'permite enviar notificaciones a usuarios'), ('gestionar administradores', 'permite administrar cuentas administrativas');")
}

// Reverse the migrations
func (m *InsertPermiso_20260521_160430) Down() {
    m.SQL("DELETE FROM sistema.permiso WHERE nombre IN ('gestionar usuarios', 'gestionar productos', 'gestionar ventas', 'ver reportes', 'gestionar pedidos', 'gestionar inventario', 'enviar notificaciones', 'gestionar administradores');")
}
