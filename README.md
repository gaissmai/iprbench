# iprbench
comparing benchmarks for some golang IP routing table implementations:

```
	github.com/tailscale/art
	github.com/gaissmai/bart
	github.com/k-sone/critbitgo
	github.com/Asphaltt/lpmtrie
	github.com/gaissmai/cidrtree
```
The ~1_000_000 **Tier1** prefix test set (IPv4 and IPv6 routes) are from a full routing table with typical ISP prefix distribution.
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


For the multibit tries `art` and `bart` the memory consumption explodes with more
than **100_000 randomly distributed IPv6** prefixes in contrast to the other algorithms,
but these two algorithms are much faster than the others.

`bart` has about a factor of 10 lower memory consumption compared to `art`.

`cidrtree` is the most economical in terms of memory consumption,
but this is also not a trie but a binary search tree and
slower by a magnitude than the other algorithms.


```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │                art/size.bm                 │
                         │    bytes     │     bytes       vs base                    │
Tier1PfxSize/100           42.81Ki ± 0%    802.97Ki ± 0%  +1775.55% (p=0.000 n=10)
Tier1PfxSize/1_000         288.9Ki ± 0%    7420.4Ki ± 0%  +2468.17% (p=0.000 n=10)
Tier1PfxSize/10_000        1.822Mi ± 0%    47.172Mi ± 0%  +2488.91% (p=0.000 n=10)
Tier1PfxSize/100_000       7.545Mi ± 0%   160.300Mi ± 0%  +2024.67% (p=0.000 n=10)
Tier1PfxSize/1_000_000     34.33Mi ± 0%    378.23Mi ± 0%  +1001.89% (p=0.000 n=10)
RandomPfx4Size/100         45.30Ki ± 0%    732.09Ki ± 0%  +1516.21% (p=0.000 n=10)
RandomPfx4Size/1_000       305.6Ki ± 0%    7394.9Ki ± 0%  +2319.60% (p=0.000 n=10)
RandomPfx4Size/10_000      2.422Mi ± 0%    57.948Mi ± 0%  +2292.50% (p=0.000 n=10)
RandomPfx4Size/100_000     18.30Mi ± 0%    523.45Mi ± 0%  +2760.41% (p=0.000 n=10)
RandomPfx4Size/1_000_000   109.8Mi ± 0%
RandomPfx6Size/100         156.8Ki ± 0%     738.5Ki ± 0%   +370.99% (p=0.000 n=10)
RandomPfx6Size/1_000       1.316Mi ± 0%     7.558Mi ± 0%   +474.33% (p=0.000 n=10)
RandomPfx6Size/10_000      12.08Mi ± 0%     65.33Mi ± 0%   +440.70% (p=0.000 n=10)
RandomPfx6Size/100_000     115.1Mi ± 0%     748.7Mi ± 0%   +550.58% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1.049Gi ± 0%
RandomPfxSize/100          64.68Ki ± 0%    675.12Ki ± 0%   +943.80% (p=0.000 n=10)
RandomPfxSize/1_000        513.7Ki ± 0%    7439.5Ki ± 0%  +1348.32% (p=0.000 n=10)
RandomPfxSize/10_000       4.367Mi ± 0%    59.579Mi ± 0%  +1264.21% (p=0.000 n=10)
RandomPfxSize/100_000      38.98Mi ± 0%    553.62Mi ± 0%  +1320.28% (p=0.000 n=10)
RandomPfxSize/1_000_000    309.1Mi ± 0%
geomean                    3.653Mi          22.63Mi       +1272.55%                ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │  bart/size.bm  │           cidrtree/size.bm            │
                         │     bytes      │    bytes      vs base                 │
Tier1PfxSize/100             42.81Ki ± 0%   12.48Ki ± 0%   -70.86% (p=0.000 n=10)
Tier1PfxSize/1_000          288.94Ki ± 0%   81.51Ki ± 0%   -71.79% (p=0.000 n=10)
Tier1PfxSize/10_000         1865.8Ki ± 0%   784.6Ki ± 0%   -57.95% (p=0.000 n=10)
Tier1PfxSize/100_000         7.545Mi ± 0%   7.633Mi ± 0%    +1.17% (p=0.000 n=10)
Tier1PfxSize/1_000_000       34.33Mi ± 0%   76.30Mi ± 0%  +122.27% (p=0.000 n=10)
RandomPfx4Size/100           45.30Ki ± 0%   11.58Ki ± 0%   -74.44% (p=0.000 n=10)
RandomPfx4Size/1_000        305.62Ki ± 0%   81.51Ki ± 0%   -73.33% (p=0.000 n=10)
RandomPfx4Size/10_000       2480.2Ki ± 0%   784.6Ki ± 0%   -68.36% (p=0.000 n=10)
RandomPfx4Size/100_000      18.300Mi ± 0%   7.633Mi ± 0%   -58.29% (p=0.000 n=10)
RandomPfx4Size/1_000_000    109.80Mi ± 0%   76.30Mi ± 0%   -30.51% (p=0.000 n=10)
RandomPfx6Size/100          156.79Ki ± 0%   11.58Ki ± 0%   -92.62% (p=0.000 n=10)
RandomPfx6Size/1_000       1347.52Ki ± 0%   81.51Ki ± 0%   -93.95% (p=0.000 n=10)
RandomPfx6Size/10_000      12372.7Ki ± 0%   784.6Ki ± 0%   -93.66% (p=0.000 n=10)
RandomPfx6Size/100_000     115.076Mi ± 0%   7.633Mi ± 0%   -93.37% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1074.54Mi ± 0%   76.30Mi ± 0%   -92.90% (p=0.000 n=10)
RandomPfxSize/100            64.68Ki ± 0%   11.58Ki ± 0%   -82.10% (p=0.000 n=10)
RandomPfxSize/1_000         513.66Ki ± 0%   81.51Ki ± 0%   -84.13% (p=0.000 n=10)
RandomPfxSize/10_000        4472.1Ki ± 0%   784.6Ki ± 0%   -82.46% (p=0.000 n=10)
RandomPfxSize/100_000       38.980Mi ± 0%   7.633Mi ± 0%   -80.42% (p=0.000 n=10)
RandomPfxSize/1_000_000     309.15Mi ± 0%   76.30Mi ± 0%   -75.32% (p=0.000 n=10)
geomean                      3.653Mi        856.4Ki        -77.11%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm  │           critbitgo/size.bm            │
                         │     bytes     │     bytes      vs base                 │
Tier1PfxSize/100            42.81Ki ± 0%    15.63Ki ± 0%   -63.49% (p=0.000 n=10)
Tier1PfxSize/1_000          288.9Ki ± 0%    115.2Ki ± 0%   -60.12% (p=0.000 n=10)
Tier1PfxSize/10_000         1.822Mi ± 0%    1.094Mi ± 0%   -39.98% (p=0.000 n=10)
Tier1PfxSize/100_000        7.545Mi ± 0%   10.913Mi ± 0%   +44.64% (p=0.000 n=10)
Tier1PfxSize/1_000_000      34.33Mi ± 0%   109.12Mi ± 0%  +217.88% (p=0.000 n=10)
RandomPfx4Size/100          45.30Ki ± 0%    14.67Ki ± 0%   -67.61% (p=0.000 n=10)
RandomPfx4Size/1_000        305.6Ki ± 0%    112.7Ki ± 0%   -63.12% (p=0.000 n=10)
RandomPfx4Size/10_000       2.422Mi ± 0%    1.071Mi ± 0%   -55.77% (p=0.000 n=10)
RandomPfx4Size/100_000      18.30Mi ± 0%    10.68Mi ± 0%   -41.61% (p=0.000 n=10)
RandomPfx4Size/1_000_000    109.8Mi ± 0%    106.8Mi ± 0%    -2.72% (p=0.000 n=10)
RandomPfx6Size/100         156.79Ki ± 0%    16.22Ki ± 0%   -89.66% (p=0.000 n=10)
RandomPfx6Size/1_000       1347.5Ki ± 0%    128.3Ki ± 0%   -90.48% (p=0.000 n=10)
RandomPfx6Size/10_000      12.083Mi ± 0%    1.224Mi ± 0%   -89.87% (p=0.000 n=10)
RandomPfx6Size/100_000     115.08Mi ± 0%    12.21Mi ± 0%   -89.39% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1074.5Mi ± 0%    122.1Mi ± 0%   -88.64% (p=0.000 n=10)
RandomPfxSize/100           64.68Ki ± 0%    14.94Ki ± 0%   -76.91% (p=0.000 n=10)
RandomPfxSize/1_000         513.7Ki ± 0%    115.8Ki ± 0%   -77.46% (p=0.000 n=10)
RandomPfxSize/10_000        4.367Mi ± 0%    1.102Mi ± 0%   -74.77% (p=0.000 n=10)
RandomPfxSize/100_000       38.98Mi ± 0%    10.99Mi ± 0%   -71.81% (p=0.000 n=10)
RandomPfxSize/1_000_000     309.1Mi ± 0%    109.9Mi ± 0%   -64.46% (p=0.000 n=10)
geomean                     3.653Mi         1.193Mi        -67.34%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm  │            lpmtrie/size.bm             │
                         │     bytes     │     bytes      vs base                 │
Tier1PfxSize/100            42.81Ki ± 0%    24.24Ki ± 0%   -43.38% (p=0.000 n=10)
Tier1PfxSize/1_000          288.9Ki ± 0%    202.3Ki ± 0%   -29.97% (p=0.000 n=10)
Tier1PfxSize/10_000         1.822Mi ± 0%    1.942Mi ± 0%    +6.58% (p=0.000 n=10)
Tier1PfxSize/100_000        7.545Mi ± 0%   19.330Mi ± 0%  +156.21% (p=0.000 n=10)
Tier1PfxSize/1_000_000      34.33Mi ± 0%   189.92Mi ± 0%  +453.28% (p=0.000 n=10)
RandomPfx4Size/100          45.30Ki ± 0%    23.20Ki ± 0%   -48.78% (p=0.000 n=10)
RandomPfx4Size/1_000        305.6Ki ± 0%    198.1Ki ± 0%   -35.17% (p=0.000 n=10)
RandomPfx4Size/10_000       2.422Mi ± 0%    1.897Mi ± 0%   -21.70% (p=0.000 n=10)
RandomPfx4Size/100_000      18.30Mi ± 0%    18.25Mi ± 0%    -0.25% (p=0.000 n=10)
RandomPfx4Size/1_000_000    109.8Mi ± 0%    163.4Mi ± 0%   +48.82% (p=0.000 n=10)
RandomPfx6Size/100         156.79Ki ± 0%    25.53Ki ± 0%   -83.72% (p=0.000 n=10)
RandomPfx6Size/1_000       1347.5Ki ± 0%    221.8Ki ± 0%   -83.54% (p=0.000 n=10)
RandomPfx6Size/10_000      12.083Mi ± 0%    2.126Mi ± 0%   -82.40% (p=0.000 n=10)
RandomPfx6Size/100_000     115.08Mi ± 0%    20.94Mi ± 0%   -81.80% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1074.5Mi ± 0%    203.9Mi ± 0%   -81.02% (p=0.000 n=10)
RandomPfxSize/100           64.68Ki ± 0%    23.56Ki ± 0%   -63.57% (p=0.000 n=10)
RandomPfxSize/1_000         513.7Ki ± 0%    202.9Ki ± 0%   -60.49% (p=0.000 n=10)
RandomPfxSize/10_000        4.367Mi ± 0%    1.941Mi ± 0%   -55.55% (p=0.000 n=10)
RandomPfxSize/100_000       38.98Mi ± 0%    18.80Mi ± 0%   -51.76% (p=0.000 n=10)
RandomPfxSize/1_000_000     309.1Mi ± 0%    171.9Mi ± 0%   -44.39% (p=0.000 n=10)
geomean                     3.653Mi         2.010Mi        -44.98%
```

