package main

// 複数のパッケージをimportする場合の書き方
import (
	"fmt"
	"time"
)

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	fmt.Println("Hello World")
	// 現在時刻の取得
	fmt.Println(time.Now())

	// 明示的な変数定義
	var i int = 100
	fmt.Println(i)
	var s string = "Hello Go"
	fmt.Println(s)

	// 一度に複数定義する方法
	var t, f bool = true, false
	fmt.Println(t, f)

	// 異なる型を一度に複数定義する方法
	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	// 定義だけする場合
	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	// 暗黙的な変数定義
	i4 := 400
	fmt.Println(i4)
	/*
		暗黙的な変数定義は関数内でのみ宣言可能
		関数外で宣言するとエラーとなる
	*/

	// 関数の呼び出し
	outer()
}
