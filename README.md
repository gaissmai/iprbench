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
	github.com/phemmer/go-iptrie 
	github.com/kentik/patricia
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
  $ export GOAMD64=v3

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
LpmTier1Pfxs/RandomMatchIP4-8            15.16n ±  9%   45.28n ± 17%  +198.78% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            24.43n ± 13%   61.06n ±  6%  +149.89% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             16.81n ±  0%   28.97n ±  1%   +72.34% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             25.72n ± 26%   35.42n ± 21%   +37.73% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     16.62n ±  0%   44.95n ±  0%  +170.46% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     19.67n ±  0%   46.10n ±  2%  +134.37% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      16.61n ± 20%   28.79n ±  0%   +73.36% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      19.68n ±  0%   28.91n ±  7%   +46.86% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    18.15n ±  0%   45.01n ±  1%  +147.99% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    19.65n ±  0%   46.47n ±  3%  +136.43% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     18.14n ±  0%   28.78n ±  1%   +58.65% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     19.67n ±  0%   29.00n ±  3%   +47.43% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   11.80n ± 39%   49.52n ±  9%  +319.63% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   25.93n ± 47%   49.70n ±  2%   +91.65% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    18.12n ±  7%   28.92n ±  1%   +59.60% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    25.95n ±  0%   33.06n ±  6%   +27.40% (p=0.000 n=20)
geomean                                  19.10n         38.15n         +99.73%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ bart/lpm.bm  │            netipds/lpm.bm             │
                                       │    sec/op    │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4-8            15.16n ±  9%   47.75n ± 16%  +215.11% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            24.43n ± 13%   45.34n ± 19%   +85.55% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             16.81n ±  0%   53.02n ±  5%  +215.41% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             25.72n ± 26%   52.74n ± 13%  +105.05% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     16.62n ±  0%   29.27n ±  4%   +76.11% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     19.67n ±  0%   31.68n ±  6%   +61.06% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      16.61n ± 20%   29.82n ±  5%   +79.53% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      19.68n ±  0%   31.59n ±  5%   +60.50% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    18.15n ±  0%   35.38n ±  5%   +94.93% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    19.65n ±  0%   35.54n ±  3%   +80.82% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     18.14n ±  0%   37.14n ±  6%  +104.77% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     19.67n ±  0%   36.55n ±  3%   +85.79% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   11.80n ± 39%   37.96n ± 14%  +221.64% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   25.93n ± 47%   40.23n ±  2%   +55.15% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    18.12n ±  7%   42.49n ±  3%  +134.52% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    25.95n ±  0%   42.34n ±  3%   +63.14% (p=0.000 n=20)
geomean                                  19.10n         38.65n        +102.37%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ bart/lpm.bm  │            critbitgo/lpm.bm             │
                                       │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4-8            15.16n ±  9%   162.50n ±  5%   +972.25% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            24.43n ± 13%   189.50n ±  4%   +675.53% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             16.81n ±  0%   629.35n ± 20%  +3643.90% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             25.72n ± 26%   481.80n ± 16%  +1773.25% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     16.62n ±  0%    85.01n ±  4%   +411.49% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     19.67n ±  0%    97.09n ±  5%   +393.59% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      16.61n ± 20%   150.60n ± 10%   +806.68% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      19.68n ±  0%   144.85n ± 11%   +635.84% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    18.15n ±  0%    98.20n ±  4%   +441.07% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    19.65n ±  0%   107.55n ±  3%   +447.19% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     18.14n ±  0%   258.45n ± 13%  +1324.75% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     19.67n ±  0%   203.75n ± 17%   +935.84% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   11.80n ± 39%   112.70n ±  4%   +854.92% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   25.93n ± 47%   120.25n ±  2%   +363.75% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    18.12n ±  7%   458.45n ± 14%  +2430.08% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    25.95n ±  0%   318.60n ± 17%  +1127.75% (p=0.000 n=20)
geomean                                  19.10n          184.8n         +867.61%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ bart/lpm.bm  │             lpmtrie/lpm.bm              │
                                       │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4-8            15.16n ±  9%   268.45n ±  6%  +1671.36% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            24.43n ± 13%   316.45n ± 16%  +1195.07% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             16.81n ±  0%   251.50n ±  5%  +1396.13% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             25.72n ± 26%   234.25n ± 13%   +810.77% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     16.62n ±  0%   130.30n ±  7%   +684.00% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     19.67n ±  0%   122.15n ± 25%   +521.00% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      16.61n ± 20%   112.80n ± 18%   +579.11% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      19.68n ±  0%    94.35n ± 34%   +379.30% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    18.15n ±  0%   150.55n ± 14%   +729.48% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    19.65n ±  0%   141.40n ±  5%   +619.41% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     18.14n ±  0%   163.15n ± 10%   +799.39% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     19.67n ±  0%   133.85n ±  5%   +580.48% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   11.80n ± 39%   194.05n ±  7%  +1544.21% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   25.93n ± 47%   186.35n ±  2%   +618.67% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    18.12n ±  7%   197.50n ±  4%   +989.96% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    25.95n ±  0%   189.90n ±  5%   +631.79% (p=0.000 n=20)
