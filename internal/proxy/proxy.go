package proxy

import (
	"fmt"
	"net/http"

	"github.com/pliavi/gocrast/internal/site_blocker_config"
	"gopkg.in/elazarl/goproxy.v1"
)

type Proxy struct {
	Config *site_blocker_config.SiteBlockerConfig
	server *goproxy.ProxyHttpServer
	Host   string
	Port   string
}

func NewProxy(host string, port string) *Proxy {
	server := goproxy.NewProxyHttpServer()
	server.Verbose = false

	return &Proxy{
		server: server,
		Host:   host,
		Port:   port,
	}
}

func (p *Proxy) Setup(pathToBlockingConfig string) *Proxy {
	return p.setupBlockingConfigs(pathToBlockingConfig)
}

func (p *Proxy) setupBlockingConfigs(pathToBlockingConfig string) *Proxy {
	p.Config = site_blocker_config.NewSiteBlockerConfig(pathToBlockingConfig)

	return p
}

func (p *Proxy) setupBlockingHandlers() *Proxy {
	p.server.OnRequest().HandleConnectFunc(func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
		url := ctx.Req.URL.Hostname()
		fmt.Println("url:", url)

		if p.Config.IsBlocked(url) {
			return goproxy.RejectConnect, host
		}

		return goproxy.OkConnect, host
	})

	return p
}

func (p *Proxy) Start() error {
	p.setupBlockingHandlers()
	return http.ListenAndServe(p.Host+":"+p.Port, p.server)
}
