package main_test

import (
	"net/netip"
	"testing"

	"local/iprbench/common"

	"github.com/aromatt/netipds"
)

func TestMatchIP(t *testing.T) {
	t.Parallel()
	pfxs := common.RandomRealWorldPrefixes(10_000)

	psb := new(netipds.PrefixSetBuilder)
	for _, route := range pfxs {
		psb.Add(route)
	}
	rt := psb.PrefixSet()

	t.Run("IPv4", func(t *testing.T) {
		t.Parallel()
		for range 1_000 {
			ip := common.MatchIP4(pfxs)
			pfx := netip.PrefixFrom(ip, ip.BitLen())
			if ok := rt.Encompasses(pfx); !ok {
				t.Fatalf("Contains(%s), expected true, got %v", ip, ok)
			}
		}
	})

	t.Run("IPv6", func(t *testing.T) {
		t.Parallel()
		for range 1_000 {
			ip := common.MatchIP6(pfxs)
			pfx := netip.PrefixFrom(ip, ip.BitLen())
			if ok := rt.Encompasses(pfx); !ok {
				t.Fatalf("Contains(%s), expected true, got %v", ip, ok)
			}
		}
	})
}

func TestMissIP(t *testing.T) {
	t.Parallel()
	pfxs := common.RandomRealWorldPrefixes(10_000)

	psb := new(netipds.PrefixSetBuilder)
	for _, route := range pfxs {
		psb.Add(route)
	}
	rt := psb.PrefixSet()

	t.Run("IPv4", func(t *testing.T) {
		t.Parallel()
		for range 1_000 {
			ip := common.MissIP4(pfxs)
			pfx := netip.PrefixFrom(ip, ip.BitLen())
			if ok := rt.Encompasses(pfx); ok {
				t.Fatalf("Contains(%s), expected false, got %v", ip, ok)
			}
		}
	})

	t.Run("IPv6", func(t *testing.T) {
		t.Parallel()
		for range 1_000 {
			ip := common.MissIP6(pfxs)
			pfx := netip.PrefixFrom(ip, ip.BitLen())
			if ok := rt.Encompasses(pfx); ok {
				t.Fatalf("Contains(%s), expected false, got %v", ip, ok)
			}
		}
	})
}
