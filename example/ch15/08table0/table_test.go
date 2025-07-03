package table

import "testing"

func TestDoMath(t *testing.T) { //liststart1
	result, err := DoMath(2, 2, "+")
	if result != 4 {
		t.Error("期待する結果 4, 結果：", result)
	}
	if err != nil {
		t.Error("エラーはnilのはずだが結果は：", err)
	}
	result2, err2 := DoMath(2, 2, "-")
	if result2 != 0 {
		t.Error("期待する結果 0, 結果：", result2)
	}
	if err2 != nil {
		t.Error("エラーはnilのはずだが結果は：", err2)
	}
	result3, err3 := DoMath(2, 2, "*")
	if result3 != 4 {
		t.Error("期待する結果 4, 結果：", result3)
	}
	if err3 != nil {
		t.Error("エラーはnilのはずだが結果は：", err3)
	}
	result4, err4 := DoMath(2, 2, "/")
	if result4 != 1 {
		t.Error("期待する結果 1, 結果：", result4)
	}
	if err4 != nil {
		t.Error("エラーはnilのはずだが結果は：", err4)
	}
	// ... 
} //listend1
