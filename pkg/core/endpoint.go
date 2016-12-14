package core

import "errors"

type Endpoint interface {
	Component()		Component
	Uri()			string
	NewConsumer(Processor)	(Consumer, error)
	NewProducer()		(Producer, error)

}

type DefaultEndpoint struct {
	component	Component
	uri		string
}

func (endpoint DefaultEndpoint) Component() Component {
	return endpoint.component
}

func (endpoint DefaultEndpoint) Uri() string {
	return endpoint.uri
}

func (endpoint DefaultEndpoint) NewConsumer(processor Processor) (Consumer, error) {
	return nil, errors.New("this endpoint doesn't provide a consumer")
}

func (endpoint DefaultEndpoint) NewProducer() (Producer, error) {
	return nil, errors.New("this endpoint doesn't provide a producer")
}