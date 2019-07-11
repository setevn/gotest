package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/test", HandleRequest)
	http.ListenAndServe(":8888", nil)
}
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", " text/html;charset=utf-8")
	if r.Method == "POST" {
		r.ParseForm()
		/* username 有两个值，默认取的是第一个的 */
		w.Write([]byte("用户名：" + r.FormValue("username") + "<br/>"))
		w.Write([]byte("<hr/>"))
		names := r.Form["username"]
		w.Write([]byte("username 有两个：" + fmt.Sprintf("%v", names)))
		w.Write([]byte("<hr/>r.Form 的 内 容 ： " + fmt.Sprintf("%v",
			r.Form)))
		w.Write([]byte("<hr/>r.PostForm 的内容：" + fmt.Sprintf("%v",
			r.Form)))
		//r.Form
	} else {
		strBody := `<form action="` + r.URL.RequestURI() + `"
method="post">
用户名：<input name="username" type="text" /><br />
用户名：<input name="username" type="text" /><br />
<input id="Submit1" type="submit" value="submit" />
</form>`
		w.Write([]byte(strBody))
		r.ParseForm()
	}
}
