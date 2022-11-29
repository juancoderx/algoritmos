package main

import (
	"math"
)

func busquedaBinaria(lista []int, item int) int {
	if len(lista) == 0 {
		return -1
	}

	puntoMedio := int(math.Floor(float64(len(lista)) / 2))

	switch {
	case lista[puntoMedio] == item:
		return puntoMedio

	case lista[puntoMedio] < item:
		return busquedaBinaria(lista[puntoMedio+1:], item)

	default:
		return busquedaBinaria(lista[:puntoMedio], item)
	}
}
