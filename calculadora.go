package main

import (
	"errors"
	"fmt"
)

// Função para adicionar dois números
func Add(a, b float64) float64 {
	return a + b
}

// Função para subtrair dois números
func Subtract(a, b float64) float64 {
	return a - b
}

// Função para multiplicar dois números
func Multiply(a, b float64) float64 {
	return a * b
}

// Função para dividir dois números
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	// Testando as funções
	fmt.Println("Resultado da adição:", Add(10, 5))
	fmt.Println("Resultado da subtração:", Subtract(10, 5))
	fmt.Println("Resultado da multiplicação:", Multiply(10, 5))

	// Teste de divisão com tratamento de erro
	result, err := Divide(10, 5)
	if err != nil {
		fmt.Println("Erro ao dividir:", err)
	} else {
		fmt.Println("Resultado da divisão:", result)
	}

	// Teste de divisão por zero com tratamento de erro
	result, err = Divide(10, 0)
	if err != nil {
		fmt.Println("Erro ao dividir:", err)
	} else {
		fmt.Println("Resultado da divisão:", result)
	}
}
