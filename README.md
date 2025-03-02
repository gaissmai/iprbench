# iprbench
comparing benchmarks for some golang IP routing table implementations:

```
	github.com/gaissmai/bart
	github.com/tailscale/art
	github.com/aromatt/netipds
	github.com/k-sone/critbitgo
	github.com/Asphaltt/lpmtrie
	github.com/yl2chen/cidranger
	github.com/gaissmai/cidrtree
```

The ~1_000_000 **Tier1** prefix test records (IPv4 and IPv6 routes) are from a full Internet routing table with typical
ISP prefix distribution.

In comparison, the prefix lengths for the _real-world_ random test sets are equally distributed between /8-28 for IPv4
and /16-56 bits for IPv6 (limited to the 2000::/3 global unicast address space).

The _real-world_ **RandomPrefixes** without IP versions labeling are composed of a distribution of 4 parts IPv4
to 1 part IPv6 random prefixes, which is approximately the current ratio in the Internet backbone routers.

## make your own benchmarks

```
  $ make dep
  $ make -B all   # takes some time!
```

## lpm (longest-prefix-match)

`bart` is the fastest software algorithm for IP-address longest-prefix-match.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │              art/lpm.bm               │
                                     │    sec/op    │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4            28.36n ±  8%   50.20n ±  1%   +77.02% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            44.84n ± 23%   66.22n ±  7%   +47.71% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             38.84n ±  5%   33.08n ±  1%   -14.84% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             49.34n ± 23%   39.25n ± 16%         ~ (p=0.114 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     23.78n ±  8%   49.69n ±  2%  +108.94% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     24.55n ±  2%   51.26n ±  1%  +108.74% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      24.75n ±  6%   32.44n ±  1%   +31.06% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      24.11n ±  3%   32.78n ±  3%   +35.96% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    30.14n ±  6%   50.74n ±  2%   +68.34% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    26.73n ± 19%   53.81n ±  5%  +101.27% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     30.57n ±  8%   32.07n ±  2%         ~ (p=0.136 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     26.86n ± 13%   33.80n ±  5%   +25.86% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   15.53n ± 64%   68.01n ±  9%  +337.93% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   35.58n ±  9%   55.94n ±  8%   +57.22% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    29.32n ± 14%   33.12n ±  4%   +12.98% (p=0.006 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    34.17n ±  6%   40.34n ±  6%   +18.07% (p=0.000 n=20)
geomean                                29.40n         43.69n         +48.60%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            netipds/lpm.bm             │
                                     │    sec/op    │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4            28.36n ±  8%   71.92n ±  6%  +153.66% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            44.84n ± 23%   57.37n ± 14%   +27.96% (p=0.001 n=20)
LpmTier1Pfxs/RandomMissIP4             38.84n ±  5%   70.43n ±  5%   +81.30% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             49.34n ± 23%   57.93n ± 15%   +17.40% (p=0.013 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     23.78n ±  8%   43.05n ±  4%   +81.06% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     24.55n ±  2%   34.61n ±  3%   +40.93% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      24.75n ±  6%   45.11n ±  4%   +82.21% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      24.11n ±  3%   36.64n ±  5%   +51.97% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    30.14n ±  6%   49.23n ±  4%   +63.33% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    26.73n ± 19%   42.10n ±  4%   +57.47% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     30.57n ±  8%   53.63n ±  3%   +75.45% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     26.86n ± 13%   42.80n ±  4%   +59.36% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   15.53n ± 64%   23.66n ± 46%   +52.32% (p=0.004 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   35.58n ±  9%   45.96n ±  4%   +29.19% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    29.32n ± 14%   61.71n ±  3%  +110.49% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    34.17n ±  6%   48.41n ±  3%   +41.67% (p=0.000 n=20)
geomean                                29.40n         47.36n         +61.11%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            critbitgo/lpm.bm             │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            28.36n ±  8%   164.25n ±  3%   +479.26% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            44.84n ± 23%   248.50n ± 10%   +454.25% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             38.84n ±  5%   786.10n ± 11%  +1923.68% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             49.34n ± 23%   579.50n ± 24%  +1074.38% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     23.78n ±  8%    88.23n ±  7%   +271.05% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     24.55n ±  2%   120.75n ±  3%   +391.75% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      24.75n ±  6%   190.55n ± 17%   +669.74% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      24.11n ±  3%   162.50n ±  5%   +573.99% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    30.14n ±  6%   108.50n ±  5%   +259.93% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    26.73n ± 19%   143.45n ±  4%   +436.56% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     30.57n ±  8%   292.50n ± 14%   +856.82% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     26.86n ± 13%   244.40n ± 14%   +809.90% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   15.53n ± 64%   130.85n ±  7%   +742.56% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   35.58n ±  9%   157.40n ±  2%   +342.38% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    29.32n ± 14%   464.40n ± 21%  +1484.17% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    34.17n ±  6%   349.25n ± 11%   +922.10% (p=0.000 n=20)
geomean                                29.40n          216.2n         +635.47%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             lpmtrie/lpm.bm              │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            28.36n ±  8%   319.55n ±  5%  +1026.96% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            44.84n ± 23%   350.65n ± 15%   +682.09% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             38.84n ±  5%   286.75n ±  9%   +638.19% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             49.34n ± 23%   274.85n ± 23%   +457.00% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     23.78n ±  8%   132.40n ± 12%   +456.77% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     24.55n ±  2%   112.95n ±  3%   +359.99% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      24.75n ±  6%   147.40n ± 21%   +495.44% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      24.11n ±  3%   109.30n ± 31%   +353.34% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    30.14n ±  6%   166.80n ±  7%   +453.33% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    26.73n ± 19%   156.15n ±  6%   +484.07% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     30.57n ±  8%   164.25n ± 20%   +437.29% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     26.86n ± 13%   143.45n ±  8%   +434.07% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   15.53n ± 64%   217.90n ±  9%  +1303.09% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   35.58n ±  9%   187.75n ± 13%   +427.68% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    29.32n ± 14%   211.70n ±  7%   +622.16% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    34.17n ±  6%   210.80n ±  4%   +516.92% (p=0.000 n=20)
geomean                                29.40n          187.7n         +538.67%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             cidranger/lpm.bm             │
                                     │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            28.36n ±  8%   342.55n ±  14%  +1108.08% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            44.84n ± 23%   365.50n ±  29%   +715.21% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             38.84n ±  5%   303.65n ±  10%   +681.70% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             49.34n ± 23%   372.90n ±  20%   +655.70% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     23.78n ±  8%   129.10n ±   8%   +442.89% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     24.55n ±  2%   162.85n ±   3%   +563.21% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      24.75n ±  6%   137.20n ±   8%   +454.23% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      24.11n ±  3%   166.20n ±   7%   +589.34% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    30.14n ±  6%   178.45n ±  10%   +491.97% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    26.73n ± 19%   213.80n ±  10%   +699.70% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     30.57n ±  8%   181.30n ±   5%   +493.07% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     26.86n ± 13%   206.55n ±   6%   +668.99% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   15.53n ± 64%    46.79n ± 103%   +201.29% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   35.58n ±  9%   258.10n ±   5%   +625.41% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    29.32n ± 14%   234.80n ±   6%   +700.96% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    34.17n ±  6%   262.00n ±   5%   +666.75% (p=0.000 n=20)
geomean                                29.40n          201.2n          +584.44%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             cidrtree/lpm.bm              │
                                     │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            28.36n ±  8%    764.10n ± 15%  +2594.76% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            44.84n ± 23%    972.10n ± 18%  +2068.17% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             38.84n ±  5%   1075.00n ± 15%  +2667.41% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             49.34n ± 23%   1297.00n ± 19%  +2528.43% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     23.78n ±  8%    422.35n ± 20%  +1676.07% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     24.55n ±  2%    392.60n ± 19%  +1498.86% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      24.75n ±  6%    637.80n ± 44%  +2476.45% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      24.11n ±  3%    576.20n ± 12%  +2289.88% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    30.14n ±  6%    481.10n ± 27%  +1495.95% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    26.73n ± 19%    643.20n ± 23%  +2305.84% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     30.57n ±  8%    699.15n ± 19%  +2187.05% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     26.86n ± 13%    765.05n ± 22%  +2748.29% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   15.53n ± 64%    638.60n ± 16%  +4012.04% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   35.58n ±  9%    819.60n ± 16%  +2203.54% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    29.32n ± 14%    908.85n ± 13%  +3000.29% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    34.17n ±  6%   1134.00n ± 19%  +3218.70% (p=0.000 n=20)
geomean                                29.40n           723.5n        +2361.08%
```

## size of the routing tables


`bart` has two orders of magnitude lower memory consumption compared to `art`
and is even better in low memory consumption to a binary search tree, like the `cidrtree`.


```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │              art/size.bm               │
                       │ bytes/route  │ bytes/route   vs base                  │
Tier1PfxSize/1_000         88.13 ± 2%   7591.00 ± 0%   +8513.41% (p=0.002 n=6)
Tier1PfxSize/10_000        66.04 ± 0%   4889.00 ± 0%   +7303.09% (p=0.002 n=6)
Tier1PfxSize/100_000       30.24 ± 0%   1669.00 ± 0%   +5419.18% (p=0.002 n=6)
Tier1PfxSize/200_000       20.48 ± 0%   1098.00 ± 0%   +5261.33% (p=0.002 n=6)
RandomPfx4Size/1_000       64.63 ± 3%   5194.00 ± 0%   +7936.52% (p=0.002 n=6)
RandomPfx4Size/10_000      42.67 ± 0%   4436.00 ± 0%  +10296.06% (p=0.002 n=6)
RandomPfx4Size/100_000     58.23 ± 0%   4692.00 ± 0%   +7957.70% (p=0.002 n=6)
RandomPfx4Size/200_000     55.16 ± 0%   4379.00 ± 0%   +7838.72% (p=0.002 n=6)
RandomPfx6Size/1_000       70.43 ± 3%   6859.00 ± 0%   +9638.75% (p=0.002 n=6)
RandomPfx6Size/10_000      84.00 ± 0%   7306.00 ± 0%   +8597.62% (p=0.002 n=6)
RandomPfx6Size/100_000     56.79 ± 0%   5730.00 ± 0%   +9989.80% (p=0.002 n=6)
RandomPfx6Size/200_000     53.37 ± 0%   5529.00 ± 0%  +10259.75% (p=0.002 n=6)
RandomPfxSize/1_000        86.59 ± 2%   7590.00 ± 0%   +8665.45% (p=0.002 n=6)
RandomPfxSize/10_000       59.15 ± 0%   6201.00 ± 0%  +10383.52% (p=0.002 n=6)
RandomPfxSize/100_000      70.99 ± 0%   5849.00 ± 0%   +8139.19% (p=0.002 n=6)
RandomPfxSize/200_000      64.32 ± 0%   5300.00 ± 0%   +8140.05% (p=0.002 n=6)
geomean                    57.27        4.669Ki        +8248.66%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │           netipds/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         88.13 ± 2%   101.00 ± 2%   +14.60% (p=0.002 n=6)
Tier1PfxSize/10_000        66.04 ± 0%    96.08 ± 0%   +45.49% (p=0.002 n=6)
Tier1PfxSize/100_000       30.24 ± 0%    94.34 ± 0%  +211.97% (p=0.002 n=6)
Tier1PfxSize/200_000       20.48 ± 0%    93.27 ± 0%  +355.42% (p=0.002 n=6)
RandomPfx4Size/1_000       64.63 ± 3%    82.98 ± 2%   +28.39% (p=0.002 n=6)
RandomPfx4Size/10_000      42.67 ± 0%    77.96 ± 0%   +82.70% (p=0.002 n=6)
RandomPfx4Size/100_000     58.23 ± 0%    75.26 ± 0%   +29.25% (p=0.002 n=6)
RandomPfx4Size/200_000     55.16 ± 0%    74.60 ± 0%   +35.24% (p=0.002 n=6)
RandomPfx6Size/1_000       70.43 ± 3%   100.20 ± 2%   +42.27% (p=0.002 n=6)
RandomPfx6Size/10_000      84.00 ± 0%    94.12 ± 0%   +12.05% (p=0.002 n=6)
RandomPfx6Size/100_000     56.79 ± 0%    87.79 ± 0%   +54.59% (p=0.002 n=6)
RandomPfx6Size/200_000     53.37 ± 0%    85.95 ± 0%   +61.05% (p=0.002 n=6)
RandomPfxSize/1_000        86.59 ± 2%    99.78 ± 2%   +15.23% (p=0.002 n=6)
RandomPfxSize/10_000       59.15 ± 0%    94.42 ± 0%   +59.63% (p=0.002 n=6)
RandomPfxSize/100_000      70.99 ± 0%    87.92 ± 0%   +23.85% (p=0.002 n=6)
RandomPfxSize/200_000      64.32 ± 0%    84.72 ± 0%   +31.72% (p=0.002 n=6)
geomean                    57.27         89.00        +55.40%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          critbitgo/size.bm          │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         88.13 ± 2%   119.50 ± 2%   +35.60% (p=0.002 n=6)
Tier1PfxSize/10_000        66.04 ± 0%   114.70 ± 0%   +73.68% (p=0.002 n=6)
Tier1PfxSize/100_000       30.24 ± 0%   114.40 ± 0%  +278.31% (p=0.002 n=6)
Tier1PfxSize/200_000       20.48 ± 0%   114.40 ± 0%  +458.59% (p=0.002 n=6)
RandomPfx4Size/1_000       64.63 ± 3%   116.10 ± 2%   +79.64% (p=0.002 n=6)
RandomPfx4Size/10_000      42.67 ± 0%   112.40 ± 0%  +163.42% (p=0.002 n=6)
RandomPfx4Size/100_000     58.23 ± 0%   112.00 ± 0%   +92.34% (p=0.002 n=6)
RandomPfx4Size/200_000     55.16 ± 0%   112.00 ± 0%  +103.05% (p=0.002 n=6)
RandomPfx6Size/1_000       70.43 ± 3%   132.40 ± 1%   +87.99% (p=0.002 n=6)
RandomPfx6Size/10_000      84.00 ± 0%   128.40 ± 0%   +52.86% (p=0.002 n=6)
RandomPfx6Size/100_000     56.79 ± 0%   128.00 ± 0%  +125.39% (p=0.002 n=6)
RandomPfx6Size/200_000     53.37 ± 0%   128.00 ± 0%  +139.84% (p=0.002 n=6)
RandomPfxSize/1_000        86.59 ± 2%   119.40 ± 2%   +37.89% (p=0.002 n=6)
RandomPfxSize/10_000       59.15 ± 0%   115.60 ± 0%   +95.44% (p=0.002 n=6)
RandomPfxSize/100_000      70.99 ± 0%   115.20 ± 0%   +62.28% (p=0.002 n=6)
RandomPfxSize/200_000      64.32 ± 0%   115.20 ± 0%   +79.10% (p=0.002 n=6)
geomean                    57.27         118.4       +106.80%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │           lpmtrie/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         88.13 ± 2%   215.40 ± 5%  +144.41% (p=0.002 n=6)
Tier1PfxSize/10_000        66.04 ± 0%   210.50 ± 5%  +218.75% (p=0.002 n=6)
Tier1PfxSize/100_000       30.24 ± 0%   209.90 ± 5%  +594.11% (p=0.002 n=6)
Tier1PfxSize/200_000       20.48 ± 0%   209.20 ± 5%  +921.48% (p=0.002 n=6)
RandomPfx4Size/1_000       64.63 ± 3%   190.80 ± 9%  +195.22% (p=0.002 n=6)
RandomPfx4Size/10_000      42.67 ± 0%   185.60 ± 8%  +334.97% (p=0.002 n=6)
RandomPfx4Size/100_000     58.23 ± 0%   182.40 ± 8%  +213.24% (p=0.002 n=6)
RandomPfx4Size/200_000     55.16 ± 0%   181.70 ± 9%  +229.41% (p=0.002 n=6)
RandomPfx6Size/1_000       70.43 ± 3%   227.80 ± 8%  +223.44% (p=0.002 n=6)
RandomPfx6Size/10_000      84.00 ± 0%   222.50 ± 8%  +164.88% (p=0.002 n=6)
RandomPfx6Size/100_000     56.79 ± 0%   213.90 ± 9%  +276.65% (p=0.002 n=6)
RandomPfx6Size/200_000     53.37 ± 0%   210.50 ± 9%  +294.42% (p=0.002 n=6)
RandomPfxSize/1_000        86.59 ± 2%   215.10 ± 5%  +148.41% (p=0.002 n=6)
RandomPfxSize/10_000       59.15 ± 0%   210.40 ± 5%  +255.71% (p=0.002 n=6)
RandomPfxSize/100_000      70.99 ± 0%   204.00 ± 7%  +187.36% (p=0.002 n=6)
RandomPfxSize/200_000      64.32 ± 0%   200.10 ± 7%  +211.10% (p=0.002 n=6)
geomean                    57.27         205.2       +258.24%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidranger/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                 │
Tier1PfxSize/1_000         88.13 ± 2%   539.60 ± 3%   +512.28% (p=0.002 n=6)
Tier1PfxSize/10_000        66.04 ± 0%   533.90 ± 3%   +708.45% (p=0.002 n=6)
Tier1PfxSize/100_000       30.24 ± 0%   527.20 ± 2%  +1643.39% (p=0.002 n=6)
Tier1PfxSize/200_000       20.48 ± 0%   522.20 ± 2%  +2449.80% (p=0.002 n=6)
RandomPfx4Size/1_000       64.63 ± 3%   447.50 ± 3%   +592.40% (p=0.002 n=6)
RandomPfx4Size/10_000      42.67 ± 0%   438.40 ± 3%   +927.42% (p=0.002 n=6)
RandomPfx4Size/100_000     58.23 ± 0%   427.30 ± 3%   +633.81% (p=0.002 n=6)
RandomPfx4Size/200_000     55.16 ± 0%   424.30 ± 3%   +669.22% (p=0.002 n=6)
RandomPfx6Size/1_000       70.43 ± 3%   595.10 ± 0%   +744.95% (p=0.002 n=6)
RandomPfx6Size/10_000      84.00 ± 0%   580.40 ± 0%   +590.95% (p=0.002 n=6)
RandomPfx6Size/100_000     56.79 ± 0%   548.00 ± 0%   +864.96% (p=0.002 n=6)
RandomPfx6Size/200_000     53.37 ± 0%   538.30 ± 0%   +908.62% (p=0.002 n=6)
RandomPfxSize/1_000        86.59 ± 2%   539.50 ± 2%   +523.05% (p=0.002 n=6)
RandomPfxSize/10_000       59.15 ± 0%   528.50 ± 2%   +793.49% (p=0.002 n=6)
RandomPfxSize/100_000      70.99 ± 0%   499.50 ± 2%   +603.62% (p=0.002 n=6)
RandomPfxSize/200_000      64.32 ± 0%   484.60 ± 2%   +653.42% (p=0.002 n=6)
geomean                    57.27         508.3        +787.51%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidrtree/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         88.13 ± 2%    69.26 ± 3%   -21.41% (p=0.002 n=6)
Tier1PfxSize/10_000        66.04 ± 0%    64.37 ± 0%    -2.53% (p=0.002 n=6)
Tier1PfxSize/100_000       30.24 ± 0%    64.04 ± 0%  +111.77% (p=0.002 n=6)
Tier1PfxSize/200_000       20.48 ± 0%    64.02 ± 0%  +212.60% (p=0.002 n=6)
RandomPfx4Size/1_000       64.63 ± 3%    68.16 ± 3%    +5.46% (p=0.002 n=6)
RandomPfx4Size/10_000      42.67 ± 0%    64.37 ± 0%   +50.86% (p=0.002 n=6)
RandomPfx4Size/100_000     58.23 ± 0%    64.04 ± 0%    +9.98% (p=0.002 n=6)
RandomPfx4Size/200_000     55.16 ± 0%    64.02 ± 0%   +16.06% (p=0.002 n=6)
RandomPfx6Size/1_000       70.43 ± 3%    68.41 ± 3%    -2.87% (p=0.002 n=6)
RandomPfx6Size/10_000      84.00 ± 0%    64.37 ± 0%   -23.37% (p=0.002 n=6)
RandomPfx6Size/100_000     56.79 ± 0%    64.04 ± 0%   +12.77% (p=0.002 n=6)
RandomPfx6Size/200_000     53.37 ± 0%    64.02 ± 0%   +19.96% (p=0.002 n=6)
RandomPfxSize/1_000        86.59 ± 2%    68.16 ± 3%   -21.28% (p=0.002 n=6)
RandomPfxSize/10_000       59.15 ± 0%    64.37 ± 0%    +8.83% (p=0.002 n=6)
RandomPfxSize/100_000      70.99 ± 0%    64.04 ± 0%    -9.79% (p=0.002 n=6)
RandomPfxSize/200_000      64.32 ± 0%    64.02 ± 0%    -0.47% (p=0.002 n=6)
geomean                    57.27         65.20        +13.86%
```

## update, insert/delete

`bart` is by far the fastest algorithm for updates.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │              art/update.bm              │
                         │   sec/route    │   sec/route     vs base                 │
InsertRandomPfxs/1_000       29.97n ±  2%   274.85n ±  17%   +817.08% (p=0.002 n=6)
InsertRandomPfxs/10_000      40.25n ± 21%   530.50n ±  11%  +1218.01% (p=0.002 n=6)
InsertRandomPfxs/100_000     114.3n ±  2%   1875.0n ±  39%  +1540.42% (p=0.002 n=6)
InsertRandomPfxs/200_000     196.3n ±  3%   1840.5n ± 189%   +837.36% (p=0.002 n=6)
DeleteRandomPfxs/1_000       15.91n ±  1%    17.70n ±  13%    +11.25% (p=0.002 n=6)
DeleteRandomPfxs/10_000      15.52n ±  1%    16.84n ±   2%     +8.51% (p=0.002 n=6)
DeleteRandomPfxs/100_000     15.72n ±  1%    17.63n ±   2%    +12.15% (p=0.002 n=6)
DeleteRandomPfxs/200_000     16.80n ±  3%    20.95n ±  14%    +24.70% (p=0.002 n=6)
geomean                      33.95n          123.9n          +264.81%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           netipds/update.bm            │
                         │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000       29.97n ±  2%   127.70n ±  9%   +326.09% (p=0.002 n=6)
InsertRandomPfxs/10_000      40.25n ± 21%   190.25n ±  3%   +372.67% (p=0.002 n=6)
InsertRandomPfxs/100_000     114.3n ±  2%    446.6n ±  3%   +290.77% (p=0.002 n=6)
InsertRandomPfxs/200_000     196.3n ±  3%    602.4n ±  1%   +206.80% (p=0.002 n=6)
DeleteRandomPfxs/1_000       15.91n ±  1%   111.10n ±  4%   +598.08% (p=0.002 n=6)
DeleteRandomPfxs/10_000      15.52n ±  1%   183.25n ±  7%  +1080.73% (p=0.002 n=6)
DeleteRandomPfxs/100_000     15.72n ±  1%   434.85n ±  9%  +2666.22% (p=0.002 n=6)
DeleteRandomPfxs/200_000     16.80n ±  3%   639.90n ± 49%  +3707.80% (p=0.002 n=6)
geomean                      33.95n          279.3n         +722.59%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │          critbitgo/update.bm          │
                         │   sec/route    │   sec/route    vs base                │
InsertRandomPfxs/1_000       29.97n ±  2%   171.95n ± 19%  +473.74% (p=0.002 n=6)
InsertRandomPfxs/10_000      40.25n ± 21%   239.00n ±  5%  +493.79% (p=0.002 n=6)
InsertRandomPfxs/100_000     114.3n ±  2%    624.9n ±  7%  +446.72% (p=0.002 n=6)
InsertRandomPfxs/200_000     196.3n ±  3%    779.3n ±  1%  +296.89% (p=0.002 n=6)
DeleteRandomPfxs/1_000       15.91n ±  1%    73.88n ± 19%  +364.18% (p=0.002 n=6)
DeleteRandomPfxs/10_000      15.52n ±  1%    79.23n ±  7%  +410.53% (p=0.002 n=6)
DeleteRandomPfxs/100_000     15.72n ±  1%    79.19n ± 11%  +403.72% (p=0.002 n=6)
DeleteRandomPfxs/200_000     16.80n ±  3%    90.64n ± 10%  +439.33% (p=0.002 n=6)
geomean                      33.95n          174.0n        +412.50%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           lpmtrie/update.bm            │
                         │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000       29.97n ±  2%   371.45n ±  5%  +1139.41% (p=0.002 n=6)
InsertRandomPfxs/10_000      40.25n ± 21%   515.85n ±  6%  +1181.61% (p=0.002 n=6)
InsertRandomPfxs/100_000     114.3n ±  2%   1239.5n ±  4%   +984.43% (p=0.002 n=6)
InsertRandomPfxs/200_000     196.3n ±  3%   1275.5n ±  6%   +549.61% (p=0.002 n=6)
DeleteRandomPfxs/1_000       15.91n ±  1%    59.47n ±  1%   +273.67% (p=0.002 n=6)
DeleteRandomPfxs/10_000      15.52n ±  1%   137.25n ±  1%   +784.34% (p=0.002 n=6)
DeleteRandomPfxs/100_000     15.72n ±  1%   340.20n ±  5%  +2064.12% (p=0.002 n=6)
DeleteRandomPfxs/200_000     16.80n ±  3%   511.55n ± 11%  +2944.03% (p=0.002 n=6)
geomean                      33.95n          379.5n        +1017.72%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │            cidranger/update.bm            │
                         │   sec/route    │    sec/route     vs base                  │
InsertRandomPfxs/1_000       29.97n ±  2%   4484.00n ±  22%  +14861.63% (p=0.002 n=6)
InsertRandomPfxs/10_000      40.25n ± 21%   6842.00n ±  26%  +16898.76% (p=0.002 n=6)
InsertRandomPfxs/100_000     114.3n ±  2%   10893.5n ±   1%   +9430.62% (p=0.002 n=6)
InsertRandomPfxs/200_000     196.3n ±  3%   12224.5n ±  21%   +6125.87% (p=0.002 n=6)
DeleteRandomPfxs/1_000       15.91n ±  1%     99.78n ±   9%    +526.96% (p=0.002 n=6)
DeleteRandomPfxs/10_000      15.52n ±  1%     90.66n ±  11%    +484.15% (p=0.002 n=6)
DeleteRandomPfxs/100_000     15.72n ±  1%    148.10n ±  16%    +842.11% (p=0.002 n=6)
DeleteRandomPfxs/200_000     16.80n ±  3%   1475.00n ± 271%   +8677.15% (p=0.002 n=6)
geomean                      33.95n           1.298µ          +3723.77%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │            cidrtree/update.bm            │
                         │   sec/route    │    sec/route     vs base                 │
InsertRandomPfxs/1_000       29.97n ±  2%   1285.00n ±   3%  +4187.62% (p=0.002 n=6)
InsertRandomPfxs/10_000      40.25n ± 21%   1953.00n ±   2%  +4752.17% (p=0.002 n=6)
InsertRandomPfxs/100_000     114.3n ±  2%    4355.0n ±  19%  +3710.15% (p=0.002 n=6)
InsertRandomPfxs/200_000     196.3n ±  3%    5557.0n ±  17%  +2730.15% (p=0.002 n=6)
DeleteRandomPfxs/1_000       15.91n ±  1%     16.82n ±  28%          ~ (p=0.368 n=6)
DeleteRandomPfxs/10_000      15.52n ±  1%     15.12n ±  17%          ~ (p=0.056 n=6)
DeleteRandomPfxs/100_000     15.72n ±  1%     23.55n ±   4%    +49.84% (p=0.002 n=6)
DeleteRandomPfxs/200_000     16.80n ±  3%    556.70n ± 821%  +3212.70% (p=0.002 n=6)
geomean                      33.95n           345.4n          +917.24%
```
