package feature

import (
	"context"
	"fmt"
	"go-test-1/domain/bank_account/model"
)

func (baf bankAccountFeature) SignUpFeature(ctx context.Context, request *model.SignUpRequest) (response model.SignUpResponse, err error) {
	resp, err := baf.bankAccountRepository.SignUpRepository(ctx, request)
	fmt.Println(err, "err in feature")
	if err != nil {
		return
	}
	response = model.SignUpResponse{
		BankNo: resp.BankNo,
		Nik:    resp.Nik,
		Name:   resp.Name,
		Phone:  resp.Phone,
	}
	return
}