## update, insert/delete

When it comes to updates, `art` and `bart` are far superior
to the other algorithms, only `critbitgo` comes close to playing in the same league .

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │             art/update.bm             │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100            53.64n ± 0%   105.60n ± 1%   +96.87% (p=0.000 n=10)
Insert/Insert_into_1_000          55.30n ± 0%   125.40n ± 1%  +126.74% (p=0.000 n=10)
Insert/Insert_into_10_000         53.67n ± 0%   129.55n ± 2%  +141.36% (p=0.000 n=10)
Insert/Insert_into_100_000        53.63n ± 0%   121.75n ± 2%  +127.00% (p=0.000 n=10)
Insert/Insert_into_1_000_000      55.50n ± 0%   120.70n ± 1%  +117.48% (p=0.000 n=10)
Delete/Delete_from_100            17.32n ± 0%    34.32n ± 4%   +98.10% (p=0.000 n=10)
Delete/Delete_from_1_000          41.12n ± 0%    59.12n ± 0%   +43.77% (p=0.000 n=10)
Delete/Delete_from_10_000         41.15n ± 0%    59.95n ± 0%   +45.70% (p=0.000 n=10)
Delete/Delete_from_100_000        41.12n ± 0%    58.71n ± 0%   +42.78% (p=0.000 n=10)
Delete/Delete_from_1_000_000      43.83n ± 0%    58.99n ± 0%   +34.58% (p=0.000 n=10)
geomean                           43.64n         79.91n        +83.13%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │            cidrtree/update.bm            │
                             │     sec/op     │    sec/op      vs base                   │
