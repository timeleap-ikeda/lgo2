package main

import (
	"fmt"
	"math"
)

func main() { //liststart5
	fmt.Println("=== 2Dバージョン ===")
	pair2Da := Pair[Point2D]{Point2D{1, 1}, Point2D{5, 5}}
	fmt.Println("1つ目のペア:", pair2Da)  // {[1,1] [5,5]}
	pair2Db := Pair[Point2D]{Point2D{10, 10}, Point2D{15, 5}}
	fmt.Println("2つ目のペア:", pair2Db)  // {[10,10] [15,5]}
	closer := FindCloser(pair2Da, pair2Db)
	fmt.Println("距離が近いペア:", closer)  // {[1,1] [5,5]}

	fmt.Println("=== 3Dバージョン ===")
	pair3Da := Pair[Point3D]{Point3D{1, 1, 10}, Point3D{5, 5, 0}}
	fmt.Println("1つ目のペア:", pair3Da)  // {[1,1,10] [5,5,0]}
	pair3Db := Pair[Point3D]{Point3D{10, 10, 10}, Point3D{11, 5, 0}}
	fmt.Println("2つ目のペア:", pair3Db)  // {[10,10,10] [11,5,0]}
	closer2 := FindCloser(pair3Da, pair3Db)
	fmt.Println("距離が近いペア:", closer2) // {[10,10,10] [11,5,0]}

	// 同じ型のペアが指定されていない場合
	// closer3 := FindCloser(pair2Da, pair3Da) // コンパイル時のエラー
	// fmt.Println(closer3)

	// Differを実装していない場合  コンパイル時のエラー
	// closer4 := FindCloser(Pair[StringerString]{"a", "b"},
	//             Pair[StringerString]{"hello", "goodbye"})
	// fmt.Println(closer4)
} //listend5

// 同じ型Tの任意の2つ値を保持する型（構造体）Pairの定義 //liststart1
type Pair[T fmt.Stringer] struct { // Tはfmt.Stringerを実装する
	// Tはfmt.Stringerをメソッドとしてもつ構造体
	// このようにインタフェースを「型」として使える
	Val1 T
	Val2 T
} //listend1

// 型パラメータをもつインタフェース
// インタフェースDifferはfmt.Stringerを埋め込み、メソッドDiffをもつ //liststart2
type Differ[T any] interface {
	fmt.Stringer  // fmt.Stringerを埋め込んでいる→メソッドString()をもつ
	Diff(T) float64
} //listend2
// fmt.Stringerについては「7.5 インタフェースを参照」

// 2組のペアを受け取り、距離が近いほうのペアを返す  //liststart3
// インタフェースDifferを満たす型Tを型パラメータとして指定
func FindCloser[T Differ[T]](pair1, pair2 Pair[T]) Pair[T] {
	d1 := pair1.Val1.Diff(pair1.Val2) // pair1の距離
	d2 := pair2.Val1.Diff(pair2.Val2) // pair2の距離
	if d1 < d2 { // pair1のほうがpair2よりも距離が近い
		return pair1
	}
	return pair2 // pair2のほうがpair1よりも近いか同じ距離
} //listend3


// 型Point2Dの定義
type Point2D struct { // 2次元の点  //liststart4
	X, Y int  // X座標、Y座標
}

// Point2Dに付随するメソッドの定義
// 1. メソッドString -- 自分の（2次元の）「座標」を出力
func (p2 Point2D) String() string {
	return fmt.Sprintf("[%d,%d]", p2.X, p2.Y)
}

// 2. メソッドDiff -- 引数で渡された2次元の点と自分との距離を計算
func (p2 Point2D) Diff(from Point2D) float64 { 
	x := p2.X - from.X
	y := p2.Y - from.Y
	return math.Sqrt(float64(x*x) + float64(y*y)) // 2次元距離
}


// 型Point3Dの定義
type Point3D struct {  // 3次元の点
	X, Y, Z int  // X座標、Y座標、Z座標
}

// Point3Dに付随するメソッドの定義
// 1. メソッドString -- 自分の（3次元の）「座標」を出力
func (p3 Point3D) String() string {
	return fmt.Sprintf("[%d,%d,%d]", p3.X, p3.Y, p3.Z)
}

// 2. メソッドDiff -- 引数で渡された3次元の点と自分との距離を計算
func (p3 Point3D) Diff(from Point3D) float64 {
	x := p3.X - from.X
	y := p3.Y - from.Y
	z := p3.Z - from.Z
	return math.Sqrt(float64(x*x) + float64(y*y) + float64(z*z)) // 3次元距離
} //listend4

type StringerString string  // メソッドStringをもつstring

// StringerStringに付随するメソッドの定義
// インタフェースDifferは満たされていない
// 1.  メソッドString
func (ss StringerString) String() string {
	return string(ss)
}
// 2. メソッドDiff はない！
