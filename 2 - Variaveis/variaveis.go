package main

import "fmt"

func main() {
	var variavel1 string = "variavel 1"
	variavel2 := "variavel 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "lalalala"
		variavel4 string = "lalalala"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variável 5", "variável 6"
	fmt.Println(variavel5, variavel6)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
