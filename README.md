# iprbench
comparing benchmarks for some golang IP routing table implementations:

```
	github.com/gaissmai/bart.Table
	github.com/gaissmai/bart.Lite  # added to suite
	github.com/gaissmai/bart.Fast  # added to suite
	github.com/aromatt/netipds
	github.com/k-sone/critbitgo
	github.com/Asphaltt/lpmtrie
	github.com/kentik/patricia

	github.com/tailscale/art        # removed from suite
	github.com/yl2chen/cidranger    # removed from suite
	github.com/gaissmai/cidrtree    # removed from suite
	github.com/phemmer/go-iptrie    # removed from suite
```

The ~1_000_000 **Tier1** prefix test records (IPv4 and IPv6 routes) are from a full Internet
routing table with typical ISP prefix distribution.

In comparison, the prefix lengths for the _real-world_ random test sets are equally distributed
between /8-28 for IPv4 and /16-56 bits for IPv6 (limited to the 2000::/3 global unicast address space).

The _real-world_ **RandomPrefixes** without IP versions labeling are composed of a distribution
of 4 parts IPv4 to 1 part IPv6 random prefixes, which is approximately the current ratio in the Internet backbone routers.

## make your own benchmarks

```
  $ # set the proper cpu feature flags, e.g.
  $ export GOAMD64=v3

  $ make dep
  $ make -B all   # takes some time!
```

## lpm (longest-prefix-match)

