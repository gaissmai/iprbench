package main_test

import (
	"net"

	"local/iprbench/common"

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

func (t *Table) Insert(p net.IPNet, val any) {
	ip := p.IP.To4()
	if ip == nil {
		ip = p.IP.To16()
	}

	ones, _ := p.Mask.Size()

	key := lpmtrie.Key{
		PrefixLen: ones,
		Data:      ip,
	}

	if common.IPis4(ip) {
		t.v4.Update(key, val)
		return
	}
	t.v6.Update(key, val)
}

func (t *Table) Delete(p net.IPNet) {
	ip := p.IP.To4()
	if ip == nil {
		ip = p.IP.To16()
	}

	ones, _ := p.Mask.Size()

	key := lpmtrie.Key{
		PrefixLen: ones,
		Data:      ip,
	}

	if common.IPis4(ip) {
		t.v4.Delete(key)
		return
	}
	t.v6.Delete(key)
}

func (t *Table) Lookup(ip net.IP) (val any, ok bool) {
	bitLen := 32
	ipv := ip.To4()
	if ipv == nil {
		ipv = ip.To16()
		bitLen = 128
	}

	key := lpmtrie.Key{
		PrefixLen: bitLen,
		Data:      ipv,
	}

	if common.IPis4(ip) {
		return t.v4.Lookup(key)
	}
	return t.v6.Lookup(key)
}
