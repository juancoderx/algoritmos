package main

import "fmt"

func main() {
	numeros := []int{4, 5, 7, 8, 6, 2, 1, 9, 0}
	fmt.Println("hola")
	fmt.Println(Insercion(numeros))
	numeros1 := []int{4, 5, 7, 8, 6, 2, 1, 9, 0}
	fmt.Println(sortManual(numeros1))
}

func Insercion(ListaDesordenada []int) []int {
	var auxiliar int
	for i := 1; i < len(ListaDesordenada); i++ {
		auxiliar = ListaDesordenada[i]
		for j := i - 1; j >= 0 && ListaDesordenada[j] > auxiliar; j-- {
			ListaDesordenada[j+1] = ListaDesordenada[j]
			ListaDesordenada[j] = auxiliar
		}
	}
	return ListaDesordenada
}
