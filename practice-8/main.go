package main

import "fmt"

// 割り算の商と余りを返す
func divmod(a, b int) (int, int){
	quotient := a / b //商
	remainder := a % b //余り
	return quotient, remainder
}

func main(){
	q, r := divmod(17,5)
	fmt.Printf("17 ÷ 5 = %d 余り %d\n", q, r)

	//別の例
	q2,r2 := divmod(100,7)
	fmt.Printf("100 ÷ 7 = %d 余り %d\n", q2,r2)
}
