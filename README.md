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
LpmTier1Pfxs/RandomMatchIP4-8            15.82n ± 10%   48.84n ± 13%  +208.63% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            25.40n ± 11%   66.89n ± 11%  +163.33% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             17.77n ±  2%   30.50n ±  1%   +71.69% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             27.56n ± 26%   38.17n ± 22%   +38.50% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     18.63n ±  3%   47.63n ±  2%  +155.69% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     21.69n ±  3%   50.71n ±  5%  +133.87% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      18.82n ± 21%   31.54n ±  2%   +67.52% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      20.67n ±  2%   32.01n ±  5%   +54.82% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    19.87n ±  2%   47.53n ±  3%  +139.20% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    21.26n ±  2%   49.87n ±  3%  +134.55% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     19.72n ±  2%   32.83n ±  3%   +66.41% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     21.46n ±  3%   31.41n ±  7%   +46.31% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   12.94n ± 34%   50.15n ± 11%  +287.56% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   27.63n ± 43%   51.72n ±  5%   +87.19% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    19.55n ±  5%   31.13n ±  6%   +59.19% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    27.30n ±  3%   36.05n ±  3%   +32.03% (p=0.000 n=20)
geomean                                  20.61n         41.09n         +99.39%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ bart/lpm.bm  │            netipds/lpm.bm             │
                                       │    sec/op    │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4-8            15.82n ± 10%   53.63n ± 17%  +238.89% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            25.40n ± 11%   49.95n ± 19%   +96.67% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             17.77n ±  2%   57.33n ±  5%  +222.71% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             27.56n ± 26%   56.86n ± 14%  +106.31% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     18.63n ±  3%   34.27n ±  9%   +83.95% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     21.69n ±  3%   32.48n ±  4%   +49.80% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      18.82n ± 21%   33.92n ± 11%   +80.19% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      20.67n ±  2%   37.94n ±  7%   +83.51% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    19.87n ±  2%   43.61n ±  8%  +119.48% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    21.26n ±  2%   40.44n ±  6%   +90.19% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     19.72n ±  2%   46.58n ±  8%  +136.12% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     21.46n ±  3%   40.08n ±  5%   +86.72% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   12.94n ± 34%   39.62n ± 12%  +206.18% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   27.63n ± 43%   42.75n ±  4%   +54.74% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    19.55n ±  5%   49.51n ± 10%  +153.18% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    27.30n ±  3%   48.95n ±  6%   +79.29% (p=0.000 n=20)
geomean                                  20.61n         43.58n        +111.48%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ bart/lpm.bm  │            critbitgo/lpm.bm             │
                                       │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4-8            15.82n ± 10%   181.15n ±  5%  +1044.71% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            25.40n ± 11%   225.50n ±  9%   +787.80% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             17.77n ±  2%   760.95n ± 18%  +4183.42% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             27.56n ± 26%   558.60n ± 15%  +1926.85% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     18.63n ±  3%   100.05n ±  4%   +437.01% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     21.69n ±  3%   112.00n ±  6%   +416.49% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      18.82n ± 21%   171.30n ± 14%   +809.96% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      20.67n ±  2%   164.60n ±  8%   +696.13% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    19.87n ±  2%   106.70n ±  5%   +436.99% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    21.26n ±  2%   130.85n ±  5%   +515.48% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     19.72n ±  2%   266.35n ± 12%  +1250.32% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     21.46n ±  3%   240.45n ± 14%  +1020.20% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   12.94n ± 34%   130.70n ±  5%   +910.05% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   27.63n ± 43%   136.75n ±  1%   +394.93% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    19.55n ±  5%   482.80n ± 15%  +2368.93% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    27.30n ±  3%   337.85n ± 21%  +1137.32% (p=0.000 n=20)
geomean                                  20.61n          210.0n         +919.10%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ bart/lpm.bm  │             lpmtrie/lpm.bm              │
                                       │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4-8            15.82n ± 10%   269.10n ±  5%  +1600.47% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            25.40n ± 11%   323.70n ± 12%  +1174.41% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             17.77n ±  2%   258.00n ±  5%  +1352.29% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             27.56n ± 26%   236.55n ± 14%   +758.31% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     18.63n ±  3%   131.70n ±  6%   +606.92% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     21.69n ±  3%   123.50n ± 25%   +469.52% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      18.82n ± 21%   116.30n ± 17%   +517.80% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      20.67n ±  2%    94.97n ± 39%   +359.32% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    19.87n ±  2%   152.50n ± 14%   +667.49% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    21.26n ±  2%   139.55n ±  6%   +556.40% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     19.72n ±  2%   163.85n ±  9%   +730.67% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     21.46n ±  3%   135.00n ±  5%   +528.93% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   12.94n ± 34%   198.00n ±  6%  +1430.14% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   27.63n ± 43%   187.15n ±  2%   +577.34% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    19.55n ±  5%   197.10n ±  3%   +907.93% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    27.30n ±  3%   192.35n ±  4%   +604.45% (p=0.000 n=20)
