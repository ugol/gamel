package core


type Exchange struct {
	In	*Message
}

func NewExchange() *Exchange  {
	return &Exchange{
		In: &Message{},
	}
}