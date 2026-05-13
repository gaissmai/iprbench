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
of 4 parts IPv4 to 1 part IPv6 random prefixes, which is approximately the current ratio in the
Internet backbone routers.

## make your own benchmarks

```
  $ # IMPORTANT: set the proper cpu feature flags for amd64, e.g.
  $ export GOAMD64=v3

  $ make dep
  $ make -B all   # takes some time!
```

## lpm (longest-prefix-match)

`bart.Fast` is by far the fastest software algorithm for IP-address longest-prefix-match.

```
goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                        │ fast/lpm.bm  │              lite/lpm.bm               │
                                        │    sec/op    │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4-16            6.571n ±  8%   12.435n ± 11%   +89.26% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-16            9.327n ± 27%   19.470n ± 19%  +108.75% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-16             6.756n ±  7%   12.855n ±  3%   +90.26% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-16             11.73n ± 24%    19.12n ± 31%   +62.96% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-16     9.947n ±  5%   13.115n ±  8%   +31.85% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-16     9.957n ±  6%   15.340n ±  5%   +54.07% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-16      9.521n ± 28%   12.595n ± 23%   +32.29% (p=0.001 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-16      10.21n ±  5%    16.01n ±  5%   +56.68% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-16    9.651n ±  7%   14.280n ±  7%   +47.97% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-16    10.24n ±  6%    15.38n ±  7%   +50.07% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-16     9.665n ±  7%   14.175n ±  2%   +46.67% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-16     10.32n ±  5%    15.87n ±  5%   +53.78% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-16   5.428n ± 23%    9.559n ± 41%   +76.08% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-16   11.45n ± 39%    18.54n ± 39%   +61.97% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-16    9.304n ± 25%   14.200n ±  6%   +52.61% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-16    12.08n ±  6%    19.18n ±  4%   +58.73% (p=0.000 n=20)
geomean                                   9.317n          14.88n         +59.67%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                        │ fast/lpm.bm  │              bart/lpm.bm               │
                                        │    sec/op    │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4-16            6.571n ±  8%   12.710n ±  5%   +93.44% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-16            9.327n ± 27%   19.300n ± 24%  +106.93% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-16             6.756n ±  7%   13.650n ±  5%  +102.03% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-16             11.73n ± 24%    19.28n ± 34%   +64.36% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-16     9.947n ±  5%   13.350n ±  5%   +34.21% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-16     9.957n ±  6%   14.030n ±  6%   +40.91% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-16      9.521n ± 28%   12.575n ± 26%   +32.08% (p=0.002 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-16      10.21n ±  5%    14.89n ±  6%   +45.81% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-16    9.651n ±  7%   15.230n ±  5%   +57.82% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-16    10.24n ±  6%    14.50n ±  5%   +41.53% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-16     9.665n ±  7%   15.435n ±  7%   +59.71% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-16     10.32n ±  5%    14.53n ±  5%   +40.75% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-16   5.428n ± 23%    9.733n ± 38%   +79.30% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-16   11.45n ± 39%    18.29n ± 36%   +59.69% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-16    9.304n ± 25%   14.145n ±  9%   +52.02% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-16    12.08n ±  6%    18.64n ±  8%   +54.30% (p=0.000 n=20)
geomean                                   9.317n          14.79n         +58.78%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                        │ fast/lpm.bm  │             netipds/lpm.bm             │
                                        │    sec/op    │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4-16            6.571n ±  8%   33.980n ± 19%  +417.16% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-16            9.327n ± 27%   34.435n ± 17%  +269.20% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-16             6.756n ±  7%   37.455n ±  6%  +454.36% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-16             11.73n ± 24%    38.62n ±  9%  +229.20% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-16     9.947n ±  5%   22.270n ±  7%  +123.89% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-16     9.957n ±  6%   22.960n ±  4%  +130.60% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-16      9.521n ± 28%   22.740n ±  4%  +138.85% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-16      10.21n ±  5%    23.34n ±  4%  +128.49% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-16    9.651n ±  7%   25.780n ±  3%  +167.14% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-16    10.24n ±  6%    27.55n ±  3%  +168.96% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-16     9.665n ±  7%   27.160n ±  3%  +181.03% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-16     10.32n ±  5%    28.14n ±  4%  +172.67% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-16   5.428n ± 23%   26.745n ± 14%  +392.68% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-16   11.45n ± 39%    30.19n ±  3%  +163.71% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-16    9.304n ± 25%   30.990n ±  5%  +233.06% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-16    12.08n ±  6%    32.18n ±  7%  +166.39% (p=0.000 n=20)
geomean                                   9.317n          28.60n        +207.02%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                        │ fast/lpm.bm  │             critbitgo/lpm.bm             │
                                        │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4-16            6.571n ±  8%   127.150n ±  7%  +1835.16% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-16            9.327n ± 27%   152.250n ±  7%  +1532.36% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-16             6.756n ±  7%   479.800n ± 21%  +7001.31% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-16             11.73n ± 24%    354.20n ± 22%  +2919.61% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-16     9.947n ±  5%    82.965n ±  8%   +734.07% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-16     9.957n ±  6%    93.810n ±  5%   +842.20% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-16      9.521n ± 28%   136.850n ± 11%  +1337.42% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-16      10.21n ±  5%    143.65n ±  8%  +1306.27% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-16    9.651n ±  7%    98.390n ±  6%   +919.53% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-16    10.24n ±  6%    110.20n ±  4%   +975.65% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-16     9.665n ±  7%   218.950n ± 10%  +2165.51% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-16     10.32n ±  5%    190.70n ± 15%  +1747.87% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-16   5.428n ± 23%   106.350n ±  4%  +1859.10% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-16   11.45n ± 39%    119.90n ±  3%   +947.16% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-16    9.304n ± 25%   309.500n ± 12%  +3226.35% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-16    12.08n ±  6%    269.45n ± 13%  +2130.55% (p=0.000 n=20)
geomean                                   9.317n           162.8n        +1646.87%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                        │ fast/lpm.bm  │              lpmtrie/lpm.bm              │
                                        │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4-16            6.571n ±  8%   218.600n ±  7%  +3226.99% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-16            9.327n ± 27%   235.150n ± 15%  +2421.18% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-16             6.756n ±  7%   202.100n ±  8%  +2891.19% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-16             11.73n ± 24%    181.85n ± 13%  +1450.30% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-16     9.947n ±  5%    95.855n ±  5%   +863.66% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-16     9.957n ±  6%    84.450n ± 11%   +748.19% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-16      9.521n ± 28%    84.900n ±  9%   +791.76% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-16      10.21n ±  5%     74.36n ± 16%   +627.95% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-16    9.651n ±  7%   124.400n ± 17%  +1189.05% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-16    10.24n ±  6%    110.15n ± 10%   +975.16% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-16     9.665n ±  7%   128.150n ± 13%  +1225.99% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-16     10.32n ±  5%    102.55n ±  5%   +893.70% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-16   5.428n ± 23%   162.750n ±  3%  +2898.07% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-16   11.45n ± 39%    148.55n ±  3%  +1197.38% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-16    9.304n ± 25%   150.450n ±  6%  +1516.96% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-16    12.08n ±  6%    148.70n ±  5%  +1130.96% (p=0.000 n=20)
geomean                                   9.317n           132.8n        +1325.83%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                        │ fast/lpm.bm  │          kentik-patricia/lpm.bm          │
                                        │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4-16            6.571n ±  8%   134.100n ± 12%  +1940.94% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-16            9.327n ± 27%   185.600n ± 15%  +1889.92% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-16             6.756n ±  7%   107.900n ±  6%  +1496.98% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-16             11.73n ± 24%    123.35n ± 15%   +951.58% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-16     9.947n ±  5%    64.685n ± 13%   +550.30% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-16     9.957n ±  6%    72.980n ±  8%   +632.99% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-16      9.521n ± 28%    53.665n ±  8%   +463.68% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-16      10.21n ±  5%     57.16n ± 12%   +459.57% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-16    9.651n ±  7%    86.525n ±  4%   +796.59% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-16    10.24n ±  6%     93.00n ± 13%   +807.81% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-16     9.665n ±  7%    74.455n ±  6%   +670.40% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-16     10.32n ±  5%     80.81n ±  4%   +682.99% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-16   5.428n ± 23%   110.350n ± 10%  +1932.79% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-16   11.45n ± 39%    111.65n ±  7%   +875.11% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-16    9.304n ± 25%    86.975n ±  6%   +834.76% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-16    12.08n ±  6%    100.61n ±  7%   +732.86% (p=0.000 n=20)
geomean                                   9.317n           91.71n         +884.35%
```

