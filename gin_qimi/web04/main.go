package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//2.解析模板
	t, err := template.ParseFiles("./tpl/index.tmpl")
	if err != nil {
		fmt.Println("Parse Tpl err")
		return
	}

	//3.渲染模板
	err = t.Execute(w, "一张小饼")
	if err != nil {
		fmt.Println("Tpl Execute err")
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Serve Err :%v", err)
		return
	}
}
