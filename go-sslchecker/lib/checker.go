package lib

import (
	"crypto/tls"
	"fmt"
	"time"
)

type SSLResponse struct {
	Issuer string `json:"issuer"`
	Expiry string `json:"expiry"`
}

func SSLCheck(addr string) (*SSLResponse, error) {
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:443", addr), nil)
	if err != nil {
		return nil, fmt.Errorf("Server doesn't support SSL certificate err: %s", err.Error())
	}

	err = conn.VerifyHostname(addr)
	if err != nil {
		return nil, fmt.Errorf("Hostname doesn't match with certificate: %s", err.Error())
	}

	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter

	return &SSLResponse{
		Issuer: conn.ConnectionState().PeerCertificates[0].Issuer.String(),
		Expiry: expiry.Format(time.RFC850),
	}, nil
}
