package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/mtvy/qr-info/internal/app"
)

func main() {

	if err := godotenv.Load("./.env"); err != nil {
		return
	}

	app.InitQRCode(os.Getenv("HOST"))
}
