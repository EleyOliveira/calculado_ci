package main

import "testing"

func TestMultiplicacao(t *testing.T) {
	resultado := Multiplicacao(10, 10)

	if resultado != 100 {
		t.Errorf("Resultado da multiplicação esperado é %d, porém o valor foi %d", 100, resultado)
	}
}
