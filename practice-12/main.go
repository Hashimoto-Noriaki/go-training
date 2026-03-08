package main

import "fmt"

func main(){
	//整数から少数への型変換
	var i int = 56
	var f float64 = float64(i)
	fmt.Printf("%.2f\n",f) 

	//少数から整数への型変換
	var x float64 = 2.3456
	var y int = int(x)
	fmt.Printf("%.2f -> %d\n",x,y) 
}
