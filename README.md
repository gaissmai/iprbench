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
cpu: Intel(R) Core(TM) i5-7500T CPU @ 2.70GHz
                         │ bart/size.bm │                art/size.bm                 │
                         │    bytes     │     bytes       vs base                    │
Tier1PfxSize/100           42.73Ki ± 0%    802.97Ki ± 0%  +1779.32% (p=0.000 n=10)
Tier1PfxSize/1_000         288.9Ki ± 0%    7420.4Ki ± 0%  +2468.17% (p=0.000 n=10)
Tier1PfxSize/10_000        1.822Mi ± 0%    47.172Mi ± 0%  +2488.91% (p=0.000 n=10)
Tier1PfxSize/100_000       7.545Mi ± 0%   160.300Mi ± 0%  +2024.67% (p=0.000 n=10)
Tier1PfxSize/1_000_000     34.33Mi ± 0%    378.23Mi ± 0%  +1001.89% (p=0.000 n=10)
RandomPfx4Size/100         43.32Ki ± 0%    700.22Ki ± 0%  +1516.38% (p=0.000 n=10)
RandomPfx4Size/1_000       299.8Ki ± 0%    7159.0Ki ± 0%  +2288.02% (p=0.000 n=10)
RandomPfx4Size/10_000      2.402Mi ± 0%    57.506Mi ± 0%  +2294.21% (p=0.000 n=10)
RandomPfx4Size/100_000     18.27Mi ± 0%    522.88Mi ± 0%  +2762.35% (p=0.000 n=10)
RandomPfx4Size/1_000_000   109.8Mi ± 0%
RandomPfx6Size/100         146.4Ki ± 0%     757.6Ki ± 0%   +417.50% (p=0.000 n=10)
RandomPfx6Size/1_000       1.273Mi ± 0%     7.477Mi ± 0%   +487.21% (p=0.000 n=10)
RandomPfx6Size/10_000      12.23Mi ± 0%     65.28Mi ± 0%   +434.00% (p=0.000 n=10)
RandomPfx6Size/100_000     115.2Mi ± 0%     747.9Mi ± 0%   +549.22% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1.049Gi ± 0%
RandomPfxSize/100          58.73Ki ± 0%    675.12Ki ± 0%  +1049.61% (p=0.000 n=10)
RandomPfxSize/1_000        487.7Ki ± 0%    7426.8Ki ± 0%  +1422.90% (p=0.000 n=10)
RandomPfxSize/10_000       4.380Mi ± 0%    59.386Mi ± 0%  +1255.94% (p=0.000 n=10)
RandomPfxSize/100_000      39.01Mi ± 0%    554.78Mi ± 0%  +1322.22% (p=0.000 n=10)
RandomPfxSize/1_000_000    309.2Mi ± 0%
geomean                    3.597Mi          22.53Mi       +1291.59%                ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-7500T CPU @ 2.70GHz
                         │  bart/size.bm  │           cidrtree/size.bm            │
                         │     bytes      │    bytes      vs base                 │
