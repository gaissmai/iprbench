# iprbench
comparing benchmarks for some golang IP routing table implementations:

```
	github.com/tailscale/art
	github.com/gaissmai/bart
	github.com/k-sone/critbitgo
	github.com/Asphaltt/lpmtrie
	github.com/gaissmai/cidrtree
	github.com/yl2chen/cidranger
```

The ~1_000_000 **Tier1** prefix test records (IPv4 and IPv6 routes) are from a full routing table with typical
ISP prefix distribution.

In comparison, the prefix lengths for the random test sets are equally distributed between 1-32 for IPv4 and 1-128 bits for IPv6,
the randomly generated _default-routes_ with prefix length 0 have been sorted out, they distorts the lookup times and there is no
lookup miss at all.

The **RandomPrefixes** without IP versions labeling are composed of a distribution of 4 parts IPv4 to 1 part IPv6 prefixes,
which is approximately the current ratio in the Internet backbone routers.

## make your own benchmarks

```
  $ make dep  
  $ make -B all
```

## size of the routing tables


The memory consumption of the multibit trie `art` with more than
**100_000 randomly distributed IPv6** prefixes brings the OOM killer in action.

`bart` has a factor of 10 lower memory consumption compared to `art`, but is by
a factor of 2 slower in lookup times.

`cidrtree` is the most economical in terms of memory consumption,
but this is also not a trie but a binary search tree and
slower by a magnitude than the other algorithms.


