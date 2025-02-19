package model

type CheckBalanceResponse struct {
	Balance int64  `json:"balance"`
	BankNo  string `json:"bankNo"`
}

type WithDrawBalanceResponse struct {
	Balance int64  `json:"balance"`
	BankNo  string `json:"bankNo"`
}

type DepositBalanceResponse struct {
	Balance int64  `json:"balance"`
	BankNo  string `json:"bankNo"`
}

type SignUpResponse struct {
	Nik    string `json:"nik"`
	Phone  string `json:"phone"`
	Name   string `json:"name"`
	BankNo string `json:"bankNo"`
}
