package app

import (
	"fmt"
	"net/http"

	"github.com/mtvy/qr-info/internal/service/qrcode"
	"github.com/mtvy/qr-info/internal/service/server"
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

	http.HandleFunc("/load/", server.InitHandler)

	http.ListenAndServe(host, nil)
}
