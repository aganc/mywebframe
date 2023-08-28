package helpers

import (
	"crypto/tls"
	"net/http"
	"time"
)

var globalTransport *http.Transport

type TransportOption struct {
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	IdleConnTimeout     time.Duration
	CustomTransport     *http.Transport
}

func InitHTTPClient() {
	opt := &TransportOption{
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     300 * time.Second,
	}
	InitHttp(opt)
}

// InitHttp 初始化全局的transport
func InitHttp(opts *TransportOption) {
	if opts == nil {
		globalTransport = &http.Transport{
			MaxIdleConns:        500,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     300 * time.Second,
		}
	} else if opts.CustomTransport != nil {
		globalTransport = opts.CustomTransport
	} else {
		globalTransport = &http.Transport{
			MaxIdleConns:        opts.MaxIdleConns,
			MaxIdleConnsPerHost: opts.MaxIdleConnsPerHost,
			IdleConnTimeout:     opts.IdleConnTimeout,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		}
	}
}
