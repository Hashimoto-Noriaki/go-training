package main

import (
	"fmt"
)

func main(){
	//Println:そのまま並べる
	fmt.Println("name:", "Messi", "age:", 38)

	//Printf: フォーマット指定で表示
	name := "大谷"
	age := 31
	height := 193.5
	fmt.Printf("%sさんは%d歳、身長は%.1fcmです\n", name, age, height)

	// 型の確認(%T)
	fmt.Printf("name の型: %T\n", name)
	fmt.Printf("age の型: %T\n", age)
}

// name: Messi age: 38
// 大谷さんは31歳、身長は193.5cmです
// name の型: string
// age の型: int
