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
Tier1PfxSize/100           45.88Ki ± 0%    802.94Ki ± 0%  +1649.97% (p=0.000 n=10)
Tier1PfxSize/1_000         309.9Ki ± 0%    7420.4Ki ± 0%  +2294.40% (p=0.000 n=10)
Tier1PfxSize/10_000        1.946Mi ± 0%    47.172Mi ± 0%  +2323.50% (p=0.000 n=10)
Tier1PfxSize/100_000       7.965Mi ± 0%   160.300Mi ± 0%  +1912.53% (p=0.000 n=10)
Tier1PfxSize/1_000_000     34.68Mi ± 0%    378.23Mi ± 0%   +990.77% (p=0.000 n=10)
RandomPfx4Size/100         49.25Ki ± 0%    725.72Ki ± 0%  +1373.54% (p=0.000 n=10)
RandomPfx4Size/1_000       320.6Ki ± 0%    7363.0Ki ± 0%  +2196.80% (p=0.000 n=10)
RandomPfx4Size/10_000      2.607Mi ± 0%    57.979Mi ± 0%  +2123.87% (p=0.000 n=10)
RandomPfx4Size/100_000     19.66Mi ± 0%    523.92Mi ± 0%  +2564.94% (p=0.000 n=10)
RandomPfx4Size/1_000_000   117.4Mi ± 0%
RandomPfx6Size/100         161.6Ki ± 0%     744.8Ki ± 0%   +360.88% (p=0.000 n=10)
RandomPfx6Size/1_000       1.420Mi ± 0%     7.527Mi ± 0%   +430.09% (p=0.000 n=10)
RandomPfx6Size/10_000      13.16Mi ± 0%     64.77Mi ± 0%   +391.98% (p=0.000 n=10)
RandomPfx6Size/100_000     125.3Mi ± 0%     748.9Mi ± 0%   +497.87% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1.144Gi ± 0%
RandomPfxSize/100          62.06Ki ± 0%    674.72Ki ± 0%   +987.16% (p=0.000 n=10)
RandomPfxSize/1_000        553.7Ki ± 0%    7541.5Ki ± 0%  +1262.09% (p=0.000 n=10)
RandomPfxSize/10_000       4.761Mi ± 0%    59.442Mi ± 0%  +1148.56% (p=0.000 n=10)
RandomPfxSize/100_000      42.18Mi ± 0%    555.76Mi ± 0%  +1217.69% (p=0.000 n=10)
RandomPfxSize/1_000_000    334.8Mi ± 0%
geomean                    3.890Mi          22.63Mi       +1192.38%                ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

                         │  bart/size.bm  │           cidrtree/size.bm            │
                         │     bytes      │    bytes      vs base                 │
