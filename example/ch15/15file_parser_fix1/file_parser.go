package file_parser

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strconv"
	//	"strings"  // 最初の段階では使わない。fixで使う
)

func ParseData(r io.Reader) ([]string, error) {
	// 最初の行には行数を表す整数
	s := bufio.NewScanner(r)
	if !s.Scan() {
		return nil, errors.New("empty")
	}
	countStr := s.Text()
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return nil, err
	}

	// 1番目のテスト結果への対応 start
	if count < 0 {  //liststart1
		return nil, errors.New("no negative numbers") // 行数の指定は0以上
	}  //listend1
	// 1番目のテスト結果への対応 end

	out := make([]string, 0, count)
	// 続く行をスライスに入れる
	for i := 0; i < count; i++ {
		hasLine := s.Scan()
		if !hasLine {
			return nil, errors.New("too few lines") // 行数が少なすぎる
		}
		line := s.Text()
		out = append(out, line)
	}
	// できあがったスライスを返す
	return out, nil
}

func ToData(s []string) []byte {
	var b bytes.Buffer
	b.WriteString(strconv.Itoa(len(s)))
	b.WriteRune('\n')
	for _, v := range s {
		b.WriteString(v)
		b.WriteRune('\n')
	}
	return b.Bytes()
}
