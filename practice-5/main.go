package main

import "fmt"

//省略記法を使わない
func addVerbose(x int, y int) int {
	return x + y
}

//足し算
func add(x, y int) int {
	return x + y
}

func main(){
	fmt.Println("省略記法なし:",addVerbose(10,20))
	fmt.Println("省略記法あり:",add(10,20))
}

//省略記法なし: 30
//省略記法あり: 30