geomean                                  20.61n          172.8n         +738.56%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ bart/lpm.bm  │            cidranger/lpm.bm             │
                                       │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4-8            15.82n ± 10%   258.30n ± 25%  +1532.23% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            25.40n ± 11%   274.85n ± 26%   +982.09% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             17.77n ±  2%   294.25n ± 10%  +1556.35% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             27.56n ± 26%   288.10n ± 17%   +945.36% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     18.63n ±  3%   118.55n ±  7%   +536.34% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     21.69n ±  3%   147.15n ± 10%   +578.58% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      18.82n ± 21%   115.60n ±  4%   +514.08% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      20.67n ±  2%   142.85n ± 12%   +590.93% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    19.87n ±  2%   153.40n ±  8%   +672.02% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    21.26n ±  2%   176.70n ± 14%   +731.14% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     19.72n ±  2%   152.65n ±  5%   +673.89% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     21.46n ±  3%   179.95n ±  6%   +738.34% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   12.94n ± 34%   156.85n ± 23%  +1112.13% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   27.63n ± 43%   223.70n ±  7%   +709.63% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    19.55n ±  5%   189.35n ±  6%   +868.29% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    27.30n ±  3%   224.55n ±  7%   +722.38% (p=0.000 n=20)
geomean                                  20.61n          185.3n         +799.07%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ bart/lpm.bm  │             cidrtree/lpm.bm              │
                                       │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4-8            15.82n ± 10%    741.15n ± 14%  +4583.41% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            25.40n ± 11%    861.80n ± 15%  +3292.91% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             17.77n ±  2%   1050.00n ± 20%  +5810.50% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             27.56n ± 26%   1191.00n ± 21%  +4221.48% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     18.63n ±  3%    351.25n ± 18%  +1785.40% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     21.69n ±  3%    370.20n ± 37%  +1607.17% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      18.82n ± 21%    448.10n ± 27%  +2280.35% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      20.67n ±  2%    603.30n ± 15%  +2818.02% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    19.87n ±  2%    425.65n ± 18%  +2042.17% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    21.26n ±  2%    497.20n ± 22%  +2238.66% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     19.72n ±  2%    627.60n ± 26%  +3081.75% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     21.46n ±  3%    873.75n ± 23%  +3970.58% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   12.94n ± 34%    620.25n ± 18%  +4693.28% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   27.63n ± 43%    678.90n ± 25%  +2357.11% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    19.55n ±  5%    766.90n ± 24%  +3821.76% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    27.30n ±  3%   1058.00n ± 14%  +3774.75% (p=0.000 n=20)
geomean                                  20.61n           653.8n        +3072.64%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ bart/lpm.bm  │            go-iptrie/lpm.bm             │
                                       │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4-8            15.82n ± 10%   196.95n ± 16%  +1144.55% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            25.40n ± 11%   183.50n ± 28%   +622.44% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             17.77n ±  2%   208.60n ±  4%  +1074.22% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             27.56n ± 26%   196.60n ± 11%   +613.35% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     18.63n ±  3%   114.95n ±  8%   +517.02% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     21.69n ±  3%   108.50n ± 18%   +400.35% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      18.82n ± 21%   111.95n ±  4%   +494.69% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      20.67n ±  2%   100.30n ±  9%   +385.13% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    19.87n ±  2%   141.40n ± 11%   +611.63% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    21.26n ±  2%   135.90n ±  5%   +539.23% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     19.72n ±  2%   145.30n ±  3%   +636.63% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     21.46n ±  3%   132.00n ±  4%   +514.95% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   12.94n ± 34%   148.85n ± 13%  +1050.31% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   27.63n ± 43%   162.55n ±  1%   +488.31% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    19.55n ±  5%   165.85n ±  5%   +748.12% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    27.30n ±  3%   172.80n ±  5%   +532.85% (p=0.000 n=20)
