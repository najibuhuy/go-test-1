package repository

import (
	"context"
	"errors"
	"go-test-1/domain/bank_account/model"
	"go-test-1/domain/shared/constant"
	Error "go-test-1/domain/shared/error"
)

func (bar bankAccountRepository) WithDrawBalanceRepository(ctx context.Context, request *model.WithDrawBalanceRequest) (response model.WithDrawBalanceResponse, err error) {

	if request.Amount <= 0 {
		err = Error.New(constant.ErrInvalidRequest, "withdraw amount must be greater than zero", err)
		return response, err
	}

	tx, err := bar.Database.Pool.Begin(ctx)
	if err != nil {
		return response, err
	}
	defer tx.Rollback(ctx)

	// Check balance before withdrawing
	var currentBalance float64
	err = tx.QueryRow(ctx, "SELECT balance FROM bank_accounts WHERE bank_no = $1", request.BankNo).Scan(&currentBalance)
	if err != nil {
		err = Error.New(constant.ErrInvalidRequest, "error select query check balance", errors.New("error query"))
		return response, err
	}

	if currentBalance < float64(request.Amount) {
		err = Error.New(constant.ErrUserInsufficientBalance, "insufficient balance", errors.New("error body"))
		return response, err
	}

	// Update balance
	_, err = tx.Exec(ctx, "UPDATE bank_accounts SET balance = balance - $1 WHERE bank_no = $2", request.Amount, request.BankNo)
	if err != nil {
		err = Error.New(constant.ErrInvalidRequest, "error update query", errors.New("error query"))
		return response, err
	}

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
	response = model.WithDrawBalanceResponse{
		Balance: int64(account.Balance),
		BankNo:  account.BankNo,
	}
	return response, nil
}
