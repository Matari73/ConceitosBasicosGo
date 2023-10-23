package main

import "fmt"

func main() {
	// aritmeticos
	soma := 1 + 2
	subt := 1 - 2
	divis := 10 / 2
	multip := 10 * 2
	restoDivis := 10 % 2

	fmt.Println(soma, subt, divis, multip, restoDivis)

	//não pode realizar operações com variaveis de tipo diferente

	//atribuição
	var variavel1 string = "string"
	variavel2 := "string2"
	fmt.Println(variavel1, variavel2)

	//relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	//logicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!falso)
	fmt.Println(!verdadeiro)

	//unário
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20

	numero *= 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)

	//ternario
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)
}