geomean                                  20.61n          148.0n         +618.29%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                       │ bart/lpm.bm  │         kentik-patricia/lpm.bm          │
                                       │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4-8            15.82n ± 10%   175.05n ±  9%  +1006.16% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6-8            25.40n ± 11%   263.20n ± 23%   +936.22% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4-8             17.77n ±  2%   138.95n ±  5%   +682.16% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6-8             27.56n ± 26%   166.55n ± 13%   +504.32% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4-8     18.63n ±  3%    80.51n ±  8%   +332.15% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6-8     21.69n ±  3%    99.72n ± 16%   +359.88% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4-8      18.82n ± 21%    66.33n ±  5%   +252.32% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6-8      20.67n ±  2%    79.01n ± 11%   +282.18% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4-8    19.87n ±  2%   103.50n ±  5%   +420.89% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6-8    21.26n ±  2%   124.55n ±  4%   +485.84% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4-8     19.72n ±  2%    91.57n ±  3%   +364.21% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6-8     21.46n ±  3%   108.40n ±  4%   +405.01% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4-8   12.94n ± 34%   131.55n ±  9%   +916.62% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6-8   27.63n ± 43%   153.40n ±  5%   +455.19% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4-8    19.55n ±  5%   107.40n ±  6%   +449.22% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6-8    27.30n ±  3%   137.30n ±  6%   +402.84% (p=0.000 n=20)
geomean                                  20.61n          119.4n         +479.48%
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
Tier1PfxSize/1_000-8           104.0 ± 2%    7591.0 ± 0%  +7199.04% (p=0.002 n=6)
Tier1PfxSize/10_000-8          75.97 ± 0%   4889.00 ± 0%  +6335.44% (p=0.002 n=6)
Tier1PfxSize/100_000-8         41.16 ± 0%   1669.00 ± 0%  +3954.91% (p=0.002 n=6)
Tier1PfxSize/200_000-8         32.85 ± 0%   1098.00 ± 0%  +3242.47% (p=0.002 n=6)
Tier1PfxSize/500_000-8         25.19 ± 0%
Tier1PfxSize/1_000_000-8       21.39 ± 0%
RandomPfx4Size/1_000-8         74.74 ± 3%   5259.00 ± 0%  +6936.39% (p=0.002 n=6)
RandomPfx4Size/10_000-8        48.22 ± 0%   4059.00 ± 0%  +8317.67% (p=0.002 n=6)
RandomPfx4Size/100_000-8       59.57 ± 0%   3938.00 ± 0%  +6510.71% (p=0.002 n=6)
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
geomean                        56.08        4.449Ki       +6732.27%               ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/size.bm │           netipds/size.bm           │
                           │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000-8          104.00 ± 2%    73.95 ± 3%   -28.89% (p=0.002 n=6)
Tier1PfxSize/10_000-8          75.97 ± 0%    69.43 ± 0%    -8.61% (p=0.002 n=6)
Tier1PfxSize/100_000-8         41.16 ± 0%    67.64 ± 0%   +64.33% (p=0.002 n=6)
Tier1PfxSize/200_000-8         32.85 ± 0%    66.90 ± 0%  +103.65% (p=0.002 n=6)
Tier1PfxSize/500_000-8         25.19 ± 0%    65.34 ± 0%  +159.39% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       21.39 ± 0%    63.72 ± 0%  +197.90% (p=0.002 n=6)
RandomPfx4Size/1_000-8         74.74 ± 3%    61.79 ± 3%   -17.33% (p=0.002 n=6)
RandomPfx4Size/10_000-8        48.22 ± 0%    51.25 ± 0%    +6.28% (p=0.002 n=6)
RandomPfx4Size/100_000-8       59.57 ± 0%    48.24 ± 0%   -19.02% (p=0.002 n=6)
RandomPfx4Size/200_000-8       52.50 ± 0%    47.57 ± 0%    -9.39% (p=0.002 n=6)
RandomPfx4Size/500_000-8       37.84 ± 0%    46.51 ± 0%   +22.91% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     31.83 ± 0%    45.65 ± 0%   +43.42% (p=0.002 n=6)
RandomPfx6Size/1_000-8         82.77 ± 2%   100.70 ± 2%   +21.66% (p=0.002 n=6)
RandomPfx6Size/10_000-8        98.78 ± 0%    94.17 ± 0%    -4.67% (p=0.002 n=6)
RandomPfx6Size/100_000-8       67.59 ± 0%    87.63 ± 0%   +29.65% (p=0.002 n=6)
RandomPfx6Size/200_000-8       64.96 ± 0%    85.90 ± 0%   +32.24% (p=0.002 n=6)
RandomPfx6Size/500_000-8       67.61 ± 0%    84.11 ± 0%   +24.40% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     73.74 ± 0%    83.23 ± 0%   +12.87% (p=0.002 n=6)
RandomPfxSize/1_000-8          99.07 ± 2%    75.34 ± 3%   -23.95% (p=0.002 n=6)
RandomPfxSize/10_000-8         71.29 ± 0%    69.29 ± 0%    -2.81% (p=0.002 n=6)
RandomPfxSize/100_000-8        76.53 ± 0%    64.28 ± 0%   -16.01% (p=0.002 n=6)
RandomPfxSize/200_000-8        65.49 ± 0%    61.55 ± 0%    -6.02% (p=0.002 n=6)
RandomPfxSize/500_000-8        47.94 ± 0%    57.42 ± 0%   +19.77% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      39.50 ± 0%    53.93 ± 0%   +36.53% (p=0.002 n=6)
geomean                        56.08         66.06        +17.79%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/size.bm │          critbitgo/size.bm          │
                           │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000-8           104.0 ± 2%    119.6 ± 2%   +15.00% (p=0.002 n=6)
