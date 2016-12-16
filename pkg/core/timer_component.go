package core

type TimerComponent struct {
	DefaultComponent
}

func (component TimerComponent) NewEndpoint(uri string) (Endpoint, error) {
	return &TimerEndpoint{
		DefaultEndpoint{
			component: component,
			uri: uri,
		},
	}, nil
}
