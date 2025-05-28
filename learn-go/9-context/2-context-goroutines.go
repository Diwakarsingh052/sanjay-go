package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Call the simulated slow function, which may take time or be cancelled
		x, err := slowFuncV2(ctx)
		if err != nil {
			// If the context was cancelled or timeout occurred, just return
			return
		}
		_ = x

		time.Sleep(3 * time.Second)

		select {
		case ch <- x: // Send the value if receiver is ready
			fmt.Println("finalizing things")
		case <-ctx.Done(): // If context expired while trying to send, just return
			fmt.Println("go routine can't finish in time")
			return
		}
	}()

	// Main goroutine waits either for:
	// 1. Value received from worker goroutine
	// 2. Timeout (context cancellation)
	select {

	// if received value in time, this case evaluates
	case val := <-ch:
		fmt.Println(val)

	case <-ctx.Done():
		fmt.Println("timer finished : receiver goroutine")
		fmt.Println(ctx.Err())
	}

	fmt.Println("main func ended")
	wg.Wait()
}

func slowFuncV2(ctx context.Context) (int, error) {
	time.Sleep(time.Second * 1)

	select {
	default:
	case <-ctx.Done():
		fmt.Println("timer finished : slowfuncV2")
		//undo all the changes done above, if needed
		return 0, ctx.Err()
	}
	fmt.Println("slowFunction: slow fn ran and add 100 records to db")
	fmt.Println()

	return 100, nil
}
