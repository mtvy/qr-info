package app

import (
	"fmt"
	"net/http"

	"github.com/mtvy/qr-info/internal/service/qrcode"
	"github.com/mtvy/qr-info/internal/service/server"
)

func InitQRCode(host string) qrcode.QRCode {
	qr := qrcode.QRCode{}

	qr.GenQRCodeBytes(host)

	qr.GenQRCodeImg()

	qr.SaveQRCode()

	fmt.Println(qr.Code_id)

	return qr
}

func InitServer(host string) bool {

	http.HandleFunc("/view/", server.MakeHandler(server.ViewHandler))
	http.HandleFunc("/edit/", server.MakeHandler(server.EditHandler))
	http.HandleFunc("/save/", server.MakeHandler(server.SaveHandler))
	http.HandleFunc("/cls/", server.ClsHandler)

	if http.ListenAndServe(host, nil) != nil {
		return false
	}

	return true
}
