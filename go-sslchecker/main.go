package main

import (
	"fmt"
	"github.com/zeindevs/sslchecker/lib"
)

func main() {
	expiry, err := lib.SSLCheck("www.zeindevs.com")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Issuer: %s\nExpiry: %v\n", expiry.Issuer, expiry.Expiry)
}