Tier1PfxSize/100             45.88Ki ± 0%   12.35Ki ± 0%   -73.08% (p=0.000 n=10)
Tier1PfxSize/1_000          309.91Ki ± 0%   81.51Ki ± 0%   -73.70% (p=0.000 n=10)
Tier1PfxSize/10_000         1993.1Ki ± 0%   784.6Ki ± 0%   -60.63% (p=0.000 n=10)
Tier1PfxSize/100_000         7.965Mi ± 0%   7.633Mi ± 0%    -4.17% (p=0.000 n=10)
Tier1PfxSize/1_000_000       34.68Mi ± 0%   76.30Mi ± 0%  +120.03% (p=0.000 n=10)
RandomPfx4Size/100           49.25Ki ± 0%   11.58Ki ± 0%   -76.49% (p=0.000 n=10)
RandomPfx4Size/1_000        320.58Ki ± 0%   81.51Ki ± 0%   -74.57% (p=0.000 n=10)
RandomPfx4Size/10_000       2669.7Ki ± 0%   784.6Ki ± 0%   -70.61% (p=0.000 n=10)
RandomPfx4Size/100_000      19.660Mi ± 0%   7.633Mi ± 0%   -61.17% (p=0.000 n=10)
RandomPfx4Size/1_000_000    117.37Mi ± 0%   76.30Mi ± 0%   -34.99% (p=0.000 n=10)
RandomPfx6Size/100          161.61Ki ± 0%   11.58Ki ± 0%   -92.84% (p=0.000 n=10)
RandomPfx6Size/1_000       1453.95Ki ± 0%   81.51Ki ± 0%   -94.39% (p=0.000 n=10)
RandomPfx6Size/10_000      13480.2Ki ± 0%   785.0Ki ± 0%   -94.18% (p=0.000 n=10)
RandomPfx6Size/100_000     125.267Mi ± 0%   7.633Mi ± 0%   -93.91% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1170.96Mi ± 0%   76.30Mi ± 0%   -93.48% (p=0.000 n=10)
RandomPfxSize/100            62.06Ki ± 0%   11.58Ki ± 0%   -81.34% (p=0.000 n=10)
RandomPfxSize/1_000         553.67Ki ± 0%   81.51Ki ± 0%   -85.28% (p=0.000 n=10)
RandomPfxSize/10_000        4875.1Ki ± 0%   784.6Ki ± 0%   -83.91% (p=0.000 n=10)
RandomPfxSize/100_000       42.176Mi ± 0%   7.633Mi ± 0%   -81.90% (p=0.000 n=10)
RandomPfxSize/1_000_000     334.82Mi ± 0%   76.30Mi ± 0%   -77.21% (p=0.000 n=10)
geomean                      3.890Mi        856.0Ki        -78.51%

                         │ bart/size.bm  │           critbitgo/size.bm            │
                         │     bytes     │     bytes      vs base                 │
Tier1PfxSize/100            45.88Ki ± 0%    15.52Ki ± 0%   -66.18% (p=0.000 n=10)
Tier1PfxSize/1_000          309.9Ki ± 0%    115.2Ki ± 0%   -62.82% (p=0.000 n=10)
Tier1PfxSize/10_000         1.946Mi ± 0%    1.094Mi ± 0%   -43.82% (p=0.000 n=10)
Tier1PfxSize/100_000        7.965Mi ± 0%   10.913Mi ± 0%   +37.01% (p=0.000 n=10)
Tier1PfxSize/1_000_000      34.68Mi ± 0%   109.12Mi ± 0%  +214.68% (p=0.000 n=10)
RandomPfx4Size/100          49.25Ki ± 0%    14.67Ki ± 0%   -70.21% (p=0.000 n=10)
RandomPfx4Size/1_000        320.6Ki ± 0%    112.7Ki ± 0%   -64.84% (p=0.000 n=10)
RandomPfx4Size/10_000       2.607Mi ± 0%    1.071Mi ± 0%   -58.91% (p=0.000 n=10)
RandomPfx4Size/100_000      19.66Mi ± 0%    10.68Mi ± 0%   -45.65% (p=0.000 n=10)
RandomPfx4Size/1_000_000    117.4Mi ± 0%    106.8Mi ± 0%    -8.99% (p=0.000 n=10)
RandomPfx6Size/100         161.61Ki ± 0%    16.22Ki ± 0%   -89.96% (p=0.000 n=10)
RandomPfx6Size/1_000       1454.0Ki ± 0%    128.3Ki ± 0%   -91.17% (p=0.000 n=10)
RandomPfx6Size/10_000      13.164Mi ± 0%    1.224Mi ± 0%   -90.70% (p=0.000 n=10)
RandomPfx6Size/100_000     125.27Mi ± 0%    12.21Mi ± 0%   -90.25% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1171.0Mi ± 0%    122.1Mi ± 0%   -89.57% (p=0.000 n=10)
RandomPfxSize/100           62.06Ki ± 0%    14.94Ki ± 0%   -75.93% (p=0.000 n=10)
RandomPfxSize/1_000         553.7Ki ± 0%    115.8Ki ± 0%   -79.09% (p=0.000 n=10)
RandomPfxSize/10_000        4.761Mi ± 0%    1.102Mi ± 0%   -76.85% (p=0.000 n=10)
RandomPfxSize/100_000       42.18Mi ± 0%    10.99Mi ± 0%   -73.94% (p=0.000 n=10)
RandomPfxSize/1_000_000     334.8Mi ± 0%    109.9Mi ± 0%   -67.19% (p=0.000 n=10)
geomean                     3.890Mi         1.193Mi        -69.33%

                         │ bart/size.bm  │            lpmtrie/size.bm             │
                         │     bytes     │     bytes      vs base                 │
