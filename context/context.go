package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx := context.Background()

	ctx, _ = context.WithTimeout(ctx, time.Second)
	//	ctx, cansel := context.WithTimeout(ctx, time.Second)
	ctx = context.WithValue(ctx, "id", 1)

	/*go func() {
		time.Sleep(time.Millisecond * 100)
		cansel()
	}()*/
	parset(ctx)
}

func parset(ctx context.Context) {
	id := ctx.Value("id")
	fmt.Println(id)

	for {
		select {
		case <-time.After(time.Second * 2): // запись в канал происходит через 2 сек
			fmt.Println("parsing complete")
			return
		case <-ctx.Done():
			fmt.Println("deadline exceded")
			return
		}
	}
}
