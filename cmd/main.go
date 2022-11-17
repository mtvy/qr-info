package main

import (
	"github.com/joho/godotenv"
	"github.com/mtvy/qr-info/internal/app"
)

func main() {

	if err := godotenv.Load("./.env"); err != nil {
		return
	}

	app.InitServer(":8080")
}
