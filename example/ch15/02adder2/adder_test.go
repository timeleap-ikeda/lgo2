package adder  //liststart1

import (
	"testing"
	"time"
	"os"
	"fmt"
)

var testTime time.Time  // パッケージレベルの変数

func TestMain(m *testing.M) {
	fmt.Println("●テスト前の準備")
	testTime = time.Now()
	exitVal := m.Run() // テスト関数を実行。下の3つを順に呼び出す
	fmt.Println("●テスト後の後片付け")
	os.Exit(exitVal)
}

func TestFirst(t *testing.T) { // m.Runの呼び出しで最初に実行
	fmt.Println(" 1. TestFirstではTestMainで準備されたもの利用する:", testTime)
}

func TestSecond(t *testing.T) { // m.Runの呼び出しで2番目に実行
	fmt.Println(" 2. TestSecondでもTestMainで準備されたもの利用する:", testTime)
}

func Test_addNumbers(t *testing.T) { // m.Runの呼び出しで3番目に実行
	fmt.Println(" 3. Test_addNumbersのテストを実行:", testTime)
	result := addNumbers(2, 3)
	if result != 5 {
		t.Error("結果が違う: 期待する結果 5, 得られた結果", result)
	}
}  //listend1

