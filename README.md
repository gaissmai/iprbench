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
goos: darwin
goarch: arm64
cpu: Apple M4 Max
                                     │ fast/lpm.bm  │              lite/lpm.bm               │
                                     │    sec/op    │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            4.059n ± 10%    6.593n ± 10%   +62.44% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            5.455n ± 27%   12.745n ± 15%  +133.64% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             4.734n ±  0%    7.494n ±  0%   +58.27% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             6.156n ± 20%   11.455n ± 56%   +86.08% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     4.839n ±  3%    7.169n ±  1%   +48.16% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     5.198n ±  4%    7.787n ±  0%   +49.81% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      4.706n ± 16%    7.148n ± 34%   +51.89% (p=0.005 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      5.198n ±  3%    7.773n ±  0%   +49.54% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    5.322n ±  0%    7.751n ±  0%   +45.66% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    5.202n ±  4%    7.785n ±  0%   +49.68% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     5.319n ±  1%    7.729n ±  0%   +45.31% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     5.199n ±  4%    7.791n ±  0%   +49.87% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   3.329n ± 39%    4.942n ± 45%   +48.49% (p=0.006 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   5.984n ± 38%   11.135n ± 46%   +86.10% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    5.319n ± 10%    7.755n ±  4%   +45.80% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    5.986n ±  4%   11.130n ±  0%   +85.95% (p=0.000 n=20)
geomean                                5.073n          8.160n         +60.83%

goos: darwin
goarch: arm64
cpu: Apple M4 Max
                                     │ fast/lpm.bm  │              bart/lpm.bm               │
                                     │    sec/op    │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            4.059n ± 10%    6.873n ±  7%   +69.36% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            5.455n ± 27%   13.800n ± 10%  +152.98% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             4.734n ±  0%    7.447n ±  0%   +57.30% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             6.156n ± 20%   11.930n ± 52%   +93.79% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     4.839n ±  3%    7.519n ±  0%   +55.37% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     5.198n ±  4%    8.006n ±  1%   +54.02% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      4.706n ± 16%    7.490n ± 37%   +59.15% (p=0.002 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      5.198n ±  3%    7.988n ±  0%   +53.66% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    5.322n ±  0%    7.854n ±  0%   +47.58% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    5.202n ±  4%    7.976n ±  1%   +53.34% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     5.319n ±  1%    7.840n ±  0%   +47.41% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     5.199n ±  4%    7.982n ±  1%   +53.54% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   3.329n ± 39%    5.095n ± 43%   +53.07% (p=0.006 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   5.984n ± 38%   11.675n ± 45%   +95.12% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    5.319n ± 10%    7.854n ±  4%   +47.65% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    5.986n ±  4%   11.675n ±  1%   +95.05% (p=0.000 n=20)
geomean                                5.073n          8.429n         +66.13%

goos: darwin
goarch: arm64
cpu: Apple M4 Max
                                     │ fast/lpm.bm  │             netipds/lpm.bm             │
                                     │    sec/op    │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            4.059n ± 10%   22.845n ± 20%  +462.89% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            5.455n ± 27%   22.200n ± 18%  +306.97% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             4.734n ±  0%   24.340n ±  8%  +414.10% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             6.156n ± 20%   24.475n ± 11%  +297.58% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     4.839n ±  3%   12.795n ±  8%  +164.41% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     5.198n ±  4%   12.845n ±  6%  +147.11% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      4.706n ± 16%   13.340n ±  5%  +183.47% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      5.198n ±  3%   13.205n ±  7%  +154.04% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    5.322n ±  0%   15.770n ±  5%  +196.35% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    5.202n ±  4%   15.410n ±  8%  +196.26% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     5.319n ±  1%   16.970n ±  3%  +219.04% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     5.199n ±  4%   16.545n ±  5%  +218.26% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   3.329n ± 39%   16.610n ± 13%  +399.02% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   5.984n ± 38%   18.750n ±  3%  +213.36% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    5.319n ± 10%   19.620n ±  3%  +268.87% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    5.986n ±  4%   19.565n ±  3%  +226.87% (p=0.000 n=20)
geomean                                5.073n          17.41n        +243.19%

goos: darwin
goarch: arm64
cpu: Apple M4 Max
                                     │ fast/lpm.bm  │             critbitgo/lpm.bm             │
                                     │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            4.059n ± 10%    58.835n ±  6%  +1349.67% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            5.455n ± 27%    83.550n ±  3%  +1431.62% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             4.734n ±  0%   224.200n ± 15%  +4635.45% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             6.156n ± 20%   198.950n ± 10%  +3131.81% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     4.839n ±  3%    33.810n ±  5%   +598.70% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     5.198n ±  4%    40.410n ±  3%   +677.41% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      4.706n ± 16%    65.610n ± 11%  +1294.18% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      5.198n ±  3%    62.035n ±  8%  +1093.44% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    5.322n ±  0%    39.955n ±  7%   +650.82% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    5.202n ±  4%    47.065n ±  4%   +804.84% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     5.319n ±  1%   106.800n ± 12%  +1907.90% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     5.199n ±  4%    90.425n ± 16%  +1639.44% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   3.329n ± 39%    47.725n ±  5%  +1333.83% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   5.984n ± 38%    52.450n ±  2%   +776.58% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    5.319n ± 10%   172.600n ± 10%  +3144.97% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    5.986n ±  4%   136.150n ± 12%  +2174.66% (p=0.000 n=20)
geomean                                5.073n           76.24n        +1402.70%

goos: darwin
goarch: arm64
cpu: Apple M4 Max
                                     │ fast/lpm.bm  │              lpmtrie/lpm.bm              │
                                     │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            4.059n ± 10%   119.600n ±  5%  +2846.90% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            5.455n ± 27%   136.050n ± 18%  +2394.04% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             4.734n ±  0%   104.400n ±  5%  +2105.09% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             6.156n ± 20%    92.880n ± 16%  +1408.77% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     4.839n ±  3%    49.905n ± 10%   +931.31% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     5.198n ±  4%    44.675n ± 20%   +759.47% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      4.706n ± 16%    43.400n ±  7%   +822.23% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      5.198n ±  3%    36.120n ± 17%   +594.88% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    5.322n ±  0%    65.850n ±  8%  +1137.43% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    5.202n ±  4%    58.400n ±  6%  +1022.75% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     5.319n ±  1%    64.805n ±  4%  +1118.37% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     5.199n ±  4%    54.865n ±  7%   +955.40% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   3.329n ± 39%    84.805n ±  4%  +2447.84% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   5.984n ± 38%    74.665n ±  3%  +1147.85% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    5.319n ± 10%    77.220n ±  6%  +1351.78% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    5.986n ±  4%    73.405n ±  6%  +1126.38% (p=0.000 n=20)
geomean                                5.073n           69.11n        +1262.14%

goos: darwin
goarch: arm64
cpu: Apple M4 Max
                                     │ fast/lpm.bm  │          kentik-patricia/lpm.bm          │
                                     │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            4.059n ± 10%    80.410n ±  7%  +1881.27% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            5.455n ± 27%   105.450n ± 19%  +1833.09% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             4.734n ±  0%    64.920n ±  5%  +1271.21% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             6.156n ± 20%    71.680n ± 14%  +1064.39% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     4.839n ±  3%    37.615n ±  8%   +677.33% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     5.198n ±  4%    40.280n ± 18%   +674.91% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      4.706n ± 16%    30.455n ±  7%   +547.15% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      5.198n ±  3%    32.015n ± 12%   +515.91% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    5.322n ±  0%    47.595n ±  6%   +794.39% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    5.202n ±  4%    50.075n ±  5%   +862.70% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     5.319n ±  1%    42.670n ±  3%   +702.22% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     5.199n ±  4%    45.130n ±  5%   +768.14% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   3.329n ± 39%    59.985n ±  8%  +1702.16% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   5.984n ± 38%    61.640n ±  3%   +930.17% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    5.319n ± 10%    49.870n ±  5%   +837.58% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    5.986n ±  4%    58.110n ±  5%   +870.85% (p=0.000 n=20)
geomean                                5.073n           52.01n         +925.14%
```

## size of the routing tables


`bart.Lite` has the lowest memory consumption under all competitors.


```
goos: darwin
goarch: arm64
cpu: Apple M4 Max
                         │ lite/size.bm │            bart/size.bm            │
                         │ bytes/route  │ bytes/route  vs base               │
Tier1PfxSize/1_000           85.36 ± 2%   105.30 ± 2%  +23.36% (p=0.002 n=6)
Tier1PfxSize/10_000          63.94 ± 0%    83.96 ± 0%  +31.31% (p=0.002 n=6)
Tier1PfxSize/100_000         36.97 ± 0%    56.72 ± 0%  +53.42% (p=0.002 n=6)
Tier1PfxSize/200_000         30.07 ± 0%    49.53 ± 0%  +64.72% (p=0.002 n=6)
Tier1PfxSize/500_000         23.63 ± 0%    42.90 ± 0%  +81.55% (p=0.002 n=6)
Tier1PfxSize/1_000_000       20.42 ± 0%    39.74 ± 0%  +94.61% (p=0.002 n=6)
RandomPfx4Size/1_000         61.90 ± 3%    82.62 ± 2%  +33.47% (p=0.002 n=6)
RandomPfx4Size/10_000        38.30 ± 0%    57.37 ± 0%  +49.79% (p=0.002 n=6)
RandomPfx4Size/100_000       49.81 ± 0%    72.09 ± 0%  +44.73% (p=0.002 n=6)
RandomPfx4Size/200_000       43.92 ± 0%    65.15 ± 0%  +48.34% (p=0.002 n=6)
RandomPfx4Size/500_000       31.18 ± 0%    52.44 ± 0%  +68.18% (p=0.002 n=6)
RandomPfx4Size/1_000_000     25.84 ± 0%    46.46 ± 0%  +79.80% (p=0.002 n=6)
RandomPfx6Size/1_000         66.27 ± 3%    83.74 ± 2%  +26.36% (p=0.002 n=6)
RandomPfx6Size/10_000        80.81 ± 0%   101.40 ± 0%  +25.48% (p=0.002 n=6)
RandomPfx6Size/100_000       53.59 ± 0%    71.69 ± 0%  +33.77% (p=0.002 n=6)
RandomPfx6Size/200_000       51.40 ± 0%    69.38 ± 0%  +34.98% (p=0.002 n=6)
RandomPfx6Size/500_000       53.93 ± 0%    72.35 ± 0%  +34.16% (p=0.002 n=6)
RandomPfx6Size/1_000_000     59.52 ± 0%    78.93 ± 0%  +32.61% (p=0.002 n=6)
RandomPfxSize/1_000          80.70 ± 2%   100.70 ± 2%  +24.78% (p=0.002 n=6)
RandomPfxSize/10_000         56.53 ± 0%    74.11 ± 0%  +31.10% (p=0.002 n=6)
RandomPfxSize/100_000        63.46 ± 0%    84.34 ± 0%  +32.90% (p=0.002 n=6)
RandomPfxSize/200_000        54.19 ± 0%    75.60 ± 0%  +39.51% (p=0.002 n=6)
RandomPfxSize/500_000        39.18 ± 0%    59.59 ± 0%  +52.09% (p=0.002 n=6)
RandomPfxSize/1_000_000      31.89 ± 0%    52.03 ± 0%  +63.15% (p=0.002 n=6)
geomean                      46.66         67.56       +44.78%

goos: darwin
goarch: arm64
cpu: Apple M4 Max
                         │ lite/size.bm │            fast/size.bm             │
                         │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000           85.36 ± 2%   172.20 ± 1%  +101.73% (p=0.002 n=6)
Tier1PfxSize/10_000          63.94 ± 0%   151.10 ± 0%  +136.32% (p=0.002 n=6)
Tier1PfxSize/100_000         36.97 ± 0%   101.80 ± 0%  +175.36% (p=0.002 n=6)
Tier1PfxSize/200_000         30.07 ± 0%    81.43 ± 0%  +170.80% (p=0.002 n=6)
Tier1PfxSize/500_000         23.63 ± 0%    62.27 ± 0%  +163.52% (p=0.002 n=6)
Tier1PfxSize/1_000_000       20.42 ± 0%    52.71 ± 0%  +158.13% (p=0.002 n=6)
RandomPfx4Size/1_000         61.90 ± 3%   147.10 ± 1%  +137.64% (p=0.002 n=6)
RandomPfx4Size/10_000        38.30 ± 0%    73.09 ± 0%   +90.84% (p=0.002 n=6)
RandomPfx4Size/100_000       49.81 ± 0%   132.60 ± 0%  +166.21% (p=0.002 n=6)
RandomPfx4Size/200_000       43.92 ± 0%   130.30 ± 0%  +196.68% (p=0.002 n=6)
RandomPfx4Size/500_000       31.18 ± 0%    88.48 ± 0%  +183.77% (p=0.002 n=6)
RandomPfx4Size/1_000_000     25.84 ± 0%    65.46 ± 0%  +153.33% (p=0.002 n=6)
RandomPfx6Size/1_000         66.27 ± 3%   108.80 ± 2%   +64.18% (p=0.002 n=6)
RandomPfx6Size/10_000        80.81 ± 0%   179.60 ± 0%  +122.25% (p=0.002 n=6)
RandomPfx6Size/100_000       53.59 ± 0%    98.32 ± 0%   +83.47% (p=0.002 n=6)
RandomPfx6Size/200_000       51.40 ± 0%    88.85 ± 0%   +72.86% (p=0.002 n=6)
RandomPfx6Size/500_000       53.93 ± 0%    96.34 ± 0%   +78.64% (p=0.002 n=6)
RandomPfx6Size/1_000_000     59.52 ± 0%   116.70 ± 0%   +96.07% (p=0.002 n=6)
RandomPfxSize/1_000          80.70 ± 2%   168.10 ± 1%  +108.30% (p=0.002 n=6)
RandomPfxSize/10_000         56.53 ± 0%    97.72 ± 0%   +72.86% (p=0.002 n=6)
RandomPfxSize/100_000        63.46 ± 0%   156.60 ± 0%  +146.77% (p=0.002 n=6)
RandomPfxSize/200_000        54.19 ± 0%   146.30 ± 0%  +169.98% (p=0.002 n=6)
RandomPfxSize/500_000        39.18 ± 0%   100.30 ± 0%  +156.00% (p=0.002 n=6)
RandomPfxSize/1_000_000      31.89 ± 0%    74.51 ± 0%  +133.65% (p=0.002 n=6)
geomean                      46.66         106.1       +127.31%

goos: darwin
goarch: arm64
cpu: Apple M4 Max
                         │ lite/size.bm │           netipds/size.bm           │
                         │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000           85.36 ± 2%    74.14 ± 3%   -13.14% (p=0.002 n=6)
Tier1PfxSize/10_000          63.94 ± 0%    68.94 ± 0%    +7.82% (p=0.002 n=6)
Tier1PfxSize/100_000         36.97 ± 0%    67.66 ± 0%   +83.01% (p=0.002 n=6)
Tier1PfxSize/200_000         30.07 ± 0%    66.91 ± 0%  +122.51% (p=0.002 n=6)
Tier1PfxSize/500_000         23.63 ± 0%    65.35 ± 0%  +176.56% (p=0.002 n=6)
Tier1PfxSize/1_000_000       20.42 ± 0%    63.72 ± 0%  +212.05% (p=0.002 n=6)
RandomPfx4Size/1_000         61.90 ± 3%    61.89 ± 3%    -0.02% (p=0.013 n=6)
RandomPfx4Size/10_000        38.30 ± 0%    51.37 ± 0%   +34.13% (p=0.002 n=6)
RandomPfx4Size/100_000       49.81 ± 0%    48.25 ± 0%    -3.13% (p=0.002 n=6)
RandomPfx4Size/200_000       43.92 ± 0%    47.57 ± 0%    +8.31% (p=0.002 n=6)
RandomPfx4Size/500_000       31.18 ± 0%    46.51 ± 0%   +49.17% (p=0.002 n=6)
RandomPfx4Size/1_000_000     25.84 ± 0%    45.66 ± 0%   +76.70% (p=0.002 n=6)
RandomPfx6Size/1_000         66.27 ± 3%   100.20 ± 2%   +51.20% (p=0.002 n=6)
RandomPfx6Size/10_000        80.81 ± 0%    94.28 ± 0%   +16.67% (p=0.002 n=6)
RandomPfx6Size/100_000       53.59 ± 0%    87.64 ± 0%   +63.54% (p=0.002 n=6)
RandomPfx6Size/200_000       51.40 ± 0%    85.90 ± 0%   +67.12% (p=0.002 n=6)
RandomPfx6Size/500_000       53.93 ± 0%    84.11 ± 0%   +55.96% (p=0.002 n=6)
RandomPfx6Size/1_000_000     59.52 ± 0%    83.23 ± 0%   +39.84% (p=0.002 n=6)
RandomPfxSize/1_000          80.70 ± 2%    74.90 ± 3%    -7.19% (p=0.002 n=6)
RandomPfxSize/10_000         56.53 ± 0%    69.40 ± 0%   +22.77% (p=0.002 n=6)
RandomPfxSize/100_000        63.46 ± 0%    64.29 ± 0%    +1.31% (p=0.002 n=6)
RandomPfxSize/200_000        54.19 ± 0%    61.55 ± 0%   +13.58% (p=0.002 n=6)
RandomPfxSize/500_000        39.18 ± 0%    57.42 ± 0%   +46.55% (p=0.002 n=6)
RandomPfxSize/1_000_000      31.89 ± 0%    53.93 ± 0%   +69.11% (p=0.002 n=6)
geomean                      46.66         66.04        +41.53%

goos: darwin
goarch: arm64
cpu: Apple M4 Max
                         │ lite/size.bm │          critbitgo/size.bm          │
                         │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000           85.36 ± 2%   119.70 ± 2%   +40.23% (p=0.002 n=6)
Tier1PfxSize/10_000          63.94 ± 0%   114.80 ± 0%   +79.54% (p=0.002 n=6)
Tier1PfxSize/100_000         36.97 ± 0%   114.40 ± 0%  +209.44% (p=0.002 n=6)
Tier1PfxSize/200_000         30.07 ± 0%   114.40 ± 0%  +280.45% (p=0.002 n=6)
Tier1PfxSize/500_000         23.63 ± 0%   114.40 ± 0%  +384.13% (p=0.002 n=6)
Tier1PfxSize/1_000_000       20.42 ± 0%   114.40 ± 0%  +460.24% (p=0.002 n=6)
RandomPfx4Size/1_000         61.90 ± 3%   116.40 ± 2%   +88.05% (p=0.002 n=6)
RandomPfx4Size/10_000        38.30 ± 0%   112.40 ± 0%  +193.47% (p=0.002 n=6)
RandomPfx4Size/100_000       49.81 ± 0%   112.00 ± 0%  +124.85% (p=0.002 n=6)
RandomPfx4Size/200_000       43.92 ± 0%   112.00 ± 0%  +155.01% (p=0.002 n=6)
RandomPfx4Size/500_000       31.18 ± 0%   112.00 ± 0%  +259.20% (p=0.002 n=6)
RandomPfx4Size/1_000_000     25.84 ± 0%   112.00 ± 0%  +333.44% (p=0.002 n=6)
RandomPfx6Size/1_000         66.27 ± 3%   132.40 ± 2%   +99.79% (p=0.002 n=6)
RandomPfx6Size/10_000        80.81 ± 0%   128.40 ± 0%   +58.89% (p=0.002 n=6)
RandomPfx6Size/100_000       53.59 ± 0%   128.00 ± 0%  +138.85% (p=0.002 n=6)
RandomPfx6Size/200_000       51.40 ± 0%   128.00 ± 0%  +149.03% (p=0.002 n=6)
RandomPfx6Size/500_000       53.93 ± 0%   128.00 ± 0%  +137.34% (p=0.002 n=6)
RandomPfx6Size/1_000_000     59.52 ± 0%   128.00 ± 0%  +115.05% (p=0.002 n=6)
RandomPfxSize/1_000          80.70 ± 2%   119.80 ± 2%   +48.45% (p=0.002 n=6)
RandomPfxSize/10_000         56.53 ± 0%   115.60 ± 0%  +104.49% (p=0.002 n=6)
RandomPfxSize/100_000        63.46 ± 0%   115.30 ± 0%   +81.69% (p=0.002 n=6)
RandomPfxSize/200_000        54.19 ± 0%   115.20 ± 0%  +112.59% (p=0.002 n=6)
RandomPfxSize/500_000        39.18 ± 0%   115.20 ± 0%  +194.03% (p=0.002 n=6)
RandomPfxSize/1_000_000      31.89 ± 0%   115.20 ± 0%  +261.24% (p=0.002 n=6)
geomean                      46.66         118.1       +153.05%

goos: darwin
goarch: arm64
cpu: Apple M4 Max
                         │ lite/size.bm │           lpmtrie/size.bm            │
                         │ bytes/route  │ bytes/route   vs base                │
Tier1PfxSize/1_000           85.36 ± 2%   215.60 ±  5%  +152.58% (p=0.002 n=6)
Tier1PfxSize/10_000          63.94 ± 0%   210.50 ±  5%  +229.21% (p=0.002 n=6)
Tier1PfxSize/100_000         36.97 ± 0%   209.90 ±  5%  +467.76% (p=0.002 n=6)
Tier1PfxSize/200_000         30.07 ± 0%   209.20 ±  5%  +595.71% (p=0.002 n=6)
Tier1PfxSize/500_000         23.63 ± 0%   207.90 ±  7%  +779.81% (p=0.002 n=6)
Tier1PfxSize/1_000_000       20.42 ± 0%   205.00 ±  7%  +903.92% (p=0.002 n=6)
RandomPfx4Size/1_000         61.90 ± 3%   205.50 ±  8%  +231.99% (p=0.002 n=6)
RandomPfx4Size/10_000        38.30 ± 0%   186.70 ±  9%  +387.47% (p=0.002 n=6)
RandomPfx4Size/100_000       49.81 ± 0%   179.60 ±  9%  +260.57% (p=0.002 n=6)
RandomPfx4Size/200_000       43.92 ± 0%   178.50 ± 10%  +306.42% (p=0.002 n=6)
RandomPfx4Size/500_000       31.18 ± 0%   176.80 ± 10%  +467.03% (p=0.002 n=6)
RandomPfx4Size/1_000_000     25.84 ± 0%   175.50 ± 10%  +579.18% (p=0.002 n=6)
RandomPfx6Size/1_000         66.27 ± 3%   228.00 ±  8%  +244.05% (p=0.002 n=6)
RandomPfx6Size/10_000        80.81 ± 0%   222.50 ±  8%  +175.34% (p=0.002 n=6)
RandomPfx6Size/100_000       53.59 ± 0%   213.90 ±  9%  +299.14% (p=0.002 n=6)
RandomPfx6Size/200_000       51.40 ± 0%   210.50 ±  9%  +309.53% (p=0.002 n=6)
RandomPfx6Size/500_000       53.93 ± 0%   206.70 ±  9%  +283.27% (p=0.002 n=6)
RandomPfx6Size/1_000_000     59.52 ± 0%   204.90 ±  9%  +244.25% (p=0.002 n=6)
RandomPfxSize/1_000          80.70 ± 2%   215.30 ±  5%  +166.79% (p=0.002 n=6)
RandomPfxSize/10_000         56.53 ± 0%   210.30 ±  5%  +272.01% (p=0.002 n=6)
RandomPfxSize/100_000        63.46 ± 0%   203.50 ±  7%  +220.67% (p=0.002 n=6)
RandomPfxSize/200_000        54.19 ± 0%   198.70 ±  8%  +266.67% (p=0.002 n=6)
RandomPfxSize/500_000        39.18 ± 0%   189.90 ±  9%  +384.69% (p=0.002 n=6)
RandomPfxSize/1_000_000      31.89 ± 0%   181.40 ± 10%  +468.83% (p=0.002 n=6)
geomean                      46.66         201.4        +331.56%

goos: darwin
goarch: arm64
cpu: Apple M4 Max
                         │ lite/size.bm │       kentik-patricia/size.bm       │
                         │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000           85.36 ± 2%   145.50 ± 1%   +70.45% (p=0.002 n=6)
Tier1PfxSize/10_000          63.94 ± 0%   200.30 ± 0%  +213.26% (p=0.002 n=6)
Tier1PfxSize/100_000         36.97 ± 0%   164.30 ± 0%  +344.41% (p=0.002 n=6)
Tier1PfxSize/200_000         30.07 ± 0%   164.20 ± 0%  +446.06% (p=0.002 n=6)
Tier1PfxSize/500_000         23.63 ± 0%   144.20 ± 0%  +510.24% (p=0.002 n=6)
Tier1PfxSize/1_000_000       20.42 ± 0%   144.20 ± 0%  +606.17% (p=0.002 n=6)
RandomPfx4Size/1_000         61.90 ± 3%   140.90 ± 1%  +127.63% (p=0.002 n=6)
RandomPfx4Size/10_000        38.30 ± 0%   109.60 ± 0%  +186.16% (p=0.002 n=6)
RandomPfx4Size/100_000       49.81 ± 0%   139.90 ± 0%  +180.87% (p=0.002 n=6)
RandomPfx4Size/200_000       43.92 ± 0%   139.80 ± 0%  +218.31% (p=0.002 n=6)
RandomPfx4Size/500_000       31.18 ± 0%   139.80 ± 0%  +348.36% (p=0.002 n=6)
RandomPfx4Size/1_000_000     25.84 ± 0%   139.70 ± 0%  +440.63% (p=0.002 n=6)
RandomPfx6Size/1_000         66.27 ± 3%   157.30 ± 1%  +137.36% (p=0.002 n=6)
RandomPfx6Size/10_000        80.81 ± 0%   201.40 ± 0%  +149.23% (p=0.002 n=6)
RandomPfx6Size/100_000       53.59 ± 0%   160.80 ± 0%  +200.06% (p=0.002 n=6)
RandomPfx6Size/200_000       51.40 ± 0%   160.80 ± 0%  +212.84% (p=0.002 n=6)
RandomPfx6Size/500_000       53.93 ± 0%   156.50 ± 0%  +190.19% (p=0.002 n=6)
RandomPfx6Size/1_000_000     59.52 ± 0%   156.50 ± 0%  +162.94% (p=0.002 n=6)
RandomPfxSize/1_000          80.70 ± 2%   144.70 ± 1%   +79.31% (p=0.002 n=6)
RandomPfxSize/10_000         56.53 ± 0%   140.20 ± 0%  +148.01% (p=0.002 n=6)
RandomPfxSize/100_000        63.46 ± 0%   180.00 ± 0%  +183.64% (p=0.002 n=6)
RandomPfxSize/200_000        54.19 ± 0%   180.00 ± 0%  +232.16% (p=0.002 n=6)
RandomPfxSize/500_000        39.18 ± 0%   144.00 ± 0%  +267.53% (p=0.002 n=6)
RandomPfxSize/1_000_000      31.89 ± 0%   144.00 ± 0%  +351.55% (p=0.002 n=6)
geomean                      46.66         152.8       +227.53%
```

## update, insert/delete

`bart.Lite` is the fastest algorithm for updates.

```
goos: darwin
goarch: arm64
cpu: Apple M4 Max
                         │ lite/update.bm │            bart/update.bm            │
                         │   sec/route    │   sec/route    vs base               │
InsertRandomPfxs/1_000        70.51n ± 1%    82.89n ±  1%  +17.56% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.11n ± 1%    70.31n ±  2%  +29.94% (p=0.002 n=6)
InsertRandomPfxs/100_000      80.22n ± 1%   105.00n ±  5%  +30.89% (p=0.002 n=6)
InsertRandomPfxs/200_000      80.27n ± 1%   115.15n ±  2%  +43.45% (p=0.002 n=6)
DeleteRandomPfxs/1_000        57.87n ± 3%    67.30n ±  3%  +16.30% (p=0.002 n=6)
DeleteRandomPfxs/10_000       37.94n ± 3%    44.25n ±  4%  +16.65% (p=0.002 n=6)
DeleteRandomPfxs/100_000      83.95n ± 7%    95.58n ± 16%  +13.86% (p=0.002 n=6)
DeleteRandomPfxs/200_000      81.17n ± 7%   104.10n ±  7%  +28.25% (p=0.002 n=6)
geomean                       66.17n         82.22n        +24.25%

goos: darwin
goarch: arm64
cpu: Apple M4 Max
                         │ lite/update.bm │            fast/update.bm             │
                         │   sec/route    │   sec/route    vs base                │
InsertRandomPfxs/1_000        70.51n ± 1%   144.20n ±  1%  +104.51% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.11n ± 1%   102.95n ±  2%   +90.26% (p=0.002 n=6)
InsertRandomPfxs/100_000      80.22n ± 1%   149.55n ±  1%   +86.42% (p=0.002 n=6)
InsertRandomPfxs/200_000      80.27n ± 1%   162.35n ±  1%  +102.25% (p=0.002 n=6)
DeleteRandomPfxs/1_000        57.87n ± 3%   136.50n ±  4%  +135.87% (p=0.002 n=6)
DeleteRandomPfxs/10_000       37.94n ± 3%    81.62n ±  2%  +115.17% (p=0.002 n=6)
DeleteRandomPfxs/100_000      83.95n ± 7%   158.10n ±  8%   +88.33% (p=0.002 n=6)
DeleteRandomPfxs/200_000      81.17n ± 7%   159.00n ± 11%   +95.89% (p=0.002 n=6)
geomean                       66.17n         133.5n        +101.77%

goos: darwin
goarch: arm64
cpu: Apple M4 Max
                         │ lite/update.bm │          netipds/update.bm           │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000        70.51n ± 1%    94.07n ± 1%   +33.42% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.11n ± 1%   138.95n ± 1%  +156.79% (p=0.002 n=6)
InsertRandomPfxs/100_000      80.22n ± 1%   188.90n ± 1%  +135.48% (p=0.002 n=6)
InsertRandomPfxs/200_000      80.27n ± 1%   211.25n ± 1%  +163.17% (p=0.002 n=6)
DeleteRandomPfxs/1_000        57.87n ± 3%    43.22n ± 5%   -25.32% (p=0.002 n=6)
DeleteRandomPfxs/10_000       37.94n ± 3%   103.75n ± 2%  +173.49% (p=0.002 n=6)
DeleteRandomPfxs/100_000      83.95n ± 7%   157.45n ± 2%   +87.55% (p=0.002 n=6)
DeleteRandomPfxs/200_000      81.17n ± 7%   175.65n ± 6%  +116.40% (p=0.002 n=6)
geomean                       66.17n         126.3n        +90.85%

goos: darwin
goarch: arm64
cpu: Apple M4 Max
                         │ lite/update.bm │         critbitgo/update.bm          │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000        70.51n ± 1%    91.51n ± 2%   +29.78% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.11n ± 1%   185.60n ± 1%  +243.00% (p=0.002 n=6)
InsertRandomPfxs/100_000      80.22n ± 1%   267.80n ± 1%  +233.83% (p=0.002 n=6)
InsertRandomPfxs/200_000      80.27n ± 1%   296.55n ± 2%  +269.44% (p=0.002 n=6)
DeleteRandomPfxs/1_000        57.87n ± 3%    38.89n ± 2%   -32.80% (p=0.002 n=6)
DeleteRandomPfxs/10_000       37.94n ± 3%   103.45n ± 2%  +172.70% (p=0.002 n=6)
DeleteRandomPfxs/100_000      83.95n ± 7%   162.50n ± 2%   +93.57% (p=0.002 n=6)
DeleteRandomPfxs/200_000      81.17n ± 7%   185.45n ± 4%  +128.47% (p=0.002 n=6)
geomean                       66.17n         141.8n       +114.31%

goos: darwin
goarch: arm64
cpu: Apple M4 Max
                         │ lite/update.bm │          lpmtrie/update.bm           │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000        70.51n ± 1%   199.80n ± 2%  +183.36% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.11n ± 1%   239.30n ± 1%  +342.25% (p=0.002 n=6)
InsertRandomPfxs/100_000      80.22n ± 1%   312.45n ± 2%  +289.49% (p=0.002 n=6)
InsertRandomPfxs/200_000      80.27n ± 1%   358.85n ± 1%  +347.05% (p=0.002 n=6)
DeleteRandomPfxs/1_000        57.87n ± 3%    59.28n ± 3%         ~ (p=0.065 n=6)
DeleteRandomPfxs/10_000       37.94n ± 3%   128.30n ± 3%  +238.21% (p=0.002 n=6)
DeleteRandomPfxs/100_000      83.95n ± 7%   201.25n ± 1%  +139.73% (p=0.002 n=6)
DeleteRandomPfxs/200_000      81.17n ± 7%   232.80n ± 4%  +186.81% (p=0.002 n=6)
geomean                       66.17n         192.8n       +191.39%

goos: darwin
goarch: arm64
cpu: Apple M4 Max
                         │ lite/update.bm │      kentik-patricia/update.bm       │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000        70.51n ± 1%   102.80n ± 0%   +45.79% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.11n ± 1%   182.05n ± 0%  +236.44% (p=0.002 n=6)
InsertRandomPfxs/100_000      80.22n ± 1%   237.25n ± 1%  +195.75% (p=0.002 n=6)
InsertRandomPfxs/200_000      80.27n ± 1%   266.50n ± 1%  +232.00% (p=0.002 n=6)
DeleteRandomPfxs/1_000        57.87n ± 3%   121.40n ± 3%  +109.78% (p=0.002 n=6)
DeleteRandomPfxs/10_000       37.94n ± 3%   175.40n ± 2%  +362.37% (p=0.002 n=6)
DeleteRandomPfxs/100_000      83.95n ± 7%   256.80n ± 2%  +205.90% (p=0.002 n=6)
DeleteRandomPfxs/200_000      81.17n ± 7%   291.60n ± 2%  +259.25% (p=0.002 n=6)
geomean                       66.17n         192.5n       +190.94%
```

