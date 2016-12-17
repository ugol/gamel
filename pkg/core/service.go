package core

type Service interface {
	Start() error
	Stop() error
	Suspend() error
	Resume() error
}
