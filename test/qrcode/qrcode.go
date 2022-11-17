package qrcode

import (
	"log"

	"github.com/mtvy/qr-info/internal/service/qrcode"
)

const (
	PRPL = "\033[35m"
	YLLW = "\033[33m"
	GRN  = "\033[32m"
	RD   = "\033[31m"
)

func UnitTest(host string) {

	log.Printf("\n%s[%sQRCODE_MAIN_FUNCS%s]\n│", PRPL, YLLW, PRPL)

	qr := qrcode.QRCode{}
	log.Printf("\n%s├[%sINIT_QRCODE%s]\n│", PRPL, YLLW, PRPL)

	val := func(expr bool) string {
		if expr {
			return "\033[32mTRUE\033[35m"
		} else {
			return "\033[31mFALSE\033[35m"
		}
	}

	log.Printf("\n%s├[%sGEN_QRCODE_V1%s]["+qr.Code_id+
		"][%sIMG_BYTES%s]["+val(qr.GenQRCodeBytes(host, "Mtvy"))+"]\n│", PRPL, YLLW, PRPL, YLLW, PRPL)

	log.Printf("\n%s├[%sGEN_QRCODE_V2%s]["+qr.Path+
		"][%sIMG_PNG%s]["+val(qr.GenQRCodeImg())+"]\n│", PRPL, YLLW, PRPL, YLLW, PRPL)

	res := func(expr []interface{}, err error) string {
		if err == nil {
			return "\033[32mTRUE\033[35m"
		} else {
			return "\033[31mFALSE\033[35m"
		}
	}

	log.Printf("\n%s├[%sSAVE_QRCODE%s]["+host+
		"][postgres]["+res(qr.SaveQRCode())+"]\n│", PRPL, YLLW, PRPL)

	log.Printf("\n%s├[%sGET_QRCODE%s]["+host+
		"][postgres]["+res(qr.GetQRCode())+"]\n│", PRPL, YLLW, PRPL)

	log.Printf("\n%s└[%sREMOVE_QRCODE%s]["+qr.Code_id+
		"]["+res(qr.RmvQRCode())+"]\n\n\n", PRPL, YLLW, PRPL)
}
