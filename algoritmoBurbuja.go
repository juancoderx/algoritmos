package main

func sortManual(paises []int) []int {
	for i := 0; i < len(paises); i++ {
		for i > 0 && paises[i-1] > paises[i] {
			copia := paises[i]
			paises[i] = paises[i-1]
			paises[i-1] = copia

			i--
		}
	}
	return paises
}
