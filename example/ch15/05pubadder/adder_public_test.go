package pubadder_test

import (
	"github.com/learning-go-book-2e/ch15/sample_code/pubadder"
	"testing"
)

func TestAddNumbers(t *testing.T) {
	result := pubadder.AddNumbers(2, 3)
	if result != 5 {
		t.Error("結果が正しくない: 想定される結果 5, 得られた結果", result)		
		//		t.Error("incorrect result: expected 5, got", result)
	}
}
