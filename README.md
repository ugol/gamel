# Gamel

Gamel is an experimental framework to reimplement most of [Apache Camel](http://camel.apache.org) in Go.

```go
	context := core.NewGamelContext()
	route, _ := core.NewRouteBuilder(context).
		From("timer:tick?period=2000").
		To("log:INFO").
		Build()

	context.AddRoute(route)
	context.Start()

	time.Sleep(20 * time.Second)
	context.Stop()
```