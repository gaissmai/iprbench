package main_test

import (
	"net"

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

func (t *Table) Insert(pfx net.IPNet, val any) {
	if common.IPis4(pfx.IP) {
		if err := t.v4.Add(&pfx, val); err != nil {
			panic(err)
		}
		return
	}
	if err := t.v6.Add(&pfx, val); err != nil {
		panic(err)
	}
}

func (t *Table) Delete(pfx net.IPNet) {
	if common.IPis4(pfx.IP) {
		if _, _, err := t.v4.Delete(&pfx); err != nil {
			panic(err)
		}
		return
	}
	if _, _, err := t.v6.Delete(&pfx); err != nil {
		panic(err)
	}
}

func (t *Table) Contains(ip net.IP) bool {
	if common.IPis4(ip) {
		ok, _ := t.v4.ContainedIP(ip)
		return ok
	}

	ok, _ := t.v6.ContainedIP(ip)
	return ok
}
