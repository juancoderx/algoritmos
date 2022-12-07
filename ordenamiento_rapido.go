package main

func quicksort(listadesordenada []int, izq int, der int) {
	pivote := listadesordenada[izq]
	i := izq
	j := der

	var aux int

	for i < j {
		for listadesordenada[i] <= pivote && i < j {
			i++
		}

		for listadesordenada[j] > pivote {
			j--
		}

		if i < j {
			aux = listadesordenada[i]
			listadesordenada[i] = listadesordenada[j]
			listadesordenada[j] = aux
		}
	}

	listadesordenada[izq] = listadesordenada[j]
	listadesordenada[j] = pivote

	if izq < j-1 {
		quicksort(listadesordenada, izq, j-1)
	}

	if j+1 < der {
		quicksort(listadesordenada, j+1, der)
	}
}
