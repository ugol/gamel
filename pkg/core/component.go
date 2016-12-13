package core

type Component interface {
	Name()	string
}

type DefaultComponent struct {
	name	string
}

func (component DefaultComponent) Name() string {
	return component.name
}