```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │                art/size.bm                │
                         │    bytes     │     bytes       vs base                   │
Tier1PfxSize/100           42.82Ki ± 4%    802.97Ki ± 0%  +1775.21% (p=0.002 n=6)
Tier1PfxSize/1_000         288.9Ki ± 1%    7420.4Ki ± 0%  +2468.17% (p=0.002 n=6)
Tier1PfxSize/10_000        1.822Mi ± 0%    47.172Mi ± 0%  +2488.91% (p=0.002 n=6)
Tier1PfxSize/100_000       7.545Mi ± 0%   160.300Mi ± 0%  +2024.67% (p=0.002 n=6)
Tier1PfxSize/1_000_000     34.33Mi ± 0%    378.23Mi ± 0%  +1001.89% (p=0.002 n=6)
RandomPfx4Size/100         43.33Ki ± 4%    706.59Ki ± 0%  +1530.80% (p=0.002 n=6)
RandomPfx4Size/1_000       301.1Ki ± 1%    7286.5Ki ± 0%  +2319.96% (p=0.002 n=6)
RandomPfx4Size/10_000      2.422Mi ± 0%    57.886Mi ± 0%  +2290.43% (p=0.002 n=6)
RandomPfx4Size/100_000     18.27Mi ± 0%    523.23Mi ± 0%  +2764.14% (p=0.002 n=6)
RandomPfx4Size/1_000_000   109.8Mi ± 0%
RandomPfx6Size/100         150.2Ki ± 1%     700.2Ki ± 0%   +366.29% (p=0.002 n=6)
RandomPfx6Size/1_000       1.319Mi ± 0%     7.477Mi ± 0%   +466.92% (p=0.002 n=6)
RandomPfx6Size/10_000      12.29Mi ± 0%     65.51Mi ± 0%   +433.02% (p=0.002 n=6)
RandomPfx6Size/100_000     115.3Mi ± 0%     749.0Mi ± 0%   +549.46% (p=0.002 n=6)
RandomPfx6Size/1_000_000   1.049Gi ± 0%
RandomPfxSize/100          68.22Ki ± 2%    694.25Ki ± 0%   +917.68% (p=0.002 n=6)
RandomPfxSize/1_000        521.2Ki ± 0%    7439.5Ki ± 0%  +1327.38% (p=0.002 n=6)
RandomPfxSize/10_000       4.400Mi ± 0%    59.641Mi ± 0%  +1255.38% (p=0.002 n=6)
RandomPfxSize/100_000      39.02Mi ± 0%    553.83Mi ± 0%  +1319.37% (p=0.002 n=6)
RandomPfxSize/1_000_000    309.1Mi ± 0%
geomean                    3.652Mi          22.52Mi       +1266.37%               ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │  bart/size.bm  │           cidrtree/size.bm            │
                         │     bytes      │     bytes      vs base                │
Tier1PfxSize/100             42.82Ki ± 4%   12.48Ki ± 14%   -70.86% (p=0.002 n=6)
Tier1PfxSize/1_000          288.94Ki ± 1%   81.51Ki ±  2%   -71.79% (p=0.002 n=6)
Tier1PfxSize/10_000         1865.8Ki ± 0%   784.6Ki ±  0%   -57.95% (p=0.002 n=6)
Tier1PfxSize/100_000         7.545Mi ± 0%   7.633Mi ±  0%    +1.17% (p=0.002 n=6)
Tier1PfxSize/1_000_000       34.33Mi ± 0%   76.30Mi ±  0%  +122.27% (p=0.002 n=6)
RandomPfx4Size/100           43.33Ki ± 4%   11.58Ki ± 14%   -73.28% (p=0.002 n=6)
RandomPfx4Size/1_000        301.10Ki ± 1%   81.51Ki ±  2%   -72.93% (p=0.002 n=6)
RandomPfx4Size/10_000       2479.7Ki ± 0%   784.6Ki ±  0%   -68.36% (p=0.002 n=6)
RandomPfx4Size/100_000      18.268Mi ± 0%   7.633Mi ±  0%   -58.22% (p=0.002 n=6)
RandomPfx4Size/1_000_000    109.79Mi ± 0%   76.30Mi ±  0%   -30.51% (p=0.002 n=6)
RandomPfx6Size/100          150.16Ki ± 1%   11.58Ki ± 14%   -92.29% (p=0.002 n=6)
RandomPfx6Size/1_000       1350.49Ki ± 0%   81.51Ki ±  2%   -93.96% (p=0.002 n=6)
RandomPfx6Size/10_000      12585.8Ki ± 0%   784.6Ki ±  0%   -93.77% (p=0.002 n=6)
RandomPfx6Size/100_000     115.320Mi ± 0%   7.633Mi ±  0%   -93.38% (p=0.002 n=6)
RandomPfx6Size/1_000_000   1074.54Mi ± 0%   76.30Mi ±  0%   -92.90% (p=0.002 n=6)
RandomPfxSize/100            68.22Ki ± 2%   11.58Ki ± 14%   -83.03% (p=0.002 n=6)
RandomPfxSize/1_000         521.20Ki ± 0%   81.51Ki ±  2%   -84.36% (p=0.002 n=6)
RandomPfxSize/10_000        4505.9Ki ± 0%   784.6Ki ±  0%   -82.59% (p=0.002 n=6)
RandomPfxSize/100_000       39.019Mi ± 0%   7.633Mi ±  0%   -80.44% (p=0.002 n=6)
RandomPfxSize/1_000_000     309.15Mi ± 0%   76.30Mi ±  0%   -75.32% (p=0.002 n=6)
geomean                      3.652Mi        856.4Ki         -77.10%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm  │           critbitgo/size.bm            │
                         │     bytes     │     bytes       vs base                │
Tier1PfxSize/100            42.82Ki ± 4%    15.63Ki ± 11%   -63.49% (p=0.002 n=6)
Tier1PfxSize/1_000          288.9Ki ± 1%    115.2Ki ±  1%   -60.12% (p=0.002 n=6)
Tier1PfxSize/10_000         1.822Mi ± 0%    1.094Mi ±  0%   -39.98% (p=0.002 n=6)
Tier1PfxSize/100_000        7.545Mi ± 0%   10.913Mi ±  0%   +44.64% (p=0.002 n=6)
Tier1PfxSize/1_000_000      34.33Mi ± 0%   109.12Mi ±  0%  +217.88% (p=0.002 n=6)
RandomPfx4Size/100          43.33Ki ± 4%    14.67Ki ± 11%   -66.14% (p=0.002 n=6)
RandomPfx4Size/1_000        301.1Ki ± 1%    112.7Ki ±  1%   -62.56% (p=0.002 n=6)
RandomPfx4Size/10_000       2.422Mi ± 0%    1.071Mi ±  0%   -55.76% (p=0.002 n=6)
RandomPfx4Size/100_000      18.27Mi ± 0%    10.68Mi ±  0%   -41.51% (p=0.002 n=6)
RandomPfx4Size/1_000_000    109.8Mi ± 0%    106.8Mi ±  0%    -2.71% (p=0.002 n=6)
RandomPfx6Size/100         150.16Ki ± 1%    16.22Ki ± 10%   -89.20% (p=0.002 n=6)
RandomPfx6Size/1_000       1350.5Ki ± 0%    128.3Ki ±  1%   -90.50% (p=0.002 n=6)
RandomPfx6Size/10_000      12.291Mi ± 0%    1.224Mi ±  0%   -90.04% (p=0.002 n=6)
RandomPfx6Size/100_000     115.32Mi ± 0%    12.21Mi ±  0%   -89.41% (p=0.002 n=6)
RandomPfx6Size/1_000_000   1074.5Mi ± 0%    122.1Mi ±  0%   -88.64% (p=0.002 n=6)
RandomPfxSize/100           68.22Ki ± 2%    14.94Ki ± 11%   -78.10% (p=0.002 n=6)
RandomPfxSize/1_000         521.2Ki ± 0%    115.8Ki ±  1%   -77.79% (p=0.002 n=6)
RandomPfxSize/10_000        4.400Mi ± 0%    1.102Mi ±  0%   -74.96% (p=0.002 n=6)
RandomPfxSize/100_000       39.02Mi ± 0%    10.99Mi ±  0%   -71.83% (p=0.002 n=6)
RandomPfxSize/1_000_000     309.1Mi ± 0%    109.9Mi ±  0%   -64.46% (p=0.002 n=6)
geomean                     3.652Mi         1.193Mi         -67.33%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm  │            lpmtrie/size.bm             │
                         │     bytes     │     bytes       vs base                │
Tier1PfxSize/100            42.82Ki ± 4%    24.25Ki ±  9%   -43.37% (p=0.002 n=6)
Tier1PfxSize/1_000          288.9Ki ± 1%    202.3Ki ±  4%   -29.97% (p=0.002 n=6)
Tier1PfxSize/10_000         1.822Mi ± 0%    1.942Mi ±  3%    +6.58% (p=0.002 n=6)
Tier1PfxSize/100_000        7.545Mi ± 0%   19.330Mi ±  3%  +156.21% (p=0.002 n=6)
Tier1PfxSize/1_000_000      34.33Mi ± 0%   189.92Mi ±  6%  +453.28% (p=0.002 n=6)
RandomPfx4Size/100          43.33Ki ± 4%    23.20Ki ±  9%   -46.45% (p=0.002 n=6)
RandomPfx4Size/1_000        301.1Ki ± 1%    198.5Ki ±  3%   -34.07% (p=0.002 n=6)
RandomPfx4Size/10_000       2.422Mi ± 0%    1.898Mi ±  3%   -21.64% (p=0.002 n=6)
RandomPfx4Size/100_000      18.27Mi ± 0%    18.28Mi ±  5%         ~ (p=0.374 n=6)
RandomPfx4Size/1_000_000    109.8Mi ± 0%    163.5Mi ±  8%   +48.93% (p=0.002 n=6)
RandomPfx6Size/100         150.16Ki ± 1%    25.53Ki ± 12%   -83.00% (p=0.002 n=6)
RandomPfx6Size/1_000       1350.5Ki ± 0%    222.0Ki ±  8%   -83.56% (p=0.002 n=6)
RandomPfx6Size/10_000      12.291Mi ± 0%    2.129Mi ±  7%   -82.68% (p=0.002 n=6)
RandomPfx6Size/100_000     115.32Mi ± 0%    20.93Mi ±  8%   -81.85% (p=0.002 n=6)
RandomPfx6Size/1_000_000   1074.5Mi ± 0%    204.0Mi ±  8%   -81.02% (p=0.002 n=6)
RandomPfxSize/100           68.22Ki ± 2%    23.56Ki ± 10%   -65.46% (p=0.002 n=6)
RandomPfxSize/1_000         521.2Ki ± 0%    202.8Ki ±  4%   -61.10% (p=0.002 n=6)
RandomPfxSize/10_000        4.400Mi ± 0%    1.943Mi ±  4%   -55.85% (p=0.002 n=6)
RandomPfxSize/100_000       39.02Mi ± 0%    18.78Mi ±  5%   -51.86% (p=0.002 n=6)
RandomPfxSize/1_000_000     309.1Mi ± 0%    171.9Mi ±  7%   -44.39% (p=0.002 n=6)
geomean                     3.652Mi         2.011Mi         -44.95%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm  │           cidranger/size.bm            │
                         │     bytes     │     bytes      vs base                 │
Tier1PfxSize/100            42.82Ki ± 4%    56.03Ki ± 5%    +30.85% (p=0.002 n=6)
Tier1PfxSize/1_000          288.9Ki ± 1%    525.9Ki ± 3%    +82.02% (p=0.002 n=6)
Tier1PfxSize/10_000         1.822Mi ± 0%    5.091Mi ± 2%   +179.41% (p=0.002 n=6)
Tier1PfxSize/100_000        7.545Mi ± 0%   50.258Mi ± 2%   +566.14% (p=0.002 n=6)
Tier1PfxSize/1_000_000      34.33Mi ± 0%   477.45Mi ± 2%  +1290.93% (p=0.002 n=6)
RandomPfx4Size/100          43.33Ki ± 4%    54.16Ki ± 6%    +24.99% (p=0.002 n=6)
RandomPfx4Size/1_000        301.1Ki ± 1%    514.0Ki ± 3%    +70.69% (p=0.002 n=6)
RandomPfx4Size/10_000       2.422Mi ± 0%    4.900Mi ± 3%   +102.35% (p=0.002 n=6)
RandomPfx4Size/100_000      18.27Mi ± 0%    45.76Mi ± 3%   +150.47% (p=0.002 n=6)
RandomPfx4Size/1_000_000    109.8Mi ± 0%    396.2Mi ± 3%   +260.87% (p=0.002 n=6)
RandomPfx6Size/100         150.16Ki ± 1%    61.08Ki ± 3%    -59.33% (p=0.002 n=6)
RandomPfx6Size/1_000       1350.5Ki ± 0%    579.3Ki ± 0%    -57.11% (p=0.002 n=6)
RandomPfx6Size/10_000      12.291Mi ± 0%    5.582Mi ± 0%    -54.59% (p=0.002 n=6)
RandomPfx6Size/100_000     115.32Mi ± 0%    54.75Mi ± 0%    -52.52% (p=0.002 n=6)
RandomPfx6Size/1_000_000   1074.5Mi ± 0%    534.1Mi ± 0%    -50.29% (p=0.002 n=6)
RandomPfxSize/100           68.22Ki ± 2%    55.07Ki ± 5%    -19.27% (p=0.002 n=6)
RandomPfxSize/1_000         521.2Ki ± 0%    524.9Ki ± 2%     +0.70% (p=0.026 n=6)
RandomPfxSize/10_000        4.400Mi ± 0%    5.024Mi ± 2%    +14.18% (p=0.002 n=6)
RandomPfxSize/100_000       39.02Mi ± 0%    47.53Mi ± 2%    +21.82% (p=0.002 n=6)
RandomPfxSize/1_000_000     309.1Mi ± 0%    424.9Mi ± 2%    +37.44% (p=0.002 n=6)
geomean                     3.652Mi         5.068Mi         +38.75%
```

