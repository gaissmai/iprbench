# iprbench
comparing benchmarks for some golang IP routing table implementations:

```
	github.com/gaissmai/bart.Table
	github.com/gaissmai/bart.Lite
	github.com/gaissmai/bart.Fast
	github.com/aromatt/netipds
	github.com/k-sone/critbitgo
	github.com/Asphaltt/lpmtrie
	github.com/kentik/patricia
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
                                     │ bart/lpm.bm  │              lite/lpm.bm              │
                                     │    sec/op    │    sec/op      vs base                │
LpmTier1Pfxs/RandomMatchIP4            12.91n ±  4%    12.09n ±  4%   -6.31% (p=0.042 n=20)
LpmTier1Pfxs/RandomMatchIP6            19.76n ± 12%    18.82n ± 17%        ~ (p=0.324 n=20)
LpmTier1Pfxs/RandomMissIP4             13.22n ±  3%    12.77n ±  0%   -3.33% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             18.58n ± 35%    18.16n ± 37%        ~ (p=0.425 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     12.96n ±  1%    13.13n ±  4%        ~ (p=0.805 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     14.55n ±  1%    15.84n ±  1%   +8.79% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      13.04n ± 26%    12.36n ± 18%        ~ (p=0.327 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      14.39n ±  3%    15.89n ±  1%  +10.46% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    14.75n ±  2%    14.83n ±  0%        ~ (p=0.664 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    14.41n ±  3%    15.25n ±  1%   +5.83% (p=0.007 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     14.62n ±  2%    14.34n ±  3%        ~ (p=0.116 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     14.39n ±  2%    15.60n ±  2%   +8.48% (p=0.003 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   9.608n ± 36%   10.092n ± 32%        ~ (p=0.129 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   18.94n ± 39%    18.18n ± 37%        ~ (p=0.616 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    14.45n ±  9%    13.94n ±  7%        ~ (p=0.084 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    18.84n ±  3%    19.48n ±  2%   +3.37% (p=0.037 n=20)
geomean                                14.73n          14.82n         +0.64%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                     │  bart/lpm.bm  │             fast/lpm.bm              │
                                     │    sec/op     │    sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4            12.910n ±  4%   9.540n ±  7%  -26.10% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6             19.76n ± 12%   12.91n ± 24%  -34.69% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4              13.22n ±  3%   10.39n ±  1%  -21.38% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6              18.58n ± 35%   13.89n ± 23%  -25.29% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4      12.96n ±  1%   10.80n ±  5%  -16.60% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6      14.55n ±  1%   10.86n ±  0%  -25.39% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4       13.04n ± 26%   10.44n ± 20%  -19.91% (p=0.006 n=20)
LpmRandomPfxs/1_000/RandomMissIP6       14.39n ±  3%   10.95n ±  5%  -23.91% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4     14.75n ±  2%   11.94n ±  2%  -19.05% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6     14.41n ±  3%   11.05n ±  5%  -23.26% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4      14.62n ±  2%   11.76n ±  1%  -19.59% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6      14.39n ±  2%   10.88n ±  6%  -24.33% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4    9.608n ± 36%   7.603n ± 40%  -20.87% (p=0.006 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6    18.94n ± 39%   13.33n ± 36%  -29.60% (p=0.002 n=20)
LpmRandomPfxs/100_000/RandomMissIP4     14.45n ±  9%   11.74n ± 12%  -18.75% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6     18.84n ±  3%   13.44n ±  7%  -28.71% (p=0.000 n=20)
geomean                                 14.73n         11.23n        -23.73%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                     │ bart/lpm.bm  │             netipds/lpm.bm             │
                                     │    sec/op    │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            12.91n ±  4%    34.18n ± 14%  +164.76% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            19.76n ± 12%    34.84n ± 17%   +76.32% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             13.22n ±  3%    36.43n ±  4%  +175.67% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             18.58n ± 35%    38.71n ± 13%  +108.31% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     12.96n ±  1%    21.36n ±  5%   +64.88% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     14.55n ±  1%    23.06n ±  5%   +58.43% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      13.04n ± 26%    22.44n ±  3%   +72.15% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      14.39n ±  3%    23.07n ±  6%   +60.32% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    14.75n ±  2%    26.01n ±  4%   +76.34% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    14.41n ±  3%    27.15n ±  5%   +88.51% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     14.62n ±  2%    27.83n ±  4%   +90.29% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     14.39n ±  2%    27.05n ±  3%   +88.01% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   9.608n ± 36%   25.980n ± 13%  +170.40% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   18.94n ± 39%    30.82n ±  3%   +62.77% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    14.45n ±  9%    31.06n ±  2%  +114.95% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    18.84n ±  3%    31.82n ±  3%   +68.85% (p=0.000 n=20)
geomean                                14.73n          28.41n         +92.89%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                     │ bart/lpm.bm  │             critbitgo/lpm.bm             │
                                     │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            12.91n ±  4%    128.35n ±  6%   +894.19% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            19.76n ± 12%    200.65n ± 11%   +915.44% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             13.22n ±  3%    412.45n ± 17%  +3021.07% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             18.58n ± 35%    332.55n ± 14%  +1689.35% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     12.96n ±  1%     83.37n ±  7%   +543.54% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     14.55n ±  1%     97.65n ±  4%   +570.90% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      13.04n ± 26%    142.70n ± 11%   +994.74% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      14.39n ±  3%    145.60n ±  8%   +911.81% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    14.75n ±  2%     97.47n ±  5%   +560.81% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    14.41n ±  3%    106.65n ±  6%   +640.37% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     14.62n ±  2%    228.40n ± 12%  +1461.71% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     14.39n ±  2%    196.25n ± 15%  +1264.27% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   9.608n ± 36%   106.100n ±  3%  +1004.29% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   18.94n ± 39%    127.35n ±  3%   +572.56% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    14.45n ±  9%    326.25n ± 11%  +2157.79% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    18.84n ±  3%    267.75n ± 17%  +1320.80% (p=0.000 n=20)
geomean                                14.73n           165.9n        +1026.20%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                     │ bart/lpm.bm  │              lpmtrie/lpm.bm               │
                                     │    sec/op    │     sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4            12.91n ±  4%    219.80n ±   7%  +1602.56% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            19.76n ± 12%    237.20n ±  15%  +1100.40% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             13.22n ±  3%    201.85n ±   4%  +1427.43% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             18.58n ± 35%    188.15n ±  18%   +912.38% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     12.96n ±  1%    111.35n ± 750%   +759.51% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     14.55n ±  1%     83.50n ±  17%   +473.72% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      13.04n ± 26%     83.91n ±   6%   +543.69% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      14.39n ±  3%     72.00n ±  14%   +400.38% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    14.75n ±  2%    121.50n ±  16%   +723.73% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    14.41n ±  3%    111.45n ±   4%   +673.69% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     14.62n ±  2%    129.40n ±  10%   +784.79% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     14.39n ±  2%    105.10n ±   5%   +630.62% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   9.608n ± 36%   170.400n ±   6%  +1673.52% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   18.94n ± 39%    152.50n ±   4%   +705.39% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    14.45n ±  9%    157.35n ±   5%   +988.93% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    18.84n ±  3%    152.10n ±   5%   +707.11% (p=0.000 n=20)
geomean                                14.73n           135.4n          +819.17%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                     │ bart/lpm.bm  │          kentik-patricia/lpm.bm          │
                                     │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            12.91n ±  4%    146.45n ±  9%  +1034.39% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            19.76n ± 12%    198.15n ± 22%   +902.78% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             13.22n ±  3%    115.55n ±  5%   +774.39% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             18.58n ± 35%    128.20n ± 14%   +589.80% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     12.96n ±  1%     67.73n ±  8%   +422.81% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     14.55n ±  1%     77.39n ± 16%   +431.71% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      13.04n ± 26%     54.45n ±  5%   +317.68% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      14.39n ±  3%     59.86n ± 10%   +315.98% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    14.75n ±  2%     85.41n ±  6%   +479.02% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    14.41n ±  3%     95.15n ±  5%   +560.50% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     14.62n ±  2%     74.88n ±  4%   +412.00% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     14.39n ±  2%     82.14n ±  5%   +471.01% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   9.608n ± 36%   108.400n ±  9%  +1028.23% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   18.94n ± 39%    116.20n ±  5%   +513.68% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    14.45n ±  9%     88.34n ±  5%   +511.38% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    18.84n ±  3%    104.05n ±  6%   +452.14% (p=0.000 n=20)
geomean                                14.73n           94.79n         +543.57%
```