## size of the routing tables


`bart.Lite` has the lowest memory consumption under all competitors.


```
goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                            │ lite/size.bm │            bart/size.bm            │
                            │ bytes/route  │ bytes/route  vs base               │
Tier1PfxSize/1_000-16           90.94 ± 2%   105.40 ± 2%  +15.90% (p=0.002 n=6)
Tier1PfxSize/10_000-16          63.79 ± 0%    83.79 ± 0%  +31.35% (p=0.002 n=6)
Tier1PfxSize/100_000-16         36.95 ± 0%    56.70 ± 0%  +53.45% (p=0.002 n=6)
Tier1PfxSize/200_000-16         30.07 ± 0%    49.53 ± 0%  +64.72% (p=0.002 n=6)
Tier1PfxSize/500_000-16         23.63 ± 0%    42.90 ± 0%  +81.55% (p=0.002 n=6)
Tier1PfxSize/1_000_000-16       20.42 ± 0%    39.74 ± 0%  +94.61% (p=0.002 n=6)
RandomPfx4Size/1_000-16         61.94 ± 3%    82.66 ± 2%  +33.45% (p=0.002 n=6)
RandomPfx4Size/10_000-16        38.09 ± 1%    57.16 ± 0%  +50.07% (p=0.002 n=6)
RandomPfx4Size/100_000-16       49.80 ± 0%    72.07 ± 0%  +44.72% (p=0.002 n=6)
RandomPfx4Size/200_000-16       43.91 ± 0%    65.14 ± 0%  +48.35% (p=0.002 n=6)
RandomPfx4Size/500_000-16       31.18 ± 0%    52.44 ± 0%  +68.18% (p=0.002 n=6)
RandomPfx4Size/1_000_000-16     25.83 ± 0%    46.46 ± 0%  +79.87% (p=0.002 n=6)
RandomPfx6Size/1_000-16         66.32 ± 3%    83.78 ± 2%  +26.33% (p=0.002 n=6)
RandomPfx6Size/10_000-16        80.59 ± 0%   101.20 ± 0%  +25.57% (p=0.002 n=6)
RandomPfx6Size/100_000-16       53.57 ± 0%    71.67 ± 0%  +33.79% (p=0.002 n=6)
RandomPfx6Size/200_000-16       51.39 ± 0%    69.37 ± 0%  +34.99% (p=0.002 n=6)
RandomPfx6Size/500_000-16       53.92 ± 0%    72.34 ± 0%  +34.16% (p=0.002 n=6)
RandomPfx6Size/1_000_000-16     59.52 ± 0%    78.92 ± 0%  +32.59% (p=0.002 n=6)
RandomPfxSize/1_000-16          80.74 ± 3%   100.70 ± 2%  +24.72% (p=0.002 n=6)
RandomPfxSize/10_000-16         56.31 ± 0%    73.94 ± 0%  +31.31% (p=0.002 n=6)
RandomPfxSize/100_000-16        63.44 ± 0%    84.32 ± 0%  +32.91% (p=0.002 n=6)
RandomPfxSize/200_000-16        54.19 ± 0%    75.59 ± 0%  +39.49% (p=0.002 n=6)
RandomPfxSize/500_000-16        39.17 ± 0%    59.59 ± 0%  +52.13% (p=0.002 n=6)
RandomPfxSize/1_000_000-16      31.88 ± 0%    52.03 ± 0%  +63.21% (p=0.002 n=6)
geomean                         46.76         67.53       +44.43%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                            │ lite/size.bm │             fast/size.bm              │
                            │ bytes/route  │ bytes/route   vs base                 │
Tier1PfxSize/1_000-16           90.94 ± 2%   1267.00 ± 0%  +1293.23% (p=0.002 n=6)
Tier1PfxSize/10_000-16          63.79 ± 0%   1252.00 ± 0%  +1862.69% (p=0.002 n=6)
Tier1PfxSize/100_000-16         36.95 ± 0%    839.50 ± 0%  +2171.99% (p=0.002 n=6)
Tier1PfxSize/200_000-16         30.07 ± 0%    600.40 ± 0%  +1896.67% (p=0.002 n=6)
Tier1PfxSize/500_000-16         23.63 ± 0%    374.50 ± 0%  +1484.85% (p=0.002 n=6)
Tier1PfxSize/1_000_000-16       20.42 ± 0%    259.00 ± 0%  +1168.36% (p=0.002 n=6)
RandomPfx4Size/1_000-16         61.94 ± 3%   1205.00 ± 0%  +1845.43% (p=0.002 n=6)
RandomPfx4Size/10_000-16        38.09 ± 1%    326.20 ± 0%   +756.39% (p=0.002 n=6)
RandomPfx4Size/100_000-16       49.80 ± 0%   1122.00 ± 0%  +2153.01% (p=0.002 n=6)
RandomPfx4Size/200_000-16       43.91 ± 0%   1200.00 ± 0%  +2632.86% (p=0.002 n=6)
RandomPfx4Size/500_000-16       31.18 ± 0%    677.30 ± 0%  +2072.23% (p=0.002 n=6)
RandomPfx4Size/1_000_000-16     25.83 ± 0%    372.80 ± 0%  +1343.28% (p=0.002 n=6)
RandomPfx6Size/1_000-16         66.32 ± 3%    515.40 ± 0%   +677.14% (p=0.002 n=6)
RandomPfx6Size/10_000-16        80.59 ± 0%   1461.00 ± 0%  +1712.88% (p=0.002 n=6)
RandomPfx6Size/100_000-16       53.57 ± 0%    531.40 ± 0%   +891.97% (p=0.002 n=6)
RandomPfx6Size/200_000-16       51.39 ± 0%    403.00 ± 0%   +684.20% (p=0.002 n=6)
RandomPfx6Size/500_000-16       53.92 ± 0%    483.70 ± 0%   +797.07% (p=0.002 n=6)
RandomPfx6Size/1_000_000-16     59.52 ± 0%    730.50 ± 0%  +1127.32% (p=0.002 n=6)
RandomPfxSize/1_000-16          80.74 ± 3%   1275.00 ± 0%  +1479.14% (p=0.002 n=6)
RandomPfxSize/10_000-16         56.31 ± 0%    481.00 ± 0%   +754.20% (p=0.002 n=6)
RandomPfxSize/100_000-16        63.44 ± 0%   1340.00 ± 0%  +2012.23% (p=0.002 n=6)
RandomPfxSize/200_000-16        54.19 ± 0%   1306.00 ± 0%  +2310.04% (p=0.002 n=6)
RandomPfxSize/500_000-16        39.17 ± 0%    765.80 ± 0%  +1855.07% (p=0.002 n=6)
RandomPfxSize/1_000_000-16      31.88 ± 0%    439.20 ± 0%  +1277.67% (p=0.002 n=6)
geomean                         46.76          703.1       +1403.83%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                            │ lite/size.bm │           netipds/size.bm           │
                            │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000-16           90.94 ± 2%    74.24 ± 3%   -18.36% (p=0.002 n=6)
Tier1PfxSize/10_000-16          63.79 ± 0%    68.73 ± 0%    +7.74% (p=0.002 n=6)
Tier1PfxSize/100_000-16         36.95 ± 0%    67.64 ± 0%   +83.06% (p=0.002 n=6)
Tier1PfxSize/200_000-16         30.07 ± 0%    66.90 ± 0%  +122.48% (p=0.002 n=6)
Tier1PfxSize/500_000-16         23.63 ± 0%    65.34 ± 0%  +176.51% (p=0.002 n=6)
Tier1PfxSize/1_000_000-16       20.42 ± 0%    63.72 ± 0%  +212.05% (p=0.002 n=6)
RandomPfx4Size/1_000-16         61.94 ± 3%    62.51 ± 3%    +0.92% (p=0.013 n=6)
RandomPfx4Size/10_000-16        38.09 ± 1%    51.21 ± 0%   +34.44% (p=0.002 n=6)
RandomPfx4Size/100_000-16       49.80 ± 0%    48.23 ± 0%    -3.15% (p=0.002 n=6)
RandomPfx4Size/200_000-16       43.91 ± 0%    47.56 ± 0%    +8.31% (p=0.002 n=6)
RandomPfx4Size/500_000-16       31.18 ± 0%    46.51 ± 0%   +49.17% (p=0.002 n=6)
RandomPfx4Size/1_000_000-16     25.83 ± 0%    45.65 ± 0%   +76.73% (p=0.002 n=6)
RandomPfx6Size/1_000-16         66.32 ± 3%   100.30 ± 2%   +51.24% (p=0.002 n=6)
RandomPfx6Size/10_000-16        80.59 ± 0%    94.07 ± 0%   +16.73% (p=0.002 n=6)
RandomPfx6Size/100_000-16       53.57 ± 0%    87.62 ± 0%   +63.56% (p=0.002 n=6)
RandomPfx6Size/200_000-16       51.39 ± 0%    85.89 ± 0%   +67.13% (p=0.002 n=6)
RandomPfx6Size/500_000-16       53.92 ± 0%    84.11 ± 0%   +55.99% (p=0.002 n=6)
RandomPfx6Size/1_000_000-16     59.52 ± 0%    83.23 ± 0%   +39.84% (p=0.002 n=6)
RandomPfxSize/1_000-16          80.74 ± 3%    74.93 ± 3%    -7.20% (p=0.002 n=6)
RandomPfxSize/10_000-16         56.31 ± 0%    69.19 ± 0%   +22.87% (p=0.002 n=6)
RandomPfxSize/100_000-16        63.44 ± 0%    64.27 ± 0%    +1.31% (p=0.002 n=6)
RandomPfxSize/200_000-16        54.19 ± 0%    61.54 ± 0%   +13.56% (p=0.002 n=6)
RandomPfxSize/500_000-16        39.17 ± 0%    57.42 ± 0%   +46.59% (p=0.002 n=6)
RandomPfxSize/1_000_000-16      31.88 ± 0%    53.93 ± 0%   +69.17% (p=0.002 n=6)
geomean                         46.76         66.04        +41.24%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                            │ lite/size.bm │          critbitgo/size.bm          │
                            │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000-16           90.94 ± 2%   119.70 ± 2%   +31.63% (p=0.002 n=6)
Tier1PfxSize/10_000-16          63.79 ± 0%   114.60 ± 0%   +79.65% (p=0.002 n=6)
Tier1PfxSize/100_000-16         36.95 ± 0%   114.40 ± 0%  +209.61% (p=0.002 n=6)
Tier1PfxSize/200_000-16         30.07 ± 0%   114.40 ± 0%  +280.45% (p=0.002 n=6)
Tier1PfxSize/500_000-16         23.63 ± 0%   114.40 ± 0%  +384.13% (p=0.002 n=6)
Tier1PfxSize/1_000_000-16       20.42 ± 0%   114.40 ± 0%  +460.24% (p=0.002 n=6)
RandomPfx4Size/1_000-16         61.94 ± 3%   116.50 ± 2%   +88.09% (p=0.002 n=6)
RandomPfx4Size/10_000-16        38.09 ± 1%   112.20 ± 0%  +194.57% (p=0.002 n=6)
RandomPfx4Size/100_000-16       49.80 ± 0%   112.00 ± 0%  +124.90% (p=0.002 n=6)
RandomPfx4Size/200_000-16       43.91 ± 0%   112.00 ± 0%  +155.07% (p=0.002 n=6)
RandomPfx4Size/500_000-16       31.18 ± 0%   112.00 ± 0%  +259.20% (p=0.002 n=6)
RandomPfx4Size/1_000_000-16     25.83 ± 0%   112.00 ± 0%  +333.60% (p=0.002 n=6)
RandomPfx6Size/1_000-16         66.32 ± 3%   132.50 ± 2%   +99.79% (p=0.002 n=6)
RandomPfx6Size/10_000-16        80.59 ± 0%   128.20 ± 0%   +59.08% (p=0.002 n=6)
RandomPfx6Size/100_000-16       53.57 ± 0%   128.00 ± 0%  +138.94% (p=0.002 n=6)
RandomPfx6Size/200_000-16       51.39 ± 0%   128.00 ± 0%  +149.08% (p=0.002 n=6)
RandomPfx6Size/500_000-16       53.92 ± 0%   128.00 ± 0%  +137.39% (p=0.002 n=6)
RandomPfx6Size/1_000_000-16     59.52 ± 0%   128.00 ± 0%  +115.05% (p=0.002 n=6)
RandomPfxSize/1_000-16          80.74 ± 3%   119.80 ± 2%   +48.38% (p=0.002 n=6)
RandomPfxSize/10_000-16         56.31 ± 0%   115.40 ± 0%  +104.94% (p=0.002 n=6)
RandomPfxSize/100_000-16        63.44 ± 0%   115.30 ± 0%   +81.75% (p=0.002 n=6)
RandomPfxSize/200_000-16        54.19 ± 0%   115.20 ± 0%  +112.59% (p=0.002 n=6)
RandomPfxSize/500_000-16        39.17 ± 0%   115.20 ± 0%  +194.10% (p=0.002 n=6)
RandomPfxSize/1_000_000-16      31.88 ± 0%   115.20 ± 0%  +261.36% (p=0.002 n=6)
geomean                         46.76         118.1       +152.49%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                            │ lite/size.bm │           lpmtrie/size.bm            │
                            │ bytes/route  │ bytes/route   vs base                │
Tier1PfxSize/1_000-16           90.94 ± 2%   215.60 ±  5%  +137.08% (p=0.002 n=6)
Tier1PfxSize/10_000-16          63.79 ± 0%   210.40 ±  5%  +229.83% (p=0.002 n=6)
Tier1PfxSize/100_000-16         36.95 ± 0%   209.90 ±  5%  +468.06% (p=0.002 n=6)
Tier1PfxSize/200_000-16         30.07 ± 0%   209.20 ±  5%  +595.71% (p=0.002 n=6)
Tier1PfxSize/500_000-16         23.63 ± 0%   207.90 ±  7%  +779.81% (p=0.002 n=6)
Tier1PfxSize/1_000_000-16       20.42 ± 0%   205.00 ±  7%  +903.92% (p=0.002 n=6)
RandomPfx4Size/1_000-16         61.94 ± 3%   206.20 ±  8%  +232.90% (p=0.002 n=6)
RandomPfx4Size/10_000-16        38.09 ± 1%   186.40 ±  9%  +389.37% (p=0.002 n=6)
RandomPfx4Size/100_000-16       49.80 ± 0%   179.60 ±  9%  +260.64% (p=0.002 n=6)
RandomPfx4Size/200_000-16       43.91 ± 0%   178.50 ± 10%  +306.51% (p=0.002 n=6)
RandomPfx4Size/500_000-16       31.18 ± 0%   176.80 ± 10%  +467.03% (p=0.002 n=6)
RandomPfx4Size/1_000_000-16     25.83 ± 0%   175.50 ± 10%  +579.44% (p=0.002 n=6)
RandomPfx6Size/1_000-16         66.32 ± 3%   228.60 ±  8%  +244.69% (p=0.002 n=6)
RandomPfx6Size/10_000-16        80.59 ± 0%   222.30 ±  8%  +175.84% (p=0.002 n=6)
RandomPfx6Size/100_000-16       53.57 ± 0%   213.90 ±  9%  +299.29% (p=0.002 n=6)
RandomPfx6Size/200_000-16       51.39 ± 0%   210.50 ±  9%  +309.61% (p=0.002 n=6)
RandomPfx6Size/500_000-16       53.92 ± 0%   206.70 ±  9%  +283.35% (p=0.002 n=6)
RandomPfx6Size/1_000_000-16     59.52 ± 0%   204.90 ±  9%  +244.25% (p=0.002 n=6)
RandomPfxSize/1_000-16          80.74 ± 3%   215.40 ±  5%  +166.78% (p=0.002 n=6)
RandomPfxSize/10_000-16         56.31 ± 0%   210.10 ±  5%  +273.11% (p=0.002 n=6)
RandomPfxSize/100_000-16        63.44 ± 0%   203.50 ±  7%  +220.78% (p=0.002 n=6)
RandomPfxSize/200_000-16        54.19 ± 0%   198.70 ±  8%  +266.67% (p=0.002 n=6)
RandomPfxSize/500_000-16        39.17 ± 0%   189.90 ±  9%  +384.81% (p=0.002 n=6)
RandomPfxSize/1_000_000-16      31.88 ± 0%   181.40 ± 10%  +469.01% (p=0.002 n=6)
geomean                         46.76         201.4        +330.75%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                            │ lite/size.bm │       kentik-patricia/size.bm       │
                            │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000-16           90.94 ± 2%   145.50 ± 1%   +60.00% (p=0.002 n=6)
Tier1PfxSize/10_000-16          63.79 ± 0%   200.10 ± 0%  +213.69% (p=0.002 n=6)
Tier1PfxSize/100_000-16         36.95 ± 0%   164.30 ± 0%  +344.65% (p=0.002 n=6)
Tier1PfxSize/200_000-16         30.07 ± 0%   164.00 ± 0%  +445.39% (p=0.002 n=6)
Tier1PfxSize/500_000-16         23.63 ± 0%   144.40 ± 0%  +511.09% (p=0.002 n=6)
Tier1PfxSize/1_000_000-16       20.42 ± 0%   144.40 ± 0%  +607.15% (p=0.002 n=6)
RandomPfx4Size/1_000-16         61.94 ± 3%   141.00 ± 1%  +127.64% (p=0.002 n=6)
RandomPfx4Size/10_000-16        38.09 ± 1%   109.40 ± 0%  +187.21% (p=0.002 n=6)
RandomPfx4Size/100_000-16       49.80 ± 0%   139.80 ± 0%  +180.72% (p=0.002 n=6)
RandomPfx4Size/200_000-16       43.91 ± 0%   139.80 ± 0%  +218.38% (p=0.002 n=6)
RandomPfx4Size/500_000-16       31.18 ± 0%   139.60 ± 0%  +347.72% (p=0.002 n=6)
RandomPfx4Size/1_000_000-16     25.83 ± 0%   139.70 ± 0%  +440.84% (p=0.002 n=6)
RandomPfx6Size/1_000-16         66.32 ± 3%   157.30 ± 1%  +137.18% (p=0.002 n=6)
RandomPfx6Size/10_000-16        80.59 ± 0%   201.20 ± 0%  +149.66% (p=0.002 n=6)
RandomPfx6Size/100_000-16       53.57 ± 0%   160.80 ± 0%  +200.17% (p=0.002 n=6)
RandomPfx6Size/200_000-16       51.39 ± 0%   160.80 ± 0%  +212.90% (p=0.002 n=6)
RandomPfx6Size/500_000-16       53.92 ± 0%   156.40 ± 0%  +190.06% (p=0.002 n=6)
RandomPfx6Size/1_000_000-16     59.52 ± 0%   156.40 ± 0%  +162.77% (p=0.002 n=6)
RandomPfxSize/1_000-16          80.74 ± 3%   145.20 ± 1%   +79.84% (p=0.002 n=6)
RandomPfxSize/10_000-16         56.31 ± 0%   140.00 ± 0%  +148.62% (p=0.002 n=6)
RandomPfxSize/100_000-16        63.44 ± 0%   180.00 ± 0%  +183.73% (p=0.002 n=6)
RandomPfxSize/200_000-16        54.19 ± 0%   180.00 ± 0%  +232.16% (p=0.002 n=6)
RandomPfxSize/500_000-16        39.17 ± 0%   144.00 ± 0%  +267.63% (p=0.002 n=6)
RandomPfxSize/1_000_000-16      31.88 ± 0%   144.00 ± 0%  +351.69% (p=0.002 n=6)
geomean                         46.76         152.8       +226.84%
```

