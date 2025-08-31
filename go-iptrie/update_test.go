package main_test

import (
	"testing"

	"local/iprbench/common"

	"github.com/phemmer/go-iptrie"
)

func BenchmarkInsertRandomPfxs(b *testing.B) {
	for _, k := range []int{1_000, 10_000, 100_000, 200_000} {
		name := common.IntMap[k]
		randomPfxs := common.RandomRealWorldPrefixes(k)

		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				rt := iptrie.NewTrie()

				b.StartTimer()
				for _, route := range randomPfxs {
					rt.Insert(route, nil)
				}
			}
			b.ReportMetric(float64(b.Elapsed())/float64(k)/float64(b.N), "ns/route")
			b.ReportMetric(0, "ns/op")
		})
	}
}

func BenchmarkDeleteRandomPfxs(b *testing.B) {
	for _, k := range []int{1_000, 10_000, 100_000, 200_000} {

		randomPfxs := common.RandomRealWorldPrefixes(k)
		name := common.IntMap[k]

		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				rt := iptrie.NewTrie()
				for _, route := range randomPfxs {
					rt.Insert(route, nil)
				}

				b.StartTimer()
				for _, route := range randomPfxs {
					rt.Remove(route)
				}
			}
			b.ReportMetric(float64(b.Elapsed())/float64(k)/float64(b.N), "ns/route")
			b.ReportMetric(0, "ns/op")
		})
	}
}
