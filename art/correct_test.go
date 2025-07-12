package main_test

import (
	"testing"

	"local/iprbench/common"

	"github.com/tailscale/art"
)

func TestMatchIP(t *testing.T) {
	t.Parallel()
	pfxs := common.RandomRealWorldPrefixes(10_000)

	rt := new(art.Table[any])
	for _, route := range pfxs {
		rt.Insert(route, nil)
	}

	t.Run("IPv4", func(t *testing.T) {
		t.Parallel()
		for range 1_000 {
			ip := common.MatchIP4(pfxs)
			if _, ok := rt.Get(ip); !ok {
				t.Fatalf("Contains(%s), expected true, got %v", ip, ok)
			}
		}
	})

	t.Run("IPv6", func(t *testing.T) {
		t.Parallel()
		for range 1_000 {
			ip := common.MatchIP6(pfxs)
			if _, ok := rt.Get(ip); !ok {
				t.Fatalf("Contains(%s), expected true, got %v", ip, ok)
			}
		}
	})
}

func TestMissIP(t *testing.T) {
	t.Parallel()
	pfxs := common.RandomRealWorldPrefixes(10_000)

	rt := new(art.Table[any])
	for _, route := range pfxs {
		rt.Insert(route, nil)
	}

	t.Run("IPv4", func(t *testing.T) {
		t.Parallel()
		for range 1_000 {
			ip := common.MissIP4(pfxs)
			if _, ok := rt.Get(ip); ok {
				t.Fatalf("Contains(%s), expected false, got %v", ip, ok)
			}
		}
	})

	t.Run("IPv6", func(t *testing.T) {
		t.Parallel()
		for range 1_000 {
			ip := common.MissIP6(pfxs)
			if _, ok := rt.Get(ip); ok {
				t.Fatalf("Contains(%s), expected false, got %v", ip, ok)
			}
		}
	})
}
