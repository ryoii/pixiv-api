// Package network implements a http client that doesn't provide SNI in https.
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

// Local dns table to avoid dns poisoning.
var localDns = map[string]*[]string{}

func NewClient() http.Client {
	var tlsCfg = &tls.Config{}
	tlsCfg.InsecureSkipVerify = true
	// hide SNI information
	tlsCfg.ServerName = ""

	var transport = &http.Transport{
		DialTLSContext: func(ctx context.Context, nw, addr string) (net.Conn, error) {
			// open a tcp connection via network.dial to avoid forced assignment of servername
			conn, err := dial(nw, lookup(addr), tlsCfg)
			return conn, err
		},
	}

	return http.Client{
		Transport: transport,
	}
}

// Transform [host:port] to [random_ip:port] via local dns table.
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

//　Init dns table
func init() {
	var lsApp []string
	localDns["oauth.secure.pixiv.net"] = &lsApp
	localDns["app-api.pixiv.net"] = &lsApp

	ips := []int{199, 219, 223, 226}

	for _, ip := range ips {
		lsApp = append(lsApp, "210.140.131."+strconv.FormatInt(int64(ip), 10))
	}
}
