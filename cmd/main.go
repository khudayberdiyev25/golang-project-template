package main

import "net/http"

func main() {
	http.ListenAndServe(":5005", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("pong"))
	}))
}
