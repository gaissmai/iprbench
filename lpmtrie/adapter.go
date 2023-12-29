package main_test

import (
	"net/netip"

	"github.com/Asphaltt/lpmtrie"
)

// adapter
type Table struct {
	v4 lpmtrie.LpmTrie
	v6 lpmtrie.LpmTrie
}

func NewTable() *Table {
	lt4, _ := lpmtrie.New(32)
	lt6, _ := lpmtrie.New(128)

	return &Table{
		v4: lt4,
		v6: lt6,
	}
}

func (t *Table) Insert(p netip.Prefix, val any) {
	ip := p.Addr()
	bits := p.Bits()

	key := lpmtrie.Key{
		PrefixLen: bits,
		Data:      ip.AsSlice(),
	}

	if ip.Is4() {
		t.v4.Update(key, val)
		return
	}
	t.v6.Update(key, val)
}

func (t *Table) Delete(p netip.Prefix) {
	ip := p.Addr()
	bits := p.Bits()

	key := lpmtrie.Key{
		PrefixLen: bits,
		Data:      ip.AsSlice(),
	}

	if ip.Is4() {
		t.v4.Delete(key)
		return
	}
	t.v6.Delete(key)
}

func (t *Table) Get(ip netip.Addr) (val any, ok bool) {
	key := lpmtrie.Key{
		PrefixLen: ip.BitLen(),
		Data:      ip.AsSlice(),
	}

	if ip.Is4() {
		return t.v4.Lookup(key)
	}
	return t.v6.Lookup(key)
}

func lpmWrapper(t *Table) func(netip.Addr) bool {
	return func(ip netip.Addr) bool {
		_, ok := t.Get(ip)
		return ok
	}
}
