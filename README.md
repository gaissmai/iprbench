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
                                     │ bart/lpm.bm │             lite/lpm.bm             │
                                     │   sec/op    │   sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4            20.95n ± 2%   20.34n ± 3%   -2.93% (p=0.003 n=10)
LpmTier1Pfxs/RandomMatchIP6            36.86n ± 4%   36.16n ± 4%        ~ (p=0.089 n=10)
LpmTier1Pfxs/RandomMissIP4             20.07n ± 2%   20.03n ± 2%        ~ (p=0.315 n=10)
LpmTier1Pfxs/RandomMissIP6             35.68n ± 4%   35.87n ± 5%        ~ (p=0.971 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     18.91n ± 2%   20.13n ± 2%   +6.48% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     16.38n ± 1%   17.29n ± 1%   +5.53% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      18.80n ± 2%   19.18n ± 1%   +1.97% (p=0.007 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      16.32n ± 1%   16.97n ± 1%   +3.98% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    19.92n ± 1%   22.33n ± 1%  +12.10% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    21.69n ± 2%   21.50n ± 2%        ~ (p=0.343 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     20.34n ± 1%   20.07n ± 1%   -1.35% (p=0.045 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     20.87n ± 1%   20.82n ± 2%        ~ (p=0.895 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   15.73n ± 3%   16.22n ± 3%   +3.08% (p=0.022 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   34.17n ± 3%   33.16n ± 3%   -2.98% (p=0.043 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    26.69n ± 3%   25.57n ± 3%   -4.20% (p=0.001 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    37.31n ± 2%   34.38n ± 3%   -7.85% (p=0.000 n=10)
geomean                                22.75n        22.89n        +0.59%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                     │ bart/lpm.bm │             fast/lpm.bm             │
                                     │   sec/op    │   sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4            20.95n ± 2%   19.76n ± 3%   -5.70% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            36.86n ± 4%   36.25n ± 3%        ~ (p=0.218 n=10)
LpmTier1Pfxs/RandomMissIP4             20.07n ± 2%   17.49n ± 3%  -12.85% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             35.68n ± 4%   38.24n ± 4%   +7.19% (p=0.002 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     18.91n ± 2%   16.89n ± 2%  -10.63% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     16.38n ± 1%   13.44n ± 1%  -17.98% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      18.80n ± 2%   17.20n ± 2%   -8.53% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      16.32n ± 1%   13.47n ± 1%  -17.49% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    19.92n ± 1%   18.14n ± 2%   -8.96% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    21.69n ± 2%   18.72n ± 2%  -13.67% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     20.34n ± 1%   18.73n ± 1%   -7.96% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     20.87n ± 1%   18.65n ± 1%  -10.66% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   15.73n ± 3%   14.34n ± 3%   -8.84% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   34.17n ± 3%   32.67n ± 3%   -4.40% (p=0.005 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    26.69n ± 3%   26.77n ± 2%        ~ (p=0.839 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    37.31n ± 2%   33.63n ± 5%   -9.88% (p=0.000 n=10)
geomean                                22.75n        20.83n        -8.44%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                     │ bart/lpm.bm │            netipds/lpm.bm             │
                                     │   sec/op    │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4            20.95n ± 2%   140.40n ± 1%  +570.01% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            36.86n ± 4%   113.65n ± 1%  +208.37% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             20.07n ± 2%   152.85n ± 0%  +661.39% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             35.68n ± 4%   137.10n ± 1%  +284.30% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     18.91n ± 2%    53.54n ± 1%  +183.23% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     16.38n ± 1%    45.25n ± 1%  +176.25% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      18.80n ± 2%    56.00n ± 2%  +197.77% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      16.32n ± 1%    47.62n ± 1%  +191.70% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    19.92n ± 1%    77.83n ± 1%  +290.71% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    21.69n ± 2%    69.19n ± 1%  +219.07% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     20.34n ± 1%    81.34n ± 0%  +299.80% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     20.87n ± 1%    69.10n ± 1%  +231.10% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   15.73n ± 3%    85.97n ± 2%  +446.57% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   34.17n ± 3%   105.70n ± 1%  +209.29% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    26.69n ± 3%   115.35n ± 1%  +332.18% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    37.31n ± 2%   107.45n ± 1%  +187.95% (p=0.000 n=10)
geomean                                22.75n         84.96n       +273.37%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                     │ bart/lpm.bm │             critbitgo/lpm.bm              │
                                     │   sec/op    │     sec/op      vs base                   │
LpmTier1Pfxs/RandomMatchIP4            20.95n ± 2%    378.90n ± 10%   +1708.16% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            36.86n ± 4%    363.25n ±  2%    +885.62% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             20.07n ± 2%   2828.50n ±  3%  +13989.66% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             35.68n ± 4%   1163.50n ±  5%   +3161.39% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     18.91n ± 2%    115.45n ±  4%    +510.69% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     16.38n ± 1%    121.30n ±  1%    +640.54% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      18.80n ± 2%    217.85n ±  1%   +1058.47% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      16.32n ± 1%    174.25n ±  1%    +967.38% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    19.92n ± 1%    173.20n ±  0%    +769.48% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    21.69n ± 2%    161.55n ±  1%    +644.99% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     20.34n ± 1%    405.65n ±  1%   +1893.86% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     20.87n ± 1%    319.95n ±  2%   +1433.06% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   15.73n ± 3%    254.50n ±  1%   +1517.93% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   34.17n ± 3%    243.75n ±  4%    +613.24% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    26.69n ± 3%    908.80n ±  3%   +3305.02% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    37.31n ± 2%    555.80n ±  1%   +1389.48% (p=0.000 n=10)
geomean                                22.75n          336.8n         +1380.14%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                     │ bart/lpm.bm │             lpmtrie/lpm.bm             │
                                     │   sec/op    │    sec/op     vs base                  │
LpmTier1Pfxs/RandomMatchIP4            20.95n ± 2%   645.45n ± 7%  +2980.17% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            36.86n ± 4%   529.15n ± 4%  +1335.76% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             20.07n ± 2%   375.65n ± 9%  +1771.23% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             35.68n ± 4%   308.30n ± 2%   +764.19% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     18.91n ± 2%   110.25n ± 1%   +483.18% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     16.38n ± 1%    96.13n ± 1%   +486.87% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      18.80n ± 2%   105.50n ± 1%   +461.02% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      16.32n ± 1%    94.70n ± 1%   +480.06% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    19.92n ± 1%   171.75n ± 1%   +762.20% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    21.69n ± 2%   143.25n ± 1%   +560.59% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     20.34n ± 1%   168.30n ± 1%   +727.23% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     20.87n ± 1%   141.40n ± 1%   +577.53% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   15.73n ± 3%   259.80n ± 1%  +1551.62% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   34.17n ± 3%   208.00n ± 2%   +508.63% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    26.69n ± 3%   240.05n ± 1%   +799.40% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    37.31n ± 2%   206.05n ± 1%   +452.19% (p=0.000 n=10)
geomean                                22.75n         200.0n        +778.97%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                     │ bart/lpm.bm │         kentik-patricia/lpm.bm         │
                                     │   sec/op    │    sec/op     vs base                  │
LpmTier1Pfxs/RandomMatchIP4            20.95n ± 2%   445.30n ± 2%  +2025.03% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            36.86n ± 4%   408.50n ± 2%  +1008.40% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             20.07n ± 2%   306.25n ± 2%  +1425.53% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             35.68n ± 4%   293.10n ± 1%   +721.58% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     18.91n ± 2%   109.75n ± 1%   +480.53% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     16.38n ± 1%    96.99n ± 1%   +492.16% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      18.80n ± 2%    97.52n ± 1%   +418.61% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      16.32n ± 1%    86.73n ± 0%   +431.27% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    19.92n ± 1%   179.65n ± 1%   +801.86% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    21.69n ± 2%   153.10n ± 1%   +606.02% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     20.34n ± 1%   154.70n ± 1%   +660.38% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     20.87n ± 1%   138.25n ± 2%   +562.43% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   15.73n ± 3%   263.05n ± 1%  +1572.28% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   34.17n ± 3%   239.05n ± 1%   +599.49% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    26.69n ± 3%   214.40n ± 1%   +703.30% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    37.31n ± 2%   205.85n ± 1%   +451.65% (p=0.000 n=10)
geomean                                22.75n         187.8n        +725.48%
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
Tier1PfxSize/200_000         49.53 ± 0%   163.70 ± 0%  +230.51% (p=0.002 n=6)
Tier1PfxSize/500_000         42.90 ± 0%   144.50 ± 0%  +236.83% (p=0.002 n=6)
Tier1PfxSize/1_000_000       39.74 ± 0%   144.30 ± 0%  +263.11% (p=0.002 n=6)
RandomPfx4Size/1_000         82.62 ± 2%   140.90 ± 1%   +70.54% (p=0.002 n=6)
RandomPfx4Size/10_000        57.37 ± 0%   109.60 ± 0%   +91.04% (p=0.002 n=6)
RandomPfx4Size/100_000       72.09 ± 0%   139.90 ± 0%   +94.06% (p=0.002 n=6)
RandomPfx4Size/200_000       65.15 ± 0%   139.80 ± 0%  +114.58% (p=0.002 n=6)
RandomPfx4Size/500_000       52.44 ± 0%   139.50 ± 0%  +166.02% (p=0.002 n=6)
RandomPfx4Size/1_000_000     46.46 ± 0%   139.70 ± 0%  +200.69% (p=0.002 n=6)
RandomPfx6Size/1_000         83.74 ± 2%   157.30 ± 1%   +87.84% (p=0.002 n=6)
RandomPfx6Size/10_000        101.4 ± 0%    201.4 ± 0%   +98.62% (p=0.002 n=6)
RandomPfx6Size/100_000       71.69 ± 0%   160.80 ± 0%  +124.30% (p=0.002 n=6)
RandomPfx6Size/200_000       69.38 ± 0%   160.80 ± 0%  +131.77% (p=0.002 n=6)
RandomPfx6Size/500_000       72.35 ± 0%   156.50 ± 0%  +116.31% (p=0.002 n=6)
RandomPfx6Size/1_000_000     78.93 ± 0%   156.40 ± 0%   +98.15% (p=0.002 n=6)
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
                         │ bart/update.bm │           lite/update.bm            │
                         │   sec/route    │  sec/route    vs base               │
InsertRandomPfxs/1_000       193.7n ±  1%   163.7n ±  1%  -15.51% (p=0.002 n=6)
InsertRandomPfxs/10_000      146.8n ±  6%   114.0n ±  0%  -22.32% (p=0.002 n=6)
InsertRandomPfxs/100_000     217.8n ±  5%   162.4n ±  7%  -25.42% (p=0.002 n=6)
InsertRandomPfxs/200_000     270.7n ±  3%   189.2n ±  8%  -30.13% (p=0.002 n=6)
DeleteRandomPfxs/1_000       138.3n ±  2%   120.1n ±  4%  -13.20% (p=0.002 n=6)
DeleteRandomPfxs/10_000      86.83n ±  2%   75.97n ±  2%  -12.51% (p=0.002 n=6)
DeleteRandomPfxs/100_000     210.6n ±  5%   152.0n ± 11%  -27.83% (p=0.002 n=6)
DeleteRandomPfxs/200_000     261.8n ± 14%   198.2n ±  7%  -24.33% (p=0.002 n=6)
geomean                      180.2n         141.1n        -21.66%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │            fast/update.bm            │
                         │   sec/route    │   sec/route    vs base               │
InsertRandomPfxs/1_000       193.7n ±  1%    297.9n ±  3%  +53.79% (p=0.002 n=6)
InsertRandomPfxs/10_000      146.8n ±  6%    198.7n ±  3%  +35.37% (p=0.002 n=6)
InsertRandomPfxs/100_000     217.8n ±  5%    323.7n ± 11%  +48.63% (p=0.002 n=6)
InsertRandomPfxs/200_000     270.7n ±  3%    380.2n ±  5%  +40.47% (p=0.002 n=6)
DeleteRandomPfxs/1_000       138.3n ±  2%    246.0n ±  2%  +77.87% (p=0.002 n=6)
DeleteRandomPfxs/10_000      86.83n ±  2%   145.95n ±  2%  +68.08% (p=0.002 n=6)
DeleteRandomPfxs/100_000     210.6n ±  5%    338.9n ±  9%  +60.92% (p=0.002 n=6)
DeleteRandomPfxs/200_000     261.8n ± 14%    295.5n ± 24%  +12.83% (p=0.026 n=6)
geomean                      180.2n          267.5n        +48.46%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │          netipds/update.bm           │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000       193.7n ±  1%    212.1n ± 3%    +9.50% (p=0.002 n=6)
InsertRandomPfxs/10_000      146.8n ±  6%    239.8n ± 3%   +63.41% (p=0.002 n=6)
InsertRandomPfxs/100_000     217.8n ±  5%    333.6n ± 5%   +53.20% (p=0.002 n=6)
InsertRandomPfxs/200_000     270.7n ±  3%    435.1n ± 4%   +60.71% (p=0.002 n=6)
DeleteRandomPfxs/1_000       138.3n ±  2%    137.1n ± 1%         ~ (p=0.180 n=6)
DeleteRandomPfxs/10_000      86.83n ±  2%   184.10n ± 1%  +112.01% (p=0.002 n=6)
DeleteRandomPfxs/100_000     210.6n ±  5%    310.7n ± 3%   +47.53% (p=0.002 n=6)
DeleteRandomPfxs/200_000     261.8n ± 14%    413.2n ± 1%   +57.82% (p=0.002 n=6)
geomean                      180.2n          264.4n        +46.79%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │         critbitgo/update.bm          │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000       193.7n ±  1%    260.6n ± 1%   +34.54% (p=0.002 n=6)
InsertRandomPfxs/10_000      146.8n ±  6%    323.3n ± 1%  +120.34% (p=0.002 n=6)
InsertRandomPfxs/100_000     217.8n ±  5%    513.2n ± 5%  +135.71% (p=0.002 n=6)
InsertRandomPfxs/200_000     270.7n ±  3%    649.9n ± 1%  +140.08% (p=0.002 n=6)
DeleteRandomPfxs/1_000       138.3n ±  2%    130.1n ± 2%    -5.93% (p=0.002 n=6)
DeleteRandomPfxs/10_000      86.83n ±  2%   163.75n ± 2%   +88.58% (p=0.002 n=6)
DeleteRandomPfxs/100_000     210.6n ±  5%    357.7n ± 1%   +69.87% (p=0.002 n=6)
DeleteRandomPfxs/200_000     261.8n ± 14%    459.3n ± 2%   +75.42% (p=0.002 n=6)
geomean                      180.2n          315.6n        +75.18%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │          lpmtrie/update.bm           │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000       193.7n ±  1%    375.1n ± 1%   +93.68% (p=0.002 n=6)
InsertRandomPfxs/10_000      146.8n ±  6%    396.4n ± 2%  +170.15% (p=0.002 n=6)
InsertRandomPfxs/100_000     217.8n ±  5%    668.9n ± 1%  +207.19% (p=0.002 n=6)
InsertRandomPfxs/200_000     270.7n ±  3%    791.8n ± 5%  +192.48% (p=0.002 n=6)
DeleteRandomPfxs/1_000       138.3n ±  2%    123.5n ± 2%   -10.70% (p=0.002 n=6)
DeleteRandomPfxs/10_000      86.83n ±  2%   181.60n ± 2%  +109.13% (p=0.002 n=6)
DeleteRandomPfxs/100_000     210.6n ±  5%    512.0n ± 4%  +143.11% (p=0.002 n=6)
DeleteRandomPfxs/200_000     261.8n ± 14%    648.4n ± 7%  +147.60% (p=0.002 n=6)
geomean                      180.2n          394.5n       +118.97%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │       kentik-patricia/update.bm       │
                         │   sec/route    │   sec/route    vs base                │
InsertRandomPfxs/1_000       193.7n ±  1%    246.5n ±  2%   +27.26% (p=0.002 n=6)
InsertRandomPfxs/10_000      146.8n ±  6%    312.0n ±  1%  +112.61% (p=0.002 n=6)
InsertRandomPfxs/100_000     217.8n ±  5%    530.8n ± 20%  +143.77% (p=0.002 n=6)
InsertRandomPfxs/200_000     270.7n ±  3%    657.3n ±  1%  +142.81% (p=0.002 n=6)
DeleteRandomPfxs/1_000       138.3n ±  2%    262.2n ±  1%   +89.59% (p=0.002 n=6)
DeleteRandomPfxs/10_000      86.83n ±  2%   301.95n ±  2%  +247.73% (p=0.002 n=6)
DeleteRandomPfxs/100_000     210.6n ±  5%    668.4n ±  3%  +217.38% (p=0.002 n=6)
DeleteRandomPfxs/200_000     261.8n ± 14%    798.0n ±  2%  +204.74% (p=0.002 n=6)
geomean                      180.2n          428.3n        +137.76%
```

