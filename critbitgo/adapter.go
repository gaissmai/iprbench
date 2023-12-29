package main_test

import (
	"net/netip"

	"local/iprbench/common"

	"github.com/k-sone/critbitgo"
)

// adapter
type Table struct {
	v4 *critbitgo.Net
	v6 *critbitgo.Net
}

func NewTable() *Table {
	return &Table{
		v4: critbitgo.NewNet(),
		v6: critbitgo.NewNet(),
	}
}

func (t *Table) Insert(p netip.Prefix, val any) {
	net := common.PfxToIPNet(p)
	if p.Addr().Is4() {
		if err := t.v4.Add(&net, val); err != nil {
			panic(err)
		}
		return
	}
	if err := t.v6.Add(&net, val); err != nil {
		panic(err)
	}
}

func (t *Table) Delete(p netip.Prefix) {
	net := common.PfxToIPNet(p)
	if p.Addr().Is4() {
		if _, _, err := t.v4.Delete(&net); err != nil {
			panic(err)
		}
		return
	}
	if _, _, err := t.v6.Delete(&net); err != nil {
		panic(err)
	}
}

func (t *Table) Get(ip netip.Addr) (val any, ok bool) {
	if ip.Is4() {
		route, val, _ := t.v4.MatchIP(common.AddrToIP(ip))
		if route != nil {
			return val, true
		}
		return val, false
	}

	route, val, _ := t.v6.MatchIP(common.AddrToIP(ip))
	if route != nil {
		return val, true
	}
	return val, false
}

func lpmWrapper(t *Table) func(netip.Addr) bool {
	return func(ip netip.Addr) bool {
		_, ok := t.Get(ip)
		return ok
	}
}
