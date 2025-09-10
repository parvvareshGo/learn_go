package main

import "fmt"

func test(x, y int) (int, int)  {
	return x, y
}

func main() {
	a,b := test(13, 14)
	fmt.Println(a, b)
}
