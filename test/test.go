package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/mtvy/qr-info/test/qrcode"
	"github.com/mtvy/qr-info/test/server"
)

func main() {

	if err := godotenv.Load("./.env"); err != nil {
		return
	}

	qrcode.UnitTest(os.Getenv("HOST"))

	server.UnitTest("http://localhost:8080/", ":8080")

}
