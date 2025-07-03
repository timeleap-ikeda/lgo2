// 3章のファイル名のの付いていないコードを試すためのものです。
// エラーになる例は、コメントになっています。

package main

import "fmt"
import "slices"
import "maps"	

func main() {
	fmt.Println("===== 3.1 配列 =====")
	{
		var x [3]int
		fmt.Println(x)
	}
	{
		var x = [3]int{10, 20, 30}
		fmt.Println(x)
	}

	{
		var x = [12]int{1, 5: 4, 6, 10: 100, 15}
		fmt.Println(x)
	}

	fmt.Println("----- 配列の比較 -----")
	{
		var x = [...]int{1, 2, 3}
		var y = [3]int{1, 2, 3}
		fmt.Println(x == y) // true
	}

	fmt.Println("----- 多次元配列のシミュレート -----")
	{
		var x [2][3]int
		fmt.Println(x)
	}

	fmt.Println("----- インデックス（添字）の指定 -----")
	{
		var x [3]int
		x[0] = 10
		fmt.Println("x[0]:", x[0])
		fmt.Println("x[2]:", x[2])

		var y [2][3]int
		y[0][2] = 12
		y[1][1] = 3
	}

	fmt.Println("----- len -----")
	{
		var x [3]int
		fmt.Println("len(x):", len(x))

		var y [2][3]int
		fmt.Println("len(y):", len(y))
	}

	fmt.Println("===== 3.2 スライス =====")
	{
		var x = []int{10, 20, 30}
		fmt.Println(x)
	}

	{
		var x = []int{1, 5: 4, 6, 10: 100, 15}
		fmt.Println(x)
	}

	{
		var x = []int{1, 5: 4, 6, 10: 100, 15}
		x[0] = 10
		fmt.Println("x[0]:", x[0])
		fmt.Println("x[2]:", x[2])
	}

	{
		var x []int
		fmt.Println(x)
	}

	fmt.Println("----- スライスの比較 -----")
	{
		var x []int
		var y []int
		// fmt.Println(x==y) // エラー
		fmt.Println("x == nil:", x == nil) // true
		fmt.Println("y != nil:", y != nil) // false
	}
		
	fmt.Println("===== 3.2.1 len =====")
	{
		var x = []int{1, 5: 4, 6, 10: 100, 15}
		fmt.Println("len(x):", len(x))
		var y = []int{}
		fmt.Println("len(y):", len(y)) // 0
	}

	fmt.Println("----- スライスのスライス -----")
	{
		var x [][]int
		fmt.Println(x)

		var y = [][]int{{0, 1}, {2, 3}, {4, 5}}
		fmt.Println(y)
		fmt.Println(y[1])    // [2 3]
		fmt.Println(y[1][1]) // 3
	}

	fmt.Println("----- slices.Equal -----")
	// 「	import "slices"」が必要
	{
		x := []int{1, 2, 3, 4, 5}
		y := []int{1, 2, 3, 4, 5}
		z := []int{1, 2, 3, 4, 5, 6}
		s := []string{"a", "b", "c"}
		fmt.Println(slices.Equal(x, y)) // true
		fmt.Println(slices.Equal(x, z)) // false
		// fmt.Println(slices.Equal(x, s)) // コンパイルエラー
		fmt.Println(s) // 3
	}
	
	fmt.Println("----- lenの引数 -----")
	{
		var x [][]int
		fmt.Println("len(x):", len(x))
		var y = []int{2, 3, 4, 5}
		fmt.Println("len(y):", len(y))
		var z string = "abc"
		fmt.Println("len(z):", len(z))

		var a = 12
		fmt.Println("a:", a)
		// fmt.Println(len(a)) // エラー
	}


	
	fmt.Println("===== 3.2.2 append =====")

	{
		var x []int
		fmt.Println(x)
		x = append(x, 10)
		fmt.Println("append(x, 10):", x)
	}

	{
		var x = []int{1, 2, 3}
		x = append(x, 4)
		fmt.Println("append(x, 4):", x)
		x = append(x, 5, 6, 7)
		fmt.Println("append(x, 5, 6, 7):", x)

		y := []int{20, 30, 40}
		x = append(x, y...)
		fmt.Println("x:", x)
		fmt.Println("append(x, y...):", append(x, y...))

		// append(x, y) // エラー
	}

	fmt.Println("===== 3.2.4 make =====")

	{
		x := make([]int, 5)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x))
		fmt.Println("x[0], x[4]:", x[0], x[4])
	}
	fmt.Println()
	{
		x := make([]int, 5)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x)) // [0 0 0 0 0] 5 5
		x = append(x, 10)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x)) // [0 0 0 0 0 10] 6 10
	}

	fmt.Println("----- キャパシティ（第3引数）を指定 -----")
	{
		x := make([]int, 5, 10)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x))
	}

	{
		x := make([]int, 0, 10)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x)) // [] 0 10
	}

	{
		x := make([]int, 0, 10)
		x = append(x, 5, 6, 7, 8)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x))
	}

	fmt.Println("===== 3.2.5 スライスのクリア =====")
	{
		s := []string{"first", "second", "third"}
		fmt.Println(s, len(s)) // [first second third] 3
		clear(s)
		fmt.Println(s, len(s)) // [  ] 3
		//		fmt.Printf("s[0]=|%s", s[0], "| s[1]=|", s[1], "| s[2]=|", s[2])
		fmt.Printf("s[0]=|%s|, s[1]=|%s|, s[2]=|%s|\n", s[0], s[1], s[2])
	}

	fmt.Println("===== 3.2.6 スライスの生成方法の選択 =====")
	{
		var data []int
		fmt.Println(data, len(data), cap(data))
		fmt.Println("data == nil:", data == nil)

		var x = []int{}
		fmt.Println(x, len(x), cap(x))
		fmt.Println("x == nil:", x == nil)
	}

	{
		data := []int{2, 4, 6, 8}
		fmt.Println(data, len(data), cap(data))
	}


	fmt.Println("===== 3.2.9 配列からスライスへの変換 =====")
	{
		xArray := [4]int{5, 6, 7, 8}
		xSlice := xArray[:]
		fmt.Println("xArray:", xArray) // xArray: [5 6 7 8]
		fmt.Println("xSlice:", xSlice) // xSlice: [5 6 7 8]

		x := [4]int{5, 6, 7, 8}
		y := x[:2]
		z := x[2:]
		fmt.Println("x:", x) // x: [5 6 7 8]
		fmt.Println("y:", y) // y: [5 6]
		fmt.Println("z:", z) // z: [7 8]
	}

	fmt.Println("===== 3.2.10 スライスから配列への変換 =====")
	{
		xArray := [4]int{5, 6, 7, 8}
		xSlice := xArray[:]
		// panicArray := [5]int(xSlice) // ←パニックになる
		// fmt.Println(panicArray)
		fmt.Println("xSlice:", xSlice)  // xSliceを使ってエラーを回避
	}


	{
		var a rune = 'x'
		var s string = string(a)
		var b byte = 'y'
		var s2 string = string(b)
		fmt.Println("a, s, b, s2:", a, s, b, s2)
	}

	{
		var x int = 65
		var y = string(x)  // go vetで警告
		fmt.Println("y:", y) // A （65ではない）
	}

	fmt.Println("===== 3.4 マップ =====")
	{
		var nilMap map[string]int
		fmt.Println("nilMap == nil:", nilMap == nil)   // true
		fmt.Println("len(nilMap):", len(nilMap))       // 0
		fmt.Println("nilMap[\"abc\"]:", nilMap["abc"]) // 0
		// nilMap["abc"] = 3 // ←パニックになる

		totalWins := map[string]int{} // マップリテラル。string→intのマップを要素なしで初期化。nilではない
		fmt.Println("totalWins == nil:", totalWins == nil) // false

		fmt.Println("totalWins[\"abc\"]:", totalWins["abc"]) //		 -> 0
		totalWins["abc"] = 3                                 // 書き込みもできる
		fmt.Println("totalWins[\"abc\"]:", totalWins["abc"]) //
	}
	{
		totalWins := map[string]int{
			"セネターズ":  14,
			"スパローズ":  15,
			"ファルコンズ": 22,
		}
		fmt.Println(totalWins)
	}
	{
		teams := map[string][]string{
			"ライターズ": []string{"夏目", "森", "国木田"},
			"ナイツ": []string{"武田", "徳川", "明智"},
			"ミュージシャンズ": []string{"ラベル", "ベートーベン", "リスト"},
		}

		fmt.Println(teams)    // map[ナイツ:[武田 徳川 明智] ミュージシャンズ:[ラベル ベートーベン リスト] ライターズ:[夏目 森 国木田]]
		fmt.Println(teams["ライターズ"]) // [夏目 森 国木田]
		fmt.Println(teams["ナイツ"])  // [武田 徳川 明智]

		/*
			teams := map[string][]string {
				"Orcas": []string{"Fred", "Ralph", "Bijou"},
				"Lions": []string{"Sarah", "Peter", "Billie"},
				"Kittens": []string{"Waldo", "Raul", "Ze"},
			}

			fmt.Println(teams) // map[Kittens:[Waldo Raul Ze] Lions:[Sarah Peter Billie] Orcas:[Fred Ralph Bijou]]
			fmt.Println(teams["Lions"]) // [Sarah Peter Billie]
			fmt.Println(teams["Kittens"]) // [Waldo Raul Ze]
		*/

		teams2 := map[string][]string{
			"シャチチーム":  []string{"謙信", "信長", "家康"},
			"ライオンチーム": []string{"レオ", "たか子", "カナ"},
			"猫チーム":    []string{"AKB", "MNB", "FNB"},
		}

		fmt.Println(teams2)           // map[シャチチーム:[謙信 信長 家康] ライオンチーム:[レオ たか子 カナ] 猫チーム:[AKB MNB FNB]]
		fmt.Println(teams2["シャチチーム"]) // [謙信 信長 家康]
		fmt.Println(teams2["チャチチーム"]) // []
		fmt.Println(teams2["猫チーム"])   // [AKB MNB FNB]

		fmt.Println("len(teams2[\"猫チーム\"]):", len(teams2["猫チーム"])) // 3
	}

	{
		ages := make(map[int][]string, 10)
		fmt.Println("ages:", ages)           // map[]
		fmt.Println("len(ages):", len(ages)) // 0
	}

	fmt.Println("===== 3.4.2 カンマOKイディオム =====")
	{
		m := map[string]int{
			"hello": 5,
			"world": 0,
		}
		v, ok := m["hello"]
		fmt.Println(v, ok) // 5 true

		v, ok = m["world"]
		fmt.Println(v, ok) // 0 true

		v, ok = m["goodbye"]
		fmt.Println(v, ok) // 0 false

	}

		fmt.Println("===== 3.4.3 マップからの削除 =====")
	{
		m := map[string]int{
			"hello": 5,
			"world": 10,
		}
		fmt.Println(m)
		v, ok := m["hello"]
		fmt.Println(v, ok)
		delete(m, "hello")
		fmt.Println(m)
		v, ok = m["hello"]
		fmt.Println(v, ok)
		delete(m, "こんにちは")
		fmt.Println(m)
	}

	fmt.Println("===== 3.4.4 マップを空にする =====")
	{
		m := map[string]int{
			"hello": 5,
			"world": 10,
		}
		fmt.Println(m, len(m)) // map[hello:5 world:10] 2
		clear(m)
		fmt.Println(m, len(m)) // map[] 0
	}

	fmt.Println("===== 3.4.5 マップの比較 =====")
	// 「import "maps"」が必要
	{
		m := map[string]int{
			"hello": 5,
			"world": 10,
		}
		n := map[string]int{
			"world": 10,
			"hello": 5,
		}
		fmt.Println(maps.Equal(m, n)) // true
	}

	
	fmt.Println("===== 3.5 構造体 =====")
	{
		type person struct {
			name string // 名前
			age  int // 年齢
			pet  string // ペット
		}

		var fred person
		fmt.Println("fred:", fred) // fred: { 0 } // 空文字列2つあり

		bob := person{} // 構造体リテラル。全フィールドがゼロ値で初期化される

		fmt.Println("bob:", bob) // bob: { 0 } // 空文字列2つあり
		bob.name = "ボブ"
		fmt.Println("bob:", bob) // bob: {ボブ 0 }

		julia := person{
			"Julia", // name
			40,      // age
			"cat",   // pet
		}
		fmt.Println("julia:", julia) // julia: {Julia 40 cat}

		/*	  コンパイル時エラー（全フィールドを指定する必要がある）
		fred := person{
			"フレッド",  // name
			"cat",    // pet
		}
		*/

		beth := person{
			age:  30,
			name: "ベス",
		}
		fmt.Println("beth:", beth) // beth: {ベス 30 }

		fmt.Println("bob.name:", bob.name) // bob.name: ボブ
		fmt.Println("bob:", bob)           // bob: {ボブ 0 }

	}

	fmt.Println("===== 3.5.1 無名構造体 =====")
	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "bob"
	person.age = 50
	person.pet = "dog"
	fmt.Println("person:", person) // person: {bob 50 dog}

	pet := struct {
		name string
		kind string
	}{
		name: "ポチ",
		kind: "dog",
	}
	fmt.Println("pet:", pet) // pet: {ポチ dog}

	fmt.Println("===== 3.5.2 構造体の比較と変換 =====")

	{
		type firstPerson struct {
			name string
			age  int
		}

		type secondPerson struct {
			name string
			age  int
		}

		type thirdPerson struct {
			age  int
			name string
		}

		x1 := firstPerson{
			"太郎",
			24,
		}
		//		x1.name = "太郎"
		//		x1.age = 24
		fmt.Println("x1:", x1)

		var x2 secondPerson
		x2 = secondPerson(x1)
		fmt.Println("x2:", x2)
		//  fmt.Println("x1==x2:", x1==x2) // コンパイル時エラー
		x12 := firstPerson(x2)
		fmt.Println("x1==x12:", x1 == x12) // x1==x12: true

		var x3 thirdPerson
		// x3 = thirdPerson(x1)  //コンパイル時エラー
		fmt.Println("x3:", x3)

		type fourthPerson struct {
			firstName string
			age       int
		}
		var x4 fourthPerson
		// x4 = fourthPerson(x1)  // コンパイル時エラー
		fmt.Println("x4:", x4)

		type fifthPerson struct {
			name          string
			age           int
			favoriteColor string // 好みの色
		}
		var x5 fifthPerson
		// x5 = fifthPerson(x1)  // コンパイル時エラー
		fmt.Println("x5:", x5)
	}

	fmt.Println("\n----- 無名構造体との変換 -----")
	{
		type firstPerson struct {
			name string
			age  int
		}
		f := firstPerson{
			name: "Bob",
			age:  50,
		}
		var g struct {
			name string
			age  int
		}
		fmt.Println("f:", f)
		fmt.Println("g:", g)
		g = f
		fmt.Println("g（fの代入後）:", g)
		fmt.Println("f==g:", f == g)
	}

} // main
