package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//Concorrência != paralelismo
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Olá mundo")
		waitGroup.Done() //-1 do Add
	}()

	go func() {
		escrever("Programando em Go")
		waitGroup.Done() //-1 do Add
	}()

	waitGroup.Wait() // 0
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
