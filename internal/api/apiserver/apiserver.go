package apiserver

import (
	"Go-http-rest-api/internal/api/store/sqlstore"
	"database/sql"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDB(config.DataBaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	store := sqlstore.New(db)
	srv := newServer(store)
	return http.ListenAndServe(config.BindAdrr, srv)
}
func newDB(dataBaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataBaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
