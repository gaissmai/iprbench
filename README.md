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
LpmTier1Pfxs/RandomMatchIP4            17.31n ±  7%   45.64n ± 16%  +163.66% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            38.38n ± 27%   62.52n ±  4%   +62.88% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             19.42n ±  1%   28.98n ±  0%   +49.23% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             29.59n ± 29%   36.27n ± 22%   +22.58% (p=0.003 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.76n ±  7%   44.97n ±  1%  +168.32% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.88n ±  0%   46.17n ±  2%  +132.33% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.65n ±  4%   28.76n ±  0%   +72.73% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.89n ±  0%   29.00n ±  7%   +45.77% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.03n ±  0%   45.10n ±  0%  +150.18% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.95n ± 33%   46.07n ±  6%  +130.93% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.02n ±  0%   28.77n ±  1%   +59.63% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.93n ±  0%   28.89n ±  6%   +44.92% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   14.84n ± 40%   49.25n ± 10%  +231.87% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   26.64n ±  3%   48.80n ±  7%   +83.18% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.17n ±  8%   28.87n ±  2%   +58.93% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    26.70n ±  3%   33.01n ±  6%   +23.66% (p=0.000 n=20)
geomean                                20.59n         38.20n         +85.52%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            netipds/lpm.bm             │
                                     │    sec/op    │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4            17.31n ±  7%   60.01n ± 12%  +246.68% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            38.38n ± 27%   47.14n ± 18%   +22.80% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             19.42n ±  1%   64.91n ±  6%  +234.24% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             29.59n ± 29%   52.71n ± 11%   +78.15% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.76n ±  7%   39.52n ±  5%  +135.77% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.88n ±  0%   32.62n ±  9%   +64.13% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.65n ±  4%   40.16n ±  3%  +141.20% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.89n ±  0%   32.46n ±  6%   +63.18% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.03n ±  0%   45.75n ±  4%  +153.81% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.95n ± 33%   37.01n ±  3%   +85.51% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.02n ±  0%   49.38n ±  4%  +174.06% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.93n ±  0%   38.69n ±  3%   +94.06% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   14.84n ± 40%   47.85n ±  9%  +222.44% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   26.64n ±  3%   42.10n ±  2%   +58.01% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.17n ±  8%   53.68n ±  4%  +195.54% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    26.70n ±  3%   44.21n ±  3%   +65.63% (p=0.000 n=20)
geomean                                20.59n         44.67n        +116.96%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            critbitgo/lpm.bm             │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            17.31n ±  7%   162.65n ±  6%   +839.63% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            38.38n ± 27%   250.05n ± 14%   +551.43% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             19.42n ±  1%   630.25n ± 18%  +3145.37% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             29.59n ± 29%   485.70n ± 12%  +1541.71% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.76n ±  7%    83.98n ±  5%   +401.10% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.88n ±  0%   105.30n ±  3%   +429.81% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.65n ±  4%   149.45n ± 13%   +797.60% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.89n ±  0%   159.20n ±  8%   +700.20% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.03n ±  0%    98.94n ±  5%   +448.90% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.95n ± 33%   117.45n ±  3%   +488.72% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.02n ±  0%   254.40n ± 10%  +1311.76% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.93n ±  0%   221.05n ± 15%  +1008.85% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   14.84n ± 40%   110.25n ±  5%   +642.92% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   26.64n ±  3%   132.15n ±  2%   +396.06% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.17n ±  8%   472.55n ± 16%  +2501.43% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    26.70n ±  3%   332.55n ± 18%  +1145.74% (p=0.000 n=20)
geomean                                20.59n          193.7n         +840.65%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             lpmtrie/lpm.bm              │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            17.31n ±  7%   279.85n ± 10%  +1516.70% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            38.38n ± 27%   323.35n ± 17%   +742.39% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             19.42n ±  1%   260.05n ±  5%  +1239.08% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             29.59n ± 29%   234.05n ± 15%   +691.11% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.76n ±  7%   135.45n ±  8%   +708.17% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.88n ±  0%   120.30n ± 23%   +505.28% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.65n ±  4%   115.15n ± 17%   +591.59% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.89n ±  0%    95.28n ± 37%   +378.91% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.03n ±  0%   153.75n ± 13%   +752.98% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.95n ± 33%   140.00n ±  9%   +601.75% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.02n ±  0%   163.00n ± 11%   +804.55% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.93n ±  0%   135.40n ±  5%   +579.21% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   14.84n ± 40%   197.85n ±  5%  +1233.22% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   26.64n ±  3%   186.65n ±  2%   +600.64% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.17n ±  8%   199.00n ±  4%   +995.51% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    26.70n ±  3%   189.40n ±  5%   +609.50% (p=0.000 n=20)
geomean                                20.59n          173.1n         +740.82%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            cidranger/lpm.bm             │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            17.31n ±  7%   248.70n ± 23%  +1336.74% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            38.38n ± 27%   276.80n ± 29%   +621.12% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             19.42n ±  1%   281.30n ± 10%  +1348.51% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             29.59n ± 29%   290.95n ± 17%   +883.44% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.76n ±  7%   119.00n ± 12%   +610.02% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.88n ±  0%   155.20n ± 11%   +680.88% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.65n ±  4%   113.55n ±  5%   +581.98% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.89n ±  0%   149.85n ± 10%   +653.20% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.03n ±  0%   157.00n ±  8%   +771.01% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.95n ± 33%   181.10n ±  6%   +807.77% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.02n ±  0%   154.55n ±  4%   +757.66% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.93n ±  0%   178.65n ±  5%   +796.16% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   14.84n ± 40%   168.90n ± 20%  +1038.14% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   26.64n ±  3%   226.85n ±  3%   +751.54% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.17n ±  8%   188.65n ±  4%   +938.54% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    26.70n ±  3%   226.55n ±  7%   +748.66% (p=0.000 n=20)
geomean                                20.59n          187.2n         +809.24%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             cidrtree/lpm.bm              │
                                     │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            17.31n ±  7%    701.40n ± 12%  +3951.99% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            38.38n ± 27%    802.75n ± 19%  +1991.31% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             19.42n ±  1%    980.15n ± 13%  +4947.12% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             29.59n ± 29%   1099.00n ± 30%  +3614.72% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.76n ±  7%    336.60n ± 12%  +1908.35% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.88n ±  0%    396.50n ± 30%  +1894.97% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.65n ±  4%    380.40n ± 17%  +2184.68% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.89n ±  0%    489.20n ± 29%  +2358.91% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.03n ±  0%    487.45n ± 15%  +2604.30% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.95n ± 33%    520.20n ± 11%  +2507.52% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.02n ±  0%    637.55n ± 18%  +3438.01% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.93n ±  0%    795.75n ± 22%  +3891.72% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   14.84n ± 40%    563.65n ± 12%  +3698.18% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   26.64n ±  3%    693.60n ± 16%  +2503.60% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.17n ±  8%    879.60n ± 19%  +4742.28% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    26.70n ±  3%   1085.00n ± 11%  +3964.43% (p=0.000 n=20)
geomean                                20.59n           636.4n        +2990.65%
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
InsertRandomPfxs/1_000        31.54n ± 2%   300.35n ±  3%   +852.28% (p=0.002 n=6)
InsertRandomPfxs/10_000       37.32n ± 3%   527.65n ±  6%  +1313.66% (p=0.002 n=6)
InsertRandomPfxs/100_000      103.6n ± 6%   2677.5n ± 21%  +2484.46% (p=0.002 n=6)
InsertRandomPfxs/200_000      158.7n ± 1%   1811.0n ± 55%  +1041.51% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.30n ± 3%    16.89n ±  1%     -2.34% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.27n ± 1%    16.65n ±  0%     -3.56% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.64n ± 1%    17.57n ±  3%          ~ (p=0.290 n=6)
DeleteRandomPfxs/200_000      18.86n ± 1%    19.68n ±  1%     +4.29% (p=0.002 n=6)
geomean                       34.32n         128.6n         +274.70%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           netipds/update.bm           │
                         │   sec/route    │  sec/route    vs base                 │
InsertRandomPfxs/1_000        31.54n ± 2%   131.30n ± 4%   +316.30% (p=0.002 n=6)
InsertRandomPfxs/10_000       37.32n ± 3%   190.50n ± 1%   +410.38% (p=0.002 n=6)
InsertRandomPfxs/100_000      103.6n ± 6%    392.8n ± 2%   +279.15% (p=0.002 n=6)
InsertRandomPfxs/200_000      158.7n ± 1%    585.4n ± 1%   +268.99% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.30n ± 3%   111.30n ± 6%   +543.35% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.27n ± 1%   166.80n ± 2%   +865.84% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.64n ± 1%   361.00n ± 2%  +1946.49% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.86n ± 1%   552.70n ± 0%  +2829.76% (p=0.002 n=6)
geomean                       34.32n         260.7n        +659.52%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │         critbitgo/update.bm          │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000        31.54n ± 2%   168.90n ± 4%  +435.51% (p=0.002 n=6)
InsertRandomPfxs/10_000       37.32n ± 3%   239.30n ± 2%  +541.13% (p=0.002 n=6)
InsertRandomPfxs/100_000      103.6n ± 6%    600.0n ± 2%  +479.10% (p=0.002 n=6)
InsertRandomPfxs/200_000      158.7n ± 1%    791.1n ± 2%  +398.68% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.30n ± 3%    71.32n ± 3%  +312.25% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.27n ± 1%    73.06n ± 2%  +323.02% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.64n ± 1%    78.97n ± 1%  +347.68% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.86n ± 1%    85.51n ± 8%  +353.27% (p=0.002 n=6)
geomean                       34.32n         169.3n       +393.33%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           lpmtrie/update.bm            │
                         │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000        31.54n ± 2%   366.15n ±  5%  +1060.91% (p=0.002 n=6)
InsertRandomPfxs/10_000       37.32n ± 3%   493.90n ±  2%  +1223.24% (p=0.002 n=6)
InsertRandomPfxs/100_000      103.6n ± 6%   1173.5n ±  4%  +1032.72% (p=0.002 n=6)
InsertRandomPfxs/200_000      158.7n ± 1%   1229.0n ± 15%   +674.66% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.30n ± 3%    71.66n ±  1%   +314.25% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.27n ± 1%   132.85n ±  4%   +669.25% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.64n ± 1%   307.60n ±  6%  +1643.76% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.86n ± 1%   485.15n ±  3%  +2471.69% (p=0.002 n=6)
geomean                       34.32n         372.5n         +985.40%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           cidranger/update.bm           │
                         │   sec/route    │   sec/route    vs base                  │
InsertRandomPfxs/1_000        31.54n ± 2%   4535.00n ± 3%  +14278.57% (p=0.002 n=6)
InsertRandomPfxs/10_000       37.32n ± 3%   6460.50n ± 4%  +17208.77% (p=0.002 n=6)
InsertRandomPfxs/100_000      103.6n ± 6%   10837.5n ± 3%  +10360.91% (p=0.002 n=6)
InsertRandomPfxs/200_000      158.7n ± 1%   11409.0n ± 4%   +7091.30% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.30n ± 3%     87.85n ± 3%    +407.83% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.27n ± 1%     90.95n ± 1%    +426.66% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.64n ± 1%    121.55n ± 2%    +589.06% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.86n ± 1%    237.75n ± 6%   +1160.27% (p=0.002 n=6)
geomean                       34.32n          977.9n        +2749.65%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           cidrtree/update.bm            │
                         │   sec/route    │   sec/route     vs base                 │
InsertRandomPfxs/1_000        31.54n ± 2%   1235.50n ±  3%  +3817.25% (p=0.002 n=6)
InsertRandomPfxs/10_000       37.32n ± 3%   1841.00n ±  2%  +4832.35% (p=0.002 n=6)
InsertRandomPfxs/100_000      103.6n ± 6%    3221.0n ±  1%  +3009.07% (p=0.002 n=6)
InsertRandomPfxs/200_000      158.7n ± 1%    3992.0n ±  6%  +2416.23% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.30n ± 3%     15.19n ±  7%    -12.20% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.27n ± 1%     15.03n ±  1%    -13.00% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.64n ± 1%     23.64n ± 11%    +34.01% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.86n ± 1%    289.40n ± 33%  +1434.06% (p=0.002 n=6)
geomean                       34.32n          286.7n         +735.47%
```
