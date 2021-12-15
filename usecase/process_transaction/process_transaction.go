package process_transaction

import (
	"github.com/HigorJardini/full-cycle/domain/entity"
	"github.com/HigorJardini/full-cycle/domain/repository"
)

type ProcessTransaction struct {
	Repository repository.TransactionRepository
}

func NewProcessTransaction(repository repository.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository}
}

func (p *ProcessTransaction) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	transaction := entity.NewTranscation()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount

	_, inValidCC := entity.NewCreditCard(input.CreditCardNumber, input.CreditCardName, input.CreditCardExpirationMonth, input.CreditCardExpirationYear, input.CreditCardCVV)

	if inValidCC != nil {
		err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, entity.REJECTED, inValidCC.Error())
		if err != nil {
			return TransactionDtoOutput{}, err
		} else {
			output := TransactionDtoOutput{
				ID:           transaction.ID,
				Status:       entity.REJECTED,
				ErrorMessage: inValidCC.Error(),
			}

			return output, nil
		}
	}

	return TransactionDtoOutput{}, nil
}
