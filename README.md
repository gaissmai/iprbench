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
                                     │ bart/lpm.bm  │             lite/lpm.bm             │
                                     │    sec/op    │   sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4            21.71n ± 13%   20.95n ± 6%        ~ (p=0.075 n=10)
LpmTier1Pfxs/RandomMatchIP6            40.45n ±  6%   35.95n ± 2%  -11.12% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             21.45n ±  8%   20.27n ± 2%   -5.48% (p=0.001 n=10)
LpmTier1Pfxs/RandomMissIP6             38.64n ± 10%   37.63n ± 4%   -2.63% (p=0.015 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     20.20n ±  2%   20.55n ± 5%        ~ (p=0.218 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     18.29n ±  4%   18.55n ± 4%        ~ (p=0.494 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      19.37n ±  5%   20.58n ± 1%   +6.27% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      17.01n ±  4%   18.18n ± 2%   +6.88% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    22.19n ±  7%   21.21n ± 5%   -4.44% (p=0.005 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    22.89n ±  9%   22.13n ± 2%   -3.30% (p=0.017 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     21.09n ±  4%   21.79n ± 4%        ~ (p=0.065 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     21.71n ± 10%   22.53n ± 1%        ~ (p=0.143 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   15.84n ± 11%   16.86n ± 2%   +6.47% (p=0.022 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   33.95n ± 12%   34.42n ± 4%        ~ (p=0.724 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    26.13n ±  8%   27.72n ± 2%   +6.11% (p=0.015 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    35.87n ± 12%   37.08n ± 1%        ~ (p=0.271 n=10)
geomean                                23.77n         23.89n        +0.51%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                     │ bart/lpm.bm  │             fast/lpm.bm             │
                                     │    sec/op    │   sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4            21.71n ± 13%   19.70n ± 3%   -9.28% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            40.45n ±  6%   33.18n ± 7%  -17.97% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             21.45n ±  8%   18.27n ± 2%  -14.83% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             38.64n ± 10%   32.67n ± 3%  -15.46% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     20.20n ±  2%   16.78n ± 3%  -16.93% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     18.29n ±  4%   14.14n ± 3%  -22.64% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      19.37n ±  5%   16.53n ± 1%  -14.64% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      17.01n ±  4%   13.91n ± 2%  -18.25% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    22.19n ±  7%   16.92n ± 2%  -23.75% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    22.89n ±  9%   17.80n ± 1%  -22.26% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     21.09n ±  4%   17.03n ± 2%  -19.25% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     21.71n ± 10%   17.55n ± 2%  -19.16% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   15.84n ± 11%   13.66n ± 3%  -13.79% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   33.95n ± 12%   31.04n ± 3%   -8.58% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    26.13n ±  8%   24.75n ± 5%   -5.28% (p=0.001 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    35.87n ± 12%   32.12n ± 1%  -10.44% (p=0.000 n=10)
geomean                                23.77n         19.98n       -15.94%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                     │ bart/lpm.bm  │            netipds/lpm.bm             │
                                     │    sec/op    │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4            21.71n ± 13%   139.95n ± 2%  +544.63% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            40.45n ±  6%   113.95n ± 1%  +181.71% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             21.45n ±  8%   154.55n ± 1%  +620.68% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             38.64n ± 10%   140.80n ± 1%  +264.34% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     20.20n ±  2%    54.17n ± 1%  +168.14% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     18.29n ±  4%    46.36n ± 1%  +153.54% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      19.37n ±  5%    56.40n ± 1%  +191.22% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      17.01n ±  4%    47.30n ± 1%  +178.02% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    22.19n ±  7%    77.72n ± 1%  +250.27% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    22.89n ±  9%    68.91n ± 1%  +201.05% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     21.09n ±  4%    81.03n ± 1%  +284.21% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     21.71n ± 10%    69.28n ± 1%  +219.09% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   15.84n ± 11%    86.53n ± 2%  +446.31% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   33.95n ± 12%   107.20n ± 1%  +215.71% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    26.13n ±  8%   115.10n ± 2%  +340.57% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    35.87n ± 12%   108.80n ± 1%  +203.36% (p=0.000 n=10)
geomean                                23.77n          85.48n       +259.61%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                     │ bart/lpm.bm  │             critbitgo/lpm.bm             │
                                     │    sec/op    │    sec/op      vs base                   │
LpmTier1Pfxs/RandomMatchIP4            21.71n ± 13%    410.75n ± 7%   +1791.99% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            40.45n ±  6%    367.55n ± 2%    +808.65% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             21.45n ±  8%   2923.00n ± 2%  +13530.22% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             38.64n ± 10%   1169.50n ± 3%   +2926.26% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     20.20n ±  2%    119.35n ± 1%    +490.84% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     18.29n ±  4%    122.25n ± 1%    +568.58% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      19.37n ±  5%    220.00n ± 1%   +1036.07% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      17.01n ±  4%    173.60n ± 1%    +920.28% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    22.19n ±  7%    169.40n ± 1%    +663.41% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    22.89n ±  9%    158.75n ± 1%    +593.53% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     21.09n ±  4%    426.05n ± 1%   +1920.15% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     21.71n ± 10%    323.50n ± 4%   +1390.10% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   15.84n ± 11%    245.05n ± 4%   +1447.03% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   33.95n ± 12%    240.80n ± 1%    +609.17% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    26.13n ±  8%    836.60n ± 3%   +3102.30% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    35.87n ± 12%    514.95n ± 2%   +1335.80% (p=0.000 n=10)
geomean                                23.77n           336.6n        +1315.94%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                     │ bart/lpm.bm  │             lpmtrie/lpm.bm              │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            21.71n ± 13%   588.20n ±  4%  +2609.35% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            40.45n ±  6%   484.75n ±  6%  +1098.39% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             21.45n ±  8%   409.30n ± 10%  +1808.60% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             38.64n ± 10%   319.90n ±  7%   +727.79% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     20.20n ±  2%   106.95n ±  2%   +429.46% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     18.29n ±  4%    94.54n ±  2%   +417.04% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      19.37n ±  5%   104.00n ±  1%   +437.05% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      17.01n ±  4%    94.71n ±  1%   +456.63% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    22.19n ±  7%   167.20n ±  1%   +653.49% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    22.89n ±  9%   139.95n ±  1%   +511.40% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     21.09n ±  4%   163.50n ±  1%   +675.25% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     21.71n ± 10%   135.15n ±  1%   +522.52% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   15.84n ± 11%   252.70n ±  3%  +1495.33% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   33.95n ± 12%   207.35n ±  1%   +510.66% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    26.13n ±  8%   234.00n ±  1%   +795.69% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    35.87n ± 12%   211.70n ±  2%   +490.27% (p=0.000 n=10)
geomean                                23.77n          196.6n         +727.21%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                     │ bart/lpm.bm  │         kentik-patricia/lpm.bm         │
                                     │    sec/op    │    sec/op     vs base                  │
LpmTier1Pfxs/RandomMatchIP4            21.71n ± 13%   438.80n ± 3%  +1921.19% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            40.45n ±  6%   387.70n ± 1%   +858.47% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             21.45n ±  8%   292.60n ± 1%  +1264.42% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             38.64n ± 10%   271.45n ± 1%   +602.42% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     20.20n ±  2%   110.55n ± 3%   +447.28% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     18.29n ±  4%    95.78n ± 1%   +423.82% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      19.37n ±  5%    94.98n ± 4%   +390.47% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      17.01n ±  4%    83.01n ± 1%   +387.86% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    22.19n ±  7%   172.15n ± 3%   +675.80% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    22.89n ±  9%   153.15n ± 0%   +569.07% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     21.09n ±  4%   150.10n ± 1%   +611.71% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     21.71n ± 10%   138.25n ± 3%   +536.80% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   15.84n ± 11%   261.90n ± 4%  +1553.41% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   33.95n ± 12%   237.25n ± 1%   +598.72% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    26.13n ±  8%   223.55n ± 7%   +755.69% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    35.87n ± 12%   220.25n ± 1%   +514.11% (p=0.000 n=10)
geomean                                23.77n          185.0n        +678.46%
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
Tier1PfxSize/200_000         49.53 ± 0%   164.00 ± 0%  +231.11% (p=0.002 n=6)
Tier1PfxSize/500_000         42.90 ± 0%   144.10 ± 0%  +235.90% (p=0.002 n=6)
Tier1PfxSize/1_000_000       39.74 ± 0%   144.30 ± 0%  +263.11% (p=0.002 n=6)
RandomPfx4Size/1_000         82.62 ± 2%   140.90 ± 1%   +70.54% (p=0.002 n=6)
RandomPfx4Size/10_000        57.37 ± 0%   109.60 ± 0%   +91.04% (p=0.002 n=6)
RandomPfx4Size/100_000       72.09 ± 0%   139.90 ± 0%   +94.06% (p=0.002 n=6)
RandomPfx4Size/200_000       65.15 ± 0%   139.80 ± 0%  +114.58% (p=0.002 n=6)
RandomPfx4Size/500_000       52.44 ± 0%   139.50 ± 0%  +166.02% (p=0.002 n=6)
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
geomean                      67.56         152.8       +126.17%
```

## update, insert/delete

`bart.Lite` is the fastest algorithm for updates.

```
goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │           lite/update.bm           │
                         │   sec/route    │  sec/route   vs base               │
InsertRandomPfxs/1_000       196.8n ±  3%   164.5n ± 1%  -16.43% (p=0.002 n=6)
InsertRandomPfxs/10_000      151.9n ±  2%   115.3n ± 1%  -24.09% (p=0.002 n=6)
InsertRandomPfxs/100_000     222.0n ± 12%   162.7n ± 4%  -26.70% (p=0.002 n=6)
InsertRandomPfxs/200_000     270.1n ± 12%   193.0n ± 4%  -28.56% (p=0.002 n=6)
DeleteRandomPfxs/1_000       139.5n ±  3%   123.1n ± 2%  -11.82% (p=0.002 n=6)
DeleteRandomPfxs/10_000      88.18n ±  4%   76.37n ± 3%  -13.39% (p=0.002 n=6)
DeleteRandomPfxs/100_000     210.7n ± 10%   161.3n ± 9%  -23.40% (p=0.002 n=6)
DeleteRandomPfxs/200_000     267.6n ± 13%   212.7n ± 8%  -20.53% (p=0.002 n=6)
geomean                      182.7n         144.7n       -20.83%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │            fast/update.bm            │
                         │   sec/route    │   sec/route    vs base               │
InsertRandomPfxs/1_000       196.8n ±  3%    286.4n ±  1%  +45.49% (p=0.002 n=6)
InsertRandomPfxs/10_000      151.9n ±  2%    199.4n ±  1%  +31.30% (p=0.002 n=6)
InsertRandomPfxs/100_000     222.0n ± 12%    330.1n ±  3%  +48.73% (p=0.002 n=6)
InsertRandomPfxs/200_000     270.1n ± 12%    374.9n ±  4%  +38.82% (p=0.002 n=6)
DeleteRandomPfxs/1_000       139.5n ±  3%    244.7n ±  3%  +75.35% (p=0.002 n=6)
DeleteRandomPfxs/10_000      88.18n ±  4%   138.80n ±  1%  +57.41% (p=0.002 n=6)
DeleteRandomPfxs/100_000     210.7n ± 10%    344.7n ±  8%  +63.64% (p=0.002 n=6)
DeleteRandomPfxs/200_000     267.6n ± 13%    304.3n ± 18%  +13.71% (p=0.002 n=6)
geomean                      182.7n          266.2n        +45.65%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │          netipds/update.bm           │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000       196.8n ±  3%    210.2n ± 2%    +6.78% (p=0.002 n=6)
InsertRandomPfxs/10_000      151.9n ±  2%    239.4n ± 1%   +57.57% (p=0.002 n=6)
InsertRandomPfxs/100_000     222.0n ± 12%    319.4n ± 2%   +43.91% (p=0.002 n=6)
InsertRandomPfxs/200_000     270.1n ± 12%    428.4n ± 2%   +58.63% (p=0.002 n=6)
DeleteRandomPfxs/1_000       139.5n ±  3%    135.6n ± 1%    -2.83% (p=0.026 n=6)
DeleteRandomPfxs/10_000      88.18n ±  4%   182.65n ± 1%  +107.13% (p=0.002 n=6)
DeleteRandomPfxs/100_000     210.7n ± 10%    305.1n ± 1%   +44.81% (p=0.002 n=6)
DeleteRandomPfxs/200_000     267.6n ± 13%    407.9n ± 1%   +52.38% (p=0.002 n=6)
geomean                      182.7n          260.5n        +42.56%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │          critbitgo/update.bm          │
                         │   sec/route    │   sec/route    vs base                │
InsertRandomPfxs/1_000       196.8n ±  3%    271.5n ±  3%   +37.92% (p=0.002 n=6)
InsertRandomPfxs/10_000      151.9n ±  2%    335.1n ±  4%  +120.61% (p=0.002 n=6)
InsertRandomPfxs/100_000     222.0n ± 12%    513.9n ±  4%  +131.56% (p=0.002 n=6)
InsertRandomPfxs/200_000     270.1n ± 12%    662.1n ±  3%  +145.11% (p=0.002 n=6)
DeleteRandomPfxs/1_000       139.5n ±  3%    129.2n ±  2%    -7.42% (p=0.002 n=6)
DeleteRandomPfxs/10_000      88.18n ±  4%   171.50n ±  2%   +94.49% (p=0.002 n=6)
DeleteRandomPfxs/100_000     210.7n ± 10%    375.5n ±  3%   +78.23% (p=0.002 n=6)
DeleteRandomPfxs/200_000     267.6n ± 13%    474.9n ± 10%   +77.43% (p=0.002 n=6)
geomean                      182.7n          324.3n         +77.46%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │           lpmtrie/update.bm           │
                         │   sec/route    │   sec/route    vs base                │
InsertRandomPfxs/1_000       196.8n ±  3%    371.9n ±  2%   +88.93% (p=0.002 n=6)
InsertRandomPfxs/10_000      151.9n ±  2%    397.6n ±  2%  +161.78% (p=0.002 n=6)
InsertRandomPfxs/100_000     222.0n ± 12%    667.8n ±  3%  +200.88% (p=0.002 n=6)
InsertRandomPfxs/200_000     270.1n ± 12%    783.1n ±  4%  +189.91% (p=0.002 n=6)
DeleteRandomPfxs/1_000       139.5n ±  3%    121.3n ±  3%   -13.08% (p=0.002 n=6)
DeleteRandomPfxs/10_000      88.18n ±  4%   181.45n ± 10%  +105.77% (p=0.002 n=6)
DeleteRandomPfxs/100_000     210.7n ± 10%    502.8n ±  9%  +138.69% (p=0.002 n=6)
DeleteRandomPfxs/200_000     267.6n ± 13%    649.4n ±  7%  +142.61% (p=0.002 n=6)
geomean                      182.7n          391.8n        +114.42%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │       kentik-patricia/update.bm       │
                         │   sec/route    │   sec/route    vs base                │
InsertRandomPfxs/1_000       196.8n ±  3%    265.8n ±  1%   +35.03% (p=0.002 n=6)
InsertRandomPfxs/10_000      151.9n ±  2%    331.2n ±  1%  +118.04% (p=0.002 n=6)
InsertRandomPfxs/100_000     222.0n ± 12%    546.0n ±  2%  +146.00% (p=0.002 n=6)
InsertRandomPfxs/200_000     270.1n ± 12%    678.0n ± 13%  +151.04% (p=0.002 n=6)
DeleteRandomPfxs/1_000       139.5n ±  3%    288.2n ±  2%  +106.49% (p=0.002 n=6)
DeleteRandomPfxs/10_000      88.18n ±  4%   330.70n ±  3%  +275.03% (p=0.002 n=6)
DeleteRandomPfxs/100_000     210.7n ± 10%    699.0n ±  3%  +231.83% (p=0.002 n=6)
DeleteRandomPfxs/200_000     267.6n ± 13%    824.3n ±  5%  +207.98% (p=0.002 n=6)
geomean                      182.7n          453.5n        +148.17%
```