## size of the routing tables


`bart.Lite` has the lowest memory consumption under all competitors.


```
goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/size.bm │            lite/size.bm            │
                         │ bytes/route  │ bytes/route  vs base               │
Tier1PfxSize/1_000          105.30 ± 2%    85.39 ± 2%  -18.91% (p=0.002 n=6)
Tier1PfxSize/10_000          83.96 ± 0%    63.94 ± 0%  -23.84% (p=0.002 n=6)
Tier1PfxSize/100_000         56.72 ± 0%    36.97 ± 0%  -34.82% (p=0.002 n=6)
Tier1PfxSize/200_000         49.53 ± 0%    30.07 ± 0%  -39.29% (p=0.002 n=6)
Tier1PfxSize/500_000         42.90 ± 0%    23.63 ± 0%  -44.92% (p=0.002 n=6)
Tier1PfxSize/1_000_000       39.74 ± 0%    20.42 ± 0%  -48.62% (p=0.002 n=6)
RandomPfx4Size/1_000         82.62 ± 2%    61.90 ± 3%  -25.08% (p=0.002 n=6)
RandomPfx4Size/10_000        57.37 ± 0%    38.30 ± 0%  -33.24% (p=0.002 n=6)
RandomPfx4Size/100_000       72.09 ± 0%    49.81 ± 0%  -30.91% (p=0.002 n=6)
RandomPfx4Size/200_000       65.15 ± 0%    43.92 ± 0%  -32.59% (p=0.002 n=6)
RandomPfx4Size/500_000       52.44 ± 0%    31.18 ± 0%  -40.54% (p=0.002 n=6)
RandomPfx4Size/1_000_000     46.46 ± 0%    25.84 ± 0%  -44.38% (p=0.002 n=6)
RandomPfx6Size/1_000         83.74 ± 2%    66.27 ± 3%  -20.86% (p=0.002 n=6)
RandomPfx6Size/10_000       101.40 ± 0%    80.81 ± 0%  -20.31% (p=0.002 n=6)
RandomPfx6Size/100_000       71.69 ± 0%    53.59 ± 0%  -25.25% (p=0.002 n=6)
RandomPfx6Size/200_000       69.38 ± 0%    51.40 ± 0%  -25.92% (p=0.002 n=6)
RandomPfx6Size/500_000       72.35 ± 0%    53.93 ± 0%  -25.46% (p=0.002 n=6)
RandomPfx6Size/1_000_000     78.93 ± 0%    59.52 ± 0%  -24.59% (p=0.002 n=6)
RandomPfxSize/1_000         100.70 ± 2%    80.70 ± 2%  -19.86% (p=0.002 n=6)
RandomPfxSize/10_000         74.11 ± 0%    56.53 ± 0%  -23.72% (p=0.002 n=6)
RandomPfxSize/100_000        84.34 ± 0%    63.46 ± 0%  -24.76% (p=0.002 n=6)
RandomPfxSize/200_000        75.60 ± 0%    54.19 ± 0%  -28.32% (p=0.002 n=6)
RandomPfxSize/500_000        59.59 ± 0%    39.18 ± 0%  -34.25% (p=0.002 n=6)
RandomPfxSize/1_000_000      52.03 ± 0%    31.89 ± 0%  -38.71% (p=0.002 n=6)
geomean                      67.56         46.66       -30.93%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/size.bm │            fast/size.bm             │
                         │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000           105.3 ± 2%    172.2 ± 1%   +63.53% (p=0.002 n=6)
Tier1PfxSize/10_000          83.96 ± 0%   151.10 ± 0%   +79.97% (p=0.002 n=6)
Tier1PfxSize/100_000         56.72 ± 0%   101.80 ± 0%   +79.48% (p=0.002 n=6)
Tier1PfxSize/200_000         49.53 ± 0%    81.43 ± 0%   +64.41% (p=0.002 n=6)
Tier1PfxSize/500_000         42.90 ± 0%    62.27 ± 0%   +45.15% (p=0.002 n=6)
Tier1PfxSize/1_000_000       39.74 ± 0%    52.71 ± 0%   +32.64% (p=0.002 n=6)
RandomPfx4Size/1_000         82.62 ± 2%   147.10 ± 1%   +78.04% (p=0.002 n=6)
RandomPfx4Size/10_000        57.37 ± 0%    73.09 ± 0%   +27.40% (p=0.002 n=6)
RandomPfx4Size/100_000       72.09 ± 0%   132.60 ± 0%   +83.94% (p=0.002 n=6)
RandomPfx4Size/200_000       65.15 ± 0%   130.30 ± 0%  +100.00% (p=0.002 n=6)
RandomPfx4Size/500_000       52.44 ± 0%    88.48 ± 0%   +68.73% (p=0.002 n=6)
RandomPfx4Size/1_000_000     46.46 ± 0%    65.46 ± 0%   +40.90% (p=0.002 n=6)
RandomPfx6Size/1_000         83.74 ± 2%   108.80 ± 2%   +29.93% (p=0.002 n=6)
RandomPfx6Size/10_000        101.4 ± 0%    179.6 ± 0%   +77.12% (p=0.002 n=6)
RandomPfx6Size/100_000       71.69 ± 0%    98.32 ± 0%   +37.15% (p=0.002 n=6)
RandomPfx6Size/200_000       69.38 ± 0%    88.85 ± 0%   +28.06% (p=0.002 n=6)
RandomPfx6Size/500_000       72.35 ± 0%    96.34 ± 0%   +33.16% (p=0.002 n=6)
RandomPfx6Size/1_000_000     78.93 ± 0%   116.70 ± 0%   +47.85% (p=0.002 n=6)
RandomPfxSize/1_000          100.7 ± 2%    168.1 ± 1%   +66.93% (p=0.002 n=6)
RandomPfxSize/10_000         74.11 ± 0%    97.72 ± 0%   +31.86% (p=0.002 n=6)
RandomPfxSize/100_000        84.34 ± 0%   156.60 ± 0%   +85.68% (p=0.002 n=6)
RandomPfxSize/200_000        75.60 ± 0%   146.30 ± 0%   +93.52% (p=0.002 n=6)
RandomPfxSize/500_000        59.59 ± 0%   100.30 ± 0%   +68.32% (p=0.002 n=6)
RandomPfxSize/1_000_000      52.03 ± 0%    74.51 ± 0%   +43.21% (p=0.002 n=6)
geomean                      67.56         106.1        +57.00%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/size.bm │          netipds/size.bm           │
                         │ bytes/route  │ bytes/route  vs base               │
Tier1PfxSize/1_000          105.30 ± 2%    74.18 ± 3%  -29.55% (p=0.002 n=6)
Tier1PfxSize/10_000          83.96 ± 0%    68.94 ± 0%  -17.89% (p=0.002 n=6)
Tier1PfxSize/100_000         56.72 ± 0%    67.66 ± 0%  +19.29% (p=0.002 n=6)
Tier1PfxSize/200_000         49.53 ± 0%    66.91 ± 0%  +35.09% (p=0.002 n=6)
Tier1PfxSize/500_000         42.90 ± 0%    65.35 ± 0%  +52.33% (p=0.002 n=6)
Tier1PfxSize/1_000_000       39.74 ± 0%    63.72 ± 0%  +60.34% (p=0.002 n=6)
RandomPfx4Size/1_000         82.62 ± 2%    61.89 ± 3%  -25.09% (p=0.002 n=6)
RandomPfx4Size/10_000        57.37 ± 0%    51.37 ± 0%  -10.46% (p=0.002 n=6)
RandomPfx4Size/100_000       72.09 ± 0%    48.25 ± 0%  -33.07% (p=0.002 n=6)
RandomPfx4Size/200_000       65.15 ± 0%    47.57 ± 0%  -26.98% (p=0.002 n=6)
RandomPfx4Size/500_000       52.44 ± 0%    46.51 ± 0%  -11.31% (p=0.002 n=6)
RandomPfx4Size/1_000_000     46.46 ± 0%    45.66 ± 0%   -1.72% (p=0.002 n=6)
RandomPfx6Size/1_000         83.74 ± 2%   100.20 ± 2%  +19.66% (p=0.002 n=6)
RandomPfx6Size/10_000       101.40 ± 0%    94.28 ± 0%   -7.02% (p=0.002 n=6)
RandomPfx6Size/100_000       71.69 ± 0%    87.64 ± 0%  +22.25% (p=0.002 n=6)
RandomPfx6Size/200_000       69.38 ± 0%    85.90 ± 0%  +23.81% (p=0.002 n=6)
RandomPfx6Size/500_000       72.35 ± 0%    84.11 ± 0%  +16.25% (p=0.002 n=6)
RandomPfx6Size/1_000_000     78.93 ± 0%    83.23 ± 0%   +5.45% (p=0.002 n=6)
RandomPfxSize/1_000         100.70 ± 2%    74.90 ± 3%  -25.62% (p=0.002 n=6)
RandomPfxSize/10_000         74.11 ± 0%    69.40 ± 0%   -6.36% (p=0.002 n=6)
RandomPfxSize/100_000        84.34 ± 0%    64.29 ± 0%  -23.77% (p=0.002 n=6)
RandomPfxSize/200_000        75.60 ± 0%    61.55 ± 0%  -18.58% (p=0.002 n=6)
RandomPfxSize/500_000        59.59 ± 0%    57.42 ± 0%   -3.64% (p=0.002 n=6)
RandomPfxSize/1_000_000      52.03 ± 0%    53.93 ± 0%   +3.65% (p=0.002 n=6)
geomean                      67.56         66.04        -2.25%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/size.bm │          critbitgo/size.bm          │
                         │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000           105.3 ± 2%    119.7 ± 2%   +13.68% (p=0.002 n=6)
Tier1PfxSize/10_000          83.96 ± 0%   114.80 ± 0%   +36.73% (p=0.002 n=6)
Tier1PfxSize/100_000         56.72 ± 0%   114.40 ± 0%  +101.69% (p=0.002 n=6)
Tier1PfxSize/200_000         49.53 ± 0%   114.40 ± 0%  +130.97% (p=0.002 n=6)
Tier1PfxSize/500_000         42.90 ± 0%   114.40 ± 0%  +166.67% (p=0.002 n=6)
Tier1PfxSize/1_000_000       39.74 ± 0%   114.40 ± 0%  +187.87% (p=0.002 n=6)
RandomPfx4Size/1_000         82.62 ± 2%   116.40 ± 2%   +40.89% (p=0.002 n=6)
RandomPfx4Size/10_000        57.37 ± 0%   112.40 ± 0%   +95.92% (p=0.002 n=6)
RandomPfx4Size/100_000       72.09 ± 0%   112.00 ± 0%   +55.36% (p=0.002 n=6)
RandomPfx4Size/200_000       65.15 ± 0%   112.00 ± 0%   +71.91% (p=0.002 n=6)
RandomPfx4Size/500_000       52.44 ± 0%   112.00 ± 0%  +113.58% (p=0.002 n=6)
RandomPfx4Size/1_000_000     46.46 ± 0%   112.00 ± 0%  +141.07% (p=0.002 n=6)
RandomPfx6Size/1_000         83.74 ± 2%   132.40 ± 2%   +58.11% (p=0.002 n=6)
RandomPfx6Size/10_000        101.4 ± 0%    128.4 ± 0%   +26.63% (p=0.002 n=6)
RandomPfx6Size/100_000       71.69 ± 0%   128.00 ± 0%   +78.55% (p=0.002 n=6)
RandomPfx6Size/200_000       69.38 ± 0%   128.00 ± 0%   +84.49% (p=0.002 n=6)
RandomPfx6Size/500_000       72.35 ± 0%   128.00 ± 0%   +76.92% (p=0.002 n=6)
RandomPfx6Size/1_000_000     78.93 ± 0%   128.00 ± 0%   +62.17% (p=0.002 n=6)
RandomPfxSize/1_000          100.7 ± 2%    119.8 ± 2%   +18.97% (p=0.002 n=6)
RandomPfxSize/10_000         74.11 ± 0%   115.60 ± 0%   +55.98% (p=0.002 n=6)
RandomPfxSize/100_000        84.34 ± 0%   115.30 ± 0%   +36.71% (p=0.002 n=6)
RandomPfxSize/200_000        75.60 ± 0%   115.20 ± 0%   +52.38% (p=0.002 n=6)
RandomPfxSize/500_000        59.59 ± 0%   115.20 ± 0%   +93.32% (p=0.002 n=6)
RandomPfxSize/1_000_000      52.03 ± 0%   115.20 ± 0%  +121.41% (p=0.002 n=6)
geomean                      67.56         118.1        +74.78%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/size.bm │           lpmtrie/size.bm            │
                         │ bytes/route  │ bytes/route   vs base                │
Tier1PfxSize/1_000           105.3 ± 2%    215.6 ±  5%  +104.75% (p=0.002 n=6)
Tier1PfxSize/10_000          83.96 ± 0%   210.50 ±  5%  +150.71% (p=0.002 n=6)
Tier1PfxSize/100_000         56.72 ± 0%   209.90 ±  5%  +270.06% (p=0.002 n=6)
Tier1PfxSize/200_000         49.53 ± 0%   209.20 ±  5%  +322.37% (p=0.002 n=6)
Tier1PfxSize/500_000         42.90 ± 0%   207.90 ±  7%  +384.62% (p=0.002 n=6)
Tier1PfxSize/1_000_000       39.74 ± 0%   205.00 ±  7%  +415.85% (p=0.002 n=6)
RandomPfx4Size/1_000         82.62 ± 2%   205.50 ±  8%  +148.73% (p=0.002 n=6)
RandomPfx4Size/10_000        57.37 ± 0%   186.70 ±  9%  +225.43% (p=0.002 n=6)
RandomPfx4Size/100_000       72.09 ± 0%   179.60 ±  9%  +149.13% (p=0.002 n=6)
RandomPfx4Size/200_000       65.15 ± 0%   178.50 ± 10%  +173.98% (p=0.002 n=6)
RandomPfx4Size/500_000       52.44 ± 0%   176.80 ± 10%  +237.15% (p=0.002 n=6)
RandomPfx4Size/1_000_000     46.46 ± 0%   175.50 ± 10%  +277.74% (p=0.002 n=6)
RandomPfx6Size/1_000         83.74 ± 2%   228.00 ±  8%  +172.27% (p=0.002 n=6)
RandomPfx6Size/10_000        101.4 ± 0%    222.5 ±  8%  +119.43% (p=0.002 n=6)
RandomPfx6Size/100_000       71.69 ± 0%   213.90 ±  9%  +198.37% (p=0.002 n=6)
RandomPfx6Size/200_000       69.38 ± 0%   210.50 ±  9%  +203.40% (p=0.002 n=6)
RandomPfx6Size/500_000       72.35 ± 0%   206.70 ±  9%  +185.69% (p=0.002 n=6)
RandomPfx6Size/1_000_000     78.93 ± 0%   204.90 ±  9%  +159.60% (p=0.002 n=6)
RandomPfxSize/1_000          100.7 ± 2%    215.3 ±  5%  +113.80% (p=0.002 n=6)
RandomPfxSize/10_000         74.11 ± 0%   210.30 ±  5%  +183.77% (p=0.002 n=6)
RandomPfxSize/100_000        84.34 ± 0%   203.50 ±  7%  +141.29% (p=0.002 n=6)
RandomPfxSize/200_000        75.60 ± 0%   198.70 ±  8%  +162.83% (p=0.002 n=6)
RandomPfxSize/500_000        59.59 ± 0%   189.90 ±  9%  +218.68% (p=0.002 n=6)
RandomPfxSize/1_000_000      52.03 ± 0%   181.40 ± 10%  +248.65% (p=0.002 n=6)
geomean                      67.56         201.4        +198.08%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/size.bm │       kentik-patricia/size.bm       │
                         │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000           105.3 ± 2%    145.5 ± 1%   +38.18% (p=0.002 n=6)
Tier1PfxSize/10_000          83.96 ± 0%   200.30 ± 0%  +138.57% (p=0.002 n=6)
Tier1PfxSize/100_000         56.72 ± 0%   164.00 ± 0%  +189.14% (p=0.002 n=6)
Tier1PfxSize/200_000         49.53 ± 0%   163.90 ± 0%  +230.91% (p=0.002 n=6)
Tier1PfxSize/500_000         42.90 ± 0%   144.50 ± 0%  +236.83% (p=0.002 n=6)
Tier1PfxSize/1_000_000       39.74 ± 0%   144.20 ± 0%  +262.86% (p=0.002 n=6)
RandomPfx4Size/1_000         82.62 ± 2%   140.90 ± 1%   +70.54% (p=0.002 n=6)
RandomPfx4Size/10_000        57.37 ± 0%   109.60 ± 0%   +91.04% (p=0.002 n=6)
RandomPfx4Size/100_000       72.09 ± 0%   139.90 ± 0%   +94.06% (p=0.002 n=6)
RandomPfx4Size/200_000       65.15 ± 0%   139.80 ± 0%  +114.58% (p=0.002 n=6)
RandomPfx4Size/500_000       52.44 ± 0%   139.70 ± 0%  +166.40% (p=0.002 n=6)
RandomPfx4Size/1_000_000     46.46 ± 0%   139.60 ± 0%  +200.47% (p=0.002 n=6)
RandomPfx6Size/1_000         83.74 ± 2%   157.30 ± 1%   +87.84% (p=0.002 n=6)
RandomPfx6Size/10_000        101.4 ± 0%    201.4 ± 0%   +98.62% (p=0.002 n=6)
RandomPfx6Size/100_000       71.69 ± 0%   160.80 ± 0%  +124.30% (p=0.002 n=6)
RandomPfx6Size/200_000       69.38 ± 0%   160.80 ± 0%  +131.77% (p=0.002 n=6)
RandomPfx6Size/500_000       72.35 ± 0%   156.50 ± 0%  +116.31% (p=0.002 n=6)
RandomPfx6Size/1_000_000     78.93 ± 0%   156.50 ± 0%   +98.28% (p=0.002 n=6)
RandomPfxSize/1_000          100.7 ± 2%    144.7 ± 1%   +43.69% (p=0.002 n=6)
RandomPfxSize/10_000         74.11 ± 0%   140.20 ± 0%   +89.18% (p=0.002 n=6)
RandomPfxSize/100_000        84.34 ± 0%   180.00 ± 0%  +113.42% (p=0.002 n=6)
RandomPfxSize/200_000        75.60 ± 0%   180.00 ± 0%  +138.10% (p=0.002 n=6)
RandomPfxSize/500_000        59.59 ± 0%   144.00 ± 0%  +141.65% (p=0.002 n=6)
RandomPfxSize/1_000_000      52.03 ± 0%   144.00 ± 0%  +176.76% (p=0.002 n=6)
geomean                      67.56         152.8       +126.19%
```

