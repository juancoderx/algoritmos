package main

import (
	"strings"
	"testing"
)

// go test -benchmem -bench=StringBuilder

const cantidadPrueba = 100_00

func BenchmarkStringBuilder_String(*testing.B) {
	var data string

	for i := 0; i < cantidadPrueba; i++ {
		data += string(rune(i))
	}

	_ = data
}

func BenchmarkStringBuilder_SliceRune(*testing.B) {
	var data []rune //nolint:prealloc // Ya se.

	for i := 0; i < cantidadPrueba; i++ {
		data = append(data, rune(i))
	}

	_ = data
}

func BenchmarkStringBuilder_Builder(*testing.B) {
	var data strings.Builder

	for i := 0; i < cantidadPrueba; i++ {
		data.WriteRune(rune(i))
	}

	_ = data
}

func BenchmarkStringBuilder_MakeWithCap(*testing.B) {
	data := make([]rune, 0, cantidadPrueba)

	for i := 0; i < cantidadPrueba; i++ {
		data = append(data, rune(i))
	}

	_ = data
}

func BenchmarkStringBuilder_MakeWithLenAndCap(*testing.B) {
	data := make([]rune, cantidadPrueba)

	for i := 0; i < cantidadPrueba; i++ {
		data[i] = rune(i)
	}

	_ = data
}