Insert/Insert_into_100            53.64n ± 0%    949.75n ± 0%   +1670.60% (p=0.000 n=10)
Insert/Insert_into_1_000          55.30n ± 0%   1673.00n ± 0%   +2925.04% (p=0.000 n=10)
Insert/Insert_into_10_000         53.67n ± 0%   2025.50n ± 0%   +3673.64% (p=0.000 n=10)
Insert/Insert_into_100_000        53.63n ± 0%   2981.00n ± 0%   +5457.94% (p=0.000 n=10)
Insert/Insert_into_1_000_000      55.50n ± 0%   3189.00n ± 0%   +5645.95% (p=0.000 n=10)
Delete/Delete_from_100            17.32n ± 0%   1189.00n ± 0%   +6762.91% (p=0.000 n=10)
Delete/Delete_from_1_000          41.12n ± 0%   2112.50n ± 0%   +5037.40% (p=0.000 n=10)
Delete/Delete_from_10_000         41.15n ± 0%   4306.50n ± 0%  +10366.64% (p=0.000 n=10)
Delete/Delete_from_100_000        41.12n ± 0%   4466.50n ± 0%  +10763.43% (p=0.000 n=10)
Delete/Delete_from_1_000_000      43.83n ± 0%   4220.00n ± 0%   +9528.11% (p=0.000 n=10)
geomean                           43.64n          2.396µ        +5390.93%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │          critbitgo/update.bm          │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100            53.64n ± 0%   109.55n ± 0%  +104.23% (p=0.000 n=10)
Insert/Insert_into_1_000          55.30n ± 0%   122.00n ± 0%  +120.59% (p=0.000 n=10)
Insert/Insert_into_10_000         53.67n ± 0%   125.80n ± 0%  +134.37% (p=0.000 n=10)
Insert/Insert_into_100_000        53.63n ± 0%   154.60n ± 1%  +188.24% (p=0.000 n=10)
Insert/Insert_into_1_000_000      55.50n ± 0%   179.80n ± 0%  +223.96% (p=0.000 n=10)
Delete/Delete_from_100            17.32n ± 0%    99.64n ± 0%  +475.12% (p=0.000 n=10)
Delete/Delete_from_1_000          41.12n ± 0%   110.65n ± 0%  +169.09% (p=0.000 n=10)
Delete/Delete_from_10_000         41.15n ± 0%   112.40n ± 0%  +173.18% (p=0.000 n=10)
Delete/Delete_from_100_000        41.12n ± 0%   129.80n ± 1%  +215.70% (p=0.000 n=10)
Delete/Delete_from_1_000_000      43.83n ± 0%   126.65n ± 1%  +188.96% (p=0.000 n=10)
geomean                           43.64n         125.3n       +187.08%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │           lpmtrie/update.bm            │
                             │     sec/op     │    sec/op     vs base                  │
