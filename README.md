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
Tier1PfxSize/100           42.79Ki ± 0%    802.94Ki ± 0%  +1776.50% (p=0.000 n=10)
Tier1PfxSize/1_000         288.7Ki ± 0%    7420.4Ki ± 0%  +2470.39% (p=0.000 n=10)
Tier1PfxSize/10_000        1.820Mi ± 0%    47.172Mi ± 0%  +2492.04% (p=0.000 n=10)
Tier1PfxSize/100_000       7.537Mi ± 0%   160.300Mi ± 0%  +2026.77% (p=0.000 n=10)
Tier1PfxSize/1_000_000     33.71Mi ± 0%    378.23Mi ± 0%  +1021.98% (p=0.000 n=10)
RandomPfx4Size/100         44.52Ki ± 0%    725.72Ki ± 0%  +1529.97% (p=0.000 n=10)
RandomPfx4Size/1_000       302.2Ki ± 0%    7363.0Ki ± 0%  +2336.64% (p=0.000 n=10)
RandomPfx4Size/10_000      2.415Mi ± 0%    57.979Mi ± 0%  +2300.91% (p=0.000 n=10)
RandomPfx4Size/100_000     18.21Mi ± 0%    523.92Mi ± 0%  +2776.41% (p=0.000 n=10)
RandomPfx4Size/1_000_000   109.4Mi ± 0%
RandomPfx6Size/100         157.1Ki ± 0%     744.8Ki ± 0%   +374.18% (p=0.000 n=10)
RandomPfx6Size/1_000       1.248Mi ± 0%     7.527Mi ± 0%   +503.20% (p=0.000 n=10)
RandomPfx6Size/10_000      12.34Mi ± 0%     64.77Mi ± 0%   +424.73% (p=0.000 n=10)
RandomPfx6Size/100_000     115.1Mi ± 0%     748.9Mi ± 0%   +550.52% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1.049Gi ± 0%
RandomPfxSize/100          59.52Ki ± 0%    674.72Ki ± 0%  +1033.53% (p=0.000 n=10)
RandomPfxSize/1_000        503.9Ki ± 0%    7541.5Ki ± 0%  +1396.59% (p=0.000 n=10)
RandomPfxSize/10_000       4.342Mi ± 0%    59.442Mi ± 0%  +1269.00% (p=0.000 n=10)
RandomPfxSize/100_000      38.80Mi ± 0%    555.76Mi ± 0%  +1332.45% (p=0.000 n=10)
RandomPfxSize/1_000_000    308.7Mi ± 0%
geomean                    3.616Mi          22.63Mi       +1288.73%                ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │  bart/size.bm  │           cidrtree/size.bm            │
                         │     bytes      │    bytes      vs base                 │
