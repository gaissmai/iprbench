package main_test

import (
	"net/netip"
	"testing"

	"local/iprbench/common"

	"github.com/gaissmai/bart"
)

func BenchmarkLpmTier1Pfxs(b *testing.B) {
	rt := new(bart.Table[any])
	for _, route := range tier1Routes {
		rt.Insert(route, nil)
	}
	lpm := lpmWrapper(rt)

	benchmarks := []struct {
		name string
		is4  bool
		fn   func(func(netip.Addr) bool, bool) netip.Addr
	}{
		{"RandomMatchIP4", true, common.MatchIP},
		{"RandomMatchIP6", false, common.MatchIP},
		{"RandomMissIP4", true, common.MissIP},
		{"RandomMissIP6", false, common.MissIP},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			ip := bm.fn(lpm, bm.is4)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				sink = lpm(ip)
			}
		})
	}
}

func BenchmarkLpmRandomPfxs100_000(b *testing.B) {
	rt := new(bart.Table[any])
	for _, route := range randomRoutes[:100_000] {
		rt.Insert(route, nil)
	}
	lpm := lpmWrapper(rt)

	benchmarks := []struct {
		name string
		is4  bool
		fn   func(func(netip.Addr) bool, bool) netip.Addr
	}{
		{"RandomMatchIP4", true, common.MatchIP},
		{"RandomMatchIP6", false, common.MatchIP},
		{"RandomMissIP4", true, common.MissIP},
		{"RandomMissIP6", false, common.MissIP},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			ip := bm.fn(lpm, bm.is4)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				sink = lpm(ip)
			}
		})
	}
}
