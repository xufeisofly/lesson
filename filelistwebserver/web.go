package main

import (
	"lession/filelistwebserver/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/list/", handler.ErrWrapper(handler.HandleList))
	http.ListenAndServe(":8889", nil)
}
