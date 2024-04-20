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
                         │ bart/size.bm │                art/size.bm                 │
                         │    bytes     │     bytes       vs base                    │
Tier1PfxSize/100           42.80Ki ± 0%    802.94Ki ± 0%  +1775.82% (p=0.000 n=10)
Tier1PfxSize/1_000         288.7Ki ± 0%    7420.4Ki ± 0%  +2470.25% (p=0.000 n=10)
Tier1PfxSize/10_000        1.820Mi ± 0%    47.172Mi ± 0%  +2492.02% (p=0.000 n=10)
Tier1PfxSize/100_000       7.537Mi ± 0%   160.300Mi ± 0%  +2026.77% (p=0.000 n=10)
Tier1PfxSize/1_000_000     33.71Mi ± 0%    378.23Mi ± 0%  +1021.98% (p=0.000 n=10)
RandomPfx4Size/100         41.47Ki ± 0%    725.72Ki ± 0%  +1650.04% (p=0.000 n=10)
RandomPfx4Size/1_000       297.5Ki ± 0%    7363.0Ki ± 0%  +2374.58% (p=0.000 n=10)
RandomPfx4Size/10_000      2.423Mi ± 0%    57.979Mi ± 0%  +2293.13% (p=0.000 n=10)
RandomPfx4Size/100_000     18.19Mi ± 0%    523.92Mi ± 0%  +2779.93% (p=0.000 n=10)
RandomPfx4Size/1_000_000   109.4Mi ± 0%
RandomPfx6Size/100         148.4Ki ± 0%     744.8Ki ± 0%   +401.88% (p=0.000 n=10)
RandomPfx6Size/1_000       1.307Mi ± 0%     7.527Mi ± 0%   +476.08% (p=0.000 n=10)
RandomPfx6Size/10_000      12.33Mi ± 0%     64.77Mi ± 0%   +425.14% (p=0.000 n=10)
RandomPfx6Size/100_000     115.5Mi ± 0%     748.9Mi ± 0%   +548.47% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1.049Gi ± 0%
RandomPfxSize/100          55.88Ki ± 0%    674.72Ki ± 0%  +1107.55% (p=0.000 n=10)
RandomPfxSize/1_000        518.2Ki ± 0%    7541.5Ki ± 0%  +1355.23% (p=0.000 n=10)
RandomPfxSize/10_000       4.432Mi ± 0%    59.442Mi ± 0%  +1241.25% (p=0.000 n=10)
RandomPfxSize/100_000      38.92Mi ± 0%    555.76Mi ± 0%  +1328.05% (p=0.000 n=10)
RandomPfxSize/1_000_000    308.7Mi ± 0%
geomean                    3.597Mi          22.63Mi       +1297.27%                ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

                         │  bart/size.bm  │           cidrtree/size.bm            │
                         │     bytes      │    bytes      vs base                 │
