package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	f := struct { //liststart
		Name string // NameのNは大文字！ 小文字だと他パッケージから見えない
		Age int
	}{}

	rec := []byte(`{"name": "小野小町", "occupation": "歌人", "age": 20}`)
	fmt.Printf("%v\n", rec) // [123 34 110 ....] // byteのスライスとして表示
	fmt.Printf("%s\n", rec) // {"name": "小野小町", ...}  // 文字列として表示
	err := json.Unmarshal(rec, &f)
	// 大文字小文字の違いを無視して、フィールドに対応付けてくれる
	if (err != nil) {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", f) // {Name:小野小町 Age:20}
	// %+v でフィールド名付きで出力   //listend
}