geomean                                  19.10n          171.0n         +795.33%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ bart/lpm.bm  │            cidranger/lpm.bm             │
                                       │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4-8            15.16n ±  9%   271.85n ± 32%  +1693.80% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            24.43n ± 13%   274.95n ± 34%  +1025.23% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             16.81n ±  0%   288.65n ±  5%  +1617.13% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             25.72n ± 26%   295.75n ± 17%  +1049.88% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     16.62n ±  0%   114.45n ±  8%   +588.63% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     19.67n ±  0%   147.45n ±  8%   +649.62% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      16.61n ± 20%   109.20n ±  6%   +557.44% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      19.68n ±  0%   140.05n ± 11%   +611.46% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    18.15n ±  0%   151.25n ±  7%   +733.33% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    19.65n ±  0%   187.85n ±  8%   +855.74% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     18.14n ±  0%   156.75n ±  7%   +764.11% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     19.67n ±  0%   182.25n ±  6%   +826.54% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   11.80n ± 39%   163.05n ± 21%  +1281.55% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   25.93n ± 47%   221.75n ±  5%   +755.19% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    18.12n ±  7%   191.50n ±  5%   +956.84% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    25.95n ±  0%   221.50n ±  6%   +753.56% (p=0.000 n=20)
geomean                                  19.10n          186.0n         +873.81%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ bart/lpm.bm  │             cidrtree/lpm.bm              │
                                       │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4-8            15.16n ±  9%    688.65n ± 13%  +4444.04% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            24.43n ± 13%    950.90n ± 15%  +3791.55% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             16.81n ±  0%   1092.50n ± 16%  +6399.11% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             25.72n ± 26%   1527.00n ± 19%  +5837.01% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     16.62n ±  0%    366.00n ± 20%  +2102.17% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     19.67n ±  0%    350.70n ± 20%  +1682.92% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      16.61n ± 20%    470.80n ± 18%  +2734.44% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      19.68n ±  0%    507.10n ± 32%  +2476.07% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    18.15n ±  0%    501.75n ± 19%  +2664.46% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    19.65n ±  0%    538.20n ±  8%  +2638.23% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     18.14n ±  0%    666.15n ± 27%  +3572.27% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     19.67n ±  0%    661.80n ± 20%  +3264.51% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   11.80n ± 39%    558.65n ± 22%  +4633.52% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   25.93n ± 47%    673.30n ± 14%  +2496.61% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    18.12n ±  7%    900.65n ± 11%  +4870.47% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    25.95n ±  0%    916.10n ± 21%  +3430.25% (p=0.000 n=20)
geomean                                  19.10n           658.1n        +3345.54%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ bart/lpm.bm  │            go-iptrie/lpm.bm             │
                                       │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4-8            15.16n ±  9%   200.45n ± 17%  +1222.67% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            24.43n ± 13%   182.65n ± 28%   +647.49% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             16.81n ±  0%   208.65n ±  4%  +1141.23% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             25.72n ± 26%   200.35n ± 11%   +678.97% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     16.62n ±  0%   114.60n ±  8%   +589.53% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     19.67n ±  0%   108.40n ± 18%   +451.09% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      16.61n ± 20%   111.75n ±  4%   +572.79% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      19.68n ±  0%   100.35n ±  9%   +409.78% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    18.15n ±  0%   141.75n ±  6%   +680.99% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    19.65n ±  0%   135.65n ±  6%   +590.16% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     18.14n ±  0%   146.20n ±  4%   +705.95% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     19.67n ±  0%   131.80n ±  4%   +570.06% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   11.80n ± 39%   148.90n ± 13%  +1161.65% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   25.93n ± 47%   162.15n ±  0%   +525.34% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    18.12n ±  7%   165.95n ±  4%   +815.84% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    25.95n ±  0%   172.80n ±  4%   +565.90% (p=0.000 n=20)
geomean                                  19.10n          148.3n         +676.44%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ bart/lpm.bm  │         kentik-patricia/lpm.bm          │
                                       │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4-8            15.16n ±  9%   174.55n ± 11%  +1051.77% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            24.43n ± 13%   260.55n ± 22%   +966.30% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             16.81n ±  0%   137.55n ±  5%   +718.26% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             25.72n ± 26%   166.75n ± 12%   +548.33% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     16.62n ±  0%    79.97n ±  7%   +381.14% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     19.67n ±  0%   100.03n ± 15%   +408.57% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      16.61n ± 20%    66.91n ±  6%   +302.83% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      19.68n ±  0%    78.96n ± 11%   +301.12% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    18.15n ±  0%   102.80n ±  5%   +466.39% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    19.65n ±  0%   124.55n ±  4%   +533.68% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     18.14n ±  0%    91.05n ±  4%   +401.93% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     19.67n ±  0%   108.50n ±  4%   +451.60% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   11.80n ± 39%   130.35n ± 11%  +1004.47% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   25.93n ± 47%   153.75n ±  5%   +492.94% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    18.12n ±  7%   106.60n ±  4%   +488.30% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    25.95n ±  0%   137.65n ±  6%   +430.44% (p=0.000 n=20)
