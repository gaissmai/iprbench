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
RandomPfx4Size/100         41.70Ki ± 0%    693.84Ki ± 0%  +1563.77% (p=0.000 n=10)
RandomPfx4Size/1_000       299.2Ki ± 0%    7197.3Ki ± 0%  +2305.29% (p=0.000 n=10)
RandomPfx4Size/10_000      2.416Mi ± 0%    58.104Mi ± 0%  +2304.95% (p=0.000 n=10)
RandomPfx4Size/100_000     18.33Mi ± 0%    522.30Mi ± 0%  +2748.81% (p=0.000 n=10)
RandomPfx4Size/1_000_000   109.8Mi ± 0%
RandomPfx6Size/100         155.4Ki ± 0%     732.1Ki ± 0%   +371.22% (p=0.000 n=10)
RandomPfx6Size/1_000       1.275Mi ± 0%     7.514Mi ± 0%   +489.48% (p=0.000 n=10)
RandomPfx6Size/10_000      12.23Mi ± 0%     65.29Mi ± 0%   +433.84% (p=0.000 n=10)
RandomPfx6Size/100_000     115.5Mi ± 0%     748.4Mi ± 0%   +547.77% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1.049Gi ± 0%
RandomPfxSize/100          68.03Ki ± 0%    707.00Ki ± 0%   +939.23% (p=0.000 n=10)
RandomPfxSize/1_000        507.7Ki ± 0%    7452.3Ki ± 0%  +1367.98% (p=0.000 n=10)
RandomPfxSize/10_000       4.347Mi ± 0%    59.560Mi ± 0%  +1270.01% (p=0.000 n=10)
RandomPfxSize/100_000      38.93Mi ± 0%    555.10Mi ± 0%  +1325.79% (p=0.000 n=10)
RandomPfxSize/1_000_000    309.1Mi ± 0%
geomean                    3.636Mi          22.57Mi       +1276.66%                ¹
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
RandomPfx4Size/100           41.70Ki ± 0%   11.58Ki ± 0%   -72.24% (p=0.000 n=10)
RandomPfx4Size/1_000        299.23Ki ± 0%   81.51Ki ± 0%   -72.76% (p=0.000 n=10)
RandomPfx4Size/10_000       2474.0Ki ± 0%   784.6Ki ± 0%   -68.28% (p=0.000 n=10)
RandomPfx4Size/100_000      18.334Mi ± 0%   7.633Mi ± 0%   -58.37% (p=0.000 n=10)
RandomPfx4Size/1_000_000    109.79Mi ± 0%   76.30Mi ± 0%   -30.51% (p=0.000 n=10)
RandomPfx6Size/100          155.36Ki ± 0%   11.58Ki ± 0%   -92.55% (p=0.000 n=10)
RandomPfx6Size/1_000       1305.30Ki ± 0%   81.51Ki ± 0%   -93.76% (p=0.000 n=10)
RandomPfx6Size/10_000      12523.4Ki ± 0%   784.6Ki ± 0%   -93.73% (p=0.000 n=10)
RandomPfx6Size/100_000     115.541Mi ± 0%   7.633Mi ± 0%   -93.39% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1074.54Mi ± 0%   76.30Mi ± 0%   -92.90% (p=0.000 n=10)
RandomPfxSize/100            68.03Ki ± 0%   11.58Ki ± 0%   -82.98% (p=0.000 n=10)
RandomPfxSize/1_000         507.66Ki ± 0%   81.51Ki ± 0%   -83.94% (p=0.000 n=10)
RandomPfxSize/10_000        4451.8Ki ± 0%   784.6Ki ± 0%   -82.37% (p=0.000 n=10)
RandomPfxSize/100_000       38.933Mi ± 0%   7.633Mi ± 0%   -80.40% (p=0.000 n=10)
RandomPfxSize/1_000_000     309.15Mi ± 0%   76.30Mi ± 0%   -75.32% (p=0.000 n=10)
geomean                      3.636Mi        856.4Ki        -77.00%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm  │           critbitgo/size.bm            │
                         │     bytes     │     bytes      vs base                 │
