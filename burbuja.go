package main

func bubbleSort(listaNumeros []int) {
	for i := 0; i < len(listaNumeros); i++ {
		for i > 0 && listaNumeros[i-1] > listaNumeros[i] {
			copia := listaNumeros[i]
			listaNumeros[i] = listaNumeros[i-1]
			listaNumeros[i-1] = copia

			i--
		}
	}
}
