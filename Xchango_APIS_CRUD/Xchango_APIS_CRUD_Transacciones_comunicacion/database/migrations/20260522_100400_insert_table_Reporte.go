package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

type InsertTableReporte_20260522_100400 struct {
	migration.Migration
}

func init() {
	m := &InsertTableReporte_20260522_100400{}
	m.Created = "20260522_100400"
	migration.Register("InsertTableReporte_20260522_100400", m)
}

func (m *InsertTableReporte_20260522_100400) Up() {
	m.SQL(`INSERT INTO "Transacciones_comunicacion"."Reporte" (id_usuario, id_publicacion, descripcion, estado) VALUES (3, 2, 'La descripción de los juegos no coincide con lo recibido. Algunos tienen rayones no mencionados.', 'en_revision')`)
}

func (m *InsertTableReporte_20260522_100400) Down() {
	m.SQL(`DELETE FROM "Transacciones_comunicacion"."Reporte" WHERE id_usuario=3 AND id_publicacion=2`)
}
