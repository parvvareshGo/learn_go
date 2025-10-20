// Each time you write a `defer`, that function call is pushed onto a **stack**.
// When the function finishes and is about to `return`, the deferred calls are executed in **last-in, first-out (LIFO)** order.

// Simply put:

// ```go
// defer A()
// defer B()
// defer C()
// ```

// Execution order when the function returns: **C → B → A**
package main

import "fmt"

func demo1() {
	defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")
	fmt.Println("work...")
}

func demo2() {
	x := 1
	defer fmt.Println("x in defer =", x) // الان x=1 ذخیره میشه
	x = 2
	fmt.Println("x now =", x)
}
func main() {
	demo1()
	fmt.Println("--------------------------------------------------")
	demo2()
}
