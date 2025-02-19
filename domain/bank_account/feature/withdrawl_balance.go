package feature

import (
	"context"
	"go-test-1/domain/bank_account/model"
)

func (baf bankAccountFeature) WithDrawBalanceFeature(ctx context.Context, request *model.WithDrawBalanceRequest) (response model.WithDrawBalanceResponse, err error) {
	resp, err := baf.bankAccountRepository.WithDrawBalanceRepository(ctx, request)
	if err != nil {
		return
	}

	response = model.WithDrawBalanceResponse{
		Balance: resp.Balance,
		BankNo:  resp.BankNo,
	}
	return
}
