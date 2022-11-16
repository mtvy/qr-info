package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/mtvy/qr-info/test/qrcode"
)

func main() {

	if err := godotenv.Load("./.env"); err != nil {
		return
	}

	qrcode.UnitTest(os.Getenv("HOST"))
}
