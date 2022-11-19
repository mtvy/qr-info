package app

import (
	"fmt"
	"log"

	"github.com/mtvy/qr-info/internal/service/qrcode"
	"github.com/mtvy/qr-info/internal/service/server"
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

func InitQRCode(host string) qrcode.QRCode {
	qr := qrcode.QRCode{}

	qr.GenQRCodeBytes(host, "Mtvy")

	qr.GenQRCodeImg()

	qr.SaveQRCode()

	fmt.Println(qr.Code_id)

	return qr
}

func InitServer(host string) {

	log.Printf("\n%s├──>[%sINIT_SERVER%s][%sHOST"+host+"%s]\n│", PRPL, YLLW, PRPL, BL, PRPL)
	go server.StartHandlers(host)
	log.Printf("\n├──>%s[%s%sRUN_SERVER%s%s][%sHOST"+host+"%s]\n│", PRPL, FLSHNG, YLLW, STATIC, PRPL, GRN, PRPL)
}
