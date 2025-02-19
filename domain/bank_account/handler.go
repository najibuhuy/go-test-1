package bank_account

import (
	"go-test-1/domain/bank_account/feature"
	"go-test-1/domain/bank_account/model"

	"go-test-1/domain/shared/context"
	Error "go-test-1/domain/shared/error"
	"go-test-1/domain/shared/response"

	"fmt"
	"go-test-1/domain/bank_account/constant"

	"github.com/gofiber/fiber/v2"
)

type BankAccountHandler interface {
	CheckBalanceHandler(c *fiber.Ctx) error
	WithdrawBalanceHandler(c *fiber.Ctx) error
	DepositBalanceHandler(c *fiber.Ctx) error
	SignUpHandler(c *fiber.Ctx) error
}

type bankAccountHandler struct {
	feature feature.BankAccountFeature
}

func NewBankAccountHandler(feature feature.BankAccountFeature) BankAccountHandler {
	return &bankAccountHandler{
		feature: feature,
	}
}

func (bah bankAccountHandler) CheckBalanceHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	bankNo := c.Params("bankNo")
	if bankNo == "" || bankNo == "0" {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrBankNoNil))
		return response.ResponseErrorWithContext(ctx, err)
	}

	results, err := bah.feature.CheckBalanceFeature(ctx, bankNo)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetBalanceSuccess, results)
}

func (bah bankAccountHandler) WithdrawBalanceHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)
	request := new(model.WithDrawBalanceRequest)
	if err := c.BodyParser(request); err != nil {
		fmt.Println(err, "err")
		err = Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
		return response.ResponseErrorWithContext(ctx, err)
	} else if request.Amount == 0 || request.BankNo == "" {
		err = Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
		return response.ResponseErrorWithContext(ctx, err)
	}
	results, err := bah.feature.WithDrawBalanceFeature(ctx, request)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgUpdateBalanceSuccess, results)
}

func (bah bankAccountHandler) DepositBalanceHandler(c *fiber.Ctx) error {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	request := new(model.DepositBalanceRequest)
	if err := c.BodyParser(request); err != nil {
		fmt.Println(err, "err")
		err = Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
		return response.ResponseErrorWithContext(ctx, err)
	} else if request.Amount == 0 || request.BankNo == "" {
		err = Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
		return response.ResponseErrorWithContext(ctx, err)
	}

	results, err := bah.feature.DepositBalanceFeature(ctx, request)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgUpdateBalanceSuccess, results)
}

func (bah bankAccountHandler) SignUpHandler(c *fiber.Ctx) error {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	request := new(model.SignUpRequest)
	if err := c.BodyParser(request); err != nil {
		fmt.Println(err, "err")
		err = Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
		return response.ResponseErrorWithContext(ctx, err)
	} else if request.Name == "" || request.Nik == "" || request.Phone == "" {
		err = Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
		return response.ResponseErrorWithContext(ctx, err)
	}

	results, err := bah.feature.SignUpFeature(ctx, request)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgSignUpSuccess, results)
}
