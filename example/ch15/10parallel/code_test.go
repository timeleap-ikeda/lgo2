package parallel

import (
	"fmt"
	"testing"
)

func TestParallelTable(t *testing.T) {  //liststart1
	data := []struct {
		name   string
		input  int
		output int
	}{
		{"a", 10, 20},
		{"b", 30, 40},
		{"c", 50, 60},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			fmt.Println(d.input, d.output)
			out := toTest(d.input)
			if out != d.output {
				t.Error("didn't match", out, d.output)
			}
		})
	}
}  //listend1

