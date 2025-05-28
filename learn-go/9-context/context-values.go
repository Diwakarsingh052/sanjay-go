package main

import (
	"context"
	"fmt"
)

type ctxKey string

const k ctxKey = "key"

func main() {

	ctx := context.Background()
	//The provided key must be comparable and should not be of type string
	//or any other built-in type to avoid collisions between packages using context
	ctx = context.WithValue(ctx, k, "123")
	fmt.Println(getReqId(ctx))
}

func getReqId(ctx context.Context) int {
	//type assertion // checking if interface is of specified type
	// ok will be false if interface is of different type
	reqId, ok := ctx.Value(k).(int)
	//id, ok := reqId.(int)
	if !ok {
		return 0
	}
	//if reqId == nil {
	//	return ""
	//}
	return reqId

}
