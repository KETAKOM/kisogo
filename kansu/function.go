package main

import "fmt"

func main() {
	//■関数
	//golangの関数の大きな特徴は、関数の戻り値として、複数の値を一度に返せることである
	
	answer := plus(1, 2)

	fmt.Println(answer)	

}

func plus(a int, b int) int {
	return a + b
}
