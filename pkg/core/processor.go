package core

type Processor interface {
	Service

	Process(exchange *Exchange)	error
}
