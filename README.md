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
LpmTier1Pfxs/RandomMatchIP4-8            8.100n ±  0%   15.000n ±  3%   +85.19% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            12.29n ±  7%    24.87n ± 20%  +102.24% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             8.401n ±  0%   16.730n ±  0%   +99.14% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             14.19n ± 26%    26.04n ± 30%   +83.54% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     13.77n ±  0%    16.89n ±  0%   +22.69% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     13.92n ±  2%    19.87n ±  4%   +42.74% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      13.75n ± 39%    16.88n ± 20%   +22.76% (p=0.003 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      13.85n ±  1%    19.89n ±  8%   +43.61% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    13.76n ±  0%    18.29n ±  0%   +32.92% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    13.86n ±  2%    19.88n ±  1%   +43.38% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     13.76n ±  0%    18.29n ±  0%   +32.92% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     13.86n ±  3%    19.87n ±  0%   +43.31% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   7.275n ± 13%   11.826n ± 37%   +62.57% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   15.86n ± 45%    25.48n ± 44%   +60.64% (p=0.002 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    13.76n ± 39%    18.28n ±  8%   +32.85% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    15.88n ±  0%    25.49n ±  1%   +60.57% (p=0.000 n=20)
geomean                                  12.59n          19.20n         +52.49%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ fast/lpm.bm  │              bart/lpm.bm               │
                                       │    sec/op    │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4-8            8.100n ±  0%   15.285n ±  6%   +88.70% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            12.29n ±  7%    25.77n ± 12%  +109.56% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             8.401n ±  0%   16.410n ±  0%   +95.33% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             14.19n ± 26%    26.01n ± 29%   +83.36% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     13.77n ±  0%    17.10n ±  3%   +24.15% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     13.92n ±  2%    19.76n ±  0%   +41.95% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      13.75n ± 39%    17.06n ± 19%   +24.07% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      13.85n ±  1%    19.76n ±  0%   +42.67% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    13.76n ±  0%    18.50n ±  1%   +34.48% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    13.86n ±  2%    19.77n ± 13%   +42.59% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     13.76n ±  0%    18.51n ±  0%   +34.52% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     13.86n ±  3%    20.35n ± 10%   +46.77% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   7.275n ± 13%   12.088n ± 39%   +66.16% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   15.86n ± 45%    26.11n ± 43%   +64.54% (p=0.001 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    13.76n ± 39%    18.50n ±  8%   +34.45% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    15.88n ±  0%    26.17n ±  3%   +64.85% (p=0.000 n=20)
geomean                                  12.59n          19.40n         +54.12%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ fast/lpm.bm  │             netipds/lpm.bm             │
                                       │    sec/op    │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4-8            8.100n ±  0%   46.750n ± 16%  +477.16% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            12.29n ±  7%    61.39n ± 21%  +399.31% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             8.401n ±  0%   51.690n ±  6%  +515.28% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             14.19n ± 26%    64.21n ± 19%  +352.63% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     13.77n ±  0%    28.50n ±  5%  +106.97% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     13.92n ±  2%    31.98n ±  6%  +129.74% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      13.75n ± 39%    29.21n ±  5%  +112.44% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      13.85n ±  1%    32.09n ± 14%  +131.73% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    13.76n ±  0%    34.37n ±  5%  +149.75% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    13.86n ±  2%    36.13n ±  7%  +160.58% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     13.76n ±  0%    36.62n ±  4%  +166.17% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     13.86n ±  3%    39.17n ± 22%  +182.47% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   7.275n ± 13%   36.535n ± 12%  +402.23% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   15.86n ± 45%    40.67n ± 12%  +156.38% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    13.76n ± 39%    41.73n ±  2%  +203.27% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    15.88n ±  0%    49.03n ±  7%  +208.85% (p=0.000 n=20)
geomean                                  12.59n          40.06n        +218.27%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ fast/lpm.bm  │             critbitgo/lpm.bm             │
                                       │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4-8            8.100n ±  0%   139.450n ±  4%  +1621.60% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            12.29n ±  7%    174.65n ±  5%  +1320.50% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             8.401n ±  0%   627.650n ± 19%  +7371.13% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             14.19n ± 26%    472.90n ± 14%  +3233.80% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     13.77n ±  0%     81.97n ±  6%   +495.28% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     13.92n ±  2%     93.95n ±  4%   +574.96% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      13.75n ± 39%    137.15n ± 13%   +897.45% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      13.85n ±  1%    143.80n ± 12%   +938.27% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    13.76n ±  0%     93.91n ±  3%   +582.52% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    13.86n ±  2%    104.20n ±  5%   +651.53% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     13.76n ±  0%    252.10n ± 17%  +1732.12% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     13.86n ±  3%    201.60n ± 16%  +1354.02% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   7.275n ± 13%   106.050n ±  5%  +1357.83% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   15.86n ± 45%    116.70n ±  2%   +635.58% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    13.76n ± 39%    452.60n ± 11%  +3189.24% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    15.88n ±  0%    323.40n ± 23%  +1937.17% (p=0.000 n=20)
geomean                                  12.59n           177.7n        +1311.72%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ fast/lpm.bm  │              lpmtrie/lpm.bm              │
                                       │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4-8            8.100n ±  0%   276.650n ± 11%  +3315.43% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            12.29n ±  7%    317.10n ± 17%  +2479.10% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             8.401n ±  0%   254.000n ±  3%  +2923.45% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             14.19n ± 26%    226.60n ± 11%  +1497.46% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     13.77n ±  0%    130.65n ±  7%   +848.80% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     13.92n ±  2%    120.30n ± 24%   +764.22% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      13.75n ± 39%    113.35n ± 19%   +724.36% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      13.85n ±  1%     94.23n ± 37%   +580.36% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    13.76n ±  0%    151.10n ± 17%   +998.11% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    13.86n ±  2%    140.00n ±  5%   +909.74% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     13.76n ±  0%    159.85n ± 10%  +1061.70% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     13.86n ±  3%    134.85n ±  6%   +872.59% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   7.275n ± 13%   194.600n ±  6%  +2575.10% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   15.86n ± 45%    185.55n ±  3%  +1069.56% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    13.76n ± 39%    196.40n ±  5%  +1327.33% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    15.88n ±  0%    188.40n ±  5%  +1086.77% (p=0.000 n=20)
geomean                                  12.59n           170.6n        +1255.51%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ fast/lpm.bm  │          kentik-patricia/lpm.bm          │
                                       │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4-8            8.100n ±  0%   174.450n ±  8%  +2053.70% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            12.29n ±  7%    259.65n ± 17%  +2011.83% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             8.401n ±  0%   138.050n ±  5%  +1543.26% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             14.19n ± 26%    168.15n ± 13%  +1085.41% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     13.77n ±  0%     79.87n ±  8%   +480.03% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     13.92n ±  2%     99.86n ± 16%   +617.35% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      13.75n ± 39%     65.84n ±  5%   +378.87% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      13.85n ±  1%     79.03n ± 10%   +470.58% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    13.76n ±  0%    109.05n ±  6%   +692.51% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    13.86n ±  2%    125.25n ±  4%   +803.35% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     13.76n ±  0%     91.31n ±  4%   +563.59% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     13.86n ±  3%    108.45n ±  4%   +682.19% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   7.275n ± 13%   130.300n ±  9%  +1691.19% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   15.86n ± 45%    153.45n ±  5%   +867.22% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    13.76n ± 39%    106.75n ±  6%   +675.80% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    15.88n ±  0%    137.55n ±  6%   +766.46% (p=0.000 n=20)
geomean                                  12.59n           119.5n         +849.43%
```

## size of the routing tables


`bart.Lite` has the lowest memory consumption under all competitors.


```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/size.bm │            bart/size.bm            │
                           │ bytes/route  │ bytes/route  vs base               │
Tier1PfxSize/1_000-8           85.21 ± 2%   105.10 ± 2%  +23.34% (p=0.002 n=6)
Tier1PfxSize/10_000-8          63.83 ± 0%    83.90 ± 0%  +31.44% (p=0.002 n=6)
Tier1PfxSize/100_000-8         36.96 ± 0%    56.71 ± 0%  +53.44% (p=0.002 n=6)
Tier1PfxSize/200_000-8         30.07 ± 0%    49.53 ± 0%  +64.72% (p=0.002 n=6)
Tier1PfxSize/500_000-8         23.63 ± 0%    42.90 ± 0%  +81.55% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       20.42 ± 0%    39.74 ± 0%  +94.61% (p=0.002 n=6)
RandomPfx4Size/1_000-8         62.35 ± 3%    82.98 ± 2%  +33.09% (p=0.002 n=6)
RandomPfx4Size/10_000-8        38.25 ± 1%    57.26 ± 0%  +49.70% (p=0.002 n=6)
RandomPfx4Size/100_000-8       49.80 ± 0%    72.08 ± 0%  +44.74% (p=0.002 n=6)
RandomPfx4Size/200_000-8       43.92 ± 0%    65.15 ± 0%  +48.34% (p=0.002 n=6)
RandomPfx4Size/500_000-8       31.18 ± 0%    52.44 ± 0%  +68.18% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     25.83 ± 0%    46.46 ± 0%  +79.87% (p=0.002 n=6)
RandomPfx6Size/1_000-8         66.18 ± 3%    84.19 ± 2%  +27.21% (p=0.002 n=6)
RandomPfx6Size/10_000-8        80.69 ± 0%   101.30 ± 0%  +25.54% (p=0.002 n=6)
RandomPfx6Size/100_000-8       53.58 ± 0%    71.68 ± 0%  +33.78% (p=0.002 n=6)
RandomPfx6Size/200_000-8       51.39 ± 0%    69.37 ± 0%  +34.99% (p=0.002 n=6)
RandomPfx6Size/500_000-8       53.93 ± 0%    72.34 ± 0%  +34.14% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     59.52 ± 0%    78.92 ± 0%  +32.59% (p=0.002 n=6)
RandomPfxSize/1_000-8          80.61 ± 2%   100.60 ± 2%  +24.80% (p=0.002 n=6)
RandomPfxSize/10_000-8         56.42 ± 0%    74.00 ± 0%  +31.16% (p=0.002 n=6)
RandomPfxSize/100_000-8        63.45 ± 0%    84.33 ± 0%  +32.91% (p=0.002 n=6)
RandomPfxSize/200_000-8        54.19 ± 0%    75.59 ± 0%  +39.49% (p=0.002 n=6)
RandomPfxSize/500_000-8        39.18 ± 0%    59.59 ± 0%  +52.09% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      31.89 ± 0%    52.03 ± 0%  +63.15% (p=0.002 n=6)
geomean                        46.65         67.56       +44.81%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/size.bm │             fast/size.bm              │
                           │ bytes/route  │ bytes/route   vs base                 │
Tier1PfxSize/1_000-8           85.21 ± 2%   1267.00 ± 0%  +1386.91% (p=0.002 n=6)
Tier1PfxSize/10_000-8          63.83 ± 0%   1252.00 ± 0%  +1861.46% (p=0.002 n=6)
Tier1PfxSize/100_000-8         36.96 ± 0%    839.50 ± 0%  +2171.37% (p=0.002 n=6)
Tier1PfxSize/200_000-8         30.07 ± 0%    600.40 ± 0%  +1896.67% (p=0.002 n=6)
Tier1PfxSize/500_000-8         23.63 ± 0%    374.50 ± 0%  +1484.85% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       20.42 ± 0%    259.00 ± 0%  +1168.36% (p=0.002 n=6)
RandomPfx4Size/1_000-8         62.35 ± 3%   1205.00 ± 0%  +1832.64% (p=0.002 n=6)
RandomPfx4Size/10_000-8        38.25 ± 1%    326.30 ± 0%   +753.07% (p=0.002 n=6)
RandomPfx4Size/100_000-8       49.80 ± 0%   1122.00 ± 0%  +2153.01% (p=0.002 n=6)
RandomPfx4Size/200_000-8       43.92 ± 0%   1200.00 ± 0%  +2632.24% (p=0.002 n=6)
RandomPfx4Size/500_000-8       31.18 ± 0%    677.30 ± 0%  +2072.23% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     25.83 ± 0%    372.80 ± 0%  +1343.28% (p=0.002 n=6)
RandomPfx6Size/1_000-8         66.18 ± 3%    515.20 ± 0%   +678.48% (p=0.002 n=6)
RandomPfx6Size/10_000-8        80.69 ± 0%   1461.00 ± 0%  +1710.63% (p=0.002 n=6)
RandomPfx6Size/100_000-8       53.58 ± 0%    531.40 ± 0%   +891.79% (p=0.002 n=6)
RandomPfx6Size/200_000-8       51.39 ± 0%    403.00 ± 0%   +684.20% (p=0.002 n=6)
RandomPfx6Size/500_000-8       53.93 ± 0%    483.70 ± 0%   +796.90% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     59.52 ± 0%    730.50 ± 0%  +1127.32% (p=0.002 n=6)
RandomPfxSize/1_000-8          80.61 ± 2%   1275.00 ± 0%  +1481.69% (p=0.002 n=6)
RandomPfxSize/10_000-8         56.42 ± 0%    481.10 ± 0%   +752.71% (p=0.002 n=6)
RandomPfxSize/100_000-8        63.45 ± 0%   1340.00 ± 0%  +2011.90% (p=0.002 n=6)
RandomPfxSize/200_000-8        54.19 ± 0%   1306.00 ± 0%  +2310.04% (p=0.002 n=6)
RandomPfxSize/500_000-8        39.18 ± 0%    765.80 ± 0%  +1854.57% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      31.89 ± 0%    439.20 ± 0%  +1277.23% (p=0.002 n=6)
geomean                        46.65          703.1       +1407.14%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/size.bm │           netipds/size.bm           │
                           │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000-8           85.21 ± 2%    74.06 ± 3%   -13.09% (p=0.002 n=6)
Tier1PfxSize/10_000-8          63.83 ± 0%    68.83 ± 0%    +7.83% (p=0.002 n=6)
Tier1PfxSize/100_000-8         36.96 ± 0%    67.64 ± 0%   +83.01% (p=0.002 n=6)
Tier1PfxSize/200_000-8         30.07 ± 0%    66.90 ± 0%  +122.48% (p=0.002 n=6)
Tier1PfxSize/500_000-8         23.63 ± 0%    65.35 ± 0%  +176.56% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       20.42 ± 0%    63.72 ± 0%  +212.05% (p=0.002 n=6)
RandomPfx4Size/1_000-8         62.35 ± 3%    61.79 ± 3%    -0.90% (p=0.013 n=6)
RandomPfx4Size/10_000-8        38.25 ± 1%    51.30 ± 0%   +34.12% (p=0.002 n=6)
RandomPfx4Size/100_000-8       49.80 ± 0%    48.24 ± 0%    -3.13% (p=0.002 n=6)
RandomPfx4Size/200_000-8       43.92 ± 0%    47.57 ± 0%    +8.31% (p=0.002 n=6)
RandomPfx4Size/500_000-8       31.18 ± 0%    46.51 ± 0%   +49.17% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     25.83 ± 0%    45.65 ± 0%   +76.73% (p=0.002 n=6)
RandomPfx6Size/1_000-8         66.18 ± 3%   100.10 ± 2%   +51.25% (p=0.002 n=6)
RandomPfx6Size/10_000-8        80.69 ± 0%    94.17 ± 0%   +16.71% (p=0.002 n=6)
RandomPfx6Size/100_000-8       53.58 ± 0%    87.64 ± 0%   +63.57% (p=0.002 n=6)
RandomPfx6Size/200_000-8       51.39 ± 0%    85.90 ± 0%   +67.15% (p=0.002 n=6)
RandomPfx6Size/500_000-8       53.93 ± 0%    84.11 ± 0%   +55.96% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     59.52 ± 0%    83.23 ± 0%   +39.84% (p=0.002 n=6)
RandomPfxSize/1_000-8          80.61 ± 2%    74.80 ± 3%    -7.21% (p=0.002 n=6)
RandomPfxSize/10_000-8         56.42 ± 0%    69.29 ± 0%   +22.81% (p=0.002 n=6)
RandomPfxSize/100_000-8        63.45 ± 0%    64.28 ± 0%    +1.31% (p=0.002 n=6)
RandomPfxSize/200_000-8        54.19 ± 0%    61.55 ± 0%   +13.58% (p=0.002 n=6)
RandomPfxSize/500_000-8        39.18 ± 0%    57.42 ± 0%   +46.55% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      31.89 ± 0%    53.93 ± 0%   +69.11% (p=0.002 n=6)
geomean                        46.65         66.01        +41.49%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/size.bm │          critbitgo/size.bm          │
                           │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000-8           85.21 ± 2%   119.50 ± 2%   +40.24% (p=0.002 n=6)
Tier1PfxSize/10_000-8          63.83 ± 0%   114.70 ± 0%   +79.70% (p=0.002 n=6)
Tier1PfxSize/100_000-8         36.96 ± 0%   114.40 ± 0%  +209.52% (p=0.002 n=6)
Tier1PfxSize/200_000-8         30.07 ± 0%   114.40 ± 0%  +280.45% (p=0.002 n=6)
Tier1PfxSize/500_000-8         23.63 ± 0%   114.40 ± 0%  +384.13% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       20.42 ± 0%   114.40 ± 0%  +460.24% (p=0.002 n=6)
RandomPfx4Size/1_000-8         62.35 ± 3%   116.40 ± 2%   +86.69% (p=0.002 n=6)
RandomPfx4Size/10_000-8        38.25 ± 1%   112.30 ± 0%  +193.59% (p=0.002 n=6)
RandomPfx4Size/100_000-8       49.80 ± 0%   112.00 ± 0%  +124.90% (p=0.002 n=6)
RandomPfx4Size/200_000-8       43.92 ± 0%   112.00 ± 0%  +155.01% (p=0.002 n=6)
RandomPfx4Size/500_000-8       31.18 ± 0%   112.00 ± 0%  +259.20% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     25.83 ± 0%   112.00 ± 0%  +333.60% (p=0.002 n=6)
RandomPfx6Size/1_000-8         66.18 ± 3%   132.90 ± 2%  +100.82% (p=0.002 n=6)
RandomPfx6Size/10_000-8        80.69 ± 0%   128.30 ± 0%   +59.00% (p=0.002 n=6)
RandomPfx6Size/100_000-8       53.58 ± 0%   128.00 ± 0%  +138.90% (p=0.002 n=6)
RandomPfx6Size/200_000-8       51.39 ± 0%   128.00 ± 0%  +149.08% (p=0.002 n=6)
RandomPfx6Size/500_000-8       53.93 ± 0%   128.00 ± 0%  +137.34% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     59.52 ± 0%   128.00 ± 0%  +115.05% (p=0.002 n=6)
RandomPfxSize/1_000-8          80.61 ± 2%   119.70 ± 2%   +48.49% (p=0.002 n=6)
RandomPfxSize/10_000-8         56.42 ± 0%   115.50 ± 0%  +104.71% (p=0.002 n=6)
RandomPfxSize/100_000-8        63.45 ± 0%   115.30 ± 0%   +81.72% (p=0.002 n=6)
RandomPfxSize/200_000-8        54.19 ± 0%   115.20 ± 0%  +112.59% (p=0.002 n=6)
RandomPfxSize/500_000-8        39.18 ± 0%   115.20 ± 0%  +194.03% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      31.89 ± 0%   115.20 ± 0%  +261.24% (p=0.002 n=6)
geomean                        46.65         118.1       +153.08%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/size.bm │           lpmtrie/size.bm            │
                           │ bytes/route  │ bytes/route   vs base                │
Tier1PfxSize/1_000-8           85.21 ± 2%   215.50 ±  5%  +152.90% (p=0.002 n=6)
Tier1PfxSize/10_000-8          63.83 ± 0%   210.50 ±  5%  +229.78% (p=0.002 n=6)
Tier1PfxSize/100_000-8         36.96 ± 0%   209.90 ±  5%  +467.91% (p=0.002 n=6)
Tier1PfxSize/200_000-8         30.07 ± 0%   209.20 ±  5%  +595.71% (p=0.002 n=6)
Tier1PfxSize/500_000-8         23.63 ± 0%   207.90 ±  7%  +779.81% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       20.42 ± 0%   205.00 ±  7%  +903.92% (p=0.002 n=6)
RandomPfx4Size/1_000-8         62.35 ± 3%   205.45 ±  8%  +229.51% (p=0.002 n=6)
RandomPfx4Size/10_000-8        38.25 ± 1%   186.50 ±  9%  +387.58% (p=0.002 n=6)
RandomPfx4Size/100_000-8       49.80 ± 0%   179.60 ±  9%  +260.64% (p=0.002 n=6)
RandomPfx4Size/200_000-8       43.92 ± 0%   178.50 ± 10%  +306.42% (p=0.002 n=6)
RandomPfx4Size/500_000-8       31.18 ± 0%   176.80 ± 10%  +467.03% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     25.83 ± 0%   175.50 ± 10%  +579.44% (p=0.002 n=6)
RandomPfx6Size/1_000-8         66.18 ± 3%   227.90 ±  8%  +244.36% (p=0.002 n=6)
RandomPfx6Size/10_000-8        80.69 ± 0%   222.40 ±  8%  +175.62% (p=0.002 n=6)
RandomPfx6Size/100_000-8       53.58 ± 0%   213.90 ±  9%  +299.22% (p=0.002 n=6)
RandomPfx6Size/200_000-8       51.39 ± 0%   210.50 ±  9%  +309.61% (p=0.002 n=6)
RandomPfx6Size/500_000-8       53.93 ± 0%   206.70 ±  9%  +283.27% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     59.52 ± 0%   204.90 ±  9%  +244.25% (p=0.002 n=6)
RandomPfxSize/1_000-8          80.61 ± 2%   215.20 ±  5%  +166.96% (p=0.002 n=6)
RandomPfxSize/10_000-8         56.42 ± 0%   210.20 ±  5%  +272.56% (p=0.002 n=6)
RandomPfxSize/100_000-8        63.45 ± 0%   203.50 ±  7%  +220.72% (p=0.002 n=6)
RandomPfxSize/200_000-8        54.19 ± 0%   198.70 ±  8%  +266.67% (p=0.002 n=6)
RandomPfxSize/500_000-8        39.18 ± 0%   189.90 ±  9%  +384.69% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      31.89 ± 0%   181.40 ± 10%  +468.83% (p=0.002 n=6)
geomean                        46.65         201.3        +331.58%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/size.bm │       kentik-patricia/size.bm       │
                           │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000-8           85.21 ± 2%   145.40 ± 1%   +70.64% (p=0.002 n=6)
Tier1PfxSize/10_000-8          63.83 ± 0%   200.20 ± 0%  +213.65% (p=0.002 n=6)
Tier1PfxSize/100_000-8         36.96 ± 0%   163.50 ± 0%  +342.37% (p=0.002 n=6)
Tier1PfxSize/200_000-8         30.07 ± 0%   164.20 ± 0%  +446.06% (p=0.002 n=6)
Tier1PfxSize/500_000-8         23.63 ± 0%   144.40 ± 0%  +511.09% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       20.42 ± 0%   144.30 ± 0%  +606.66% (p=0.002 n=6)
RandomPfx4Size/1_000-8         62.35 ± 3%   140.80 ± 1%  +125.82% (p=0.002 n=6)
RandomPfx4Size/10_000-8        38.25 ± 1%   109.60 ± 0%  +186.54% (p=0.002 n=6)
RandomPfx4Size/100_000-8       49.80 ± 0%   139.80 ± 0%  +180.72% (p=0.002 n=6)
RandomPfx4Size/200_000-8       43.92 ± 0%   139.80 ± 0%  +218.31% (p=0.002 n=6)
RandomPfx4Size/500_000-8       31.18 ± 0%   139.60 ± 0%  +347.72% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     25.83 ± 0%   139.60 ± 0%  +440.46% (p=0.002 n=6)
RandomPfx6Size/1_000-8         66.18 ± 3%   157.20 ± 1%  +137.53% (p=0.002 n=6)
RandomPfx6Size/10_000-8        80.69 ± 0%   201.30 ± 0%  +149.47% (p=0.002 n=6)
RandomPfx6Size/100_000-8       53.58 ± 0%   160.80 ± 0%  +200.11% (p=0.002 n=6)
RandomPfx6Size/200_000-8       51.39 ± 0%   160.80 ± 0%  +212.90% (p=0.002 n=6)
RandomPfx6Size/500_000-8       53.93 ± 0%   156.50 ± 0%  +190.19% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     59.52 ± 0%   156.40 ± 0%  +162.77% (p=0.002 n=6)
RandomPfxSize/1_000-8          80.61 ± 2%   144.60 ± 1%   +79.38% (p=0.002 n=6)
RandomPfxSize/10_000-8         56.42 ± 0%   140.10 ± 0%  +148.32% (p=0.002 n=6)
RandomPfxSize/100_000-8        63.45 ± 0%   180.00 ± 0%  +183.69% (p=0.002 n=6)
RandomPfxSize/200_000-8        54.19 ± 0%   180.00 ± 0%  +232.16% (p=0.002 n=6)
RandomPfxSize/500_000-8        39.18 ± 0%   144.00 ± 0%  +267.53% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      31.89 ± 0%   144.00 ± 0%  +351.55% (p=0.002 n=6)
geomean                        46.65         152.8       +227.45%
```

## update, insert/delete

`bart.Lite` is the fastest algorithm for updates.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/update.bm │           bart/update.bm           │
                           │   sec/route    │  sec/route   vs base               │
InsertRandomPfxs/1_000-8        140.2n ± 2%   165.0n ± 2%  +17.69% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       104.8n ± 2%   128.1n ± 2%  +22.29% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      146.8n ± 1%   188.2n ± 2%  +28.20% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      165.2n ± 4%   225.8n ± 4%  +36.69% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        107.2n ± 5%   114.8n ± 8%   +7.09% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       72.54n ± 1%   83.42n ± 1%  +15.00% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      147.8n ± 2%   176.5n ± 3%  +19.42% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      181.6n ± 2%   237.5n ± 4%  +30.77% (p=0.002 n=6)
geomean                         128.5n        156.5n       +21.82%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/update.bm │            fast/update.bm            │
                           │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000-8        140.2n ± 2%    363.0n ± 2%  +158.88% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       104.8n ± 2%    198.9n ± 3%   +89.88% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      146.8n ± 1%    461.1n ± 8%  +214.13% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      165.2n ± 4%    589.1n ± 4%  +256.68% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        107.2n ± 5%    140.5n ± 2%   +31.02% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       72.54n ± 1%   102.15n ± 2%   +40.82% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      147.8n ± 2%    233.4n ± 4%   +57.92% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      181.6n ± 2%    300.9n ± 5%   +65.65% (p=0.002 n=6)
geomean                         128.5n         258.2n       +100.94%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/update.bm │          netipds/update.bm           │
                           │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000-8        140.2n ± 2%    224.9n ± 7%   +60.45% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       104.8n ± 2%    285.2n ± 2%  +172.32% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      146.8n ± 1%    382.8n ± 1%  +160.80% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      165.2n ± 4%    461.1n ± 5%  +179.17% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        107.2n ± 5%    160.4n ± 8%   +49.67% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       72.54n ± 1%   203.95n ± 1%  +181.16% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      147.8n ± 2%    308.5n ± 1%  +108.69% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      181.6n ± 2%    407.3n ± 1%  +124.22% (p=0.002 n=6)
geomean                         128.5n         287.4n       +123.67%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/update.bm │         critbitgo/update.bm          │
                           │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000-8        140.2n ± 2%    292.1n ± 2%  +108.31% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       104.8n ± 2%    375.2n ± 1%  +258.19% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      146.8n ± 1%    547.1n ± 6%  +272.68% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      165.2n ± 4%    710.9n ± 3%  +330.43% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        107.2n ± 5%    153.3n ± 1%   +42.96% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       72.54n ± 1%   204.45n ± 1%  +181.84% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      147.8n ± 2%    374.7n ± 1%  +153.55% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      181.6n ± 2%    482.7n ± 2%  +165.73% (p=0.002 n=6)
geomean                         128.5n         353.1n       +174.77%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/update.bm │          lpmtrie/update.bm           │
                           │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000-8        140.2n ± 2%    437.6n ± 2%  +212.09% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       104.8n ± 2%    514.5n ± 4%  +391.17% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      146.8n ± 1%    737.4n ± 3%  +402.32% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      165.2n ± 4%    843.8n ± 2%  +410.93% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        107.2n ± 5%    152.6n ± 1%   +42.30% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       72.54n ± 1%   233.40n ± 1%  +221.75% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      147.8n ± 2%    491.1n ± 1%  +232.27% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      181.6n ± 2%    647.9n ± 2%  +256.67% (p=0.002 n=6)
geomean                         128.5n         446.8n       +247.67%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ lite/update.bm │      kentik-patricia/update.bm       │
                           │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000-8        140.2n ± 2%    238.1n ± 0%   +69.83% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       104.8n ± 2%    334.8n ± 1%  +219.57% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      146.8n ± 1%    525.0n ± 1%  +257.60% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      165.2n ± 4%    656.6n ± 4%  +297.58% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        107.2n ± 5%    306.8n ± 1%  +186.15% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       72.54n ± 1%   370.00n ± 0%  +410.06% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      147.8n ± 2%    675.8n ± 2%  +357.24% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      181.6n ± 2%    832.1n ± 1%  +358.08% (p=0.002 n=6)
geomean                         128.5n         452.4n       +252.04%
```
