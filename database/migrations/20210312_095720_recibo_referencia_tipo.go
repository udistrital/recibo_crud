package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ReciboReferenciaTipo_20210312_095720 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ReciboReferenciaTipo_20210312_095720{}
	m.Created = "20210312_095720"

	migration.Register("ReciboReferenciaTipo_20210312_095720", m)
}

// Run the migrations
func (m *ReciboReferenciaTipo_20210312_095720) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210312_095720_recibo_referencia_tipo.up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}

// Reverse the migrations
func (m *ReciboReferenciaTipo_20210312_095720) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210312_095720_recibo_referencia_tipo.down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}