Tier1PfxSize/100             42.79Ki ± 0%   12.35Ki ± 0%   -71.13% (p=0.000 n=10)
Tier1PfxSize/1_000          288.69Ki ± 0%   81.51Ki ± 0%   -71.77% (p=0.000 n=10)
Tier1PfxSize/10_000         1863.5Ki ± 0%   784.6Ki ± 0%   -57.90% (p=0.000 n=10)
Tier1PfxSize/100_000         7.537Mi ± 0%   7.633Mi ± 0%    +1.27% (p=0.000 n=10)
Tier1PfxSize/1_000_000       33.71Mi ± 0%   76.30Mi ± 0%  +126.33% (p=0.000 n=10)
RandomPfx4Size/100           44.52Ki ± 0%   11.58Ki ± 0%   -74.00% (p=0.000 n=10)
RandomPfx4Size/1_000        302.18Ki ± 0%   81.51Ki ± 0%   -73.03% (p=0.000 n=10)
RandomPfx4Size/10_000       2472.8Ki ± 0%   784.6Ki ± 0%   -68.27% (p=0.000 n=10)
RandomPfx4Size/100_000      18.214Mi ± 0%   7.633Mi ± 0%   -58.09% (p=0.000 n=10)
RandomPfx4Size/1_000_000    109.41Mi ± 0%   76.30Mi ± 0%   -30.26% (p=0.000 n=10)
RandomPfx6Size/100          157.08Ki ± 0%   11.58Ki ± 0%   -92.63% (p=0.000 n=10)
RandomPfx6Size/1_000       1277.72Ki ± 0%   81.51Ki ± 0%   -93.62% (p=0.000 n=10)
RandomPfx6Size/10_000      12638.8Ki ± 0%   785.0Ki ± 0%   -93.79% (p=0.000 n=10)
RandomPfx6Size/100_000     115.128Mi ± 0%   7.633Mi ± 0%   -93.37% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1074.68Mi ± 0%   76.30Mi ± 0%   -92.90% (p=0.000 n=10)
RandomPfxSize/100            59.52Ki ± 0%   11.58Ki ± 0%   -80.55% (p=0.000 n=10)
RandomPfxSize/1_000         503.91Ki ± 0%   81.51Ki ± 0%   -83.83% (p=0.000 n=10)
RandomPfxSize/10_000        4446.2Ki ± 0%   784.6Ki ± 0%   -82.35% (p=0.000 n=10)
RandomPfxSize/100_000       38.798Mi ± 0%   7.633Mi ± 0%   -80.33% (p=0.000 n=10)
RandomPfxSize/1_000_000     308.68Mi ± 0%   76.30Mi ± 0%   -75.28% (p=0.000 n=10)
geomean                      3.616Mi        856.0Ki        -76.88%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm  │           critbitgo/size.bm            │
                         │     bytes     │     bytes      vs base                 │
Tier1PfxSize/100            42.79Ki ± 0%    15.52Ki ± 0%   -63.74% (p=0.000 n=10)
Tier1PfxSize/1_000          288.7Ki ± 0%    115.2Ki ± 0%   -60.09% (p=0.000 n=10)
Tier1PfxSize/10_000         1.820Mi ± 0%    1.094Mi ± 0%   -39.91% (p=0.000 n=10)
Tier1PfxSize/100_000        7.537Mi ± 0%   10.913Mi ± 0%   +44.78% (p=0.000 n=10)
Tier1PfxSize/1_000_000      33.71Mi ± 0%   109.12Mi ± 0%  +223.68% (p=0.000 n=10)
RandomPfx4Size/100          44.52Ki ± 0%    14.67Ki ± 0%   -67.05% (p=0.000 n=10)
RandomPfx4Size/1_000        302.2Ki ± 0%    112.7Ki ± 0%   -62.70% (p=0.000 n=10)
RandomPfx4Size/10_000       2.415Mi ± 0%    1.071Mi ± 0%   -55.63% (p=0.000 n=10)
RandomPfx4Size/100_000      18.21Mi ± 0%    10.68Mi ± 0%   -41.34% (p=0.000 n=10)
RandomPfx4Size/1_000_000    109.4Mi ± 0%    106.8Mi ± 0%    -2.37% (p=0.000 n=10)
RandomPfx6Size/100         157.08Ki ± 0%    16.22Ki ± 0%   -89.67% (p=0.000 n=10)
RandomPfx6Size/1_000       1277.7Ki ± 0%    128.3Ki ± 0%   -89.96% (p=0.000 n=10)
RandomPfx6Size/10_000      12.343Mi ± 0%    1.224Mi ± 0%   -90.08% (p=0.000 n=10)
RandomPfx6Size/100_000     115.13Mi ± 0%    12.21Mi ± 0%   -89.39% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1074.7Mi ± 0%    122.1Mi ± 0%   -88.64% (p=0.000 n=10)
RandomPfxSize/100           59.52Ki ± 0%    14.94Ki ± 0%   -74.90% (p=0.000 n=10)
RandomPfxSize/1_000         503.9Ki ± 0%    115.8Ki ± 0%   -77.03% (p=0.000 n=10)
RandomPfxSize/10_000        4.342Mi ± 0%    1.102Mi ± 0%   -74.62% (p=0.000 n=10)
RandomPfxSize/100_000       38.80Mi ± 0%    10.99Mi ± 0%   -71.67% (p=0.000 n=10)
RandomPfxSize/1_000_000     308.7Mi ± 0%    109.9Mi ± 0%   -64.41% (p=0.000 n=10)
geomean                     3.616Mi         1.193Mi        -67.01%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm  │            lpmtrie/size.bm             │
                         │     bytes     │     bytes      vs base                 │
