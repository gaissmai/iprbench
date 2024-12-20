package main_test

import (
	"testing"

	"local/iprbench/common"
)

func BenchmarkInsert(b *testing.B) {
	for k := 100; k <= 1_000_000; k *= 10 {
		rt := NewTable()
		for _, route := range tier1Routes[:k] {
			rt.Insert(common.PfxToIPNet(route), nil)
		}

		name := "Insert into " + common.IntMap[k]
		cidrProbe := common.PfxToIPNet(probe)
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				rt.Insert(cidrProbe, nil)
			}
		})
	}
}

func BenchmarkDelete(b *testing.B) {
	for k := 100; k <= 1_000_000; k *= 10 {
		rt := NewTable()
		for _, route := range tier1Routes[:k] {
			rt.Insert(common.PfxToIPNet(route), nil)
		}

		name := "Delete from " + common.IntMap[k]
		cidrProbe := common.PfxToIPNet(probe)
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				rt.Delete(cidrProbe)
			}
		})
	}
}
