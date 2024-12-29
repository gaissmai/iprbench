package main_test

import (
	"net/netip"
	"testing"

	"local/iprbench/common"

	"github.com/yl2chen/cidranger"
)

func BenchmarkLpmTier1Pfxs(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func([]netip.Prefix) netip.Addr
	}{
		{"RandomMatchIP4", common.MatchIP4},
		{"RandomMatchIP6", common.MatchIP6},
		{"RandomMissIP4", common.MissIP4},
		{"RandomMissIP6", common.MissIP6},
	}

	rt := cidranger.NewPCTrieRanger()
	for _, route := range tier1Routes {
		rt.Insert(cidranger.NewBasicRangerEntry(common.PfxToIPNet(route)))
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			netIP := bm.fn(tier1Routes)
			ip := common.AddrToIP(netIP)
			b.ResetTimer()
			for range b.N {
				_, _ = rt.Contains(ip)
			}
		})
	}
}

func BenchmarkLpmRandomPfxs(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func([]netip.Prefix) netip.Addr
	}{
		{"RandomMatchIP4", common.MatchIP4},
		{"RandomMatchIP6", common.MatchIP6},
		{"RandomMissIP4", common.MissIP4},
		{"RandomMissIP6", common.MissIP6},
	}

	for _, k := range []int{1_000, 10_000, 100_000} {
		for _, bm := range benchmarks {

			rt := cidranger.NewPCTrieRanger()
			for _, route := range randomRoutes[:k] {
				rt.Insert(cidranger.NewBasicRangerEntry(common.PfxToIPNet(route)))
			}

			b.Run(common.IntMap[k]+"/"+bm.name, func(b *testing.B) {
				netIP := bm.fn(randomRoutes[:k]) // get a random matching or missing ip
				ip := common.AddrToIP(netIP)
				b.ResetTimer()
				for range b.N {
					_, _ = rt.Contains(ip)
				}
			})
		}
	}
}
