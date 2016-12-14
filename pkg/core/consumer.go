package core

type Consumer interface {
	Service
	Endpoint()	Endpoint
}


type DefaultConsumer struct {

	endpoint	Endpoint

}

func (consumer DefaultConsumer) Endpoint() Endpoint {
	return consumer.endpoint
}

func (consumer *DefaultConsumer) Start() error {
	return nil
}

func (consumer *DefaultConsumer) Stop() error {
	return nil
}