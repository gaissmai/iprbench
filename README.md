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
LpmTier1Pfxs/RandomMatchIP4            20.89n ± 3%   20.41n ± 2%   -2.32% (p=0.005 n=10)
LpmTier1Pfxs/RandomMatchIP6            40.76n ± 2%   34.13n ± 3%  -16.27% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             21.09n ± 4%   19.42n ± 1%   -7.92% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             40.25n ± 5%   34.85n ± 2%  -13.43% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     19.47n ± 1%   18.80n ± 1%   -3.47% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     17.06n ± 3%   15.98n ± 1%   -6.30% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      19.44n ± 3%   18.65n ± 1%   -4.06% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      16.40n ± 1%   15.88n ± 1%   -3.20% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    20.45n ± 1%   19.77n ± 4%   -3.30% (p=0.008 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    21.21n ± 3%   20.83n ± 2%        ~ (p=0.072 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     20.43n ± 1%   20.23n ± 1%   -0.98% (p=0.010 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     20.59n ± 2%   20.24n ± 1%   -1.70% (p=0.001 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   16.05n ± 2%   16.17n ± 3%        ~ (p=0.739 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   32.70n ± 7%   32.50n ± 1%        ~ (p=0.644 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    27.32n ± 3%   25.04n ± 2%   -8.36% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    34.86n ± 4%   33.95n ± 3%        ~ (p=0.061 n=10)
geomean                                23.19n        22.07n        -4.83%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                     │ bart/lpm.bm │             fast/lpm.bm             │
                                     │   sec/op    │   sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4            20.89n ± 3%   18.41n ± 5%  -11.87% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            40.76n ± 2%   32.49n ± 3%  -20.29% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             21.09n ± 4%   17.04n ± 5%  -19.18% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             40.25n ± 5%   32.17n ± 6%  -20.07% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     19.47n ± 1%   15.56n ± 4%  -20.08% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     17.06n ± 3%   12.62n ± 1%  -26.03% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      19.44n ± 3%   15.34n ± 1%  -21.09% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      16.40n ± 1%   12.72n ± 1%  -22.43% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    20.45n ± 1%   16.23n ± 1%  -20.64% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    21.21n ± 3%   17.27n ± 2%  -18.56% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     20.43n ± 1%   16.88n ± 4%  -17.38% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     20.59n ± 2%   17.29n ± 2%  -16.05% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   16.05n ± 2%   13.50n ± 3%  -15.89% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   32.70n ± 7%   32.64n ± 4%        ~ (p=0.755 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    27.32n ± 3%   23.65n ± 5%  -13.43% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    34.86n ± 4%   33.37n ± 1%   -4.27% (p=0.005 n=10)
geomean                                23.19n        19.26n       -16.95%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                     │ bart/lpm.bm │            netipds/lpm.bm             │
                                     │   sec/op    │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4            20.89n ± 3%   137.55n ± 2%  +558.45% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            40.76n ± 2%   113.10n ± 1%  +177.48% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             21.09n ± 4%   156.25n ± 1%  +640.87% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             40.25n ± 5%   141.65n ± 5%  +251.93% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     19.47n ± 1%    49.61n ± 1%  +154.80% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     17.06n ± 3%    43.60n ± 3%  +155.57% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      19.44n ± 3%    50.13n ± 1%  +157.87% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      16.40n ± 1%    44.29n ± 1%  +170.01% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    20.45n ± 1%    72.01n ± 3%  +252.10% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    21.21n ± 3%    65.40n ± 4%  +208.44% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     20.43n ± 1%    74.36n ± 1%  +263.95% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     20.59n ± 2%    67.86n ± 3%  +229.60% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   16.05n ± 2%    85.55n ± 3%  +433.05% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   32.70n ± 7%   104.00n ± 2%  +218.04% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    27.32n ± 3%   114.50n ± 1%  +319.11% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    34.86n ± 4%   106.45n ± 1%  +205.36% (p=0.000 n=10)
geomean                                23.19n         82.10n       +254.08%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                     │ bart/lpm.bm │             critbitgo/lpm.bm              │
                                     │   sec/op    │     sec/op      vs base                   │
LpmTier1Pfxs/RandomMatchIP4            20.89n ± 3%    395.65n ±  3%   +1793.97% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            40.76n ± 2%    379.85n ±  3%    +831.92% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             21.09n ± 4%   2935.00n ±  2%  +13816.55% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             40.25n ± 5%   1270.50n ±  8%   +3056.52% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     19.47n ± 1%    118.85n ±  2%    +510.43% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     17.06n ± 3%    122.40n ±  1%    +617.47% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      19.44n ± 3%    218.30n ±  1%   +1022.94% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      16.40n ± 1%    176.65n ±  5%    +976.81% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    20.45n ± 1%    175.10n ±  1%    +756.23% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    21.21n ± 3%    164.95n ±  1%    +677.88% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     20.43n ± 1%    417.10n ±  4%   +1941.61% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     20.59n ± 2%    324.65n ±  3%   +1476.74% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   16.05n ± 2%    259.40n ±  1%   +1516.20% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   32.70n ± 7%    249.80n ±  4%    +663.91% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    27.32n ± 3%    924.20n ± 10%   +3282.87% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    34.86n ± 4%    576.95n ±  2%   +1555.05% (p=0.000 n=10)
geomean                                23.19n          346.1n         +1392.86%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                     │ bart/lpm.bm │             lpmtrie/lpm.bm              │
                                     │   sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            20.89n ± 3%   616.70n ±  8%  +2852.13% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            40.76n ± 2%   491.50n ±  9%  +1105.84% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             21.09n ± 4%   380.60n ± 12%  +1704.65% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             40.25n ± 5%   315.65n ±  2%   +684.22% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     19.47n ± 1%   115.05n ±  1%   +490.91% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     17.06n ± 3%   101.50n ±  1%   +494.96% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      19.44n ± 3%   110.45n ±  1%   +468.16% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      16.40n ± 1%    99.56n ±  1%   +506.92% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    20.45n ± 1%   176.45n ±  1%   +762.84% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    21.21n ± 3%   147.35n ±  1%   +594.88% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     20.43n ± 1%   174.80n ±  1%   +755.60% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     20.59n ± 2%   145.35n ±  1%   +605.93% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   16.05n ± 2%   269.30n ±  4%  +1577.88% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   32.70n ± 7%   216.70n ±  1%   +562.69% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    27.32n ± 3%   245.35n ±  0%   +798.06% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    34.86n ± 4%   207.90n ±  3%   +496.39% (p=0.000 n=10)
geomean                                23.19n         204.3n         +781.00%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                                     │ bart/lpm.bm │         kentik-patricia/lpm.bm         │
                                     │   sec/op    │    sec/op     vs base                  │
LpmTier1Pfxs/RandomMatchIP4            20.89n ± 3%   463.35n ± 4%  +2118.05% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            40.76n ± 2%   398.90n ± 1%   +878.66% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             21.09n ± 4%   316.70n ± 1%  +1401.66% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             40.25n ± 5%   282.15n ± 1%   +600.99% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     19.47n ± 1%   114.25n ± 2%   +486.80% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     17.06n ± 3%    96.78n ± 1%   +467.29% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      19.44n ± 3%   104.20n ± 0%   +436.01% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      16.40n ± 1%    85.03n ± 0%   +418.32% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    20.45n ± 1%   188.55n ± 4%   +822.00% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    21.21n ± 3%   153.60n ± 0%   +624.36% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     20.43n ± 1%   156.75n ± 1%   +667.25% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     20.59n ± 2%   140.20n ± 6%   +580.91% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   16.05n ± 2%   297.45n ± 3%  +1753.27% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   32.70n ± 7%   252.85n ± 4%   +673.24% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    27.32n ± 3%   247.30n ± 3%   +805.20% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    34.86n ± 4%   229.30n ± 5%   +557.77% (p=0.000 n=10)
geomean                                23.19n         195.1n        +741.24%
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
Tier1PfxSize/100_000         56.72 ± 0%   163.80 ± 0%  +188.79% (p=0.002 n=6)
Tier1PfxSize/200_000         49.53 ± 0%   164.00 ± 0%  +231.11% (p=0.002 n=6)
Tier1PfxSize/500_000         42.90 ± 0%   144.20 ± 0%  +236.13% (p=0.002 n=6)
Tier1PfxSize/1_000_000       39.74 ± 0%   144.50 ± 0%  +263.61% (p=0.002 n=6)
RandomPfx4Size/1_000         82.62 ± 2%   140.90 ± 1%   +70.54% (p=0.002 n=6)
RandomPfx4Size/10_000        57.37 ± 0%   109.60 ± 0%   +91.04% (p=0.002 n=6)
RandomPfx4Size/100_000       72.09 ± 0%   139.90 ± 0%   +94.06% (p=0.002 n=6)
RandomPfx4Size/200_000       65.15 ± 0%   139.80 ± 0%  +114.58% (p=0.002 n=6)
RandomPfx4Size/500_000       52.44 ± 0%   139.80 ± 0%  +166.59% (p=0.002 n=6)
RandomPfx4Size/1_000_000     46.46 ± 0%   139.60 ± 0%  +200.47% (p=0.002 n=6)
RandomPfx6Size/1_000         83.74 ± 2%   157.30 ± 1%   +87.84% (p=0.002 n=6)
RandomPfx6Size/10_000        101.4 ± 0%    201.4 ± 0%   +98.62% (p=0.002 n=6)
RandomPfx6Size/100_000       71.69 ± 0%   160.80 ± 0%  +124.30% (p=0.002 n=6)
RandomPfx6Size/200_000       69.38 ± 0%   160.80 ± 0%  +131.77% (p=0.002 n=6)
RandomPfx6Size/500_000       72.35 ± 0%   156.50 ± 0%  +116.31% (p=0.002 n=6)
RandomPfx6Size/1_000_000     78.93 ± 0%   156.30 ± 0%   +98.02% (p=0.002 n=6)
RandomPfxSize/1_000          100.7 ± 2%    144.7 ± 1%   +43.69% (p=0.002 n=6)
RandomPfxSize/10_000         74.11 ± 0%   140.20 ± 0%   +89.18% (p=0.002 n=6)
RandomPfxSize/100_000        84.34 ± 0%   180.00 ± 0%  +113.42% (p=0.002 n=6)
RandomPfxSize/200_000        75.60 ± 0%   180.00 ± 0%  +138.10% (p=0.002 n=6)
RandomPfxSize/500_000        59.59 ± 0%   144.00 ± 0%  +141.65% (p=0.002 n=6)
RandomPfxSize/1_000_000      52.03 ± 0%   144.00 ± 0%  +176.76% (p=0.002 n=6)
geomean                      67.56         152.8       +126.18%
```

## update, insert/delete

`bart.Lite` is the fastest algorithm for updates.

```
goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │           lite/update.bm            │
                         │   sec/route    │  sec/route    vs base               │
InsertRandomPfxs/1_000       202.5n ±  4%   167.8n ±  2%  -17.11% (p=0.002 n=6)
InsertRandomPfxs/10_000      160.2n ±  1%   117.9n ±  2%  -26.43% (p=0.002 n=6)
InsertRandomPfxs/100_000     243.2n ±  9%   161.7n ±  9%  -33.53% (p=0.002 n=6)
InsertRandomPfxs/200_000     264.5n ± 21%   187.5n ±  5%  -29.09% (p=0.002 n=6)
DeleteRandomPfxs/1_000       142.4n ±  4%   125.6n ±  4%  -11.83% (p=0.002 n=6)
DeleteRandomPfxs/10_000      87.01n ±  5%   75.39n ±  4%  -13.37% (p=0.002 n=6)
DeleteRandomPfxs/100_000     218.9n ±  7%   170.1n ±  8%  -22.31% (p=0.002 n=6)
DeleteRandomPfxs/200_000     245.2n ± 23%   203.6n ± 11%  -16.99% (p=0.002 n=6)
geomean                      185.3n         145.1n        -21.68%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │            fast/update.bm             │
                         │   sec/route    │   sec/route     vs base               │
InsertRandomPfxs/1_000       202.5n ±  4%    309.8n ±   3%  +52.99% (p=0.002 n=6)
InsertRandomPfxs/10_000      160.2n ±  1%    219.5n ±   1%  +36.97% (p=0.002 n=6)
InsertRandomPfxs/100_000     243.2n ±  9%    346.1n ±   8%  +42.31% (p=0.002 n=6)
InsertRandomPfxs/200_000     264.5n ± 21%    392.7n ±   7%  +48.45% (p=0.002 n=6)
DeleteRandomPfxs/1_000       142.4n ±  4%    253.3n ±   4%  +77.85% (p=0.002 n=6)
DeleteRandomPfxs/10_000      87.01n ±  5%   139.25n ±   4%  +60.03% (p=0.002 n=6)
DeleteRandomPfxs/100_000     218.9n ±  7%    350.8n ±  19%  +60.22% (p=0.002 n=6)
DeleteRandomPfxs/200_000     245.2n ± 23%    322.2n ± 361%  +31.40% (p=0.002 n=6)
geomean                      185.3n          279.1n         +50.65%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │          netipds/update.bm          │
                         │   sec/route    │  sec/route    vs base               │
InsertRandomPfxs/1_000       202.5n ±  4%    209.1n ± 1%   +3.26% (p=0.026 n=6)
InsertRandomPfxs/10_000      160.2n ±  1%    237.4n ± 3%  +48.14% (p=0.002 n=6)
InsertRandomPfxs/100_000     243.2n ±  9%    316.2n ± 2%  +30.04% (p=0.002 n=6)
InsertRandomPfxs/200_000     264.5n ± 21%    415.5n ± 4%  +57.09% (p=0.002 n=6)
DeleteRandomPfxs/1_000       142.4n ±  4%    123.8n ± 4%  -13.13% (p=0.002 n=6)
DeleteRandomPfxs/10_000      87.01n ±  5%   165.15n ± 3%  +89.79% (p=0.002 n=6)
DeleteRandomPfxs/100_000     218.9n ±  7%    293.8n ± 3%  +34.19% (p=0.002 n=6)
DeleteRandomPfxs/200_000     245.2n ± 23%    441.2n ± 5%  +79.91% (p=0.002 n=6)
geomean                      185.3n          253.9n       +37.04%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │          critbitgo/update.bm           │
                         │   sec/route    │   sec/route     vs base                │
InsertRandomPfxs/1_000       202.5n ±  4%    259.3n ±   2%   +28.05% (p=0.002 n=6)
InsertRandomPfxs/10_000      160.2n ±  1%    333.3n ±   3%  +108.02% (p=0.002 n=6)
InsertRandomPfxs/100_000     243.2n ±  9%    524.8n ±   5%  +115.81% (p=0.002 n=6)
InsertRandomPfxs/200_000     264.5n ± 21%    720.4n ±  10%  +172.38% (p=0.002 n=6)
DeleteRandomPfxs/1_000       142.4n ±  4%    132.4n ±   2%    -7.06% (p=0.002 n=6)
DeleteRandomPfxs/10_000      87.01n ±  5%   175.50n ±   4%  +101.69% (p=0.002 n=6)
DeleteRandomPfxs/100_000     218.9n ±  7%    363.3n ± 240%   +65.93% (p=0.002 n=6)
DeleteRandomPfxs/200_000     245.2n ± 23%    447.6n ±   2%   +82.57% (p=0.002 n=6)
geomean                      185.3n          324.7n          +75.24%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │          lpmtrie/update.bm           │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000       202.5n ±  4%    378.5n ± 1%   +86.89% (p=0.002 n=6)
InsertRandomPfxs/10_000      160.2n ±  1%    401.4n ± 2%  +150.51% (p=0.002 n=6)
InsertRandomPfxs/100_000     243.2n ±  9%    662.1n ± 3%  +172.25% (p=0.002 n=6)
InsertRandomPfxs/200_000     264.5n ± 21%    788.6n ± 4%  +198.15% (p=0.002 n=6)
DeleteRandomPfxs/1_000       142.4n ±  4%    126.4n ± 4%   -11.30% (p=0.002 n=6)
DeleteRandomPfxs/10_000      87.01n ±  5%   185.60n ± 2%  +113.30% (p=0.002 n=6)
DeleteRandomPfxs/100_000     218.9n ±  7%    495.5n ± 2%  +126.31% (p=0.002 n=6)
DeleteRandomPfxs/200_000     245.2n ± 23%    641.0n ± 3%  +161.44% (p=0.002 n=6)
geomean                      185.3n          394.9n       +113.11%

goos: linux
goarch: amd64
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
                         │ bart/update.bm │        kentik-patricia/update.bm        │
                         │   sec/route    │   sec/route     vs base                 │
InsertRandomPfxs/1_000       202.5n ±  4%    256.6n ±   2%    +26.72% (p=0.002 n=6)
InsertRandomPfxs/10_000      160.2n ±  1%    325.8n ±   1%   +103.31% (p=0.002 n=6)
InsertRandomPfxs/100_000     243.2n ±  9%    541.5n ±   2%   +122.68% (p=0.002 n=6)
InsertRandomPfxs/200_000     264.5n ± 21%   4380.5n ±  84%  +1556.14% (p=0.002 n=6)
DeleteRandomPfxs/1_000       142.4n ±  4%    259.8n ±  15%    +82.34% (p=0.002 n=6)
DeleteRandomPfxs/10_000      87.01n ±  5%   302.35n ± 846%   +247.47% (p=0.002 n=6)
DeleteRandomPfxs/100_000     218.9n ±  7%    655.4n ± 589%   +199.34% (p=0.002 n=6)
DeleteRandomPfxs/200_000     245.2n ± 23%    785.8n ±   2%   +220.45% (p=0.002 n=6)
geomean                      185.3n          547.1n          +195.25%
```

