package db

type Balance struct {
	ID     int `db:"id" json:"id"`
	UserID int `db:"user_id" json:"user_id"`
	Amount int `db:"amount" json:"amount"`
}

func (c Balance) TableName() string {
	return "balances"
}
