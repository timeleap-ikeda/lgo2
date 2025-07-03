package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

func main() {
//liststart1		
	const streamData = `
		{"name": "フレッド", "age": 40}
		{"name": "メアリ", "age": 21}
		{"name": "パット", "age": 30}
	`  //listend1
	var t struct {  //liststart2
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	dec := json.NewDecoder(strings.NewReader(streamData))
	var b bytes.Buffer
	enc := json.NewEncoder(&b)  // bにデコードした結果を入れる
	for {
		err := dec.Decode(&t)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}  //listend2
		fmt.Println(t)
		err = enc.Encode(t)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("------")		
	out := b.String()
	fmt.Println(out)
}
