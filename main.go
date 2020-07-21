package main

import (
	"fmt"

	algorithms "github.com/RiccardoBiosas/golang-bioinformatics-algorithms/algorithms"
)

func main() {
	for i := 1; i <= 2; i++ {
		result := algorithms.RunAlgorithms(i)
		fmt.Println(result)
	}
}
