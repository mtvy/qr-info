package service

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"image"
	"image/png"
	"os"
	"runtime"

	"rsc.io/qr"
)

type status_t bool

type MetaData struct {
	url     string
	code_id string
}

type FileData struct {
	folder string
	name   string
	path   string
	img_b  []byte
}

type QStatus struct {
	status status_t
}

type ProcQRCode interface {
	GenQRCode()
	GenQRCodeImg()
	IsValid()
}

type QRCode struct {
	MetaData
	FileData
	QStatus
}

type Vertex struct {
	Lat, Long float64
}

const (
	SYMBS = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	WORDS = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	NUMS  = "0123456789"
)

const (
	HOST = "rutubeto.ru"
)

func randStr(strSize int, dict string) string {

	bytes := make([]byte, strSize)

	rand.Read(bytes)

	for ind, num := range bytes {
		bytes[ind] = dict[num%byte(len(dict))]
	}

	return string(bytes)
}

func (qrcode *QRCode) IsValid() status_t {

	// TODO: conf validation

	if qrcode.status && qrcode.img_b == nil {
		qrcode.url = ""
		qrcode.code_id = ""
		qrcode.status = false
	}
	return qrcode.status
}

type stream_t func(...any) (int, error)

func (qrcode *QRCode) clearQRcode(err error, out stream_t) status_t {

	out(err)

	qrcode.code_id = ""
	qrcode.folder = ""
	qrcode.name = ""
	qrcode.path = ""
	qrcode.url = ""
	qrcode.img_b = nil
	qrcode.status = false

	return qrcode.status
}

func (qrcode *QRCode) GenQRCodeImg() status_t {

	if !qrcode.IsValid() {
		return qrcode.status
	}

	img, _, err := image.Decode(bytes.NewReader(qrcode.img_b))
	if err != nil {
		return qrcode.clearQRcode(err, fmt.Println)
	}

	qrcode.folder = "qrcodes/"
	qrcode.name = qrcode.code_id + ".png"
	qrcode.path = qrcode.folder + qrcode.name

	out, err := os.Create(qrcode.path)
	if err != nil {
		return qrcode.clearQRcode(err, fmt.Println)
	}

	err = png.Encode(out, img)
	if err != nil {
		return qrcode.clearQRcode(err, fmt.Println)
	}

	return qrcode.status
}

func (qrcode *QRCode) GenQRCode() status_t {

	runtime.GOMAXPROCS(runtime.NumCPU())

	qrcode.code_id = randStr(8, SYMBS)

	qrcode.url = HOST + "/?code=" + qrcode.code_id

	code, err := qr.Encode(qrcode.url, qr.H)
	if err != nil {
		fmt.Println(err)
		return qrcode.status
	}

	qrcode.img_b = code.PNG()

	qrcode.status = true
	return qrcode.status
}
