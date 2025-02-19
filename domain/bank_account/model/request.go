package model

type WithDrawBalanceRequest struct {
	Amount int64  `json:"amount"`
	BankNo string `json:"bankNo"`
}

type DepositBalanceRequest struct {
	Amount int64  `json:"amount"`
	BankNo string `json:"bankNo"`
}

type SignUpRequest struct {
	Nik   string `json:"nik"`
	Phone string `json:"phone"`
	Name  string `json:"name"`
}
