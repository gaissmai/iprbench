package main_test

import (
	"runtime"
	"testing"

	"local/iprbench/common"

	"github.com/yl2chen/cidranger"
)

func BenchmarkTier1PfxSize(b *testing.B) {
	var startMem, endMem runtime.MemStats

	for k := 100; k <= 1_000_000; k *= 10 {
		tree := cidranger.NewPCTrieRanger()
		runtime.GC()
		runtime.ReadMemStats(&startMem)

		b.Run(common.IntMap[k], func(b *testing.B) {
			for range b.N {
				for _, cidr := range tier1Routes[:k] {
					_ = tree.Insert(cidranger.NewBasicRangerEntry(common.PfxToIPNet(cidr)))
				}
			}
			runtime.GC()
			runtime.ReadMemStats(&endMem)

			b.ReportMetric(float64(endMem.HeapAlloc-startMem.HeapAlloc), "bytes")
			b.ReportMetric(0, "ns/op")
		})
	}
}

func BenchmarkRandomPfx4Size(b *testing.B) {
	var startMem, endMem runtime.MemStats

	for k := 100; k <= 1_000_000; k *= 10 {
		tree := cidranger.NewPCTrieRanger()
		runtime.GC()
		runtime.ReadMemStats(&startMem)

		b.Run(common.IntMap[k], func(b *testing.B) {
			for range b.N {
				for _, cidr := range randomRoutes4[:k] {
					_ = tree.Insert(cidranger.NewBasicRangerEntry(common.PfxToIPNet(cidr)))
				}
			}
			runtime.GC()
			runtime.ReadMemStats(&endMem)

			b.ReportMetric(float64(endMem.HeapAlloc-startMem.HeapAlloc), "bytes")
			b.ReportMetric(0, "ns/op")
		})
	}
}

func BenchmarkRandomPfx6Size(b *testing.B) {
	var startMem, endMem runtime.MemStats

	for k := 100; k <= 1_000_000; k *= 10 {
		tree := cidranger.NewPCTrieRanger()
		runtime.GC()
		runtime.ReadMemStats(&startMem)

		b.Run(common.IntMap[k], func(b *testing.B) {
			for range b.N {
				for _, cidr := range randomRoutes6[:k] {
					_ = tree.Insert(cidranger.NewBasicRangerEntry(common.PfxToIPNet(cidr)))
				}
			}
			runtime.GC()
			runtime.ReadMemStats(&endMem)

			b.ReportMetric(float64(endMem.HeapAlloc-startMem.HeapAlloc), "bytes")
			b.ReportMetric(0, "ns/op")
		})
	}
}

func BenchmarkRandomPfxSize(b *testing.B) {
	var startMem, endMem runtime.MemStats

	for k := 100; k <= 1_000_000; k *= 10 {
		tree := cidranger.NewPCTrieRanger()
		runtime.GC()
		runtime.ReadMemStats(&startMem)

		b.Run(common.IntMap[k], func(b *testing.B) {
			for range b.N {
				for _, cidr := range randomRoutes[:k] {
					_ = tree.Insert(cidranger.NewBasicRangerEntry(common.PfxToIPNet(cidr)))
				}
			}
			runtime.GC()
			runtime.ReadMemStats(&endMem)

			b.ReportMetric(float64(endMem.HeapAlloc-startMem.HeapAlloc), "bytes")
			b.ReportMetric(0, "ns/op")
		})
	}
}
