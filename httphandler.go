package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/containous/yaegi/interp"
	"github.com/mkevac/yaegitesting/user"
)

type myHTTPHandler struct {
	yaegi *interp.Interpreter
	db    *UserDB
}

func (m myHTTPHandler) ServeHTTP(handler http.ResponseWriter, req *http.Request) {
	start := time.Now()

	bodyBuf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("error while reading body: %v", err)
		handler.WriteHeader(http.StatusInternalServerError)
		return
	}
	bodyString := string(bodyBuf)
	log.Printf("read body '%s'", bodyString)

	_, err = m.yaegi.Eval(bodyString)
	if err != nil {
		log.Printf("error while evaluating trough yaegi: %v", err)
		handler.WriteHeader(http.StatusInternalServerError)
		return
	}

	v, err := m.yaegi.Eval("filter")
	if err != nil {
		log.Printf("error while getting process func: %v", err)
		handler.WriteHeader(http.StatusInternalServerError)
		return
	}

	filter := v.Interface().(func(*user.User) bool)

	total := len(m.db.users)
	left := 0
	for i := 0; i < total; i++ {
		if ok := filter(m.db.users[i]); ok {
			left++
		}
	}
	log.Printf("%d users left from %d", left, total)
	log.Printf("processing took %v", time.Since(start))
}
