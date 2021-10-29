package User

import (
	"encoding/json"
	"errors"
	"mnc-backend/formatter"
	"mnc-backend/http/enum"
	"mnc-backend/http/utils"
	"mnc-backend/models/requests"
	"mnc-backend/repositories/balance"
	"mnc-backend/repositories/transfer_history"
	"mnc-backend/repositories/user"
	"net/http"

	"github.com/gocraft/dbr"
)

type Transfer struct {
	DBConn *dbr.Connection
}

func (c Transfer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("Email").(string)

	var request requests.TransferRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&request)
	if err != nil {
		formatter.ERROR(w, http.StatusBadRequest, err)
		return
	}

	sess := c.DBConn.NewSession(nil)

	err = utils.CheckAuthorization(r, sess)
	if err != nil {
		formatter.ERROR(w, http.StatusForbidden, err)
		return
	}

	currentUser, err := user.GetUserByEmail(sess, email)
	if err != nil {
		formatter.ERROR(w, http.StatusBadRequest, err)
		return
	}

	currentBalance, err := balance.GetBalance(sess, currentUser.ID)
	if err != nil {
		formatter.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if currentBalance.Amount < request.Amount {
		formatter.ERROR(w, http.StatusBadRequest, errors.New(enum.InsufficientBalance))
		return
	}

	receiver, err := balance.GetBalance(sess, request.ReceiverID)
	if err != nil {
		formatter.ERROR(w, http.StatusBadRequest, err)
		return
	}

	senderNewBalance := currentBalance.Amount - request.Amount
	receiverNewBalance := receiver.Amount + request.Amount

	tx, err := sess.Begin()
	if err != nil {
		tx.Rollback()
		formatter.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	data := make(map[int]int)
	data[currentUser.ID] = senderNewBalance
	data[receiver.ID] = receiverNewBalance

	err = transfer_history.CreateTransfer(tx, request, currentUser.ID)
	if err != nil {
		tx.Rollback()
		formatter.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = balance.UpdateBalance(tx, data)
	if err != nil {
		tx.Rollback()
		formatter.ERROR(w, http.StatusBadRequest, err)
		return
	}

	tx.Commit()
	return
}
