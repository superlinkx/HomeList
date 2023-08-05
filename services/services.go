package services

import (
	"database/sql"

	"github.com/superlinkx/HomeList/db/sqlite"
	"github.com/superlinkx/HomeList/services/list"
)

type Services struct {
	List list.Service
}

func Init(db *sql.DB) Services {
	queries := sqlite.New(db)
	return Services{
		List: list.NewService(queries),
	}
}
