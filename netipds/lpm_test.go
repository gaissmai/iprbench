package main_test

import (
	"net/netip"
	"testing"

	"local/iprbench/common"

	"github.com/aromatt/netipds"
)

func BenchmarkLpmTier1Pfxs(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func(int, []netip.Prefix) []netip.Addr
	}{
		{"RandomMatchIP4", common.MatchManyIP4},
		{"RandomMatchIP6", common.MatchManyIP6},
		{"RandomMissIP4", common.MissManyIP4},
		{"RandomMissIP6", common.MissManyIP6},
	}

	psb := new(netipds.PrefixSetBuilder)
	for _, route := range tier1Routes {
		psb.Add(route)
	}
	ps := psb.PrefixSet()

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			manyIPs := bm.fn(common.N, tier1Routes)
			manyPfxs := common.AddrsToPfxs(manyIPs)

			i := 0
			ok := false
			for b.Loop() {
				ok = ps.Encompasses(manyPfxs[i&common.Mask])
				i++
			}
			common.Sink = ok
		})
	}
}

func BenchmarkLpmRandomPfxs(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func(int, []netip.Prefix) []netip.Addr
	}{
		{"RandomMatchIP4", common.MatchManyIP4},
		{"RandomMatchIP6", common.MatchManyIP6},
		{"RandomMissIP4", common.MissManyIP4},
		{"RandomMissIP6", common.MissManyIP6},
	}

	for _, k := range []int{1_000, 10_000, 100_000} {
		for _, bm := range benchmarks {

			psb := new(netipds.PrefixSetBuilder)
			for _, route := range randomRoutes[:k] {
				psb.Add(route)
			}
			ps := psb.PrefixSet()

			b.Run(common.IntMap[k]+"/"+bm.name, func(b *testing.B) {
				manyIPs := bm.fn(common.N, randomRoutes[:k])
				manyPfxs := common.AddrsToPfxs(manyIPs)

				i := 0
				ok := false
				for b.Loop() {
					ok = ps.Encompasses(manyPfxs[i&common.Mask])
					i++
				}
				common.Sink = ok
			})
		}
	}
}
