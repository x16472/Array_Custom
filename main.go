package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Printf("請輸入數值：")
	fmt.Scanln(&num)
	for i := 0; i < num; i++ {
		for e := 0; e < i; e++ {
			fmt.Printf("*")
		}
		for f := num - i; f > 0; f-- {
			fmt.Printf("-")
		}
		fmt.Println(" ")
	}
	for i := 1; i < num; i++ {
		for j := num - i; j > 0; j-- {
			fmt.Printf("*")
		}
		for k := 0; k < i; k++ {
			fmt.Printf("-")
		}
		fmt.Println(" ")
		//蝞銵é蝞
	}
}
