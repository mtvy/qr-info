package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mtvy/qr-info/internal/service/psql"
	"github.com/mtvy/qr-info/internal/service/qrcode"
)

type RespQRCodeInit struct {
	Code_id string `json:"code_id"`
	Initer  string `json:"initer"`
	Err     string `json:"error"`
}

type RespQRCodeShow struct {
	Initer string        `json:"initer"`
	Img_b  []interface{} `json:"img_b"`
	Err    string        `json:"error"`
}

type RespQRCodeDel struct {
	Initer string `json:"initer"`
	Status string `json:"status"`
	Err    string `json:"error"`
}

const (
	MIN_PARAMS_LEN int = 0
	INITER_LEN         = 1
	URL_LEN            = 3
)

func MakeRequest(req_url string) string {
	resp, err := http.Get(req_url)

	if err != nil {
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(body)
}

func RespQRCodeJson(w http.ResponseWriter, r *http.Request, qr_resp any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(qr_resp)
}

func InitHandler(w http.ResponseWriter, r *http.Request) {

	qr := qrcode.QRCode{}
	url := r.URL.Query()["url"]
	initer := r.URL.Query()["initer"]

	if len(initer)*len(url) > MIN_PARAMS_LEN && len(initer[0]) > INITER_LEN && len(url[0]) > URL_LEN {

		qr.GenQRCodeBytes(url[0], initer[0])

		qr.SaveQRCode()

		RespQRCodeJson(w, r, RespQRCodeInit{
			Code_id: qr.Code_id,
			Initer:  qr.Initer,
		})

	} else {
		RespQRCodeJson(w, r, RespQRCodeInit{
			Err: "Request should contain 'url' and 'initer'",
		})
	}
}

func ShowHandler(w http.ResponseWriter, r *http.Request) {

	initer := r.URL.Query()["initer"]

	if len(initer) > 0 && len(initer[0]) > INITER_LEN {
		rows, err := psql.Get(initer[0])
		if err != nil {
			RespQRCodeJson(w, r, RespQRCodeShow{
				Initer: initer[0],
				Err:    err.Error(),
			})
		} else {
			RespQRCodeJson(w, r, RespQRCodeShow{
				Initer: initer[0],
				Img_b:  rows,
			})
		}

	} else {
		RespQRCodeJson(w, r, RespQRCodeShow{
			Err: "Request should contain 'initer'",
		})
	}
}

func DelHandler(w http.ResponseWriter, r *http.Request) {
	initer := r.URL.Query()["initer"]

	if len(initer) > 0 && len(initer[0]) > INITER_LEN {
		_, err := psql.Delete(initer[0])
		if err != nil {
			RespQRCodeJson(w, r, RespQRCodeDel{
				Initer: initer[0],
				Status: "fault",
				Err:    err.Error(),
			})
		} else {
			RespQRCodeJson(w, r, RespQRCodeDel{
				Initer: initer[0],
				Status: "deleted",
			})
		}

	} else {
		RespQRCodeJson(w, r, RespQRCodeDel{
			Err: "Request should contain 'initer'",
		})
	}
}

func InitHandlers(host string) {

	http.HandleFunc("/init", InitHandler)
	http.HandleFunc("/show", ShowHandler)
	http.HandleFunc("/del", DelHandler)

	http.ListenAndServe(host, nil)
}

func ClsHandler() {

	os.Exit(1)
}
