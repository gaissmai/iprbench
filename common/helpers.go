package common

import (
	"bufio"
	"compress/gzip"
	"log"
	"math/rand/v2"
	"net"
	"net/netip"
	"os"
	"strings"

	"github.com/gaissmai/bart"
	"github.com/gaissmai/extnetip"
)

var IntMap = map[int]string{
	1:         "1",
	2:         "2",
	5:         "5",
	10:        "10",
	20:        "20",
	50:        "50",
	100:       "100",
	500:       "500",
	1_000:     "1_000",
	10_000:    "10_000",
	50_000:    "50_000",
	100_000:   "100_000",
	200_000:   "200_000",
	500_000:   "500_000",
	1_000_000: "1_000_000",
	2_000_000: "2_000_000",
}

var mpp = netip.MustParsePrefix

var Prng = rand.New(rand.NewPCG(42, 42))

func IPis4(ip net.IP) bool {
	return ip.To4() != nil
}

func IPis6(ip net.IP) bool {
	return ip.To4() == nil && ip.To16() != nil
}

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
// The matching IP is found with the help of the art algorithm, it's the fastest algo.
func MatchIP4(routes []netip.Prefix) netip.Addr {
	rt := new(bart.Table[struct{}])
	for _, r := range routes {
		if r.Addr().Is4() {
			rt.Insert(r, struct{}{})
		}
	}

	i := 0
	for {
		// choose a random route from input
		pfx := routes[Prng.IntN(len(routes))]
		ip := pfx.Addr()
		if ip.Is6() {
			continue // wrong IP version
		}

		// just don't take the start IP
		ip = ip.Next()
		if ok := rt.Contains(ip); ok {
			return ip
		}
		i++
		if i > 500_000_000 {
			panic("couldn't find a matching IP, giving up!")
		}
	}
}

// MatchIP6 returns a random IP covered by the routing table.
// The matching IP is found with the help of the art algorithm, it's the fastest algo.
func MatchIP6(routes []netip.Prefix) netip.Addr {
	rt := new(bart.Table[struct{}])
	for _, r := range routes {
		if r.Addr().Is6() {
			rt.Insert(r, struct{}{})
		}
	}

	i := 0
	for {
		// choose a random route from input
		pfx := routes[Prng.IntN(len(routes))]
		ip := pfx.Addr()
		if ip.Is4() {
			continue // wrong IP version
		}

		// just don't take the start IP
		ip = ip.Next()
		if ok := rt.Contains(ip); ok {
			return ip
		}
		i++
		if i > 500_000_000 {
			panic("couldn't find a matching IP, giving up!")
		}
	}
}

// MissIP4 returns a random IP NOT covered by the routing table.
// The missing IP is found with the help of the art algorithm, it's the fastest algo.
func MissIP4(routes []netip.Prefix) netip.Addr {
	rt := new(bart.Table[struct{}])
	for _, r := range routes {
		if r.Addr().Is4() {
			rt.Insert(r, struct{}{})
		}
	}

	i := 0
	for {
		// choose a random route from input
		pfx := routes[Prng.IntN(len(routes))]
		ip := pfx.Addr()
		if ip.Is6() {
			continue // wrong IP version
		}

		// take last IP from prefix ...
		_, last := extnetip.Range(pfx)
		// ... add one
		ip = last.Next()
		if ok := rt.Contains(ip); !ok {
			return ip
		}
		i++
		if i > 500_000_000 {
			panic("couldn't find a missing IP, giving up!")
		}
	}
}

// MissIP6 returns a random IP NOT covered by the routing table.
// The missing IP is found with the help of the art algorithm, it's the fastest algo.
func MissIP6(routes []netip.Prefix) netip.Addr {
	rt := new(bart.Table[struct{}])
	for _, r := range routes {
		if r.Addr().Is6() {
			rt.Insert(r, struct{}{})
		}
	}

	i := 0
	for {
		// choose a random route from input
		pfx := routes[Prng.IntN(len(routes))]
		ip := pfx.Addr()
		if ip.Is4() {
			continue // wrong IP version
		}

		// take last IP from prefix ...
		_, last := extnetip.Range(pfx)
		// ... add one
		ip = last.Next()
		if ok := rt.Contains(ip); !ok {
			return ip
		}
		i++
		if i > 500_000_000 {
			panic("couldn't find a missing IP, giving up!")
		}
	}
}

// RandomRealWorldPrefixes returns n randomly generated prefixes with TODO
// IPv4 and IPv6 Prefixes are naturally distributed 4:1.
func RandomRealWorldPrefixes(n int) []netip.Prefix {
	ret := make([]netip.Prefix, 0, n)

	n4 := 4 * n / 5
	n6 := n / 5
	for n4+n6 < n {
		n4++
	}

	ret = append(ret, RandomRealWorldPrefixes4(n4)...)
	ret = append(ret, RandomRealWorldPrefixes6(n6)...)

	Prng.Shuffle(len(ret), func(i, j int) {
		ret[i], ret[j] = ret[j], ret[i]
	})
	return ret
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

// #########################################################

func RandomRealWorldPrefixes4(n int) []netip.Prefix {
	set := map[netip.Prefix]netip.Prefix{}
	pfxs := make([]netip.Prefix, 0, n)

	for {
		pfx := randomPrefix4()

		// skip too small or too big masks
		if pfx.Bits() < 8 && pfx.Bits() > 28 {
			continue
		}

		// skip multicast ...
		if pfx.Overlaps(mpp("240.0.0.0/8")) {
			continue
		}

		if _, ok := set[pfx]; !ok {
			set[pfx] = pfx
			pfxs = append(pfxs, pfx)
		}

		if len(set) >= n {
			break
		}
	}
	return pfxs
}

func RandomRealWorldPrefixes6(n int) []netip.Prefix {
	set := map[netip.Prefix]netip.Prefix{}
	pfxs := make([]netip.Prefix, 0, n)

	for {
		pfx := randomPrefix6()

		// skip too small or too big masks
		if pfx.Bits() < 16 || pfx.Bits() > 56 {
			continue
		}

		// skip non global routes seen in the real world
		if !pfx.Overlaps(mpp("2000::/3")) {
			continue
		}
		if pfx.Addr().Compare(mpp("2c0f::/16").Addr()) == 1 {
			continue
		}

		if _, ok := set[pfx]; !ok {
			set[pfx] = pfx
			pfxs = append(pfxs, pfx)
		}
		if len(set) >= n {
			break
		}
	}
	return pfxs
}

func randomPrefix4() netip.Prefix {
	bits := Prng.IntN(33)
	pfx, err := randomIP4().Prefix(bits)
	if err != nil {
		panic(err)
	}
	return pfx
}

func randomPrefix6() netip.Prefix {
	bits := Prng.IntN(129)
	pfx, err := randomIP6().Prefix(bits)
	if err != nil {
		panic(err)
	}
	return pfx
}

func randomIP4() netip.Addr {
	var b [4]byte
	for i := range b {
		b[i] = byte(Prng.Uint32() & 0xff)
	}
	return netip.AddrFrom4(b)
}

func randomIP6() netip.Addr {
	var b [16]byte
	for i := range b {
		b[i] = byte(Prng.Uint32() & 0xff)
	}
	return netip.AddrFrom16(b)
}
