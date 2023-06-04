// Code adapted from:
// https://github.com/MarioCarrion/todo-api-microservice-example

package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"sync"
	"time"

	zapadapter "github.com/jackc/pgx-zap"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jackc/pgx/v5/tracelog"
	"go.uber.org/zap"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
)

var pgxAfterConnectLock = sync.Mutex{}

// New instantiates the PostgreSQL database using configuration defined in environment variables.
func New(logger *zap.SugaredLogger) (*pgxpool.Pool, *sql.DB, error) {
	cfg := internal.Config()
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.Postgres.User, cfg.Postgres.Password),
		Host:   fmt.Sprintf("%s:%s", cfg.Postgres.Server, fmt.Sprint(cfg.Postgres.Port)),
		Path:   cfg.Postgres.DB,
	}

	q := dsn.Query()
	q.Add("sslmode", os.Getenv("DATABASE_SSLMODE"))

	dsn.RawQuery = q.Encode()

	poolConfig, err := pgxpool.ParseConfig(dsn.String())
	if err != nil {
		return nil, nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "pgx.ParseConfig")
	}

	if cfg.Postgres.TraceEnabled {
		poolConfig.ConnConfig.Tracer = &tracelog.TraceLog{
			Logger:   zapadapter.NewLogger(logger.Desugar()),
			LogLevel: tracelog.LogLevelTrace,
		}
	}

	searchPaths := []string{"public", "xo_tests"}
	conn, err := pgx.Connect(context.Background(), dsn.String())
	if err != nil {
		return nil, nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "could not connect to database")
	}
	typeNames, err := queryDatabaseTypeNames(context.Background(), conn, searchPaths...)
	if err != nil {
		return nil, nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "could not query database types")
	}

	afterConnectRun := false

	// called after a connection is established, but before it is added to the pool.
	// Will run once.
	poolConfig.AfterConnect = func(ctx context.Context, c *pgx.Conn) error {
		pgxAfterConnectLock.Lock()
		defer pgxAfterConnectLock.Unlock()

		err = registerDataTypes(context.Background(), c, typeNames)
		if err != nil {
			return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "could not register data types")
		}

		afterConnectRun = true

		return nil
	}

	pgxPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "pgxpool.New")
	}

	if err := pgxPool.Ping(context.Background()); err != nil {
		return nil, nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "db.Ping")
	}

	sqlPool, err := sql.Open("pgx", pgxPool.Config().ConnString())
	if err != nil {
		return nil, nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "sql.Open")
	}

	for !afterConnectRun {
		time.Sleep(200 * time.Millisecond)
	}

	return pgxPool, sqlPool, nil
}

// registerDataTypes automatically registers all enums and tables types
// for proper encoding/decoding in pgx.
// See https://pkg.go.dev/github.com/jackc/pgx/v5@v5.3.1/pgtype#hdr-New_PostgreSQL_Type_Support
func registerDataTypes(ctx context.Context, conn *pgx.Conn, typeNames []string) error {
	for _, typeName := range typeNames {
		// fmt.Printf("registering %v\n", typeName)
		dataType, err := conn.LoadType(ctx, typeName)
		if err != nil {
			return err
		}
		conn.TypeMap().RegisterType(dataType)
	}

	return nil
}

func queryDatabaseTypeNames(ctx context.Context, conn *pgx.Conn, searchPaths ...string) ([]string, error) {
	var typeNames []string
	for _, sp := range searchPaths {
		query := fmt.Sprintf(`SELECT table_name
	FROM information_schema.tables
	WHERE table_schema IN ('%s')`, sp)
		tableTypes, err := queryTypeNames(conn, query, sp)
		if err != nil {
			return []string{}, fmt.Errorf("querying table names: %w", err)
		}

		query = fmt.Sprintf(`SELECT t.typname AS enum_name
	FROM pg_type t
	INNER JOIN pg_namespace n ON n.oid = t.typnamespace
	WHERE t.typtype = 'e' AND n.nspname IN ('%s');`, sp)
		enumTypes, err := queryTypeNames(conn, query, sp)
		if err != nil {
			return []string{}, fmt.Errorf("querying enum names: %w", err)
		}

		// register enumTypes first, in case they're used in tables
		tn := append(enumTypes, tableTypes...)
		typeNames = append(typeNames, tn...)
	}

	return typeNames, nil
}

func queryTypeNames(conn *pgx.Conn, query string, searchPath string) ([]string, error) {
	names := []string{}
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return []string{}, fmt.Errorf("conn.Query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var enumName string
		err = rows.Scan(&enumName)
		if err != nil {
			return []string{}, fmt.Errorf("rows.Scan: %w", err)
		}
		names = append(names, searchPath+"."+enumName)
		names = append(names, searchPath+"."+"_"+enumName) // postgres internal array type automatically created
	}
	if err = rows.Err(); err != nil {
		return []string{}, fmt.Errorf("rows.Next: %w", err)
	}
	return names, err
}
