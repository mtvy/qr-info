package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mtvy/qr-info/internal/service/server"
)

const (
	PRPL = "\033[35m"
	YLLW = "\033[33m"
	GRN  = "\033[32m"
	RD   = "\033[31m"
)

const (
	INIT1 = "init?"
	INIT2 = "init?initer=Mtvy"
	INIT3 = "init?initer=Mtvy&url=rutubeto.ru"
	SHOW1 = "show?"
	SHOW2 = "show?initer=Mtvy"
	DEL1  = "del?"
	DEL2  = "del?initer=Mtvy"
)

func UnitTest(url string, host string) {

	log.Printf("\n%s[%sSERVER_MAIN_REQS%s]\n│", PRPL, YLLW, PRPL)

	res := func(sc string) string {
		if sc != "" {

			rs := sc[:len(sc)-1]
			if len(sc) > 128 {
				rs = sc[:64] + "..."
			}

			return "├──>\033[32m" + rs + "\033[35m"

		} else {
			return "├──>\033[31mFALSE\033[35m"
		}
	}

	log.Printf("\n%s├[%sMAKE_REQ%s]["+url+"]\n"+res(server.MakeRequest(url))+"\n│", PRPL, YLLW, PRPL)
	log.Printf("\n%s├[%sINIT_HANDLERS%s]\n│", PRPL, YLLW, PRPL)

	go server.InitHandlers(host)

	log.Printf("\n%s├[%sMAKE_REQ%s]["+url+INIT1+"]\n"+res(server.MakeRequest(url+INIT1))+"\n│", PRPL, YLLW, PRPL)
	log.Printf("\n%s├[%sMAKE_REQ%s]["+url+INIT2+"]\n"+res(server.MakeRequest(url+INIT2))+"\n│", PRPL, YLLW, PRPL)
	log.Printf("\n%s├[%sMAKE_REQ%s]["+url+INIT3+"]\n"+res(server.MakeRequest(url+INIT3))+"\n│", PRPL, YLLW, PRPL)
	log.Printf("\n%s├[%sMAKE_REQ%s]["+url+SHOW1+"]\n"+res(server.MakeRequest(url+SHOW1))+"\n│", PRPL, YLLW, PRPL)
	log.Printf("\n%s├[%sMAKE_REQ%s]["+url+SHOW2+"]\n"+res(server.MakeRequest(url+SHOW2))+"\n│", PRPL, YLLW, PRPL)
	log.Printf("\n%s├[%sMAKE_REQ%s]["+url+DEL1+"]\n"+res(server.MakeRequest(url+DEL1))+"\n│", PRPL, YLLW, PRPL)
	log.Printf("\n%s├[%sMAKE_REQ%s]["+url+DEL2+"]\n"+res(server.MakeRequest(url+DEL2))+"\n│", PRPL, YLLW, PRPL)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	server.ClsHandler()

}
