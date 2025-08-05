package main 

import "fmt"

func get(x string, y string) (string , string){
	return x, y
}

func main(){
	a, b := get("hi","alireza")
	fmt.Println(a, b)
}
