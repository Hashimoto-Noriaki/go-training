package main

import "fmt"

func main(){
	//初期化と後処理を省略(条件だけ)
	sum := 1
	for ; sum < 100; {
		sum += sum
	}
	fmt.Println("最終結果:", sum)

	// 後処理だけ省略(ループ内で更新)
	for j := 0; j < 5; {
		fmt.Println(j)
		j += 2
	}
}

// 最終結果: 128
// 0
// 2
// 4
