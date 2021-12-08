package factory

import "github.com/HigorJardini/full-cycle/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
