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

	app.InitServer(":8080")

	app.InitQRCode(os.Getenv("HOST"))

}
