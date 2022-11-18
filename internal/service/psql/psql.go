package psql

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	INSERT = "INSERT INTO qrcodes_tb (url, code_id, folder, name, path, initer, img_b) VALUES($1, $2, $3, $4, $5, $6, $7)"
	DELETE = "DELETE FROM qrcodes_tb WHERE initer=$1"
	GET    = "SELECT img_b FROM qrcodes_tb WHERE initer=$1"
)

func createCon() *pgxpool.Pool {

	dbUrl := os.Getenv("DB_URL")

	dbPool, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return dbPool
}

func dbReq(msg string, args ...any) ([]interface{}, error) {

	dbPool := createCon()

	rows, err := dbPool.Query(context.Background(), msg, args...)

	if rows.Next() {
		return rows.Values()
	}

	return nil, err
}

func Insert(args ...any) ([]interface{}, error) {
	return dbReq(INSERT, args...)
}

func Delete(args ...any) ([]interface{}, error) {
	return dbReq(DELETE, args...)
}

func Get(args ...any) ([]interface{}, error) {
	return dbReq(GET, args...)
}