Tier1PfxSize/10_000-8          75.97 ± 0%   114.70 ± 0%   +50.98% (p=0.002 n=6)
Tier1PfxSize/100_000-8         41.16 ± 0%   114.40 ± 0%  +177.94% (p=0.002 n=6)
Tier1PfxSize/200_000-8         32.85 ± 0%   114.40 ± 0%  +248.25% (p=0.002 n=6)
Tier1PfxSize/500_000-8         25.19 ± 0%   114.40 ± 0%  +354.15% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       21.39 ± 0%   114.40 ± 0%  +434.83% (p=0.002 n=6)
RandomPfx4Size/1_000-8         74.74 ± 3%   116.40 ± 2%   +55.74% (p=0.002 n=6)
RandomPfx4Size/10_000-8        48.22 ± 0%   112.30 ± 0%  +132.89% (p=0.002 n=6)
RandomPfx4Size/100_000-8       59.57 ± 0%   112.00 ± 0%   +88.01% (p=0.002 n=6)
RandomPfx4Size/200_000-8       52.50 ± 0%   112.00 ± 0%  +113.33% (p=0.002 n=6)
RandomPfx4Size/500_000-8       37.84 ± 0%   112.00 ± 0%  +195.98% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     31.83 ± 0%   112.00 ± 0%  +251.87% (p=0.002 n=6)
RandomPfx6Size/1_000-8         82.77 ± 2%   132.90 ± 2%   +60.57% (p=0.002 n=6)
RandomPfx6Size/10_000-8        98.78 ± 0%   128.30 ± 0%   +29.88% (p=0.002 n=6)
RandomPfx6Size/100_000-8       67.59 ± 0%   128.00 ± 0%   +89.38% (p=0.002 n=6)
RandomPfx6Size/200_000-8       64.96 ± 0%   128.00 ± 0%   +97.04% (p=0.002 n=6)
RandomPfx6Size/500_000-8       67.61 ± 0%   128.00 ± 0%   +89.32% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     73.74 ± 0%   128.00 ± 0%   +73.58% (p=0.002 n=6)
RandomPfxSize/1_000-8          99.07 ± 2%   119.70 ± 2%   +20.82% (p=0.002 n=6)
RandomPfxSize/10_000-8         71.29 ± 0%   115.50 ± 0%   +62.01% (p=0.002 n=6)
RandomPfxSize/100_000-8        76.53 ± 0%   115.30 ± 0%   +50.66% (p=0.002 n=6)
RandomPfxSize/200_000-8        65.49 ± 0%   115.20 ± 0%   +75.90% (p=0.002 n=6)
RandomPfxSize/500_000-8        47.94 ± 0%   115.20 ± 0%  +140.30% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      39.50 ± 0%   115.20 ± 0%  +191.65% (p=0.002 n=6)
geomean                        56.08         118.1       +110.53%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/size.bm │           lpmtrie/size.bm            │
                           │ bytes/route  │ bytes/route   vs base                │
Tier1PfxSize/1_000-8           104.0 ± 2%    215.4 ±  5%  +107.12% (p=0.002 n=6)
Tier1PfxSize/10_000-8          75.97 ± 0%   210.40 ±  5%  +176.95% (p=0.002 n=6)
Tier1PfxSize/100_000-8         41.16 ± 0%   209.90 ±  5%  +409.96% (p=0.002 n=6)
Tier1PfxSize/200_000-8         32.85 ± 0%   209.20 ±  5%  +536.83% (p=0.002 n=6)
Tier1PfxSize/500_000-8         25.19 ± 0%   207.90 ±  7%  +725.33% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       21.39 ± 0%   205.00 ±  7%  +858.39% (p=0.002 n=6)
RandomPfx4Size/1_000-8         74.74 ± 3%   205.40 ±  8%  +174.82% (p=0.002 n=6)
RandomPfx4Size/10_000-8        48.22 ± 0%   186.50 ±  9%  +286.77% (p=0.002 n=6)
RandomPfx4Size/100_000-8       59.57 ± 0%   179.60 ±  9%  +201.49% (p=0.002 n=6)
RandomPfx4Size/200_000-8       52.50 ± 0%   178.50 ± 10%  +240.00% (p=0.002 n=6)
RandomPfx4Size/500_000-8       37.84 ± 0%   176.80 ± 10%  +367.23% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     31.83 ± 0%   175.50 ± 10%  +451.37% (p=0.002 n=6)
RandomPfx6Size/1_000-8         82.77 ± 2%   228.40 ±  8%  +175.95% (p=0.002 n=6)
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
geomean                        56.08         201.4        +259.03%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/size.bm │           cidranger/size.bm            │
                           │ bytes/route  │ bytes/route  vs base                   │
