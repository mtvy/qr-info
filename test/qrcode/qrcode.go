package qrcode

import (
	"fmt"
	"strconv"

	"github.com/mtvy/qr-info/internal/qrcode"
)

func UnitTest(host string) {

	qr := qrcode.QRCode{}
	fmt.Println("[INIT_QRCODE]")

	fmt.Println("[GEN_QRCODE_V1][" + qr.Code_id + "][IMG_BYTES][" + strconv.FormatBool(qr.GenQRCodeBytes(host)) + "]")

	fmt.Println("[GEN_QRCODE_V2][" + qr.Path + "][IMG_PNG][" + strconv.FormatBool(qr.GenQRCodeImg()) + "]")

	fmt.Println("[SAVE_QRCODE][" + host + "][postgres][" + strconv.FormatBool(qr.SaveQRCode()) + "]")

	fmt.Println("[REMOVE_QRCODE][" + qr.Code_id + "][" + strconv.FormatBool(qr.RmvQRCode()) + "]")
}