geomean                                  19.10n          119.1n         +523.61%
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
Tier1PfxSize/1_000-8           104.1 ± 2%    7591.0 ± 0%  +7192.03% (p=0.002 n=6)
Tier1PfxSize/10_000-8          76.02 ± 0%   4889.00 ± 0%  +6331.20% (p=0.002 n=6)
Tier1PfxSize/100_000-8         41.16 ± 0%   1669.00 ± 0%  +3954.91% (p=0.002 n=6)
Tier1PfxSize/200_000-8         32.85 ± 0%   1098.00 ± 0%  +3242.47% (p=0.002 n=6)
Tier1PfxSize/500_000-8         25.19 ± 0%
Tier1PfxSize/1_000_000-8       21.39 ± 0%
RandomPfx4Size/1_000-8         74.74 ± 3%   5260.00 ± 0%  +6937.73% (p=0.002 n=6)
RandomPfx4Size/10_000-8        48.17 ± 0%   4059.00 ± 0%  +8326.41% (p=0.002 n=6)
RandomPfx4Size/100_000-8       59.56 ± 0%   3938.00 ± 0%  +6511.82% (p=0.002 n=6)
RandomPfx4Size/200_000-8       52.50 ± 0%   3476.00 ± 0%  +6520.95% (p=0.002 n=6)
RandomPfx4Size/500_000-8       37.84 ± 0%
RandomPfx4Size/1_000_000-8     31.83 ± 0%
RandomPfx6Size/1_000-8         82.77 ± 2%   6761.00 ± 0%  +8068.42% (p=0.002 n=6)
RandomPfx6Size/10_000-8        98.78 ± 0%   7333.00 ± 0%  +7323.57% (p=0.002 n=6)
RandomPfx6Size/100_000-8       67.59 ± 0%   5708.00 ± 0%  +8345.04% (p=0.002 n=6)
RandomPfx6Size/200_000-8       64.96 ± 0%   5526.00 ± 0%  +8406.77% (p=0.002 n=6)
RandomPfx6Size/500_000-8       67.61 ± 0%
RandomPfx6Size/1_000_000-8     73.74 ± 0%
RandomPfxSize/1_000-8          99.07 ± 2%   7538.00 ± 0%  +7508.76% (p=0.002 n=6)
RandomPfxSize/10_000-8         71.29 ± 0%   6058.00 ± 0%  +8397.69% (p=0.002 n=6)
RandomPfxSize/100_000-8        76.53 ± 0%   5300.00 ± 0%  +6825.39% (p=0.002 n=6)
RandomPfxSize/200_000-8        65.49 ± 0%   4586.00 ± 0%  +6902.60% (p=0.002 n=6)
RandomPfxSize/500_000-8        47.94 ± 0%
RandomPfxSize/1_000_000-8      39.50 ± 0%
geomean                        56.08        4.450Ki       +6732.17%               ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/size.bm │           netipds/size.bm           │
                           │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000-8          104.10 ± 2%    74.10 ± 3%   -28.81% (p=0.002 n=6)
Tier1PfxSize/10_000-8          76.02 ± 0%    68.83 ± 0%    -9.46% (p=0.002 n=6)
Tier1PfxSize/100_000-8         41.16 ± 0%    67.65 ± 0%   +64.36% (p=0.002 n=6)
Tier1PfxSize/200_000-8         32.85 ± 0%    66.90 ± 0%  +103.65% (p=0.002 n=6)
Tier1PfxSize/500_000-8         25.19 ± 0%    65.35 ± 0%  +159.43% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       21.39 ± 0%    63.72 ± 0%  +197.90% (p=0.002 n=6)
RandomPfx4Size/1_000-8         74.74 ± 3%    61.79 ± 3%   -17.33% (p=0.002 n=6)
RandomPfx4Size/10_000-8        48.17 ± 0%    51.25 ± 0%    +6.39% (p=0.002 n=6)
RandomPfx4Size/100_000-8       59.56 ± 0%    48.24 ± 0%   -19.01% (p=0.002 n=6)
RandomPfx4Size/200_000-8       52.50 ± 0%    47.57 ± 0%    -9.39% (p=0.002 n=6)
RandomPfx4Size/500_000-8       37.84 ± 0%    46.51 ± 0%   +22.91% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     31.83 ± 0%    45.65 ± 0%   +43.42% (p=0.002 n=6)
RandomPfx6Size/1_000-8         82.77 ± 2%   100.10 ± 2%   +20.94% (p=0.002 n=6)
RandomPfx6Size/10_000-8        98.78 ± 0%    94.17 ± 0%    -4.67% (p=0.002 n=6)
RandomPfx6Size/100_000-8       67.59 ± 0%    87.63 ± 0%   +29.65% (p=0.002 n=6)
RandomPfx6Size/200_000-8       64.96 ± 0%    85.90 ± 0%   +32.24% (p=0.002 n=6)
RandomPfx6Size/500_000-8       67.61 ± 0%    84.11 ± 0%   +24.40% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     73.74 ± 0%    83.23 ± 0%   +12.87% (p=0.002 n=6)
RandomPfxSize/1_000-8          99.07 ± 2%    74.80 ± 3%   -24.50% (p=0.002 n=6)
RandomPfxSize/10_000-8         71.29 ± 0%    69.34 ± 0%    -2.74% (p=0.002 n=6)
RandomPfxSize/100_000-8        76.53 ± 0%    64.28 ± 0%   -16.01% (p=0.002 n=6)
RandomPfxSize/200_000-8        65.49 ± 0%    61.55 ± 0%    -6.02% (p=0.002 n=6)
RandomPfxSize/500_000-8        47.94 ± 0%    57.42 ± 0%   +19.77% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      39.50 ± 0%    53.93 ± 0%   +36.53% (p=0.002 n=6)
geomean                        56.08         66.01        +17.70%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/size.bm │          critbitgo/size.bm          │
                           │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000-8           104.1 ± 2%    119.5 ± 2%   +14.79% (p=0.002 n=6)
