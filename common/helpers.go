package common

import (
	"bufio"
	"compress/gzip"
	"log"
	"math/rand"
	"net"
	"net/netip"
	"os"
	"strings"

	"github.com/tailscale/art"
)

var IntMap = map[int]string{
	1:         "1",
	10:        "10",
	100:       "100",
	1_000:     "1_000",
	10_000:    "10_000",
	100_000:   "100_000",
	1_000_000: "1_000_000",
}

var Prng = rand.New(rand.NewSource(42))

// AddrToIP converts a netip.Addr to net.IP.
func AddrToIP(addr netip.Addr) net.IP {
	return net.IP(addr.AsSlice())
}

// PfxToIPNet converts a netip.Prefix to net.IPNet.
func PfxToIPNet(p netip.Prefix) net.IPNet {
	return net.IPNet{
		IP:   p.Addr().AsSlice(),
		Mask: net.CIDRMask(p.Bits(), p.Addr().BitLen()),
	}
}

// MatchIP4 returns a random IP covered by the routing table.
func MatchIP4(routes []netip.Prefix) netip.Addr {
	rt := new(art.Table[struct{}])
	for _, r := range routes {
		if r.Addr().Is4() {
			rt.Insert(r, struct{}{})
		}
	}

	i := 0
	for {
		ip := RandomAddr4()
		if _, ok := rt.Get(ip); ok {
			return ip
		}
		i++
		if i > 20_000_000 {
			panic("logic error")
		}
	}
}

// MatchIP6 returns a random IP covered by the routing table.
func MatchIP6(routes []netip.Prefix) netip.Addr {
	rt := new(art.Table[struct{}])
	for _, r := range routes {
		if r.Addr().Is6() {
			rt.Insert(r, struct{}{})
		}
	}

	i := 0
	for {
		ip := RandomAddr6()
		if _, ok := rt.Get(ip); ok {
			return ip
		}
		i++
		if i > 20_000_000 {
			panic("logic error")
		}
	}
}

// MissIP4 returns a random IP covered by the routing table.
func MissIP4(routes []netip.Prefix) netip.Addr {
	rt := new(art.Table[struct{}])
	for _, r := range routes {
		if r.Addr().Is4() {
			rt.Insert(r, struct{}{})
		}
	}

	i := 0
	for {
		ip := RandomAddr4()
		if _, ok := rt.Get(ip); !ok {
			return ip
		}
		i++
		if i > 20_000_000 {
			panic("logic error")
		}
	}
}

// MissIP6 returns a random IP covered by the routing table.
func MissIP6(routes []netip.Prefix) netip.Addr {
	rt := new(art.Table[struct{}])
	for _, r := range routes {
		if r.Addr().Is6() {
			rt.Insert(r, struct{}{})
		}
	}

	i := 0
	for {
		ip := RandomAddr6()
		if _, ok := rt.Get(ip); !ok {
			return ip
		}
		i++
		if i > 20_000_000 {
			panic("logic error")
		}
	}
}

// RandomPrefixes returns n randomly generated prefixes without default routes.
// IPv6 and IPv6 Prefixes are naturally distributed 4:1.
func RandomPrefixes(n int) []netip.Prefix {
	ret := make([]netip.Prefix, 0, n)

	ret = append(ret, RandomPrefixes4(4*n/5)...)
	ret = append(ret, RandomPrefixes6(n/5)...)

	Prng.Shuffle(len(ret), func(i, j int) {
		ret[i], ret[j] = ret[j], ret[i]
	})
	return ret
}

// RandomPrefixes4 returns n randomly generated IPv4 prefixes without default route.
func RandomPrefixes4(n int) []netip.Prefix {
	pfxs := map[netip.Prefix]bool{}

	for len(pfxs) < n {
		bits := Prng.Intn(32)
		// skip default routes
		bits += 1
		pfx, err := RandomAddr4().Prefix(bits)
		if err != nil {
			panic(err)
		}
		pfxs[pfx] = true
	}

	ret := make([]netip.Prefix, 0, len(pfxs))
	for pfx := range pfxs {
		ret = append(ret, pfx)
	}

	return ret
}

// RandomPrefixes6 returns n randomly generated IPv6 prefixes without default route.
func RandomPrefixes6(n int) []netip.Prefix {
	pfxs := map[netip.Prefix]bool{}

	for len(pfxs) < n {
		bits := Prng.Intn(128)
		// skip default routes
		bits += 1
		pfx, err := RandomAddr6().Prefix(bits)
		if err != nil {
			panic(err)
		}
		pfxs[pfx] = true
	}

	ret := make([]netip.Prefix, 0, len(pfxs))
	for pfx := range pfxs {
		ret = append(ret, pfx)
	}

	return ret
}

// RandomAddr returns a randomly generated IP address.
// IPv4 and IPv6 Prefixes are distributed 1:1
func RandomAddr() netip.Addr {
	if Prng.Intn(2) == 1 {
		return RandomAddr6()
	}
	return RandomAddr4()
}

// RandomAddr4 returns a randomly generated IPv4 address.
func RandomAddr4() netip.Addr {
	var b [4]byte
	if _, err := Prng.Read(b[:]); err != nil {
		panic(err)
	}
	return netip.AddrFrom4(b)
}

// RandomAddr6 returns a randomly generated IPv6 address.
func RandomAddr6() netip.Addr {
	var b [16]byte
	if _, err := Prng.Read(b[:]); err != nil {
		panic(err)
	}
	return netip.AddrFrom16(b)
}

func ReadFullTableShuffled(pfxFname string) []netip.Prefix {
	file, err := os.Open(pfxFname)
	if err != nil {
		log.Fatal(err)
	}

	rgz, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}

	var ret []netip.Prefix
	scanner := bufio.NewScanner(rgz)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		cidr := netip.MustParsePrefix(line)
		ret = append(ret, cidr)
	}
	if err := scanner.Err(); err != nil {
		log.Printf("reading from %v, %v", rgz, err)
	}

	Prng.Shuffle(len(ret), func(i, j int) {
		ret[i], ret[j] = ret[j], ret[i]
	})
	return ret
}
