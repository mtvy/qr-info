package app

import (
	"github.com/mtvy/qr-info/internal/service"
)

func InitQRCode(host string) service.QRCode {
	qr := service.QRCode{}

	qr.GenQRCodeBytes(host)

	qr.GenQRCodeImg()

	return qr
}
