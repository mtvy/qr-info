package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

const (
	PRPL = "\033[35m"
	YLLW = "\033[33m"
	GRN  = "\033[32m"
	RD   = "\033[31m"
	BL   = "\033[36m"

	FLSHNG = "\033[5m"
	STATIC = "\033[25m"
)

const (
	RESP_INIT_LOG = "\n%s├──>[%sRespQRCodeInit%s][%sINITER_%s%s][%sQR_%s%s][%sTRUE%s]\n│"
	RESP_SHOW_LOG = "\n%s├──>[%sRespQRCodeShow%s][%sINITER_%s%s][%sROWS_%d%s][%sTRUE%s]\n│"
	RESP_DEL_LOG  = "\n%s├──>[%sRespQRCodeDel%s][%sINITER_%s%s][%sSTATUS_%s%s][%sTRUE%s]\n│"

	F_RESP_INIT_LOG = "\n%s├──>[%sRespQRCodeInit%s][%sFALSE%s]\n│"
	F_RESP_SHOW_LOG = "\n%s├──>[%sRespQRCodeShow%s][%sFALSE%s]\n│"
	F_RESP_DEL_LOG  = "\n%s├──>[%sRespQRCodeDel%s][%sFALSE%s]\n│"

	E_RESP_INIT_MSG = "Request should contain 'url' and 'initer'"
	E_RESP_SHOW_MSG = "Request should contain 'initer'"
	E_RESP_DEL_MSG  = "Request should contain 'initer'"
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

func respQRCodeJson(w http.ResponseWriter, r *http.Request, qr_resp any) {
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

		respQRCodeJson(w, r, RespQRCodeInit{
			Code_id: qr.Code_id,
			Initer:  qr.Initer,
		})

		log.Printf(RESP_INIT_LOG, PRPL, YLLW, PRPL, BL, qr.Initer, PRPL, BL, qr.Code_id, PRPL, GRN, PRPL)

	} else {
		respQRCodeJson(w, r, RespQRCodeInit{Err: E_RESP_INIT_MSG})
		log.Printf(F_RESP_INIT_LOG, PRPL, YLLW, PRPL, RD, PRPL)
	}
}

func ShowHandler(w http.ResponseWriter, r *http.Request) {

	initer := r.URL.Query()["initer"]

	if len(initer) > 0 && len(initer[0]) > INITER_LEN {
		rows, err := psql.Get(initer[0])
		if err != nil {
			respQRCodeJson(w, r, RespQRCodeShow{
				Initer: initer[0],
				Err:    err.Error(),
			})
			log.Printf(F_RESP_SHOW_LOG, PRPL, YLLW, PRPL, RD, PRPL)
		} else {
			respQRCodeJson(w, r, RespQRCodeShow{
				Initer: initer[0],
				Img_b:  rows,
			})
			log.Printf(RESP_SHOW_LOG, PRPL, YLLW, PRPL, BL, initer[0], PRPL, BL, len(rows), PRPL, GRN, PRPL)
		}

	} else {
		respQRCodeJson(w, r, RespQRCodeShow{Err: E_RESP_SHOW_MSG})
		log.Printf(F_RESP_SHOW_LOG, PRPL, YLLW, PRPL, RD, PRPL)
	}
}

func DelHandler(w http.ResponseWriter, r *http.Request) {

	initer := r.URL.Query()["initer"]

	if len(initer) > 0 && len(initer[0]) > INITER_LEN {
		_, err := psql.Delete(initer[0])
		if err != nil {
			respQRCodeJson(w, r, RespQRCodeDel{
				Initer: initer[0],
				Status: "fault",
				Err:    err.Error(),
			})
			log.Printf(F_RESP_DEL_LOG, PRPL, YLLW, PRPL, RD, PRPL)
		} else {
			respQRCodeJson(w, r, RespQRCodeDel{
				Initer: initer[0],
				Status: "deleted",
			})
			log.Printf(RESP_DEL_LOG, PRPL, YLLW, PRPL, BL, initer[0], PRPL, BL, "deleted", PRPL, GRN, PRPL)
		}

	} else {
		respQRCodeJson(w, r, RespQRCodeDel{Err: E_RESP_DEL_MSG})
		log.Printf(F_RESP_DEL_LOG, PRPL, YLLW, PRPL, RD, PRPL)
	}
}

func StartHandlers(host string) {
	http.HandleFunc("/init", InitHandler)
	http.HandleFunc("/show", ShowHandler)
	http.HandleFunc("/del", DelHandler)

	http.ListenAndServe(host, nil)
}

func ClsHandler() {
	os.Exit(1)
}
