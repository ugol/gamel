package core

import (
	"fmt"
	"errors"
)

type GamelContext struct {
	Name	string
	Status	ContextStatus
}

type ContextStatus int

const (
	Idle ContextStatus = iota
	Started
)

func (context *GamelContext) Start() {
	context.Status = Started
}

func (context GamelContext) GetComponent(name string) (Component, error){

	switch name {
	case "timer":
		return &TimerComponent{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("cannot find component %s", name))
	}
}