package main

import "fmt"

func main() {
	//関数から返却された多値を直接他の関数に渡す
	f2(f1())	
}

//多値を返す関数
func f1() (int, string, float32){
	return 0, "xyz", 3.14 
}

func f2(a int, b string, c interface{}) {
	fmt.Println(a, b, c)
}

