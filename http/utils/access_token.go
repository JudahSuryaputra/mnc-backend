package utils

import (
	"errors"
	"mnc-backend/http/enum"
	"mnc-backend/models/requests"
	"mnc-backend/repositories/access_token"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gocraft/dbr"
	"github.com/spf13/viper"
)

func EncodeAuthToken(id int, username, email, userType string) (string, error) {
	claims := jwt.MapClaims{}
	claims["ID"] = id
	claims["FirstName"] = username
	claims["Email"] = email
	claims["UserType"] = userType
	claims["ExpiresAt"] = time.Now().Add(time.Minute * 30).Unix()

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	return token.SignedString([]byte(viper.GetString("MNC_SECRET")))
}

func CheckAuthorization(r *http.Request, sess *dbr.Session) error {
	authorization := r.Header.Get("Authorization")
	accessToken := strings.Split(authorization, "Bearer ")
	if len(accessToken) != 2 {
		return errors.New(enum.UnauthorizedUser)
	}
	currentToken, err := access_token.GetUserAccessToken(sess, accessToken[1])
	if currentToken != nil || err != nil {
		return err
	}
	return nil
}

func InsertAccessToken(sess *dbr.Session, email string, accessToken string) error {
	insertAccessTokenRequest := requests.InsertAccessTokenRequest{
		Email:       email,
		AccessToken: accessToken,
	}

	err := access_token.InsertAccessToken(sess, insertAccessTokenRequest)
	if err != nil {
		return err
	}

	return nil
}
