package core

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestContextCreation(t *testing.T) {
	context := GamelContext{
		Name: "gamel-1",
	}
	assert.Equal(t, "gamel-1", context.Name)
	assert.Equal(t, Idle, context.Status)

	context.Start()
	assert.Equal(t, Started, context.Status)
}

func TestComponentFactory(t *testing.T) {
	context := GamelContext{
		Name: "gamel-1",
	}

	timer, err := context.GetComponent("timer")
	assert.Nil(t, err)

	_, err2 := timer.(TimerComponent)
	assert.False(t, err2)
}

func TestComponentFactoryUnknown(t *testing.T) {
	context := GamelContext{
		Name: "gamel-1",
	}

	_, err := context.GetComponent("ayeye")
	assert.NotNil(t, err)
}