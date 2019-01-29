package main


import (
	"net/http"
	"apis/controller"
)


func main() {
    http.HandleFunc("/", controller.HelloWorld)
    http.ListenAndServe(":8080", nil)
}