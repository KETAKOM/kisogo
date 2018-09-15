package main

import "fmt"

func main() {
	//int型からInteger型を宣言
	//type 型名 元になる型
	type myInteger int

	var i myInteger = 123

	//元の型と同じ演算が可能
	i += 1

	fmt.Println(i)

	//新しい構造体を作成
	type myStruct struct {
		a int
		b int
	}
}
