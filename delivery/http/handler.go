package http

import (
	"go-test-1/delivery/container"
	"go-test-1/domain/bank_account"
)

type handler struct {
	bankAccountHandler bank_account.BankAccountHandler
}

func SetupHandler(container container.Container) handler {
	return handler{
		bankAccountHandler: bank_account.NewBankAccountHandler(container.BankAccountFeature),
	}
}