## update, insert/delete

`bart.Lite` is the fastest algorithm for updates.

```
goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │           lite/update.bm            │
                         │   sec/route    │  sec/route    vs base               │
InsertRandomPfxs/1_000       203.2n ±  1%   164.2n ±  0%  -19.21% (p=0.002 n=6)
InsertRandomPfxs/10_000      150.5n ±  2%   116.2n ±  3%  -22.76% (p=0.002 n=6)
InsertRandomPfxs/100_000     227.5n ±  5%   164.2n ±  9%  -27.82% (p=0.002 n=6)
InsertRandomPfxs/200_000     265.9n ± 11%   191.7n ±  3%  -27.91% (p=0.002 n=6)
DeleteRandomPfxs/1_000       141.6n ±  2%   124.0n ±  1%  -12.39% (p=0.002 n=6)
DeleteRandomPfxs/10_000      90.20n ±  4%   76.64n ±  2%  -15.04% (p=0.002 n=6)
DeleteRandomPfxs/100_000     202.1n ±  4%   162.4n ± 13%  -19.62% (p=0.002 n=6)
DeleteRandomPfxs/200_000     233.9n ±  6%   198.3n ±  6%  -15.24% (p=0.002 n=6)
geomean                      180.3n         143.9n        -20.19%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │            fast/update.bm            │
                         │   sec/route    │   sec/route    vs base               │
InsertRandomPfxs/1_000       203.2n ±  1%    286.8n ±  1%  +41.11% (p=0.002 n=6)
InsertRandomPfxs/10_000      150.5n ±  2%    195.1n ±  2%  +29.67% (p=0.002 n=6)
InsertRandomPfxs/100_000     227.5n ±  5%    319.4n ±  6%  +40.42% (p=0.002 n=6)
InsertRandomPfxs/200_000     265.9n ± 11%    376.6n ±  8%  +41.64% (p=0.002 n=6)
DeleteRandomPfxs/1_000       141.6n ±  2%    238.9n ±  4%  +68.75% (p=0.002 n=6)
DeleteRandomPfxs/10_000      90.20n ±  4%   137.65n ±  3%  +52.61% (p=0.002 n=6)
DeleteRandomPfxs/100_000     202.1n ±  4%    315.9n ± 16%  +56.31% (p=0.002 n=6)
DeleteRandomPfxs/200_000     233.9n ±  6%    309.4n ±  8%  +32.27% (p=0.002 n=6)
geomean                      180.3n          261.2n        +44.85%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │          netipds/update.bm          │
                         │   sec/route    │  sec/route    vs base               │
InsertRandomPfxs/1_000       203.2n ±  1%    209.3n ± 1%   +2.98% (p=0.002 n=6)
InsertRandomPfxs/10_000      150.5n ±  2%    236.7n ± 0%  +57.31% (p=0.002 n=6)
InsertRandomPfxs/100_000     227.5n ±  5%    304.6n ± 2%  +33.87% (p=0.002 n=6)
InsertRandomPfxs/200_000     265.9n ± 11%    398.6n ± 1%  +49.93% (p=0.002 n=6)
DeleteRandomPfxs/1_000       141.6n ±  2%    120.5n ± 2%  -14.87% (p=0.002 n=6)
DeleteRandomPfxs/10_000      90.20n ±  4%   157.05n ± 2%  +74.11% (p=0.002 n=6)
DeleteRandomPfxs/100_000     202.1n ±  4%    283.5n ± 4%  +40.28% (p=0.002 n=6)
DeleteRandomPfxs/200_000     233.9n ±  6%    391.2n ± 1%  +67.24% (p=0.002 n=6)
geomean                      180.3n          244.2n       +35.41%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │          critbitgo/update.bm          │
                         │   sec/route    │   sec/route    vs base                │
InsertRandomPfxs/1_000       203.2n ±  1%    264.3n ±  1%   +30.06% (p=0.002 n=6)
InsertRandomPfxs/10_000      150.5n ±  2%    321.2n ±  2%  +113.39% (p=0.002 n=6)
InsertRandomPfxs/100_000     227.5n ±  5%    515.6n ±  2%  +126.64% (p=0.002 n=6)
InsertRandomPfxs/200_000     265.9n ± 11%    659.8n ±  1%  +148.19% (p=0.002 n=6)
DeleteRandomPfxs/1_000       141.6n ±  2%    129.2n ±  1%    -8.76% (p=0.002 n=6)
DeleteRandomPfxs/10_000      90.20n ±  4%   168.70n ±  4%   +87.03% (p=0.002 n=6)
DeleteRandomPfxs/100_000     202.1n ±  4%    362.7n ±  3%   +79.47% (p=0.002 n=6)
DeleteRandomPfxs/200_000     233.9n ±  6%    461.8n ± 10%   +97.41% (p=0.002 n=6)
geomean                      180.3n          318.3n         +76.55%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │          lpmtrie/update.bm           │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000       203.2n ±  1%    377.9n ± 1%   +85.93% (p=0.002 n=6)
InsertRandomPfxs/10_000      150.5n ±  2%    404.2n ± 2%  +168.60% (p=0.002 n=6)
InsertRandomPfxs/100_000     227.5n ±  5%    674.1n ± 4%  +196.29% (p=0.002 n=6)
InsertRandomPfxs/200_000     265.9n ± 11%    793.9n ± 4%  +198.63% (p=0.002 n=6)
DeleteRandomPfxs/1_000       141.6n ±  2%    126.6n ± 3%   -10.56% (p=0.002 n=6)
DeleteRandomPfxs/10_000      90.20n ±  4%   185.15n ± 3%  +105.27% (p=0.002 n=6)
DeleteRandomPfxs/100_000     202.1n ±  4%    503.6n ± 5%  +149.18% (p=0.002 n=6)
DeleteRandomPfxs/200_000     233.9n ±  6%    690.2n ± 6%  +195.04% (p=0.002 n=6)
geomean                      180.3n          400.8n       +122.30%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │      kentik-patricia/update.bm       │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000       203.2n ±  1%    248.9n ± 3%   +22.48% (p=0.002 n=6)
InsertRandomPfxs/10_000      150.5n ±  2%    326.3n ± 1%  +116.78% (p=0.002 n=6)
InsertRandomPfxs/100_000     227.5n ±  5%    532.5n ± 2%  +134.09% (p=0.002 n=6)
InsertRandomPfxs/200_000     265.9n ± 11%    662.1n ± 3%  +149.03% (p=0.002 n=6)
DeleteRandomPfxs/1_000       141.6n ±  2%    263.3n ± 3%   +85.95% (p=0.002 n=6)
DeleteRandomPfxs/10_000      90.20n ±  4%   305.75n ± 1%  +238.97% (p=0.002 n=6)
DeleteRandomPfxs/100_000     202.1n ±  4%    652.9n ± 1%  +223.06% (p=0.002 n=6)
DeleteRandomPfxs/200_000     233.9n ±  6%    779.7n ± 4%  +233.28% (p=0.002 n=6)
geomean                      180.3n          430.2n       +138.60%
```

