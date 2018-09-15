package main

import "fmt"

func main() {

	//goto文は指定したラベルをつけた文にプログラムの実行を移動する
	//goto文の使用には制限があり、それまでスコープ外だった変数がスコープ内に入る場合はコンパイルエラーになる

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			//EVENへジャンプ
			goto EVEN
		}
		fmt.Println("add", i)
		
		EVEN:
	}

}