Tier1PfxSize/100            42.79Ki ± 0%    24.22Ki ± 0%   -43.40% (p=0.000 n=10)
Tier1PfxSize/1_000          288.7Ki ± 0%    202.3Ki ± 0%   -29.91% (p=0.000 n=10)
Tier1PfxSize/10_000         1.820Mi ± 0%    1.942Mi ± 0%    +6.71% (p=0.000 n=10)
Tier1PfxSize/100_000        7.537Mi ± 0%   19.330Mi ± 0%  +156.46% (p=0.000 n=10)
Tier1PfxSize/1_000_000      33.71Mi ± 0%   189.92Mi ± 0%  +463.37% (p=0.000 n=10)
RandomPfx4Size/100          44.52Ki ± 0%    23.20Ki ± 0%   -47.89% (p=0.000 n=10)
RandomPfx4Size/1_000        302.2Ki ± 0%    198.0Ki ± 0%   -34.46% (p=0.000 n=10)
RandomPfx4Size/10_000       2.415Mi ± 0%    1.894Mi ± 0%   -21.56% (p=0.000 n=10)
RandomPfx4Size/100_000      18.21Mi ± 0%    18.25Mi ± 0%    +0.21% (p=0.001 n=10)
RandomPfx4Size/1_000_000    109.4Mi ± 0%    163.6Mi ± 0%   +49.54% (p=0.000 n=10)
RandomPfx6Size/100         157.08Ki ± 0%    25.53Ki ± 0%   -83.75% (p=0.000 n=10)
RandomPfx6Size/1_000       1277.7Ki ± 0%    221.9Ki ± 0%   -82.63% (p=0.000 n=10)
RandomPfx6Size/10_000      12.343Mi ± 0%    2.132Mi ± 0%   -82.73% (p=0.000 n=10)
RandomPfx6Size/100_000     115.13Mi ± 0%    20.92Mi ± 0%   -81.83% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1074.7Mi ± 0%    203.9Mi ± 0%   -81.02% (p=0.000 n=10)
RandomPfxSize/100           59.52Ki ± 0%    23.56Ki ± 0%   -60.41% (p=0.000 n=10)
RandomPfxSize/1_000         503.9Ki ± 0%    203.1Ki ± 0%   -59.70% (p=0.000 n=10)
RandomPfxSize/10_000        4.342Mi ± 0%    1.943Mi ± 0%   -55.26% (p=0.000 n=10)
RandomPfxSize/100_000       38.80Mi ± 0%    18.78Mi ± 0%   -51.59% (p=0.000 n=10)
RandomPfxSize/1_000_000     308.7Mi ± 0%    171.9Mi ± 0%   -44.30% (p=0.000 n=10)
geomean                     3.616Mi         2.010Mi        -44.41%
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
Insert/Insert_into_100            59.14n ± 0%   109.25n ± 3%   +84.75% (p=0.000 n=10)
Insert/Insert_into_1_000          59.15n ± 0%   129.90n ± 4%  +119.61% (p=0.000 n=10)
Insert/Insert_into_10_000         59.13n ± 0%   133.70n ± 1%  +126.13% (p=0.000 n=10)
Insert/Insert_into_100_000        59.21n ± 0%   124.15n ± 1%  +109.70% (p=0.000 n=10)
Insert/Insert_into_1_000_000      60.01n ± 0%   126.05n ± 0%  +110.05% (p=0.000 n=10)
Delete/Delete_from_100            23.66n ± 0%    29.66n ± 2%   +25.36% (p=0.000 n=10)
Delete/Delete_from_1_000          49.06n ± 0%    59.63n ± 2%   +21.55% (p=0.000 n=10)
Delete/Delete_from_10_000         49.06n ± 0%    58.27n ± 1%   +18.76% (p=0.000 n=10)
Delete/Delete_from_100_000        49.03n ± 0%    58.03n ± 1%   +18.34% (p=0.000 n=10)
Delete/Delete_from_1_000_000      49.82n ± 0%    57.94n ± 0%   +16.30% (p=0.000 n=10)
geomean                           50.23n         79.66n        +58.59%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │            cidrtree/update.bm            │
                             │     sec/op     │    sec/op      vs base                   │
