package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableIntentoLogin_20260521_163535 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableIntentoLogin_20260521_163535{}
	m.Created = "20260521_163535"

	migration.Register("InsertTableIntentoLogin_20260521_163535", m)
}

// Run the migrations
func (m *InsertTableIntentoLogin_20260521_163535) Up() {
	m.SQL("INSERT INTO usuario_seguridad.IntentoLogin (id_usuario, email_ingresado, exitoso, motivo_fallo, ip_origen, fecha) VALUES ('1', 'sofia.ramirez@gmail.com', 'true', 'contrasena_incorrecta', '181.55.23.104', '2025-05-18 08:30:00')")

}

// Reverse the migrations
func (m *InsertTableIntentoLogin_20260521_163535) Down() {
	m.SQL("DELETE FROM usuario_seguridad.IntentoLogin WHERE email_ingresado='sofia.ramirez@gmail.com'") 

}
