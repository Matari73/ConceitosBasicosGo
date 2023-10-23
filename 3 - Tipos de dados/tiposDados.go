package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 10000000000
	fmt.Println(numero)

	var numero2 uint = 10000
	fmt.Println(numero2)

	//alias
	//int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	//byte = uint
	var numero4 byte = 123
	fmt.Println(numero4)

	//numero real
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123333.45
	fmt.Println(numeroReal2)

	//string
	var str string = "texto"
	fmt.Println(str)

	str2 := "texto"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	//Bool
	var booleano1 bool
	fmt.Println(booleano1)

	//erro
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
