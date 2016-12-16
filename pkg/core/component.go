package core

type Component interface {
	Name()			string
	NewEndpoint(uri string)	(Endpoint, error)
}

type DefaultComponent struct {
	name	string
}

func (component DefaultComponent) Name() string {
	return component.name
}

