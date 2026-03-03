// package main

// import (
// 	"fmt"
// )

// func main(){
// 	//Println:そのまま並べる
// 	fmt.Println("name:", "Messi", "age:", 38)

// 	//Printf: フォーマット指定で表示
// 	name := "大谷"
// 	age := 31
// 	height := 193.5
// 	fmt.Printf("%sさんは%d歳、身長は%.1fcmです\n", name, age, height)

// 	// 型の確認(%T)
// 	fmt.Printf("name の型: %T\n", name)
// 	fmt.Printf("age の型: %T\n", age)
// }

// package main

// import (
// 	"fmt"
// 	_"math"
// )

// func main(){
// 	// fmt.Printf("Now you have %g problems.", math.Sqrt(7))
// 	fmt.Printf("Now you have %g problems.")
// }

// import (
// 	"fmt"
// 	"math"
// )

//func main(){
	// fmt.Printf("Now you have %g problems.", math.Sqrt(7))
	// fmt.Printf("Now you have %g problems.")
	//fmt.Println(math.Pi)

// package main 

// import "fmt"

// func add(x int, y int) int {
// 	return x + y
// }

// func main(){
// 	fmt.Println(add(12,13))
// }

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
