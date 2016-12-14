package core


type RouteBuilder struct {
	context GamelContext
	input 	string
	output	string
	// TODO only 1 input and output for now. To be extended to generic processors
}

func NewRouteBuilder(context GamelContext) *RouteBuilder {
	return &RouteBuilder{
		context: context,
	}
}

func (builder *RouteBuilder) From(uri string) *RouteBuilder {
	builder.input = uri
	return builder
}

func (builder *RouteBuilder) To(uri string) *RouteBuilder {
	builder.output = uri
	return builder
}

func (builder *RouteBuilder) Build() (Route, error) {
	targetComponent, err := builder.context.GetComponentFromURI(builder.output)
	if err != nil {
		return nil, err
	}

	targetEndpoint, err := targetComponent.NewEndpoint(builder.output)
	if err != nil {
		return nil, err
	}

	target, err := targetEndpoint.NewProducer()
	if err != nil {
		return nil, err
	}

	sourceComponent, err := builder.context.GetComponentFromURI(builder.input)
	if err != nil {
		return nil, err
	}

	sourceEndpoint, err := sourceComponent.NewEndpoint(builder.input)
	if err != nil {
		return nil, err
	}

	source, err:= sourceEndpoint.NewConsumer(target)
	if err != nil {
		return nil, err
	}

	return &DefaultRoute{
		services: []Service{source, target},
	}, nil
}
