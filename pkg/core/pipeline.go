package core


type Pipeline struct {
	steps	[]Processor
}

func NewPipeline(processors ...Processor) *Pipeline {
	return &Pipeline{
		steps: processors,
	}
}

// TODO connect all processors in the pipeline (output with input)