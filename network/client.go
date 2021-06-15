package network

import (
	"context"
	"crypto/tls"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"strings"
)

var localDns = map[string]*[]string{}

func NewClient() http.Client {
	var tlsCfg = &tls.Config{}
	tlsCfg.InsecureSkipVerify = true
	tlsCfg.ServerName = ""

	var transport = &http.Transport{
		DialTLSContext: func(ctx context.Context, nw, addr string) (net.Conn, error) {
			conn, err := dial(nw, lookup(addr), tlsCfg)
			return conn, err
		},
	}

	return http.Client{
		Transport: transport,
	}
}

func lookup(addr string) string {
	colonPos := strings.LastIndex(addr, ":")
	if colonPos == -1 {
		colonPos = len(addr)
	}
	return randElement(*localDns[addr[:colonPos]]) + addr[colonPos:]
}

func randElement(array []string) string {
	return array[rand.Intn(len(array))]
}

func init() {
	var lsApp []string
	localDns["oauth.secure.pixiv.net"] = &lsApp
	localDns["app-api.pixiv.net"] = &lsApp

	for i := 199; i <= 223; i++ {
		lsApp = append(lsApp, "210.140.131."+strconv.FormatInt(int64(i), 10))
	}
}
