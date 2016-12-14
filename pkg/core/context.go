package core

import (
	"fmt"
	"errors"
	"net/url"
)

type GamelContext interface {
	Service
	Name()			string
	WithName(string)	GamelContext

	Status()		ContextStatus

	GetComponent(name string) (Component, error)
	GetComponentFromURI(uri string) (Component, error)

	AddRoute(route Route)	GamelContext
}

type ContextStatus int

const (
	Idle ContextStatus = iota
	Started
	Error
	Stopped
)


type DefaultGamelContext struct {
	name	string
	status	ContextStatus
	routes	[]Route
}

func NewGamelContext() GamelContext {
	return &DefaultGamelContext{
		name: "gamel",
		status: Idle,
	}
}

func (context *DefaultGamelContext) Start() error {
	fmt.Println("Starting Gamel Context...")
	for i:=0; i<len(context.routes); i++ {
		err := context.routes[i].Start()
		if err != nil {
			context.status = Error
			return err
		}
	}
	context.status = Started
	fmt.Println("Gamel Context Started")
	return nil
}

func (context *DefaultGamelContext) Stop() error {
	fmt.Println("Stopping Gamel Context...")
	for i:=0; i<len(context.routes); i++ {
		err := context.routes[i].Stop()
		if err != nil {
			context.status = Error
			return err
		}
	}
	context.status = Stopped
	fmt.Println("Gamel Context Stopped")
	return nil
}

func (context DefaultGamelContext) Name() string {
	return context.name
}

func (context *DefaultGamelContext) WithName(name string) GamelContext {
	context.name = name
	return context
}

func (context DefaultGamelContext) Status() ContextStatus {
	return context.status
}

func (context DefaultGamelContext) GetComponent(name string) (Component, error){

	switch name {
	case "timer":
		return &TimerComponent{}, nil
	case "log":
		return &LogComponent{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("cannot find component %s", name))
	}
}

func (context DefaultGamelContext) GetComponentFromURI(uri string) (Component, error){
	url, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	name := url.Scheme
	return context.GetComponent(name)
}

func (context *DefaultGamelContext) AddRoute(route Route) GamelContext {
	context.routes = append(context.routes, route)
	return context
}