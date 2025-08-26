package main

import (
	"fmt"
	"os"
)

func main() {
	var filename string
	var content string
	fmt.Print("Nombre del nuevo archivo: ")
	fmt.Scan(&filename)

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error: al crear el archio")
		return
	}
	defer file.Close()

	fmt.Printf("El archivo %s se creo exitosamente\n", filename)

	fmt.Print("Contenido del archivo: ")
	fmt.Scan(&content)

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error: al escribir en el archivo")
		return
	}

	fmt.Println("Contenido escrito exitosamente")

}