Tier1PfxSize/100             42.80Ki ± 0%   12.35Ki ± 0%   -71.14% (p=0.000 n=10)
Tier1PfxSize/1_000          288.70Ki ± 0%   81.51Ki ± 0%   -71.77% (p=0.000 n=10)
Tier1PfxSize/10_000         1863.6Ki ± 0%   784.6Ki ± 0%   -57.90% (p=0.000 n=10)
Tier1PfxSize/100_000         7.537Mi ± 0%   7.633Mi ± 0%    +1.27% (p=0.000 n=10)
Tier1PfxSize/1_000_000       33.71Mi ± 0%   76.30Mi ± 0%  +126.33% (p=0.000 n=10)
RandomPfx4Size/100           41.47Ki ± 0%   11.58Ki ± 0%   -72.08% (p=0.000 n=10)
RandomPfx4Size/1_000        297.55Ki ± 0%   81.51Ki ± 0%   -72.61% (p=0.000 n=10)
RandomPfx4Size/10_000       2480.9Ki ± 0%   784.6Ki ± 0%   -68.37% (p=0.000 n=10)
RandomPfx4Size/100_000      18.192Mi ± 0%   7.633Mi ± 0%   -58.04% (p=0.000 n=10)
RandomPfx4Size/1_000_000    109.41Mi ± 0%   76.30Mi ± 0%   -30.27% (p=0.000 n=10)
RandomPfx6Size/100          148.41Ki ± 0%   11.58Ki ± 0%   -92.20% (p=0.000 n=10)
RandomPfx6Size/1_000       1337.88Ki ± 0%   81.51Ki ± 0%   -93.91% (p=0.000 n=10)
RandomPfx6Size/10_000      12628.9Ki ± 0%   785.0Ki ± 0%   -93.78% (p=0.000 n=10)
RandomPfx6Size/100_000     115.492Mi ± 0%   7.633Mi ± 0%   -93.39% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1074.68Mi ± 0%   76.30Mi ± 0%   -92.90% (p=0.000 n=10)
RandomPfxSize/100            55.88Ki ± 0%   11.58Ki ± 0%   -79.28% (p=0.000 n=10)
RandomPfxSize/1_000         518.23Ki ± 0%   81.51Ki ± 0%   -84.27% (p=0.000 n=10)
RandomPfxSize/10_000        4538.2Ki ± 0%   784.6Ki ± 0%   -82.71% (p=0.000 n=10)
RandomPfxSize/100_000       38.917Mi ± 0%   7.633Mi ± 0%   -80.39% (p=0.000 n=10)
RandomPfxSize/1_000_000     308.68Mi ± 0%   76.30Mi ± 0%   -75.28% (p=0.000 n=10)
geomean                      3.597Mi        856.0Ki        -76.76%

                         │ bart/size.bm  │           critbitgo/size.bm            │
                         │     bytes     │     bytes      vs base                 │
Tier1PfxSize/100            42.80Ki ± 0%    15.52Ki ± 0%   -63.75% (p=0.000 n=10)
Tier1PfxSize/1_000          288.7Ki ± 0%    115.2Ki ± 0%   -60.09% (p=0.000 n=10)
Tier1PfxSize/10_000         1.820Mi ± 0%    1.094Mi ± 0%   -39.91% (p=0.000 n=10)
Tier1PfxSize/100_000        7.537Mi ± 0%   10.913Mi ± 0%   +44.78% (p=0.000 n=10)
Tier1PfxSize/1_000_000      33.71Mi ± 0%   109.12Mi ± 0%  +223.68% (p=0.000 n=10)
RandomPfx4Size/100          41.47Ki ± 0%    14.67Ki ± 0%   -64.62% (p=0.000 n=10)
RandomPfx4Size/1_000        297.5Ki ± 0%    112.7Ki ± 0%   -62.11% (p=0.000 n=10)
RandomPfx4Size/10_000       2.423Mi ± 0%    1.071Mi ± 0%   -55.78% (p=0.000 n=10)
RandomPfx4Size/100_000      18.19Mi ± 0%    10.68Mi ± 0%   -41.27% (p=0.000 n=10)
RandomPfx4Size/1_000_000    109.4Mi ± 0%    106.8Mi ± 0%    -2.37% (p=0.000 n=10)
RandomPfx6Size/100         148.41Ki ± 0%    16.22Ki ± 0%   -89.07% (p=0.000 n=10)
RandomPfx6Size/1_000       1337.9Ki ± 0%    128.3Ki ± 0%   -90.41% (p=0.000 n=10)
RandomPfx6Size/10_000      12.333Mi ± 0%    1.224Mi ± 0%   -90.07% (p=0.000 n=10)
RandomPfx6Size/100_000     115.49Mi ± 0%    12.21Mi ± 0%   -89.43% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1074.7Mi ± 0%    122.1Mi ± 0%   -88.64% (p=0.000 n=10)
RandomPfxSize/100           55.88Ki ± 0%    14.94Ki ± 0%   -73.27% (p=0.000 n=10)
RandomPfxSize/1_000         518.2Ki ± 0%    115.8Ki ± 0%   -77.66% (p=0.000 n=10)
RandomPfxSize/10_000        4.432Mi ± 0%    1.102Mi ± 0%   -75.13% (p=0.000 n=10)
RandomPfxSize/100_000       38.92Mi ± 0%    10.99Mi ± 0%   -71.76% (p=0.000 n=10)
RandomPfxSize/1_000_000     308.7Mi ± 0%    109.9Mi ± 0%   -64.41% (p=0.000 n=10)
geomean                     3.597Mi         1.193Mi        -66.84%

                         │ bart/size.bm  │            lpmtrie/size.bm             │
                         │     bytes     │     bytes      vs base                 │
