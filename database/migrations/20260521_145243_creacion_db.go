package main

import (
	_ "embed"
	"strings"

	"github.com/beego/beego/v2/client/orm/migration"
)

//go:embed 20260521_145243_creacion_db_up.sql
var upSQL string

//go:embed 20260521_145243_creacion_db_down.sql
var downSQL string

type CreacionDb_20260521_145243 struct {
	migration.Migration
}

func init() {
	m := &CreacionDb_20260521_145243{}
	m.Created = "20260521_145243"
	migration.Register("CreacionDb_20260521_145243", m)
}

func (m *CreacionDb_20260521_145243) Up() {
	for _, request := range strings.Split(upSQL, ";") {
		trimmed := strings.TrimSpace(request)
		if trimmed == "" {
			continue
		}
		m.SQL(trimmed)
	}
}

func (m *CreacionDb_20260521_145243) Down() {
	for _, request := range strings.Split(downSQL, ";") {
		trimmed := strings.TrimSpace(request)
		if trimmed == "" {
			continue
		}
		m.SQL(trimmed)
	}
}