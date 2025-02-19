package feature

import (
	"context"
	"go-test-1/domain/bank_account/model"
)

func (baf bankAccountFeature) CheckBalanceFeature(ctx context.Context, bankNo string) (response model.CheckBalanceResponse, err error) {

	result, err := baf.bankAccountRepository.CheckBalanceRepository(ctx, bankNo)
	if err != nil {
		return
	}

	response = model.CheckBalanceResponse{
		Balance: result.Balance,
		BankNo:  result.BankNo,
	}
	return
}