Tier1PfxSize/10_000-8          76.02 ± 0%   114.70 ± 0%   +50.88% (p=0.002 n=6)
Tier1PfxSize/100_000-8         41.16 ± 0%   114.40 ± 0%  +177.94% (p=0.002 n=6)
Tier1PfxSize/200_000-8         32.85 ± 0%   114.40 ± 0%  +248.25% (p=0.002 n=6)
Tier1PfxSize/500_000-8         25.19 ± 0%   114.40 ± 0%  +354.15% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       21.39 ± 0%   114.40 ± 0%  +434.83% (p=0.002 n=6)
RandomPfx4Size/1_000-8         74.74 ± 3%   116.80 ± 2%   +56.28% (p=0.002 n=6)
RandomPfx4Size/10_000-8        48.17 ± 0%   112.30 ± 0%  +133.13% (p=0.002 n=6)
RandomPfx4Size/100_000-8       59.56 ± 0%   112.00 ± 0%   +88.05% (p=0.002 n=6)
RandomPfx4Size/200_000-8       52.50 ± 0%   112.00 ± 0%  +113.33% (p=0.002 n=6)
RandomPfx4Size/500_000-8       37.84 ± 0%   112.00 ± 0%  +195.98% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     31.83 ± 0%   112.00 ± 0%  +251.87% (p=0.002 n=6)
RandomPfx6Size/1_000-8         82.77 ± 2%   132.30 ± 1%   +59.84% (p=0.002 n=6)
RandomPfx6Size/10_000-8        98.78 ± 0%   128.30 ± 0%   +29.88% (p=0.002 n=6)
RandomPfx6Size/100_000-8       67.59 ± 0%   128.00 ± 0%   +89.38% (p=0.002 n=6)
RandomPfx6Size/200_000-8       64.96 ± 0%   128.00 ± 0%   +97.04% (p=0.002 n=6)
RandomPfx6Size/500_000-8       67.61 ± 0%   128.00 ± 0%   +89.32% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     73.74 ± 0%   128.00 ± 0%   +73.58% (p=0.002 n=6)
RandomPfxSize/1_000-8          99.07 ± 2%   120.20 ± 2%   +21.33% (p=0.002 n=6)
RandomPfxSize/10_000-8         71.29 ± 0%   115.50 ± 0%   +62.01% (p=0.002 n=6)
RandomPfxSize/100_000-8        76.53 ± 0%   115.30 ± 0%   +50.66% (p=0.002 n=6)
RandomPfxSize/200_000-8        65.49 ± 0%   115.20 ± 0%   +75.90% (p=0.002 n=6)
RandomPfxSize/500_000-8        47.94 ± 0%   115.20 ± 0%  +140.30% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      39.50 ± 0%   115.20 ± 0%  +191.65% (p=0.002 n=6)
geomean                        56.08         118.1       +110.55%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/size.bm │           lpmtrie/size.bm            │
                           │ bytes/route  │ bytes/route   vs base                │
