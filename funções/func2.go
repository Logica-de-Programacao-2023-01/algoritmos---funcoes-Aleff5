package main

import "fmt"

func main() {
	vogais := "aeiouAEIOU"
	count := 0

	for _, char := range texto {
		for _, vogal := range vogais {
			if char == vogal {
				count++
				break
			}
		}
	}

	return count
}

func main() {
	texto := "Hello, World!"
	quantidadeVogais := contarVogais(texto)
	fmt.Println("Quantidade de vogais:", quantidadeVogais)
}