Insert/Insert_into_100            59.14n ± 0%    743.05n ± 3%   +1156.53% (p=0.000 n=10)
Insert/Insert_into_1_000          59.15n ± 0%   1336.50n ± 0%   +2159.51% (p=0.000 n=10)
Insert/Insert_into_10_000         59.13n ± 0%   1759.50n ± 0%   +2875.90% (p=0.000 n=10)
Insert/Insert_into_100_000        59.21n ± 0%   2552.50n ± 1%   +4211.29% (p=0.000 n=10)
Insert/Insert_into_1_000_000      60.01n ± 0%   2749.00n ± 0%   +4480.90% (p=0.000 n=10)
Delete/Delete_from_100            23.66n ± 0%   1312.50n ± 1%   +5447.34% (p=0.000 n=10)
Delete/Delete_from_1_000          49.06n ± 0%   2354.00n ± 2%   +4698.21% (p=0.000 n=10)
Delete/Delete_from_10_000         49.06n ± 0%   2081.00n ± 1%   +4141.31% (p=0.000 n=10)
Delete/Delete_from_100_000        49.03n ± 0%   5773.50n ± 2%  +11674.24% (p=0.000 n=10)
Delete/Delete_from_1_000_000      49.82n ± 0%   4469.50n ± 1%   +8871.30% (p=0.000 n=10)
geomean                           50.23n          2.142µ        +4164.63%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │          critbitgo/update.bm          │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100            59.14n ± 0%   109.65n ± 1%   +85.42% (p=0.000 n=10)
Insert/Insert_into_1_000          59.15n ± 0%   124.70n ± 4%  +110.82% (p=0.000 n=10)
Insert/Insert_into_10_000         59.13n ± 0%   121.35n ± 1%  +105.24% (p=0.000 n=10)
Insert/Insert_into_100_000        59.21n ± 0%   147.90n ± 1%  +149.81% (p=0.000 n=10)
Insert/Insert_into_1_000_000      60.01n ± 0%   149.30n ± 0%  +148.79% (p=0.000 n=10)
Delete/Delete_from_100            23.66n ± 0%   105.30n ± 3%  +345.05% (p=0.000 n=10)
Delete/Delete_from_1_000          49.06n ± 0%   110.15n ± 3%  +124.52% (p=0.000 n=10)
Delete/Delete_from_10_000         49.06n ± 0%   120.60n ± 5%  +145.80% (p=0.000 n=10)
Delete/Delete_from_100_000        49.03n ± 0%   148.20n ± 6%  +202.23% (p=0.000 n=10)
Delete/Delete_from_1_000_000      49.82n ± 0%   148.90n ± 2%  +198.88% (p=0.000 n=10)
geomean                           50.23n         127.5n       +153.74%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │           lpmtrie/update.bm            │
                             │     sec/op     │    sec/op      vs base                 │
