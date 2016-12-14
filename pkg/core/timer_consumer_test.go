package core

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTimerConsumerCreation(t *testing.T) {
	context := GamelContext{
		Name: "gamel",
	}

	component, _ := context.GetComponent("timer")

	endpoint, _ := component.NewEndpoint("timer:tick?period=3000")

	uri := endpoint.Uri()
	assert.Equal(t, "timer:tick?period=3000", uri)

	consumer, _ := endpoint.NewConsumer()
	assert.NotNil(t, consumer)

	consumer.Start()
}
