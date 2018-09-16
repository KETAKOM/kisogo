package main

import "fmt"

//■関数型
//関数を変数に格納するには関数型を使う

func main() {
	//関数型の変数を宣言
	var f func(int, int) int

	//関数型の変数に、関数リテラルの値を代入
	f = func(a int, b int) int {
		return a + b
	}

	//関数型の変数経由で関数を呼び出す
	fmt.Println(f(1, 2))
	
	//関数型の変数に関数(の値)を代入
	f = multiply

	fmt.Println(f(1, 2))
}

func multiply(x int, y int) int {
	return x * y
}
