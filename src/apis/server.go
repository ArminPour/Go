package main


import (
	"net/http"
	"apis/controller"
)


func main() {
    http.HandleFunc("/", controller.HomePage)
    http.ListenAndServe(":8080", nil)
}