Tier1PfxSize/100            42.81Ki ± 0%    15.55Ki ± 0%   -63.69% (p=0.000 n=10)
Tier1PfxSize/1_000          288.9Ki ± 0%    115.2Ki ± 0%   -60.12% (p=0.000 n=10)
Tier1PfxSize/10_000         1.822Mi ± 0%    1.094Mi ± 0%   -39.98% (p=0.000 n=10)
Tier1PfxSize/100_000        7.545Mi ± 0%   10.913Mi ± 0%   +44.64% (p=0.000 n=10)
Tier1PfxSize/1_000_000      34.33Mi ± 0%   109.12Mi ± 0%  +217.88% (p=0.000 n=10)
RandomPfx4Size/100          41.70Ki ± 0%    14.67Ki ± 0%   -64.82% (p=0.000 n=10)
RandomPfx4Size/1_000        299.2Ki ± 0%    112.7Ki ± 0%   -62.33% (p=0.000 n=10)
RandomPfx4Size/10_000       2.416Mi ± 0%    1.071Mi ± 0%   -55.65% (p=0.000 n=10)
RandomPfx4Size/100_000      18.33Mi ± 0%    10.68Mi ± 0%   -41.72% (p=0.000 n=10)
RandomPfx4Size/1_000_000    109.8Mi ± 0%    106.8Mi ± 0%    -2.71% (p=0.000 n=10)
RandomPfx6Size/100         155.36Ki ± 0%    16.22Ki ± 0%   -89.56% (p=0.000 n=10)
RandomPfx6Size/1_000       1305.3Ki ± 0%    128.3Ki ± 0%   -90.17% (p=0.000 n=10)
RandomPfx6Size/10_000      12.230Mi ± 0%    1.224Mi ± 0%   -89.99% (p=0.000 n=10)
RandomPfx6Size/100_000     115.54Mi ± 0%    12.21Mi ± 0%   -89.43% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1074.5Mi ± 0%    122.1Mi ± 0%   -88.64% (p=0.000 n=10)
RandomPfxSize/100           68.03Ki ± 0%    14.94Ki ± 0%   -78.04% (p=0.000 n=10)
RandomPfxSize/1_000         507.7Ki ± 0%    115.8Ki ± 0%   -77.19% (p=0.000 n=10)
RandomPfxSize/10_000        4.347Mi ± 0%    1.102Mi ± 0%   -74.65% (p=0.000 n=10)
RandomPfxSize/100_000       38.93Mi ± 0%    10.99Mi ± 0%   -71.77% (p=0.000 n=10)
RandomPfxSize/1_000_000     309.1Mi ± 0%    109.9Mi ± 0%   -64.46% (p=0.000 n=10)
geomean                     3.636Mi         1.193Mi        -67.19%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm  │            lpmtrie/size.bm             │
                         │     bytes     │     bytes      vs base                 │
Tier1PfxSize/100            42.81Ki ± 0%    24.25Ki ± 0%   -43.36% (p=0.000 n=10)
Tier1PfxSize/1_000          288.9Ki ± 0%    202.3Ki ± 0%   -29.97% (p=0.000 n=10)
Tier1PfxSize/10_000         1.822Mi ± 0%    1.942Mi ± 0%    +6.58% (p=0.000 n=10)
Tier1PfxSize/100_000        7.545Mi ± 0%   19.330Mi ± 0%  +156.21% (p=0.000 n=10)
Tier1PfxSize/1_000_000      34.33Mi ± 0%   189.92Mi ± 0%  +453.28% (p=0.000 n=10)
RandomPfx4Size/100          41.70Ki ± 0%    23.20Ki ± 0%   -44.36% (p=0.000 n=10)
RandomPfx4Size/1_000        299.2Ki ± 0%    198.9Ki ± 0%   -33.53% (p=0.000 n=10)
RandomPfx4Size/10_000       2.416Mi ± 0%    1.893Mi ± 0%   -21.65% (p=0.000 n=10)
RandomPfx4Size/100_000      18.33Mi ± 0%    18.29Mi ± 0%    -0.26% (p=0.000 n=10)
RandomPfx4Size/1_000_000    109.8Mi ± 0%    163.5Mi ± 0%   +48.92% (p=0.000 n=10)
RandomPfx6Size/100         155.36Ki ± 0%    25.53Ki ± 0%   -83.57% (p=0.000 n=10)
RandomPfx6Size/1_000       1305.3Ki ± 0%    222.0Ki ± 0%   -82.99% (p=0.000 n=10)
RandomPfx6Size/10_000      12.230Mi ± 0%    2.129Mi ± 0%   -82.59% (p=0.000 n=10)
RandomPfx6Size/100_000     115.54Mi ± 0%    20.93Mi ± 0%   -81.89% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1074.5Mi ± 0%    203.9Mi ± 0%   -81.03% (p=0.000 n=10)
RandomPfxSize/100           68.03Ki ± 0%    23.56Ki ± 0%   -65.37% (p=0.000 n=10)
RandomPfxSize/1_000         507.7Ki ± 0%    202.9Ki ± 0%   -60.04% (p=0.000 n=10)
RandomPfxSize/10_000        4.347Mi ± 0%    1.942Mi ± 0%   -55.34% (p=0.000 n=10)
RandomPfxSize/100_000       38.93Mi ± 0%    18.78Mi ± 0%   -51.75% (p=0.000 n=10)
RandomPfxSize/1_000_000     309.1Mi ± 0%    171.9Mi ± 0%   -44.41% (p=0.000 n=10)
geomean                     3.636Mi         2.011Mi        -44.70%
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
Insert/Insert_into_100            55.34n ± 0%   104.25n ± 1%   +88.36% (p=0.000 n=10)
Insert/Insert_into_1_000          55.33n ± 0%   125.30n ± 1%  +126.44% (p=0.000 n=10)
Insert/Insert_into_10_000         55.34n ± 0%   129.45n ± 1%  +133.92% (p=0.000 n=10)
Insert/Insert_into_100_000        55.96n ± 0%   120.85n ± 2%  +115.96% (p=0.000 n=10)
Insert/Insert_into_1_000_000      57.49n ± 1%   119.70n ± 0%  +108.21% (p=0.000 n=10)
Delete/Delete_from_100            17.31n ± 0%    32.01n ± 3%   +84.92% (p=0.000 n=10)
Delete/Delete_from_1_000          41.58n ± 0%    58.37n ± 0%   +40.38% (p=0.000 n=10)
Delete/Delete_from_10_000         41.58n ± 3%    59.21n ± 0%   +42.37% (p=0.000 n=10)
Delete/Delete_from_100_000        42.78n ± 0%    58.66n ± 0%   +37.12% (p=0.000 n=10)
Delete/Delete_from_1_000_000      45.35n ± 1%    58.22n ± 0%   +28.38% (p=0.000 n=10)
geomean                           44.67n         78.81n        +76.41%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │            cidrtree/update.bm            │
                             │     sec/op     │    sec/op      vs base                   │
