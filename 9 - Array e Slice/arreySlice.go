package main

import (
	"fmt"
)

func main() {
	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]string{"posição 1", "posição 2", "posição 3", "posição 4", "posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice = append(slice, 6)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	//arrays internos
	fmt.Println("----------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) //capacidade

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

}