Tier1PfxSize/1_000-8           104.0 ± 2%    539.7 ± 3%   +418.94% (p=0.002 n=6)
Tier1PfxSize/10_000-8          75.97 ± 0%   533.80 ± 3%   +602.65% (p=0.002 n=6)
Tier1PfxSize/100_000-8         41.16 ± 0%   527.20 ± 2%  +1180.86% (p=0.002 n=6)
Tier1PfxSize/200_000-8         32.85 ± 0%   522.20 ± 2%  +1489.65% (p=0.002 n=6)
Tier1PfxSize/500_000-8         25.19 ± 0%
Tier1PfxSize/1_000_000-8       21.39 ± 0%
RandomPfx4Size/1_000-8         74.74 ± 3%   482.30 ± 3%   +545.30% (p=0.002 n=6)
RandomPfx4Size/10_000-8        48.22 ± 0%   433.00 ± 3%   +797.97% (p=0.002 n=6)
RandomPfx4Size/100_000-8       59.57 ± 0%   413.70 ± 3%   +594.48% (p=0.002 n=6)
RandomPfx4Size/200_000-8       52.50 ± 0%   409.10 ± 3%   +679.24% (p=0.002 n=6)
RandomPfx4Size/500_000-8       37.84 ± 0%
RandomPfx4Size/1_000_000-8     31.83 ± 0%
RandomPfx6Size/1_000-8         82.77 ± 2%   595.40 ± 0%   +619.34% (p=0.002 n=6)
RandomPfx6Size/10_000-8        98.78 ± 0%   581.00 ± 0%   +488.18% (p=0.002 n=6)
RandomPfx6Size/100_000-8       67.59 ± 0%   547.20 ± 0%   +709.59% (p=0.002 n=6)
RandomPfx6Size/200_000-8       64.96 ± 0%   538.10 ± 0%   +728.36% (p=0.002 n=6)
RandomPfx6Size/500_000-8       67.61 ± 0%
RandomPfx6Size/1_000_000-8     73.74 ± 0%
RandomPfxSize/1_000-8          99.07 ± 2%   540.50 ± 2%   +445.57% (p=0.002 n=6)
RandomPfxSize/10_000-8         71.29 ± 0%   528.10 ± 2%   +640.78% (p=0.002 n=6)
RandomPfxSize/100_000-8        76.53 ± 0%   495.50 ± 2%   +547.46% (p=0.002 n=6)
RandomPfxSize/200_000-8        65.49 ± 0%   477.60 ± 2%   +629.27% (p=0.002 n=6)
RandomPfxSize/500_000-8        47.94 ± 0%
RandomPfxSize/1_000_000-8      39.50 ± 0%
geomean                        56.08         507.4        +660.83%               ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/size.bm │          cidrtree/size.bm           │
                           │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000-8          104.00 ± 2%    69.29 ± 3%   -33.38% (p=0.002 n=6)
