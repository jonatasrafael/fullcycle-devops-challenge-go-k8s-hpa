package main

import (
	"testing"
)

func TestSomaRaizQuadrada_1(t *testing.T) {
	resultado := somaRaizQuadrada(4)
	if resultado != 6 {
		t.Errorf("Texto esperado: %d, obtido: %f", 6, resultado)
	}
}

func TestSomaRaizQuadrada_2(t *testing.T) {
	resultado := somaRaizQuadrada(0.0001)
	if resultado != 0.0101 {
		t.Errorf("Texto esperado: %f, obtido: %f", 0.0101, resultado)
	}
}

func TestSomaRaizQuadrada_3(t *testing.T) {
	resultado := somaRaizQuadrada(36)
	if resultado != 42 {
		t.Errorf("Texto esperado: %d, obtido: %f", 42, resultado)
	}
}

func TestSomaRaizQuadrada_4(t *testing.T) {
	resultado := somaRaizQuadrada(0.01)
	if resultado != 0.11 {
		t.Errorf("Texto esperado: %f, obtido: %f", 0.11, resultado)
	}
}
