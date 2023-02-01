package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
	"xm/config"
	"xm/internal/handlers"
	"xm/internal/repositories"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

var logger zerolog.Logger
var cfg config.Configuration

func init() {
	logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	cfg = config.Load(logger)
}

func main() {
	var err error
	var db *sqlx.DB

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err = sqlx.Connect(cfg.Postgres.SqlDriver, cfg.Postgres.SqlDSN)
	if err != nil {
		logger.Info().Timestamp().Msg(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	companyRepository := repositories.NewCompanyRepository(ctx, db)
	companyHandler := handlers.NewCompanyHandler(ctx, companyRepository, logger)

	router := chi.NewRouter()

	router.Group(func(router chi.Router) {
		router.Get("/v1/company/{id}", companyHandler.GetCompany)
	})

	router.Group(func(router chi.Router) {
		router.Use(validateToken)

		router.Post("/v1/company", companyHandler.CreateCompany)
		router.Delete("/v1/company/{id}", companyHandler.DeleteCompany)
		router.Patch("/v1/company/{id}", companyHandler.UpdateCompany)
	})

	server := http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Port),
		WriteTimeout:      cfg.HTTPServerTimeout,
		ReadTimeout:       cfg.HTTPServerTimeout,
		IdleTimeout:       time.Second,
		ReadHeaderTimeout: 0,
		Handler:           router,
	}

	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router)
	if err != nil {
		err := server.Shutdown(ctx)
		if err != nil {
			logger.Info().Timestamp().Msg(err.Error())
			os.Exit(1)
		}
	}
}

func validateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := jwtauth.TokenFromHeader(r)

		if token != cfg.Token {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte(errors.New("invalid token").Error()))
			if err != nil {
				logger.Info().Timestamp().Msg(err.Error())
			}
			return
		}

		next.ServeHTTP(w, r)
	})
}
