package psql

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	INSERT = "INSERT INTO qrcodes_tb (url, code_id, folder, name, path, img_b) VALUES($1, $2, $3, $4, $5, $6)"
	DELETE = "DELETE FROM qrcodes_tb WHERE code_id =$1"
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

func dbReq(msg string, args ...any) bool {

	dbPool := createCon()

	_, err := dbPool.Query(context.Background(), msg, args...)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func Insert(args ...any) bool {
	return dbReq(INSERT, args...)
}

func Delete(args ...any) bool {
	return dbReq(DELETE, args...)
}
