package main

import(
    "io"
	"log"
	"net/http"
)

func main(){
   mux := http.NewServeMux()
   mux.Handle("/",&myHandler{})
   mux.HandleFunc("/hello", sayHello)
   err:=http.ListenAndServe(":8080",mux)// 设置监听的端口
   if err!=nil{
       log.Fatal(err);
   }
}
type myHandler struct{}

func (*myHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "URL: "+r.URL.String())
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world, this is version 1.")
}