Insert/Insert_into_100            55.34n ± 0%   1195.50n ± 4%   +2060.09% (p=0.000 n=10)
Insert/Insert_into_1_000          55.33n ± 0%   1250.00n ± 1%   +2158.97% (p=0.000 n=10)
Insert/Insert_into_10_000         55.34n ± 0%   1824.00n ± 4%   +3195.99% (p=0.000 n=10)
Insert/Insert_into_100_000        55.96n ± 0%   2683.50n ± 3%   +4695.39% (p=0.000 n=10)
Insert/Insert_into_1_000_000      57.49n ± 1%   2329.50n ± 0%   +3952.01% (p=0.000 n=10)
Delete/Delete_from_100            17.31n ± 0%   1071.00n ± 0%   +6087.18% (p=0.000 n=10)
Delete/Delete_from_1_000          41.58n ± 0%   3481.00n ± 0%   +8271.81% (p=0.000 n=10)
Delete/Delete_from_10_000         41.58n ± 3%   4059.00n ± 0%   +9660.73% (p=0.000 n=10)
Delete/Delete_from_100_000        42.78n ± 0%   4881.50n ± 0%  +11310.71% (p=0.000 n=10)
Delete/Delete_from_1_000_000      45.35n ± 1%   4461.00n ± 0%   +9736.82% (p=0.000 n=10)
geomean                           44.67n          2.371µ        +5207.41%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │          critbitgo/update.bm           │
                             │     sec/op     │    sec/op      vs base                 │
Insert/Insert_into_100            55.34n ± 0%   111.25n ±  0%  +101.01% (p=0.000 n=10)
Insert/Insert_into_1_000          55.33n ± 0%   123.70n ±  0%  +123.55% (p=0.000 n=10)
Insert/Insert_into_10_000         55.34n ± 0%   127.10n ±  1%  +129.67% (p=0.000 n=10)
Insert/Insert_into_100_000        55.96n ± 0%   157.45n ±  1%  +181.36% (p=0.000 n=10)
Insert/Insert_into_1_000_000      57.49n ± 1%   153.55n ±  0%  +167.09% (p=0.000 n=10)
Delete/Delete_from_100            17.31n ± 0%    98.14n ± 11%  +466.96% (p=0.000 n=10)
Delete/Delete_from_1_000          41.58n ± 0%   117.05n ±  6%  +181.51% (p=0.000 n=10)
Delete/Delete_from_10_000         41.58n ± 3%   120.35n ± 10%  +189.41% (p=0.000 n=10)
Delete/Delete_from_100_000        42.78n ± 0%   142.60n ±  6%  +233.33% (p=0.000 n=10)
Delete/Delete_from_1_000_000      45.35n ± 1%   147.15n ±  1%  +224.48% (p=0.000 n=10)
geomean                           44.67n         128.5n        +187.62%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │            lpmtrie/update.bm            │
                             │     sec/op     │    sec/op      vs base                  │
