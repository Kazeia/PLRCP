package main

import (
	"fmt"
	"net/http"
)

const (
	port = ":8000"
)

var (
	calls = 0
)

//HelloWorld funcion a enrutar
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	calls++
	fmt.Fprintf(w, "Hola mundo en %d veces.\n", calls)
}

func main() {
	fmt.Printf("Servidor corriendo en http://localhost%v.\n", port)
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(port, nil)
}
