package handlers

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

var tcp4client *http.Client

// TCP4 method
func TCP4(w http.ResponseWriter, r *http.Request) {

	resp := externalCall(getClient("tcp4"))
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, resp)
}

func getClient(networkType string) *http.Client {
	if tcp4client != nil {
		return tcp4client
	}
	dialer := &net.Dialer{}
	transport := &http.Transport{
		DialContext: func(ctx context.Context, network string, addr string) (net.Conn, error) {
			return dialer.DialContext(ctx, networkType, addr)
		},
	}
	tcp4client = &http.Client{
		Timeout:   time.Duration(2) * time.Second,
		Transport: transport,
	}
	return tcp4client
}
