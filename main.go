package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("0.0.0.0:8110")
	http.ListenAndServe("0.0.0.0:8110",http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("xx"))
	}))
}
