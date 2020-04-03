package arrays

import (
	"fmt"
)

func main() {
	fmt.Println()
}

func sum(numbers []int) int {
	var arraySum int
	for _, value := range numbers {
		arraySum += value
	}
	return arraySum
}

func sumAll(slices ...[]int) []int {
	var slicesSum []int
	for _, slice := range slices {
		slicesSum = append(slicesSum, sum(slice))
	}
	return slicesSum
}

func sumTails(slices ...[]int) []int {
	var slicesSum []int
	for _, slice := range slices {
		if len(slice) == 0 {
			slicesSum = append(slicesSum, 0)
		} else {
			sliceWithOutHead := slice[1:]
			slicesSum = append(slicesSum, sum(sliceWithOutHead))
		}
	}
	return slicesSum
}
