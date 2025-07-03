package main

import (
	"fmt"
	"time"
)

type Counter struct { // 構造体Counterの定義 //liststart
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() { // *Counterをレシーバーとするメソッドの定義
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {  // Counterをレシーバーとするメソッドの定義
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

type Incrementer interface { // インタフェースIncrementerの定義
	Increment()  // 関数IncrementをもてばIncrementerになれる
}

func main() {
	var pointerCounter *Counter  // 構造体Counterを指すポインタ。ゼロ値はnil
	fmt.Println(pointerCounter == nil) // true // 値はnil（ゼロ値）なので
	fmt.Printf("%T\n", pointerCounter);
	var incrementer Incrementer  // インタフェースIncrementerを満たす変数の定義
	                             // ゼロ値はnil
	fmt.Println(incrementer == nil) // true
	fmt.Printf("%T\n", incrementer);
	incrementer = pointerCounter  // pointerCounterは*CounterなのでIncrementの実装をもつので、incrementerに代入できる
	fmt.Println(incrementer == nil) // false  // pointerCounter
	fmt.Printf("%T\n", incrementer);
	fmt.Printf("%v\n", incrementer);
}  //listend
