package main

import "fmt"

func main(){
	// 基本的な switch
	day := "金曜日"
	switch day{
	case "月曜日":
		fmt.Println("仕事はじめ")
	case "火曜日", "水曜日", "木曜日":
		fmt.Println("平日")
	case "金曜日":
		fmt.Println("花金")
	default:
		fmt.Println("おやすみ")
	}
}
