package main

import (
	"fmt"
)

type Score int // intから新たな型（ユーザー定義型）Scoreを定義 //liststart1
type HighScore Score // ユーザー定義型Scoreから別のユーザー定義型HighScoreを定義

type Person struct { // ユーザー定義型Person（人）を定義
	LastName string  // 姓
	FirstName string // 名
	Age int  // 年齢
} 
type Employee Person // Personを使ってEmployee（従業員）を定義 //listend1


func (s Score) Double() Score {  //liststart3
	return s * 2
}  //listend3

func main() {
	// 型のない定数の代入は認められる //liststart2
	var i int = 300
	var s Score = 100  // Scoreの基底型はintなので100は代入できる
	var hs HighScore = 200
	// hs = s    // コンパイル時のエラー！ 型が違う（ScoreとHighScore）
	// s = i     // コンパイル時のエラー！ 型が違う（Scoreとint）
	s = Score(i)   // 型変換後に代入
	hs = HighScore(s)  // 型変換後に代入
	fmt.Println(s, hs) // 300 300
	hhs := hs + 20  // 基底型（int）に対して使える演算子（+）は使える
	fmt.Println(hhs)  // 320

	var s2 Score = 50
	scoreWithBonus := s2 + 100 // scoreWithBonusの型はScore  //listend2
	fmt.Println("scoreWithBonus:", scoreWithBonus)  // 150

	s = 200   //liststart4
	hs = 300
	fmt.Println(s.Double()) // 400
	fmt.Println(Score(hs).Double()) // 600
	// fmt.Println(hs.Double()) // コンパイル時のエラー //listend4
} 
