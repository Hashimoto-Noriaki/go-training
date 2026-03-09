package main

import "fmt"

func main(){
	// 基本的な if-else
	age := 15
	if age >= 18 {
		fmt.Println("大人です。")
	} else {
		fmt.Println("未成年です。")
	}

	//else ifの連鎖
	score := 65
	if score >= 80 {
		fmt.Println("優秀")
	} else if  score >=60 {
		fmt.Println("普通")
	} else {
		fmt.Println("もっと頑張れ")
	}
}
