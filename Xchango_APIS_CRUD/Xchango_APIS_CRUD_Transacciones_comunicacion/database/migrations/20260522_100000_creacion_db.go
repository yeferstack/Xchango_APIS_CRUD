package main

import (
	_ "embed"
	"strings"

	"github.com/beego/beego/v2/client/orm/migration"
)

//go:embed 20260522_100000_creacion_db_up.sql
var upSQL string

//go:embed 20260522_100000_creacion_db_down.sql
var downSQL string

type CreacionDb_20260522_100000 struct {
	migration.Migration
}

func init() {
	m := &CreacionDb_20260522_100000{}
	m.Created = "20260522_100000"
	migration.Register("CreacionDb_20260522_100000", m)
}

func (m *CreacionDb_20260522_100000) Up() {
	for _, request := range strings.Split(upSQL, ";") {
		trimmed := strings.TrimSpace(request)
		if trimmed == "" {
			continue
		}
		m.SQL(trimmed)
	}
}

func (m *CreacionDb_20260522_100000) Down() {
	for _, request := range strings.Split(downSQL, ";") {
		trimmed := strings.TrimSpace(request)
		if trimmed == "" {
			continue
		}
		m.SQL(trimmed)
	}
}
