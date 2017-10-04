package main

import (
	"fmt"
	"net/http"

	Entities "../Entities"
)

const (
	port = ":8000"
)

var (
	calls = 0
)

//HelloWorld funcion a enrutar
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	var person Entities.Person
	person.FirstName = "David"
	person.LastName = "Silva"

	calls++
	fmt.Fprintf(w, "Hola %s %d veces.\n", person.FirstName, calls)
}

func main() {
	fmt.Printf("Servidor corriendo en http://localhost%v.\n", port)
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(port, nil)
}