Tier1PfxSize/1_000-8           104.1 ± 2%    215.4 ±  5%  +106.92% (p=0.002 n=6)
Tier1PfxSize/10_000-8          76.02 ± 0%   210.50 ±  5%  +176.90% (p=0.002 n=6)
Tier1PfxSize/100_000-8         41.16 ± 0%   209.90 ±  5%  +409.96% (p=0.002 n=6)
Tier1PfxSize/200_000-8         32.85 ± 0%   209.20 ±  5%  +536.83% (p=0.002 n=6)
Tier1PfxSize/500_000-8         25.19 ± 0%   207.90 ±  7%  +725.33% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       21.39 ± 0%   205.00 ±  7%  +858.39% (p=0.002 n=6)
RandomPfx4Size/1_000-8         74.74 ± 3%   211.50 ± 10%  +182.98% (p=0.002 n=6)
RandomPfx4Size/10_000-8        48.17 ± 0%   186.50 ±  9%  +287.17% (p=0.002 n=6)
RandomPfx4Size/100_000-8       59.56 ± 0%   179.60 ±  9%  +201.54% (p=0.002 n=6)
RandomPfx4Size/200_000-8       52.50 ± 0%   178.50 ± 10%  +240.00% (p=0.002 n=6)
RandomPfx4Size/500_000-8       37.84 ± 0%   176.80 ± 10%  +367.23% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     31.83 ± 0%   175.50 ± 10%  +451.37% (p=0.002 n=6)
RandomPfx6Size/1_000-8         82.77 ± 2%   227.90 ±  8%  +175.34% (p=0.002 n=6)
RandomPfx6Size/10_000-8        98.78 ± 0%   222.40 ±  8%  +125.15% (p=0.002 n=6)
RandomPfx6Size/100_000-8       67.59 ± 0%   213.90 ±  9%  +216.47% (p=0.002 n=6)
RandomPfx6Size/200_000-8       64.96 ± 0%   210.50 ±  9%  +224.05% (p=0.002 n=6)
RandomPfx6Size/500_000-8       67.61 ± 0%   206.70 ±  9%  +205.72% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     73.74 ± 0%   204.90 ±  9%  +177.87% (p=0.002 n=6)
RandomPfxSize/1_000-8          99.07 ± 2%   215.20 ±  5%  +117.22% (p=0.002 n=6)
RandomPfxSize/10_000-8         71.29 ± 0%   210.20 ±  5%  +194.85% (p=0.002 n=6)
RandomPfxSize/100_000-8        76.53 ± 0%   203.50 ±  7%  +165.91% (p=0.002 n=6)
RandomPfxSize/200_000-8        65.49 ± 0%   198.70 ±  8%  +203.41% (p=0.002 n=6)
RandomPfxSize/500_000-8        47.94 ± 0%   189.90 ±  9%  +296.12% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      39.50 ± 0%   181.40 ± 10%  +359.24% (p=0.002 n=6)
geomean                        56.08         201.6        +259.43%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/size.bm │           cidranger/size.bm            │
                           │ bytes/route  │ bytes/route  vs base                   │
Tier1PfxSize/1_000-8           104.1 ± 2%    539.7 ± 3%   +418.44% (p=0.002 n=6)
Tier1PfxSize/10_000-8          76.02 ± 0%   533.80 ± 3%   +602.18% (p=0.002 n=6)
Tier1PfxSize/100_000-8         41.16 ± 0%   527.20 ± 2%  +1180.86% (p=0.002 n=6)
Tier1PfxSize/200_000-8         32.85 ± 0%   522.20 ± 2%  +1489.65% (p=0.002 n=6)
Tier1PfxSize/500_000-8         25.19 ± 0%
Tier1PfxSize/1_000_000-8       21.39 ± 0%
RandomPfx4Size/1_000-8         74.74 ± 3%   481.70 ± 3%   +544.50% (p=0.002 n=6)
RandomPfx4Size/10_000-8        48.17 ± 0%   433.00 ± 3%   +798.90% (p=0.002 n=6)
RandomPfx4Size/100_000-8       59.56 ± 0%   413.70 ± 3%   +594.59% (p=0.002 n=6)
RandomPfx4Size/200_000-8       52.50 ± 0%   409.10 ± 3%   +679.24% (p=0.002 n=6)
RandomPfx4Size/500_000-8       37.84 ± 0%
RandomPfx4Size/1_000_000-8     31.83 ± 0%
RandomPfx6Size/1_000-8         82.77 ± 2%   595.40 ± 0%   +619.34% (p=0.002 n=6)
RandomPfx6Size/10_000-8        98.78 ± 0%   581.00 ± 0%   +488.18% (p=0.002 n=6)
RandomPfx6Size/100_000-8       67.59 ± 0%   547.20 ± 0%   +709.59% (p=0.002 n=6)
RandomPfx6Size/200_000-8       64.96 ± 0%   538.10 ± 0%   +728.36% (p=0.002 n=6)
RandomPfx6Size/500_000-8       67.61 ± 0%
RandomPfx6Size/1_000_000-8     73.74 ± 0%
RandomPfxSize/1_000-8          99.07 ± 2%   541.00 ± 2%   +446.08% (p=0.002 n=6)
RandomPfxSize/10_000-8         71.29 ± 0%   528.10 ± 2%   +640.78% (p=0.002 n=6)
RandomPfxSize/100_000-8        76.53 ± 0%   495.50 ± 2%   +547.46% (p=0.002 n=6)
RandomPfxSize/200_000-8        65.49 ± 0%   477.60 ± 2%   +629.27% (p=0.002 n=6)
RandomPfxSize/500_000-8        47.94 ± 0%
RandomPfxSize/1_000_000-8      39.50 ± 0%
geomean                        56.08         507.4        +660.79%               ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/size.bm │          cidrtree/size.bm           │
                           │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000-8          104.10 ± 2%    69.44 ± 3%   -33.29% (p=0.002 n=6)
