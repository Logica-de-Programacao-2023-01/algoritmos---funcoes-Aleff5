package main

import "fmt"

func findPosition(slice []int, value int) int {
	for i, num := range slice {
		if num == value {
			return i
		}
	}
	return -1
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	value := 3
	position := findPosition(slice, value)
	fmt.Printf("Valor encontrado na posição: %d\n", position)
}
