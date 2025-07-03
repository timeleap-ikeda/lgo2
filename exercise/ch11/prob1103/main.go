package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed udhr/text/english_rights.txt
var englishRights string

//go:embed udhr/text/french_rights.txt
var frenchRights string

//go:embed udhr/text/spanish_rights.txt
var spanishRights string

//go:embed udhr/text/japanese_rights.txt
var japaneseRights string

//go:embed udhr/text/korean_rights.txt
var koreanRights string

//go:embed udhr/text/chinese_rights.txt
var chineseRights string

func main() {
	process(os.Args)
}

func process(args []string) {
	if len(args) == 1 {
		fmt.Println("no language provided")
		return
	}
	switch strings.ToLower(args[1]) {
	case "english":
		fmt.Println(englishRights)
	case "spanish":
		fmt.Println(spanishRights)
	case "french":
		fmt.Println(frenchRights)
	case "japanese":
		fmt.Println(japaneseRights)
	case "korean":
		fmt.Println(koreanRights)
	case "chinese":
		fmt.Println(chineseRights)
	default:
		fmt.Println("unknown language:", os.Args[1])
	}
}
