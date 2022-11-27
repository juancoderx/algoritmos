package main

func cocktailSort(lista []int) {
	listaAnterior := len(lista) - 1
	for {
		cambio := false
		for i := 0; i < listaAnterior; i++ {
			if lista[i] > lista[i+1] {
				lista[i], lista[i+1] = lista[i+1], lista[i]
				cambio = true
			}
		}
		if !cambio {
			return
		}
		cambio = false
		for i := listaAnterior - 1; i >= 0; i-- {
			if lista[i] > lista[i+1] {
				lista[i], lista[i+1] = lista[i+1], lista[i]
				cambio = true
			}
		}
		if !cambio {
			return
		}
	}
}
