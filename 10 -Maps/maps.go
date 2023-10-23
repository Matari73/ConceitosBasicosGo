package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"segundo":  "Pedro",
		},
		"curso": {
			"nome":      "Engenharia",
			"faculdade": "UTFPR",
		},
	}
	fmt.Println(usuario2)

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}
	fmt.Println(usuario2)
}