Tier1PfxSize/100            42.80Ki ± 0%    24.22Ki ± 0%   -43.42% (p=0.000 n=10)
Tier1PfxSize/1_000          288.7Ki ± 0%    202.3Ki ± 0%   -29.92% (p=0.000 n=10)
Tier1PfxSize/10_000         1.820Mi ± 0%    1.942Mi ± 0%    +6.71% (p=0.000 n=10)
Tier1PfxSize/100_000        7.537Mi ± 0%   19.330Mi ± 0%  +156.46% (p=0.000 n=10)
Tier1PfxSize/1_000_000      33.71Mi ± 0%   189.92Mi ± 0%  +463.37% (p=0.000 n=10)
RandomPfx4Size/100          41.47Ki ± 0%    23.20Ki ± 0%   -44.05% (p=0.000 n=10)
RandomPfx4Size/1_000        297.5Ki ± 0%    198.0Ki ± 0%   -33.44% (p=0.000 n=10)
RandomPfx4Size/10_000       2.423Mi ± 0%    1.894Mi ± 0%   -21.82% (p=0.000 n=10)
RandomPfx4Size/100_000      18.19Mi ± 0%    18.25Mi ± 0%    +0.34% (p=0.001 n=10)
RandomPfx4Size/1_000_000    109.4Mi ± 0%    163.6Mi ± 0%   +49.54% (p=0.000 n=10)
RandomPfx6Size/100         148.41Ki ± 0%    25.53Ki ± 0%   -82.80% (p=0.000 n=10)
RandomPfx6Size/1_000       1337.9Ki ± 0%    221.9Ki ± 0%   -83.41% (p=0.000 n=10)
RandomPfx6Size/10_000      12.333Mi ± 0%    2.132Mi ± 0%   -82.72% (p=0.000 n=10)
RandomPfx6Size/100_000     115.49Mi ± 0%    20.92Mi ± 0%   -81.88% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1074.7Mi ± 0%    203.9Mi ± 0%   -81.02% (p=0.000 n=10)
RandomPfxSize/100           55.88Ki ± 0%    23.56Ki ± 0%   -57.83% (p=0.000 n=10)
RandomPfxSize/1_000         518.2Ki ± 0%    203.1Ki ± 0%   -60.82% (p=0.000 n=10)
RandomPfxSize/10_000        4.432Mi ± 0%    1.943Mi ± 0%   -56.16% (p=0.000 n=10)
RandomPfxSize/100_000       38.92Mi ± 0%    18.78Mi ± 0%   -51.74% (p=0.000 n=10)
RandomPfxSize/1_000_000     308.7Mi ± 0%    171.9Mi ± 0%   -44.30% (p=0.000 n=10)
geomean                     3.597Mi         2.010Mi        -44.12%
```

## update, insert/delete

When it comes to updates, `art` and `bart` are far superior
to the other algorithms, only `critbitgo` comes close to playing in the same league .

```
                             │ bart/update.bm │            art/update.bm             │
                             │     sec/op     │    sec/op     vs base                │
Insert/Insert_into_100            69.75n ± 3%   109.25n ± 3%  +56.62% (p=0.000 n=10)
Insert/Insert_into_1_000          75.12n ± 1%   129.90n ± 4%  +72.93% (p=0.000 n=10)
Insert/Insert_into_10_000         69.44n ± 1%   133.70n ± 1%  +92.55% (p=0.000 n=10)
Insert/Insert_into_100_000        69.49n ± 1%   124.15n ± 1%  +78.65% (p=0.000 n=10)
Insert/Insert_into_1_000_000      72.07n ± 2%   126.05n ± 0%  +74.90% (p=0.000 n=10)
Delete/Delete_from_100            32.91n ± 4%    29.66n ± 2%   -9.86% (p=0.000 n=10)
Delete/Delete_from_1_000          57.86n ± 4%    59.63n ± 2%   +3.07% (p=0.035 n=10)
Delete/Delete_from_10_000         56.66n ± 0%    58.27n ± 1%   +2.84% (p=0.000 n=10)
Delete/Delete_from_100_000        57.23n ± 0%    58.03n ± 1%   +1.40% (p=0.000 n=10)
Delete/Delete_from_1_000_000      56.88n ± 0%    57.94n ± 0%   +1.85% (p=0.000 n=10)
geomean                           60.34n         79.66n       +32.02%

                             │ bart/update.bm │           cidrtree/update.bm            │
                             │     sec/op     │    sec/op      vs base                  │
