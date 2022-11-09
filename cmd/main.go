package main

import (
	"fmt"

	"github.com/mtvy/qr-info/internal/app"
)

func main() {
	qr_meta_data := app.GetQrMetaInfo()
	fmt.Println(qr_meta_data)
	app.InitQRCode()
}
