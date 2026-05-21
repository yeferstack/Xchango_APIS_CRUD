package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreacionDb_20260521_145243 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionDb_20260521_145243{}
	m.Created = "20260521_145243"

	migration.Register("CreacionDb_20260521_145243", m)
}

// Run the migrations
func (m *CreacionDb_20260521_145243) Up() {
	file, err:= ioutil.ReadFile("../scripts/20260514_125045_creacion_db_up.sql")

	if err !=nil{
		fmt.Println(err)
	}

	requests:= strings.Split(string(file),";")

	for _, request:= range requests{
		fmt.Println(request)
		m.SQL(request)
	}
}

// Reverse the migrations
func (m *CreacionDb_20260521_145243) Down() {
	file, err:= ioutil.ReadFile("../scripts/20260514_125045_creacion_db_down.sql")

	if err !=nil{
		fmt.Println(err)
	}

	requests:= strings.Split(string(file),";")

	for _, request:= range requests{
		fmt.Println(request)
		m.SQL(request)
	}
}
