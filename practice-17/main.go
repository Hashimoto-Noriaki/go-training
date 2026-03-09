package main

import "fmt"

func main(){
	//条件だけのループ(while風)
	sum := 1
	for sum < 100 {
		sum+= sum
	}
	fmt.Println("最終結果:", sum)

	//カウントアップ(条件だけで制御)
	count :=0
	for count < 5 {
		fmt.Println(count)
		count ++ //count = count + 1 と同じ 
	}
}
