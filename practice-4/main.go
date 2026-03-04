package main

import "fmt"

//足し算
func add(x int, y int) int {
	return x + y
}

//掛け算
func multply(x int, y int) int {
	return x * y 
}

// 結果表示
func printResult(label string, value int) {
	fmt.Printf("%s: %d\n", label, value)
}

func main(){
	a := 10
	b := 5

	sum := add(a,b)
	product := multply(a,b)

	printResult("足し算",sum)
	printResult("掛け算",product)
}

//足し算: 15
//掛け算: 50