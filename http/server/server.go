package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(b))
}
func f2(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()

	name := param.Get("name")
	age := param.Get("age")
	fmt.Println("request method:", r.Method)
	fmt.Println("name:", name)
	fmt.Println("age:", age)
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println("body", body)
	w.Write([]byte("read OK"))

}
func main() {
	http.HandleFunc("/post/go/1", f1)
	http.HandleFunc("/get/peg", f2)
	http.ListenAndServe("0.0.0.0:9999", nil)
}