Tier1PfxSize/10_000-8          75.97 ± 0%    64.35 ± 0%   -15.30% (p=0.002 n=6)
Tier1PfxSize/100_000-8         41.16 ± 0%    64.03 ± 0%   +55.56% (p=0.002 n=6)
Tier1PfxSize/200_000-8         32.85 ± 0%    64.01 ± 0%   +94.86% (p=0.002 n=6)
Tier1PfxSize/500_000-8         25.19 ± 0%    64.01 ± 0%  +154.11% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       21.39 ± 0%    64.00 ± 0%  +199.21% (p=0.002 n=6)
RandomPfx4Size/1_000-8         74.74 ± 3%    68.38 ± 3%    -8.51% (p=0.002 n=6)
RandomPfx4Size/10_000-8        48.22 ± 0%    64.29 ± 0%   +33.33% (p=0.002 n=6)
RandomPfx4Size/100_000-8       59.57 ± 0%    64.04 ± 0%    +7.50% (p=0.002 n=6)
RandomPfx4Size/200_000-8       52.50 ± 0%    64.02 ± 0%   +21.94% (p=0.002 n=6)
RandomPfx4Size/500_000-8       37.84 ± 0%    64.01 ± 0%   +69.16% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     31.83 ± 0%    64.00 ± 0%  +101.07% (p=0.002 n=6)
RandomPfx6Size/1_000-8         82.77 ± 2%    68.93 ± 3%   -16.72% (p=0.002 n=6)
RandomPfx6Size/10_000-8        98.78 ± 0%    64.35 ± 0%   -34.86% (p=0.002 n=6)
RandomPfx6Size/100_000-8       67.59 ± 0%    64.03 ± 0%    -5.27% (p=0.002 n=6)
RandomPfx6Size/200_000-8       64.96 ± 0%    64.02 ± 0%    -1.45% (p=0.002 n=6)
RandomPfx6Size/500_000-8       67.61 ± 0%    64.01 ± 0%    -5.32% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     73.74 ± 0%    64.00 ± 0%   -13.21% (p=0.002 n=6)
RandomPfxSize/1_000-8          99.07 ± 2%    68.93 ± 3%   -30.42% (p=0.002 n=6)
RandomPfxSize/10_000-8         71.29 ± 0%    64.35 ± 0%    -9.73% (p=0.002 n=6)
RandomPfxSize/100_000-8        76.53 ± 0%    64.03 ± 0%   -16.33% (p=0.002 n=6)
RandomPfxSize/200_000-8        65.49 ± 0%    64.01 ± 0%    -2.26% (p=0.002 n=6)
RandomPfxSize/500_000-8        47.94 ± 0%    64.01 ± 0%   +33.52% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      39.50 ± 0%    64.00 ± 0%   +62.03% (p=0.002 n=6)
geomean                        56.08         64.86        +15.64%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/size.bm │          go-iptrie/size.bm          │
                           │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000-8           104.0 ± 2%    165.1 ± 1%   +58.80% (p=0.002 n=6)
