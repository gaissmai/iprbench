package main_test

import (
	"net/netip"

	"github.com/kentik/patricia"
	kptree "github.com/kentik/patricia/generics_tree"
)

// adapter
type Table[V any] struct {
	v4 *kptree.TreeV4[V]
	v6 *kptree.TreeV6[V]
}

func NewTable[V any]() *Table[V] {
	kp4 := kptree.NewTreeV4[V]()
	kp6 := kptree.NewTreeV6[V]()

	return &Table[V]{
		v4: kp4,
		v6: kp6,
	}
}

func (t *Table[V]) Insert(pfx netip.Prefix, val V) {
	ip := pfx.Addr()

	if ip.Is4() {
		addr := patricia.NewIPv4AddressFromBytes(ip.AsSlice(), uint(pfx.Bits()))
		t.v4.Set(addr, val)
		return
	}
	addr := patricia.NewIPv6Address(ip.AsSlice(), uint(pfx.Bits()))
	t.v6.Set(addr, val)
	return
}

func (t *Table[V]) Delete(pfx netip.Prefix) {
	var zero V
	matchFunc := func(V, V) bool { return true } // always true

	ip := pfx.Addr()

	if ip.Is4() {
		addr := patricia.NewIPv4AddressFromBytes(ip.AsSlice(), uint(pfx.Bits()))
		t.v4.Delete(addr, matchFunc, zero)
		return
	}
	addr := patricia.NewIPv6Address(ip.AsSlice(), uint(pfx.Bits()))
	t.v6.Delete(addr, matchFunc, zero)
	return
}

func (t *Table[V]) Lookup(ip netip.Addr) (val V, ok bool) {
	if ip.Is4() {
		addr := patricia.NewIPv4AddressFromBytes(ip.AsSlice(), 32)
		ok, val = t.v4.FindDeepestTag(addr)
		return
	}
	addr := patricia.NewIPv6Address(ip.AsSlice(), 128)

	// switch order
	ok, val = t.v6.FindDeepestTag(addr)
	return
}
