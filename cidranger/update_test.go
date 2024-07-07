package main_test

import (
	"local/iprbench/common"
	"testing"

	"github.com/yl2chen/cidranger"
)

func BenchmarkInsert(b *testing.B) {
	for k := 100; k <= 1_000_000; k *= 10 {
		rt := cidranger.NewPCTrieRanger()
		for _, route := range tier1Routes[:k] {
			_ = rt.Insert(cidranger.NewBasicRangerEntry(common.PfxToIPNet(route)))
		}

		name := "Insert into " + common.IntMap[k]
		b.Run(name, func(b *testing.B) {
			ipNetProbe := common.PfxToIPNet(probe)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = rt.Insert(cidranger.NewBasicRangerEntry(ipNetProbe))
			}
		})
	}
}

func BenchmarkDelete(b *testing.B) {
	for k := 100; k <= 1_000_000; k *= 10 {
		rt := cidranger.NewPCTrieRanger()
		for _, route := range tier1Routes[:k] {
			_ = rt.Insert(cidranger.NewBasicRangerEntry(common.PfxToIPNet(route)))
		}

		name := "Delete from " + common.IntMap[k]
		ipNetProbe := common.PfxToIPNet(probe)
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				sink, _ = rt.Remove(ipNetProbe)
			}
		})
	}
}
