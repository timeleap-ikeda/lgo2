package main

import "fmt"

// 型Personを定義   //liststart1
type Person struct {
	FirstName string // 名
	LastName string  // 姓
	Age int  // 年齢
}  //listend1

// 型Personに付随するメソッドStringを定義（PersonにメソッドStringを付加） //liststart2
// メソッドは「レシーバー」という特別な引数をもつ関数
// レシーバーpにはPerson型のインスタンスが指定されて呼び出される
func (p Person) String() string { //「(p Person)」がレシーバーの指定
	return fmt.Sprintf("%s %s：年齢%d歳", p.LastName, p.FirstName, p.Age)
	// fmt.Sprintfはfmt.Printfの、文字列生成用のバージョン。
	// 第1引数に指定したフォーマットで、第2引数以降に指定した値を文字列に変換する。
	// フォーマットの指定方法は「B.2 fmtパッケージのverb」に解説あり。
} //listend2

func main() { //liststart3
	p := Person { // 上で定義したPerson型の変数pの宣言と初期化
		LastName: "武田",
		FirstName: "信玄",
		Age: 52,
	}
	
	output := p.String() // 型Personに付随するメソッドStringを呼び出す
	fmt.Println(output) // 武田 信玄：年齢52歳
} //listend3
