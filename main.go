package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	fmt.Printf("Starting server at localhost:8080 from %s\n", dir)
	filesystem := http.FileServer(http.Dir(dir))
	http.Handle("/", filesystem)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
