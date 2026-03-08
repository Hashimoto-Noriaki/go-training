package main

import "fmt"

func main(){
	//型推論(正数)
	var a = 10
	var b = -5
	var c = 10000

	//型推論(少数)
	var d = 3.15
	var e = 4.5678
	var f = 1.2356

	//型推論(文字列)
	var g = "Goの練習"
	var h = true
	var i = 'C'

	fmt.Println("===整数リテラル===")
	fmt.Printf("a = %d (型: %T)\n",a, a)
	fmt.Printf("b = %d (型: %T)\n",b, b)
	fmt.Printf("c = %d (型: %T)\n",c, c)

	fmt.Println("===浮動少数点リテラル===")
	fmt.Printf("d = %.2f (型: %T)\n",d, d)
	fmt.Printf("e = %.2f (型: %T)\n",e, e)
	fmt.Printf("f = %.2f (型: %T)\n",f, f)

	fmt.Println("===文字列リテラル===")
	fmt.Printf("g = %s (型: %T)\n",g, g)
	fmt.Printf("h = %t (型: %T)\n",h, h)
	fmt.Printf("i = %c (型: %T)\n",i, i)
}
