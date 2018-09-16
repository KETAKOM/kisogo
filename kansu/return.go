package main

import "fmt"

//■戻り値に名前をつける
//golangの戻り値の特徴は、多値を返せるだけでなく、返り値に名前をつけることができる

//四則演算を行う関数の宣言
func calc(a int, b int) (add int, sub int, mul int, div float32) {
	add = a + b
	sub = a - b
	mul = a * b
	div = float32(a) / float32(b)

	//この場合、returnだけで戻り値に指定された変数の値を返す
	//変数に値を代入していない場合は、ゼロ値が自動的の値が返る
	return
} 


func main() {
	//calc関数を呼び出す
	fmt.Println(calc(1, 2))
}
