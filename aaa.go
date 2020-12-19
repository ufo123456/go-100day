package main

import "fmt"

func main() {
	//vi := "hoda"
	//a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sitys := [...]string{"北京", "山海", "江西"}
	for i := 0; i < len(sitys); i++ {
		fmt.Println(sitys[i])
	}
	//fmt.Println(a1)
	//fmt.Println(a1)
	//fmt.Printf(vi)
}
