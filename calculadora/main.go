package main

import (
	"fmt"
)

func main() {
	var primernumero int
	var segundonumero int
	var operacion string
	var resultado int

	fmt.Print("Primer numero: ")
	fmt.Scanln(&primernumero)
	fmt.Print("Operacion(+, *, /): ")
	fmt.Scanln(&operacion)
	fmt.Print("Segundo numero: ")
	fmt.Scanln(&segundonumero)

	if operacion == "+" {
		resultado = primernumero + segundonumero
	} else if operacion == "*" {
		resultado = primernumero * segundonumero
	} else if operacion == "/" {
		// importante validar que no se divida entre 0
		if segundonumero == 0 {
			fmt.Println("error: no se puede dividir entre 0")
			return
		}
		resultado = primernumero / segundonumero
	} else {
		fmt.Println("Operacion invalida")
		return
	}

	fmt.Printf("El resultado es %d\n", resultado)
}