Insert/Insert_into_100            69.75n ± 3%    743.05n ± 3%   +965.23% (p=0.000 n=10)
Insert/Insert_into_1_000          75.12n ± 1%   1336.50n ± 0%  +1679.27% (p=0.000 n=10)
Insert/Insert_into_10_000         69.44n ± 1%   1759.50n ± 0%  +2434.02% (p=0.000 n=10)
Insert/Insert_into_100_000        69.49n ± 1%   2552.50n ± 1%  +3572.93% (p=0.000 n=10)
Insert/Insert_into_1_000_000      72.07n ± 2%   2749.00n ± 0%  +3714.35% (p=0.000 n=10)
Delete/Delete_from_100            32.91n ± 4%   1312.50n ± 1%  +3888.76% (p=0.000 n=10)
Delete/Delete_from_1_000          57.86n ± 4%   2354.00n ± 2%  +3968.79% (p=0.000 n=10)
Delete/Delete_from_10_000         56.66n ± 0%   2081.00n ± 1%  +3572.79% (p=0.000 n=10)
Delete/Delete_from_100_000        57.23n ± 0%   5773.50n ± 2%  +9988.24% (p=0.000 n=10)
Delete/Delete_from_1_000_000      56.88n ± 0%   4469.50n ± 1%  +7757.08% (p=0.000 n=10)
geomean                           60.34n          2.142µ       +3450.02%

                             │ bart/update.bm │          critbitgo/update.bm          │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100            69.75n ± 3%   109.65n ± 1%   +57.19% (p=0.000 n=10)
Insert/Insert_into_1_000          75.12n ± 1%   124.70n ± 4%   +66.01% (p=0.000 n=10)
Insert/Insert_into_10_000         69.44n ± 1%   121.35n ± 1%   +74.77% (p=0.000 n=10)
Insert/Insert_into_100_000        69.49n ± 1%   147.90n ± 1%  +112.82% (p=0.000 n=10)
Insert/Insert_into_1_000_000      72.07n ± 2%   149.30n ± 0%  +107.16% (p=0.000 n=10)
Delete/Delete_from_100            32.91n ± 4%   105.30n ± 3%  +220.01% (p=0.000 n=10)
Delete/Delete_from_1_000          57.86n ± 4%   110.15n ± 3%   +90.39% (p=0.000 n=10)
Delete/Delete_from_10_000         56.66n ± 0%   120.60n ± 5%  +112.85% (p=0.000 n=10)
Delete/Delete_from_100_000        57.23n ± 0%   148.20n ± 6%  +158.96% (p=0.000 n=10)
Delete/Delete_from_1_000_000      56.88n ± 0%   148.90n ± 2%  +161.76% (p=0.000 n=10)
geomean                           60.34n         127.5n       +111.22%

                             │ bart/update.bm │           lpmtrie/update.bm            │
                             │     sec/op     │    sec/op      vs base                 │
