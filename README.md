# dial-poc
POC for custom go http transport with tcp4 only dns resolution 

Service runs on port `8998`

We have 2 APIs with different setup of http client

## `/tcp`

It uses default http transport and client implementation, default network type is "tcp" hence creating 2 dns queries A(ipv4) and AAAA(ipv6)


Client Implementation

```
dialer := &net.Dialer{}
transport := &http.Transport{
  DialContext: func(ctx context.Context, network string, addr string) (net.Conn, error) {
    return dialer.DialContext(ctx, network, addr)
  },
}
tcp4client = &http.Client{
  Timeout:   time.Duration(2) * time.Second,
  Transport: transport,
	}
```


Curl Command: `curl http://localhost:8998/tcp`


## `/tcp4`

Customized the transport implementation to use "tcp4", hence creating single dns queries A(ipv4)


```
dialer := &net.Dialer{}
transport := &http.Transport{
  DialContext: func(ctx context.Context, network string, addr string) (net.Conn, error) {
    return dialer.DialContext(ctx, "tcp4", addr)
  },
}
tcp4client = &http.Client{
  Timeout:   time.Duration(2) * time.Second,
  Transport: transport,
}
```

Curl Command: `curl http://localhost:8998/tcp4`
