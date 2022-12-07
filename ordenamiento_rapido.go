package main

func quicksort(listadesordenada []int, izq int, der int) {
	pivote := listadesordenada[izq]
	// tomamos primer elemento como pivote
	i := izq // i realiza la búsqueda de izquierda a derecha
	j := der
	// j realiza la búsqueda de derecha a izquierda
	var aux int

	for i < j { // mientras no se crucen las búsquedas
		for listadesordenada[i] <= pivote && i < j {
			i++
		}
		// busca elemento mayor que pivote
		for listadesordenada[j] > pivote {
			j--
		}
		// busca elemento menor que pivote
		if i < j { // si no se han cruzado
			aux = listadesordenada[i] // los intercambia
			listadesordenada[i] = listadesordenada[j]
			listadesordenada[j] = aux
		}
	}

	listadesordenada[izq] = listadesordenada[j] // se coloca el pivote en su lugar de forma que tendremos
	listadesordenada[j] = pivote

	// los menores a su izquierda y los mayores a su derecha
	if izq < j-1 {
		quicksort(listadesordenada, izq, j-1) // ordenamos subarray izquierdo
	}

	if j+1 < der {
		quicksort(listadesordenada, j+1, der) // ordenamos subarray derecho
	}
}
