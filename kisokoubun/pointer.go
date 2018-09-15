package main

import "fmt"

func main() {
	//変数が存在するメモリ上の場所→アドレス
	//そのアドレスを格納可能な変数→ポインタ

	//int型のポインタ変数を定義
	var ptr *int

	var i = 122345

	//&で変数のアドレスを取得する
	ptr = &i
	fmt.Println(i)
	fmt.Println(&i)
	fmt.Println(ptr)

	//ポインタ経由で変数iの値を変更する
	*ptr = 23456
	fmt.Println(*ptr)
	fmt.Println(i)
}
