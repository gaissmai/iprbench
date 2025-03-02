package main_test

import (
	"runtime"
	"testing"

	"local/iprbench/common"

	"github.com/yl2chen/cidranger"
)

func BenchmarkInsertRandomPfxs(b *testing.B) {
	for _, k := range []int{1_000, 10_000, 100_000, 200_000} {
		name := common.IntMap[k]

		randomPfxs := common.RandomRealWorldPrefixes(k)
		randomRangerEntries := make([]cidranger.RangerEntry, 0, k)
		for _, pfx := range randomPfxs {
			randomRangerEntries = append(randomRangerEntries, cidranger.NewBasicRangerEntry(common.PfxToIPNet(pfx)))
		}

		b.Run(name, func(b *testing.B) {
			rt := cidranger.NewPCTrieRanger()

			runtime.GC()
			for b.Loop() {
				for _, route := range randomRangerEntries {
					rt.Insert(route)
				}
			}
			b.ReportMetric(float64(b.Elapsed())/float64(k)/float64(b.N), "ns/route")
			b.ReportMetric(0, "ns/op")
		})
	}
}

func BenchmarkDeleteRandomPfxs(b *testing.B) {
	for _, k := range []int{1_000, 10_000, 100_000, 200_000} {
		name := common.IntMap[k]

		randomPfxs := common.RandomRealWorldPrefixes(k)
		randomRangerEntries := make([]cidranger.RangerEntry, 0, k)
		for _, pfx := range randomPfxs {
			randomRangerEntries = append(randomRangerEntries, cidranger.NewBasicRangerEntry(common.PfxToIPNet(pfx)))
		}

		b.Run(name, func(b *testing.B) {
			rt := cidranger.NewPCTrieRanger()
			for _, route := range randomRangerEntries {
				rt.Insert(route)
			}

			runtime.GC()
			for b.Loop() {
				for _, route := range randomRangerEntries {
					rt.Remove(route.Network())
				}
			}
			b.ReportMetric(float64(b.Elapsed())/float64(k)/float64(b.N), "ns/route")
			b.ReportMetric(0, "ns/op")
		})
	}
}
