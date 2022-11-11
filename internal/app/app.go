package app

import (
	"github.com/mtvy/qr-info/internal/service"
)

func GetQrMetaInfo() string {
	return "[META_DATA]"
}

func InitQRCode() service.QRCode {
	qr := service.QRCode{}

	qr.GenQRCodeBytes()

	qr.GenQRCodeImg()

	return qr
}
