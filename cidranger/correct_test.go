package main_test

import (
	"testing"

	"local/iprbench/common"

	"github.com/yl2chen/cidranger"
)

func TestMatchIP(t *testing.T) {
	t.Parallel()
	pfxs := common.RandomRealWorldPrefixes(10_000)

	rt := cidranger.NewPCTrieRanger()
	for _, route := range pfxs {
		rt.Insert(cidranger.NewBasicRangerEntry(common.PfxToIPNet(route)))
	}

	t.Run("IPv4", func(t *testing.T) {
		t.Parallel()
		for range 1_000 {
			netIP := common.MatchIP4(pfxs)
			ip := common.AddrToIP(netIP)
			if ok, _ := rt.Contains(ip); !ok {
				t.Fatalf("Contains(%s), expected true, got %v", ip, ok)
			}
		}
	})

	t.Run("IPv6", func(t *testing.T) {
		t.Parallel()
		for range 1_000 {
			netIP := common.MatchIP6(pfxs)
			ip := common.AddrToIP(netIP)
			if ok, _ := rt.Contains(ip); !ok {
				t.Fatalf("Contains(%s), expected true, got %v", ip, ok)
			}
		}
	})
}

func TestMissIP(t *testing.T) {
	t.Parallel()
	pfxs := common.RandomRealWorldPrefixes(10_000)

	rt := cidranger.NewPCTrieRanger()
	for _, route := range pfxs {
		rt.Insert(cidranger.NewBasicRangerEntry(common.PfxToIPNet(route)))
	}

	t.Run("IPv4", func(t *testing.T) {
		t.Parallel()
		for range 1_000 {
			netIP := common.MissIP4(pfxs)
			t.Log(netIP)
			ip := common.AddrToIP(netIP)
			if ok, _ := rt.Contains(ip); ok {
				t.Fatalf("Contains(%s), expected false, got %v", ip, ok)
			}
		}
	})

	t.Run("IPv6", func(t *testing.T) {
		t.Parallel()
		for range 1_000 {
			netIP := common.MissIP6(pfxs)
			ip := common.AddrToIP(netIP)
			if ok, _ := rt.Contains(ip); ok {
				t.Fatalf("Contains(%s), expected false, got %v", ip, ok)
			}
		}
	})
}
