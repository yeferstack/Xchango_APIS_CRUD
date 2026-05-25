package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableSesion_20260521_160534 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableSesion_20260521_160534{}
	m.Created = "20260521_160534"

	migration.Register("InsertTableSesion_20260521_160534", m)
}

// Run the migrations
func (m *InsertTableSesion_20260521_160534) Up() {
	m.SQL("INSERT INTO usuario_seguridad.Sesion (id_usuario, token_sesion, ip_origen, user_agent, fecha_inicio, fecha_expiracion, revocada) VALUES ('1', 'sess_sofia_token_a1b2c3', '181.55.23.104', 'Mozilla/5.0 (Android 13; SM-G998B)', '2025-05-18 08:30:00', '2025-05-19 08:30:00', 'false')")

}

// Reverse the migrations
func (m *InsertTableSesion_20260521_160534) Down() {
	m.SQL("DELETE FROM usuario_seguridad.Sesion WHERE token_sesion='sess_sofia_token_a1b2c3'") 

}
