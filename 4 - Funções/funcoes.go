package main

import "fmt"

func somar(n1, n2 int) int {
	return n1 + n2
}

func calculos(n1 int, n2 int) (int, int) {
	soma := n1 + n2
	subt := n1 - n2

	return soma, subt
}

func main() {
	soma := somar(1, 1)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	resultSoma, resultSubt := calculos(10, 10)
	fmt.Println(resultSoma, resultSubt)

	resultSoma, _ = calculos(10, 10)
	fmt.Println(resultSoma)
}
