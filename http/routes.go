package http

import (
	"log"
	"mnc-backend/http/handlers/Public"
	"mnc-backend/http/handlers/User"
	"mnc-backend/http/middlewares"
	"net/http"
	"os"

	"github.com/gocraft/dbr"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(dbConn *dbr.Connection) {
	a.Router = mux.NewRouter().StrictSlash(true)

	a.publicRoutes(dbConn)
	a.userRoutes(dbConn)
}

func (a *App) publicRoutes(dbConn *dbr.Connection) {
	a.Router.Use(middlewares.LoggingMiddleware, middlewares.SetContentTypeMiddleware)

	register := Public.Register{DBConn: dbConn}
	login := Public.Login{DBConn: dbConn}

	a.Router.Handle("/v1/register", register).Methods(http.MethodPost)
	a.Router.Handle("/v1/login", login).Methods(http.MethodPost)
}

func (a *App) userRoutes(dbConn *dbr.Connection) {
	user := a.Router.PathPrefix("/v1/api").Subrouter()
	user.Use(middlewares.UserAuthJwtVerify)

	transfer := User.Transfer{DBConn: dbConn}
	logout := User.Logout{DBConn: dbConn}

	user.Handle("/transfer", transfer).Methods(http.MethodPost)
	user.Handle("/logout", logout).Methods(http.MethodPatch)
}

func (a *App) RunServer() {
	port := viper.GetString("PORT")

	headersOK := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"POST", "PUT"})

	log.Printf("\nServer starting on Port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CombinedLoggingHandler(os.Stderr, handlers.CORS(headersOK, originsOK, methodsOK)(a.Router))))
}
