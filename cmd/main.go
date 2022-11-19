package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/mtvy/qr-info/internal/app"
)

const (
	PRPL = "\033[35m"
	YLLW = "\033[33m"
	GRN  = "\033[32m"
	RD   = "\033[31m"
	BL   = "\033[36m"

	FLSHNG = "\033[5m"
	STATIC = "\033[25m"
)

func main() {

	fmt.Printf("\n%s[%sMAIN%s](ctrl+c to terminate)\n│\n", PRPL, YLLW, PRPL)

	if err := godotenv.Load("./.env"); err != nil {
		return
	}

	go app.InitServer(":8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Printf("\n%s├──>[%sSTOPED_SERVER%s][%sHOST:"+os.Getenv("HOST")+"%s]\n│", PRPL, YLLW, PRPL, BL, PRPL)
	log.Printf("\n%s└──>[%sEXIT%s]\n", PRPL, YLLW, PRPL)
}
