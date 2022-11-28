package main

import "fmt"

func main() {
	fmt.Println("Algoritmo de Ordenamiento por Insercion:")

	numeros := []int{23, 27, 43, 38, 33, 34, 36, 48, 32, 25, 39, 34, 23, 47, 51}

	fmt.Println(" Antes:")
	fmt.Println("  ", numeros)
	fmt.Println(" Despues:")
	insertionSort(numeros)
	fmt.Println("  ", numeros)
	fmt.Println()

	// ---

	fmt.Println("Algoritmo de Ordenamiento Burbuja:")

	numeros1 := []int{21, 19, 35, 20, 24, 32, 81, 35, 22, 10, 13, 27, 21, 16, 30}

	fmt.Println(" Antes:")
	fmt.Println("  ", numeros1)
	fmt.Println(" Despues:")
	bubbleSort(numeros1)
	fmt.Println("  ", numeros1)
	fmt.Println()

	// ---

	fmt.Println("Algoritmo de Ordenamiento Burbuja Bidireccional:")

	numeros2 := []int{61, 29, 31, 62, 21, 25, 36, 22, 19, 64, 74, 59, 45, 58, 29}

	fmt.Println(" Antes:")
	fmt.Println("  ", numeros2)
	fmt.Println(" Despues:")
	cocktailSort(numeros2)
	fmt.Println("  ", numeros2)
	fmt.Println()

	// ---

	fmt.Println("Algoritmo de Ordenamiento por Cuentas:")

	numeros3 := []int{10, 46, 71, 69, 57, 29, 15, 57, 31, 53, 2, 61, 5, 38, 52}

	fmt.Println(" Antes:")
	fmt.Println("  ", numeros3)
	fmt.Println(" Despues:")

	fmt.Println("  ", countSort(numeros3))
	fmt.Println()

	// ---

	fmt.Println("Busqueda Secuencial:")

	numeros4 := []int{10, 100, 1000, 20, 200, 2000}

	fmt.Println(" Datos:", numeros4)
	fmt.Println(" 200 - Indice:", busquedaSecuencial(numeros4, 200))
	fmt.Println()

	// ---

	fmt.Println("Busqueda Binaria:")

	numeros5 := []int{100, 200, 300, 400, 800}

	fmt.Println(" Datos:", numeros5)
	fmt.Println(" 100 - Indice:", busquedaBinaria(numeros5, 100, 0))
	fmt.Println(" 200 - Indice:", busquedaBinaria(numeros5, 200, 0))
	fmt.Println(" 300 - Indice:", busquedaBinaria(numeros5, 300, 0))
	fmt.Println(" 400 - Indice:", busquedaBinaria(numeros5, 400, 0))
	fmt.Println(" 500 - Indice:", busquedaBinaria(numeros5, 500, 0))
}
