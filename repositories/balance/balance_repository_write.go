package balance

import (
	"mnc-backend/models/db"

	"github.com/gocraft/dbr"
)

func CreateBalance(sess *dbr.Session, userID int) error {
	balance := db.Balance{
		UserID: userID,
		Amount: 1000000,
	}

	columns := []string{
		"user_id",
		"amount",
	}

	_, err := sess.InsertInto(db.Balance{}.TableName()).
		Columns(columns...).
		Record(balance).
		Exec()
	if err != nil {
		return err
	}

	return nil
}

func UpdateBalance(tx *dbr.Tx, data map[int]int) error {
	for id := range data {
		_, err := tx.Update(db.Balance{}.TableName()).
			Set("amount", data[id]).
			Where("user_id = ?", id).
			Exec()
		if err != nil {
			return err
		}
	}

	return nil
}
