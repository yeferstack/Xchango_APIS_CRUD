package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

type InsertTableCalificacion_20260522_100300 struct {
	migration.Migration
}

func init() {
	m := &InsertTableCalificacion_20260522_100300{}
	m.Created = "20260522_100300"
	migration.Register("InsertTableCalificacion_20260522_100300", m)
}

func (m *InsertTableCalificacion_20260522_100300) Up() {
	m.SQL(`INSERT INTO "Transacciones_comunicacion"."Calificacion" (id_intercambio, id_usuario_califica, id_usuario_calificado, puntuacion, comentario) VALUES (1, 2, 1, 5, 'Sofía fue muy amable, la cámara estaba tal como la describió. Excelente experiencia.')`)
}

func (m *InsertTableCalificacion_20260522_100300) Down() {
	m.SQL(`DELETE FROM "Transacciones_comunicacion"."Calificacion" WHERE id_intercambio=1 AND id_usuario_califica=2`)
}
