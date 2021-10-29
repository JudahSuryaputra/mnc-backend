package transfer_history

import (
	"mnc-backend/models/db"
	"mnc-backend/models/requests"

	"github.com/gocraft/dbr"
)

func CreateTransfer(tx *dbr.Tx, r requests.TransferRequest, userID int) error {
	transfer := db.TransferHistory{
		SenderID:   userID,
		ReceiverID: r.ReceiverID,
		Amount:     r.Amount,
		Message:    r.Message,
	}

	columns := []string{
		"sender_id",
		"receiver_id",
		"amount",
		"message",
	}

	_, err := tx.InsertInto(db.TransferHistory{}.TableName()).
		Columns(columns...).
		Record(transfer).
		Exec()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
