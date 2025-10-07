// Constants
// Constants are declared like variables, but with the const keyword.
// 
// Constants can be character, string, boolean, or numeric values.
// 
// Constants cannot be declared using the := syntax.

package main
import "fmt"

func main() {
	const word = "this is const"
	fmt.Println("Hello", word)

	
	const Truth = true
	fmt.Println("Go rules?", Truth)

}
