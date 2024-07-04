package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadFile("./hello.txt")
	_, _ = fmt.Fprintln(w, string(body))
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http serve err:%v\n", err)
		return
	}
}
