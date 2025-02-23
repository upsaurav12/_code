package order

import (
	"testing"

	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
)

func TestComputeTotal(t *testing.T) {

	t.Run("nominal case", func(t *testing.T) {
		o := Order{
			ID:                "45",
			CurrencyAlphacode: "EUR",
			Items: []Item{
				{
					ID:        "458",
					Quantity:  2,
					UnitPrice: money.New(100, "EUR"),
				},
			},
		}
		total, err := o.ComputerTotal()
		assert.NoError(t, err)
		assert.Equal(t, int64(200), total.Amount())
		assert.Equal(t, "EUR", total.Currency().Code)
	})

	t.Run("Currency Issues", func(t *testing.T) {
		o := Order{
			ID:                "123",
			CurrencyAlphacode: "USD",
			Items: []Item{
				{
					ID:        "353",
					Quantity:  2,
					UnitPrice: money.New(100, "USD"),
				},
			},
		}
		_, err := o.ComputerTotal()
		assert.NoError(t, err)
	})
}
