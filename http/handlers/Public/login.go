package Public

import (
	"encoding/json"
	"errors"
	"mnc-backend/formatter"
	"mnc-backend/http/enum"
	"mnc-backend/http/utils"
	"mnc-backend/models/requests"
	"mnc-backend/repositories/access_token"
	"mnc-backend/repositories/user"
	"net/http"

	"github.com/gocraft/dbr"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	DBConn *dbr.Connection
}

func (c Login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request requests.LoginRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&request)
	if err != nil {
		formatter.ERROR(w, http.StatusBadRequest, err)
		return
	}

	sess := c.DBConn.NewSession(nil)

	currentUser, err := user.GetUserByEmail(sess, request.Email)
	if err != nil {
		formatter.ERROR(w, http.StatusUnauthorized, errors.New(enum.PasswordIncorrect))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(request.Password))
	if err != nil {
		formatter.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	accessToken, err := utils.EncodeAuthToken(currentUser.ID, currentUser.FullName, currentUser.Email, currentUser.UserType)
	if err != nil {
		formatter.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	insertAccessTokenRequest := requests.InsertAccessTokenRequest{
		Email:       request.Email,
		AccessToken: accessToken,
	}
	err = access_token.InsertorUpdateAccessToken(sess, insertAccessTokenRequest)
	if err != nil {
		formatter.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	formatter.JSON(w, http.StatusOK, struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: accessToken,
	})
	return
}
