package main_test

import (
	"net/netip"
	"testing"

	"local/iprbench/common"
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

	rt := NewTable[any]()
	for _, route := range tier1Routes {
		rt.Insert(route, nil)
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			manyIPs := bm.fn(common.N, tier1Routes)

			i := 0
			for b.Loop() {
				rt.Lookup(manyIPs[i&common.Mask])
				i++
			}

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

			rt := NewTable[any]()
			for _, route := range randomRoutes[:k] {
				rt.Insert(route, nil)
			}

			b.Run(common.IntMap[k]+"/"+bm.name, func(b *testing.B) {
				manyIPs := bm.fn(common.N, randomRoutes[:k])

				i := 0
				for b.Loop() {
					rt.Lookup(manyIPs[i&common.Mask])
					i++
				}

			})
		}
	}
}