Tier1PfxSize/100            45.88Ki ± 0%    24.22Ki ± 0%   -47.22% (p=0.000 n=10)
Tier1PfxSize/1_000          309.9Ki ± 0%    202.3Ki ± 0%   -34.71% (p=0.000 n=10)
Tier1PfxSize/10_000         1.946Mi ± 0%    1.942Mi ± 0%    -0.22% (p=0.000 n=10)
Tier1PfxSize/100_000        7.965Mi ± 0%   19.330Mi ± 0%  +142.69% (p=0.000 n=10)
Tier1PfxSize/1_000_000      34.68Mi ± 0%   189.92Mi ± 0%  +447.70% (p=0.000 n=10)
RandomPfx4Size/100          49.25Ki ± 0%    23.20Ki ± 0%   -52.89% (p=0.000 n=10)
RandomPfx4Size/1_000        320.6Ki ± 0%    198.0Ki ± 0%   -38.22% (p=0.000 n=10)
RandomPfx4Size/10_000       2.607Mi ± 0%    1.894Mi ± 0%   -27.35% (p=0.000 n=10)
RandomPfx4Size/100_000      19.66Mi ± 0%    18.25Mi ± 0%    -7.15% (p=0.000 n=10)
RandomPfx4Size/1_000_000    117.4Mi ± 0%    163.6Mi ± 0%   +39.40% (p=0.000 n=10)
RandomPfx6Size/100         161.61Ki ± 0%    25.53Ki ± 0%   -84.20% (p=0.000 n=10)
RandomPfx6Size/1_000       1454.0Ki ± 0%    221.9Ki ± 0%   -84.74% (p=0.000 n=10)
RandomPfx6Size/10_000      13.164Mi ± 0%    2.132Mi ± 0%   -83.81% (p=0.000 n=10)
RandomPfx6Size/100_000     125.27Mi ± 0%    20.92Mi ± 0%   -83.30% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1171.0Mi ± 0%    203.9Mi ± 0%   -82.58% (p=0.000 n=10)
RandomPfxSize/100           62.06Ki ± 0%    23.56Ki ± 0%   -62.03% (p=0.000 n=10)
RandomPfxSize/1_000         553.7Ki ± 0%    203.1Ki ± 0%   -63.33% (p=0.000 n=10)
RandomPfxSize/10_000        4.761Mi ± 0%    1.943Mi ± 0%   -59.19% (p=0.000 n=10)
RandomPfxSize/100_000       42.18Mi ± 0%    18.78Mi ± 0%   -55.47% (p=0.000 n=10)
RandomPfxSize/1_000_000     334.8Mi ± 0%    171.9Mi ± 0%   -48.65% (p=0.000 n=10)
geomean                     3.890Mi         2.010Mi        -48.32%
```

## update, insert/delete

When it comes to updates, `art` and `bart` are far superior
to the other algorithms, only `critbitgo` comes close to playing in the same league .

```
                             │ bart/update.bm │            art/update.bm             │
                             │     sec/op     │    sec/op     vs base                │
Insert/Insert_into_100            69.14n ± 1%   109.25n ± 3%  +58.01% (p=0.000 n=10)
Insert/Insert_into_1_000          70.45n ± 1%   129.90n ± 4%  +84.37% (p=0.000 n=10)
Insert/Insert_into_10_000         75.67n ± 1%   133.70n ± 1%  +76.70% (p=0.000 n=10)
Insert/Insert_into_100_000        68.60n ± 1%   124.15n ± 1%  +80.98% (p=0.000 n=10)
Insert/Insert_into_1_000_000      70.47n ± 0%   126.05n ± 0%  +78.86% (p=0.000 n=10)
Delete/Delete_from_100            32.27n ± 1%    29.66n ± 2%   -8.09% (p=0.000 n=10)
Delete/Delete_from_1_000          55.96n ± 0%    59.63n ± 2%   +6.56% (p=0.000 n=10)
Delete/Delete_from_10_000         61.59n ± 0%    58.27n ± 1%   -5.38% (p=0.000 n=10)
Delete/Delete_from_100_000        58.80n ± 0%    58.03n ± 1%   -1.32% (p=0.001 n=10)
Delete/Delete_from_1_000_000      56.13n ± 0%    57.94n ± 0%   +3.23% (p=0.000 n=10)
geomean                           60.47n         79.66n       +31.73%

                             │ bart/update.bm │           cidrtree/update.bm            │
                             │     sec/op     │    sec/op      vs base                  │
