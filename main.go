package main

import(
    "io"
	"log"
	"net/http"
)

func main(){
 //设置路由
   http.HandleFunc("/", sayHello)
   err:=http.ListenAndServe(":8080",nil)// 设置监听的端口
   if err!=nil{
       log.Fatal(err);
   }
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world, this is version 1.")
}