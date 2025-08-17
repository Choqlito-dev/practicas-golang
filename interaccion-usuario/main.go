package main

import "fmt"

func main() {
	var nombre string
	var apellido string

	fmt.Print("Cual es tu nombre: ")
	fmt.Scanln(&nombre)
	fmt.Print("Cual es tu apellido: ")
	fmt.Scanln(&apellido)
	fmt.Println("¡Hola,", nombre, apellido, "espero te encuentres bien")

}
