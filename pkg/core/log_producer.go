package core

import "fmt"

type LogProducer struct {
	DefaultProducer
}

func (producer *LogProducer) Process(exchange *Exchange) error {
	fmt.Println(exchange.In.Body)
	return nil;
}