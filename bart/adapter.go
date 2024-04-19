package main_test

import (
	"net/netip"

	"github.com/gaissmai/bart"
)

func lpmWrapper(t *bart.Table[any]) func(netip.Addr) bool {
	return func(ip netip.Addr) bool {
		_, ok := t.Lookup(ip)
		return ok
	}
}
