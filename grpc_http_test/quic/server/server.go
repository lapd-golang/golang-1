package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	_ "net/http/pprof"

	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
)


type tracingHandler struct {
	handler http.Handler
}

var _ http.Handler = &tracingHandler{}

func (h *tracingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}

func setupHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
	//	fmt.Println(r.Form)
	//	fmt.Println(r.Form["numa"])
	//	fmt.Println(r.Form["numb"])
		numa, _:= strconv.Atoi(r.Form["numa"][0])
		numb, _:= strconv.Atoi(r.Form["numb"][0])
		fmt.Fprintf(w, "%d\n", numa+numb)
	})
	return &tracingHandler{handler: mux}
}

var (
	certFile = "cert/server.crt"
	keyFile = "cert/server.key"
)

func main() {

	addr := flag.String("addr", ":8088", "bind to")
	tcp := flag.Bool("tcp", false, "also listen on TCP")
	flag.Parse()

	handler := setupHandler()
	quicConf := &quic.Config{}

	var err error
	if *tcp {
		err = http3.ListenAndServe(*addr, certFile, keyFile, nil)
	} else {
		server := http3.Server{
			Server:     &http.Server{Handler: handler, Addr: *addr},
			QuicConfig: quicConf,
		}
		err = server.ListenAndServeTLS(certFile, keyFile)
	}
	if err != nil {
		fmt.Println(err)
	}
}
