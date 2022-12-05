package main

import (
	"testing"
)

// Se ejecutan con: go test -bench=.

func BenchmarkBusquedaBinaria_DebajoDeLaMitad(b *testing.B) {
	cantidad := 10_000_000

	data := make([]int, cantidad)

	for i := 0; i < len(data); i++ {
		data[i] = i
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		busquedaBinaria(data, 3_000_000, 0)
	}
}

func BenchmarkBusquedaBinariaSinRecursion_DebajoDeLaMitad(b *testing.B) {
	cantidad := 10_000_000

	data := make([]int, cantidad)

	for i := 0; i < len(data); i++ {
		data[i] = i
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		busquedaBinariaSinRecursion(data, 3_000_000)
	}
}

func BenchmarkBusquedaBinaria_ArribaDeLaMitad(b *testing.B) {
	cantidad := 10_000_000

	data := make([]int, cantidad)

	for i := 0; i < len(data); i++ {
		data[i] = i
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		busquedaBinaria(data, 9_000_000, 0)
	}
}

func BenchmarkBusquedaBinariaSinRecursion_ArribaDeLaMitad(b *testing.B) {
	cantidad := 10_000_000

	data := make([]int, cantidad)

	for i := 0; i < len(data); i++ {
		data[i] = i
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		busquedaBinariaSinRecursion(data, 9_000_000)
	}
}
