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

`bart` is the fastest software algorithm for IP-address longest-prefix-match.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │              art/lpm.bm               │
                                     │    sec/op    │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4            18.60n ± 32%   45.64n ± 16%  +145.38% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            38.49n ± 18%   62.52n ±  4%   +62.41% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             20.51n ±  1%   28.98n ±  0%   +41.33% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             35.83n ± 25%   36.27n ± 22%         ~ (p=0.251 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     18.48n ±  4%   44.97n ±  1%  +143.34% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     21.85n ±  0%   46.17n ±  2%  +111.33% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      18.46n ±  2%   28.76n ±  0%   +55.80% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      21.89n ±  0%   29.00n ±  7%   +32.48% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    19.08n ±  1%   45.10n ±  0%  +136.35% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    21.93n ±  1%   46.07n ±  6%  +110.13% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     19.11n ±  3%   28.77n ±  1%   +50.52% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     21.96n ± 36%   28.89n ±  6%   +31.56% (p=0.002 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   15.74n ± 42%   49.25n ± 10%  +212.80% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   29.90n ±  2%   48.80n ±  7%   +63.24% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.96n ± 42%   28.87n ±  2%   +52.27% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    29.86n ±  2%   33.01n ±  6%   +10.55% (p=0.000 n=20)
geomean                                22.40n         38.20n         +70.53%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            netipds/lpm.bm             │
                                     │    sec/op    │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4            18.60n ± 32%   60.01n ± 12%  +222.63% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            38.49n ± 18%   47.14n ± 18%   +22.44% (p=0.018 n=20)
LpmTier1Pfxs/RandomMissIP4             20.51n ±  1%   64.91n ±  6%  +216.56% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             35.83n ± 25%   52.71n ± 11%   +47.12% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     18.48n ±  4%   39.52n ±  5%  +113.83% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     21.85n ±  0%   32.62n ±  9%   +49.29% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      18.46n ±  2%   40.16n ±  3%  +117.55% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      21.89n ±  0%   32.46n ±  6%   +48.31% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    19.08n ±  1%   45.75n ±  4%  +139.78% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    21.93n ±  1%   37.01n ±  3%   +68.80% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     19.11n ±  3%   49.38n ±  4%  +158.42% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     21.96n ± 36%   38.69n ±  3%   +76.16% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   15.74n ± 42%   47.85n ±  9%  +203.91% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   29.90n ±  2%   42.10n ±  2%   +40.81% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.96n ± 42%   53.68n ±  4%  +183.15% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    29.86n ±  2%   44.21n ±  3%   +48.07% (p=0.000 n=20)
geomean                                22.40n         44.67n         +99.44%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            critbitgo/lpm.bm             │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            18.60n ± 32%   162.65n ±  6%   +774.46% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            38.49n ± 18%   250.05n ± 14%   +549.56% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             20.51n ±  1%   630.25n ± 18%  +2973.64% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             35.83n ± 25%   485.70n ± 12%  +1255.76% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     18.48n ±  4%    83.98n ±  5%   +354.46% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     21.85n ±  0%   105.30n ±  3%   +381.92% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      18.46n ±  2%   149.45n ± 13%   +709.59% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      21.89n ±  0%   159.20n ±  8%   +627.27% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    19.08n ±  1%    98.94n ±  5%   +418.55% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    21.93n ±  1%   117.45n ±  3%   +435.69% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     19.11n ±  3%   254.40n ± 10%  +1231.24% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     21.96n ± 36%   221.05n ± 15%   +906.60% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   15.74n ± 42%   110.25n ±  5%   +600.22% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   29.90n ±  2%   132.15n ±  2%   +342.05% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.96n ± 42%   472.55n ± 16%  +2392.35% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    29.86n ±  2%   332.55n ± 18%  +1013.70% (p=0.000 n=20)
geomean                                22.40n          193.7n         +764.66%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             lpmtrie/lpm.bm              │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            18.60n ± 32%   279.85n ± 10%  +1404.57% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            38.49n ± 18%   323.35n ± 17%   +739.98% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             20.51n ±  1%   260.05n ±  5%  +1168.23% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             35.83n ± 25%   234.05n ± 15%   +553.31% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     18.48n ±  4%   135.45n ±  8%   +632.95% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     21.85n ±  0%   120.30n ± 23%   +450.57% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      18.46n ±  2%   115.15n ± 17%   +523.78% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      21.89n ±  0%    95.28n ± 37%   +335.27% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    19.08n ±  1%   153.75n ± 13%   +705.82% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    21.93n ±  1%   140.00n ±  9%   +538.54% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     19.11n ±  3%   163.00n ± 11%   +752.96% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     21.96n ± 36%   135.40n ±  5%   +516.58% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   15.74n ± 42%   197.85n ±  5%  +1156.59% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   29.90n ±  2%   186.65n ±  2%   +524.35% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.96n ± 42%   199.00n ±  4%   +949.58% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    29.86n ±  2%   189.40n ±  5%   +534.29% (p=0.000 n=20)
geomean                                22.40n          173.1n         +672.90%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            cidranger/lpm.bm             │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            18.60n ± 32%   248.70n ± 23%  +1237.10% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            38.49n ± 18%   276.80n ± 29%   +619.05% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             20.51n ±  1%   281.30n ± 10%  +1271.86% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             35.83n ± 25%   290.95n ± 17%   +712.14% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     18.48n ±  4%   119.00n ± 12%   +543.94% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     21.85n ±  0%   155.20n ± 11%   +610.30% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      18.46n ±  2%   113.55n ±  5%   +515.11% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      21.89n ±  0%   149.85n ± 10%   +584.56% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    19.08n ±  1%   157.00n ±  8%   +722.85% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    21.93n ±  1%   181.10n ±  6%   +726.00% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     19.11n ±  3%   154.55n ±  4%   +708.74% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     21.96n ± 36%   178.65n ±  5%   +713.52% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   15.74n ± 42%   168.90n ± 20%   +972.72% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   29.90n ±  2%   226.85n ±  3%   +658.82% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.96n ± 42%   188.65n ±  4%   +894.99% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    29.86n ±  2%   226.55n ±  7%   +658.71% (p=0.000 n=20)
geomean                                22.40n          187.2n         +735.79%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             cidrtree/lpm.bm              │
                                     │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            18.60n ± 32%    701.40n ± 12%  +3670.97% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            38.49n ± 18%    802.75n ± 19%  +1985.34% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             20.51n ±  1%    980.15n ± 13%  +4680.05% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             35.83n ± 25%   1099.00n ± 30%  +2967.69% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     18.48n ±  4%    336.60n ± 12%  +1721.43% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     21.85n ±  0%    396.50n ± 30%  +1714.65% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      18.46n ±  2%    380.40n ± 17%  +1960.67% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      21.89n ±  0%    489.20n ± 29%  +2134.81% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    19.08n ±  1%    487.45n ± 15%  +2454.77% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    21.93n ±  1%    520.20n ± 11%  +2272.63% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     19.11n ±  3%    637.55n ± 18%  +3236.21% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     21.96n ± 36%    795.75n ± 22%  +3523.63% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   15.74n ± 42%    563.65n ± 12%  +3479.87% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   29.90n ±  2%    693.60n ± 16%  +2220.12% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.96n ± 42%    879.60n ± 19%  +4539.24% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    29.86n ±  2%   1085.00n ± 11%  +3533.62% (p=0.000 n=20)
geomean                                22.40n           636.4n        +2740.97%
```

## size of the routing tables


`bart` has two orders of magnitude lower memory consumption compared to `art`
and is in par to the binary search tree `cidrtree`, but orders of magintude faster.


```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │              art/size.bm              │
                       │ bytes/route  │ bytes/route   vs base                 │
Tier1PfxSize/1_000         102.3 ± 2%    7591.0 ± 0%  +7320.33% (p=0.002 n=6)
Tier1PfxSize/10_000        88.50 ± 0%   4889.00 ± 0%  +5424.29% (p=0.002 n=6)
Tier1PfxSize/100_000       66.05 ± 0%   1669.00 ± 0%  +2426.87% (p=0.002 n=6)
Tier1PfxSize/200_000       59.36 ± 0%   1098.00 ± 0%  +1749.73% (p=0.002 n=6)
RandomPfx4Size/1_000       73.02 ± 3%   5259.00 ± 0%  +7102.14% (p=0.002 n=6)
RandomPfx4Size/10_000      49.60 ± 0%   4059.00 ± 0%  +8083.47% (p=0.002 n=6)
RandomPfx4Size/100_000     58.99 ± 0%   3938.00 ± 0%  +6575.71% (p=0.002 n=6)
RandomPfx4Size/200_000     51.84 ± 0%   3476.00 ± 0%  +6605.25% (p=0.002 n=6)
RandomPfx6Size/1_000       82.90 ± 2%   6761.00 ± 0%  +8055.61% (p=0.002 n=6)
RandomPfx6Size/10_000      95.44 ± 0%   7333.00 ± 0%  +7583.36% (p=0.002 n=6)
RandomPfx6Size/100_000     67.20 ± 0%   5708.00 ± 0%  +8394.05% (p=0.002 n=6)
RandomPfx6Size/200_000     64.99 ± 0%   5526.00 ± 0%  +8402.85% (p=0.002 n=6)
RandomPfxSize/1_000        96.22 ± 2%   7537.00 ± 0%  +7733.09% (p=0.002 n=6)
RandomPfxSize/10_000       72.19 ± 0%   6058.00 ± 0%  +8291.74% (p=0.002 n=6)
RandomPfxSize/100_000      75.13 ± 0%   5300.00 ± 0%  +6954.44% (p=0.002 n=6)
RandomPfxSize/200_000      64.40 ± 0%   4586.00 ± 0%  +7021.12% (p=0.002 n=6)
geomean                    71.40        4.449Ki       +6280.97%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          netipds/size.bm           │
                       │ bytes/route  │ bytes/route  vs base               │
Tier1PfxSize/1_000         102.3 ± 2%    101.0 ± 2%   -1.27% (p=0.013 n=6)
Tier1PfxSize/10_000        88.50 ± 0%    96.08 ± 0%   +8.56% (p=0.002 n=6)
Tier1PfxSize/100_000       66.05 ± 0%    94.34 ± 0%  +42.83% (p=0.002 n=6)
Tier1PfxSize/200_000       59.36 ± 0%    93.27 ± 0%  +57.13% (p=0.002 n=6)
RandomPfx4Size/1_000       73.02 ± 3%    90.32 ± 2%  +23.69% (p=0.002 n=6)
RandomPfx4Size/10_000      49.60 ± 0%    76.82 ± 0%  +54.88% (p=0.002 n=6)
RandomPfx4Size/100_000     58.99 ± 0%    72.35 ± 0%  +22.65% (p=0.002 n=6)
RandomPfx4Size/200_000     51.84 ± 0%    71.35 ± 0%  +37.64% (p=0.002 n=6)
RandomPfx6Size/1_000       82.90 ± 2%   100.20 ± 2%  +20.87% (p=0.002 n=6)
RandomPfx6Size/10_000      95.44 ± 0%    94.25 ± 0%   -1.25% (p=0.002 n=6)
RandomPfx6Size/100_000     67.20 ± 0%    87.63 ± 0%  +30.40% (p=0.002 n=6)
RandomPfx6Size/200_000     64.99 ± 0%    85.90 ± 0%  +32.17% (p=0.002 n=6)
RandomPfxSize/1_000        96.22 ± 2%    99.78 ± 2%   +3.70% (p=0.002 n=6)
RandomPfxSize/10_000       72.19 ± 0%    94.34 ± 0%  +30.68% (p=0.002 n=6)
RandomPfxSize/100_000      75.13 ± 0%    87.05 ± 0%  +15.87% (p=0.002 n=6)
RandomPfxSize/200_000      64.40 ± 0%    83.21 ± 0%  +29.21% (p=0.002 n=6)
geomean                    71.40         88.75       +24.30%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          critbitgo/size.bm          │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         102.3 ± 2%    119.5 ± 2%   +16.81% (p=0.002 n=6)
Tier1PfxSize/10_000        88.50 ± 0%   114.70 ± 0%   +29.60% (p=0.002 n=6)
Tier1PfxSize/100_000       66.05 ± 0%   114.40 ± 0%   +73.20% (p=0.002 n=6)
Tier1PfxSize/200_000       59.36 ± 0%   114.40 ± 0%   +92.72% (p=0.002 n=6)
RandomPfx4Size/1_000       73.02 ± 3%   116.10 ± 2%   +59.00% (p=0.002 n=6)
RandomPfx4Size/10_000      49.60 ± 0%   112.40 ± 0%  +126.61% (p=0.002 n=6)
RandomPfx4Size/100_000     58.99 ± 0%   112.00 ± 0%   +89.86% (p=0.002 n=6)
RandomPfx4Size/200_000     51.84 ± 0%   112.00 ± 0%  +116.05% (p=0.002 n=6)
RandomPfx6Size/1_000       82.90 ± 2%   132.40 ± 1%   +59.71% (p=0.002 n=6)
RandomPfx6Size/10_000      95.44 ± 0%   128.40 ± 0%   +34.53% (p=0.002 n=6)
RandomPfx6Size/100_000     67.20 ± 0%   128.00 ± 0%   +90.48% (p=0.002 n=6)
RandomPfx6Size/200_000     64.99 ± 0%   128.00 ± 0%   +96.95% (p=0.002 n=6)
RandomPfxSize/1_000        96.22 ± 2%   119.50 ± 2%   +24.19% (p=0.002 n=6)
RandomPfxSize/10_000       72.19 ± 0%   115.60 ± 0%   +60.13% (p=0.002 n=6)
RandomPfxSize/100_000      75.13 ± 0%   115.30 ± 0%   +53.47% (p=0.002 n=6)
RandomPfxSize/200_000      64.40 ± 0%   115.20 ± 0%   +78.88% (p=0.002 n=6)
geomean                    71.40         118.4        +65.88%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │           lpmtrie/size.bm            │
                       │ bytes/route  │ bytes/route   vs base                │
Tier1PfxSize/1_000         102.3 ± 2%    215.4 ±  5%  +110.56% (p=0.002 n=6)
Tier1PfxSize/10_000        88.50 ± 0%   210.50 ±  5%  +137.85% (p=0.002 n=6)
Tier1PfxSize/100_000       66.05 ± 0%   209.90 ±  5%  +217.79% (p=0.002 n=6)
Tier1PfxSize/200_000       59.36 ± 0%   209.20 ±  5%  +252.43% (p=0.002 n=6)
RandomPfx4Size/1_000       73.02 ± 3%   205.20 ±  8%  +181.02% (p=0.002 n=6)
RandomPfx4Size/10_000      49.60 ± 0%   186.60 ±  9%  +276.21% (p=0.002 n=6)
RandomPfx4Size/100_000     58.99 ± 0%   179.60 ±  9%  +204.46% (p=0.002 n=6)
RandomPfx4Size/200_000     51.84 ± 0%   178.50 ± 10%  +244.33% (p=0.002 n=6)
RandomPfx6Size/1_000       82.90 ± 2%   228.00 ±  8%  +175.03% (p=0.002 n=6)
RandomPfx6Size/10_000      95.44 ± 0%   222.50 ±  8%  +133.13% (p=0.002 n=6)
RandomPfx6Size/100_000     67.20 ± 0%   213.90 ±  9%  +218.30% (p=0.002 n=6)
RandomPfx6Size/200_000     64.99 ± 0%   210.50 ±  9%  +223.90% (p=0.002 n=6)
RandomPfxSize/1_000        96.22 ± 2%   215.00 ±  5%  +123.45% (p=0.002 n=6)
RandomPfxSize/10_000       72.19 ± 0%   210.30 ±  5%  +191.31% (p=0.002 n=6)
RandomPfxSize/100_000      75.13 ± 0%   203.50 ±  7%  +170.86% (p=0.002 n=6)
RandomPfxSize/200_000      64.40 ± 0%   198.70 ±  8%  +208.54% (p=0.002 n=6)
geomean                    71.40         205.6        +187.96%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidranger/size.bm          │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         102.3 ± 2%    539.6 ± 3%  +427.47% (p=0.002 n=6)
Tier1PfxSize/10_000        88.50 ± 0%   533.90 ± 3%  +503.28% (p=0.002 n=6)
Tier1PfxSize/100_000       66.05 ± 0%   527.20 ± 2%  +698.18% (p=0.002 n=6)
Tier1PfxSize/200_000       59.36 ± 0%   522.20 ± 2%  +779.72% (p=0.002 n=6)
RandomPfx4Size/1_000       73.02 ± 3%   481.50 ± 3%  +559.41% (p=0.002 n=6)
RandomPfx4Size/10_000      49.60 ± 0%   433.10 ± 3%  +773.19% (p=0.002 n=6)
RandomPfx4Size/100_000     58.99 ± 0%   413.70 ± 3%  +601.31% (p=0.002 n=6)
RandomPfx4Size/200_000     51.84 ± 0%   409.10 ± 3%  +689.16% (p=0.002 n=6)
RandomPfx6Size/1_000       82.90 ± 2%   595.10 ± 0%  +617.85% (p=0.002 n=6)
RandomPfx6Size/10_000      95.44 ± 0%   581.00 ± 0%  +508.76% (p=0.002 n=6)
RandomPfx6Size/100_000     67.20 ± 0%   547.20 ± 0%  +714.29% (p=0.002 n=6)
RandomPfx6Size/200_000     64.99 ± 0%   538.10 ± 0%  +727.97% (p=0.002 n=6)
RandomPfxSize/1_000        96.22 ± 2%   540.30 ± 2%  +461.53% (p=0.002 n=6)
RandomPfxSize/10_000       72.19 ± 0%   528.10 ± 2%  +631.54% (p=0.002 n=6)
RandomPfxSize/100_000      75.13 ± 0%   495.50 ± 2%  +559.52% (p=0.002 n=6)
RandomPfxSize/200_000      64.40 ± 0%   477.60 ± 2%  +641.61% (p=0.002 n=6)
geomean                    71.40         507.3       +610.48%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidrtree/size.bm          │
                       │ bytes/route  │ bytes/route  vs base               │
Tier1PfxSize/1_000        102.30 ± 2%    69.26 ± 3%  -32.30% (p=0.002 n=6)
Tier1PfxSize/10_000        88.50 ± 0%    64.37 ± 0%  -27.27% (p=0.002 n=6)
Tier1PfxSize/100_000       66.05 ± 0%    64.04 ± 0%   -3.04% (p=0.002 n=6)
Tier1PfxSize/200_000       59.36 ± 0%    64.02 ± 0%   +7.85% (p=0.002 n=6)
RandomPfx4Size/1_000       73.02 ± 3%    68.16 ± 3%   -6.66% (p=0.002 n=6)
RandomPfx4Size/10_000      49.60 ± 0%    64.37 ± 0%  +29.78% (p=0.002 n=6)
RandomPfx4Size/100_000     58.99 ± 0%    64.04 ± 0%   +8.56% (p=0.002 n=6)
RandomPfx4Size/200_000     51.84 ± 0%    64.02 ± 0%  +23.50% (p=0.002 n=6)
RandomPfx6Size/1_000       82.90 ± 2%    68.41 ± 3%  -17.48% (p=0.002 n=6)
RandomPfx6Size/10_000      95.44 ± 0%    64.37 ± 0%  -32.55% (p=0.002 n=6)
RandomPfx6Size/100_000     67.20 ± 0%    64.04 ± 0%   -4.70% (p=0.002 n=6)
RandomPfx6Size/200_000     64.99 ± 0%    64.02 ± 0%   -1.49% (p=0.002 n=6)
RandomPfxSize/1_000        96.22 ± 2%    68.16 ± 3%  -29.16% (p=0.002 n=6)
RandomPfxSize/10_000       72.19 ± 0%    64.37 ± 0%  -10.83% (p=0.002 n=6)
RandomPfxSize/100_000      75.13 ± 0%    64.04 ± 0%  -14.76% (p=0.002 n=6)
RandomPfxSize/200_000      64.40 ± 0%    64.02 ± 0%   -0.59% (p=0.002 n=6)
geomean                    71.40         65.20        -8.68%
```

## update, insert/delete

`bart` is by far the fastest algorithm for updates.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │              art/update.bm              │
                         │   sec/route    │   sec/route     vs base                 │
InsertRandomPfxs/1_000        26.31n ± 1%    300.55n ±  3%  +1042.12% (p=0.002 n=6)
InsertRandomPfxs/10_000       31.22n ± 2%    525.55n ±  4%  +1583.38% (p=0.002 n=6)
InsertRandomPfxs/100_000      81.93n ± 4%   1886.50n ± 34%  +2202.43% (p=0.002 n=6)
InsertRandomPfxs/200_000      131.2n ± 3%    1759.0n ± 52%  +1240.70% (p=0.002 n=6)
DeleteRandomPfxs/1_000        15.58n ± 0%     16.86n ±  1%     +8.22% (p=0.002 n=6)
DeleteRandomPfxs/10_000       15.60n ± 0%     16.65n ±  0%     +6.80% (p=0.002 n=6)
DeleteRandomPfxs/100_000      15.97n ± 0%     17.48n ±  1%     +9.49% (p=0.002 n=6)
DeleteRandomPfxs/200_000      16.64n ± 1%     18.94n ±  3%    +13.76% (p=0.002 n=6)
geomean                       29.48n          121.9n         +313.48%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           netipds/update.bm           │
                         │   sec/route    │  sec/route    vs base                 │
InsertRandomPfxs/1_000        26.31n ± 1%   130.60n ± 1%   +396.29% (p=0.002 n=6)
InsertRandomPfxs/10_000       31.22n ± 2%   191.60n ± 1%   +513.71% (p=0.002 n=6)
InsertRandomPfxs/100_000      81.93n ± 4%   380.85n ± 2%   +364.82% (p=0.002 n=6)
InsertRandomPfxs/200_000      131.2n ± 3%    552.5n ± 1%   +321.11% (p=0.002 n=6)
DeleteRandomPfxs/1_000        15.58n ± 0%   110.60n ± 1%   +609.88% (p=0.002 n=6)
DeleteRandomPfxs/10_000       15.60n ± 0%   168.55n ± 1%   +980.80% (p=0.002 n=6)
DeleteRandomPfxs/100_000      15.97n ± 0%   347.70n ± 2%  +2077.21% (p=0.002 n=6)
DeleteRandomPfxs/200_000      16.64n ± 1%   514.80n ± 1%  +2992.82% (p=0.002 n=6)
geomean                       29.48n         254.4n        +763.12%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │         critbitgo/update.bm          │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000        26.31n ± 1%   168.55n ± 5%  +540.51% (p=0.002 n=6)
InsertRandomPfxs/10_000       31.22n ± 2%   237.45n ± 2%  +660.57% (p=0.002 n=6)
InsertRandomPfxs/100_000      81.93n ± 4%   545.70n ± 1%  +566.02% (p=0.002 n=6)
InsertRandomPfxs/200_000      131.2n ± 3%    745.6n ± 2%  +468.33% (p=0.002 n=6)
DeleteRandomPfxs/1_000        15.58n ± 0%    71.11n ± 3%  +356.42% (p=0.002 n=6)
DeleteRandomPfxs/10_000       15.60n ± 0%    72.94n ± 1%  +367.75% (p=0.002 n=6)
DeleteRandomPfxs/100_000      15.97n ± 0%    76.67n ± 0%  +380.12% (p=0.002 n=6)
DeleteRandomPfxs/200_000      16.64n ± 1%    82.55n ± 4%  +395.94% (p=0.002 n=6)
geomean                       29.48n         164.4n       +457.81%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │            lpmtrie/update.bm            │
                         │   sec/route    │   sec/route     vs base                 │
InsertRandomPfxs/1_000        26.31n ± 1%    359.35n ±  5%  +1265.57% (p=0.002 n=6)
InsertRandomPfxs/10_000       31.22n ± 2%    474.85n ±  4%  +1420.98% (p=0.002 n=6)
InsertRandomPfxs/100_000      81.93n ± 4%   1088.50n ±  4%  +1228.49% (p=0.002 n=6)
InsertRandomPfxs/200_000      131.2n ± 3%    1385.5n ± 14%   +956.02% (p=0.002 n=6)
DeleteRandomPfxs/1_000        15.58n ± 0%     71.07n ±  1%   +356.16% (p=0.002 n=6)
DeleteRandomPfxs/10_000       15.60n ± 0%    131.00n ±  0%   +740.01% (p=0.002 n=6)
DeleteRandomPfxs/100_000      15.97n ± 0%    303.30n ±  1%  +1799.19% (p=0.002 n=6)
DeleteRandomPfxs/200_000      16.64n ± 1%    462.25n ±  3%  +2677.11% (p=0.002 n=6)
geomean                       29.48n          367.9n        +1148.11%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           cidranger/update.bm            │
                         │   sec/route    │   sec/route     vs base                  │
InsertRandomPfxs/1_000        26.31n ± 1%   4615.50n ±  4%  +17439.43% (p=0.002 n=6)
InsertRandomPfxs/10_000       31.22n ± 2%   6324.50n ±  3%  +20157.85% (p=0.002 n=6)
InsertRandomPfxs/100_000      81.93n ± 4%   9944.50n ±  5%  +12037.06% (p=0.002 n=6)
InsertRandomPfxs/200_000      131.2n ± 3%   10488.0n ± 10%   +7893.90% (p=0.002 n=6)
DeleteRandomPfxs/1_000        15.58n ± 0%     88.48n ±  2%    +467.91% (p=0.002 n=6)
DeleteRandomPfxs/10_000       15.60n ± 0%     92.86n ±  3%    +495.48% (p=0.002 n=6)
DeleteRandomPfxs/100_000      15.97n ± 0%    119.10n ±  4%    +645.77% (p=0.002 n=6)
DeleteRandomPfxs/200_000      16.64n ± 1%    228.65n ±  5%   +1273.69% (p=0.002 n=6)
geomean                       29.48n          953.2n         +3133.25%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           cidrtree/update.bm            │
                         │   sec/route    │   sec/route     vs base                 │
InsertRandomPfxs/1_000        26.31n ± 1%   1241.50n ±  3%  +4617.84% (p=0.002 n=6)
InsertRandomPfxs/10_000       31.22n ± 2%   1837.00n ±  1%  +5784.05% (p=0.002 n=6)
InsertRandomPfxs/100_000      81.93n ± 4%   3048.50n ±  5%  +3620.63% (p=0.002 n=6)
InsertRandomPfxs/200_000      131.2n ± 3%    3755.0n ±  2%  +2762.04% (p=0.002 n=6)
DeleteRandomPfxs/1_000        15.58n ± 0%     14.74n ±  2%     -5.36% (p=0.002 n=6)
DeleteRandomPfxs/10_000       15.60n ± 0%     15.01n ±  0%     -3.78% (p=0.002 n=6)
DeleteRandomPfxs/100_000      15.97n ± 0%     22.48n ±  6%    +40.76% (p=0.002 n=6)
DeleteRandomPfxs/200_000      16.64n ± 1%    136.60n ± 19%   +720.67% (p=0.002 n=6)
geomean                       29.48n          254.7n         +764.15%
```
