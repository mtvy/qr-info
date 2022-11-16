package app

import (
	"fmt"

	"github.com/mtvy/qr-info/internal/qrcode"
)

func InitQRCode(host string) qrcode.QRCode {
	qr := qrcode.QRCode{}

	qr.GenQRCodeBytes(host)

	qr.GenQRCodeImg()

	qr.SaveQRCode()

	fmt.Println(qr.Code_id)

	qr.RmvQRCode()

	return qr
}
