package main

import (
	"flag"
	"log"

	"github.com/pliavi/gocrast/internal/proxy"
	"github.com/pliavi/gocrast/pkg/network"
)

func main() {
	configPath := flag.String("config", "", "path to site blocking config file")
	port := flag.String("port", "62222", "port to run the proxy server")
	flag.Parse()

	ips, err := network.GetLocalIPs()
	if err != nil {
		log.Fatal(err)
	}

	proxy := proxy.NewProxy(ips[0].String(), *port).Setup(*configPath)

	log.Println("Starting proxy server on port", *port)
	log.Println("Config file:", *configPath)
	log.Println("Blocking sites:", proxy.Config.Everyday.Sites)
	log.Println("Blocking from:", proxy.Config.Everyday.BlockStart)
	log.Println("Blocking to:", proxy.Config.Everyday.BlockEnd)
	log.Println("-------------------------------------------------------")
	log.Println("You can use this proxy server with:")
	log.Println(" - address:", ips[0])
	log.Println(" - port:", *port)
	log.Println("Press Ctrl+C to stop")
	log.Fatalln(proxy.Start())
}
