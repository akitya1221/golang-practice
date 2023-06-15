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

	// 複数行の文字列を表示したい場合
	fmt.Println(`test
	test
	test`)

	// 文字の切り出し
	// ※string()をつけない場合、バイト型で表示される
	fmt.Println(string(s[0]))

	// byte型
	// スライス（配列）
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	// string型をbyte型に変換
	c := []byte("HI")
	fmt.Println(c)

	// 配列型
	var arr1 [3]int
	fmt.Println(arr1)
	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)
	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)
	arr4 := [...]string{"c", "d"}
	fmt.Println(arr4)

	arr2[2] = "C"
	fmt.Println(arr2)

	// 要素数の取得
	fmt.Println(len(arr4))

	// interface型
	// 特殊な型でいろんな型と互換性があるのが特徴
	var x interface{}
	x = 1
	x = 0.1
	x = "test"
	x = [3]int{1, 2, 3}
	fmt.Println(x)
}
