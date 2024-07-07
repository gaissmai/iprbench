package main_test

import (
	"net/netip"
	"testing"

	"local/iprbench/common"

	"github.com/yl2chen/cidranger"
)

var rt1 = cidranger.NewPCTrieRanger()
var rt2 = cidranger.NewPCTrieRanger()

func init() {
	for _, route := range tier1Routes {
		_ = rt1.Insert(cidranger.NewBasicRangerEntry(common.PfxToIPNet(route)))
	}
}

func init() {
	for _, route := range randomRoutes[:100_000] {
		_ = rt2.Insert(cidranger.NewBasicRangerEntry(common.PfxToIPNet(route)))
	}
}

func BenchmarkLpmTier1Pfxs(b *testing.B) {
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
			ip := common.AddrToIP(bm.fn(bm.routes))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				sink, _ = rt1.ContainingNetworks(ip)
			}
		})
	}
}

func BenchmarkLpmRandomPfxs100_000(b *testing.B) {
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
			ip := common.AddrToIP(bm.fn(bm.routes))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				sink, _ = rt2.ContainingNetworks(ip)
			}
		})
	}
}
