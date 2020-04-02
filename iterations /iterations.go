package iterations

import (
	"fmt"
)

func main() {
	fmt.Println(repeat("a", 4))
}

func repeat(char string, numberToRepeat int) string {
	var repeatString string
	for i :=0; i<numberToRepeat; i++ {
		repeatString += char
	}
	return repeatString
}