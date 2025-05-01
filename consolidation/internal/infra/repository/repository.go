package repository

import (
	"database/sql"
	"errors"

	"github.com/johnkgs/imersao11-consolidation/internal/infra/db"
)

var ErrQueriesNotSet = errors.New("queries not set")

type Repository struct {
	dbConn *sql.DB
	*db.Queries
}

func (r *Repository) SetQuery(q *db.Queries) {
	r.Queries = q
}

func (r *Repository) Validade() error {
	if r.Queries == nil {
		return ErrQueriesNotSet
	}
	return nil
}
