package main_test

import (
	"runtime"
	"testing"

	"local/iprbench/common"

	"github.com/aromatt/netipds"
)

func BenchmarkTier1PfxSize(b *testing.B) {
	var startMem, endMem runtime.MemStats

	for _, k := range []int{1_000, 10_000, 100_000, 200_000} {
		psb := new(netipds.PrefixSetBuilder)

		runtime.GC()
		runtime.ReadMemStats(&startMem)

		b.Run(common.IntMap[k], func(b *testing.B) {
			for range b.N {
				for _, cidr := range tier1Routes[:k] {
					psb.Add(cidr)
				}
			}
			runtime.GC()
			runtime.ReadMemStats(&endMem)

			b.ReportMetric(float64(endMem.HeapAlloc-startMem.HeapAlloc)/(float64(k)), "bytes/route")
			b.ReportMetric(0, "ns/op")
		})
	}
}

func BenchmarkRandomPfx4Size(b *testing.B) {
	var startMem, endMem runtime.MemStats

	for _, k := range []int{1_000, 10_000, 100_000, 200_000} {
		psb := new(netipds.PrefixSetBuilder)

		runtime.GC()
		runtime.ReadMemStats(&startMem)

		b.Run(common.IntMap[k], func(b *testing.B) {
			for range b.N {
				for _, cidr := range randomRoutes4[:k] {
					psb.Add(cidr)
				}
			}
			runtime.GC()
			runtime.ReadMemStats(&endMem)

			b.ReportMetric(float64(endMem.HeapAlloc-startMem.HeapAlloc)/(float64(k)), "bytes/route")
			b.ReportMetric(0, "ns/op")
		})
	}
}

func BenchmarkRandomPfx6Size(b *testing.B) {
	var startMem, endMem runtime.MemStats

	for _, k := range []int{1_000, 10_000, 100_000, 200_000} {
		psb := new(netipds.PrefixSetBuilder)

		runtime.GC()
		runtime.ReadMemStats(&startMem)

		b.Run(common.IntMap[k], func(b *testing.B) {
			for range b.N {
				for _, cidr := range randomRoutes6[:k] {
					psb.Add(cidr)
				}
			}
			runtime.GC()
			runtime.ReadMemStats(&endMem)

			b.ReportMetric(float64(endMem.HeapAlloc-startMem.HeapAlloc)/(float64(k)), "bytes/route")
			b.ReportMetric(0, "ns/op")
		})
	}
}

func BenchmarkRandomPfxSize(b *testing.B) {
	var startMem, endMem runtime.MemStats

	for _, k := range []int{1_000, 10_000, 100_000, 200_000} {
		psb := new(netipds.PrefixSetBuilder)

		runtime.GC()
		runtime.ReadMemStats(&startMem)

		b.Run(common.IntMap[k], func(b *testing.B) {
			for range b.N {
				for _, cidr := range randomRoutes[:k] {
					psb.Add(cidr)
				}
			}
			runtime.GC()
			runtime.ReadMemStats(&endMem)

			b.ReportMetric(float64(endMem.HeapAlloc-startMem.HeapAlloc)/(float64(k)), "bytes/route")
			b.ReportMetric(0, "ns/op")
		})
	}
}
