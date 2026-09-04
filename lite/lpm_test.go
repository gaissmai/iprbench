package main_test

import (
	"net/netip"
	"testing"

	"local/iprbench/common"

	"github.com/gaissmai/bart"
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

	rt := new(bart.Lite)
	for _, route := range tier1Routes {
		rt.Insert(route)
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			manyIPs := bm.fn(common.N, tier1Routes)

			i := 0
			ok := false
			for b.Loop() {
				ok = rt.Contains(manyIPs[i&common.Mask])
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

			rt := new(bart.Lite)
			for _, route := range randomRoutes[:k] {
				rt.Insert(route)
			}

			b.Run(common.IntMap[k]+"/"+bm.name, func(b *testing.B) {
				manyIPs := bm.fn(common.N, randomRoutes[:k])

				i := 0
				ok := false
				for b.Loop() {
					ok = rt.Contains(manyIPs[i&common.Mask])
					i++
				}
				common.Sink = ok
			})
		}
	}
}
