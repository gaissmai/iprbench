package main_test

import (
	"net/netip"
	"testing"

	"local/iprbench/common"

	"github.com/gaissmai/cidrtree"
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

	rt := new(cidrtree.Table[any])
	for _, route := range tier1Routes {
		rt.Insert(route, nil)
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			ip := bm.fn(tier1Routes)
			b.ResetTimer()
			for range b.N {
				_, _, _ = rt.Lookup(ip)
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

			rt := new(cidrtree.Table[any])
			for _, route := range randomRoutes[:k] {
				rt.Insert(route, nil)
			}

			b.Run(common.IntMap[k]+"/"+bm.name, func(b *testing.B) {
				ip := bm.fn(randomRoutes[:k]) // get a random matching or missing ip
				b.ResetTimer()
				for range b.N {
					_, _, _ = rt.Lookup(ip)
				}
			})
		}
	}
}
