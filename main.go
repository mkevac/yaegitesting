package main

import (
	"log"
	"net/http"
	"reflect"

	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"
)

var Symbols = map[string]map[string]reflect.Value{}

func main() {
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	i.Use(Symbols)

	userDB := NewDBWithRandomUsers(100000)
	log.Printf("%d userd in db", len(userDB.users))
	if err := http.ListenAndServe(":8080", myHTTPHandler{i, userDB}); err != nil {
		log.Printf("error while listening: %v", err)
	}
}
