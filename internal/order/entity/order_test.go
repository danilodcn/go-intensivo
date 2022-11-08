package entity

import (
	"testing"

	// "github.com/danilodcn/go-intensivo/internal/order/entity"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyID_WhenCreateANewOrder_ThenShouldReceiveAnRandomID(t *testing.T) {
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid ID")
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "Casa"}
	assert.Error(t, order.IsValid(), "invalid Price")
}

func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "Casa"}
	assert.Error(t, order.IsValid(), "invalid Tax")
}

func TestGivenAnInvalidTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "Casa", Tax: 100, Price: 34}
	assert.Nil(t, order.IsValid())

	order = Order{ID: "Casa", Tax: -100, Price: 34}
	assert.Nil(t, order.IsValid())

	order = Order{ID: "Casa", Tax: -130, Price: 34}
	assert.Error(t, order.IsValid(), "Invalid price")
}

func TestGivenAnValidOrder_WhenICallNewOrder_ThenShouldReceiveACreateOrderWithAllParams(t *testing.T) {
	order, err := NewOrder(234, 32)
	assert.Nil(t, order.IsValid())
	assert.Equal(t, err, nil)
}

func TestGivenAPriceAndTax_WhenCallCalculateFinalPrice_ThenIShouldSetFinalPrice(t *testing.T) {
	order, err := NewOrder(100, 2)

	assert.Nil(t, err)
	err = order.CalculateFinalPrice()
	assert.Nil(t, err)
	assert.Equal(t, order.FinalPrice, 102.0)
}
