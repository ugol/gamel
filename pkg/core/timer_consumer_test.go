package core

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTimerConsumerCreation(t *testing.T) {
	context := NewGamelContext()

	component, _ := context.GetComponent("timer")

	endpoint, _ := component.NewEndpoint("timer:tick?period=3000")

	uri := endpoint.Uri()
	assert.Equal(t, "timer:tick?period=3000", uri)
	period := endpoint.Parameter("period")
	assert.Equal(t, "3000", period)

	logComponent, _ := context.GetComponent("log")
	logEndpoint, _ := logComponent.NewEndpoint("log:INFO")
	logProducer, _ := logEndpoint.NewProducer()

	consumer, _ := endpoint.NewConsumer(logProducer)
	assert.NotNil(t, consumer)
}
