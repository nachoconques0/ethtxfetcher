//go:build dev
// +build dev

package main

import (
	"os"
)

func init() {
	os.Setenv("HTTP_PORT", "8080")
	os.Setenv("NODE_ETH_URL", "https://ethereum-rpc.publicnode.com")
}
