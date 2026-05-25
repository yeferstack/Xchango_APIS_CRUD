package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

type InsertTableContactoWhatsApp_20260522_100200 struct {
	migration.Migration
}

func init() {
	m := &InsertTableContactoWhatsApp_20260522_100200{}
	m.Created = "20260522_100200"
	migration.Register("InsertTableContactoWhatsApp_20260522_100200", m)
}

func (m *InsertTableContactoWhatsApp_20260522_100200) Up() {
	m.SQL(`INSERT INTO "Transacciones_comunicacion"."ContactoWhatsApp" (id_publicacion, id_usuario_interesado, id_usuario_propietario) VALUES (1, 2, 1)`)
}

func (m *InsertTableContactoWhatsApp_20260522_100200) Down() {
	m.SQL(`DELETE FROM "Transacciones_comunicacion"."ContactoWhatsApp" WHERE id_publicacion=1 AND id_usuario_interesado=2`)
}
