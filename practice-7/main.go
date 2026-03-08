package main

import "fmt"

var (
	appName string = "HashiApp"
	appVersion string = "22.2.1"
	maxUsers int = 10000
)

func printInfo(){
	fmt.Println("アプリ名:", appName)
	fmt.Println("バージョン名:", appVersion)
	fmt.Println("最大ユーザー数:", maxUsers)
}

func main() {
	printInfo()

	//関数内でも使える
	fmt.Println("---")
	fmt.Printf("%s (v%s)\n", appName, appVersion)
}
