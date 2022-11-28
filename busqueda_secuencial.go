package main

func busquedaSecuencial(numeros []int, valor int) int {
	for k, v := range numeros {
		if v == valor {
			return k
		}
	}

	return -1
}
