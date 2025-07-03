package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	total := 0
	count := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("合計:", total, "  反復回数:", count, "    タイムアウト")
			return
		default:
		}
		newNum := rand.Intn(100_000_000)
		if newNum == 1_234 {
			fmt.Println("合計:", total, "  反復回数:", count, "    成功！")
			return
		}
		total += newNum
		count++
	}
}
