package main

import (
	"math"
)

func busquedaBinariaSinRecursion(arreglo []int, busqueda int) (indice int) {
	izquierda := 0
	derecha := len(arreglo) - 1

	for izquierda <= derecha {
		indiceDelMedio := int(math.Floor((float64(izquierda+derecha) / 2)))

		elementoDelMedio := arreglo[indiceDelMedio]

		if elementoDelMedio == busqueda {
			return indiceDelMedio
		}

		if busqueda < elementoDelMedio {
			derecha = indiceDelMedio - 1
		} else {
			izquierda = indiceDelMedio + 1
		}
	}

	return -1
}
