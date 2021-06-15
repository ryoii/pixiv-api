package network

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
)

func NewClient() http.Client {
	var tlsCfg = &tls.Config{}
	tlsCfg.InsecureSkipVerify = true
	tlsCfg.ServerName = ""

	var transport = &http.Transport{
		DialTLSContext: func(ctx context.Context, nw, addr string) (net.Conn, error) {
			conn, err := dial(nw, addr, tlsCfg)
			return conn, err
		},
	}

	return http.Client{
		Transport: transport,
	}
}
