package main

import "fmt"

func main() {

	list := NewList()

	fmt.Println(list.Size())
	list.Append("A")
	//list.Append("B")
	//list.Append("C")

	fmt.Println(list)

	fmt.Println(list.Size())

}
