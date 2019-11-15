package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type TablasEsquemaRecibo_20191115_054616 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &TablasEsquemaRecibo_20191115_054616{}
	m.Created = "20191115_054616"

	migration.Register("TablasEsquemaRecibo_20191115_054616", m)
}

// Run the migrations
func (m *TablasEsquemaRecibo_20191115_054616) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20191115_054616_tablas_esquema_recibo.up.sql")

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
func (m *TablasEsquemaRecibo_20191115_054616) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/scripts/20191115_054616_tablas_esquema_recibo.down.sql")

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
