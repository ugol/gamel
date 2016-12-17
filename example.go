package main

import (
	"github.com/ugol/gamel/pkg/core"
	"time"
	"fmt"
)

func main() {

	context := core.NewGamelContext()

	route, _ := core.NewRouteBuilder(context).
		From("timer:tick?period=1000").
		To("log:INFO").
		Build()

	context.AddRoute(route)
	context.Start()

	time.Sleep(20 * time.Second)
	context.Stop()

	fmt.Println("Bye :)")
}
