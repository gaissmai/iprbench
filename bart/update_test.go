package main_test

import (
	"local/iprbench/common"
	"testing"

	"github.com/gaissmai/bart"
)

func BenchmarkInsert(b *testing.B) {
	for k := 100; k <= 1_000_000; k *= 10 {
		rt := new(bart.Table[any])
		for _, route := range tier1Routes[:k] {
			rt.Insert(route, nil)
		}

		name := "Insert into " + common.IntMap[k]
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				rt.Insert(probe, nil)
			}
		})
	}
}

func BenchmarkDelete(b *testing.B) {
	for k := 100; k <= 1_000_000; k *= 10 {
		rt := new(bart.Table[any])
		for _, route := range tier1Routes[:k] {
			rt.Insert(route, nil)
		}

		name := "Delete from " + common.IntMap[k]
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				rt.Delete(probe)
			}
		})
	}
}
