package model

type BankAccount struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Email      string `json:"email" db:"email"`
	Nik        string `json:"nik" db:"nik"`
	Phone      string `json:"phone" db:"phone"`
	BankNo     string `json:"bank_no" db:"bank_no"`
	Balance    uint64 `json:"balance" db:"balance"`
	CreatedAt  string `json:"created_at" db:"created_at"`
	ModifiedAt string `json:"modified_at" db:"modified_at"`
	CreatedBy  string `json:"created_by" db:"created_by"`
	ModifiedBy string `json:"modified_by" db:"modified_by"`
}
