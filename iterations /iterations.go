package iterations

import (
	"fmt"
)

func main() {
	fmt.Println(repeat("a"))
}

func repeat(char string) string {
	var repeatString string
	for i :=0; i<5; i++ {
		repeatString += char
	}
	return repeatString
}