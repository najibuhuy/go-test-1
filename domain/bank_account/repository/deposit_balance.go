package repository

import (
	"context"
	"errors"
	"go-test-1/domain/bank_account/model"
	"go-test-1/domain/shared/constant"
	Error "go-test-1/domain/shared/error"
)

func (bar *bankAccountRepository) DepositBalanceRepository(ctx context.Context, request *model.DepositBalanceRequest) (response model.DepositBalanceResponse, err error) {

	if request.Amount <= 0 {
		err = Error.New(constant.ErrInvalidRequest, "deposit amount must be greater than zero", errors.New("error amount"))
		return response, err
	}

	tx, err := bar.Database.Pool.Begin(ctx)
	if err != nil {
		err = Error.New(constant.ErrInvalidRequest, "error begin db", errors.New("error begin db"))
		return response, err
	}
	defer tx.Rollback(ctx)

	// Update balance
	_, err = tx.Exec(ctx, "UPDATE bank_accounts SET balance = balance + $1 WHERE bank_no = $2", request.Amount, request.BankNo)
	if err != nil {
		err = Error.New(constant.ErrInvalidRequest, "error update query", errors.New("error query"))
		return response, err
	}

	// Fetch updated account details
	row := tx.QueryRow(ctx, "SELECT bank_no, balance FROM bank_accounts WHERE bank_no = $1", request.BankNo)
	var account model.BankAccount
	if err := row.Scan(&account.BankNo, &account.Balance); err != nil {
		err = Error.New(constant.ErrInvalidRequest, "error select query", errors.New("error query"))
		return response, err
	}

	// Commit transaction
	if err := tx.Commit(ctx); err != nil {
		err = Error.New(constant.ErrInvalidRequest, "error committing transaction", errors.New("error commit"))
		return response, err
	}

	response = model.DepositBalanceResponse{
		Balance: int64(account.Balance),
		BankNo:  account.BankNo,
	}
	return
}
