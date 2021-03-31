package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		img, _ := ioutil.ReadAll(r.Body) // []byte
		fmt.Println(img)

		r.Body.Close()

		ioutil.WriteFile("ejemplo.png", img, 0666) //permisos base

		log.Println("img guardada")

	}
}

func main() {
	setupRoutes()
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":2020", nil)
}
