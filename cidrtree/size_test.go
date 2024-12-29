package main_test

import (
	"runtime"
	"testing"

	"local/iprbench/common"

	"github.com/gaissmai/cidrtree"
)

func BenchmarkTier1PfxSize(b *testing.B) {
	var startMem, endMem runtime.MemStats

	for _, k := range []int{1_000, 10_000, 100_000, 200_000} {
		tree := new(cidrtree.Table[any])
		runtime.GC()
		runtime.ReadMemStats(&startMem)

		b.Run(common.IntMap[k], func(b *testing.B) {
			for _, cidr := range tier1Routes[:k] {
				tree.Insert(cidr, nil)
			}
			runtime.GC()
			runtime.ReadMemStats(&endMem)

			b.ReportMetric(float64(endMem.HeapAlloc-startMem.HeapAlloc)/float64(k), "bytes/route")
			b.ReportMetric(0, "ns/op")
		})
	}
}

func BenchmarkRandomPfx4Size(b *testing.B) {
	var startMem, endMem runtime.MemStats

	for _, k := range []int{1_000, 10_000, 100_000, 200_000} {
		tree := new(cidrtree.Table[any])
		runtime.GC()
		runtime.ReadMemStats(&startMem)

		b.Run(common.IntMap[k], func(b *testing.B) {
			for _, cidr := range randomRoutes4[:k] {
				tree.Insert(cidr, nil)
			}
			runtime.GC()
			runtime.ReadMemStats(&endMem)

			b.ReportMetric(float64(endMem.HeapAlloc-startMem.HeapAlloc)/float64(k), "bytes/route")
			b.ReportMetric(0, "ns/op")
		})
	}
}

func BenchmarkRandomPfx6Size(b *testing.B) {
	var startMem, endMem runtime.MemStats

	for _, k := range []int{1_000, 10_000, 100_000, 200_000} {
		tree := new(cidrtree.Table[any])
		runtime.GC()
		runtime.ReadMemStats(&startMem)

		b.Run(common.IntMap[k], func(b *testing.B) {
			for _, cidr := range randomRoutes6[:k] {
				tree.Insert(cidr, nil)
			}
			runtime.GC()
			runtime.ReadMemStats(&endMem)

			b.ReportMetric(float64(endMem.HeapAlloc-startMem.HeapAlloc)/float64(k), "bytes/route")
			b.ReportMetric(0, "ns/op")
		})
	}
}

func BenchmarkRandomPfxSize(b *testing.B) {
	var startMem, endMem runtime.MemStats

	for _, k := range []int{1_000, 10_000, 100_000, 200_000} {
		tree := new(cidrtree.Table[any])
		runtime.GC()
		runtime.ReadMemStats(&startMem)

		b.Run(common.IntMap[k], func(b *testing.B) {
			for _, cidr := range randomRoutes[:k] {
				tree.Insert(cidr, nil)
			}
			runtime.GC()
			runtime.ReadMemStats(&endMem)

			b.ReportMetric(float64(endMem.HeapAlloc-startMem.HeapAlloc)/float64(k), "bytes/route")
			b.ReportMetric(0, "ns/op")
		})
	}
}