Tier1PfxSize/10_000-8          76.02 ± 0%    64.29 ± 0%   -15.43% (p=0.002 n=6)
Tier1PfxSize/100_000-8         41.16 ± 0%    64.03 ± 0%   +55.56% (p=0.002 n=6)
Tier1PfxSize/200_000-8         32.85 ± 0%    64.01 ± 0%   +94.86% (p=0.002 n=6)
Tier1PfxSize/500_000-8         25.19 ± 0%    64.01 ± 0%  +154.11% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       21.39 ± 0%    64.00 ± 0%  +199.21% (p=0.002 n=6)
RandomPfx4Size/1_000-8         74.74 ± 3%    68.40 ± 3%    -8.48% (p=0.002 n=6)
RandomPfx4Size/10_000-8        48.17 ± 0%    64.29 ± 0%   +33.46% (p=0.002 n=6)
RandomPfx4Size/100_000-8       59.56 ± 0%    64.03 ± 0%    +7.51% (p=0.002 n=6)
RandomPfx4Size/200_000-8       52.50 ± 0%    64.01 ± 0%   +21.92% (p=0.002 n=6)
RandomPfx4Size/500_000-8       37.84 ± 0%    64.01 ± 0%   +69.16% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     31.83 ± 0%    64.00 ± 0%  +101.07% (p=0.002 n=6)
RandomPfx6Size/1_000-8         82.77 ± 2%    68.38 ± 3%   -17.39% (p=0.002 n=6)
RandomPfx6Size/10_000-8        98.78 ± 0%    64.29 ± 0%   -34.92% (p=0.002 n=6)
RandomPfx6Size/100_000-8       67.59 ± 0%    64.04 ± 0%    -5.25% (p=0.002 n=6)
RandomPfx6Size/200_000-8       64.96 ± 0%    64.01 ± 0%    -1.46% (p=0.002 n=6)
RandomPfx6Size/500_000-8       67.61 ± 0%    64.01 ± 0%    -5.32% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     73.74 ± 0%    64.00 ± 0%   -13.21% (p=0.002 n=6)
RandomPfxSize/1_000-8          99.07 ± 2%    68.38 ± 3%   -30.98% (p=0.002 n=6)
RandomPfxSize/10_000-8         71.29 ± 0%    64.29 ± 0%    -9.82% (p=0.002 n=6)
RandomPfxSize/100_000-8        76.53 ± 0%    64.03 ± 0%   -16.33% (p=0.002 n=6)
RandomPfxSize/200_000-8        65.49 ± 0%    64.01 ± 0%    -2.26% (p=0.002 n=6)
RandomPfxSize/500_000-8        47.94 ± 0%    64.01 ± 0%   +33.52% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      39.50 ± 0%    64.00 ± 0%   +62.03% (p=0.002 n=6)
geomean                        56.08         64.81        +15.56%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/size.bm │          go-iptrie/size.bm          │
                           │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000-8           104.1 ± 2%    165.1 ± 1%   +58.60% (p=0.002 n=6)
Tier1PfxSize/10_000-8          76.02 ± 0%   159.80 ± 0%  +110.21% (p=0.002 n=6)
Tier1PfxSize/100_000-8         41.16 ± 0%   157.20 ± 0%  +281.92% (p=0.002 n=6)
Tier1PfxSize/200_000-8         32.85 ± 0%   155.40 ± 0%  +373.06% (p=0.002 n=6)
Tier1PfxSize/500_000-8         25.19 ± 0%   151.70 ± 0%  +502.22% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       21.39 ± 0%   147.80 ± 0%  +590.98% (p=0.002 n=6)
RandomPfx4Size/1_000-8         74.74 ± 3%   148.50 ± 1%   +98.69% (p=0.002 n=6)
RandomPfx4Size/10_000-8        48.17 ± 0%   127.80 ± 0%  +165.31% (p=0.002 n=6)
RandomPfx4Size/100_000-8       59.56 ± 0%   120.50 ± 0%  +102.32% (p=0.002 n=6)
RandomPfx4Size/200_000-8       52.50 ± 0%   118.90 ± 0%  +126.48% (p=0.002 n=6)
RandomPfx4Size/500_000-8       37.84 ± 0%   116.30 ± 0%  +207.35% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     31.83 ± 0%   114.10 ± 0%  +258.47% (p=0.002 n=6)
RandomPfx6Size/1_000-8         82.77 ± 2%   164.00 ± 1%   +98.14% (p=0.002 n=6)
RandomPfx6Size/10_000-8        98.78 ± 0%   156.70 ± 0%   +58.64% (p=0.002 n=6)
RandomPfx6Size/100_000-8       67.59 ± 0%   146.00 ± 0%  +116.01% (p=0.002 n=6)
RandomPfx6Size/200_000-8       64.96 ± 0%   143.20 ± 0%  +120.44% (p=0.002 n=6)
RandomPfx6Size/500_000-8       67.61 ± 0%   140.20 ± 0%  +107.37% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     73.74 ± 0%   138.70 ± 0%   +88.09% (p=0.002 n=6)
RandomPfxSize/1_000-8          99.07 ± 2%   163.80 ± 1%   +65.34% (p=0.002 n=6)
RandomPfxSize/10_000-8         71.29 ± 0%   156.90 ± 0%  +120.09% (p=0.002 n=6)
RandomPfxSize/100_000-8        76.53 ± 0%   145.10 ± 0%   +89.60% (p=0.002 n=6)
RandomPfxSize/200_000-8        65.49 ± 0%   138.70 ± 0%  +111.79% (p=0.002 n=6)
RandomPfxSize/500_000-8        47.94 ± 0%   128.80 ± 0%  +168.67% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      39.50 ± 0%   120.50 ± 0%  +205.06% (p=0.002 n=6)
geomean                        56.08         141.8       +152.85%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/size.bm │       kentik-patricia/size.bm       │
                           │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000-8           104.1 ± 2%    145.3 ± 1%   +39.58% (p=0.002 n=6)
