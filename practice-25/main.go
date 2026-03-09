package main

import "fmt"

func main(){
	defer fmt.Println("1. 最初にdefer(最後に実行)")
	defer fmt.Println("2. 次にdefer")
	defer fmt.Println("3. 最後にdefer(最初に実行)")
	fmt.Println("通常の処理")
}
