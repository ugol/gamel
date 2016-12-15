package core

type Route interface {
	Service
	Services() []Service
}

type DefaultRoute struct {
	services []Service
}

func NewRoute(services ...Service) Route {
	return &DefaultRoute{
		services: services,
	}
}

func (route *DefaultRoute) Start() error {
	// TODO return a composite error
	var globalErr error
	for i := len(route.services) - 1; i >= 0; i-- {
		err := route.services[i].Start()
		if err != nil {
			globalErr = err
		}
	}
	return globalErr
}

func (route *DefaultRoute) Stop() error {
	// TODO return a composite error
	var globalErr error
	for i := 0; i < len(route.services); i++ {
		err := route.services[i].Stop()
		if err != nil {
			globalErr = err
		}
	}
	return globalErr
}

func (route DefaultRoute) Services() []Service {
	return route.services
}
