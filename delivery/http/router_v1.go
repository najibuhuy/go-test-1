package http

import (
	"github.com/gofiber/fiber/v2"
)

func RouterGroupV1(app *fiber.App, handler handler) {

	v1 := app.Group("/v1")
	bank_account := v1.Group("/bank-account")
	{
		bank_account.Post("/daftar", handler.bankAccountHandler.SignUpHandler)
		bank_account.Get("/saldo/:bankNo", handler.bankAccountHandler.CheckBalanceHandler)
		bank_account.Post("/tarik", handler.bankAccountHandler.WithdrawBalanceHandler)
		bank_account.Post("/tabung", handler.bankAccountHandler.DepositBalanceHandler)
	}

	//route > http handler > feature > repository

}
