package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	resultado := Add(10, 5)
	if resultado != 15 {
		t.Errorf("Esperado 15, mas obteve %f", resultado)
	}
}

func TestSubtract(t *testing.T) {
	resultado := Subtract(10, 5)
	if resultado != 5 {
		t.Errorf("Esperado 5, mas obteve %f", resultado)
	}
}

func TestMultiply(t *testing.T) {
	resultado := Multiply(10, 5)
	if resultado != 50 {
		t.Errorf("Esperado 50, mas obteve %f", resultado)
	}
}

func TestDivide(t *testing.T) {
	resultado, err := Divide(10, 5)
	if err != nil || resultado != 2 {
		t.Errorf("Esperado 2, mas obteve %f com erro %v", resultado, err)
	}

	_, err = Divide(10, 0)
	if err == nil {
		t.Error("Esperado erro, mas obteve nada")
	}
}
