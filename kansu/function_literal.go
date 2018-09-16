package main

import "fmt"


//■関数リテラル
//関数リテラルとは匿名関数のである。
//通常の関数とは異なり、関数リテラルは他の関数内に記述する
//関数リテラルは、関数リテラル外で宣言された変数にアクセス可能→クロージャ


func main() {
	val := 123

	//関数リテラルの呼び出しを同時に行う
	func(i int) {
		//関数リテラル外の変数にアクセス可能
		fmt.Println(i * val)
	}(10)

	//関数リテラルを変数に代入
	f := func(s string) {
		fmt.Println(s)
	}

	//関数の呼び出し
	f("hogehoge")
}
