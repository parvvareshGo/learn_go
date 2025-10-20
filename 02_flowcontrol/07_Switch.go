// A switch statement is a shorter way to write a sequence of
// if - else statements. It runs the first case whose value is equal to the condition expression.

// Go's switch is like the one in C, C++, Java, JavaScript, and PHP,
// except that Go only runs the selected case, not all the cases that follow. In effect,
// the break statement that is needed at the end of each case in those languages is provided automatically in Go.
// Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

package main

import "fmt"

func main() {
	day := "monday"
	switch day {
	case "monday":
		fmt.Println("Start of the work week.")
	case "wednesday":
		fmt.Println("Midweek day.")
	case "friday":
		fmt.Println("End of the work week.")
	default:
		fmt.Println("Just another day.")
	}
}

// default is the case that runs if none of the other cases match.

// It’s optional — you can leave it out.
// If you don’t include it and no case matches, nothing will be executed.

// Quick notes:

// Only one default is allowed in a switch.

// The position of default doesn’t matter; it can be in the middle or at the end, but it’s usually placed last.

// Like other cases, it doesn’t need a break — Go automatically stops after the matching case.
