package format

import "fmt"

func AddNote(s string, num int) string {
	return fmt.Sprintf("%s: %d", s, num)
}
