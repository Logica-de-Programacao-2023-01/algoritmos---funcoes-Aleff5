package main

import (
	"errors"
	"fmt"
	"strings"
)

func joinStrings(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("o slice está vazio")
	}

	result := strings.Join(slice, ",")
	return result, nil
}

func main() {
	slice := []string{"foo", "bar", "baz"}
	joined, err := joinStrings(slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings concatenadas:", joined)
	}

	emptySlice := []string{}
	emptyJoined, emptyErr := joinStrings(emptySlice)
	if emptyErr != nil {
		fmt.Println("Erro:", emptyErr)
	} else {
		fmt.Println("Strings concatenadas:", emptyJoined)
	}
}
