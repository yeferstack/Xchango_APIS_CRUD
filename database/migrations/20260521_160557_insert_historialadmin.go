package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertHistorialadmin_20260521_160557 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertHistorialadmin_20260521_160557{}
	m.Created = "20260521_160557"

	migration.Register("InsertHistorialadmin_20260521_160557", m)
}

// Run the migrations
func (m *InsertHistorialadmin_20260521_160557) Up() {
    m.SQL("INSERT INTO sistema.historialadmin (id_admin, accion, descripcion, tipo_objeto, id_objeto) VALUES (1, 'crear', 'creó un nuevo usuario', 'usuario', 11), (1, 'eliminar', 'eliminó un producto del inventario', 'producto', 22), (2, 'actualizar', 'actualizó información de un pedido', 'pedido', 33), (3, 'consultar', 'consultó reportes de ventas', 'reporte', 44), (4, 'enviar', 'envió una notificación masiva', 'notificacion', 55), (5, 'modificar', 'modificó permisos de acceso', 'permiso', 66);")
}

// Reverse the migrations
func (m *InsertHistorialadmin_20260521_160557) Down() {
    m.SQL("DELETE FROM sistema.historialadmin WHERE id_admin IN (1, 2, 3, 4, 5);")
}
