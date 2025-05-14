package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func main() {
	// context -> cancellation signals
	//     		-> passing values in different layers of a http request

	// we need a container to hold the context object

	//Background returns a non-nil, empty Context.
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*1)
	defer cancel() // clean up the resources taken up by the context
	// if we don't put cancel in defer timer would be immediately cancelled
	//req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.google.com", nil)
	//_, err := http.DefaultClient.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	i, err := Slow(ctx, "10")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)

	fmt.Println("ok")

}

// context should be the first argument passed to function

func Slow(ctx context.Context, input string) (int, error) {

	//sql.DB{}.ExecContext()

	i, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	time.Sleep(time.Second * 2)
	select {
	default:
		fmt.Println("moving on, timer is not over yet")
	case <-ctx.Done():
		//undo all the changes done above, if needed
		return 0, ctx.Err()
	}
	return i, nil

}
