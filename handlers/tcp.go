package handlers

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

var tcpclient *http.Client

// TCP method
func TCP(w http.ResponseWriter, r *http.Request) {
	resp := externalCall(getTCPClient("tcp"))
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, resp)
}

func getTCPClient(networkType string) *http.Client {
	if tcpclient != nil {
		return tcpclient
	}
	dialer := &net.Dialer{}
	transport := &http.Transport{
		DialContext: func(ctx context.Context, network string, addr string) (net.Conn, error) {
			return dialer.DialContext(ctx, networkType, addr)
		},
	}
	tcpclient = &http.Client{
		Timeout:   time.Duration(2) * time.Second,
		Transport: transport,
	}
	return tcpclient
}
