package main_test

import (
	"net/netip"
	"testing"

	"github.com/aromatt/netipds"
	"local/iprbench/common"
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

	psb := new(netipds.PrefixSetBuilder)
	for _, route := range tier1Routes {
		psb.Add(route)
	}
	ps := psb.PrefixSet()

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			ip := bm.fn(tier1Routes)
			pfx := netip.PrefixFrom(ip, ip.BitLen())
			b.ResetTimer()
			for range b.N {
				_ = ps.Encompasses(pfx)
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

			psb := new(netipds.PrefixSetBuilder)
			for _, route := range randomRoutes[:k] {
				psb.Add(route)
			}
			ps := psb.PrefixSet()

			b.Run(common.IntMap[k]+"/"+bm.name, func(b *testing.B) {
				ip := bm.fn(randomRoutes[:k]) // get a random matching or missing ip
				pfx := netip.PrefixFrom(ip, ip.BitLen())
				b.ResetTimer()
				for range b.N {
					_ = ps.Encompasses(pfx)
				}
			})
		}
	}
}
