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

The ~1_000_000 **Tier1** prefix test records (IPv4 and IPv6 routes) are from a full Internet
routing table with typical ISP prefix distribution.

In comparison, the prefix lengths for the _real-world_ random test sets are equally distributed
between /8-28 for IPv4 and /16-56 bits for IPv6 (limited to the 2000::/3 global unicast address space).

The _real-world_ **RandomPrefixes** without IP versions labeling are composed of a distribution
of 4 parts IPv4 to 1 part IPv6 random prefixes, which is approximately the current ratio in the Internet backbone routers.

## make your own benchmarks

```
  $ # set the proper cpu feature flags, e.g.
  $ export GOAMD64='v2'

  $ make dep
  $ make -B all   # takes some time!
```

## lpm (longest-prefix-match)

`bart` is by far the fastest software algorithm for IP-address longest-prefix-match.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │              art/lpm.bm               │
                                     │    sec/op    │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4            17.16n ±  4%   45.64n ± 16%  +165.97% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            35.05n ± 24%   62.52n ±  4%   +78.35% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             19.54n ±  1%   28.98n ±  0%   +48.35% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             26.77n ± 30%   36.27n ± 22%   +35.49% (p=0.001 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.73n ±  9%   44.97n ±  1%  +168.80% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.26n ±  0%   46.17n ±  2%  +139.75% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.84n ±  4%   28.76n ±  0%   +70.73% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.25n ±  0%   29.00n ±  7%   +50.65% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.18n ±  0%   45.10n ±  0%  +148.05% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.29n ±  1%   46.07n ±  6%  +138.83% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.18n ±  0%   28.77n ±  1%   +58.22% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.27n ±  0%   28.89n ±  6%   +49.92% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   18.27n ± 21%   49.25n ± 10%  +169.64% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   25.72n ± 24%   48.80n ±  7%   +89.70% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.38n ± 38%   28.87n ±  2%   +57.07% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    25.89n ±  1%   33.01n ±  6%   +27.53% (p=0.000 n=20)
geomean                                20.40n         38.20n         +87.20%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            netipds/lpm.bm             │
                                     │    sec/op    │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4            17.16n ±  4%   60.01n ± 12%  +249.71% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            35.05n ± 24%   47.14n ± 18%   +34.46% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             19.54n ±  1%   64.91n ±  6%  +232.28% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             26.77n ± 30%   52.71n ± 11%   +96.92% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.73n ±  9%   39.52n ±  5%  +136.19% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.26n ±  0%   32.62n ±  9%   +69.37% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.84n ±  4%   40.16n ±  3%  +138.41% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.25n ±  0%   32.46n ±  6%   +68.65% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.18n ±  0%   45.75n ±  4%  +151.65% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.29n ±  1%   37.01n ±  3%   +91.86% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.18n ±  0%   49.38n ±  4%  +171.64% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.27n ±  0%   38.69n ±  3%  +100.75% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   18.27n ± 21%   47.85n ±  9%  +161.98% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   25.72n ± 24%   42.10n ±  2%   +63.63% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.38n ± 38%   53.68n ±  4%  +192.08% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    25.89n ±  1%   44.21n ±  3%   +70.81% (p=0.000 n=20)
geomean                                20.40n         44.67n        +118.94%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            critbitgo/lpm.bm             │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            17.16n ±  4%   162.65n ±  6%   +847.84% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            35.05n ± 24%   250.05n ± 14%   +613.31% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             19.54n ±  1%   630.25n ± 18%  +3126.26% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             26.77n ± 30%   485.70n ± 12%  +1714.68% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.73n ±  9%    83.98n ±  5%   +402.00% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.26n ±  0%   105.30n ±  3%   +446.73% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.84n ±  4%   149.45n ± 13%   +787.21% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.25n ±  0%   159.20n ±  8%   +727.01% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.18n ±  0%    98.94n ±  5%   +444.22% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.29n ±  1%   117.45n ±  3%   +508.86% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.18n ±  0%   254.40n ± 10%  +1299.34% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.27n ±  0%   221.05n ± 15%  +1047.12% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   18.27n ± 21%   110.25n ±  5%   +503.61% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   25.72n ± 24%   132.15n ±  2%   +413.70% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.38n ± 38%   472.55n ± 16%  +2471.00% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    25.89n ±  1%   332.55n ± 18%  +1184.72% (p=0.000 n=20)
geomean                                20.40n          193.7n         +849.20%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             lpmtrie/lpm.bm              │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            17.16n ±  4%   279.85n ± 10%  +1530.83% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            35.05n ± 24%   323.35n ± 17%   +822.41% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             19.54n ±  1%   260.05n ±  5%  +1231.20% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             26.77n ± 30%   234.05n ± 15%   +774.46% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.73n ±  9%   135.45n ±  8%   +709.62% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.26n ±  0%   120.30n ± 23%   +524.61% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.84n ±  4%   115.15n ± 17%   +583.59% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.25n ±  0%    95.28n ± 37%   +394.96% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.18n ±  0%   153.75n ± 13%   +745.71% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.29n ±  1%   140.00n ±  9%   +625.76% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.18n ±  0%   163.00n ± 11%   +796.59% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.27n ±  0%   135.40n ±  5%   +602.65% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   18.27n ± 21%   197.85n ±  5%   +983.22% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   25.72n ± 24%   186.65n ±  2%   +625.56% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.38n ± 38%   199.00n ±  4%   +982.70% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    25.89n ±  1%   189.40n ±  5%   +631.70% (p=0.000 n=20)
geomean                                20.40n          173.1n         +748.47%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            cidranger/lpm.bm             │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            17.16n ±  4%   248.70n ± 23%  +1349.30% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            35.05n ± 24%   276.80n ± 29%   +689.62% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             19.54n ±  1%   281.30n ± 10%  +1339.98% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             26.77n ± 30%   290.95n ± 17%   +987.05% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.73n ±  9%   119.00n ± 12%   +611.30% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.26n ±  0%   155.20n ± 11%   +705.82% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.84n ±  4%   113.55n ±  5%   +574.09% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.25n ±  0%   149.85n ± 10%   +678.44% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.18n ±  0%   157.00n ±  8%   +763.59% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.29n ±  1%   181.10n ±  6%   +838.83% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.18n ±  0%   154.55n ±  4%   +750.11% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.27n ±  0%   178.65n ±  5%   +827.09% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   18.27n ± 21%   168.90n ± 20%   +824.72% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   25.72n ± 24%   226.85n ±  3%   +781.83% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.38n ± 38%   188.65n ±  4%   +926.39% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    25.89n ±  1%   226.55n ±  7%   +775.22% (p=0.000 n=20)
geomean                                20.40n          187.2n         +817.50%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             cidrtree/lpm.bm              │
                                     │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            17.16n ±  4%    701.40n ± 12%  +3987.41% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            35.05n ± 24%    802.75n ± 19%  +2189.97% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             19.54n ±  1%    980.15n ± 13%  +4917.40% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             26.77n ± 30%   1099.00n ± 30%  +4006.11% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.73n ±  9%    336.60n ± 12%  +1911.95% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.26n ±  0%    396.50n ± 30%  +1958.67% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.84n ±  4%    380.40n ± 17%  +2158.24% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.25n ±  0%    489.20n ± 29%  +2441.30% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.18n ±  0%    487.45n ± 15%  +2581.24% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.29n ±  1%    520.20n ± 11%  +2596.73% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.18n ±  0%    637.55n ± 18%  +3406.88% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.27n ±  0%    795.75n ± 22%  +4029.48% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   18.27n ± 21%    563.65n ± 12%  +2985.96% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   25.72n ± 24%    693.60n ± 16%  +2596.21% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.38n ± 38%    879.60n ± 19%  +4685.64% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    25.89n ±  1%   1085.00n ± 11%  +4091.62% (p=0.000 n=20)
geomean                                20.40n           636.4n        +3018.74%
```

## size of the routing tables


`bart` has two orders of magnitude lower memory consumption compared to `art`
and becomes more efficient the more prefixes are stored in the table.

For `art` and `cidranger` no benchmarks are made for 500_000 and 1_000_000,
with `art` the memory consumption is too high and with cidranger the insert takes too long.


```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │               art/size.bm               │
                         │ bytes/route  │ bytes/route   vs base                   │
Tier1PfxSize/1_000           103.9 ± 2%    7591.0 ± 0%  +7206.06% (p=0.002 n=6)
Tier1PfxSize/10_000          76.05 ± 0%   4889.00 ± 0%  +6328.67% (p=0.002 n=6)
Tier1PfxSize/100_000         41.16 ± 0%   1669.00 ± 0%  +3954.91% (p=0.002 n=6)
Tier1PfxSize/200_000         32.85 ± 0%   1098.00 ± 0%  +3242.47% (p=0.002 n=6)
Tier1PfxSize/500_000         25.19 ± 0%
Tier1PfxSize/1_000_000       21.39 ± 0%
RandomPfx4Size/1_000         74.51 ± 2%   5259.00 ± 0%  +6958.11% (p=0.002 n=6)
RandomPfx4Size/10_000        48.24 ± 0%   4059.00 ± 0%  +8314.18% (p=0.002 n=6)
RandomPfx4Size/100_000       59.57 ± 0%   3938.00 ± 0%  +6510.71% (p=0.002 n=6)
RandomPfx4Size/200_000       52.51 ± 0%   3476.00 ± 0%  +6519.69% (p=0.002 n=6)
RandomPfx4Size/500_000       37.84 ± 0%
RandomPfx4Size/1_000_000     31.83 ± 0%
RandomPfx6Size/1_000         82.54 ± 2%   6761.00 ± 0%  +8091.18% (p=0.002 n=6)
RandomPfx6Size/10_000        98.86 ± 0%   7333.00 ± 0%  +7317.56% (p=0.002 n=6)
RandomPfx6Size/100_000       67.59 ± 0%   5708.00 ± 0%  +8345.04% (p=0.002 n=6)
RandomPfx6Size/200_000       64.96 ± 0%   5526.00 ± 0%  +8406.77% (p=0.002 n=6)
RandomPfx6Size/500_000       67.62 ± 0%
RandomPfx6Size/1_000_000     73.74 ± 0%
RandomPfxSize/1_000          98.85 ± 2%   7537.00 ± 0%  +7524.68% (p=0.002 n=6)
RandomPfxSize/10_000         71.37 ± 0%   6058.00 ± 0%  +8388.16% (p=0.002 n=6)
RandomPfxSize/100_000        76.54 ± 0%   5300.00 ± 0%  +6824.48% (p=0.002 n=6)
RandomPfxSize/200_000        65.49 ± 0%   4586.00 ± 0%  +6902.60% (p=0.002 n=6)
RandomPfxSize/500_000        47.94 ± 0%
RandomPfxSize/1_000_000      39.50 ± 0%
geomean                      56.07        4.449Ki       +6734.49%               ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │           netipds/size.bm           │
                         │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000           103.9 ± 2%    101.0 ± 2%    -2.79% (p=0.002 n=6)
Tier1PfxSize/10_000          76.05 ± 0%    96.08 ± 0%   +26.34% (p=0.002 n=6)
Tier1PfxSize/100_000         41.16 ± 0%    94.34 ± 0%  +129.20% (p=0.002 n=6)
Tier1PfxSize/200_000         32.85 ± 0%    93.27 ± 0%  +183.93% (p=0.002 n=6)
Tier1PfxSize/500_000         25.19 ± 0%    91.05 ± 0%  +261.45% (p=0.002 n=6)
Tier1PfxSize/1_000_000       21.39 ± 0%    88.69 ± 0%  +314.63% (p=0.002 n=6)
RandomPfx4Size/1_000         74.51 ± 2%    90.32 ± 2%   +21.22% (p=0.002 n=6)
RandomPfx4Size/10_000        48.24 ± 0%    76.82 ± 0%   +59.25% (p=0.002 n=6)
RandomPfx4Size/100_000       59.57 ± 0%    72.35 ± 0%   +21.45% (p=0.002 n=6)
RandomPfx4Size/200_000       52.51 ± 0%    71.35 ± 0%   +35.88% (p=0.002 n=6)
RandomPfx4Size/500_000       37.84 ± 0%    69.77 ± 0%   +84.38% (p=0.002 n=6)
RandomPfx4Size/1_000_000     31.83 ± 0%    68.48 ± 0%  +115.14% (p=0.002 n=6)
RandomPfx6Size/1_000         82.54 ± 2%    99.92 ± 2%   +21.06% (p=0.002 n=6)
RandomPfx6Size/10_000        98.86 ± 0%    94.25 ± 0%    -4.66% (p=0.002 n=6)
RandomPfx6Size/100_000       67.59 ± 0%    87.64 ± 0%   +29.66% (p=0.002 n=6)
RandomPfx6Size/200_000       64.96 ± 0%    85.90 ± 0%   +32.24% (p=0.002 n=6)
RandomPfx6Size/500_000       67.62 ± 0%    84.11 ± 0%   +24.39% (p=0.002 n=6)
RandomPfx6Size/1_000_000     73.74 ± 0%    83.23 ± 0%   +12.87% (p=0.002 n=6)
RandomPfxSize/1_000          98.85 ± 2%    99.78 ± 2%    +0.94% (p=0.013 n=6)
RandomPfxSize/10_000         71.37 ± 0%    94.34 ± 0%   +32.18% (p=0.002 n=6)
RandomPfxSize/100_000        76.54 ± 0%    87.05 ± 0%   +13.73% (p=0.002 n=6)
RandomPfxSize/200_000        65.49 ± 0%    83.21 ± 0%   +27.06% (p=0.002 n=6)
RandomPfxSize/500_000        47.94 ± 0%    77.29 ± 0%   +61.22% (p=0.002 n=6)
RandomPfxSize/1_000_000      39.50 ± 0%    72.31 ± 0%   +83.06% (p=0.002 n=6)
geomean                      56.07         85.35        +52.21%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │          critbitgo/size.bm          │
                         │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000           103.9 ± 2%    119.5 ± 2%   +15.01% (p=0.002 n=6)
Tier1PfxSize/10_000          76.05 ± 0%   114.70 ± 0%   +50.82% (p=0.002 n=6)
Tier1PfxSize/100_000         41.16 ± 0%   114.40 ± 0%  +177.94% (p=0.002 n=6)
Tier1PfxSize/200_000         32.85 ± 0%   114.40 ± 0%  +248.25% (p=0.002 n=6)
Tier1PfxSize/500_000         25.19 ± 0%   114.40 ± 0%  +354.15% (p=0.002 n=6)
Tier1PfxSize/1_000_000       21.39 ± 0%   114.40 ± 0%  +434.83% (p=0.002 n=6)
RandomPfx4Size/1_000         74.51 ± 2%   116.10 ± 2%   +55.82% (p=0.002 n=6)
RandomPfx4Size/10_000        48.24 ± 0%   112.40 ± 0%  +133.00% (p=0.002 n=6)
RandomPfx4Size/100_000       59.57 ± 0%   112.00 ± 0%   +88.01% (p=0.002 n=6)
RandomPfx4Size/200_000       52.51 ± 0%   112.00 ± 0%  +113.29% (p=0.002 n=6)
RandomPfx4Size/500_000       37.84 ± 0%   112.00 ± 0%  +195.98% (p=0.002 n=6)
RandomPfx4Size/1_000_000     31.83 ± 0%   112.00 ± 0%  +251.87% (p=0.002 n=6)
RandomPfx6Size/1_000         82.54 ± 2%   132.10 ± 1%   +60.04% (p=0.002 n=6)
RandomPfx6Size/10_000        98.86 ± 0%   128.40 ± 0%   +29.88% (p=0.002 n=6)
RandomPfx6Size/100_000       67.59 ± 0%   128.00 ± 0%   +89.38% (p=0.002 n=6)
RandomPfx6Size/200_000       64.96 ± 0%   128.00 ± 0%   +97.04% (p=0.002 n=6)
RandomPfx6Size/500_000       67.62 ± 0%   128.00 ± 0%   +89.29% (p=0.002 n=6)
RandomPfx6Size/1_000_000     73.74 ± 0%   128.00 ± 0%   +73.58% (p=0.002 n=6)
RandomPfxSize/1_000          98.85 ± 2%   119.50 ± 2%   +20.89% (p=0.002 n=6)
RandomPfxSize/10_000         71.37 ± 0%   115.60 ± 0%   +61.97% (p=0.002 n=6)
RandomPfxSize/100_000        76.54 ± 0%   115.30 ± 0%   +50.64% (p=0.002 n=6)
RandomPfxSize/200_000        65.49 ± 0%   115.20 ± 0%   +75.90% (p=0.002 n=6)
RandomPfxSize/500_000        47.94 ± 0%   115.20 ± 0%  +140.30% (p=0.002 n=6)
RandomPfxSize/1_000_000      39.50 ± 0%   115.20 ± 0%  +191.65% (p=0.002 n=6)
geomean                      56.07         118.0       +110.50%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │           lpmtrie/size.bm            │
                         │ bytes/route  │ bytes/route   vs base                │
Tier1PfxSize/1_000           103.9 ± 2%    215.4 ±  5%  +107.31% (p=0.002 n=6)
Tier1PfxSize/10_000          76.05 ± 0%   210.50 ±  5%  +176.79% (p=0.002 n=6)
Tier1PfxSize/100_000         41.16 ± 0%   209.90 ±  5%  +409.96% (p=0.002 n=6)
Tier1PfxSize/200_000         32.85 ± 0%   209.20 ±  5%  +536.83% (p=0.002 n=6)
Tier1PfxSize/500_000         25.19 ± 0%   207.90 ±  7%  +725.33% (p=0.002 n=6)
Tier1PfxSize/1_000_000       21.39 ± 0%   205.00 ±  7%  +858.39% (p=0.002 n=6)
RandomPfx4Size/1_000         74.51 ± 2%   205.20 ±  8%  +175.40% (p=0.002 n=6)
RandomPfx4Size/10_000        48.24 ± 0%   186.60 ±  9%  +286.82% (p=0.002 n=6)
RandomPfx4Size/100_000       59.57 ± 0%   179.60 ±  9%  +201.49% (p=0.002 n=6)
RandomPfx4Size/200_000       52.51 ± 0%   178.50 ± 10%  +239.94% (p=0.002 n=6)
RandomPfx4Size/500_000       37.84 ± 0%   176.80 ± 10%  +367.23% (p=0.002 n=6)
RandomPfx4Size/1_000_000     31.83 ± 0%   175.50 ± 10%  +451.37% (p=0.002 n=6)
RandomPfx6Size/1_000         82.54 ± 2%   227.70 ±  8%  +175.87% (p=0.002 n=6)
RandomPfx6Size/10_000        98.86 ± 0%   222.50 ±  8%  +125.07% (p=0.002 n=6)
RandomPfx6Size/100_000       67.59 ± 0%   213.90 ±  9%  +216.47% (p=0.002 n=6)
RandomPfx6Size/200_000       64.96 ± 0%   210.50 ±  9%  +224.05% (p=0.002 n=6)
RandomPfx6Size/500_000       67.62 ± 0%   206.70 ±  9%  +205.68% (p=0.002 n=6)
RandomPfx6Size/1_000_000     73.74 ± 0%   204.90 ±  9%  +177.87% (p=0.002 n=6)
RandomPfxSize/1_000          98.85 ± 2%   215.00 ±  5%  +117.50% (p=0.002 n=6)
RandomPfxSize/10_000         71.37 ± 0%   210.30 ±  5%  +194.66% (p=0.002 n=6)
RandomPfxSize/100_000        76.54 ± 0%   203.50 ±  7%  +165.87% (p=0.002 n=6)
RandomPfxSize/200_000        65.49 ± 0%   198.70 ±  8%  +203.41% (p=0.002 n=6)
RandomPfxSize/500_000        47.94 ± 0%   189.90 ±  9%  +296.12% (p=0.002 n=6)
RandomPfxSize/1_000_000      39.50 ± 0%   181.40 ± 10%  +359.24% (p=0.002 n=6)
geomean                      56.07         201.3        +259.06%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │           cidranger/size.bm            │
                         │ bytes/route  │ bytes/route  vs base                   │
Tier1PfxSize/1_000           103.9 ± 2%    539.5 ± 3%   +419.25% (p=0.002 n=6)
Tier1PfxSize/10_000          76.05 ± 0%   533.90 ± 3%   +602.04% (p=0.002 n=6)
Tier1PfxSize/100_000         41.16 ± 0%   527.20 ± 2%  +1180.86% (p=0.002 n=6)
Tier1PfxSize/200_000         32.85 ± 0%   522.20 ± 2%  +1489.65% (p=0.002 n=6)
Tier1PfxSize/500_000         25.19 ± 0%
Tier1PfxSize/1_000_000       21.39 ± 0%
RandomPfx4Size/1_000         74.51 ± 2%   481.50 ± 3%   +546.22% (p=0.002 n=6)
RandomPfx4Size/10_000        48.24 ± 0%   433.10 ± 3%   +797.80% (p=0.002 n=6)
RandomPfx4Size/100_000       59.57 ± 0%   413.70 ± 3%   +594.48% (p=0.002 n=6)
RandomPfx4Size/200_000       52.51 ± 0%   409.10 ± 3%   +679.09% (p=0.002 n=6)
RandomPfx4Size/500_000       37.84 ± 0%
RandomPfx4Size/1_000_000     31.83 ± 0%
RandomPfx6Size/1_000         82.54 ± 2%   595.10 ± 0%   +620.98% (p=0.002 n=6)
RandomPfx6Size/10_000        98.86 ± 0%   581.00 ± 0%   +487.70% (p=0.002 n=6)
RandomPfx6Size/100_000       67.59 ± 0%   547.20 ± 0%   +709.59% (p=0.002 n=6)
RandomPfx6Size/200_000       64.96 ± 0%   538.10 ± 0%   +728.36% (p=0.002 n=6)
RandomPfx6Size/500_000       67.62 ± 0%
RandomPfx6Size/1_000_000     73.74 ± 0%
RandomPfxSize/1_000          98.85 ± 2%   540.30 ± 2%   +446.59% (p=0.002 n=6)
RandomPfxSize/10_000         71.37 ± 0%   528.10 ± 2%   +639.95% (p=0.002 n=6)
RandomPfxSize/100_000        76.54 ± 0%   495.50 ± 2%   +547.37% (p=0.002 n=6)
RandomPfxSize/200_000        65.49 ± 0%   477.60 ± 2%   +629.27% (p=0.002 n=6)
RandomPfxSize/500_000        47.94 ± 0%
RandomPfxSize/1_000_000      39.50 ± 0%
geomean                      56.07         507.3        +660.96%               ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │          cidrtree/size.bm           │
                         │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000          103.90 ± 2%    69.26 ± 3%   -33.34% (p=0.002 n=6)
Tier1PfxSize/10_000          76.05 ± 0%    64.37 ± 0%   -15.36% (p=0.002 n=6)
Tier1PfxSize/100_000         41.16 ± 0%    64.04 ± 0%   +55.59% (p=0.002 n=6)
Tier1PfxSize/200_000         32.85 ± 0%    64.02 ± 0%   +94.89% (p=0.002 n=6)
Tier1PfxSize/500_000         25.19 ± 0%    64.01 ± 0%  +154.11% (p=0.002 n=6)
Tier1PfxSize/1_000_000       21.39 ± 0%    64.00 ± 0%  +199.21% (p=0.002 n=6)
RandomPfx4Size/1_000         74.51 ± 2%    68.16 ± 3%    -8.52% (p=0.002 n=6)
RandomPfx4Size/10_000        48.24 ± 0%    64.37 ± 0%   +33.44% (p=0.002 n=6)
RandomPfx4Size/100_000       59.57 ± 0%    64.04 ± 0%    +7.50% (p=0.002 n=6)
RandomPfx4Size/200_000       52.51 ± 0%    64.02 ± 0%   +21.92% (p=0.002 n=6)
RandomPfx4Size/500_000       37.84 ± 0%    64.01 ± 0%   +69.16% (p=0.002 n=6)
RandomPfx4Size/1_000_000     31.83 ± 0%    64.00 ± 0%  +101.07% (p=0.002 n=6)
RandomPfx6Size/1_000         82.54 ± 2%    68.16 ± 3%   -17.42% (p=0.002 n=6)
RandomPfx6Size/10_000        98.86 ± 0%    64.37 ± 0%   -34.89% (p=0.002 n=6)
RandomPfx6Size/100_000       67.59 ± 0%    64.04 ± 0%    -5.25% (p=0.002 n=6)
RandomPfx6Size/200_000       64.96 ± 0%    64.02 ± 0%    -1.45% (p=0.002 n=6)
RandomPfx6Size/500_000       67.62 ± 0%    64.01 ± 0%    -5.34% (p=0.002 n=6)
RandomPfx6Size/1_000_000     73.74 ± 0%    64.00 ± 0%   -13.21% (p=0.002 n=6)
RandomPfxSize/1_000          98.85 ± 2%    68.16 ± 3%   -31.05% (p=0.002 n=6)
RandomPfxSize/10_000         71.37 ± 0%    64.37 ± 0%    -9.81% (p=0.002 n=6)
RandomPfxSize/100_000        76.54 ± 0%    64.04 ± 0%   -16.33% (p=0.002 n=6)
RandomPfxSize/200_000        65.49 ± 0%    64.02 ± 0%    -2.24% (p=0.002 n=6)
RandomPfxSize/500_000        47.94 ± 0%    64.01 ± 0%   +33.52% (p=0.002 n=6)
RandomPfxSize/1_000_000      39.50 ± 0%    64.00 ± 0%   +62.03% (p=0.002 n=6)
geomean                      56.07         64.79        +15.55%
```

## update, insert/delete

`bart` is by far the fastest algorithm for updates.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │             art/update.bm              │
                         │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000        30.13n ± 3%   300.35n ±  3%   +896.68% (p=0.002 n=6)
InsertRandomPfxs/10_000       36.26n ± 1%   527.65n ±  6%  +1355.18% (p=0.002 n=6)
InsertRandomPfxs/100_000      100.9n ± 2%   2677.5n ± 21%  +2552.30% (p=0.002 n=6)
InsertRandomPfxs/200_000      155.3n ± 1%   1811.0n ± 55%  +1065.75% (p=0.002 n=6)
DeleteRandomPfxs/1_000        16.74n ± 0%    16.89n ±  1%     +0.93% (p=0.006 n=6)
DeleteRandomPfxs/10_000       16.79n ± 1%    16.65n ±  0%     -0.77% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.23n ± 0%    17.57n ±  3%     +1.94% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.33n ± 1%    19.68n ±  1%     +7.34% (p=0.002 n=6)
geomean                       33.33n         128.6n         +285.86%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           netipds/update.bm           │
                         │   sec/route    │  sec/route    vs base                 │
InsertRandomPfxs/1_000        30.13n ± 3%   131.30n ± 4%   +335.71% (p=0.002 n=6)
InsertRandomPfxs/10_000       36.26n ± 1%   190.50n ± 1%   +425.37% (p=0.002 n=6)
InsertRandomPfxs/100_000      100.9n ± 2%    392.8n ± 2%   +289.10% (p=0.002 n=6)
InsertRandomPfxs/200_000      155.3n ± 1%    585.4n ± 1%   +276.83% (p=0.002 n=6)
DeleteRandomPfxs/1_000        16.74n ± 0%   111.30n ± 6%   +564.87% (p=0.002 n=6)
DeleteRandomPfxs/10_000       16.79n ± 1%   166.80n ± 2%   +893.74% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.23n ± 0%   361.00n ± 2%  +1994.57% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.33n ± 1%   552.70n ± 0%  +2915.28% (p=0.002 n=6)
geomean                       33.33n         260.7n        +682.15%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │         critbitgo/update.bm          │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000        30.13n ± 3%   168.90n ± 4%  +460.48% (p=0.002 n=6)
InsertRandomPfxs/10_000       36.26n ± 1%   239.30n ± 2%  +559.96% (p=0.002 n=6)
InsertRandomPfxs/100_000      100.9n ± 2%    600.0n ± 2%  +494.30% (p=0.002 n=6)
InsertRandomPfxs/200_000      155.3n ± 1%    791.1n ± 2%  +409.27% (p=0.002 n=6)
DeleteRandomPfxs/1_000        16.74n ± 0%    71.32n ± 3%  +326.05% (p=0.002 n=6)
DeleteRandomPfxs/10_000       16.79n ± 1%    73.06n ± 2%  +335.24% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.23n ± 0%    78.97n ± 1%  +358.20% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.33n ± 1%    85.51n ± 8%  +366.50% (p=0.002 n=6)
geomean                       33.33n         169.3n       +408.03%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           lpmtrie/update.bm            │
                         │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000        30.13n ± 3%   366.15n ±  5%  +1115.03% (p=0.002 n=6)
InsertRandomPfxs/10_000       36.26n ± 1%   493.90n ±  2%  +1262.11% (p=0.002 n=6)
InsertRandomPfxs/100_000      100.9n ± 2%   1173.5n ±  4%  +1062.46% (p=0.002 n=6)
InsertRandomPfxs/200_000      155.3n ± 1%   1229.0n ± 15%   +691.12% (p=0.002 n=6)
DeleteRandomPfxs/1_000        16.74n ± 0%    71.66n ±  1%   +328.11% (p=0.002 n=6)
DeleteRandomPfxs/10_000       16.79n ± 1%   132.85n ±  4%   +691.48% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.23n ± 0%   307.60n ±  6%  +1684.74% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.33n ± 1%   485.15n ±  3%  +2546.75% (p=0.002 n=6)
geomean                       33.33n         372.5n        +1017.73%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           cidranger/update.bm           │
                         │   sec/route    │   sec/route    vs base                  │
InsertRandomPfxs/1_000        30.13n ± 3%   4535.00n ± 3%  +14948.95% (p=0.002 n=6)
InsertRandomPfxs/10_000       36.26n ± 1%   6460.50n ± 4%  +17717.15% (p=0.002 n=6)
InsertRandomPfxs/100_000      100.9n ± 2%   10837.5n ± 3%  +10635.51% (p=0.002 n=6)
InsertRandomPfxs/200_000      155.3n ± 1%   11409.0n ± 4%   +7244.06% (p=0.002 n=6)
DeleteRandomPfxs/1_000        16.74n ± 0%     87.85n ± 3%    +424.82% (p=0.002 n=6)
DeleteRandomPfxs/10_000       16.79n ± 1%     90.95n ± 1%    +441.88% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.23n ± 0%    121.55n ± 2%    +605.25% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.33n ± 1%    237.75n ± 6%   +1197.05% (p=0.002 n=6)
geomean                       33.33n          977.9n        +2834.54%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           cidrtree/update.bm            │
                         │   sec/route    │   sec/route     vs base                 │
InsertRandomPfxs/1_000        30.13n ± 3%   1235.50n ±  3%  +3999.88% (p=0.002 n=6)
InsertRandomPfxs/10_000       36.26n ± 1%   1841.00n ±  2%  +4977.22% (p=0.002 n=6)
InsertRandomPfxs/100_000      100.9n ± 2%    3221.0n ±  1%  +3090.69% (p=0.002 n=6)
InsertRandomPfxs/200_000      155.3n ± 1%    3992.0n ±  6%  +2469.68% (p=0.002 n=6)
DeleteRandomPfxs/1_000        16.74n ± 0%     15.19n ±  7%     -9.26% (p=0.002 n=6)
DeleteRandomPfxs/10_000       16.79n ± 1%     15.03n ±  1%    -10.49% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.23n ± 0%     23.64n ± 11%    +37.16% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.33n ± 1%    289.40n ± 33%  +1478.83% (p=0.002 n=6)
geomean                       33.33n          286.7n         +760.36%
```
