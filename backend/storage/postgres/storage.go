package postgres

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vointini/vointini/backend/storage/storageimpl"
	"log"
	"time"
)

// Check storageimpl.Storage implementation for StoragePostgreSQL
var _ storageimpl.Storage = &StoragePostgreSQL{}

type StoragePostgreSQL struct {
	dbctx context.Context
	db    *pgxpool.Pool
}

func New(user, password, database, host string, port uint16) (s *StoragePostgreSQL) {
	if port == 0 {
		port = 5432
	}

	ctx := context.Background()
	db, err := pgxpool.Connect(ctx, fmt.Sprintf(
		`postgres://%s:%s@%s:%d/%s?application_name=vointini`,
		user, password, host, port, database,
	))

	if err != nil {
		panic(err)
	}

	go func(database *pgxpool.Pool) {
		for range time.Tick(time.Minute * 1) {
			err := database.Ping(context.TODO())
			if err != nil {
				log.Printf(`ping failed: %v`, err)
			}
		}
	}(db)

	s = &StoragePostgreSQL{
		db:    db,
		dbctx: ctx,
	}

	appName, err := s.getVal(`application_name`)
	if err != nil {
		panic(err)
	}

	version, err := s.getVal(`server_version`)
	if err != nil {
		panic(err)
	}

	encoding, err := s.getVal(`server_encoding`)
	if err != nil {
		panic(err)
	}

	timezone, err := s.getVal(`timezone`)
	if err != nil {
		panic(err)
	}

	log.Printf(`PostgreSQL (ver: %s, app: %s) is using encoding %q and timezone %q`,
		version, appName, encoding, timezone,
	)

	return s
}

func (s *StoragePostgreSQL) getVal(v string) (value string, err error) {
	err = pgxscan.Get(context.TODO(), s.db, &value,
		`SHOW `+v,
	)

	if err != nil {
		return ``, err
	}

	return value, nil
}
