package core

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestExchangeCreation(t *testing.T) {
	exchange := NewExchange()

	assert.NotNil(t, exchange.In)
	assert.Nil(t, exchange.In.Body)
}
