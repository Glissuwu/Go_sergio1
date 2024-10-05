package main

import "fmt"

func main() {

	

	var centimetro, metro float64 = 0, 0
	fmt.Println("CONVERSOR DE CM A METROS")
	fmt.Println("INGRESE CM")
	fmt.Scanln(&centimetro)

	metro = centimetro / 100
	fmt.Println("la conversion de", centimetro, "cm es de", metro, "m")
}
