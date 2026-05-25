package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/beego/beego/v2/client/orm/migration"

)

// DO NOT MODIFY
type CreacionDb_20260521_153457 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionDb_20260521_153457{}
	m.Created = "20260521_153457"

	migration.Register("CreacionDb_20260521_153457", m)
}

// Run the migrations
func (m *CreacionDb_20260521_153457) Up() {
	file, err:= ioutil.ReadFile("database/scripst/20260521_153457_creacion_db_up.sql")

	if err != nil {
		fmt.Print(err)
	}

	requests := strings.Split(string(file), ";")
	
	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}

}

// Reverse the migrations
func (m *CreacionDb_20260521_153457) Down() {
	file, err:= ioutil.ReadFile("database/scripst/20260521_153457_creacion_db_down.sql")

	if err != nil {
		fmt.Print(err)
	}

	requests := strings.Split(string(file), ";")
	
	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
