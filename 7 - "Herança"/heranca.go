package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
	altura    float64
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"João", "Pedro", 20, 1.78}
	fmt.Println(p1)

	e1 := estudante{p1, "engenharia", "UTFPR"}
	fmt.Println(e1.altura)
}