Insert/Insert_into_100            55.34n ± 0%   321.40n ± 10%   +480.72% (p=0.000 n=10)
Insert/Insert_into_1_000          55.33n ± 0%   346.10n ±  2%   +525.46% (p=0.000 n=10)
Insert/Insert_into_10_000         55.34n ± 0%   370.70n ±  1%   +569.86% (p=0.000 n=10)
Insert/Insert_into_100_000        55.96n ± 0%   528.40n ±  3%   +844.25% (p=0.000 n=10)
Insert/Insert_into_1_000_000      57.49n ± 1%   691.85n ±  2%  +1103.43% (p=0.000 n=10)
Delete/Delete_from_100            17.31n ± 0%    76.59n ±  1%   +342.43% (p=0.000 n=10)
Delete/Delete_from_1_000          41.58n ± 0%   125.50n ±  1%   +201.83% (p=0.000 n=10)
Delete/Delete_from_10_000         41.58n ± 3%   145.20n ±  0%   +249.16% (p=0.000 n=10)
Delete/Delete_from_100_000        42.78n ± 0%   254.15n ±  2%   +494.09% (p=0.000 n=10)
Delete/Delete_from_1_000_000      45.35n ± 1%   319.55n ±  0%   +604.63% (p=0.000 n=10)
geomean                           44.67n         265.0n         +493.21%
```

## lookup (longest-prefix-match)

In the lookup, `art` is the champion, closely followed by `bart`. 

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │            art/lookup.bm             │
                                    │     sec/op     │    sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4             68.63n ± 27%   46.24n ± 13%  -32.62% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             61.18n ± 49%   46.03n ± 11%  -24.75% (p=0.001 n=10)
LpmTier1Pfxs/RandomMissIP4              82.97n ± 47%   28.21n ±  1%  -65.99% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              36.13n ±  2%   28.05n ±  5%  -22.35% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     80.08n ± 25%   49.75n ±  9%  -37.87% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     80.63n ± 35%   49.63n ±  8%  -38.44% (p=0.001 n=10)
LpmRandomPfxs100_000/RandomMissIP4     116.10n ±  2%   32.25n ± 11%  -72.22% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      81.59n ± 39%   30.11n ±  8%  -63.10% (p=0.000 n=10)
geomean                                 72.58n         37.65n        -48.12%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │            cidrtree/lookup.bm            │
                                    │     sec/op     │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             68.63n ± 27%   1167.50n ± 29%  +1601.15% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             61.18n ± 49%   1483.50n ± 20%  +2324.81% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              82.97n ± 47%   1188.50n ± 47%  +1332.53% (p=0.002 n=10)
LpmTier1Pfxs/RandomMissIP6              36.13n ±  2%     68.84n ± 18%    +90.52% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     80.08n ± 25%   1121.00n ± 36%  +1299.85% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     80.63n ± 35%   1252.00n ± 35%  +1452.77% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      116.1n ±  2%    1282.5n ± 20%  +1004.65% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      81.59n ± 39%   1312.50n ± 10%  +1508.55% (p=0.000 n=10)
geomean                                 72.58n           872.1n        +1101.68%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │           critbitgo/lookup.bm           │
                                    │     sec/op     │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             68.63n ± 27%   387.80n ± 26%   +465.06% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             61.18n ± 49%   572.60n ± 18%   +835.93% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              82.97n ± 47%   899.75n ± 31%   +984.49% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              36.13n ±  2%   699.05n ± 26%  +1834.82% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     80.08n ± 25%   376.00n ± 41%   +369.53% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     80.63n ± 35%   363.75n ± 17%   +351.13% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      116.1n ±  2%    646.4n ± 19%   +456.80% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      81.59n ± 39%   476.65n ± 22%   +484.17% (p=0.000 n=10)
geomean                                 72.58n          526.3n         +625.15%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │           lpmtrie/lookup.bm            │
                                    │     sec/op     │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4             68.63n ± 27%   240.05n ± 15%  +249.77% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             61.18n ± 49%   339.40n ± 31%  +454.76% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              82.97n ± 47%   245.55n ± 67%  +195.97% (p=0.019 n=10)
LpmTier1Pfxs/RandomMissIP6              36.13n ±  2%   119.40n ± 40%  +230.47% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     80.08n ± 25%   280.70n ± 10%  +250.52% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     80.63n ± 35%   225.55n ±  7%  +179.73% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      116.1n ±  2%    309.6n ± 12%  +166.67% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      81.59n ± 39%   275.35n ± 18%  +237.46% (p=0.000 n=10)
geomean                                 72.58n          244.8n        +237.28%
```
