package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		img, _ := ioutil.ReadAll(r.Body) // []byte

		sum := sha256.Sum256(img) //codigo sha
		fmt.Printf("%x", sum)

		r.Body.Close()

		ioutil.WriteFile("ejemplo.png", img, 0666) //permisos base

		//log.Println("img guardada")

	}
}

func main() {
	setupRoutes()
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":2020", nil)
}
