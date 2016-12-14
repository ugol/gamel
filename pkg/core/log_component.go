package core

type LogComponent struct {
	DefaultComponent

}

func (component LogComponent) NewEndpoint(uri string) (Endpoint, error) {
	return &LogEndpoint{
		DefaultEndpoint{
			component: component,
			uri: uri,
		},
	}, nil
}
