package main

import "fmt"

func main() {
	//■可変長パラメータ
	
	//関数の最終パラメータだけは、「可変長パラメータ」として使用できる
	//可変長パラメータでは、関数を呼び出す側で、そのパラメータに渡す値の数を自由に決めれる
	//可変長パラメータを使用するには、関数宣言パラメータの前に「...」を記述する
	
	//可変長パラメータを持つ関数の呼び出し
	holiday(1, "元旦", "成人の日")
	holiday(2, "建国記念日")
	holiday(3, "春分の日")	
}

//可変長パラメータdaysを持つ関数
func holiday(month int, days ...string) {
	fmt.Printf("%d月の祝日は%d日あります。\n", month, len(days))

	for _, day := range days {
		fmt.Println(day)
	}
	fmt.Println()
}
