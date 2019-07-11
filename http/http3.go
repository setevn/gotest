package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	//更底层实现
	server := &http.Server{
		Addr:        ":8080",
		Handler:     &myHandler{}, //Handler的底层实现
		ReadTimeout: 5 * time.Second,
	}
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/hello"] = sayHello
	mux["/bye"] = sayBye

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//HandleFunc
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "url:"+r.URL.String())
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}
func sayBye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "bye hello world")
}

func drink(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "喝饮料！")
}
