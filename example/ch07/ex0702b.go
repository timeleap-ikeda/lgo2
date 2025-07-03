package main  //liststart1

import (
	"fmt"
	"time"
)

type Counter struct { // Counter型を定義
	total       int  // 更新回数
	lastUpdated time.Time  // 更新時刻
}

// 型Counterに付随するメソッドIncrement（1増やす）を定義
func (c *Counter) Increment() { // ポインタレシーバー（cはポインタ）
	c.total++
	c.lastUpdated = time.Now()
}

// 型Counterに付随するメソッドStringを定義
func (c Counter) String() string { // 値レシーバー（cにはコピーが渡される）
	return fmt.Sprintf("更新回数: %d, 更新時刻: %v", c.total, c.lastUpdated)
} //listend1


func main() {
	c := &Counter{}  //liststart2
	fmt.Println(c.String()) // 「(*c).String()」と書かなくてもよい 
	c.Increment()
	fmt.Println(c.String())  //listend2
} 
