package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetAnErrorIfIDIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "id is riquired")
}

func Test_If_It_Get_An_Error_If_Price_Is_Blank(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validate(), "price must be greater tan zero")
}

func Test_If_It_Get_An_Error_If_Tax_Is_Blank(t *testing.T) {
	order := Order{ID: "123", Price: 10.0}
	assert.Error(t, order.Validate(), "price must be greater tan zero")
}

func Test_If_It_Get_No_Error(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 0.3}
	assert.Nil(t, order.Validate(), "ok")
}

func Test_Final_Price(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 0.3}
	order.CalculateFinalPrice()
	assert.Equal(t, order.FinalPrice, 10.3)
}