Insert/Insert_into_100            69.14n ± 1%    743.05n ± 3%   +974.70% (p=0.000 n=10)
Insert/Insert_into_1_000          70.45n ± 1%   1336.50n ± 0%  +1796.96% (p=0.000 n=10)
Insert/Insert_into_10_000         75.67n ± 1%   1759.50n ± 0%  +2225.38% (p=0.000 n=10)
Insert/Insert_into_100_000        68.60n ± 1%   2552.50n ± 1%  +3620.85% (p=0.000 n=10)
Insert/Insert_into_1_000_000      70.47n ± 0%   2749.00n ± 0%  +3800.67% (p=0.000 n=10)
Delete/Delete_from_100            32.27n ± 1%   1312.50n ± 1%  +3967.25% (p=0.000 n=10)
Delete/Delete_from_1_000          55.96n ± 0%   2354.00n ± 2%  +4106.58% (p=0.000 n=10)
Delete/Delete_from_10_000         61.59n ± 0%   2081.00n ± 1%  +3279.07% (p=0.000 n=10)
Delete/Delete_from_100_000        58.80n ± 0%   5773.50n ± 2%  +9718.04% (p=0.000 n=10)
Delete/Delete_from_1_000_000      56.13n ± 0%   4469.50n ± 1%  +7863.47% (p=0.000 n=10)
geomean                           60.47n          2.142µ       +3442.24%

                             │ bart/update.bm │          critbitgo/update.bm          │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100            69.14n ± 1%   109.65n ± 1%   +58.59% (p=0.000 n=10)
Insert/Insert_into_1_000          70.45n ± 1%   124.70n ± 4%   +76.99% (p=0.000 n=10)
Insert/Insert_into_10_000         75.67n ± 1%   121.35n ± 1%   +60.38% (p=0.000 n=10)
Insert/Insert_into_100_000        68.60n ± 1%   147.90n ± 1%  +115.60% (p=0.000 n=10)
Insert/Insert_into_1_000_000      70.47n ± 0%   149.30n ± 0%  +111.85% (p=0.000 n=10)
Delete/Delete_from_100            32.27n ± 1%   105.30n ± 3%  +226.31% (p=0.000 n=10)
Delete/Delete_from_1_000          55.96n ± 0%   110.15n ± 3%   +96.84% (p=0.000 n=10)
Delete/Delete_from_10_000         61.59n ± 0%   120.60n ± 5%   +95.83% (p=0.000 n=10)
Delete/Delete_from_100_000        58.80n ± 0%   148.20n ± 6%  +152.02% (p=0.000 n=10)
Delete/Delete_from_1_000_000      56.13n ± 0%   148.90n ± 2%  +165.30% (p=0.000 n=10)
geomean                           60.47n         127.5n       +110.76%

                             │ bart/update.bm │           lpmtrie/update.bm            │
                             │     sec/op     │    sec/op      vs base                 │
