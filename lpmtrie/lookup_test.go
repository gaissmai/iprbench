package main_test

import (
	"net/netip"
	"testing"

	"local/iprbench/common"
)

func BenchmarkLpmTier1Pfxs(b *testing.B) {
	rt := NewTable()
	for _, route := range tier1Routes {
		rt.Insert(route, nil)
	}

	benchmarks := []struct {
		name   string
		routes []netip.Prefix
		fn     func([]netip.Prefix) netip.Addr
	}{
		{"RandomMatchIP4", tier1Routes, common.MatchIP4},
		{"RandomMatchIP6", tier1Routes, common.MatchIP6},
		{"RandomMissIP4", tier1Routes, common.MissIP4},
		{"RandomMissIP6", tier1Routes, common.MissIP6},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			ip := bm.fn(bm.routes)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, sink = rt.Get(ip)
			}
		})
	}
}

func BenchmarkLpmRandomPfxs100_000(b *testing.B) {
	rt := NewTable()
	for _, route := range randomRoutes[:100_000] {
		rt.Insert(route, nil)
	}

	benchmarks := []struct {
		name   string
		routes []netip.Prefix
		fn     func([]netip.Prefix) netip.Addr
	}{
		{"RandomMatchIP4", randomRoutes[:100_000], common.MatchIP4},
		{"RandomMatchIP6", randomRoutes[:100_000], common.MatchIP6},
		{"RandomMissIP4", randomRoutes[:100_000], common.MissIP4},
		{"RandomMissIP6", randomRoutes[:100_000], common.MissIP6},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			ip := bm.fn(bm.routes)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, sink = rt.Get(ip)
			}
		})
	}
}
