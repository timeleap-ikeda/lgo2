// 2章のリスト番号の付いていないコードを試すためのものです。
// エラーになる例は、コメントになっています。


package main
import "fmt"
func main() {

	fmt.Println("===== 2.1 事前宣言された型 =====")		
	fmt.Println("===== 2.1.1 ゼロ値 =====")

	{
		var x int
		var y float32
		var z string

		fmt.Println("x:", x) // x: 0
		fmt.Println("y:", y)  // y: 0
		fmt.Printf("z:|%s|\n", z)	 // z:|| （空文字列）
	}

	fmt.Println("===== 2.1.2 リテラル =====")
	fmt.Println("===== 2.1.2.1 整数リテラル =====")
	fmt.Println(1_234_567)  // 1234567
	fmt.Println(0b11010)  // 26 = 16+8+0+2+0
	fmt.Println(0x9F3)  //  2547 = 16*16*9 + 16*15 + 3
	fmt.Println(6.03e23)  //  6.03e+23
	fmt.Println(6.03e3)  //  6.03*(10*10*10) = 6030
	fmt.Println(6.03e-3)  // 0.00603 = 6.03*((1/10)*(1/10)*(1/10))
	fmt.Println("\141")  // "a" と同じ
	fmt.Println("\x61")  // "a" と同じ
	fmt.Println("\u0061")  // "a" と同じ
	fmt.Println("\U00000061")  // "a" と同じ
	fmt.Println("a")
	fmt.Println()
	
	fmt.Println(0x2p5) // 2^5  （ここで^は冪乗。たとえば、2^5 は 2*2*2*2*2）
	fmt.Println(0x12.34p5) //  582.5  = ( 1*16^1 + 2*16^0 + 3*16^(-1) + 4*16^(-2) ) * 2^5
	// p5は、2のべき乗で「スケーリング」する。ここでは、2の5乗倍する（32倍する）
	// よく意味がわからない方はChatGPTなどに尋ねると教えてくれます


	fmt.Println("===== 2.1.2.2 浮動小数点数リテラル =====")
	fmt.Println(1_234.567_89)  // 1234.56789

	fmt.Println(1_234.567_89)  // 1234.56789
	
	
	fmt.Println("===== 2.1.2.3 runeリテラル =====")
	{
		// A
		fmt.Println("1->\x41")
		fmt.Println("2->\u0041")
		fmt.Println("3->\U00000041")
		// あ
		fmt.Println("1->\u3042")
		fmt.Println("2->\U00003042")
		// 絵文字
		fmt.Println("1->\U0001F600") 

		// 「\」（バックスラッシュ）でエスケープされたruneリテラル
		fmt.Println(" --「\\」（バックスラッシュ）でエスケープされたruneリテラル --");
		fmt.Println('\n')  // 改行
		fmt.Println('\t')  // タブ
		fmt.Println('\'')  // 一重引用符（シングルクオート）
		fmt.Println('\\')  // バックスラッシュ
		
		fmt.Println("===== 2.1.2.4 文字列リテラル =====")
		x :=	`バッククオートを使って"ロー文字列リテラル"を書くことで
改行や二重引用符（ダブルクオート）を文字列中に含めることができる`
		fmt.Println(x)
	}

	fmt.Println("===== 2.1.2.5 リテラルと型 =====")
	{
		var x float32 = 2 * 0.23 * 3 // エラーにならない
		fmt.Println(x)
		var a float64 = 3.14
		var b float64 = 10/a
		fmt.Println(b) // 3.184713375796178  // エラーにならない

		// var c int = "abc"  // コンパイル時のエラー
		// var s string = 12  // コンパイル時のエラー

		var bb byte;
		// bb = 1000; // コンパイル時のエラー。オーバースローする
		fmt.Println("bb:", bb) //
		
		var d int
		// d = 12.3 // コンパイル時のエラー （桁落ちする）
		d = 12.0 // エラーにならない（桁落ちしない）
		fmt.Println("d:", d) // d: 12
	}
	
	
	fmt.Println("===== 2.1.3 論理型 =====")
	{
		var flag bool // 値が代入されないとfalseになる（ゼロ値）
		var isAwesome = true  // 素晴らしい！
		fmt.Println(flag)  // false
		fmt.Println(isAwesome)  // true
	}
	
	fmt.Println("===== オーバーフロー（あふれ） =====")			
	{
	    //	var x byte = 1000 // エラー。たとえば1000をbyte型の変数に代入しようとするとエラーになる
		var x byte = 100
		fmt.Println(x)
	}

	fmt.Println("===== 2.1.4.4 整数関連の演算子 =====")				
	{ 
		var x int = 10
		x *= 2
		fmt.Println(x)  // 20
	}

	fmt.Println("===== 2.1.4.5 浮動小数点数型 =====")					
	{
		var x float64 = 0
		fmt.Println(x)  // 0
		x += 1.333456
		fmt.Println(x)  // 1.333456

		x = 1.5
		var y float64 = 2.5
		x /= y
		fmt.Println(x) // 0.6
		// 		x = x % y    // エラー。 floatにたいして「%」は使えない

		fmt.Println("x/0=", x/0)  // +Inf
		fmt.Println("-x/0=", -x/0)  // -Inf
	}


	fmt.Println("===== 2.1.5 文字列とrune =====")						
	{
		var a, b string;
		a = "イロハ"
		b = "あいうえお"
		fmt.Println(a==b) // false
		fmt.Println(a!=b) // true
		fmt.Println(`"あ" < "い":`, "あ" < "い") // true
		fmt.Println(`"ア" < "あ":`, "ア" < "あ") // false
		fmt.Println("a+b:", a+b)  // イロハあいうえお
	}

	fmt.Println("===== 2.1.7 リテラルと型 =====")						
	{
		var x float64 = 200.3 * 5
		var y float64 = 10
		fmt.Println(`x: `, x) // 1001.5
		fmt.Println(`y: `, y) // 10

		// var s float64 = "abc"   // コンパイル時のエラー
		// var s float64 = "abc"   // コンパイル時のエラー
		// var i1 int = 10.1  // コンパイル時のエラー
		var i2 int = 10.00  // これは大丈夫（切り捨てが行われない）
		fmt.Println(`i2: `, i2) // 10
	}
	
	fmt.Println("===== 2.2 変数の宣言 =====")							
	{
		var x, y int = 10, 20
		fmt.Println(x, y)  // 10 20
	}

	{
		var x, y int
		fmt.Println(x, y)  // 0 0
	}

	{
		var x, y = 10, "hello"
		fmt.Println(x, y) // 10 hello
	}

	{
		var (
			x    int
			y        = 20
			z    int = 30
			d, e     = 40, "hello"
			f, g string
		)
		fmt.Println(x, y, z, d, e) // 0 20 30 40 hello

		fmt.Println("f=|", f, "| g=|", g, "|" ) // f=|  | g=|  |
	}

	{
		var x = 10
		fmt.Println(x) // 10
	}

	{
		x := 10
		fmt.Println(x) // 10
	}


	{
		var x = 10
		// var y int32 = x  // エラー
		var y int = x
		fmt.Println(y) // 10
	}
	{
		x := 10
		var y = 10;
		var z int64 = 10;
		
		fmt.Println(x==y)
		//	fmt.Println(x==z)  // エラー  x と z は型が異なる
		fmt.Println( int64(x)==z ) // true
	}


	fmt.Println("===== 2.3 定数 =====")
	{
		x := 5
		y := 10
		// const z = x + y // ★コンパイルエラー！★
		fmt.Println(x, y)
	}

	fmt.Println("===== 2.4 型付きの定数と型のない定数 =====")
	{
		const x = 10
		var y int = x
		var z float64 = x
		var d byte = x

		fmt.Println("x, y, z, d:", x, y, z, d)

		const typedX int = 10
		fmt.Println("typedX:", typedX)
		fmt.Println(x == typedX) // true
	}
	{
		const typedX int = 10
		// var z float64 = typedX  // コンパイル時のエラー
		// fmt.Println("z:", z)
	}

	fmt.Println("===== 2.5 未使用変数 =====")	
	{
			x := 10
			x = 20
			fmt.Println(x)
			x = 30
	}

}
