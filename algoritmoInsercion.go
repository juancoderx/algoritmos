package main

func insertionSort(listaDesordenada []int) []int {
	var auxiliar int
	for i := 1; i < len(listaDesordenada); i++ {
		auxiliar = listaDesordenada[i]
		for j := i - 1; j >= 0 && listaDesordenada[j] > auxiliar; j-- {
			listaDesordenada[j+1] = listaDesordenada[j]
			listaDesordenada[j] = auxiliar
		}
	}
	return listaDesordenada
}
