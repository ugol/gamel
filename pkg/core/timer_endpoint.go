package core

type TimerEndpoint struct {
	DefaultEndpoint
}

func (endpoint TimerEndpoint) NewConsumer() (Consumer, error) {
	return &TimerConsumer{
		DefaultConsumer{endpoint: endpoint},
	}, nil
}