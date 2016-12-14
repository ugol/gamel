package core

type Consumer interface {
	Service
	Endpoint()	Endpoint
	Processor()	Processor
}


type DefaultConsumer struct {
	endpoint	Endpoint
	processor	Processor
}

func (consumer DefaultConsumer) Endpoint() Endpoint {
	return consumer.endpoint
}

func (consumer DefaultConsumer) Processor() Processor{
	return consumer.processor
}

func (consumer *DefaultConsumer) Start() error {
	return nil
}

func (consumer *DefaultConsumer) Stop() error {
	return nil
}