Tier1PfxSize/10_000-8          75.97 ± 0%   159.80 ± 0%  +110.35% (p=0.002 n=6)
Tier1PfxSize/100_000-8         41.16 ± 0%   157.20 ± 0%  +281.92% (p=0.002 n=6)
Tier1PfxSize/200_000-8         32.85 ± 0%   155.40 ± 0%  +373.06% (p=0.002 n=6)
Tier1PfxSize/500_000-8         25.19 ± 0%   151.70 ± 0%  +502.22% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       21.39 ± 0%   147.80 ± 0%  +590.98% (p=0.002 n=6)
RandomPfx4Size/1_000-8         74.74 ± 3%   148.00 ± 1%   +98.02% (p=0.002 n=6)
RandomPfx4Size/10_000-8        48.22 ± 0%   127.80 ± 0%  +165.04% (p=0.002 n=6)
RandomPfx4Size/100_000-8       59.57 ± 0%   120.50 ± 0%  +102.28% (p=0.002 n=6)
RandomPfx4Size/200_000-8       52.50 ± 0%   118.90 ± 0%  +126.48% (p=0.002 n=6)
RandomPfx4Size/500_000-8       37.84 ± 0%   116.30 ± 0%  +207.35% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     31.83 ± 0%   114.10 ± 0%  +258.47% (p=0.002 n=6)
RandomPfx6Size/1_000-8         82.77 ± 2%   164.50 ± 1%   +98.74% (p=0.002 n=6)
RandomPfx6Size/10_000-8        98.78 ± 0%   156.70 ± 0%   +58.64% (p=0.002 n=6)
RandomPfx6Size/100_000-8       67.59 ± 0%   146.00 ± 0%  +116.01% (p=0.002 n=6)
RandomPfx6Size/200_000-8       64.96 ± 0%   143.20 ± 0%  +120.44% (p=0.002 n=6)
RandomPfx6Size/500_000-8       67.61 ± 0%   140.20 ± 0%  +107.37% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     73.74 ± 0%   138.70 ± 0%   +88.09% (p=0.002 n=6)
RandomPfxSize/1_000-8          99.07 ± 2%   163.70 ± 1%   +65.24% (p=0.002 n=6)
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
Tier1PfxSize/1_000-8           104.0 ± 2%    145.4 ± 1%   +39.81% (p=0.002 n=6)
Tier1PfxSize/10_000-8          75.97 ± 0%   200.20 ± 0%  +163.53% (p=0.002 n=6)
Tier1PfxSize/100_000-8         41.16 ± 0%   164.00 ± 0%  +298.45% (p=0.002 n=6)
Tier1PfxSize/200_000-8         32.85 ± 0%   164.00 ± 0%  +399.24% (p=0.002 n=6)
Tier1PfxSize/500_000-8         25.19 ± 0%   144.50 ± 0%  +473.64% (p=0.002 n=6)
Tier1PfxSize/1_000_000-8       21.39 ± 0%   144.30 ± 0%  +574.61% (p=0.002 n=6)
RandomPfx4Size/1_000-8         74.74 ± 3%   140.80 ± 1%   +88.39% (p=0.002 n=6)
RandomPfx4Size/10_000-8        48.22 ± 0%   109.50 ± 0%  +127.08% (p=0.002 n=6)
RandomPfx4Size/100_000-8       59.57 ± 0%   139.80 ± 0%  +134.68% (p=0.002 n=6)
RandomPfx4Size/200_000-8       52.50 ± 0%   139.80 ± 0%  +166.29% (p=0.002 n=6)
RandomPfx4Size/500_000-8       37.84 ± 0%   139.70 ± 0%  +269.19% (p=0.002 n=6)
RandomPfx4Size/1_000_000-8     31.83 ± 0%   139.60 ± 0%  +338.58% (p=0.002 n=6)
RandomPfx6Size/1_000-8         82.77 ± 2%   157.20 ± 1%   +89.92% (p=0.002 n=6)
RandomPfx6Size/10_000-8        98.78 ± 0%   201.30 ± 0%  +103.79% (p=0.002 n=6)
RandomPfx6Size/100_000-8       67.59 ± 0%   160.80 ± 0%  +137.91% (p=0.002 n=6)
RandomPfx6Size/200_000-8       64.96 ± 0%   160.80 ± 0%  +147.54% (p=0.002 n=6)
RandomPfx6Size/500_000-8       67.61 ± 0%   156.50 ± 0%  +131.47% (p=0.002 n=6)
RandomPfx6Size/1_000_000-8     73.74 ± 0%   156.50 ± 0%  +112.23% (p=0.002 n=6)
RandomPfxSize/1_000-8          99.07 ± 2%   144.60 ± 1%   +45.96% (p=0.002 n=6)
RandomPfxSize/10_000-8         71.29 ± 0%   140.10 ± 0%   +96.52% (p=0.002 n=6)
RandomPfxSize/100_000-8        76.53 ± 0%   180.00 ± 0%  +135.20% (p=0.002 n=6)
RandomPfxSize/200_000-8        65.49 ± 0%   180.00 ± 0%  +174.85% (p=0.002 n=6)
RandomPfxSize/500_000-8        47.94 ± 0%   144.00 ± 0%  +200.38% (p=0.002 n=6)
RandomPfxSize/1_000_000-8      39.50 ± 0%   144.00 ± 0%  +264.56% (p=0.002 n=6)
geomean                        56.08         152.8       +172.42%
```

## update, insert/delete

`bart` is by far the fastest algorithm for updates.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/update.bm │             art/update.bm              │
                           │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000-8        152.6n ± 2%   1433.5n ±  3%   +839.69% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       111.5n ± 1%   1276.0n ±  3%  +1044.91% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      164.0n ± 2%   1490.0n ± 43%   +808.26% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      184.4n ± 4%   1705.5n ± 25%   +824.89% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        113.1n ± 3%    678.6n ±  2%   +500.00% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       79.53n ± 0%   744.95n ±  5%   +836.75% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      164.8n ± 3%    736.9n ± 14%   +347.15% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      209.3n ± 2%    595.7n ±  7%   +184.55% (p=0.002 n=6)
geomean                         141.4n         1.004µ         +610.08%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/update.bm │           netipds/update.bm           │
                           │   sec/route    │   sec/route    vs base                │
InsertRandomPfxs/1_000-8        152.6n ± 2%    226.5n ±  5%   +48.48% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       111.5n ± 1%    286.2n ±  3%  +156.80% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      164.0n ± 2%    392.7n ± 42%  +139.38% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      184.4n ± 4%    472.3n ±  4%  +156.13% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        113.1n ± 3%    149.8n ±  1%   +32.49% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       79.53n ± 0%   202.65n ±  6%  +154.83% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      164.8n ± 3%    326.8n ± 12%   +98.33% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      209.3n ± 2%    407.6n ±  1%   +94.72% (p=0.002 n=6)
geomean                         141.4n         289.0n        +104.41%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/update.bm │          critbitgo/update.bm          │
                           │   sec/route    │   sec/route    vs base                │
InsertRandomPfxs/1_000-8        152.6n ± 2%    293.1n ±  1%   +92.17% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       111.5n ± 1%    374.2n ±  2%  +235.80% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      164.0n ± 2%    550.0n ±  5%  +235.29% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      184.4n ± 4%    703.7n ±  3%  +281.62% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        113.1n ± 3%    150.1n ±  2%   +32.71% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       79.53n ± 0%   203.80n ±  2%  +156.27% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      164.8n ± 3%    391.8n ± 22%  +137.74% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      209.3n ± 2%    499.0n ±  2%  +138.36% (p=0.002 n=6)
geomean                         141.4n         355.3n        +151.32%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/update.bm │          lpmtrie/update.bm           │
                           │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000-8        152.6n ± 2%    439.0n ± 2%  +187.77% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       111.5n ± 1%    513.6n ± 1%  +360.88% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      164.0n ± 2%    753.1n ± 2%  +359.10% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      184.4n ± 4%    894.1n ± 2%  +384.90% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        113.1n ± 3%    152.5n ± 1%   +34.79% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       79.53n ± 0%   235.30n ± 1%  +195.88% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      164.8n ± 3%    512.0n ± 1%  +210.71% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      209.3n ± 2%    676.1n ± 1%  +222.93% (p=0.002 n=6)
geomean                         141.4n         456.5n       +222.89%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/update.bm │          cidranger/update.bm           │
                           │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000-8        152.6n ± 2%    4520.0n ± 1%  +2862.96% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       111.5n ± 1%    5717.5n ± 4%  +5030.10% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      164.0n ± 2%    7814.0n ± 8%  +4663.18% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      184.4n ± 4%    8554.5n ± 5%  +4539.10% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        113.1n ± 3%     694.3n ± 1%   +513.88% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       79.53n ± 0%   1186.50n ± 1%  +1391.98% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      164.8n ± 3%    2486.5n ± 7%  +1408.80% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      209.3n ± 2%    2978.0n ± 6%  +1322.50% (p=0.002 n=6)
geomean                         141.4n          3.183µ       +2151.46%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/update.bm │           cidrtree/update.bm            │
                           │   sec/route    │   sec/route     vs base                 │
InsertRandomPfxs/1_000-8        152.6n ± 2%    1029.5n ±  3%   +574.86% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       111.5n ± 1%    1629.5n ±  6%  +1362.09% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      164.0n ± 2%    2595.0n ± 10%  +1481.83% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      184.4n ± 4%    3046.5n ±  5%  +1552.11% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        113.1n ± 3%    1437.5n ± 11%  +1171.00% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       79.53n ± 0%   2474.50n ±  8%  +3011.60% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      164.8n ± 3%    3663.0n ±  3%  +2122.69% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      209.3n ± 2%    4302.5n ±  5%  +1955.17% (p=0.002 n=6)
geomean                         141.4n          2.285µ        +1516.33%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/update.bm │         go-iptrie/update.bm          │
                           │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000-8        152.6n ± 2%    254.9n ± 2%   +67.09% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       111.5n ± 1%    322.0n ± 7%  +188.92% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      164.0n ± 2%    514.2n ± 5%  +213.44% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      184.4n ± 4%    639.9n ± 5%  +246.99% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        113.1n ± 3%    149.2n ± 2%   +31.92% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       79.53n ± 0%   224.70n ± 7%  +182.55% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      164.8n ± 3%    508.8n ± 4%  +208.74% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      209.3n ± 2%    624.2n ± 6%  +198.18% (p=0.002 n=6)
geomean                         141.4n         360.9n       +155.25%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                           │ bart/update.bm │       kentik-patricia/update.bm       │
                           │   sec/route    │   sec/route    vs base                │
InsertRandomPfxs/1_000-8        152.6n ± 2%    245.6n ±  3%   +61.00% (p=0.002 n=6)
InsertRandomPfxs/10_000-8       111.5n ± 1%    355.9n ±  5%  +219.29% (p=0.002 n=6)
InsertRandomPfxs/100_000-8      164.0n ± 2%    549.6n ± 10%  +235.02% (p=0.002 n=6)
InsertRandomPfxs/200_000-8      184.4n ± 4%    712.5n ± 19%  +286.36% (p=0.002 n=6)
DeleteRandomPfxs/1_000-8        113.1n ± 3%    317.8n ± 10%  +180.95% (p=0.002 n=6)
DeleteRandomPfxs/10_000-8       79.53n ± 0%   402.00n ± 16%  +405.50% (p=0.002 n=6)
DeleteRandomPfxs/100_000-8      164.8n ± 3%    711.3n ±  5%  +331.61% (p=0.002 n=6)
DeleteRandomPfxs/200_000-8      209.3n ± 2%    871.8n ±  1%  +316.43% (p=0.002 n=6)
geomean                         141.4n         477.7n        +237.88%
```
