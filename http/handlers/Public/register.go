package Public

import (
	"encoding/json"
	"errors"
	"mnc-backend/formatter"
	"mnc-backend/http/enum"
	"mnc-backend/models/requests"
	"mnc-backend/repositories/balance"
	"mnc-backend/repositories/user"
	"net/http"

	"github.com/gocraft/dbr"
	"golang.org/x/crypto/bcrypt"
)

type Register struct {
	DBConn *dbr.Connection
}

func (c Register) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request requests.RegisterRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&request)
	if err != nil {
		formatter.ERROR(w, http.StatusBadRequest, err)
		return
	}

	sess := c.DBConn.NewSession(nil)

	err = checkExistingUser(sess, request.Email)
	if err != nil {
		formatter.ERROR(w, http.StatusBadRequest, err)
		return
	}

	hashedPassword, err := HashPassword(request.Password)
	if err != nil {
		formatter.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	request.Password = hashedPassword

	userID, err := user.CreateUser(sess, request)
	if err != nil {
		formatter.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = balance.CreateBalance(sess, userID)
	if err != nil {
		formatter.ERROR(w, http.StatusBadRequest, err)
		return
	}

	return
}

func checkExistingUser(sess *dbr.Session, email string) error {
	currentUser, err := user.GetUserByEmail(sess, email)
	if currentUser != nil && err == nil {
		return errors.New(enum.EmailDuplicate)
	}

	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
