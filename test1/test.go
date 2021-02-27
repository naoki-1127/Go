package main

import (
	"fmt"
)

// 変数:=[要素数]データ型[データ1,データ2,...]

func main() {

	a := [4]string{"yamamoto", "enomoto", "sato"}
	a[3] = "tanaka"

	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
	fmt.Println(a[3])
}
