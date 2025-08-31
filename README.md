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
LpmTier1Pfxs/RandomMatchIP4            16.38n ±  5%   45.19n ± 17%  +175.80% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            28.70n ± 30%   61.20n ±  7%  +113.19% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             18.61n ±  0%   28.77n ±  1%   +54.55% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             29.07n ± 36%   35.39n ± 21%   +21.72% (p=0.001 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.54n ±  0%   44.90n ±  1%  +171.41% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.24n ±  0%   46.02n ±  2%  +139.22% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.59n ±  9%   28.79n ±  1%   +73.54% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.23n ±  0%   28.93n ±  7%   +50.44% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.19n ±  0%   45.02n ±  1%  +147.50% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.25n ±  0%   46.58n ±  3%  +141.97% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.16n ±  0%   28.79n ±  1%   +58.51% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.27n ± 33%   28.94n ±  3%   +50.22% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   14.88n ± 41%   49.26n ± 10%  +230.97% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   25.71n ±  0%   50.01n ±  5%   +94.50% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.32n ± 39%   29.06n ±  1%   +58.62% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    25.71n ±  0%   32.95n ±  6%   +28.16% (p=0.000 n=20)
geomean                                19.83n         38.13n         +92.31%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            netipds/lpm.bm             │
                                     │    sec/op    │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4            16.38n ±  5%   46.04n ± 14%  +181.02% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            28.70n ± 30%   61.17n ± 23%  +113.12% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             18.61n ±  0%   51.22n ±  5%  +175.13% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             29.07n ± 36%   64.70n ± 20%  +122.57% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.54n ±  0%   28.68n ±  5%   +73.35% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.24n ±  0%   31.94n ±  6%   +66.03% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.59n ±  9%   29.36n ±  6%   +77.00% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.23n ±  0%   31.95n ± 15%   +66.17% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.19n ±  0%   34.57n ±  5%   +90.02% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.25n ±  0%   37.29n ±  6%   +93.69% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.16n ±  0%   36.95n ±  5%  +103.50% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.27n ± 33%   43.38n ± 15%  +125.15% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   14.88n ± 41%   36.89n ± 14%  +147.87% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   25.71n ±  0%   45.32n ± 12%   +76.25% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.32n ± 39%   41.85n ±  1%  +128.44% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    25.71n ±  0%   49.30n ± 13%   +91.73% (p=0.000 n=20)
geomean                                19.83n         40.72n        +105.38%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            critbitgo/lpm.bm             │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            16.38n ±  5%   139.90n ±  2%   +753.83% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            28.70n ± 30%   202.55n ±  7%   +605.63% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             18.61n ±  0%   614.45n ± 20%  +3200.83% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             29.07n ± 36%   475.80n ± 13%  +1536.74% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.54n ±  0%    82.08n ±  4%   +396.10% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.24n ±  0%   103.70n ±  2%   +438.98% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.59n ±  9%   137.10n ± 14%   +726.40% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.23n ±  0%   151.55n ± 12%   +688.09% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.19n ±  0%    95.50n ±  4%   +425.01% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.25n ±  0%   113.35n ±  5%   +488.83% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.16n ±  0%   246.70n ± 14%  +1258.48% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.27n ± 33%   205.60n ± 14%   +967.22% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   14.88n ± 41%   109.80n ±  4%   +637.66% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   25.71n ±  0%   130.00n ±  3%   +405.64% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.32n ± 39%   449.95n ± 18%  +2356.06% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    25.71n ±  0%   329.25n ± 14%  +1180.63% (p=0.000 n=20)
geomean                                19.83n          183.8n         +827.13%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             lpmtrie/lpm.bm              │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            16.38n ±  5%   272.95n ±  7%  +1565.85% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            28.70n ± 30%   318.30n ± 14%  +1008.87% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             18.61n ±  0%   258.00n ±  7%  +1285.98% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             29.07n ± 36%   235.90n ± 15%   +711.49% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.54n ±  0%   132.30n ±  5%   +699.64% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.24n ±  0%   122.55n ± 24%   +536.95% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.59n ±  9%   114.70n ± 16%   +591.38% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.23n ±  0%    95.12n ± 37%   +394.62% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.19n ±  0%   153.25n ± 13%   +742.50% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.25n ±  0%   141.30n ±  5%   +634.03% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.16n ±  0%   163.10n ± 11%   +798.13% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.27n ± 33%   135.75n ±  6%   +604.65% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   14.88n ± 41%   196.15n ±  5%  +1217.77% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   25.71n ±  0%   187.50n ±  3%   +629.29% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.32n ± 39%   196.20n ±  4%   +970.96% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    25.71n ±  0%   189.90n ±  5%   +638.62% (p=0.000 n=20)
geomean                                19.83n          172.5n         +769.97%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            cidranger/lpm.bm             │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            16.38n ±  5%   252.85n ± 24%  +1443.18% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            28.70n ± 30%   271.10n ± 29%   +844.43% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             18.61n ±  0%   286.55n ±  5%  +1439.35% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             29.07n ± 36%   291.70n ± 18%   +903.44% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.54n ±  0%   116.45n ±  9%   +603.84% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.24n ±  0%   156.20n ± 12%   +711.85% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.59n ±  9%   112.05n ±  7%   +575.41% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.23n ±  0%   141.65n ±  9%   +636.61% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.19n ±  0%   155.40n ±  9%   +754.32% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.25n ±  0%   185.20n ±  6%   +862.08% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.16n ±  0%   152.65n ±  7%   +740.58% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.27n ± 33%   180.80n ±  6%   +838.49% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   14.88n ± 41%   155.55n ± 23%   +945.01% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   25.71n ±  0%   225.40n ±  4%   +776.70% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.32n ± 39%   191.55n ±  2%   +945.58% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    25.71n ±  0%   220.45n ±  6%   +757.45% (p=0.000 n=20)
geomean                                19.83n          185.4n         +834.94%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             cidrtree/lpm.bm              │
                                     │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            16.38n ±  5%    771.75n ± 12%  +4610.10% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            28.70n ± 30%    954.40n ± 15%  +3224.86% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             18.61n ±  0%   1260.50n ± 18%  +6671.42% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             29.07n ± 36%   1298.00n ± 10%  +4365.08% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.54n ±  0%    345.05n ± 21%  +1985.52% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.24n ±  0%    328.35n ± 19%  +1606.60% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.59n ±  9%    479.80n ± 16%  +2792.10% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.23n ±  0%    409.70n ± 12%  +2030.53% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.19n ±  0%    537.40n ± 11%  +2854.37% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.25n ±  0%    559.55n ± 11%  +2806.75% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.16n ±  0%    611.45n ± 22%  +3267.02% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.27n ± 33%    673.05n ± 19%  +3393.64% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   14.88n ± 41%    586.05n ± 23%  +3837.19% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   25.71n ±  0%    683.25n ± 20%  +2557.53% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.32n ± 39%    843.90n ± 19%  +4506.44% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    25.71n ±  0%   1058.00n ± 13%  +4015.13% (p=0.000 n=20)
geomean                                19.83n           656.4n        +3210.44%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            go-iptrie/lpm.bm             │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            16.38n ±  5%   227.70n ± 11%  +1289.69% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            28.70n ± 30%   211.10n ± 27%   +635.41% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             18.61n ±  0%   246.15n ±  3%  +1222.32% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             29.07n ± 36%   207.75n ±  9%   +614.65% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.54n ±  0%   115.95n ±  3%   +600.82% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.24n ±  0%   103.65n ±  7%   +438.72% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.59n ±  9%   112.25n ±  3%   +576.61% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.23n ±  0%   105.25n ±  6%   +447.32% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.19n ±  0%   142.20n ±  2%   +681.75% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.25n ±  0%   134.15n ±  6%   +596.88% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.16n ±  0%   141.15n ±  1%   +677.26% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.27n ± 33%   125.90n ±  7%   +553.52% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   14.88n ± 41%   162.25n ±  8%   +990.02% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   25.71n ±  0%   160.75n ±  4%   +525.24% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.32n ± 39%   175.65n ±  5%   +858.79% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    25.71n ±  0%   162.35n ±  6%   +531.47% (p=0.000 n=20)
geomean                                19.83n          152.8n         +670.46%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │         kentik-patricia/lpm.bm          │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            16.38n ±  5%   181.05n ±  9%  +1004.97% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            28.70n ± 30%   259.65n ± 22%   +804.55% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             18.61n ±  0%   143.75n ±  5%   +672.23% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             29.07n ± 36%   166.80n ± 13%   +473.79% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     16.54n ±  0%    81.81n ±  8%   +394.44% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     19.24n ±  0%    99.82n ± 16%   +418.81% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      16.59n ±  9%    68.03n ±  5%   +310.07% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      19.23n ±  0%    78.61n ± 11%   +308.81% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    18.19n ±  0%   105.90n ±  7%   +482.19% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    19.25n ±  0%   124.70n ±  3%   +547.79% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     18.16n ±  0%    93.60n ±  4%   +415.45% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.27n ± 33%   108.40n ±  4%   +462.68% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   14.88n ± 41%   134.00n ±  8%   +800.24% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   25.71n ±  0%   153.60n ±  3%   +497.43% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    18.32n ± 39%   109.45n ±  5%   +497.43% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    25.71n ±  0%   136.70n ±  6%   +431.70% (p=0.000 n=20)
geomean                                19.83n          120.7n         +508.78%
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
RandomPfx4Size/10_000        48.26 ± 0%   4060.00 ± 0%  +8312.76% (p=0.002 n=6)
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
RandomPfxSize/1_000          98.98 ± 2%   7538.00 ± 0%  +7515.68% (p=0.002 n=6)
RandomPfxSize/10_000         71.38 ± 0%   6058.00 ± 0%  +8386.97% (p=0.002 n=6)
RandomPfxSize/100_000        76.54 ± 0%   5300.00 ± 0%  +6824.48% (p=0.002 n=6)
RandomPfxSize/200_000        65.49 ± 0%   4586.00 ± 0%  +6902.60% (p=0.002 n=6)
RandomPfxSize/500_000        47.95 ± 0%
RandomPfxSize/1_000_000      39.50 ± 0%
geomean                      56.09        4.450Ki       +6731.41%               ¹
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
Tier1PfxSize/1_000           104.1 ± 2%    119.6 ± 2%   +14.89% (p=0.002 n=6)
Tier1PfxSize/10_000          76.06 ± 0%   114.80 ± 0%   +50.93% (p=0.002 n=6)
Tier1PfxSize/100_000         41.17 ± 0%   114.40 ± 0%  +177.87% (p=0.002 n=6)
Tier1PfxSize/200_000         32.85 ± 0%   114.40 ± 0%  +248.25% (p=0.002 n=6)
Tier1PfxSize/500_000         25.19 ± 0%   114.40 ± 0%  +354.15% (p=0.002 n=6)
Tier1PfxSize/1_000_000       21.39 ± 0%   114.40 ± 0%  +434.83% (p=0.002 n=6)
RandomPfx4Size/1_000         74.64 ± 3%   116.30 ± 2%   +55.81% (p=0.002 n=6)
RandomPfx4Size/10_000        48.26 ± 0%   112.40 ± 0%  +132.91% (p=0.002 n=6)
RandomPfx4Size/100_000       59.57 ± 0%   112.00 ± 0%   +88.01% (p=0.002 n=6)
RandomPfx4Size/200_000       52.51 ± 0%   112.00 ± 0%  +113.29% (p=0.002 n=6)
RandomPfx4Size/500_000       37.84 ± 0%   112.00 ± 0%  +195.98% (p=0.002 n=6)
RandomPfx4Size/1_000_000     31.83 ± 0%   112.00 ± 0%  +251.87% (p=0.002 n=6)
RandomPfx6Size/1_000         82.67 ± 2%   132.20 ± 1%   +59.91% (p=0.002 n=6)
RandomPfx6Size/10_000        98.87 ± 0%   128.40 ± 0%   +29.87% (p=0.002 n=6)
RandomPfx6Size/100_000       67.59 ± 0%   128.00 ± 0%   +89.38% (p=0.002 n=6)
RandomPfx6Size/200_000       64.96 ± 0%   128.00 ± 0%   +97.04% (p=0.002 n=6)
RandomPfx6Size/500_000       67.62 ± 0%   128.00 ± 0%   +89.29% (p=0.002 n=6)
RandomPfx6Size/1_000_000     73.74 ± 0%   128.00 ± 0%   +73.58% (p=0.002 n=6)
RandomPfxSize/1_000          98.98 ± 2%   119.60 ± 2%   +20.83% (p=0.002 n=6)
RandomPfxSize/10_000         71.38 ± 0%   115.60 ± 0%   +61.95% (p=0.002 n=6)
RandomPfxSize/100_000        76.54 ± 0%   115.30 ± 0%   +50.64% (p=0.002 n=6)
RandomPfxSize/200_000        65.49 ± 0%   115.20 ± 0%   +75.90% (p=0.002 n=6)
RandomPfxSize/500_000        47.95 ± 0%   115.20 ± 0%  +140.25% (p=0.002 n=6)
RandomPfxSize/1_000_000      39.50 ± 0%   115.20 ± 0%  +191.65% (p=0.002 n=6)
geomean                      56.09         118.1       +110.48%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │           lpmtrie/size.bm            │
                         │ bytes/route  │ bytes/route   vs base                │
Tier1PfxSize/1_000           104.1 ± 2%    215.5 ±  5%  +107.01% (p=0.002 n=6)
Tier1PfxSize/10_000          76.06 ± 0%   210.50 ±  5%  +176.76% (p=0.002 n=6)
Tier1PfxSize/100_000         41.17 ± 0%   209.90 ±  5%  +409.84% (p=0.002 n=6)
Tier1PfxSize/200_000         32.85 ± 0%   209.20 ±  5%  +536.83% (p=0.002 n=6)
Tier1PfxSize/500_000         25.19 ± 0%   207.90 ±  7%  +725.33% (p=0.002 n=6)
Tier1PfxSize/1_000_000       21.39 ± 0%   205.00 ±  7%  +858.39% (p=0.002 n=6)
RandomPfx4Size/1_000         74.64 ± 3%   205.30 ±  8%  +175.05% (p=0.002 n=6)
RandomPfx4Size/10_000        48.26 ± 0%   186.60 ±  9%  +286.66% (p=0.002 n=6)
RandomPfx4Size/100_000       59.57 ± 0%   179.60 ±  9%  +201.49% (p=0.002 n=6)
RandomPfx4Size/200_000       52.51 ± 0%   178.50 ± 10%  +239.94% (p=0.002 n=6)
RandomPfx4Size/500_000       37.84 ± 0%   176.80 ± 10%  +367.23% (p=0.002 n=6)
RandomPfx4Size/1_000_000     31.83 ± 0%   175.50 ± 10%  +451.37% (p=0.002 n=6)
RandomPfx6Size/1_000         82.67 ± 2%   227.80 ±  8%  +175.55% (p=0.002 n=6)
RandomPfx6Size/10_000        98.87 ± 0%   222.50 ±  8%  +125.04% (p=0.002 n=6)
RandomPfx6Size/100_000       67.59 ± 0%   213.90 ±  9%  +216.47% (p=0.002 n=6)
RandomPfx6Size/200_000       64.96 ± 0%   210.50 ±  9%  +224.05% (p=0.002 n=6)
RandomPfx6Size/500_000       67.62 ± 0%   206.70 ±  9%  +205.68% (p=0.002 n=6)
RandomPfx6Size/1_000_000     73.74 ± 0%   204.90 ±  9%  +177.87% (p=0.002 n=6)
RandomPfxSize/1_000          98.98 ± 2%   215.10 ±  5%  +117.32% (p=0.002 n=6)
RandomPfxSize/10_000         71.38 ± 0%   210.30 ±  5%  +194.62% (p=0.002 n=6)
RandomPfxSize/100_000        76.54 ± 0%   203.50 ±  7%  +165.87% (p=0.002 n=6)
RandomPfxSize/200_000        65.49 ± 0%   198.70 ±  8%  +203.41% (p=0.002 n=6)
RandomPfxSize/500_000        47.95 ± 0%   189.90 ±  9%  +296.04% (p=0.002 n=6)
RandomPfxSize/1_000_000      39.50 ± 0%   181.40 ± 10%  +359.24% (p=0.002 n=6)
geomean                      56.09         201.3        +258.97%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │           cidranger/size.bm            │
                         │ bytes/route  │ bytes/route  vs base                   │
Tier1PfxSize/1_000           104.1 ± 2%    539.7 ± 3%   +418.44% (p=0.002 n=6)
Tier1PfxSize/10_000          76.06 ± 0%   533.90 ± 3%   +601.95% (p=0.002 n=6)
Tier1PfxSize/100_000         41.17 ± 0%   527.20 ± 2%  +1180.54% (p=0.002 n=6)
Tier1PfxSize/200_000         32.85 ± 0%   522.20 ± 2%  +1489.65% (p=0.002 n=6)
Tier1PfxSize/500_000         25.19 ± 0%
Tier1PfxSize/1_000_000       21.39 ± 0%
RandomPfx4Size/1_000         74.64 ± 3%   481.60 ± 3%   +545.23% (p=0.002 n=6)
RandomPfx4Size/10_000        48.26 ± 0%   433.10 ± 3%   +797.43% (p=0.002 n=6)
RandomPfx4Size/100_000       59.57 ± 0%   413.70 ± 3%   +594.48% (p=0.002 n=6)
RandomPfx4Size/200_000       52.51 ± 0%   409.10 ± 3%   +679.09% (p=0.002 n=6)
RandomPfx4Size/500_000       37.84 ± 0%
RandomPfx4Size/1_000_000     31.83 ± 0%
RandomPfx6Size/1_000         82.67 ± 2%   595.30 ± 0%   +620.09% (p=0.002 n=6)
RandomPfx6Size/10_000        98.87 ± 0%   581.00 ± 0%   +487.64% (p=0.002 n=6)
RandomPfx6Size/100_000       67.59 ± 0%   547.20 ± 0%   +709.59% (p=0.002 n=6)
RandomPfx6Size/200_000       64.96 ± 0%   538.10 ± 0%   +728.36% (p=0.002 n=6)
RandomPfx6Size/500_000       67.62 ± 0%
RandomPfx6Size/1_000_000     73.74 ± 0%
RandomPfxSize/1_000          98.98 ± 2%   540.40 ± 2%   +445.97% (p=0.002 n=6)
RandomPfxSize/10_000         71.38 ± 0%   528.20 ± 2%   +639.98% (p=0.002 n=6)
RandomPfxSize/100_000        76.54 ± 0%   495.50 ± 2%   +547.37% (p=0.002 n=6)
RandomPfxSize/200_000        65.49 ± 0%   477.60 ± 2%   +629.27% (p=0.002 n=6)
RandomPfxSize/500_000        47.95 ± 0%
RandomPfxSize/1_000_000      39.50 ± 0%
geomean                      56.09         507.3        +660.66%               ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │          cidrtree/size.bm           │
                         │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000          104.10 ± 2%    69.39 ± 3%   -33.34% (p=0.002 n=6)
Tier1PfxSize/10_000          76.06 ± 0%    64.39 ± 0%   -15.34% (p=0.002 n=6)
Tier1PfxSize/100_000         41.17 ± 0%    64.04 ± 0%   +55.55% (p=0.002 n=6)
Tier1PfxSize/200_000         32.85 ± 0%    64.02 ± 0%   +94.89% (p=0.002 n=6)
Tier1PfxSize/500_000         25.19 ± 0%    64.01 ± 0%  +154.11% (p=0.002 n=6)
Tier1PfxSize/1_000_000       21.39 ± 0%    64.00 ± 0%  +199.21% (p=0.002 n=6)
RandomPfx4Size/1_000         74.64 ± 3%    68.29 ± 3%    -8.51% (p=0.002 n=6)
RandomPfx4Size/10_000        48.26 ± 0%    64.39 ± 0%   +33.42% (p=0.002 n=6)
RandomPfx4Size/100_000       59.57 ± 0%    64.04 ± 0%    +7.50% (p=0.002 n=6)
RandomPfx4Size/200_000       52.51 ± 0%    64.02 ± 0%   +21.92% (p=0.002 n=6)
RandomPfx4Size/500_000       37.84 ± 0%    64.01 ± 0%   +69.16% (p=0.002 n=6)
RandomPfx4Size/1_000_000     31.83 ± 0%    64.00 ± 0%  +101.07% (p=0.002 n=6)
RandomPfx6Size/1_000         82.67 ± 2%    68.29 ± 3%   -17.39% (p=0.002 n=6)
RandomPfx6Size/10_000        98.87 ± 0%    64.39 ± 0%   -34.87% (p=0.002 n=6)
RandomPfx6Size/100_000       67.59 ± 0%    64.04 ± 0%    -5.25% (p=0.002 n=6)
RandomPfx6Size/200_000       64.96 ± 0%    64.02 ± 0%    -1.45% (p=0.002 n=6)
RandomPfx6Size/500_000       67.62 ± 0%    64.01 ± 0%    -5.34% (p=0.002 n=6)
RandomPfx6Size/1_000_000     73.74 ± 0%    64.00 ± 0%   -13.21% (p=0.002 n=6)
RandomPfxSize/1_000          98.98 ± 2%    68.29 ± 3%   -31.01% (p=0.002 n=6)
RandomPfxSize/10_000         71.38 ± 0%    64.39 ± 0%    -9.79% (p=0.002 n=6)
RandomPfxSize/100_000        76.54 ± 0%    64.04 ± 0%   -16.33% (p=0.002 n=6)
RandomPfxSize/200_000        65.49 ± 0%    64.02 ± 0%    -2.24% (p=0.002 n=6)
RandomPfxSize/500_000        47.95 ± 0%    64.01 ± 0%   +33.49% (p=0.002 n=6)
RandomPfxSize/1_000_000      39.50 ± 0%    64.00 ± 0%   +62.03% (p=0.002 n=6)
geomean                      56.09         64.82        +15.56%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │          go-iptrie/size.bm          │
                         │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000           104.1 ± 2%    165.1 ± 1%   +58.60% (p=0.002 n=6)
Tier1PfxSize/10_000          76.06 ± 0%   159.90 ± 0%  +110.23% (p=0.002 n=6)
Tier1PfxSize/100_000         41.17 ± 0%   157.20 ± 0%  +281.83% (p=0.002 n=6)
Tier1PfxSize/200_000         32.85 ± 0%   155.40 ± 0%  +373.06% (p=0.002 n=6)
Tier1PfxSize/500_000         25.19 ± 0%   151.70 ± 0%  +502.22% (p=0.002 n=6)
Tier1PfxSize/1_000_000       21.39 ± 0%   147.80 ± 0%  +590.98% (p=0.002 n=6)
RandomPfx4Size/1_000         74.64 ± 3%   147.90 ± 1%   +98.15% (p=0.002 n=6)
RandomPfx4Size/10_000        48.26 ± 0%   127.80 ± 0%  +164.82% (p=0.002 n=6)
RandomPfx4Size/100_000       59.57 ± 0%   120.60 ± 0%  +102.45% (p=0.002 n=6)
RandomPfx4Size/200_000       52.51 ± 0%   118.90 ± 0%  +126.43% (p=0.002 n=6)
RandomPfx4Size/500_000       37.84 ± 0%   116.30 ± 0%  +207.35% (p=0.002 n=6)
RandomPfx4Size/1_000_000     31.83 ± 0%   114.10 ± 0%  +258.47% (p=0.002 n=6)
RandomPfx6Size/1_000         82.67 ± 2%   163.90 ± 1%   +98.26% (p=0.002 n=6)
RandomPfx6Size/10_000        98.87 ± 0%   156.80 ± 0%   +58.59% (p=0.002 n=6)
RandomPfx6Size/100_000       67.59 ± 0%   146.00 ± 0%  +116.01% (p=0.002 n=6)
RandomPfx6Size/200_000       64.96 ± 0%   143.20 ± 0%  +120.44% (p=0.002 n=6)
RandomPfx6Size/500_000       67.62 ± 0%   140.20 ± 0%  +107.34% (p=0.002 n=6)
RandomPfx6Size/1_000_000     73.74 ± 0%   138.70 ± 0%   +88.09% (p=0.002 n=6)
RandomPfxSize/1_000          98.98 ± 2%   163.60 ± 1%   +65.29% (p=0.002 n=6)
RandomPfxSize/10_000         71.38 ± 0%   157.00 ± 0%  +119.95% (p=0.002 n=6)
RandomPfxSize/100_000        76.54 ± 0%   145.10 ± 0%   +89.57% (p=0.002 n=6)
RandomPfxSize/200_000        65.49 ± 0%   138.70 ± 0%  +111.79% (p=0.002 n=6)
RandomPfxSize/500_000        47.95 ± 0%   128.80 ± 0%  +168.61% (p=0.002 n=6)
RandomPfxSize/1_000_000      39.50 ± 0%   120.50 ± 0%  +205.06% (p=0.002 n=6)
geomean                      56.09         141.8       +152.79%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │       kentik-patricia/size.bm       │
                         │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000           104.1 ± 2%    145.4 ± 1%   +39.67% (p=0.002 n=6)
Tier1PfxSize/10_000          76.06 ± 0%   200.20 ± 0%  +163.21% (p=0.002 n=6)
Tier1PfxSize/100_000         41.17 ± 0%   164.00 ± 0%  +298.35% (p=0.002 n=6)
Tier1PfxSize/200_000         32.85 ± 0%   164.00 ± 0%  +399.24% (p=0.002 n=6)
Tier1PfxSize/500_000         25.19 ± 0%   144.50 ± 0%  +473.64% (p=0.002 n=6)
Tier1PfxSize/1_000_000       21.39 ± 0%   144.50 ± 0%  +575.55% (p=0.002 n=6)
RandomPfx4Size/1_000         74.64 ± 3%   140.70 ± 1%   +88.50% (p=0.002 n=6)
RandomPfx4Size/10_000        48.26 ± 0%   109.60 ± 0%  +127.10% (p=0.002 n=6)
RandomPfx4Size/100_000       59.57 ± 0%   139.80 ± 0%  +134.68% (p=0.002 n=6)
RandomPfx4Size/200_000       52.51 ± 0%   139.80 ± 0%  +166.24% (p=0.002 n=6)
RandomPfx4Size/500_000       37.84 ± 0%   139.70 ± 0%  +269.19% (p=0.002 n=6)
RandomPfx4Size/1_000_000     31.83 ± 0%   139.80 ± 0%  +339.21% (p=0.002 n=6)
RandomPfx6Size/1_000         82.67 ± 2%   157.10 ± 1%   +90.03% (p=0.002 n=6)
RandomPfx6Size/10_000        98.87 ± 0%   201.30 ± 0%  +103.60% (p=0.002 n=6)
RandomPfx6Size/100_000       67.59 ± 0%   160.80 ± 0%  +137.91% (p=0.002 n=6)
RandomPfx6Size/200_000       64.96 ± 0%   160.80 ± 0%  +147.54% (p=0.002 n=6)
RandomPfx6Size/500_000       67.62 ± 0%   156.30 ± 0%  +131.14% (p=0.002 n=6)
RandomPfx6Size/1_000_000     73.74 ± 0%   156.50 ± 0%  +112.23% (p=0.002 n=6)
RandomPfxSize/1_000          98.98 ± 2%   144.50 ± 1%   +45.99% (p=0.002 n=6)
RandomPfxSize/10_000         71.38 ± 0%   140.20 ± 0%   +96.41% (p=0.002 n=6)
RandomPfxSize/100_000        76.54 ± 0%   180.00 ± 0%  +135.17% (p=0.002 n=6)
RandomPfxSize/200_000        65.49 ± 0%   180.00 ± 0%  +174.85% (p=0.002 n=6)
RandomPfxSize/500_000        47.95 ± 0%   144.00 ± 0%  +200.31% (p=0.002 n=6)
RandomPfxSize/1_000_000      39.50 ± 0%   144.00 ± 0%  +264.56% (p=0.002 n=6)
geomean                      56.09         152.8       +172.40%
```

## update, insert/delete

`bart` is by far the fastest algorithm for updates.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │             art/update.bm              │
                         │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000       245.6n ±  5%   2941.5n ± 11%  +1097.92% (p=0.002 n=6)
InsertRandomPfxs/10_000      177.8n ±  3%   2290.0n ±  2%  +1188.33% (p=0.002 n=6)
InsertRandomPfxs/100_000     266.4n ±  1%   2944.0n ± 72%  +1004.90% (p=0.002 n=6)
InsertRandomPfxs/200_000     269.7n ±  3%   2983.0n ±  7%  +1006.25% (p=0.002 n=6)
DeleteRandomPfxs/1_000       386.2n ±  4%   3367.5n ±  5%   +771.84% (p=0.002 n=6)
DeleteRandomPfxs/10_000      248.6n ±  3%   3266.5n ±  3%  +1214.22% (p=0.002 n=6)
DeleteRandomPfxs/100_000     454.9n ± 12%   4125.5n ± 27%   +806.80% (p=0.002 n=6)
DeleteRandomPfxs/200_000     487.6n ± 14%   3924.5n ± 12%   +704.78% (p=0.002 n=6)
geomean                      300.7n          3.183µ         +958.62%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │          netipds/update.bm           │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000       245.6n ±  5%   268.0n ±  4%    +9.14% (p=0.002 n=6)
InsertRandomPfxs/10_000      177.8n ±  3%   324.4n ±  4%   +82.50% (p=0.002 n=6)
InsertRandomPfxs/100_000     266.4n ±  1%   426.2n ±  3%   +59.95% (p=0.002 n=6)
InsertRandomPfxs/200_000     269.7n ±  3%   538.1n ±  6%   +99.54% (p=0.002 n=6)
DeleteRandomPfxs/1_000       386.2n ±  4%   402.7n ±  6%    +4.26% (p=0.002 n=6)
DeleteRandomPfxs/10_000      248.6n ±  3%   546.6n ± 11%  +119.90% (p=0.002 n=6)
DeleteRandomPfxs/100_000     454.9n ± 12%   689.6n ± 20%   +51.59% (p=0.002 n=6)
DeleteRandomPfxs/200_000     487.6n ± 14%   853.5n ±  5%   +75.01% (p=0.002 n=6)
geomean                      300.7n         474.8n         +57.91%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │         critbitgo/update.bm          │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000       245.6n ±  5%    357.4n ± 2%   +45.57% (p=0.002 n=6)
InsertRandomPfxs/10_000      177.8n ±  3%    446.1n ± 4%  +150.97% (p=0.002 n=6)
InsertRandomPfxs/100_000     266.4n ±  1%    606.2n ± 2%  +127.49% (p=0.002 n=6)
InsertRandomPfxs/200_000     269.7n ±  3%    797.6n ± 3%  +195.81% (p=0.002 n=6)
DeleteRandomPfxs/1_000       386.2n ±  4%    511.6n ± 4%   +32.47% (p=0.002 n=6)
DeleteRandomPfxs/10_000      248.6n ±  3%    618.1n ± 1%  +148.70% (p=0.002 n=6)
DeleteRandomPfxs/100_000     454.9n ± 12%   1056.0n ± 2%  +132.11% (p=0.002 n=6)
DeleteRandomPfxs/200_000     487.6n ± 14%   1268.5n ± 4%  +160.13% (p=0.002 n=6)
geomean                      300.7n          652.0n       +116.85%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │          lpmtrie/update.bm           │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000       245.6n ±  5%    538.3n ± 3%  +119.22% (p=0.002 n=6)
InsertRandomPfxs/10_000      177.8n ±  3%    617.9n ± 1%  +247.65% (p=0.002 n=6)
InsertRandomPfxs/100_000     266.4n ±  1%    835.8n ± 3%  +213.70% (p=0.002 n=6)
InsertRandomPfxs/200_000     269.7n ±  3%   1038.5n ± 4%  +285.13% (p=0.002 n=6)
DeleteRandomPfxs/1_000       386.2n ±  4%    687.0n ± 5%   +77.85% (p=0.002 n=6)
DeleteRandomPfxs/10_000      248.6n ±  3%    845.4n ± 6%  +240.11% (p=0.002 n=6)
DeleteRandomPfxs/100_000     454.9n ± 12%   1526.5n ± 3%  +235.53% (p=0.002 n=6)
DeleteRandomPfxs/200_000     487.6n ± 14%   1709.0n ± 1%  +250.46% (p=0.002 n=6)
geomean                      300.7n          901.8n       +199.93%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │          cidranger/update.bm           │
                         │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000       245.6n ±  5%    4975.0n ± 1%  +1926.06% (p=0.002 n=6)
InsertRandomPfxs/10_000      177.8n ±  3%    6117.0n ± 2%  +3341.35% (p=0.002 n=6)
InsertRandomPfxs/100_000     266.4n ±  1%    9244.0n ± 3%  +3369.32% (p=0.002 n=6)
InsertRandomPfxs/200_000     269.7n ±  3%    9906.5n ± 4%  +3573.84% (p=0.002 n=6)
DeleteRandomPfxs/1_000       386.2n ±  4%    5556.5n ± 4%  +1338.58% (p=0.002 n=6)
DeleteRandomPfxs/10_000      248.6n ±  3%    7282.5n ± 4%  +2829.99% (p=0.002 n=6)
DeleteRandomPfxs/100_000     454.9n ± 12%   11649.0n ± 2%  +2460.50% (p=0.002 n=6)
DeleteRandomPfxs/200_000     487.6n ± 14%   13645.0n ± 2%  +2698.11% (p=0.002 n=6)
geomean                      300.7n           8.066µ       +2582.82%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │          cidrtree/update.bm           │
                         │   sec/route    │  sec/route    vs base                 │
InsertRandomPfxs/1_000       245.6n ±  5%    988.2n ± 0%   +302.46% (p=0.002 n=6)
InsertRandomPfxs/10_000      177.8n ±  3%   1523.0n ± 1%   +756.82% (p=0.002 n=6)
InsertRandomPfxs/100_000     266.4n ±  1%   2315.0n ± 2%   +768.83% (p=0.002 n=6)
InsertRandomPfxs/200_000     269.7n ±  3%   2868.5n ± 2%   +963.79% (p=0.002 n=6)
DeleteRandomPfxs/1_000       386.2n ±  4%   2406.0n ± 1%   +522.91% (p=0.002 n=6)
DeleteRandomPfxs/10_000      248.6n ±  3%   3888.5n ± 1%  +1464.47% (p=0.002 n=6)
DeleteRandomPfxs/100_000     454.9n ± 12%   5959.0n ± 2%  +1209.81% (p=0.002 n=6)
DeleteRandomPfxs/200_000     487.6n ± 14%   7220.5n ± 2%  +1380.67% (p=0.002 n=6)
geomean                      300.7n          2.822µ        +838.61%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │         go-iptrie/update.bm          │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000       245.6n ±  5%    316.5n ± 3%   +28.89% (p=0.002 n=6)
InsertRandomPfxs/10_000      177.8n ±  3%    375.5n ± 1%  +111.22% (p=0.002 n=6)
InsertRandomPfxs/100_000     266.4n ±  1%    569.9n ± 3%  +113.91% (p=0.002 n=6)
InsertRandomPfxs/200_000     269.7n ±  3%    660.1n ± 3%  +144.82% (p=0.002 n=6)
DeleteRandomPfxs/1_000       386.2n ±  4%    481.8n ± 2%   +24.72% (p=0.002 n=6)
DeleteRandomPfxs/10_000      248.6n ±  3%    609.0n ± 4%  +145.04% (p=0.002 n=6)
DeleteRandomPfxs/100_000     454.9n ± 12%   1034.5n ± 3%  +127.39% (p=0.002 n=6)
DeleteRandomPfxs/200_000     487.6n ± 14%   1351.0n ± 3%  +177.04% (p=0.002 n=6)
geomean                      300.7n          606.6n       +101.75%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │      kentik-patricia/update.bm       │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000       245.6n ±  5%    331.4n ± 1%   +34.96% (p=0.002 n=6)
InsertRandomPfxs/10_000      177.8n ±  3%    430.6n ± 1%  +142.25% (p=0.002 n=6)
InsertRandomPfxs/100_000     266.4n ±  1%    636.0n ± 5%  +138.71% (p=0.002 n=6)
InsertRandomPfxs/200_000     269.7n ±  3%    761.0n ± 2%  +182.22% (p=0.002 n=6)
DeleteRandomPfxs/1_000       386.2n ±  4%    721.2n ± 2%   +86.72% (p=0.002 n=6)
DeleteRandomPfxs/10_000      248.6n ±  3%    916.6n ± 2%  +268.80% (p=0.002 n=6)
DeleteRandomPfxs/100_000     454.9n ± 12%   1423.0n ± 4%  +212.78% (p=0.002 n=6)
DeleteRandomPfxs/200_000     487.6n ± 14%   1714.0n ± 2%  +251.48% (p=0.002 n=6)
geomean                      300.7n          760.1n       +152.79%
```
