package core

type Service interface {
	Start()		error
	Stop()		error
}