Insert/Insert_into_100            69.75n ± 3%   305.75n ±  2%  +338.32% (p=0.000 n=10)
Insert/Insert_into_1_000          75.12n ± 1%   347.75n ±  3%  +362.96% (p=0.000 n=10)
Insert/Insert_into_10_000         69.44n ± 1%   368.30n ± 12%  +430.42% (p=0.000 n=10)
Insert/Insert_into_100_000        69.49n ± 1%   479.45n ±  2%  +589.91% (p=0.000 n=10)
Insert/Insert_into_1_000_000      72.07n ± 2%   490.00n ± 21%  +579.89% (p=0.000 n=10)
Delete/Delete_from_100            32.91n ± 4%    74.22n ±  1%  +125.54% (p=0.000 n=10)
Delete/Delete_from_1_000          57.86n ± 4%   123.95n ±  1%  +114.24% (p=0.000 n=10)
Delete/Delete_from_10_000         56.66n ± 0%   140.50n ±  1%  +147.97% (p=0.000 n=10)
Delete/Delete_from_100_000        57.23n ± 0%   271.40n ±  1%  +374.23% (p=0.000 n=10)
Delete/Delete_from_1_000_000      56.88n ± 0%   272.90n ±  1%  +379.74% (p=0.000 n=10)
geomean                           60.34n         248.0n        +311.03%
```

## lookup (longest-prefix-match)

In the lookup, `art` is the champion, closely followed by `bart`. 

```
                                    │ bart/lookup.bm │            art/lookup.bm             │
                                    │     sec/op     │    sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4             66.66n ± 32%   48.33n ± 19%  -27.51% (p=0.043 n=10)
LpmTier1Pfxs/RandomMatchIP6             62.77n ± 23%   48.61n ± 10%  -22.56% (p=0.002 n=10)
LpmTier1Pfxs/RandomMissIP4              84.81n ± 48%   32.00n ±  0%  -62.27% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              19.80n ±  0%   31.91n ±  0%  +61.10% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     73.29n ± 51%   56.37n ± 12%  -23.08% (p=0.003 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     78.12n ± 26%   50.19n ± 17%  -35.76% (p=0.001 n=10)
LpmRandomPfxs100_000/RandomMissIP4      89.42n ± 34%   31.93n ±  3%  -64.29% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      94.82n ± 13%   31.86n ±  3%  -66.39% (p=0.000 n=10)
geomean                                 65.56n         40.26n        -38.59%

                                    │ bart/lookup.bm │             cidrtree/lookup.bm             │
                                    │     sec/op     │      sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4             66.66n ± 32%   1033.00n ±   60%  +1449.54% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             62.77n ± 23%   1235.50n ±   20%  +1868.45% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              84.81n ± 48%   1342.50n ±   98%  +1482.86% (p=0.023 n=10)
LpmTier1Pfxs/RandomMissIP6              19.80n ±  0%     54.49n ± 1264%   +175.13% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     73.29n ± 51%   1328.50n ±   48%  +1712.79% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     78.12n ± 26%   1767.50n ±   30%  +2162.40% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      89.42n ± 34%   1222.50n ±   17%  +1267.22% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      94.82n ± 13%   1490.50n ±   35%  +1472.01% (p=0.000 n=10)
geomean                                 65.56n           891.6n          +1260.08%

                                    │ bart/lookup.bm │           critbitgo/lookup.bm           │
                                    │     sec/op     │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             66.66n ± 32%   365.25n ± 13%   +447.89% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             62.77n ± 23%   449.30n ± 15%   +615.84% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              84.81n ± 48%   682.65n ± 49%   +704.87% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              19.80n ±  0%   638.60n ± 17%  +3124.44% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     73.29n ± 51%   325.30n ± 22%   +343.88% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     78.12n ± 26%   386.85n ± 14%   +395.17% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      89.42n ± 34%   541.50n ± 21%   +505.60% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      94.82n ± 13%   461.20n ± 17%   +386.42% (p=0.000 n=10)
geomean                                 65.56n          466.6n         +611.81%

                                    │ bart/lookup.bm │           lpmtrie/lookup.bm            │
                                    │     sec/op     │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4             66.66n ± 32%   240.60n ± 16%  +260.91% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             62.77n ± 23%   238.80n ± 56%  +280.47% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              84.81n ± 48%   233.60n ± 68%         ~ (p=0.052 n=10)
LpmTier1Pfxs/RandomMissIP6              19.80n ±  0%    94.30n ± 18%  +376.17% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     73.29n ± 51%   277.70n ±  7%  +278.93% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     78.12n ± 26%   205.15n ±  7%  +162.59% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      89.42n ± 34%   282.30n ±  5%  +215.72% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      94.82n ± 13%   233.90n ± 17%  +146.69% (p=0.000 n=10)
geomean                                 65.56n          216.1n        +229.68%
```