Insert/Insert_into_100            53.64n ± 0%   318.45n ± 1%   +493.68% (p=0.000 n=10)
Insert/Insert_into_1_000          55.30n ± 0%   348.75n ± 1%   +530.59% (p=0.000 n=10)
Insert/Insert_into_10_000         53.67n ± 0%   361.75n ± 1%   +573.96% (p=0.000 n=10)
Insert/Insert_into_100_000        53.63n ± 0%   502.35n ± 1%   +836.61% (p=0.000 n=10)
Insert/Insert_into_1_000_000      55.50n ± 0%   682.75n ± 1%  +1130.18% (p=0.000 n=10)
Delete/Delete_from_100            17.32n ± 0%    76.59n ± 0%   +342.05% (p=0.000 n=10)
Delete/Delete_from_1_000          41.12n ± 0%   125.65n ± 0%   +205.57% (p=0.000 n=10)
Delete/Delete_from_10_000         41.15n ± 0%   145.25n ± 0%   +253.02% (p=0.000 n=10)
Delete/Delete_from_100_000        41.12n ± 0%   261.60n ± 2%   +536.26% (p=0.000 n=10)
Delete/Delete_from_1_000_000      43.83n ± 0%   298.00n ± 0%   +579.90% (p=0.000 n=10)
geomean                           43.64n         261.6n        +499.50%
```

## lookup (longest-prefix-match)

In the lookup, `art` is the champion, closely followed by `bart`. 

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │            art/lookup.bm             │
                                    │     sec/op     │    sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4             61.92n ± 22%   46.12n ± 12%  -25.51% (p=0.005 n=10)
LpmTier1Pfxs/RandomMatchIP6             59.40n ± 51%   45.26n ± 13%  -23.80% (p=0.023 n=10)
LpmTier1Pfxs/RandomMissIP4              78.69n ± 46%   28.18n ±  0%  -64.19% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              35.34n ±  1%   27.81n ±  0%  -21.31% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     83.16n ± 24%   46.06n ± 12%  -44.61% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     80.39n ± 44%   45.02n ± 17%  -44.00% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4     109.15n ± 29%   29.87n ±  7%  -72.63% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      84.91n ± 30%   29.80n ±  7%  -64.91% (p=0.000 n=10)
geomean                                 70.83n         36.31n        -48.74%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │            cidrtree/lookup.bm             │
                                    │     sec/op     │     sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4             61.92n ± 22%   1177.50n ±  24%  +1801.65% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             59.40n ± 51%   1174.00n ±  19%  +1876.43% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              78.69n ± 46%   1321.50n ±  23%  +1579.48% (p=0.002 n=10)
LpmTier1Pfxs/RandomMissIP6              35.34n ±  1%     67.13n ± 579%    +89.94% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     83.16n ± 24%   1083.50n ±  52%  +1202.99% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     80.39n ± 44%   1519.50n ±  27%  +1790.04% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      109.2n ± 29%    1126.5n ±  17%   +932.07% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      84.91n ± 30%   1401.50n ±  31%  +1550.67% (p=0.000 n=10)
geomean                                 70.83n           866.8n         +1123.76%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │           critbitgo/lookup.bm           │
                                    │     sec/op     │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             61.92n ± 22%   352.65n ± 11%   +469.53% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             59.40n ± 51%   509.85n ± 18%   +758.33% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              78.69n ± 46%   696.15n ± 40%   +784.73% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              35.34n ±  1%   695.45n ± 14%  +1867.88% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     83.16n ± 24%   390.05n ± 39%   +369.06% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     80.39n ± 44%   408.45n ± 18%   +408.05% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      109.2n ± 29%    466.3n ± 39%   +327.21% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      84.91n ± 30%   454.15n ± 20%   +434.89% (p=0.000 n=10)
geomean                                 70.83n          482.5n         +581.17%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │           lpmtrie/lookup.bm            │
                                    │     sec/op     │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4             61.92n ± 22%   236.50n ± 16%  +281.94% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             59.40n ± 51%   264.50n ± 58%  +345.29% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              78.69n ± 46%   238.90n ± 66%  +203.62% (p=0.004 n=10)
LpmTier1Pfxs/RandomMissIP6              35.34n ±  1%   106.03n ± 21%  +200.04% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     83.16n ± 24%   259.30n ±  8%  +211.83% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     80.39n ± 44%   206.45n ±  9%  +156.79% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      109.2n ± 29%    291.4n ±  7%  +167.02% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      84.91n ± 30%   230.60n ±  9%  +171.60% (p=0.000 n=10)
geomean                                 70.83n          221.1n        +212.08%
```
