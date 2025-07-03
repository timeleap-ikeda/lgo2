package main

import (
	"fmt"
	"sync"
)

func ProcessData() {
	ch := make(chan int)
	// 2つのウェイトグループを使う
	// 最初のウェイトグループがチャネルのクローズをコントロールする
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i*100 + 1
		}
	}()
	// 書き込みをする2つのゴルーチンが終了したときにチャネルをクローズするヘルパーゴルーチンを起動
	go func() {
		wg.Wait()
		close(ch)
	}()
	// 2番目のウェイトグループは、読み込み側のゴルーチンが終了したときに知らせる
	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()
	wg2.Wait()
}

func main() {
	ProcessData()
}
