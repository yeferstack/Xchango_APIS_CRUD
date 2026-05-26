package main

import (
		"fmt"
	"io/ioutil"
	_"embed"
	"strings"
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreacionDb_20260525_200622 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionDb_20260525_200622{}
	m.Created = "20260525_200622"

	migration.Register("CreacionDb_20260525_200622", m)
}

// Run the migrations
func (m *CreacionDb_20260525_200622) Up() {
	file, err:= ioutil.ReadFile("../scripts/20260525_200622_creacion_db_up.sql")

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
func (m *CreacionDb_20260525_200622) Down() {
	file, err:= ioutil.ReadFile("../scripts/20260525_200622_creacion_db_down.sql")

	if err !=nil{
		fmt.Println(err)
	}

	requests:= strings.Split(string(file),";")

	for _, request:= range requests{
		fmt.Println(request)
		m.SQL(request)
	}

}
