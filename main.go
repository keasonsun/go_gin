package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc(
		"/hello",
		func(w http.ResponseWriter, r *http.Request) {
			//解析模板
			t, err2 := template.ParseFiles("./tmpl/hello.tmpl")
			//渲染模板
			if err2 != nil {
				panic(err2)
			}

			t.Execute(w, "你好这是我的第一个go模板")

		})
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
}