## lookup (longest-prefix-match)

In the lookup, `art` is the champion, closely followed by `bart`. 

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │            art/lookup.bm             │
                                    │     sec/op     │    sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4             59.84n ± 31%   48.52n ± 16%        ~ (p=0.075 n=10)
LpmTier1Pfxs/RandomMatchIP6             72.45n ± 38%   47.73n ±  4%  -34.12% (p=0.022 n=10)
LpmTier1Pfxs/RandomMissIP4              97.48n ± 23%   31.57n ±  0%  -67.61% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              21.00n ±  4%   31.20n ±  0%  +48.54% (p=0.001 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     77.81n ± 28%   56.62n ± 15%  -27.24% (p=0.019 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     63.97n ± 26%   48.90n ± 17%        ~ (p=0.052 n=10)
LpmRandomPfxs100_000/RandomMissIP4      92.05n ± 20%   31.51n ±  1%  -65.77% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      77.34n ± 37%   31.16n ±  1%  -59.71% (p=0.000 n=10)
geomean                                 64.89n         39.72n        -38.79%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │             cidrtree/lookup.bm             │
                                    │     sec/op     │      sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4             59.84n ± 31%   1144.50n ±   35%  +1812.60% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             72.45n ± 38%   1443.50n ±   15%  +1892.55% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              97.48n ± 23%   1183.50n ±   44%  +1114.10% (p=0.002 n=10)
LpmTier1Pfxs/RandomMissIP6              21.00n ±  4%     70.51n ± 2012%   +235.71% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     77.81n ± 28%   1236.50n ±   34%  +1489.03% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     63.97n ± 26%   1432.00n ±   32%  +2138.72% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      92.05n ± 20%   1286.50n ±   23%  +1297.61% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      77.34n ± 37%   1431.50n ±   14%  +1750.92% (p=0.000 n=10)
geomean                                 64.89n           904.9n          +1294.42%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │           critbitgo/lookup.bm           │
                                    │     sec/op     │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             59.84n ± 31%   373.90n ± 16%   +524.83% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             72.45n ± 38%   699.95n ± 56%   +866.18% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              97.48n ± 23%   772.05n ± 49%   +692.01% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              21.00n ±  4%   757.05n ± 21%  +3504.14% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     77.81n ± 28%   337.50n ± 53%   +333.72% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     63.97n ± 26%   697.65n ± 12%   +990.67% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      92.05n ± 20%   560.25n ± 40%   +508.64% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      77.34n ± 37%   631.00n ± 24%   +715.88% (p=0.000 n=10)
geomean                                 64.89n          579.6n         +793.11%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │           lpmtrie/lookup.bm            │
                                    │     sec/op     │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4             59.84n ± 31%   245.50n ± 19%  +310.26% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             72.45n ± 38%   250.45n ± 63%  +245.71% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              97.48n ± 23%   233.20n ± 67%  +139.23% (p=0.029 n=10)
LpmTier1Pfxs/RandomMissIP6              21.00n ±  4%    88.63n ± 21%  +321.92% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     77.81n ± 28%   259.75n ±  9%  +233.80% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     63.97n ± 26%   199.95n ±  7%  +212.59% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      92.05n ± 20%   278.10n ±  5%  +202.12% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      77.34n ± 37%   235.70n ±  4%  +204.76% (p=0.000 n=10)
geomean                                 64.89n          213.6n        +229.08%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │           cidranger/lookup.bm            │
                                    │     sec/op     │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             59.84n ± 31%    351.55n ± 44%   +487.48% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             72.45n ± 38%    292.70n ± 42%   +304.03% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              97.48n ± 23%    206.50n ± 56%   +111.84% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              21.00n ±  4%    160.50n ± 24%   +664.10% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     77.81n ± 28%   1352.50n ± 61%  +1638.10% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     63.97n ± 26%    481.05n ± 43%   +652.05% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      92.05n ± 20%    452.45n ± 24%   +391.53% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      77.34n ± 37%    400.55n ± 21%   +417.91% (p=0.000 n=10)
geomean                                 64.89n           376.3n         +479.88%
```

## update, insert/delete

`bart` is by far the fastest algorithm for updates under all competitors.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │             art/update.bm             │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100            54.36n ± 0%   105.85n ± 1%   +94.72% (p=0.000 n=10)
Insert/Insert_into_1_000          54.35n ± 0%   125.45n ± 1%  +130.82% (p=0.000 n=10)
Insert/Insert_into_10_000         54.36n ± 0%   129.05n ± 1%  +137.40% (p=0.000 n=10)
Insert/Insert_into_100_000        54.37n ± 0%   120.55n ± 2%  +121.74% (p=0.000 n=10)
Insert/Insert_into_1_000_000      64.71n ± 0%   121.00n ± 0%   +86.99% (p=0.000 n=10)
Delete/Delete_from_100            18.53n ± 0%    34.04n ± 3%   +83.70% (p=0.000 n=10)
Delete/Delete_from_1_000          43.07n ± 0%    59.06n ± 0%   +37.13% (p=0.000 n=10)
Delete/Delete_from_10_000         43.09n ± 0%    58.93n ± 0%   +36.75% (p=0.000 n=10)
Delete/Delete_from_100_000        45.93n ± 0%    58.71n ± 0%   +27.82% (p=0.000 n=10)
Delete/Delete_from_1_000_000      45.68n ± 0%    58.50n ± 0%   +28.08% (p=0.000 n=10)
geomean                           45.82n         79.57n        +73.66%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │            cidrtree/update.bm            │
                             │     sec/op     │    sec/op      vs base                   │
Insert/Insert_into_100            54.36n ± 0%    945.20n ± 0%   +1638.78% (p=0.000 n=10)
Insert/Insert_into_1_000          54.35n ± 0%   1452.00n ± 0%   +2571.57% (p=0.000 n=10)
Insert/Insert_into_10_000         54.36n ± 0%   2125.00n ± 0%   +3809.12% (p=0.000 n=10)
Insert/Insert_into_100_000        54.37n ± 0%   1595.00n ± 0%   +2833.87% (p=0.000 n=10)
Insert/Insert_into_1_000_000      64.71n ± 0%   2627.50n ± 0%   +3960.42% (p=0.000 n=10)
Delete/Delete_from_100            18.53n ± 0%   1502.00n ± 0%   +8005.77% (p=0.000 n=10)
Delete/Delete_from_1_000          43.07n ± 0%   1765.00n ± 0%   +3997.98% (p=0.000 n=10)
Delete/Delete_from_10_000         43.09n ± 0%   3444.00n ± 0%   +7892.57% (p=0.000 n=10)
Delete/Delete_from_100_000        45.93n ± 0%   3493.50n ± 0%   +7506.14% (p=0.000 n=10)
Delete/Delete_from_1_000_000      45.68n ± 0%   4819.00n ± 0%  +10449.47% (p=0.000 n=10)
geomean                           45.82n          2.125µ        +4538.26%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │          critbitgo/update.bm          │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100            54.36n ± 0%   109.95n ± 1%  +102.26% (p=0.000 n=10)
Insert/Insert_into_1_000          54.35n ± 0%   121.75n ± 1%  +124.01% (p=0.000 n=10)
Insert/Insert_into_10_000         54.36n ± 0%   125.50n ± 1%  +130.87% (p=0.000 n=10)
Insert/Insert_into_100_000        54.37n ± 0%   147.50n ± 1%  +171.31% (p=0.000 n=10)
Insert/Insert_into_1_000_000      64.71n ± 0%   152.20n ± 2%  +135.20% (p=0.000 n=10)
Delete/Delete_from_100            18.53n ± 0%    99.93n ± 1%  +439.29% (p=0.000 n=10)
Delete/Delete_from_1_000          43.07n ± 0%   108.40n ± 1%  +151.68% (p=0.000 n=10)
Delete/Delete_from_10_000         43.09n ± 0%   112.55n ± 1%  +161.20% (p=0.000 n=10)
Delete/Delete_from_100_000        45.93n ± 0%   131.40n ± 2%  +186.09% (p=0.000 n=10)
Delete/Delete_from_1_000_000      45.68n ± 0%   127.40n ± 2%  +178.90% (p=0.000 n=10)
geomean                           45.82n         122.6n       +167.67%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │           lpmtrie/update.bm           │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100            54.36n ± 0%   314.55n ± 0%  +478.64% (p=0.000 n=10)
Insert/Insert_into_1_000          54.35n ± 0%   352.45n ± 2%  +548.48% (p=0.000 n=10)
Insert/Insert_into_10_000         54.36n ± 0%   362.60n ± 1%  +567.03% (p=0.000 n=10)
Insert/Insert_into_100_000        54.37n ± 0%   507.25n ± 1%  +833.05% (p=0.000 n=10)
Insert/Insert_into_1_000_000      64.71n ± 0%   656.80n ± 1%  +914.99% (p=0.000 n=10)
Delete/Delete_from_100            18.53n ± 0%    77.03n ± 3%  +315.73% (p=0.000 n=10)
Delete/Delete_from_1_000          43.07n ± 0%   123.30n ± 0%  +186.28% (p=0.000 n=10)
Delete/Delete_from_10_000         43.09n ± 0%   149.50n ± 1%  +246.95% (p=0.000 n=10)
Delete/Delete_from_100_000        45.93n ± 0%   266.55n ± 0%  +480.34% (p=0.000 n=10)
Delete/Delete_from_1_000_000      45.68n ± 0%   309.70n ± 3%  +577.98% (p=0.000 n=10)
geomean                           45.82n         262.8n       +473.53%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │           cidranger/update.bm            │
                             │     sec/op     │    sec/op      vs base                   │
Insert/Insert_into_100            54.36n ± 0%   2442.50n ± 1%   +4393.19% (p=0.000 n=10)
Insert/Insert_into_1_000          54.35n ± 0%   4130.00n ± 2%   +7498.90% (p=0.000 n=10)
Insert/Insert_into_10_000         54.36n ± 0%   4849.00n ± 1%   +8820.16% (p=0.000 n=10)
Insert/Insert_into_100_000        54.37n ± 0%   8110.50n ± 1%  +14818.61% (p=0.000 n=10)
Insert/Insert_into_1_000_000      64.71n ± 0%   7579.00n ± 1%  +11612.25% (p=0.000 n=10)
Delete/Delete_from_100            18.53n ± 0%    385.05n ± 3%   +1977.98% (p=0.000 n=10)
Delete/Delete_from_1_000          43.07n ± 0%    421.45n ± 0%    +878.52% (p=0.000 n=10)
Delete/Delete_from_10_000         43.09n ± 0%    442.30n ± 1%    +926.46% (p=0.000 n=10)
Delete/Delete_from_100_000        45.93n ± 0%    534.85n ± 8%   +1064.49% (p=0.000 n=10)
Delete/Delete_from_1_000_000      45.68n ± 0%    549.85n ± 0%   +1103.70% (p=0.000 n=10)
geomean                           45.82n          1.514µ        +3205.41%
```
