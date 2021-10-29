package middlewares

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mnc-backend/formatter"
	"mnc-backend/http/enum"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func SetContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func UserAuthJwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, status, err := extractToken(r)
		if err != nil {
			formatter.ERROR(w, status, err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "Email", claims["Email"])
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := os.OpenFile("log_history.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}

		log.SetOutput(file)
		logHistory := fmt.Sprintf("method=%s, host=%s, url=%s, user_agent=%s", r.Method, r.Host, r.URL, r.UserAgent())
		log.Println(logHistory)
		next.ServeHTTP(w, r)
	})
}

func extractToken(r *http.Request) (*jwt.Token, int, error) {
	header := r.Header.Get("Authorization")
	header = strings.TrimSpace(header)

	if header == "" {
		return nil, http.StatusForbidden, errors.New(enum.UnauthorizedUser)
	}

	headerSplit := strings.Split(header, "Bearer ")
	if len(headerSplit) != 2 {
		return nil, http.StatusForbidden, errors.New(enum.UnauthorizedUser)
	}

	token, err := jwt.Parse(headerSplit[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("MNC_SECRET")), nil
	})
	if err != nil {
		return nil, http.StatusForbidden, err
	}
	
	return token, http.StatusOK, err
}
