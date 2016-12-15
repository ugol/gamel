package core

type Producer interface {
	Processor
	Endpoint() Endpoint
}

type DefaultProducer struct {
	endpoint Endpoint
}

func (producer DefaultProducer) Endpoint() Endpoint {
	return producer.endpoint
}

func (producer *DefaultProducer) Start() error {
	return nil
}

func (producer *DefaultProducer) Stop() error {
	return nil
}

func (producer *DefaultProducer) Suspend() error {
	return nil
}

func (producer *DefaultProducer) Resume() error {
	return nil
}
