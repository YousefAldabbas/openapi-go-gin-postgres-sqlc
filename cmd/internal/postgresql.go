// Code adapted from:
// https://github.com/MarioCarrion/todo-api-microservice-example

package internal

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
)

// NewPostgreSQL instantiates the PostgreSQL database using configuration defined in environment variables.
func NewPostgreSQL(conf *envvar.Configuration) (*pgxpool.Pool, error) {
	get := func(v string) string {
		res, err := conf.Get(v)
		if err != nil {
			log.Fatalf("Couldn't get configuration value for %s: %s", v, err)
		}

		return res
	}

	// XXX: We will revisit this code in future episodes replacing it with another solution
	databaseHost := get("DATABASE_HOST")
	databasePort := get("DATABASE_PORT")
	databaseUsername := get("DATABASE_USERNAME")
	databasePassword := get("DATABASE_PASSWORD")
	databaseName := get("DATABASE_NAME")
	databaseSSLMode := get("DATABASE_SSLMODE")
	// XXX: -

	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(databaseUsername, databasePassword),
		Host:   fmt.Sprintf("%s:%s", databaseHost, databasePort),
		Path:   databaseName,
	}

	q := dsn.Query()
	q.Add("sslmode", databaseSSLMode)

	dsn.RawQuery = q.Encode()

	pool, err := pgxpool.Connect(context.Background(), dsn.String())
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "pgxpool.Connect")
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "db.Ping")
	}

	return pool, nil
}
