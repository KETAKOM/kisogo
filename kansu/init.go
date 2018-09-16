package main

import "fmt"

//■init関数
//init関数は特殊な関数で、initを宣言しておくとmain関数より先に、自動で呼び出される
//init関数には初期化処理を記述する。

//初期化処理
func init() {
	fmt.Println("初期化処理")
}

//main関数
func main() {
	fmt.Println("main関数")
}
