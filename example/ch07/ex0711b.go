package main

import (
	"fmt"
	"time"
)

type Counter struct {  //liststart1
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {  // 型*Counterを対象とするメソッド
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("合計: %d, 更新: %v", c.total, c.lastUpdated)
}

type Incrementer interface { // 型IncrementerはメソッドIncrementを実装する必要がある
	Increment()
}

func main() {
	var myStringer fmt.Stringer
	var myIncrementer Incrementer
	pointerCounter := &Counter{}
	valueCounter := Counter{}

	myStringer = pointerCounter    // ok
	myStringer = valueCounter      // ok
	myIncrementer = pointerCounter // ok
	myIncrementer = valueCounter   // コンパイル時のエラー

	fmt.Println(myStringer)
	fmt.Println(myIncrementer)
}  //listend1
