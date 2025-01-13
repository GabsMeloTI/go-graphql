package repository

import (
	"context"
	"database/sql"
	db "github.com/GabsMeloTI/go-graphql/db/sqlc"
)

type InterfaceRepository interface {
	CreateBasicInfo(context.Context, db.CreateBasicInfoParams) (db.BasicInfo, error)
}

type Repository struct {
	Conn    *sql.DB
	DBtx    db.DBTX
	Queries *db.Queries
	SqlConn *sql.DB
}

func NewBasicInfoRepository(Conn *sql.DB) *Repository {
	q := db.New(Conn)
	return &Repository{
		Conn:    Conn,
		DBtx:    Conn,
		Queries: q,
		SqlConn: Conn,
	}
}

func (r *Repository) CreateBasicInfo(ctx context.Context, arg db.CreateBasicInfoParams) (db.BasicInfo, error) {
	return r.Queries.CreateBasicInfo(ctx, arg)
}
