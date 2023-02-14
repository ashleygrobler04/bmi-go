package main

import (
	"fmt"
	"math"
)

func calcBmi(height float64, weight float64) string {
	res := fmt.Sprint(math.Round(weight / (height * height)))
	return res
}

func main() {
	result := calcBmi(1.82, 74.9)
	fmt.Println(result)
}
