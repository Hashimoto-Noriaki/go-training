//名前付き戻り値
package main

import "fmt"

//通常版
func splitNormal(sum int) (int, int) {
	x := sum * 4 / 9
	y := sum - x
	return x, y
}

//名前付き版
func splitNamed(sum int) (x int, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main(){
	a1, b1 := splitNormal(17)
	fmt.Println("通常版:", a1, b1)

	a2, b2 := splitNamed(17)
	fmt.Println("名前付き版:", a2, b2)
}
