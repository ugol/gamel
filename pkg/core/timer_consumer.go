package core

import (
	"time"
	"strconv"
)

type TimerConsumer struct {
	DefaultConsumer

	quit	chan struct{}
}

func (consumer *TimerConsumer) Start() error {
	period, _ := strconv.Atoi(consumer.endpoint.Parameter("period"))
	ticker := time.NewTicker(time.Duration(period) * time.Millisecond)
	consumer.quit = make(chan struct{})
	counter:=1
	go func() {
		for {
			select {
			case <- ticker.C:

				message := NewMessage()
				message.Body = counter
				counter+=1

				exchange := NewExchange()
				exchange.In = message

				// async again
				go consumer.processor.Process(exchange)

			case <- consumer.quit:
				ticker.Stop()
				return
			}
		}
	}()

	return nil
}

func (consumer *TimerConsumer) Stop() error {
	close(consumer.quit)
	consumer.quit = nil

	return nil
}