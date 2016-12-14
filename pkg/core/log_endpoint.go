package core

type LogEndpoint struct {
	DefaultEndpoint
}

func (endpoint LogEndpoint) NewProducer() (Producer, error) {
	return &LogProducer{
		DefaultProducer{endpoint: endpoint},
	}, nil
}