Tier1PfxSize/100             42.73Ki ± 0%   12.47Ki ± 0%   -70.82% (p=0.000 n=10)
Tier1PfxSize/1_000          288.94Ki ± 0%   81.51Ki ± 0%   -71.79% (p=0.000 n=10)
Tier1PfxSize/10_000         1865.8Ki ± 0%   784.6Ki ± 0%   -57.95% (p=0.000 n=10)
Tier1PfxSize/100_000         7.545Mi ± 0%   7.633Mi ± 0%    +1.17% (p=0.000 n=10)
Tier1PfxSize/1_000_000       34.33Mi ± 0%   76.30Mi ± 0%  +122.27% (p=0.000 n=10)
RandomPfx4Size/100           43.32Ki ± 0%   11.58Ki ± 0%   -73.27% (p=0.000 n=10)
RandomPfx4Size/1_000        299.79Ki ± 0%   81.51Ki ± 0%   -72.81% (p=0.000 n=10)
RandomPfx4Size/10_000       2459.5Ki ± 0%   784.6Ki ± 0%   -68.10% (p=0.000 n=10)
RandomPfx4Size/100_000      18.268Mi ± 0%   7.633Mi ± 0%   -58.22% (p=0.000 n=10)
RandomPfx4Size/1_000_000    109.79Mi ± 0%   76.30Mi ± 0%   -30.50% (p=0.000 n=10)
RandomPfx6Size/100          146.39Ki ± 0%   11.58Ki ± 0%   -92.09% (p=0.000 n=10)
RandomPfx6Size/1_000       1303.84Ki ± 0%   81.51Ki ± 0%   -93.75% (p=0.000 n=10)
RandomPfx6Size/10_000      12518.4Ki ± 0%   784.6Ki ± 0%   -93.73% (p=0.000 n=10)
RandomPfx6Size/100_000     115.196Mi ± 0%   7.633Mi ± 0%   -93.37% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1074.54Mi ± 0%   76.30Mi ± 0%   -92.90% (p=0.000 n=10)
RandomPfxSize/100            58.73Ki ± 0%   11.58Ki ± 0%   -80.28% (p=0.000 n=10)
RandomPfxSize/1_000         487.67Ki ± 0%   81.51Ki ± 0%   -83.29% (p=0.000 n=10)
RandomPfxSize/10_000        4484.8Ki ± 0%   784.6Ki ± 0%   -82.50% (p=0.000 n=10)
RandomPfxSize/100_000       39.008Mi ± 0%   7.633Mi ± 0%   -80.43% (p=0.000 n=10)
RandomPfxSize/1_000_000     309.16Mi ± 0%   76.30Mi ± 0%   -75.32% (p=0.000 n=10)
geomean                      3.597Mi        856.4Ki        -76.75%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-7500T CPU @ 2.70GHz
                         │ bart/size.bm  │           critbitgo/size.bm            │
                         │     bytes     │     bytes      vs base                 │
Tier1PfxSize/100            42.73Ki ± 0%    15.64Ki ± 0%   -63.39% (p=0.000 n=10)
Tier1PfxSize/1_000          288.9Ki ± 0%    115.2Ki ± 0%   -60.12% (p=0.000 n=10)
Tier1PfxSize/10_000         1.822Mi ± 0%    1.094Mi ± 0%   -39.98% (p=0.000 n=10)
Tier1PfxSize/100_000        7.545Mi ± 0%   10.913Mi ± 0%   +44.64% (p=0.000 n=10)
Tier1PfxSize/1_000_000      34.33Mi ± 0%   109.12Mi ± 0%  +217.88% (p=0.000 n=10)
RandomPfx4Size/100          43.32Ki ± 0%    14.67Ki ± 0%   -66.13% (p=0.000 n=10)
RandomPfx4Size/1_000        299.8Ki ± 0%    112.7Ki ± 0%   -62.40% (p=0.000 n=10)
RandomPfx4Size/10_000       2.402Mi ± 0%    1.071Mi ± 0%   -55.39% (p=0.000 n=10)
RandomPfx4Size/100_000      18.27Mi ± 0%    10.68Mi ± 0%   -41.51% (p=0.000 n=10)
RandomPfx4Size/1_000_000    109.8Mi ± 0%    106.8Mi ± 0%    -2.71% (p=0.000 n=10)
RandomPfx6Size/100         146.39Ki ± 0%    16.22Ki ± 0%   -88.92% (p=0.000 n=10)
RandomPfx6Size/1_000       1303.8Ki ± 0%    128.3Ki ± 0%   -90.16% (p=0.000 n=10)
RandomPfx6Size/10_000      12.225Mi ± 0%    1.224Mi ± 0%   -89.99% (p=0.000 n=10)
RandomPfx6Size/100_000     115.20Mi ± 0%    12.21Mi ± 0%   -89.40% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1074.5Mi ± 0%    122.1Mi ± 0%   -88.64% (p=0.000 n=10)
RandomPfxSize/100           58.73Ki ± 0%    14.94Ki ± 0%   -74.56% (p=0.000 n=10)
RandomPfxSize/1_000         487.7Ki ± 0%    115.8Ki ± 0%   -76.26% (p=0.000 n=10)
RandomPfxSize/10_000        4.380Mi ± 0%    1.102Mi ± 0%   -74.84% (p=0.000 n=10)
RandomPfxSize/100_000       39.01Mi ± 0%    10.99Mi ± 0%   -71.83% (p=0.000 n=10)
RandomPfxSize/1_000_000     309.2Mi ± 0%    109.9Mi ± 0%   -64.46% (p=0.000 n=10)
geomean                     3.597Mi         1.193Mi        -66.83%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-7500T CPU @ 2.70GHz
                         │ bart/size.bm  │            lpmtrie/size.bm             │
                         │     bytes     │     bytes      vs base                 │
