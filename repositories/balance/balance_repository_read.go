package balance

import (
	"mnc-backend/models/db"

	"github.com/gocraft/dbr"
)

func GetBalance(sess *dbr.Session, userID int) (*db.Balance, error) {
	var balance *db.Balance

	query := sess.Select("*").
		From(db.Balance{}.TableName()).
		Where("user_id = ?", userID)

	err := query.LoadOne(&balance)
	if err != nil {
		return nil, err
	}

	return balance, nil
}