Insert/Insert_into_100            59.14n ± 0%   305.75n ±  2%  +417.04% (p=0.000 n=10)
Insert/Insert_into_1_000          59.15n ± 0%   347.75n ±  3%  +487.91% (p=0.000 n=10)
Insert/Insert_into_10_000         59.13n ± 0%   368.30n ± 12%  +522.92% (p=0.000 n=10)
Insert/Insert_into_100_000        59.21n ± 0%   479.45n ±  2%  +709.81% (p=0.000 n=10)
Insert/Insert_into_1_000_000      60.01n ± 0%   490.00n ± 21%  +716.53% (p=0.000 n=10)
Delete/Delete_from_100            23.66n ± 0%    74.22n ±  1%  +213.67% (p=0.000 n=10)
Delete/Delete_from_1_000          49.06n ± 0%   123.95n ±  1%  +152.65% (p=0.000 n=10)
Delete/Delete_from_10_000         49.06n ± 0%   140.50n ±  1%  +186.35% (p=0.000 n=10)
Delete/Delete_from_100_000        49.03n ± 0%   271.40n ±  1%  +453.48% (p=0.000 n=10)
Delete/Delete_from_1_000_000      49.82n ± 0%   272.90n ±  1%  +447.77% (p=0.000 n=10)
geomean                           50.23n         248.0n        +393.77%
```

## lookup (longest-prefix-match)

In the lookup, `art` is the champion, closely followed by `bart`. 

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │            art/lookup.bm             │
                                    │     sec/op     │    sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4             73.20n ± 21%   48.33n ± 19%  -33.98% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             63.67n ± 48%   48.61n ± 10%  -23.66% (p=0.002 n=10)
LpmTier1Pfxs/RandomMissIP4              87.50n ± 43%   32.00n ±  0%  -63.43% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              19.43n ±  0%   31.91n ±  0%  +64.16% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     90.59n ± 21%   56.37n ± 12%  -37.77% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     73.84n ± 14%   50.19n ± 17%  -32.03% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4     121.60n ± 19%   31.93n ±  3%  -73.74% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      83.64n ± 43%   31.86n ±  3%  -61.90% (p=0.000 n=10)
geomean                                 69.42n         40.26n        -42.00%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │             cidrtree/lookup.bm             │
                                    │     sec/op     │      sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4             73.20n ± 21%   1033.00n ±   60%  +1311.20% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             63.67n ± 48%   1235.50n ±   20%  +1840.47% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              87.50n ± 43%   1342.50n ±   98%  +1434.20% (p=0.023 n=10)
LpmTier1Pfxs/RandomMissIP6              19.43n ±  0%     54.49n ± 1264%   +180.37% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     90.59n ± 21%   1328.50n ±   48%  +1366.50% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     73.84n ± 14%   1767.50n ±   30%  +2293.69% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      121.6n ± 19%    1222.5n ±   17%   +905.35% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      83.64n ± 43%   1490.50n ±   35%  +1681.94% (p=0.000 n=10)
geomean                                 69.42n           891.6n          +1184.42%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │           critbitgo/lookup.bm           │
                                    │     sec/op     │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             73.20n ± 21%   365.25n ± 13%   +398.98% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             63.67n ± 48%   449.30n ± 15%   +605.67% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              87.50n ± 43%   682.65n ± 49%   +680.13% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              19.43n ±  0%   638.60n ± 17%  +3185.82% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     90.59n ± 21%   325.30n ± 22%   +259.09% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     73.84n ± 14%   386.85n ± 14%   +423.90% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      121.6n ± 19%    541.5n ± 21%   +345.31% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      83.64n ± 43%   461.20n ± 17%   +451.38% (p=0.000 n=10)
geomean                                 69.42n          466.6n         +572.21%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │           lpmtrie/lookup.bm            │
                                    │     sec/op     │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4             73.20n ± 21%   240.60n ± 16%  +228.69% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             63.67n ± 48%   238.80n ± 56%  +275.06% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              87.50n ± 43%   233.60n ± 68%         ~ (p=0.089 n=10)
LpmTier1Pfxs/RandomMissIP6              19.43n ±  0%    94.30n ± 18%  +385.23% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     90.59n ± 21%   277.70n ±  7%  +206.55% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     73.84n ± 14%   205.15n ±  7%  +177.83% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      121.6n ± 19%    282.3n ±  5%  +132.15% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      83.64n ± 43%   233.90n ± 17%  +179.63% (p=0.000 n=10)
geomean                                 69.42n          216.1n        +211.34%
```