Tier1PfxSize/100            42.73Ki ± 0%    24.25Ki ± 0%   -43.24% (p=0.000 n=10)
Tier1PfxSize/1_000          288.9Ki ± 0%    202.3Ki ± 0%   -29.97% (p=0.000 n=10)
Tier1PfxSize/10_000         1.822Mi ± 0%    1.942Mi ± 0%    +6.58% (p=0.000 n=10)
Tier1PfxSize/100_000        7.545Mi ± 0%   19.330Mi ± 0%  +156.21% (p=0.000 n=10)
Tier1PfxSize/1_000_000      34.33Mi ± 0%   189.92Mi ± 0%  +453.28% (p=0.000 n=10)
RandomPfx4Size/100          43.32Ki ± 0%    23.20Ki ± 0%   -46.44% (p=0.000 n=10)
RandomPfx4Size/1_000        299.8Ki ± 0%    198.2Ki ± 0%   -33.88% (p=0.000 n=10)
RandomPfx4Size/10_000       2.402Mi ± 0%    1.897Mi ± 0%   -21.03% (p=0.000 n=10)
RandomPfx4Size/100_000      18.27Mi ± 0%    18.25Mi ± 0%    -0.11% (p=0.000 n=10)
RandomPfx4Size/1_000_000    109.8Mi ± 0%    163.5Mi ± 0%   +48.95% (p=0.000 n=10)
RandomPfx6Size/100         146.39Ki ± 0%    25.31Ki ± 0%   -82.71% (p=0.000 n=10)
RandomPfx6Size/1_000       1303.8Ki ± 0%    221.6Ki ± 0%   -83.01% (p=0.000 n=10)
RandomPfx6Size/10_000      12.225Mi ± 0%    2.125Mi ± 0%   -82.62% (p=0.000 n=10)
RandomPfx6Size/100_000     115.20Mi ± 0%    20.91Mi ± 0%   -81.85% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1074.5Mi ± 0%    203.9Mi ± 0%   -81.02% (p=0.000 n=10)
RandomPfxSize/100           58.73Ki ± 0%    23.56Ki ± 0%   -59.88% (p=0.000 n=10)
RandomPfxSize/1_000         487.7Ki ± 0%    203.4Ki ± 0%   -58.30% (p=0.000 n=10)
RandomPfxSize/10_000        4.380Mi ± 0%    1.945Mi ± 0%   -55.59% (p=0.000 n=10)
RandomPfxSize/100_000       39.01Mi ± 0%    18.78Mi ± 0%   -51.85% (p=0.000 n=10)
RandomPfxSize/1_000_000     309.2Mi ± 0%    171.9Mi ± 0%   -44.39% (p=0.000 n=10)
geomean                     3.597Mi         2.009Mi        -44.14%
```

## update, insert/delete

When it comes to updates, `art` and `bart` are far superior
to the other algorithms, only `critbitgo` comes close to playing in the same league .

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-7500T CPU @ 2.70GHz
                             │ bart/update.bm │             art/update.bm             │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100            62.86n ± 2%   106.90n ± 1%   +70.07% (p=0.000 n=10)
Insert/Insert_into_1_000          62.69n ± 2%   131.95n ± 1%  +110.50% (p=0.000 n=10)
Insert/Insert_into_10_000         62.62n ± 1%   134.05n ± 1%  +114.09% (p=0.000 n=10)
Insert/Insert_into_100_000        62.81n ± 1%   124.10n ± 2%   +97.56% (p=0.000 n=10)
Insert/Insert_into_1_000_000      63.52n ± 3%   123.40n ± 1%   +94.27% (p=0.000 n=10)
Delete/Delete_from_100            24.24n ± 4%    32.91n ± 4%   +35.79% (p=0.000 n=10)
Delete/Delete_from_1_000          49.31n ± 1%    61.47n ± 4%   +24.66% (p=0.000 n=10)
Delete/Delete_from_10_000         49.24n ± 0%    60.41n ± 2%   +22.68% (p=0.000 n=10)
Delete/Delete_from_100_000        49.40n ± 2%    61.88n ± 5%   +25.26% (p=0.000 n=10)
Delete/Delete_from_1_000_000      50.63n ± 1%    59.45n ± 1%   +17.40% (p=0.000 n=10)
geomean                           52.01n         81.56n        +56.80%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-7500T CPU @ 2.70GHz
                             │ bart/update.bm │           cidrtree/update.bm            │
                             │     sec/op     │    sec/op      vs base                  │
Insert/Insert_into_100            62.86n ± 2%   1143.50n ± 0%  +1719.27% (p=0.000 n=10)
Insert/Insert_into_1_000          62.69n ± 2%   1492.50n ± 0%  +2280.95% (p=0.000 n=10)
Insert/Insert_into_10_000         62.62n ± 1%   2472.00n ± 0%  +3847.94% (p=0.000 n=10)
Insert/Insert_into_100_000        62.81n ± 1%   2345.00n ± 4%  +3633.18% (p=0.000 n=10)
Insert/Insert_into_1_000_000      63.52n ± 3%   2627.00n ± 3%  +4035.71% (p=0.000 n=10)
Delete/Delete_from_100            24.24n ± 4%   1674.50n ± 2%  +6808.00% (p=0.000 n=10)
Delete/Delete_from_1_000          49.31n ± 1%   2764.50n ± 3%  +5505.80% (p=0.000 n=10)
Delete/Delete_from_10_000         49.24n ± 0%   3962.50n ± 2%  +7946.50% (p=0.000 n=10)
Delete/Delete_from_100_000        49.40n ± 2%   4306.00n ± 1%  +8615.72% (p=0.000 n=10)
Delete/Delete_from_1_000_000      50.63n ± 1%   3746.50n ± 1%  +7299.03% (p=0.000 n=10)
geomean                           52.01n          2.447µ       +4604.05%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-7500T CPU @ 2.70GHz
                             │ bart/update.bm │          critbitgo/update.bm          │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100            62.86n ± 2%   115.00n ± 1%   +82.96% (p=0.000 n=10)
Insert/Insert_into_1_000          62.69n ± 2%   127.50n ± 0%  +103.40% (p=0.000 n=10)
Insert/Insert_into_10_000         62.62n ± 1%   134.05n ± 0%  +114.09% (p=0.000 n=10)
Insert/Insert_into_100_000        62.81n ± 1%   154.30n ± 1%  +145.64% (p=0.000 n=10)
Insert/Insert_into_1_000_000      63.52n ± 3%   150.85n ± 0%  +137.48% (p=0.000 n=10)
Delete/Delete_from_100            24.24n ± 4%   101.45n ± 0%  +318.52% (p=0.000 n=10)
Delete/Delete_from_1_000          49.31n ± 1%   109.35n ± 1%  +121.74% (p=0.000 n=10)
Delete/Delete_from_10_000         49.24n ± 0%   116.40n ± 1%  +136.37% (p=0.000 n=10)
Delete/Delete_from_100_000        49.40n ± 2%   139.10n ± 2%  +181.55% (p=0.000 n=10)
Delete/Delete_from_1_000_000      50.63n ± 1%   135.40n ± 0%  +167.40% (p=0.000 n=10)
geomean                           52.01n         127.2n       +144.63%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-7500T CPU @ 2.70GHz
                             │ bart/update.bm │           lpmtrie/update.bm            │
                             │     sec/op     │    sec/op      vs base                 │
Insert/Insert_into_100            62.86n ± 2%   322.00n ±  1%  +412.29% (p=0.000 n=10)
Insert/Insert_into_1_000          62.69n ± 2%   383.15n ±  4%  +511.23% (p=0.000 n=10)
Insert/Insert_into_10_000         62.62n ± 1%   419.75n ±  4%  +570.37% (p=0.000 n=10)
Insert/Insert_into_100_000        62.81n ± 1%   522.70n ±  5%  +732.13% (p=0.000 n=10)
Insert/Insert_into_1_000_000      63.52n ± 3%   657.60n ± 13%  +935.26% (p=0.000 n=10)
Delete/Delete_from_100            24.24n ± 4%    77.22n ±  2%  +218.56% (p=0.000 n=10)
Delete/Delete_from_1_000          49.31n ± 1%   130.10n ±  3%  +163.81% (p=0.000 n=10)
Delete/Delete_from_10_000         49.24n ± 0%   153.30n ±  4%  +211.30% (p=0.000 n=10)
Delete/Delete_from_100_000        49.40n ± 2%   265.50n ±  4%  +437.40% (p=0.000 n=10)
Delete/Delete_from_1_000_000      50.63n ± 1%   313.25n ±  1%  +518.64% (p=0.000 n=10)
geomean                           52.01n         272.8n        +424.41%
```

