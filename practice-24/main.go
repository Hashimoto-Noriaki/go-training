package main

import "fmt"

func main(){
	defer fmt.Println("3. 最後に実行")
	fmt.Println("1. 最初に実行")
	fmt.Println("2. 次に実行")
}

//1. 最初に実行
//2. 次に実行
//3. 最後に実行
