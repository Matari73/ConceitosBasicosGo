package main

import "fmt"

type usuario struct {
	nome     string
	idade    int
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int
}

func main() {

	var u usuario
	u.nome = "Mariana"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos bobos", 0}

	u2 := usuario{"Mariana", 21, enderecoExemplo}
	fmt.Println(u2)

	u3 := usuario{nome: "Mariana"}
	fmt.Println(u3)
}
