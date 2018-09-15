package main

import "fmt"

func main() {

	//オブジェクト指向には「クラス」の概念があるが、golangには「クラス」はなく「型」しかない

	//bool型の変数を宣言
	var b bool
	
	//bool型の変数にリテラル定数のtrueとfalseを代入
	b = true
	b = false

	//bool型の変数に論理演算した結果を代入
	b = true || false
	
	//出力
	fmt.Println(b)
}
