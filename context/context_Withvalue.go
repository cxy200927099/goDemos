/**
 * context取值
 */
package main

import (
	"context"
	"fmt"
)

func process(ctx context.Context){
	ret,ok := ctx.Value("trace_id").(int)
	if !ok {
		ret = 9999999
	}

	fmt.Println("ret:%d",ret)

	s, _ := ctx.Value("session").(string)
	fmt.Println("session:%s",s)
}

func main()  {
	ctx := context.WithValue(context.Background(), "trace_id", 13146878)
	ctx = context.WithValue(ctx, "session", "abcdefghijklmn")
	process(ctx)

}