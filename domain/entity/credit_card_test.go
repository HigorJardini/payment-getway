package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("90000000000000000", "Higor Teste Card", 11, 2026, 123)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("5521919080210388", "Higor Teste Card", 11, 2026, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("5521919080210388", "Higor Teste Card", 13, 2026, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5521919080210388", "Higor Teste Card", 0, 2026, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5521919080210388", "Higor Teste Card", 11, 2026, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0)

	_, err := NewCreditCard("5521919080210388", "Higor Teste Card", 12, lastYear.Year(), 123)
	assert.Equal(t, "invalid expiration year", err.Error())
}

func TestCreditCardExpirationCvv(t *testing.T) {
	_, err := NewCreditCard("5521919080210388", "Higor Teste Card", 12, 2026, 1234)
	assert.Equal(t, "invalid expiration cvv", err.Error())

	_, err = NewCreditCard("5521919080210388", "Higor Teste Card", 11, 2026, 123)
	assert.Nil(t, err)
}
