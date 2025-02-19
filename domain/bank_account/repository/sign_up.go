package repository

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"go-test-1/domain/bank_account/model"
	"go-test-1/domain/shared/constant"
	Error "go-test-1/domain/shared/error"
	"math/big"
)

func generateBankNo() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(9000000000))
	return fmt.Sprintf("10%09d", n.Int64()) // Always starts with 10
}

func (bar bankAccountRepository) SignUpRepository(ctx context.Context, request *model.SignUpRequest) (response model.SignUpResponse, err error) {

	tx, err := bar.Database.Pool.Begin(ctx)
	if err != nil {
		return response, err
	}
	defer tx.Rollback(ctx)

	// Check if ID number or Phone already exists
	var exists bool
	err = tx.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM bank_accounts WHERE id_number=$1 OR phone=$2)",
		request.Nik, request.Phone).Scan(&exists)

	if err != nil {
		err = Error.New(constant.ErrInvalidRequest, "error select query check accounts exist or not", errors.New("error query"))
		return response, err
	}

	if exists {
		err = Error.New(constant.ErrUserExist, "Account Existed", errors.New("duplicate entry"))
		fmt.Println(err, "err")
		return response, err
	}

	// Generate a unique bank number
	bankNo := generateBankNo()
	fmt.Println(request.Nik, "request.Nik")

	// Insert into the database
	_, err = tx.Exec(ctx,
		"INSERT INTO bank_accounts (name, id_number, phone, bank_no, balance) VALUES ($1, $2, $3, $4, $5)",
		request.Name, request.Nik, request.Phone, bankNo, 0.00)
	if err != nil {
		err = Error.New(constant.ErrInvalidRequest, "error insert bank account query", errors.New("error insert"))
		return response, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		err = Error.New(constant.ErrInvalidRequest, "error committing transaction", errors.New("error commit"))
		return response, err
	}

	response = model.SignUpResponse{
		BankNo: bankNo,
		Nik:    request.Nik,
		Name:   request.Name,
		Phone:  request.Phone,
	}
	return
}
