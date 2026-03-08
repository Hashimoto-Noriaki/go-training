package main

import "fmt"

func main(){
	name := "大谷"
	age := 25
	isPlayer := true

	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("isPlayer:", isPlayer)

	fmt.Printf("nameの型: %T\n", name)
	fmt.Printf("ageの型: %T\n", age)
	fmt.Printf("isPlayerの型: %T\n", isPlayer)
}
