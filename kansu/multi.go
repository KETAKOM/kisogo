package main

import "fmt"

//四則演算を行う関数の宣言、4つの返り値を持つ
func calc(a int, b int) (int, int, int, float32) {
	
	//戻り値を返す
	return a + b, a - b, a * b, float32(a) / float32(b)
}

func main() {
	//calc関数を呼び出し、多値を受け取る
	add, sub, mul, div := calc(1, 2)

	fmt.Println(add, sub, mul, div)
}
