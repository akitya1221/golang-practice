package main

// 複数のパッケージをimportする場合の書き方
import (
	"fmt"
	"time"
)

// 定数
const (
	// iotaは連続した整数の連番を生成してくれる
	c0 = iota
	c1
	c2
)

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

// 関数
// 戻り値の型は()の右側に記載する
func Sum(x int, y int) int {
	return x + y
}

// 複数の戻り値がある場合
func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

// 関数を返す関数
func ReturnFunc() func() {
	return func() {
		fmt.Println("return function")
	}
}

// 引数に関数を使用する関数
func CallFunc(f func()) {
	f()
}

// クロージャー、ジェネレーター
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
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

	// 定数
	fmt.Println(c0, c1, c2)

	// 複数の戻り値がある場合
	t_int1, t_int2 := Div(2, 3)
	fmt.Println(t_int1, t_int2)
	// 戻り値の破棄
	// 「_」を使用することで、戻り値を破棄することができる
	t_int3, _ := Div(2, 3)
	fmt.Println(t_int3)

	// 無名関数
	test_func1 := func(x1, y1 int) int {
		return x1 + y1
	}
	func_result := test_func1(1, 2)
	fmt.Println(func_result)
	// 値を渡してしまうことも可能
	test_func2 := func(x1, y1 int) int {
		return x1 + y1
	}(1, 2)
	fmt.Println(test_func2)

	// 関数を返す関数
	return_func := ReturnFunc()
	fmt.Println(return_func)

	// 引数に関数を使用する関数
	CallFunc(return_func)

	// クロージャー、ジェネレーター
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	// for文
	j := 0
	for {
		j++
		if j == 5 {
			break
		}
		fmt.Println("Loop")
	}

	point := 0
	for point < 5 {
		fmt.Println(point)
		point++
	}

	for j := 0; j < 5; j++ {
		if j == 3 {
			continue
		}
		fmt.Println(j)
	}

	arr := [5]int{1, 2, 3, 4, 5}
	for j := 0; j < len(arr); j++ {
		fmt.Println(arr[j])
	}

	// インデックス番号を表示することができる
	for j, v := range arr {
		fmt.Println(j, v)
	}

	// インデックス番号を表示しない方法
	for _, v := range arr {
		fmt.Println(v)
	}

	// インデックス番号を表示する方法
	for j := range arr {
		fmt.Println(j)
	}

	// switch式
	switch sw := 1; sw {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("それ以外")
	}
}
