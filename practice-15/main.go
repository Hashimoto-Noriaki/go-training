package main

import "fmt"

func main(){
	//0+1+2+...+9の合計
	sum := 0
	for i :=0; i < 10; i++ {
		sum+= i
	}
	fmt.Println("合計:", sum)//45

	//カウントダウン
	for i:=3; i > 0; i-- {
		fmt.Println(i)
	}
	fmt.Println("golang楽しい")
	//3
	//2
	//1
	//golang楽しい
}
