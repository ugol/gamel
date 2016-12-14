package core

type TimerEndpoint struct {
	DefaultEndpoint
}

func (endpoint TimerEndpoint) NewConsumer(processor Processor) (Consumer, error) {
	return &TimerConsumer{
		DefaultConsumer: DefaultConsumer{
			endpoint: endpoint,
			processor: processor,
		},
		quit: nil,
	}, nil
}