## lookup (longest-prefix-match)

In the lookup, `art` is the champion, closely followed by `bart`. 

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-7500T CPU @ 2.70GHz
                                    │ bart/lookup.bm │            art/lookup.bm             │
                                    │     sec/op     │    sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4             61.77n ± 36%   47.87n ± 13%  -22.50% (p=0.001 n=10)
LpmTier1Pfxs/RandomMatchIP6             67.49n ± 34%   46.45n ± 17%        ~ (p=0.089 n=10)
LpmTier1Pfxs/RandomMissIP4              77.89n ± 40%   29.40n ±  7%  -62.26% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              39.59n ± 81%   28.67n ±  1%  -27.60% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     86.36n ± 20%   46.88n ± 22%  -45.71% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     76.73n ± 49%   46.38n ± 18%  -39.55% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4     115.55n ± 30%   30.98n ±  7%  -73.18% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6     104.55n ± 24%   28.90n ±  6%  -72.36% (p=0.000 n=10)
geomean                                 75.26n         37.18n        -50.60%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-7500T CPU @ 2.70GHz
                                    │ bart/lookup.bm │             cidrtree/lookup.bm             │
                                    │     sec/op     │      sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4             61.77n ± 36%   1107.50n ±   25%  +1692.94% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             67.49n ± 34%   1211.50n ±    5%  +1695.08% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              77.89n ± 40%   1247.50n ±   97%          ~ (p=0.143 n=10)
