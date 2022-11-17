package qrcode

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"image"
	"image/png"
	"os"
	"runtime"

	"github.com/mtvy/qr-info/internal/service/psql"
	"rsc.io/qr"
)

type stream_t func(...any) (int, error)

type MetaData struct {
	Url     string
	Code_id string
}

type FileData struct {
	Folder string
	Name   string
	Path   string
	Img_b  []byte
}

type QStatus struct {
	Status bool
}

type ProcQRCode interface {
	GenQRCodeBytes()
	GenQRCodeImg()
	IsValid()
	SaveQRCode()
	RmvQRCode()
}

type QRCode struct {
	MetaData
	FileData
	QStatus
}

const (
	SYMBS = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	WORDS = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	NUMS  = "0123456789"
)

func randStr(strSize int, dict string) string {

	bytes := make([]byte, strSize)

	rand.Read(bytes)

	for ind, num := range bytes {
		bytes[ind] = dict[num%byte(len(dict))]
	}

	return string(bytes)
}

func (qrcode *QRCode) IsValid() bool {

	// TODO: conf validation

	if qrcode.Status && qrcode.Img_b == nil {
		qrcode.Url = ""
		qrcode.Code_id = ""
		qrcode.Status = false
	}
	return qrcode.Status
}

func (qrcode *QRCode) clearQRcode(err error, out stream_t) bool {

	out(err)

	qrcode.Code_id = ""
	qrcode.Folder = ""
	qrcode.Name = ""
	qrcode.Path = ""
	qrcode.Url = ""
	qrcode.Img_b = nil
	qrcode.Status = false

	return qrcode.Status
}

func (qrcode *QRCode) GenQRCodeImg() bool {

	if !qrcode.IsValid() {
		return qrcode.Status
	}

	img, _, err := image.Decode(bytes.NewReader(qrcode.Img_b))
	if err != nil {
		return qrcode.clearQRcode(err, fmt.Println)
	}

	qrcode.Folder = "assets/qrcodes/" + qrcode.Code_id

	if err = os.Mkdir(qrcode.Folder, 0777); err != nil {
		return qrcode.clearQRcode(err, fmt.Println)
	}

	qrcode.Name = qrcode.Code_id + ".png"
	qrcode.Path = qrcode.Folder + "/" + qrcode.Name

	out, err := os.Create(qrcode.Path)
	if err != nil {
		return qrcode.clearQRcode(err, fmt.Println)
	}

	if err = png.Encode(out, img); err != nil {
		return qrcode.clearQRcode(err, fmt.Println)
	}

	return qrcode.Status
}

func (qrcode *QRCode) GenQRCodeBytes(host string) bool {

	runtime.GOMAXPROCS(runtime.NumCPU())

	qrcode.Code_id = randStr(8, SYMBS)

	qrcode.Url = host + "/?code=" + qrcode.Code_id

	code, err := qr.Encode(qrcode.Url, qr.H)
	if err != nil {
		fmt.Println(err)
		return qrcode.Status
	}

	qrcode.Img_b = code.PNG()

	qrcode.Status = true
	return qrcode.Status
}

func (qrcode *QRCode) SaveQRCode() bool {
	return psql.Insert(qrcode.Url, qrcode.Code_id, qrcode.Folder, qrcode.Name, qrcode.Path, qrcode.Img_b)
}

func (qrcode *QRCode) RmvQRCode() bool {
	if os.RemoveAll("./assets/qrcodes/"+qrcode.Code_id) != nil {
		return false
	}
	return psql.Delete(qrcode.Code_id)
}
