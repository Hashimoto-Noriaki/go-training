package main

import "fmt"

func main(){
	//型推論
	var a = 10
	var b = -5
	var c = 10000

	fmt.Println("===整数リテラル===")
	fmt.Printf("a = %d (型: %T)\n",a, a)
	fmt.Printf("b = %d (型: %T)\n",b, b)
	fmt.Printf("c = %d (型: %T)\n",c, c)
}
