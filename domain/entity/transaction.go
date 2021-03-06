package entity

import "errors"

type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	CreditCard   CreditCard
	Status       string
	ErrorMessage string
}

func NewTranscation() *Transaction {
	return &Transaction{}
}

func (t *Transaction) IsValid() error {

	if t.Amount > 1000 {
		return errors.New("you dont have limit for this transaction")
	} else if t.Amount < 1 {
		return errors.New("the amount must be greater than 1")
	} else {
		return nil
	}
}

func (t *Transaction) SetCreditCard(card CreditCard) {
	t.CreditCard = card
}
