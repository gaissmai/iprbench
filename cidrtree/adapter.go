package main_test

import (
	"net/netip"

	"github.com/gaissmai/cidrtree"
)

func lpmWrapper(t *cidrtree.Table[any]) func(netip.Addr) bool {
	return func(ip netip.Addr) bool {
		_, _, ok := t.Lookup(ip)
		return ok
	}
}
