package format  // ←ディレクトリ名はdo-formatだがパッケージ名はformat

import "fmt"

func Number(num int) string {
	return fmt.Sprintf("The number is %d", num)
}