Tier1PfxSize/10_000-8          76.02 ± 0%   200.20 ± 0%  +163.35% (p=0.002 n=6)
Tier1PfxSize/100_000-8         41.16 ± 0%   163.80 ± 0%  +297.96% (p=0.002 n=6)
Tier1PfxSize/200_000-8         32.85 ± 0%   164.00 ± 0%  +399.24% (p=0.002 n=6)
Tier1PfxSize/500_000-8         25.19 ± 0%   144.20 ± 0%  +472.45% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       21.39 ± 0%   144.20 ± 0%  +574.15% (p=0.002 n=6)
RandomPfx4Size/1_000-8         74.74 ± 3%   140.80 ± 1%   +88.39% (p=0.002 n=6)
RandomPfx4Size/10_000-8        48.17 ± 0%   109.50 ± 0%  +127.32% (p=0.002 n=6)
RandomPfx4Size/100_000-8       59.56 ± 0%   139.80 ± 0%  +134.72% (p=0.002 n=6)
RandomPfx4Size/200_000-8       52.50 ± 0%   139.80 ± 0%  +166.29% (p=0.002 n=6)
RandomPfx4Size/500_000-8       37.84 ± 0%   139.70 ± 0%  +269.19% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     31.83 ± 0%   139.80 ± 0%  +339.21% (p=0.002 n=6)
RandomPfx6Size/1_000-8         82.77 ± 2%   157.70 ± 1%   +90.53% (p=0.002 n=6)
RandomPfx6Size/10_000-8        98.78 ± 0%   201.30 ± 0%  +103.79% (p=0.002 n=6)
RandomPfx6Size/100_000-8       67.59 ± 0%   160.80 ± 0%  +137.91% (p=0.002 n=6)
RandomPfx6Size/200_000-8       64.96 ± 0%   160.80 ± 0%  +147.54% (p=0.002 n=6)
RandomPfx6Size/500_000-8       67.61 ± 0%   156.50 ± 0%  +131.47% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     73.74 ± 0%   156.40 ± 0%  +112.10% (p=0.002 n=6)
RandomPfxSize/1_000-8          99.07 ± 2%   144.60 ± 1%   +45.96% (p=0.002 n=6)
RandomPfxSize/10_000-8         71.29 ± 0%   140.10 ± 0%   +96.52% (p=0.002 n=6)
RandomPfxSize/100_000-8        76.53 ± 0%   180.00 ± 0%  +135.20% (p=0.002 n=6)
RandomPfxSize/200_000-8        65.49 ± 0%   180.00 ± 0%  +174.85% (p=0.002 n=6)
RandomPfxSize/500_000-8        47.94 ± 0%   144.00 ± 0%  +200.38% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      39.50 ± 0%   144.00 ± 0%  +264.56% (p=0.002 n=6)
geomean                        56.08         152.8       +172.41%
```

## update, insert/delete

`bart` is by far the fastest algorithm for updates.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/update.bm │             art/update.bm              │
                           │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000-8        150.2n ± 1%   1385.0n ±  4%   +822.10% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       110.3n ± 3%   1221.5n ±  3%  +1007.43% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      159.5n ± 2%   1418.0n ± 50%   +789.31% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      179.8n ± 3%   1512.0n ± 50%   +740.70% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        113.1n ± 2%    656.7n ±  1%   +480.59% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       80.31n ± 1%   716.80n ±  4%   +792.54% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      161.4n ± 4%    765.6n ±  5%   +374.50% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      202.8n ± 3%    569.6n ±  4%   +180.91% (p=0.002 n=6)
geomean                         139.2n         963.8n         +592.25%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/update.bm │          netipds/update.bm           │
                           │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000-8        150.2n ± 1%    222.3n ± 4%   +48.00% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       110.3n ± 3%    285.9n ± 3%  +159.20% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      159.5n ± 2%    380.5n ± 3%  +138.66% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      179.8n ± 3%    459.9n ± 4%  +155.74% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        113.1n ± 2%    147.0n ± 6%   +29.97% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       80.31n ± 1%   201.20n ± 0%  +150.53% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      161.4n ± 4%    304.9n ± 3%   +88.94% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      202.8n ± 3%    405.6n ± 1%  +100.05% (p=0.002 n=6)
geomean                         139.2n         282.6n       +102.98%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/update.bm │         critbitgo/update.bm          │
                           │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000-8        150.2n ± 1%    296.6n ± 3%   +97.50% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       110.3n ± 3%    375.4n ± 2%  +240.30% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      159.5n ± 2%    559.0n ± 5%  +250.61% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      179.8n ± 3%    711.4n ± 3%  +295.55% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        113.1n ± 2%    149.3n ± 1%   +32.05% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       80.31n ± 1%   205.45n ± 1%  +155.82% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      161.4n ± 4%    372.9n ± 2%  +131.11% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      202.8n ± 3%    484.0n ± 2%  +138.72% (p=0.002 n=6)
geomean                         139.2n         353.8n       +154.09%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/update.bm │          lpmtrie/update.bm           │
                           │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000-8        150.2n ± 1%    431.9n ± 2%  +187.52% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       110.3n ± 3%    507.1n ± 2%  +359.75% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      159.5n ± 2%    734.7n ± 1%  +360.77% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      179.8n ± 3%    840.7n ± 3%  +367.45% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        113.1n ± 2%    152.3n ± 2%   +34.70% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       80.31n ± 1%   234.10n ± 0%  +191.50% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      161.4n ± 4%    490.4n ± 1%  +203.90% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      202.8n ± 3%    647.0n ± 1%  +219.14% (p=0.002 n=6)
geomean                         139.2n         444.7n       +219.44%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/update.bm │          cidranger/update.bm           │
                           │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000-8        150.2n ± 1%    4497.5n ± 1%  +2894.34% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       110.3n ± 3%    5682.5n ± 4%  +5051.86% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      159.5n ± 2%    7624.5n ± 6%  +4681.75% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      179.8n ± 3%    8299.0n ± 6%  +4514.40% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        113.1n ± 2%     678.7n ± 0%   +500.09% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       80.31n ± 1%   1152.00n ± 1%  +1334.44% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      161.4n ± 4%    2395.0n ± 3%  +1384.35% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      202.8n ± 3%    2794.0n ± 1%  +1278.05% (p=0.002 n=6)
geomean                         139.2n          3.097µ       +2124.38%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/update.bm │           cidrtree/update.bm           │
                           │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000-8        150.2n ± 1%     990.6n ± 4%   +559.49% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       110.3n ± 3%    1524.0n ± 3%  +1281.69% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      159.5n ± 2%    2288.5n ± 1%  +1335.25% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      179.8n ± 3%    2821.5n ± 2%  +1468.81% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        113.1n ± 2%    1402.5n ± 0%  +1140.05% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       80.31n ± 1%   2367.00n ± 1%  +2847.33% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      161.4n ± 4%    3575.0n ± 3%  +2115.68% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      202.8n ± 3%    4187.5n ± 2%  +1965.35% (p=0.002 n=6)
geomean                         139.2n          2.166µ       +1455.73%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/update.bm │         go-iptrie/update.bm          │
                           │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000-8        150.2n ± 1%    246.6n ± 1%   +64.18% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       110.3n ± 3%    312.9n ± 2%  +183.73% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      159.5n ± 2%    470.2n ± 2%  +194.86% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      179.8n ± 3%    555.2n ± 2%  +208.70% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        113.1n ± 2%    145.0n ± 1%   +28.25% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       80.31n ± 1%   217.40n ± 0%  +170.70% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      161.4n ± 4%    446.2n ± 2%  +176.54% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      202.8n ± 3%    570.1n ± 1%  +181.21% (p=0.002 n=6)
geomean                         139.2n         335.8n       +141.18%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/update.bm │      kentik-patricia/update.bm       │
                           │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000-8        150.2n ± 1%    239.3n ± 2%   +59.32% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       110.3n ± 3%    333.6n ± 1%  +202.49% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      159.5n ± 2%    523.9n ± 2%  +228.57% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      179.8n ± 3%    656.5n ± 3%  +265.03% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        113.1n ± 2%    304.8n ± 1%  +169.45% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       80.31n ± 1%   371.30n ± 1%  +362.33% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      161.4n ± 4%    677.2n ± 2%  +319.71% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      202.8n ± 3%    830.9n ± 1%  +309.82% (p=0.002 n=6)
geomean                         139.2n         452.2n       +224.80%
```