## update, insert/delete

`bart.Lite` is the fastest algorithm for updates.

```
goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                            │ lite/update.bm │            bart/update.bm             │
                            │   sec/route    │   sec/route    vs base                │
InsertRandomPfxs/1_000-16      119.9n ±   4%    158.5n ±  6%   +32.19% (p=0.002 n=6)
InsertRandomPfxs/10_000-16     91.81n ±   5%   120.45n ±  5%   +31.19% (p=0.002 n=6)
InsertRandomPfxs/100_000-16    126.2n ±   5%    178.1n ±  2%   +41.07% (p=0.002 n=6)
InsertRandomPfxs/200_000-16    146.1n ±   2%    212.9n ±  4%   +45.77% (p=0.002 n=6)
DeleteRandomPfxs/1_000-16      85.15n ±   4%    93.30n ±  3%    +9.57% (p=0.002 n=6)
DeleteRandomPfxs/10_000-16     59.30n ±   2%    67.50n ±  1%   +13.83% (p=0.002 n=6)
DeleteRandomPfxs/100_000-16    129.1n ±   2%    164.1n ±  1%   +27.20% (p=0.002 n=6)
DeleteRandomPfxs/200_000-16    220.9n ± 424%   1338.0n ± 83%  +505.57% (p=0.041 n=6)
geomean                        114.3n           177.9n         +55.54%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                            │ lite/update.bm │             fast/update.bm             │
                            │   sec/route    │   sec/route     vs base                │
InsertRandomPfxs/1_000-16      119.9n ±   4%    545.9n ±   7%  +355.30% (p=0.002 n=6)
InsertRandomPfxs/10_000-16     91.81n ±   5%   249.00n ±   5%  +171.21% (p=0.002 n=6)
InsertRandomPfxs/100_000-16    126.2n ±   5%    601.9n ±   7%  +376.75% (p=0.002 n=6)
InsertRandomPfxs/200_000-16    146.1n ±   2%    672.4n ±   2%  +360.36% (p=0.002 n=6)
DeleteRandomPfxs/1_000-16      85.15n ±   4%   186.25n ±   3%  +118.73% (p=0.002 n=6)
DeleteRandomPfxs/10_000-16     59.30n ±   2%   153.20n ± 339%  +158.33% (p=0.002 n=6)
DeleteRandomPfxs/100_000-16    129.1n ±   2%    277.5n ± 376%  +115.03% (p=0.002 n=6)
DeleteRandomPfxs/200_000-16    220.9n ± 424%    340.4n ±   5%         ~ (p=0.394 n=6)
geomean                        114.3n           332.2n         +190.51%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                            │ lite/update.bm │          netipds/update.bm           │
                            │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000-16      119.9n ±   4%    201.9n ± 3%   +68.39% (p=0.002 n=6)
InsertRandomPfxs/10_000-16     91.81n ±   5%   236.60n ± 5%  +157.71% (p=0.002 n=6)
InsertRandomPfxs/100_000-16    126.2n ±   5%    302.9n ± 3%  +139.92% (p=0.002 n=6)
InsertRandomPfxs/200_000-16    146.1n ±   2%    406.4n ± 2%  +178.23% (p=0.002 n=6)
DeleteRandomPfxs/1_000-16      85.15n ±   4%   122.65n ± 3%   +44.04% (p=0.002 n=6)
DeleteRandomPfxs/10_000-16     59.30n ±   2%   158.80n ± 4%  +167.77% (p=0.002 n=6)
DeleteRandomPfxs/100_000-16    129.1n ±   2%    276.9n ± 2%  +114.57% (p=0.002 n=6)
DeleteRandomPfxs/200_000-16    220.9n ± 424%    391.0n ± 2%         ~ (p=0.372 n=6)
geomean                        114.3n           243.6n       +113.03%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                            │ lite/update.bm │         critbitgo/update.bm          │
                            │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000-16      119.9n ±   4%    253.5n ± 2%  +111.38% (p=0.002 n=6)
InsertRandomPfxs/10_000-16     91.81n ±   5%   313.50n ± 1%  +241.47% (p=0.002 n=6)
InsertRandomPfxs/100_000-16    126.2n ±   5%    491.6n ± 5%  +289.39% (p=0.002 n=6)
InsertRandomPfxs/200_000-16    146.1n ±   2%    660.2n ± 3%  +352.04% (p=0.002 n=6)
DeleteRandomPfxs/1_000-16      85.15n ±   4%   129.25n ± 2%   +51.79% (p=0.002 n=6)
DeleteRandomPfxs/10_000-16     59.30n ±   2%   167.95n ± 1%  +183.20% (p=0.002 n=6)
DeleteRandomPfxs/100_000-16    129.1n ±   2%    351.4n ± 2%  +172.30% (p=0.002 n=6)
DeleteRandomPfxs/200_000-16    220.9n ± 424%    450.1n ± 1%         ~ (p=0.394 n=6)
geomean                        114.3n           311.4n       +172.37%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                            │ lite/update.bm │          lpmtrie/update.bm           │
                            │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000-16      119.9n ±   4%    339.8n ± 3%  +183.40% (p=0.002 n=6)
InsertRandomPfxs/10_000-16     91.81n ±   5%   370.10n ± 3%  +303.12% (p=0.002 n=6)
InsertRandomPfxs/100_000-16    126.2n ±   5%    602.5n ± 2%  +377.19% (p=0.002 n=6)
InsertRandomPfxs/200_000-16    146.1n ±   2%    747.3n ± 3%  +411.64% (p=0.002 n=6)
DeleteRandomPfxs/1_000-16      85.15n ±   4%   132.00n ± 3%   +55.02% (p=0.002 n=6)
DeleteRandomPfxs/10_000-16     59.30n ±   2%   185.20n ± 1%  +212.28% (p=0.002 n=6)
DeleteRandomPfxs/100_000-16    129.1n ±   2%    478.3n ± 1%  +270.63% (p=0.002 n=6)
DeleteRandomPfxs/200_000-16    220.9n ± 424%    642.1n ± 2%         ~ (p=0.240 n=6)
geomean                        114.3n           378.9n       +231.39%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                            │ lite/update.bm │       kentik-patricia/update.bm        │
                            │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000-16      119.9n ±   4%    260.8n ±  3%   +117.56% (p=0.002 n=6)
InsertRandomPfxs/10_000-16     91.81n ±   5%   333.05n ±  7%   +262.76% (p=0.002 n=6)
InsertRandomPfxs/100_000-16    126.2n ±   5%    527.5n ±  4%   +317.78% (p=0.002 n=6)
InsertRandomPfxs/200_000-16    146.1n ±   2%    650.3n ±  3%   +345.26% (p=0.002 n=6)
DeleteRandomPfxs/1_000-16      85.15n ±   4%   282.85n ±  4%   +232.18% (p=0.002 n=6)
DeleteRandomPfxs/10_000-16     59.30n ±   2%   316.75n ±  3%   +434.10% (p=0.002 n=6)
DeleteRandomPfxs/100_000-16    129.1n ±   2%   3302.5n ± 81%  +2459.09% (p=0.002 n=6)
DeleteRandomPfxs/200_000-16    220.9n ± 424%    769.9n ±  3%          ~ (p=0.058 n=6)
geomean                        114.3n           535.8n         +368.54%
```

