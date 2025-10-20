// Switch without a condition is the same as switch true.

// This construct can be a clean way to write long if-then-else chains.

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Current time:", t.Format("15:04"))
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
