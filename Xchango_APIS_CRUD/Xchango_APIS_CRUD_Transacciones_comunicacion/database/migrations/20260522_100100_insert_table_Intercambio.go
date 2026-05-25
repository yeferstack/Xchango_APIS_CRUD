package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

type InsertTableIntercambio_20260522_100100 struct {
	migration.Migration
}

func init() {
	m := &InsertTableIntercambio_20260522_100100{}
	m.Created = "20260522_100100"
	migration.Register("InsertTableIntercambio_20260522_100100", m)
}

func (m *InsertTableIntercambio_20260522_100100) Up() {
	m.SQL(`INSERT INTO "Transacciones_comunicacion"."Intercambio" (id_publicacion, id_usuario_interesado, estado, mensaje_inicial) VALUES (1, 2, 'aceptado', 'Hola Sofía, me interesa la cámara, tengo juegos de PS3 para cambiar.')`)
}

func (m *InsertTableIntercambio_20260522_100100) Down() {
	m.SQL(`DELETE FROM "Transacciones_comunicacion"."Intercambio" WHERE id_publicacion=1 AND id_usuario_interesado=2`)
}
