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
LpmTier1Pfxs/RandomMatchIP4            17.21n ±  6%   45.64n ± 16%  +165.12% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            35.12n ± 25%   62.52n ±  4%   +78.02% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             19.51n ±  0%   28.98n ±  0%   +48.54% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             30.09n ± 37%   36.27n ± 22%   +20.54% (p=0.003 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.97n ±  7%   44.97n ±  1%  +165.00% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.27n ±  0%   46.17n ±  2%  +139.56% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.71n ±  1%   28.76n ±  0%   +72.16% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.27n ±  0%   29.00n ±  7%   +50.49% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.24n ±  1%   45.10n ±  0%  +147.23% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.62n ±  9%   46.07n ±  6%  +134.81% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.21n ±  2%   28.77n ±  1%   +57.92% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.44n ± 33%   28.89n ±  6%   +48.61% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   14.68n ± 63%   49.25n ± 10%  +235.49% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   26.01n ±  1%   48.80n ±  7%   +87.62% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.47n ±  9%   28.87n ±  2%   +56.27% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    26.11n ± 10%   33.01n ±  6%   +26.43% (p=0.000 n=20)
geomean                                20.36n         38.20n         +87.61%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            netipds/lpm.bm             │
                                     │    sec/op    │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4            17.21n ±  6%   51.57n ± 19%  +199.56% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            35.12n ± 25%   53.54n ± 21%   +52.43% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             19.51n ±  0%   55.87n ±  4%  +186.34% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             30.09n ± 37%   54.15n ± 12%   +79.99% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.97n ±  7%   30.74n ±  5%   +81.14% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.27n ±  0%   32.32n ±  8%   +67.68% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.71n ±  1%   32.47n ±  6%   +94.37% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.27n ±  0%   33.64n ±  7%   +74.57% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.24n ±  1%   37.56n ±  5%  +105.89% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.62n ±  9%   36.59n ± 12%   +86.47% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.21n ±  2%   37.80n ±  5%  +107.49% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.44n ± 33%   38.01n ±  4%   +95.52% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   14.68n ± 63%   38.03n ± 13%  +159.06% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   26.01n ±  1%   41.96n ±  6%   +61.30% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.47n ±  9%   42.62n ±  3%  +130.66% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    26.11n ± 10%   43.08n ±  4%   +64.99% (p=0.000 n=20)
geomean                                20.36n         40.50n         +98.92%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            critbitgo/lpm.bm             │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            17.21n ±  6%   162.65n ±  6%   +844.82% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            35.12n ± 25%   250.05n ± 14%   +611.99% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             19.51n ±  0%   630.25n ± 18%  +3130.39% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             30.09n ± 37%   485.70n ± 12%  +1514.43% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.97n ±  7%    83.98n ±  5%   +394.90% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.27n ±  0%   105.30n ±  3%   +446.30% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.71n ±  1%   149.45n ± 13%   +794.64% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.27n ±  0%   159.20n ±  8%   +726.15% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.24n ±  1%    98.94n ±  5%   +442.43% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.62n ±  9%   117.45n ±  3%   +498.62% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.21n ±  2%   254.40n ± 10%  +1296.65% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.44n ± 33%   221.05n ± 15%  +1037.09% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   14.68n ± 63%   110.25n ±  5%   +651.02% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   26.01n ±  1%   132.15n ±  2%   +408.07% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.47n ±  9%   472.55n ± 16%  +2457.78% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    26.11n ± 10%   332.55n ± 18%  +1173.65% (p=0.000 n=20)
geomean                                20.36n          193.7n         +851.24%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             lpmtrie/lpm.bm              │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            17.21n ±  6%   279.85n ± 10%  +1525.62% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            35.12n ± 25%   323.35n ± 17%   +820.70% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             19.51n ±  0%   260.05n ±  5%  +1232.91% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             30.09n ± 37%   234.05n ± 15%   +677.96% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.97n ±  7%   135.45n ±  8%   +698.17% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.27n ±  0%   120.30n ± 23%   +524.12% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.71n ±  1%   115.15n ± 17%   +589.31% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.27n ±  0%    95.28n ± 37%   +394.45% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.24n ±  1%   153.75n ± 13%   +742.93% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.62n ±  9%   140.00n ±  9%   +613.56% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.21n ±  2%   163.00n ± 11%   +794.87% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.44n ± 33%   135.40n ±  5%   +596.50% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   14.68n ± 63%   197.85n ±  5%  +1247.75% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   26.01n ±  1%   186.65n ±  2%   +617.61% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.47n ±  9%   199.00n ±  4%   +977.13% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    26.11n ± 10%   189.40n ±  5%   +625.39% (p=0.000 n=20)
geomean                                20.36n          173.1n         +750.29%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            cidranger/lpm.bm             │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            17.21n ±  6%   248.70n ± 23%  +1344.67% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            35.12n ± 25%   276.80n ± 29%   +688.15% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             19.51n ±  0%   281.30n ± 10%  +1341.82% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             30.09n ± 37%   290.95n ± 17%   +867.09% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.97n ±  7%   119.00n ± 12%   +601.24% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.27n ±  0%   155.20n ± 11%   +705.19% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.71n ±  1%   113.55n ±  5%   +579.74% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.27n ±  0%   149.85n ± 10%   +677.63% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.24n ±  1%   157.00n ±  8%   +760.75% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.62n ±  9%   181.10n ±  6%   +823.04% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.21n ±  2%   154.55n ±  4%   +748.48% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.44n ± 33%   178.65n ±  5%   +818.98% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   14.68n ± 63%   168.90n ± 20%  +1050.54% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   26.01n ±  1%   226.85n ±  3%   +772.16% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.47n ±  9%   188.65n ±  4%   +921.11% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    26.11n ± 10%   226.55n ±  7%   +767.68% (p=0.000 n=20)
geomean                                20.36n          187.2n         +819.47%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             cidrtree/lpm.bm              │
                                     │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            17.21n ±  6%    701.40n ± 12%  +3974.35% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            35.12n ± 25%    802.75n ± 19%  +2185.73% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             19.51n ±  0%    980.15n ± 13%  +4923.83% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             30.09n ± 37%   1099.00n ± 30%  +3552.98% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.97n ±  7%    336.60n ± 12%  +1883.50% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.27n ±  0%    396.50n ± 30%  +1957.07% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.71n ±  1%    380.40n ± 17%  +2177.16% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.27n ±  0%    489.20n ± 29%  +2438.66% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.24n ±  1%    487.45n ± 15%  +2572.42% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.62n ±  9%    520.20n ± 11%  +2551.38% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.21n ±  2%    637.55n ± 18%  +3400.14% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.44n ± 33%    795.75n ± 22%  +3993.36% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   14.68n ± 63%    563.65n ± 12%  +3739.58% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   26.01n ±  1%    693.60n ± 16%  +2566.67% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.47n ±  9%    879.60n ± 19%  +4661.03% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    26.11n ± 10%   1085.00n ± 11%  +4055.50% (p=0.000 n=20)
geomean                                20.36n           636.4n        +3025.43%
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
Tier1PfxSize/1_000           104.1 ± 2%    7591.0 ± 0%  +7192.03% (p=0.002 n=6)
Tier1PfxSize/10_000          76.06 ± 0%   4889.00 ± 0%  +6327.82% (p=0.002 n=6)
Tier1PfxSize/100_000         41.17 ± 0%   1669.00 ± 0%  +3953.92% (p=0.002 n=6)
Tier1PfxSize/200_000         32.85 ± 0%   1098.00 ± 0%  +3242.47% (p=0.002 n=6)
Tier1PfxSize/500_000         25.19 ± 0%
Tier1PfxSize/1_000_000       21.39 ± 0%
RandomPfx4Size/1_000         74.64 ± 3%   5259.00 ± 0%  +6945.82% (p=0.002 n=6)
RandomPfx4Size/10_000        48.26 ± 0%   4059.00 ± 0%  +8310.69% (p=0.002 n=6)
RandomPfx4Size/100_000       59.57 ± 0%   3938.00 ± 0%  +6510.71% (p=0.002 n=6)
RandomPfx4Size/200_000       52.51 ± 0%   3476.00 ± 0%  +6519.69% (p=0.002 n=6)
RandomPfx4Size/500_000       37.84 ± 0%
RandomPfx4Size/1_000_000     31.83 ± 0%
RandomPfx6Size/1_000         82.67 ± 2%   6761.00 ± 0%  +8078.30% (p=0.002 n=6)
RandomPfx6Size/10_000        98.87 ± 0%   7333.00 ± 0%  +7316.81% (p=0.002 n=6)
RandomPfx6Size/100_000       67.59 ± 0%   5708.00 ± 0%  +8345.04% (p=0.002 n=6)
RandomPfx6Size/200_000       64.96 ± 0%   5526.00 ± 0%  +8406.77% (p=0.002 n=6)
RandomPfx6Size/500_000       67.62 ± 0%
RandomPfx6Size/1_000_000     73.74 ± 0%
RandomPfxSize/1_000          98.98 ± 2%   7537.00 ± 0%  +7514.67% (p=0.002 n=6)
RandomPfxSize/10_000         71.38 ± 0%   6058.00 ± 0%  +8386.97% (p=0.002 n=6)
RandomPfxSize/100_000        76.54 ± 0%   5300.00 ± 0%  +6824.48% (p=0.002 n=6)
RandomPfxSize/200_000        65.49 ± 0%   4586.00 ± 0%  +6902.60% (p=0.002 n=6)
RandomPfxSize/500_000        47.95 ± 0%
RandomPfxSize/1_000_000      39.50 ± 0%
geomean                      56.09        4.449Ki       +6731.25%               ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │           netipds/size.bm           │
                         │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000          104.10 ± 2%    74.06 ± 3%   -28.86% (p=0.002 n=6)
Tier1PfxSize/10_000          76.06 ± 0%    68.92 ± 0%    -9.39% (p=0.002 n=6)
Tier1PfxSize/100_000         41.17 ± 0%    67.65 ± 0%   +64.32% (p=0.002 n=6)
Tier1PfxSize/200_000         32.85 ± 0%    66.91 ± 0%  +103.68% (p=0.002 n=6)
Tier1PfxSize/500_000         25.19 ± 0%    65.35 ± 0%  +159.43% (p=0.002 n=6)
Tier1PfxSize/1_000_000       21.39 ± 0%    63.72 ± 0%  +197.90% (p=0.002 n=6)
RandomPfx4Size/1_000         74.64 ± 3%    61.70 ± 3%   -17.34% (p=0.002 n=6)
RandomPfx4Size/10_000        48.26 ± 0%    51.35 ± 0%    +6.40% (p=0.002 n=6)
RandomPfx4Size/100_000       59.57 ± 0%    48.25 ± 0%   -19.00% (p=0.002 n=6)
RandomPfx4Size/200_000       52.51 ± 0%    47.57 ± 0%    -9.41% (p=0.002 n=6)
RandomPfx4Size/500_000       37.84 ± 0%    46.51 ± 0%   +22.91% (p=0.002 n=6)
RandomPfx4Size/1_000_000     31.83 ± 0%    45.66 ± 0%   +43.45% (p=0.002 n=6)
RandomPfx6Size/1_000         82.67 ± 2%   100.00 ± 2%   +20.96% (p=0.002 n=6)
RandomPfx6Size/10_000        98.87 ± 0%    94.26 ± 0%    -4.66% (p=0.002 n=6)
RandomPfx6Size/100_000       67.59 ± 0%    87.64 ± 0%   +29.66% (p=0.002 n=6)
RandomPfx6Size/200_000       64.96 ± 0%    85.90 ± 0%   +32.24% (p=0.002 n=6)
RandomPfx6Size/500_000       67.62 ± 0%    84.11 ± 0%   +24.39% (p=0.002 n=6)
RandomPfx6Size/1_000_000     73.74 ± 0%    83.23 ± 0%   +12.87% (p=0.002 n=6)
RandomPfxSize/1_000          98.98 ± 2%    74.70 ± 3%   -24.53% (p=0.002 n=6)
RandomPfxSize/10_000         71.38 ± 0%    69.38 ± 0%    -2.80% (p=0.002 n=6)
RandomPfxSize/100_000        76.54 ± 0%    64.29 ± 0%   -16.00% (p=0.002 n=6)
RandomPfxSize/200_000        65.49 ± 0%    61.55 ± 0%    -6.02% (p=0.002 n=6)
RandomPfxSize/500_000        47.95 ± 0%    57.42 ± 0%   +19.75% (p=0.002 n=6)
RandomPfxSize/1_000_000      39.50 ± 0%    53.93 ± 0%   +36.53% (p=0.002 n=6)
geomean                      56.09         66.01        +17.69%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │          critbitgo/size.bm          │
                         │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000           104.1 ± 2%    119.5 ± 2%   +14.79% (p=0.002 n=6)
Tier1PfxSize/10_000          76.06 ± 0%   114.70 ± 0%   +50.80% (p=0.002 n=6)
Tier1PfxSize/100_000         41.17 ± 0%   114.40 ± 0%  +177.87% (p=0.002 n=6)
Tier1PfxSize/200_000         32.85 ± 0%   114.40 ± 0%  +248.25% (p=0.002 n=6)
Tier1PfxSize/500_000         25.19 ± 0%   114.40 ± 0%  +354.15% (p=0.002 n=6)
Tier1PfxSize/1_000_000       21.39 ± 0%   114.40 ± 0%  +434.83% (p=0.002 n=6)
RandomPfx4Size/1_000         74.64 ± 3%   116.10 ± 2%   +55.55% (p=0.002 n=6)
RandomPfx4Size/10_000        48.26 ± 0%   112.40 ± 0%  +132.91% (p=0.002 n=6)
RandomPfx4Size/100_000       59.57 ± 0%   112.00 ± 0%   +88.01% (p=0.002 n=6)
RandomPfx4Size/200_000       52.51 ± 0%   112.00 ± 0%  +113.29% (p=0.002 n=6)
RandomPfx4Size/500_000       37.84 ± 0%   112.00 ± 0%  +195.98% (p=0.002 n=6)
RandomPfx4Size/1_000_000     31.83 ± 0%   112.00 ± 0%  +251.87% (p=0.002 n=6)
RandomPfx6Size/1_000         82.67 ± 2%   132.10 ± 1%   +59.79% (p=0.002 n=6)
RandomPfx6Size/10_000        98.87 ± 0%   128.40 ± 0%   +29.87% (p=0.002 n=6)
RandomPfx6Size/100_000       67.59 ± 0%   128.00 ± 0%   +89.38% (p=0.002 n=6)
RandomPfx6Size/200_000       64.96 ± 0%   128.00 ± 0%   +97.04% (p=0.002 n=6)
RandomPfx6Size/500_000       67.62 ± 0%   128.00 ± 0%   +89.29% (p=0.002 n=6)
RandomPfx6Size/1_000_000     73.74 ± 0%   128.00 ± 0%   +73.58% (p=0.002 n=6)
RandomPfxSize/1_000          98.98 ± 2%   119.50 ± 2%   +20.73% (p=0.002 n=6)
RandomPfxSize/10_000         71.38 ± 0%   115.60 ± 0%   +61.95% (p=0.002 n=6)
RandomPfxSize/100_000        76.54 ± 0%   115.30 ± 0%   +50.64% (p=0.002 n=6)
RandomPfxSize/200_000        65.49 ± 0%   115.20 ± 0%   +75.90% (p=0.002 n=6)
RandomPfxSize/500_000        47.95 ± 0%   115.20 ± 0%  +140.25% (p=0.002 n=6)
RandomPfxSize/1_000_000      39.50 ± 0%   115.20 ± 0%  +191.65% (p=0.002 n=6)
geomean                      56.09         118.0       +110.44%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │           lpmtrie/size.bm            │
                         │ bytes/route  │ bytes/route   vs base                │
Tier1PfxSize/1_000           104.1 ± 2%    215.4 ±  5%  +106.92% (p=0.002 n=6)
Tier1PfxSize/10_000          76.06 ± 0%   210.50 ±  5%  +176.76% (p=0.002 n=6)
Tier1PfxSize/100_000         41.17 ± 0%   209.90 ±  5%  +409.84% (p=0.002 n=6)
Tier1PfxSize/200_000         32.85 ± 0%   209.20 ±  5%  +536.83% (p=0.002 n=6)
Tier1PfxSize/500_000         25.19 ± 0%   207.90 ±  7%  +725.33% (p=0.002 n=6)
Tier1PfxSize/1_000_000       21.39 ± 0%   205.00 ±  7%  +858.39% (p=0.002 n=6)
RandomPfx4Size/1_000         74.64 ± 3%   205.20 ±  8%  +174.92% (p=0.002 n=6)
RandomPfx4Size/10_000        48.26 ± 0%   186.60 ±  9%  +286.66% (p=0.002 n=6)
RandomPfx4Size/100_000       59.57 ± 0%   179.60 ±  9%  +201.49% (p=0.002 n=6)
RandomPfx4Size/200_000       52.51 ± 0%   178.50 ± 10%  +239.94% (p=0.002 n=6)
RandomPfx4Size/500_000       37.84 ± 0%   176.80 ± 10%  +367.23% (p=0.002 n=6)
RandomPfx4Size/1_000_000     31.83 ± 0%   175.50 ± 10%  +451.37% (p=0.002 n=6)
RandomPfx6Size/1_000         82.67 ± 2%   227.70 ±  8%  +175.43% (p=0.002 n=6)
RandomPfx6Size/10_000        98.87 ± 0%   222.50 ±  8%  +125.04% (p=0.002 n=6)
RandomPfx6Size/100_000       67.59 ± 0%   213.90 ±  9%  +216.47% (p=0.002 n=6)
RandomPfx6Size/200_000       64.96 ± 0%   210.50 ±  9%  +224.05% (p=0.002 n=6)
RandomPfx6Size/500_000       67.62 ± 0%   206.70 ±  9%  +205.68% (p=0.002 n=6)
RandomPfx6Size/1_000_000     73.74 ± 0%   204.90 ±  9%  +177.87% (p=0.002 n=6)
RandomPfxSize/1_000          98.98 ± 2%   215.00 ±  5%  +117.22% (p=0.002 n=6)
RandomPfxSize/10_000         71.38 ± 0%   210.30 ±  5%  +194.62% (p=0.002 n=6)
RandomPfxSize/100_000        76.54 ± 0%   203.50 ±  7%  +165.87% (p=0.002 n=6)
RandomPfxSize/200_000        65.49 ± 0%   198.70 ±  8%  +203.41% (p=0.002 n=6)
RandomPfxSize/500_000        47.95 ± 0%   189.90 ±  9%  +296.04% (p=0.002 n=6)
RandomPfxSize/1_000_000      39.50 ± 0%   181.40 ± 10%  +359.24% (p=0.002 n=6)
geomean                      56.09         201.3        +258.94%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │           cidranger/size.bm            │
                         │ bytes/route  │ bytes/route  vs base                   │
Tier1PfxSize/1_000           104.1 ± 2%    539.5 ± 3%   +418.25% (p=0.002 n=6)
Tier1PfxSize/10_000          76.06 ± 0%   533.90 ± 3%   +601.95% (p=0.002 n=6)
Tier1PfxSize/100_000         41.17 ± 0%   527.20 ± 2%  +1180.54% (p=0.002 n=6)
Tier1PfxSize/200_000         32.85 ± 0%   522.20 ± 2%  +1489.65% (p=0.002 n=6)
Tier1PfxSize/500_000         25.19 ± 0%
Tier1PfxSize/1_000_000       21.39 ± 0%
RandomPfx4Size/1_000         74.64 ± 3%   481.50 ± 3%   +545.10% (p=0.002 n=6)
RandomPfx4Size/10_000        48.26 ± 0%   433.10 ± 3%   +797.43% (p=0.002 n=6)
RandomPfx4Size/100_000       59.57 ± 0%   413.70 ± 3%   +594.48% (p=0.002 n=6)
RandomPfx4Size/200_000       52.51 ± 0%   409.10 ± 3%   +679.09% (p=0.002 n=6)
RandomPfx4Size/500_000       37.84 ± 0%
RandomPfx4Size/1_000_000     31.83 ± 0%
RandomPfx6Size/1_000         82.67 ± 2%   595.10 ± 0%   +619.85% (p=0.002 n=6)
RandomPfx6Size/10_000        98.87 ± 0%   581.00 ± 0%   +487.64% (p=0.002 n=6)
RandomPfx6Size/100_000       67.59 ± 0%   547.20 ± 0%   +709.59% (p=0.002 n=6)
RandomPfx6Size/200_000       64.96 ± 0%   538.10 ± 0%   +728.36% (p=0.002 n=6)
RandomPfx6Size/500_000       67.62 ± 0%
RandomPfx6Size/1_000_000     73.74 ± 0%
RandomPfxSize/1_000          98.98 ± 2%   540.30 ± 2%   +445.87% (p=0.002 n=6)
RandomPfxSize/10_000         71.38 ± 0%   528.10 ± 2%   +639.84% (p=0.002 n=6)
RandomPfxSize/100_000        76.54 ± 0%   495.50 ± 2%   +547.37% (p=0.002 n=6)
RandomPfxSize/200_000        65.49 ± 0%   477.60 ± 2%   +629.27% (p=0.002 n=6)
RandomPfxSize/500_000        47.95 ± 0%
RandomPfxSize/1_000_000      39.50 ± 0%
geomean                      56.09         507.3        +660.60%               ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │          cidrtree/size.bm           │
                         │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000          104.10 ± 2%    69.26 ± 3%   -33.47% (p=0.002 n=6)
Tier1PfxSize/10_000          76.06 ± 0%    64.37 ± 0%   -15.37% (p=0.002 n=6)
Tier1PfxSize/100_000         41.17 ± 0%    64.04 ± 0%   +55.55% (p=0.002 n=6)
Tier1PfxSize/200_000         32.85 ± 0%    64.02 ± 0%   +94.89% (p=0.002 n=6)
Tier1PfxSize/500_000         25.19 ± 0%    64.01 ± 0%  +154.11% (p=0.002 n=6)
Tier1PfxSize/1_000_000       21.39 ± 0%    64.00 ± 0%  +199.21% (p=0.002 n=6)
RandomPfx4Size/1_000         74.64 ± 3%    68.16 ± 3%    -8.68% (p=0.002 n=6)
RandomPfx4Size/10_000        48.26 ± 0%    64.37 ± 0%   +33.38% (p=0.002 n=6)
RandomPfx4Size/100_000       59.57 ± 0%    64.04 ± 0%    +7.50% (p=0.002 n=6)
RandomPfx4Size/200_000       52.51 ± 0%    64.02 ± 0%   +21.92% (p=0.002 n=6)
RandomPfx4Size/500_000       37.84 ± 0%    64.01 ± 0%   +69.16% (p=0.002 n=6)
RandomPfx4Size/1_000_000     31.83 ± 0%    64.00 ± 0%  +101.07% (p=0.002 n=6)
RandomPfx6Size/1_000         82.67 ± 2%    68.16 ± 3%   -17.55% (p=0.002 n=6)
RandomPfx6Size/10_000        98.87 ± 0%    64.37 ± 0%   -34.89% (p=0.002 n=6)
RandomPfx6Size/100_000       67.59 ± 0%    64.04 ± 0%    -5.25% (p=0.002 n=6)
RandomPfx6Size/200_000       64.96 ± 0%    64.02 ± 0%    -1.45% (p=0.002 n=6)
RandomPfx6Size/500_000       67.62 ± 0%    64.01 ± 0%    -5.34% (p=0.002 n=6)
RandomPfx6Size/1_000_000     73.74 ± 0%    64.00 ± 0%   -13.21% (p=0.002 n=6)
RandomPfxSize/1_000          98.98 ± 2%    68.16 ± 3%   -31.14% (p=0.002 n=6)
RandomPfxSize/10_000         71.38 ± 0%    64.37 ± 0%    -9.82% (p=0.002 n=6)
RandomPfxSize/100_000        76.54 ± 0%    64.04 ± 0%   -16.33% (p=0.002 n=6)
RandomPfxSize/200_000        65.49 ± 0%    64.02 ± 0%    -2.24% (p=0.002 n=6)
RandomPfxSize/500_000        47.95 ± 0%    64.01 ± 0%   +33.49% (p=0.002 n=6)
RandomPfxSize/1_000_000      39.50 ± 0%    64.00 ± 0%   +62.03% (p=0.002 n=6)
geomean                      56.09         64.79        +15.52%
```

## update, insert/delete

`bart` is by far the fastest algorithm for updates.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │             art/update.bm              │
                         │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000        30.41n ± 1%   300.35n ±  3%   +887.67% (p=0.002 n=6)
InsertRandomPfxs/10_000       36.35n ± 0%   527.65n ±  6%  +1351.58% (p=0.002 n=6)
InsertRandomPfxs/100_000      101.8n ± 2%   2677.5n ± 21%  +2530.16% (p=0.002 n=6)
InsertRandomPfxs/200_000      159.0n ± 1%   1811.0n ± 55%  +1038.99% (p=0.002 n=6)
DeleteRandomPfxs/1_000        16.89n ± 3%    16.89n ±  1%          ~ (p=0.942 n=6)
DeleteRandomPfxs/10_000       16.85n ± 1%    16.65n ±  0%     -1.13% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.34n ± 2%    17.57n ±  3%     +1.36% (p=0.017 n=6)
DeleteRandomPfxs/200_000      18.51n ± 2%    19.68n ±  1%     +6.29% (p=0.002 n=6)
geomean                       33.62n         128.6n         +282.46%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           netipds/update.bm            │
                         │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000        30.41n ± 1%   160.35n ±  4%   +427.29% (p=0.002 n=6)
InsertRandomPfxs/10_000       36.35n ± 0%   229.10n ±  9%   +530.26% (p=0.002 n=6)
InsertRandomPfxs/100_000      101.8n ± 2%    414.2n ±  8%   +306.88% (p=0.002 n=6)
InsertRandomPfxs/200_000      159.0n ± 1%    630.2n ± 16%   +296.38% (p=0.002 n=6)
DeleteRandomPfxs/1_000        16.89n ± 3%   146.20n ± 11%   +765.86% (p=0.002 n=6)
DeleteRandomPfxs/10_000       16.85n ± 1%   212.40n ±  7%  +1160.91% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.34n ± 2%   483.45n ± 23%  +2688.87% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.51n ± 2%   575.70n ± 11%  +3010.21% (p=0.002 n=6)
geomean                       33.62n         308.9n         +818.73%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │         critbitgo/update.bm          │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000        30.41n ± 1%   168.90n ± 4%  +455.41% (p=0.002 n=6)
InsertRandomPfxs/10_000       36.35n ± 0%   239.30n ± 2%  +558.32% (p=0.002 n=6)
InsertRandomPfxs/100_000      101.8n ± 2%    600.0n ± 2%  +489.34% (p=0.002 n=6)
InsertRandomPfxs/200_000      159.0n ± 1%    791.1n ± 2%  +397.58% (p=0.002 n=6)
DeleteRandomPfxs/1_000        16.89n ± 3%    71.32n ± 3%  +322.39% (p=0.002 n=6)
DeleteRandomPfxs/10_000       16.85n ± 1%    73.06n ± 2%  +333.69% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.34n ± 2%    78.97n ± 1%  +355.55% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.51n ± 2%    85.51n ± 8%  +361.97% (p=0.002 n=6)
geomean                       33.62n         169.3n       +403.55%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           lpmtrie/update.bm            │
                         │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000        30.41n ± 1%   366.15n ±  5%  +1104.04% (p=0.002 n=6)
InsertRandomPfxs/10_000       36.35n ± 0%   493.90n ±  2%  +1258.73% (p=0.002 n=6)
InsertRandomPfxs/100_000      101.8n ± 2%   1173.5n ±  4%  +1052.75% (p=0.002 n=6)
InsertRandomPfxs/200_000      159.0n ± 1%   1229.0n ± 15%   +672.96% (p=0.002 n=6)
DeleteRandomPfxs/1_000        16.89n ± 3%    71.66n ±  1%   +324.43% (p=0.002 n=6)
DeleteRandomPfxs/10_000       16.85n ± 1%   132.85n ±  4%   +688.66% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.34n ± 2%   307.60n ±  6%  +1674.44% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.51n ± 2%   485.15n ±  3%  +2521.02% (p=0.002 n=6)
geomean                       33.62n         372.5n        +1007.87%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           cidranger/update.bm           │
                         │   sec/route    │   sec/route    vs base                  │
InsertRandomPfxs/1_000        30.41n ± 1%   4535.00n ± 3%  +14812.86% (p=0.002 n=6)
InsertRandomPfxs/10_000       36.35n ± 0%   6460.50n ± 4%  +17673.04% (p=0.002 n=6)
InsertRandomPfxs/100_000      101.8n ± 2%   10837.5n ± 3%  +10545.87% (p=0.002 n=6)
InsertRandomPfxs/200_000      159.0n ± 1%   11409.0n ± 4%   +7075.47% (p=0.002 n=6)
DeleteRandomPfxs/1_000        16.89n ± 3%     87.85n ± 3%    +420.31% (p=0.002 n=6)
DeleteRandomPfxs/10_000       16.85n ± 1%     90.95n ± 1%    +439.95% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.34n ± 2%    121.55n ± 2%    +601.18% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.51n ± 2%    237.75n ± 6%   +1184.44% (p=0.002 n=6)
geomean                       33.62n          977.9n        +2808.64%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           cidrtree/update.bm            │
                         │   sec/route    │   sec/route     vs base                 │
InsertRandomPfxs/1_000        30.41n ± 1%   1235.50n ±  3%  +3962.81% (p=0.002 n=6)
InsertRandomPfxs/10_000       36.35n ± 0%   1841.00n ±  2%  +4964.65% (p=0.002 n=6)
InsertRandomPfxs/100_000      101.8n ± 2%    3221.0n ±  1%  +3064.05% (p=0.002 n=6)
InsertRandomPfxs/200_000      159.0n ± 1%    3992.0n ±  6%  +2410.69% (p=0.002 n=6)
DeleteRandomPfxs/1_000        16.89n ± 3%     15.19n ±  7%    -10.04% (p=0.002 n=6)
DeleteRandomPfxs/10_000       16.85n ± 1%     15.03n ±  1%    -10.80% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.34n ± 2%     23.64n ± 11%    +36.37% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.51n ± 2%    289.40n ± 33%  +1463.48% (p=0.002 n=6)
geomean                       33.62n          286.7n         +752.76%
```