`bart.Fast` is by far the fastest software algorithm for IP-address longest-prefix-match.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ fast/lpm.bm  │              lite/lpm.bm               │
                                       │    sec/op    │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4-8            8.785n ±  5%   16.050n ±  6%   +82.70% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            13.28n ±  6%    26.78n ± 17%  +101.73% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             9.080n ±  3%   17.420n ±  2%   +91.84% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             15.22n ± 24%    27.29n ± 27%   +79.24% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     14.54n ±  2%    18.11n ±  4%   +24.54% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     15.04n ±  3%    20.93n ±  2%   +39.16% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      14.32n ± 37%    17.56n ± 17%   +22.66% (p=0.001 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      14.81n ±  1%    20.94n ±  3%   +41.36% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    14.37n ±  4%    19.39n ±  1%   +34.93% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    14.80n ±  3%    21.48n ±  3%   +45.14% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     14.45n ±  2%    19.38n ±  3%   +34.15% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     15.03n ±  5%    21.07n ±  4%   +40.19% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   7.897n ± 14%   12.660n ± 36%   +60.33% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   16.82n ± 43%    26.63n ± 41%   +58.31% (p=0.002 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    14.34n ± 37%    19.26n ±  5%   +34.27% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    17.12n ±  3%    26.88n ±  3%   +56.99% (p=0.000 n=20)
geomean                                  13.44n          20.32n         +51.25%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ fast/lpm.bm  │              bart/lpm.bm              │
                                       │    sec/op    │    sec/op      vs base                │
LpmTier1Pfxs/RandomMatchIP4-8            8.785n ±  5%   16.405n ±  6%  +86.74% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            13.28n ±  6%    26.52n ± 18%  +99.81% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             9.080n ±  3%   17.485n ±  2%  +92.56% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             15.22n ± 24%    26.87n ± 29%  +76.45% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     14.54n ±  2%    18.16n ±  1%  +24.89% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     15.04n ±  3%    20.97n ±  3%  +39.43% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      14.32n ± 37%    17.62n ± 22%  +23.01% (p=0.002 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      14.81n ±  1%    20.95n ±  2%  +41.42% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    14.37n ±  4%    19.88n ±  4%  +38.31% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    14.80n ±  3%    20.71n ±  6%  +39.97% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     14.45n ±  2%    19.61n ±  3%  +35.74% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     15.03n ±  5%    20.80n ±  3%  +38.42% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   7.897n ± 14%   12.880n ± 35%  +63.11% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   16.82n ± 43%    26.85n ± 43%  +59.58% (p=0.002 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    14.34n ± 37%    19.14n ±  7%  +33.47% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    17.12n ±  3%    27.69n ±  3%  +61.69% (p=0.000 n=20)
geomean                                  13.44n          20.38n        +51.68%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ fast/lpm.bm  │             netipds/lpm.bm             │
                                       │    sec/op    │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4-8            8.785n ±  5%   51.175n ± 13%  +482.53% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            13.28n ±  6%    60.12n ± 10%  +352.92% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             9.080n ±  3%   56.115n ±  5%  +517.97% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             15.22n ± 24%    58.10n ± 16%  +281.61% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     14.54n ±  2%    31.23n ±  5%  +114.75% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     15.04n ±  3%    33.88n ±  2%  +125.23% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      14.32n ± 37%    32.47n ±  5%  +126.75% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      14.81n ±  1%    34.65n ±  7%  +133.96% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    14.37n ±  4%    37.61n ±  5%  +161.76% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    14.80n ±  3%    38.70n ±  5%  +161.52% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     14.45n ±  2%    39.81n ±  4%  +175.54% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     15.03n ±  5%    42.07n ±  6%  +179.94% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   7.897n ± 14%   39.590n ± 11%  +401.36% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   16.82n ± 43%    42.99n ±  3%  +155.51% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    14.34n ± 37%    45.45n ±  2%  +216.95% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    17.12n ±  3%    46.59n ±  5%  +172.09% (p=0.000 n=20)
geomean                                  13.44n          42.30n        +214.82%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ fast/lpm.bm  │             critbitgo/lpm.bm             │
                                       │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4-8            8.785n ±  5%   152.950n ±  4%  +1641.04% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            13.28n ±  6%    189.30n ±  3%  +1325.99% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             9.080n ±  3%   683.150n ± 18%  +7423.26% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             15.22n ± 24%    508.20n ± 18%  +3237.93% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     14.54n ±  2%     87.52n ±  4%   +501.68% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     15.04n ±  3%    103.30n ±  8%   +586.84% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      14.32n ± 37%    149.55n ± 15%   +944.34% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      14.81n ±  1%    153.50n ±  8%   +936.46% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    14.37n ±  4%    103.05n ±  5%   +617.12% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    14.80n ±  3%    114.10n ±  3%   +670.95% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     14.45n ±  2%    277.05n ± 12%  +1817.30% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     15.03n ±  5%    225.30n ± 18%  +1399.00% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   7.897n ± 14%   117.750n ±  4%  +1391.17% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   16.82n ± 43%    128.75n ±  3%   +665.23% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    14.34n ± 37%    464.90n ± 15%  +3141.98% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    17.12n ±  3%    343.90n ± 17%  +1908.18% (p=0.000 n=20)
geomean                                  13.44n           193.0n        +1336.62%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ fast/lpm.bm  │              lpmtrie/lpm.bm              │
                                       │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4-8            8.785n ±  5%   287.700n ±  9%  +3174.90% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            13.28n ±  6%    339.80n ± 17%  +2459.70% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             9.080n ±  3%   267.800n ±  5%  +2849.18% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             15.22n ± 24%    253.10n ± 14%  +1562.40% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     14.54n ±  2%    138.05n ±  6%   +849.12% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     15.04n ±  3%    127.75n ± 23%   +749.40% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      14.32n ± 37%    119.40n ± 19%   +733.80% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      14.81n ±  1%     99.20n ± 35%   +569.85% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    14.37n ±  4%    160.75n ± 13%  +1018.65% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    14.80n ±  3%    148.10n ±  8%   +900.68% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     14.45n ±  2%    169.20n ± 11%  +1070.93% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     15.03n ±  5%    142.95n ±  3%   +851.10% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   7.897n ± 14%   203.350n ±  6%  +2475.19% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   16.82n ± 43%    196.30n ±  5%  +1066.72% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    14.34n ± 37%    204.75n ±  2%  +1327.82% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    17.12n ±  3%    198.30n ±  3%  +1057.96% (p=0.000 n=20)
geomean                                  13.44n           180.7n        +1244.93%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ fast/lpm.bm  │          kentik-patricia/lpm.bm          │
                                       │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4-8            8.785n ±  5%   192.950n ±  7%  +2096.36% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            13.28n ±  6%    283.00n ± 19%  +2031.83% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             9.080n ±  3%   148.700n ±  5%  +1537.58% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             15.22n ± 24%    178.20n ± 11%  +1070.44% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     14.54n ±  2%     87.73n ±  7%   +503.20% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     15.04n ±  3%    107.15n ± 13%   +612.43% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      14.32n ± 37%     70.73n ±  7%   +393.96% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      14.81n ±  1%     84.11n ± 11%   +467.93% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    14.37n ±  4%    112.80n ±  5%   +684.97% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    14.80n ±  3%    133.10n ±  8%   +799.32% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     14.45n ±  2%     99.08n ±  3%   +585.64% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     15.03n ±  5%    118.50n ±  7%   +688.42% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   7.897n ± 14%   142.200n ±  9%  +1700.80% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   16.82n ± 43%    163.80n ±  4%   +873.55% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    14.34n ± 37%    114.95n ±  5%   +701.60% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    17.12n ±  3%    146.25n ±  4%   +754.01% (p=0.000 n=20)
geomean                                  13.44n           128.6n         +857.07%
```

## size of the routing tables


`bart.Lite` has the lowest memory consumption under all competitors.


```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/size.bm │            bart/size.bm            │
                           │ bytes/route  │ bytes/route  vs base               │
Tier1PfxSize/1_000-8           85.18 ± 2%   105.20 ± 2%  +23.50% (p=0.002 n=6)
Tier1PfxSize/10_000-8          63.83 ± 0%    83.84 ± 0%  +31.35% (p=0.002 n=6)
Tier1PfxSize/100_000-8         36.96 ± 0%    56.71 ± 0%  +53.44% (p=0.002 n=6)
Tier1PfxSize/200_000-8         30.07 ± 0%    49.53 ± 0%  +64.72% (p=0.002 n=6)
Tier1PfxSize/500_000-8         23.63 ± 0%    42.90 ± 0%  +81.55% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       20.42 ± 0%    39.74 ± 0%  +94.61% (p=0.002 n=6)
RandomPfx4Size/1_000-8         62.26 ± 3%    82.53 ± 2%  +32.56% (p=0.002 n=6)
RandomPfx4Size/10_000-8        38.19 ± 0%    57.26 ± 0%  +49.93% (p=0.002 n=6)
RandomPfx4Size/100_000-8       49.81 ± 0%    72.08 ± 0%  +44.71% (p=0.002 n=6)
RandomPfx4Size/200_000-8       43.92 ± 0%    65.15 ± 0%  +48.34% (p=0.002 n=6)
RandomPfx4Size/500_000-8       31.18 ± 0%    52.44 ± 0%  +68.18% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     25.83 ± 0%    46.46 ± 0%  +79.87% (p=0.002 n=6)
RandomPfx6Size/1_000-8         66.18 ± 3%    83.65 ± 2%  +26.40% (p=0.002 n=6)
RandomPfx6Size/10_000-8        80.69 ± 0%   101.30 ± 0%  +25.54% (p=0.002 n=6)
RandomPfx6Size/100_000-8       53.58 ± 0%    71.68 ± 0%  +33.78% (p=0.002 n=6)
RandomPfx6Size/200_000-8       51.40 ± 0%    69.37 ± 0%  +34.96% (p=0.002 n=6)
RandomPfx6Size/500_000-8       53.93 ± 0%    72.34 ± 0%  +34.14% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     59.52 ± 0%    78.92 ± 0%  +32.59% (p=0.002 n=6)
RandomPfxSize/1_000-8          80.61 ± 2%   100.60 ± 2%  +24.80% (p=0.002 n=6)
RandomPfxSize/10_000-8         56.42 ± 0%    74.00 ± 0%  +31.16% (p=0.002 n=6)
RandomPfxSize/100_000-8        63.45 ± 0%    84.33 ± 0%  +32.91% (p=0.002 n=6)
RandomPfxSize/200_000-8        54.19 ± 0%    75.59 ± 0%  +39.49% (p=0.002 n=6)
RandomPfxSize/500_000-8        39.18 ± 0%    59.59 ± 0%  +52.09% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      31.89 ± 0%    52.03 ± 0%  +63.15% (p=0.002 n=6)
geomean                        46.65         67.53       +44.76%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/size.bm │             fast/size.bm              │
                           │ bytes/route  │ bytes/route   vs base                 │
Tier1PfxSize/1_000-8           85.18 ± 2%   1267.00 ± 0%  +1387.44% (p=0.002 n=6)
Tier1PfxSize/10_000-8          63.83 ± 0%   1252.00 ± 0%  +1861.46% (p=0.002 n=6)
Tier1PfxSize/100_000-8         36.96 ± 0%    839.50 ± 0%  +2171.37% (p=0.002 n=6)
Tier1PfxSize/200_000-8         30.07 ± 0%    600.40 ± 0%  +1896.67% (p=0.002 n=6)
Tier1PfxSize/500_000-8         23.63 ± 0%    374.50 ± 0%  +1484.85% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       20.42 ± 0%    259.00 ± 0%  +1168.36% (p=0.002 n=6)
RandomPfx4Size/1_000-8         62.26 ± 3%   1206.00 ± 0%  +1837.04% (p=0.002 n=6)
RandomPfx4Size/10_000-8        38.19 ± 0%    326.30 ± 0%   +754.41% (p=0.002 n=6)
RandomPfx4Size/100_000-8       49.81 ± 0%   1122.00 ± 0%  +2152.56% (p=0.002 n=6)
RandomPfx4Size/200_000-8       43.92 ± 0%   1200.00 ± 0%  +2632.24% (p=0.002 n=6)
RandomPfx4Size/500_000-8       31.18 ± 0%    677.30 ± 0%  +2072.23% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     25.83 ± 0%    372.80 ± 0%  +1343.28% (p=0.002 n=6)
RandomPfx6Size/1_000-8         66.18 ± 3%    515.70 ± 0%   +679.24% (p=0.002 n=6)
RandomPfx6Size/10_000-8        80.69 ± 0%   1461.00 ± 0%  +1710.63% (p=0.002 n=6)
RandomPfx6Size/100_000-8       53.58 ± 0%    531.40 ± 0%   +891.79% (p=0.002 n=6)
RandomPfx6Size/200_000-8       51.40 ± 0%    403.00 ± 0%   +684.05% (p=0.002 n=6)
RandomPfx6Size/500_000-8       53.93 ± 0%    483.70 ± 0%   +796.90% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     59.52 ± 0%    730.50 ± 0%  +1127.32% (p=0.002 n=6)
RandomPfxSize/1_000-8          80.61 ± 2%   1275.00 ± 0%  +1481.69% (p=0.002 n=6)
RandomPfxSize/10_000-8         56.42 ± 0%    481.10 ± 0%   +752.71% (p=0.002 n=6)
RandomPfxSize/100_000-8        63.45 ± 0%   1340.00 ± 0%  +2011.90% (p=0.002 n=6)
RandomPfxSize/200_000-8        54.19 ± 0%   1306.00 ± 0%  +2310.04% (p=0.002 n=6)
RandomPfxSize/500_000-8        39.18 ± 0%    765.80 ± 0%  +1854.57% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      31.89 ± 0%    439.20 ± 0%  +1277.23% (p=0.002 n=6)
geomean                        46.65          703.2       +1407.44%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/size.bm │           netipds/size.bm           │
                           │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000-8           85.18 ± 2%    73.95 ± 3%   -13.18% (p=0.002 n=6)
Tier1PfxSize/10_000-8          63.83 ± 0%    68.88 ± 0%    +7.91% (p=0.002 n=6)
Tier1PfxSize/100_000-8         36.96 ± 0%    67.65 ± 0%   +83.04% (p=0.002 n=6)
Tier1PfxSize/200_000-8         30.07 ± 0%    66.90 ± 0%  +122.48% (p=0.002 n=6)
Tier1PfxSize/500_000-8         23.63 ± 0%    65.34 ± 0%  +176.51% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       20.42 ± 0%    63.72 ± 0%  +212.05% (p=0.002 n=6)
RandomPfx4Size/1_000-8         62.26 ± 3%    62.35 ± 3%    +0.14% (p=0.013 n=6)
RandomPfx4Size/10_000-8        38.19 ± 0%    51.25 ± 0%   +34.20% (p=0.002 n=6)
RandomPfx4Size/100_000-8       49.81 ± 0%    48.24 ± 0%    -3.15% (p=0.002 n=6)
RandomPfx4Size/200_000-8       43.92 ± 0%    47.57 ± 0%    +8.31% (p=0.002 n=6)
RandomPfx4Size/500_000-8       31.18 ± 0%    46.51 ± 0%   +49.17% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     25.83 ± 0%    45.65 ± 0%   +76.73% (p=0.002 n=6)
RandomPfx6Size/1_000-8         66.18 ± 3%   100.10 ± 2%   +51.25% (p=0.002 n=6)
RandomPfx6Size/10_000-8        80.69 ± 0%    94.17 ± 0%   +16.71% (p=0.002 n=6)
RandomPfx6Size/100_000-8       53.58 ± 0%    87.63 ± 0%   +63.55% (p=0.002 n=6)
RandomPfx6Size/200_000-8       51.40 ± 0%    85.90 ± 0%   +67.12% (p=0.002 n=6)
RandomPfx6Size/500_000-8       53.93 ± 0%    84.11 ± 0%   +55.96% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     59.52 ± 0%    83.23 ± 0%   +39.84% (p=0.002 n=6)
RandomPfxSize/1_000-8          80.61 ± 2%    74.82 ± 3%    -7.18% (p=0.002 n=6)
RandomPfxSize/10_000-8         56.42 ± 0%    69.29 ± 0%   +22.81% (p=0.002 n=6)
RandomPfxSize/100_000-8        63.45 ± 0%    64.28 ± 0%    +1.31% (p=0.002 n=6)
RandomPfxSize/200_000-8        54.19 ± 0%    61.55 ± 0%   +13.58% (p=0.002 n=6)
RandomPfxSize/500_000-8        39.18 ± 0%    57.42 ± 0%   +46.55% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      31.89 ± 0%    53.93 ± 0%   +69.11% (p=0.002 n=6)
geomean                        46.65         66.03        +41.55%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/size.bm │          critbitgo/size.bm          │
                           │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000-8           85.18 ± 2%   119.50 ± 2%   +40.29% (p=0.002 n=6)
Tier1PfxSize/10_000-8          63.83 ± 0%   114.70 ± 0%   +79.70% (p=0.002 n=6)
Tier1PfxSize/100_000-8         36.96 ± 0%   114.40 ± 0%  +209.52% (p=0.002 n=6)
Tier1PfxSize/200_000-8         30.07 ± 0%   114.40 ± 0%  +280.45% (p=0.002 n=6)
Tier1PfxSize/500_000-8         23.63 ± 0%   114.40 ± 0%  +384.13% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       20.42 ± 0%   114.40 ± 0%  +460.24% (p=0.002 n=6)
RandomPfx4Size/1_000-8         62.26 ± 3%   116.40 ± 2%   +86.96% (p=0.002 n=6)
RandomPfx4Size/10_000-8        38.19 ± 0%   112.30 ± 0%  +194.06% (p=0.002 n=6)
RandomPfx4Size/100_000-8       49.81 ± 0%   112.00 ± 0%  +124.85% (p=0.002 n=6)
RandomPfx4Size/200_000-8       43.92 ± 0%   112.00 ± 0%  +155.01% (p=0.002 n=6)
RandomPfx4Size/500_000-8       31.18 ± 0%   112.00 ± 0%  +259.20% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     25.83 ± 0%   112.00 ± 0%  +333.60% (p=0.002 n=6)
RandomPfx6Size/1_000-8         66.18 ± 3%   132.30 ± 1%   +99.91% (p=0.002 n=6)
RandomPfx6Size/10_000-8        80.69 ± 0%   128.30 ± 0%   +59.00% (p=0.002 n=6)
RandomPfx6Size/100_000-8       53.58 ± 0%   128.00 ± 0%  +138.90% (p=0.002 n=6)
RandomPfx6Size/200_000-8       51.40 ± 0%   128.00 ± 0%  +149.03% (p=0.002 n=6)
RandomPfx6Size/500_000-8       53.93 ± 0%   128.00 ± 0%  +137.34% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     59.52 ± 0%   128.00 ± 0%  +115.05% (p=0.002 n=6)
RandomPfxSize/1_000-8          80.61 ± 2%   120.20 ± 2%   +49.11% (p=0.002 n=6)
RandomPfxSize/10_000-8         56.42 ± 0%   115.50 ± 0%  +104.71% (p=0.002 n=6)
RandomPfxSize/100_000-8        63.45 ± 0%   115.30 ± 0%   +81.72% (p=0.002 n=6)
RandomPfxSize/200_000-8        54.19 ± 0%   115.20 ± 0%  +112.59% (p=0.002 n=6)
RandomPfxSize/500_000-8        39.18 ± 0%   115.20 ± 0%  +194.03% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      31.89 ± 0%   115.20 ± 0%  +261.24% (p=0.002 n=6)
geomean                        46.65         118.1       +153.11%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/size.bm │           lpmtrie/size.bm            │
                           │ bytes/route  │ bytes/route   vs base                │
Tier1PfxSize/1_000-8           85.18 ± 2%   215.50 ±  5%  +152.99% (p=0.002 n=6)
Tier1PfxSize/10_000-8          63.83 ± 0%   210.50 ±  5%  +229.78% (p=0.002 n=6)
Tier1PfxSize/100_000-8         36.96 ± 0%   209.90 ±  5%  +467.91% (p=0.002 n=6)
Tier1PfxSize/200_000-8         30.07 ± 0%   209.20 ±  5%  +595.71% (p=0.002 n=6)
Tier1PfxSize/500_000-8         23.63 ± 0%   207.90 ±  7%  +779.81% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       20.42 ± 0%   205.00 ±  7%  +903.92% (p=0.002 n=6)
RandomPfx4Size/1_000-8         62.26 ± 3%   205.40 ±  8%  +229.91% (p=0.002 n=6)
RandomPfx4Size/10_000-8        38.19 ± 0%   186.50 ±  9%  +388.35% (p=0.002 n=6)
RandomPfx4Size/100_000-8       49.81 ± 0%   179.60 ±  9%  +260.57% (p=0.002 n=6)
RandomPfx4Size/200_000-8       43.92 ± 0%   178.50 ± 10%  +306.42% (p=0.002 n=6)
RandomPfx4Size/500_000-8       31.18 ± 0%   176.80 ± 10%  +467.03% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     25.83 ± 0%   175.50 ± 10%  +579.44% (p=0.002 n=6)
RandomPfx6Size/1_000-8         66.18 ± 3%   227.90 ±  8%  +244.36% (p=0.002 n=6)
RandomPfx6Size/10_000-8        80.69 ± 0%   222.40 ±  8%  +175.62% (p=0.002 n=6)
RandomPfx6Size/100_000-8       53.58 ± 0%   213.90 ±  9%  +299.22% (p=0.002 n=6)
RandomPfx6Size/200_000-8       51.40 ± 0%   210.50 ±  9%  +309.53% (p=0.002 n=6)
RandomPfx6Size/500_000-8       53.93 ± 0%   206.70 ±  9%  +283.27% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     59.52 ± 0%   204.90 ±  9%  +244.25% (p=0.002 n=6)
RandomPfxSize/1_000-8          80.61 ± 2%   215.20 ±  5%  +166.96% (p=0.002 n=6)
RandomPfxSize/10_000-8         56.42 ± 0%   210.20 ±  5%  +272.56% (p=0.002 n=6)
RandomPfxSize/100_000-8        63.45 ± 0%   203.50 ±  7%  +220.72% (p=0.002 n=6)
RandomPfxSize/200_000-8        54.19 ± 0%   198.70 ±  8%  +266.67% (p=0.002 n=6)
RandomPfxSize/500_000-8        39.18 ± 0%   189.90 ±  9%  +384.69% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      31.89 ± 0%   181.40 ± 10%  +468.83% (p=0.002 n=6)
geomean                        46.65         201.3        +331.63%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/size.bm │       kentik-patricia/size.bm       │
                           │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000-8           85.18 ± 2%   145.30 ± 1%   +70.58% (p=0.002 n=6)
Tier1PfxSize/10_000-8          63.83 ± 0%   200.20 ± 0%  +213.65% (p=0.002 n=6)
Tier1PfxSize/100_000-8         36.96 ± 0%   164.00 ± 0%  +343.72% (p=0.002 n=6)
Tier1PfxSize/200_000-8         30.07 ± 0%   163.90 ± 0%  +445.06% (p=0.002 n=6)
Tier1PfxSize/500_000-8         23.63 ± 0%   144.30 ± 0%  +510.66% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       20.42 ± 0%   144.20 ± 0%  +606.17% (p=0.002 n=6)
RandomPfx4Size/1_000-8         62.26 ± 3%   140.80 ± 1%  +126.15% (p=0.002 n=6)
RandomPfx4Size/10_000-8        38.19 ± 0%   109.60 ± 0%  +186.99% (p=0.002 n=6)
RandomPfx4Size/100_000-8       49.81 ± 0%   139.80 ± 0%  +180.67% (p=0.002 n=6)
RandomPfx4Size/200_000-8       43.92 ± 0%   139.80 ± 0%  +218.31% (p=0.002 n=6)
RandomPfx4Size/500_000-8       31.18 ± 0%   139.80 ± 0%  +348.36% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     25.83 ± 0%   139.60 ± 0%  +440.46% (p=0.002 n=6)
RandomPfx6Size/1_000-8         66.18 ± 3%   157.70 ± 1%  +138.29% (p=0.002 n=6)
RandomPfx6Size/10_000-8        80.69 ± 0%   201.30 ± 0%  +149.47% (p=0.002 n=6)
RandomPfx6Size/100_000-8       53.58 ± 0%   160.80 ± 0%  +200.11% (p=0.002 n=6)
RandomPfx6Size/200_000-8       51.40 ± 0%   160.80 ± 0%  +212.84% (p=0.002 n=6)
RandomPfx6Size/500_000-8       53.93 ± 0%   156.50 ± 0%  +190.19% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     59.52 ± 0%   156.50 ± 0%  +162.94% (p=0.002 n=6)
RandomPfxSize/1_000-8          80.61 ± 2%   144.60 ± 1%   +79.38% (p=0.002 n=6)
RandomPfxSize/10_000-8         56.42 ± 0%   140.10 ± 0%  +148.32% (p=0.002 n=6)
RandomPfxSize/100_000-8        63.45 ± 0%   180.00 ± 0%  +183.69% (p=0.002 n=6)
RandomPfxSize/200_000-8        54.19 ± 0%   180.00 ± 0%  +232.16% (p=0.002 n=6)
RandomPfxSize/500_000-8        39.18 ± 0%   144.00 ± 0%  +267.53% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      31.89 ± 0%   144.00 ± 0%  +351.55% (p=0.002 n=6)
geomean                        46.65         152.8       +227.55%
```

## update, insert/delete

`bart.Lite` is the fastest algorithm for updates.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/update.bm │           bart/update.bm           │
                           │   sec/route    │  sec/route   vs base               │
InsertRandomPfxs/1_000-8        154.3n ± 3%   181.0n ± 4%  +17.34% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       112.8n ± 4%   140.5n ± 3%  +24.51% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      167.8n ± 3%   215.6n ± 2%  +28.55% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      195.8n ± 2%   263.4n ± 2%  +34.53% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        116.2n ± 1%   126.4n ± 1%   +8.73% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       80.36n ± 2%   93.70n ± 1%  +16.61% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      173.2n ± 3%   206.1n ± 2%  +19.00% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      209.2n ± 3%   266.5n ± 1%  +27.36% (p=0.002 n=6)
geomean                         144.8n        176.4n       +21.84%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/update.bm │            fast/update.bm            │
                           │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000-8        154.3n ± 3%    400.2n ± 2%  +159.36% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       112.8n ± 4%    219.4n ± 2%   +94.50% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      167.8n ± 3%    526.4n ± 7%  +213.80% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      195.8n ± 2%    667.1n ± 8%  +240.70% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        116.2n ± 1%    172.4n ± 2%   +48.30% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       80.36n ± 2%   119.45n ± 1%   +48.64% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      173.2n ± 3%    277.2n ± 4%   +60.09% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      209.2n ± 3%    355.9n ± 8%   +70.08% (p=0.002 n=6)
geomean                         144.8n         298.3n       +105.96%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/update.bm │          netipds/update.bm           │
                           │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000-8        154.3n ± 3%    248.6n ± 3%   +61.11% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       112.8n ± 4%    313.1n ± 5%  +177.57% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      167.8n ± 3%    467.5n ± 3%  +178.69% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      195.8n ± 2%    527.9n ± 8%  +169.59% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        116.2n ± 1%    179.2n ± 7%   +54.19% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       80.36n ± 2%   225.05n ± 3%  +180.05% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      173.2n ± 3%    386.2n ± 5%  +123.04% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      209.2n ± 3%    496.4n ± 3%  +137.25% (p=0.002 n=6)
geomean                         144.8n         332.3n       +129.43%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/update.bm │         critbitgo/update.bm          │
                           │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000-8        154.3n ± 3%    322.6n ± 3%  +109.11% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       112.8n ± 4%    419.8n ± 3%  +272.12% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      167.8n ± 3%    673.2n ± 1%  +301.34% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      195.8n ± 2%    824.8n ± 5%  +321.22% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        116.2n ± 1%    164.9n ± 2%   +41.85% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       80.36n ± 2%   230.05n ± 3%  +186.27% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      173.2n ± 3%    442.5n ± 2%  +155.56% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      209.2n ± 3%    539.3n ± 2%  +157.73% (p=0.002 n=6)
geomean                         144.8n         401.9n       +177.52%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/update.bm │          lpmtrie/update.bm           │
                           │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000-8        154.3n ± 3%    469.4n ± 5%  +204.21% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       112.8n ± 4%    558.0n ± 7%  +394.64% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      167.8n ± 3%    821.0n ± 4%  +389.39% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      195.8n ± 2%    977.1n ± 4%  +399.03% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        116.2n ± 1%    163.1n ± 1%   +40.26% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       80.36n ± 2%   266.25n ± 2%  +231.32% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      173.2n ± 3%    568.9n ± 1%  +228.56% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      209.2n ± 3%    719.1n ± 3%  +243.68% (p=0.002 n=6)
geomean                         144.8n         497.1n       +243.28%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/update.bm │      kentik-patricia/update.bm       │
                           │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000-8        154.3n ± 3%    256.2n ± 2%   +66.07% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       112.8n ± 4%    359.4n ± 3%  +218.66% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      167.8n ± 3%    592.7n ± 3%  +253.32% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      195.8n ± 2%    740.3n ± 7%  +278.09% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        116.2n ± 1%    332.9n ± 1%  +186.32% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       80.36n ± 2%   415.70n ± 3%  +417.30% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      173.2n ± 3%    780.5n ± 4%  +350.77% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      209.2n ± 3%    928.1n ± 1%  +343.56% (p=0.002 n=6)
geomean                         144.8n         502.3n       +246.83%
```
