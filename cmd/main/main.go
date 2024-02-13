package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/pliavi/gocrast/internal"
	"github.com/pliavi/gocrast/internal/utils"
	"gopkg.in/elazarl/goproxy.v1"
)

func main() {
	configPath := flag.String("file", "", "path to site blocking config file")
	port := flag.String("port", "62222", "port to run the proxy server")
	flag.Parse()

	ips, err := utils.GetLocalIPs()
	if err != nil {
		log.Fatal(err)
	}

	siteBlocker := internal.CreateSiteBlockerConfig(*configPath)

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false

	proxy.OnRequest().HandleConnectFunc(func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
		url := ctx.Req.URL.Hostname()
		fmt.Println("url:", url)

		if siteBlocker.IsBlocked(url) {
			return goproxy.RejectConnect, host
		}

		return goproxy.OkConnect, host
	})

	log.Println("Starting proxy server on port", *port)
	log.Println("Config file:", *configPath)
	log.Println("Blocking sites:", siteBlocker.Everyday.Sites)
	log.Println("Blocking from:", siteBlocker.Everyday.BlockStart)
	log.Println("Blocking to:", siteBlocker.Everyday.BlockEnd)
	log.Println("-------------------------------------------------------")
	log.Println("You can use this proxy server with:")
	log.Println(" - address:", ips[0])
	log.Println(" - port:", *port)
	log.Println("Press Ctrl+C to stop")
	log.Fatalln(http.ListenAndServe(ips[0].String()+":"+*port, proxy))
}
