package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total             int
	lastUpdated time.Time
}

func (c *Counter) Increment() { // cのアドレスが渡される
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string { // cのコピーが渡される
	return fmt.Sprintf("更新回数: %d, 更新時刻: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) { // 間違い  //liststart1
	c.Increment()  // mainのcのコピーに対してIncrementが行われる
	fmt.Println("Wrong:", c.String())
}

func doUpdateRight(c *Counter) { // 正しい
	c.Increment()  // mainのcに対してIncrementが行われる
	fmt.Println("Right:", c.String())
}

func main() {
	var c Counter
	fmt.Println("main1:", c.String())	
	doUpdateWrong(c)
	fmt.Println("main2:", c.String())
	doUpdateRight(&c)
	fmt.Println("main3:", c.String())
	doUpdateRight(&c)
	fmt.Println("main4:", c.String())	
}
