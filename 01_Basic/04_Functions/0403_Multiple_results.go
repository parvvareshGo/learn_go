// 
// Multiple results
// A function can return any number of results.
// 
// The swap function returns two strings.

package main
import "fmt"

func return_number(x, y int) (int, int) {
	return x, y
}

func main() {
	a, b := return_number(12, 4)
	fmt.Println(a, b)
}
