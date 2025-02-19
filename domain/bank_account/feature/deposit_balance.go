package feature

import (
	"context"
	"go-test-1/domain/bank_account/model"
)

func (baf bankAccountFeature) DepositBalanceFeature(ctx context.Context, request *model.DepositBalanceRequest) (response model.DepositBalanceResponse, err error) {
	resp, err := baf.bankAccountRepository.DepositBalanceRepository(ctx, request)
	if err != nil {
		return
	}

	response = model.DepositBalanceResponse{
		Balance: resp.Balance,
		BankNo:  resp.BankNo,
	}
	return
}
