package main

import "fmt"

func main() {
	var name string = "ana"
	fmt.Println(name)
	
	var ponteiro *string = &name
	fmt.Println(ponteiro)

	*ponteiro = "Eduardy"
	fmt.Println(name) 
}