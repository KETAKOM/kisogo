package main

import "fmt"

func main () {

LBL:
	for i := 0; i < 10; i++ {
		
		switch {
			case i == 5:
				//forへbreak
				break LBL
			default:
				fmt.Println(i)
		}
	}


}
