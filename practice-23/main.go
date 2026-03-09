package main

import "fmt"

func main(){
	// 最初に一致した　case で止まる(上から評価)
	i := 0
	switch i {
	case 0:
		fmt.Println("case 0: 一致")
	case checkCase(1);
		fmt.Println("case 1: 一致")
	}
}

func checkCase(n int) int {
	fmt.Println("  checkCase(%d) が呼ばれました\n", n)
	return n
}