LpmTier1Pfxs/RandomMissIP6              39.59n ± 81%     56.83n ± 2478%    +43.55% (p=0.003 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     86.36n ± 20%   1294.50n ±  106%  +1398.96% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     76.73n ± 49%   1963.50n ±   46%  +2458.81% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      115.6n ± 30%    1223.5n ±   10%   +958.85% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      104.6n ± 24%    1754.5n ±   13%  +1578.14% (p=0.000 n=10)
geomean                                 75.26n           921.3n          +1124.15%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-7500T CPU @ 2.70GHz
                                    │ bart/lookup.bm │           critbitgo/lookup.bm           │
                                    │     sec/op     │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             61.77n ± 36%   378.45n ± 16%   +512.68% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             67.49n ± 34%   530.55n ± 23%   +686.12% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              77.89n ± 40%   829.45n ± 58%   +964.83% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              39.59n ± 81%   562.65n ± 26%  +1321.19% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     86.36n ± 20%   368.20n ± 20%   +326.35% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     76.73n ± 49%   383.45n ±  9%   +399.71% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      115.6n ± 30%    543.2n ± 53%   +370.14% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      104.6n ± 24%    368.5n ± 32%   +252.46% (p=0.000 n=10)
geomean                                 75.26n          476.3n         +532.86%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-7500T CPU @ 2.70GHz
                                    │ bart/lookup.bm │           lpmtrie/lookup.bm            │
                                    │     sec/op     │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4             61.77n ± 36%   250.25n ± 14%  +305.13% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             67.49n ± 34%   266.35n ± 67%  +294.65% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              77.89n ± 40%   247.00n ± 69%  +217.09% (p=0.009 n=10)
LpmTier1Pfxs/RandomMissIP6              39.59n ± 81%    93.39n ± 28%  +135.91% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     86.36n ± 20%   288.35n ±  3%  +233.89% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     76.73n ± 49%   213.95n ±  9%  +178.82% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      115.6n ± 30%    293.8n ±  3%  +154.22% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      104.6n ± 24%    225.9n ± 12%  +116.07% (p=0.000 n=10)
geomean                                 75.26n          223.8n        +197.38%
```
