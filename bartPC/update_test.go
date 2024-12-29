package main_test

import (
	"runtime"
	"testing"

	"local/iprbench/common"

	"github.com/gaissmai/bart"
)

func BenchmarkInsertRandomPfxs(b *testing.B) {
	for _, k := range []int{1_000, 10_000, 100_000, 200_000} {
		name := common.IntMap[k]
		randomPfxs := common.RandomPrefixes(k)

		b.Run(name, func(b *testing.B) {
			rt := new(bart.Table[any]).WithPathCompression()

			runtime.GC()
			b.ResetTimer()
			for range b.N {
				for _, route := range randomPfxs {
					rt.Insert(route, nil)
				}
			}
			b.StopTimer()
			b.ReportMetric(float64(b.Elapsed())/float64(k)/float64(b.N), "ns/route")
			b.ReportMetric(0, "ns/op")
		})
	}
}

func BenchmarkDeleteRandomPfxs(b *testing.B) {
	for _, k := range []int{1_000, 10_000, 100_000, 200_000} {

		randomPfxs := common.RandomPrefixes(k)
		name := common.IntMap[k]

		b.Run(name, func(b *testing.B) {
			rt := new(bart.Table[any]).WithPathCompression()
			for _, route := range randomPfxs {
				rt.Insert(route, nil)
			}

			runtime.GC()
			b.ResetTimer()
			for range b.N {
				for _, route := range randomPfxs {
					rt.Delete(route)
				}
			}
			b.StopTimer()
			b.ReportMetric(float64(b.Elapsed())/float64(k)/float64(b.N), "ns/route")
			b.ReportMetric(0, "ns/op")
		})
	}
}