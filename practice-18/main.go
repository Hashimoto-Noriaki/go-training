package main

import "fmt"

func main(){
	//無限のループ + breakで抜ける
	count := 0
	for {
		count ++
		if count >= 4 {
			break
		}
		fmt.Println(count)
	}
	fmt.Println("最終結果:", count)
}
