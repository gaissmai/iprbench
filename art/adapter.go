package main_test

import (
	"net/netip"

	"github.com/tailscale/art"
)

func lpmWrapper(t *art.Table[any]) func(netip.Addr) bool {
	return func(ip netip.Addr) bool {
		_, ok := t.Get(ip)
		return ok
	}
}
