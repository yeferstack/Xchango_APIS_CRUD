package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertNotificacion_20260521_160515 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertNotificacion_20260521_160515{}
	m.Created = "20260521_160515"

	migration.Register("InsertNotificacion_20260521_160515", m)
}

// Run the migrations
func (m *InsertNotificacion_20260521_160515) Up() {
    m.SQL("INSERT INTO sistema.notificacion (id_usuario, titulo, mensaje, tipo, id_referencia, tipo_referencia, leido) VALUES (1, 'nuevo usuario registrado', 'se registró un nuevo usuario en el sistema', 'registro', 101, 'usuario', FALSE), (2, 'pedido actualizado', 'el pedido fue actualizado correctamente', 'pedido', 202, 'pedido', TRUE), (3, 'inventario bajo', 'algunos productos tienen inventario bajo', 'inventario', 303, 'producto', FALSE), (4, 'nueva venta realizada', 'se registró una nueva venta', 'venta', 404, 'venta', TRUE), (5, 'cambio de permisos', 'tus permisos fueron modificados', 'seguridad', 505, 'permiso', FALSE);")
}

// Reverse the migrations
func (m *InsertNotificacion_20260521_160515) Down() {
    m.SQL("DELETE FROM sistema.notificacion WHERE id_usuario IN (1, 2, 3, 4, 5);")
}