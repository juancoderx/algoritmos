package main

import (
	"math"
)

func BusquedaBinaria(lista []int, item int, pos int) int {
	if len(lista) <= 0 {
		return -1
	}

	puntoMedio := int(math.Floor(float64(len(lista)) / 2))

	switch {

	case lista[puntoMedio] == item:
		return pos + puntoMedio

	case lista[puntoMedio] < item:
		return BusquedaBinaria(lista[puntoMedio+1:], item, pos+puntoMedio+1)

	default:
		return BusquedaBinaria(lista[:puntoMedio], item, pos)

	}
}
