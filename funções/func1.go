package main

import "fmt"

func calcularMedia(slice []int) float64 {
	total := 0

	for _, valor := range slice {
		total += valor
	}

	media := float64(total) / float64(len(slice))
	return media
}

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	media := calcularMedia(numeros)
	fmt.Println("A média é:", media)