Insert/Insert_into_100            69.14n ± 1%   305.75n ±  2%  +342.22% (p=0.000 n=10)
Insert/Insert_into_1_000          70.45n ± 1%   347.75n ±  3%  +393.58% (p=0.000 n=10)
Insert/Insert_into_10_000         75.67n ± 1%   368.30n ± 12%  +386.75% (p=0.000 n=10)
Insert/Insert_into_100_000        68.60n ± 1%   479.45n ±  2%  +598.91% (p=0.000 n=10)
Insert/Insert_into_1_000_000      70.47n ± 0%   490.00n ± 21%  +595.28% (p=0.000 n=10)
Delete/Delete_from_100            32.27n ± 1%    74.22n ±  1%  +129.98% (p=0.000 n=10)
Delete/Delete_from_1_000          55.96n ± 0%   123.95n ±  1%  +121.50% (p=0.000 n=10)
Delete/Delete_from_10_000         61.59n ± 0%   140.50n ±  1%  +128.14% (p=0.000 n=10)
Delete/Delete_from_100_000        58.80n ± 0%   271.40n ±  1%  +361.53% (p=0.000 n=10)
Delete/Delete_from_1_000_000      56.13n ± 0%   272.90n ±  1%  +386.24% (p=0.000 n=10)
geomean                           60.47n         248.0n        +310.13%
```

## lookup (longest-prefix-match)

In the lookup, `art` is the champion, closely followed by `bart`. 

```
                                    │ bart/lookup.bm │            art/lookup.bm             │
                                    │     sec/op     │    sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4             64.93n ± 36%   48.33n ± 19%  -25.57% (p=0.001 n=10)
LpmTier1Pfxs/RandomMatchIP6             68.18n ± 34%   48.61n ± 10%        ~ (p=0.143 n=10)
LpmTier1Pfxs/RandomMissIP4              73.73n ± 35%   32.00n ±  0%  -56.60% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              19.75n ± 65%   31.91n ±  0%  +61.50% (p=0.022 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     84.22n ± 41%   56.37n ± 12%  -33.07% (p=0.007 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     82.61n ± 14%   50.19n ± 17%  -39.25% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4     107.39n ± 22%   31.93n ±  3%  -70.27% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      82.83n ±  2%   31.86n ±  3%  -61.53% (p=0.000 n=10)
geomean                                 66.85n         40.26n        -39.78%

                                    │ bart/lookup.bm │             cidrtree/lookup.bm             │
                                    │     sec/op     │      sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4             64.93n ± 36%   1033.00n ±   60%  +1490.94% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             68.18n ± 34%   1235.50n ±   20%  +1712.25% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              73.73n ± 35%   1342.50n ±   98%  +1720.71% (p=0.023 n=10)
LpmTier1Pfxs/RandomMissIP6              19.75n ± 65%     54.49n ± 1264%   +175.83% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     84.22n ± 41%   1328.50n ±   48%  +1477.32% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     82.61n ± 14%   1767.50n ±   30%  +2039.44% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      107.4n ± 22%    1222.5n ±   17%  +1038.32% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      82.83n ±  2%   1490.50n ±   35%  +1699.58% (p=0.000 n=10)
geomean                                 66.85n           891.6n          +1233.68%

                                    │ bart/lookup.bm │           critbitgo/lookup.bm           │
                                    │     sec/op     │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             64.93n ± 36%   365.25n ± 13%   +462.53% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             68.18n ± 34%   449.30n ± 15%   +559.04% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              73.73n ± 35%   682.65n ± 49%   +825.82% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              19.75n ± 65%   638.60n ± 17%  +3132.60% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     84.22n ± 41%   325.30n ± 22%   +286.23% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     82.61n ± 14%   386.85n ± 14%   +368.26% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      107.4n ± 22%    541.5n ± 21%   +404.21% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      82.83n ±  2%   461.20n ± 17%   +456.84% (p=0.000 n=10)
geomean                                 66.85n          466.6n         +597.99%

                                    │ bart/lookup.bm │           lpmtrie/lookup.bm            │
                                    │     sec/op     │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4             64.93n ± 36%   240.60n ± 16%  +270.55% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             68.18n ± 34%   238.80n ± 56%  +250.28% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              73.73n ± 35%   233.60n ± 68%  +216.81% (p=0.007 n=10)
LpmTier1Pfxs/RandomMissIP6              19.75n ± 65%    94.30n ± 18%  +377.37% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     84.22n ± 41%   277.70n ±  7%  +229.71% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     82.61n ± 14%   205.15n ±  7%  +148.32% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      107.4n ± 22%    282.3n ±  5%  +162.86% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      82.83n ±  2%   233.90n ± 17%  +182.40% (p=0.000 n=10)
geomean                                 66.85n          216.1n        +223.28%
```
