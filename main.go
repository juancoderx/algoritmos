package main

import "fmt"

func main() {
	numeros := []int{23, 27, 43, 38, 33, 34, 36, 48, 32, 25, 39, 34, 23, 47, 51}
	fmt.Println("Algoritmo Inserción")
	fmt.Println(insertionSort(numeros))

	numeros1 := []int{21, 19, 35, 20, 24, 32, 81, 35, 22, 10, 13, 27, 21, 16, 30}
	fmt.Println("Algoritmo de Burbuja")
	fmt.Println(bubbleSort(numeros1))

	numeros2 := []int{61, 29, 31, 62, 21, 25, 36, 22, 19, 64, 74, 59, 45, 58, 29}
	fmt.Println("Algoritmo de Burbuja Bidireccional")
	cocktailSort(numeros2)
	fmt.Println(numeros2)

	numeros3 := []int{10, 46, 71, 69, 57, 29, 15, 57, 31, 53, 2, 61, 5, 38, 52}
	fmt.Println("Algoritmo por Cuentas")
	fmt.Println(countSort(numeros3))

	numero4 := []int{10, 100, 1000, 20, 200, 2000}
	fmt.Println("Busqueda secuencial, encontrara la posicion del 200:", BusquedaSecuencial(numero4, 200))

	nums := []int{100, 200, 300, 400, 800}
	fmt.Println("Busqueda Binaria")
	fmt.Println("100", BusquedaBinaria(nums, 100, 0))
	fmt.Println("200", BusquedaBinaria(nums, 200, 0))
	fmt.Println("300", BusquedaBinaria(nums, 300, 0))
	fmt.Println("400", BusquedaBinaria(nums, 400, 0))
	fmt.Println("500", BusquedaBinaria(nums, 500, 0))
}
