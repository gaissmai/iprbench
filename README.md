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


For the multibit tries `art` and `bart` the memory consumption explodes with more than **100_000 randomly distributed IPv6**
prefixes in contrast to the other algorithms, but these two algorithms are much faster than the others.

`bart` has about a factor of 10 lower memory consumption compared to `art`.

`cidrtree` is the most economical in terms of memory consumption, but this is also not a trie but a binary search tree and
slower by a magnitude than the other algorithms.

```
                         │ bart/size.bm │                art/size.bm                 │
                         │    bytes     │     bytes       vs base                    │
Tier1PfxSize/100           45.88Ki ± 0%    802.94Ki ± 0%  +1650.27% (p=0.000 n=10)
Tier1PfxSize/1_000         309.9Ki ± 0%    7420.4Ki ± 0%  +2294.40% (p=0.000 n=10)
Tier1PfxSize/10_000        1.946Mi ± 0%    47.172Mi ± 0%  +2323.50% (p=0.000 n=10)
Tier1PfxSize/100_000       7.965Mi ± 0%   160.300Mi ± 0%  +1912.53% (p=0.000 n=10)
Tier1PfxSize/1_000_000     34.68Mi ± 0%    378.23Mi ± 0%   +990.77% (p=0.000 n=10)
RandomPfx4Size/100         48.23Ki ± 0%    738.47Ki ± 0%  +1431.00% (p=0.000 n=10)
RandomPfx4Size/1_000       335.2Ki ± 0%    7299.3Ki ± 0%  +2077.36% (p=0.000 n=10)
RandomPfx4Size/10_000      2.615Mi ± 0%    58.147Mi ± 0%  +2123.57% (p=0.000 n=10)
RandomPfx4Size/100_000     19.63Mi ± 0%    523.82Mi ± 0%  +2568.09% (p=0.000 n=10)
RandomPfx4Size/1_000_000   117.5Mi ± 0%
RandomPfx6Size/100         164.3Ki ± 0%     732.1Ki ± 0%   +345.46% (p=0.000 n=10)
RandomPfx6Size/1_000       1.424Mi ± 0%     7.483Mi ± 0%   +425.37% (p=0.000 n=10)
RandomPfx6Size/10_000      13.48Mi ± 0%     65.25Mi ± 0%   +383.99% (p=0.000 n=10)
RandomPfx6Size/100_000     125.6Mi ± 0%     747.8Mi ± 0%   +495.30% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1.145Gi ± 0%
RandomPfxSize/100          72.19Ki ± 0%    547.20Ki ± 0%   +658.03% (p=0.000 n=10)
RandomPfxSize/1_000        745.3Ki ± 0%    6087.0Ki ± 0%   +716.75% (p=0.000 n=10)
RandomPfxSize/10_000       6.947Mi ± 0%    46.487Mi ± 0%   +569.21% (p=0.000 n=10)
RandomPfxSize/100_000      64.52Mi ± 0%    473.40Mi ± 0%   +633.67% (p=0.000 n=10)
RandomPfxSize/1_000_000    576.3Mi ± 0%
geomean                    4.271Mi          21.54Mi       +1038.25%                ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

                         │  bart/size.bm  │           cidrtree/size.bm            │
                         │     bytes      │    bytes      vs base                 │
Tier1PfxSize/100             45.88Ki ± 0%   12.45Ki ± 0%   -72.87% (p=0.000 n=10)
Tier1PfxSize/1_000          309.91Ki ± 0%   81.51Ki ± 0%   -73.70% (p=0.000 n=10)
Tier1PfxSize/10_000         1993.1Ki ± 0%   784.6Ki ± 0%   -60.63% (p=0.000 n=10)
Tier1PfxSize/100_000         7.965Mi ± 0%   7.633Mi ± 0%    -4.17% (p=0.000 n=10)
Tier1PfxSize/1_000_000       34.68Mi ± 0%   76.30Mi ± 0%  +120.03% (p=0.000 n=10)
RandomPfx4Size/100           48.23Ki ± 0%   11.58Ki ± 0%   -76.00% (p=0.000 n=10)
RandomPfx4Size/1_000        335.23Ki ± 0%   81.51Ki ± 0%   -75.69% (p=0.000 n=10)
RandomPfx4Size/10_000       2677.8Ki ± 0%   784.6Ki ± 0%   -70.70% (p=0.000 n=10)
RandomPfx4Size/100_000      19.633Mi ± 0%   7.633Mi ± 0%   -61.12% (p=0.000 n=10)
RandomPfx4Size/1_000_000    117.50Mi ± 0%   76.30Mi ± 0%   -35.07% (p=0.000 n=10)
RandomPfx6Size/100          164.34Ki ± 0%   11.58Ki ± 0%   -92.95% (p=0.000 n=10)
RandomPfx6Size/1_000       1458.51Ki ± 0%   81.51Ki ± 0%   -94.41% (p=0.000 n=10)
RandomPfx6Size/10_000      13805.4Ki ± 0%   785.0Ki ± 0%   -94.31% (p=0.000 n=10)
RandomPfx6Size/100_000     125.622Mi ± 0%   7.633Mi ± 0%   -93.92% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1172.13Mi ± 0%   76.30Mi ± 0%   -93.49% (p=0.000 n=10)
RandomPfxSize/100            72.19Ki ± 0%   11.50Ki ± 0%   -84.07% (p=0.000 n=10)
RandomPfxSize/1_000         745.27Ki ± 0%   76.35Ki ± 0%   -89.76% (p=0.000 n=10)
RandomPfxSize/10_000        7113.3Ki ± 0%   688.8Ki ± 0%   -90.32% (p=0.000 n=10)
RandomPfxSize/100_000       64.525Mi ± 0%   6.190Mi ± 0%   -90.41% (p=0.000 n=10)
RandomPfxSize/1_000_000     576.25Mi ± 0%   56.85Mi ± 0%   -90.13% (p=0.000 n=10)
geomean                      4.271Mi        826.6Ki        -81.10%

                         │ bart/size.bm  │           critbitgo/size.bm            │
                         │     bytes     │     bytes      vs base                 │
Tier1PfxSize/100            45.88Ki ± 0%    15.52Ki ± 0%   -66.18% (p=0.000 n=10)
Tier1PfxSize/1_000          309.9Ki ± 0%    115.2Ki ± 0%   -62.82% (p=0.000 n=10)
Tier1PfxSize/10_000         1.946Mi ± 0%    1.094Mi ± 0%   -43.82% (p=0.000 n=10)
Tier1PfxSize/100_000        7.965Mi ± 0%   10.913Mi ± 0%   +37.01% (p=0.000 n=10)
Tier1PfxSize/1_000_000      34.68Mi ± 0%   109.12Mi ± 0%  +214.68% (p=0.000 n=10)
RandomPfx4Size/100          48.23Ki ± 0%    14.67Ki ± 0%   -69.58% (p=0.000 n=10)
RandomPfx4Size/1_000        335.2Ki ± 0%    112.7Ki ± 0%   -66.37% (p=0.000 n=10)
RandomPfx4Size/10_000       2.615Mi ± 0%    1.071Mi ± 0%   -59.03% (p=0.000 n=10)
RandomPfx4Size/100_000      19.63Mi ± 0%    10.68Mi ± 0%   -45.58% (p=0.000 n=10)
RandomPfx4Size/1_000_000    117.5Mi ± 0%    106.8Mi ± 0%    -9.09% (p=0.000 n=10)
RandomPfx6Size/100         164.34Ki ± 0%    16.22Ki ± 0%   -90.13% (p=0.000 n=10)
RandomPfx6Size/1_000       1458.5Ki ± 0%    128.3Ki ± 0%   -91.20% (p=0.000 n=10)
RandomPfx6Size/10_000      13.482Mi ± 0%    1.224Mi ± 0%   -90.92% (p=0.000 n=10)
RandomPfx6Size/100_000     125.62Mi ± 0%    12.21Mi ± 0%   -90.28% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1172.1Mi ± 0%    122.1Mi ± 0%   -89.59% (p=0.000 n=10)
RandomPfxSize/100           72.19Ki ± 0%    15.28Ki ± 0%   -78.83% (p=0.000 n=10)
RandomPfxSize/1_000         745.3Ki ± 0%    113.2Ki ± 0%   -84.81% (p=0.000 n=10)
RandomPfxSize/10_000        6.947Mi ± 0%    1.014Mi ± 0%   -85.40% (p=0.000 n=10)
RandomPfxSize/100_000      64.525Mi ± 0%    9.377Mi ± 0%   -85.47% (p=0.000 n=10)
RandomPfxSize/1_000_000    576.25Mi ± 0%    86.53Mi ± 0%   -84.98% (p=0.000 n=10)
geomean                     4.271Mi         1.165Mi        -72.73%

                         │ bart/size.bm  │            lpmtrie/size.bm             │
                         │     bytes     │     bytes      vs base                 │
Tier1PfxSize/100            45.88Ki ± 0%    24.22Ki ± 0%   -47.21% (p=0.000 n=10)
Tier1PfxSize/1_000          309.9Ki ± 0%    202.3Ki ± 0%   -34.71% (p=0.000 n=10)
Tier1PfxSize/10_000         1.946Mi ± 0%    1.942Mi ± 0%    -0.22% (p=0.000 n=10)
Tier1PfxSize/100_000        7.965Mi ± 0%   19.330Mi ± 0%  +142.69% (p=0.000 n=10)
Tier1PfxSize/1_000_000      34.68Mi ± 0%   189.92Mi ± 0%  +447.70% (p=0.000 n=10)
RandomPfx4Size/100          48.23Ki ± 0%    23.20Ki ± 0%   -51.90% (p=0.000 n=10)
RandomPfx4Size/1_000        335.2Ki ± 0%    198.4Ki ± 0%   -40.81% (p=0.000 n=10)
RandomPfx4Size/10_000       2.615Mi ± 0%    1.895Mi ± 0%   -27.53% (p=0.000 n=10)
RandomPfx4Size/100_000      19.63Mi ± 0%    18.26Mi ± 0%    -7.00% (p=0.000 n=10)
RandomPfx4Size/1_000_000    117.5Mi ± 0%    163.5Mi ± 0%   +39.19% (p=0.000 n=10)
RandomPfx6Size/100         164.34Ki ± 0%    25.53Ki ± 0%   -84.46% (p=0.000 n=10)
RandomPfx6Size/1_000       1458.5Ki ± 0%    221.5Ki ± 0%   -84.82% (p=0.000 n=10)
RandomPfx6Size/10_000      13.482Mi ± 0%    2.127Mi ± 0%   -84.22% (p=0.000 n=10)
RandomPfx6Size/100_000     125.62Mi ± 0%    20.92Mi ± 0%   -83.35% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1172.1Mi ± 0%    204.0Mi ± 0%   -82.59% (p=0.000 n=10)
RandomPfxSize/100           72.19Ki ± 0%    23.28Ki ± 0%   -67.75% (p=0.000 n=10)
RandomPfxSize/1_000         745.3Ki ± 0%    184.9Ki ± 0%   -75.19% (p=0.000 n=10)
RandomPfxSize/10_000        6.947Mi ± 0%    1.672Mi ± 0%   -75.93% (p=0.000 n=10)
RandomPfxSize/100_000       64.52Mi ± 0%    15.45Mi ± 0%   -76.06% (p=0.000 n=10)
RandomPfxSize/1_000_000     576.3Mi ± 0%    142.2Mi ± 0%   -75.32% (p=0.000 n=10)
geomean                     4.271Mi         1.947Mi        -54.42%
```

## update, insert/delete

When it comes to updates, `art` and `bart` are far superior to the other algorithms, only `critbitgo` comes close to playing in the same league 

```
                             │ bart/update.bm │            art/update.bm             │
                             │     sec/op     │    sec/op     vs base                │
Insert/Insert_into_100            67.99n ± 1%   129.30n ± 2%  +90.18% (p=0.000 n=10)
Insert/Insert_into_1_000          68.18n ± 1%   130.20n ± 1%  +90.98% (p=0.000 n=10)
Insert/Insert_into_10_000         68.48n ± 1%   122.90n ± 3%  +79.48% (p=0.000 n=10)
Insert/Insert_into_100_000        70.14n ± 1%   120.90n ± 1%  +72.36% (p=0.000 n=10)
Insert/Insert_into_1_000_000      67.59n ± 1%   120.40n ± 1%  +78.13% (p=0.000 n=10)
Delete/Delete_from_100            44.97n ± 1%    48.42n ± 3%   +7.68% (p=0.000 n=10)
Delete/Delete_from_1_000          44.98n ± 1%    35.87n ± 1%  -20.26% (p=0.000 n=10)
Delete/Delete_from_10_000         58.24n ± 1%    59.38n ± 0%   +1.95% (p=0.001 n=10)
Delete/Delete_from_100_000        59.94n ± 1%    59.40n ± 1%        ~ (p=0.079 n=10)
Delete/Delete_from_1_000_000      58.76n ± 2%    58.41n ± 1%        ~ (p=0.089 n=10)
geomean                           60.19n         80.03n       +32.95%

                             │ bart/update.bm │           cidrtree/update.bm            │
                             │     sec/op     │    sec/op      vs base                  │
Insert/Insert_into_100            67.99n ± 1%    816.10n ± 0%  +1100.32% (p=0.000 n=10)
Insert/Insert_into_1_000          68.18n ± 1%   1510.00n ± 1%  +2114.89% (p=0.000 n=10)
Insert/Insert_into_10_000         68.48n ± 1%   2068.00n ± 1%  +2920.08% (p=0.000 n=10)
Insert/Insert_into_100_000        70.14n ± 1%   2548.50n ± 1%  +3533.19% (p=0.000 n=10)
Insert/Insert_into_1_000_000      67.59n ± 1%   2461.00n ± 0%  +3541.07% (p=0.000 n=10)
Delete/Delete_from_100            44.97n ± 1%   1524.50n ± 0%  +3290.41% (p=0.000 n=10)
Delete/Delete_from_1_000          44.98n ± 1%   1989.00n ± 1%  +4321.97% (p=0.000 n=10)
Delete/Delete_from_10_000         58.24n ± 1%   3426.00n ± 0%  +5782.55% (p=0.000 n=10)
Delete/Delete_from_100_000        59.94n ± 1%   4769.00n ± 0%  +7855.63% (p=0.000 n=10)
Delete/Delete_from_1_000_000      58.76n ± 2%   4395.00n ± 0%  +7379.58% (p=0.000 n=10)
geomean                           60.19n          2.260µ       +3655.06%

                             │ bart/update.bm │          critbitgo/update.bm          │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100            67.99n ± 1%   107.10n ± 1%   +57.52% (p=0.000 n=10)
Insert/Insert_into_1_000          68.18n ± 1%   110.05n ± 1%   +61.42% (p=0.000 n=10)
Insert/Insert_into_10_000         68.48n ± 1%   116.70n ± 1%   +70.43% (p=0.000 n=10)
Insert/Insert_into_100_000        70.14n ± 1%   117.50n ± 5%   +67.51% (p=0.000 n=10)
Insert/Insert_into_1_000_000      67.59n ± 1%   136.10n ± 1%  +101.36% (p=0.000 n=10)
Delete/Delete_from_100            44.97n ± 1%   100.09n ± 2%  +122.58% (p=0.000 n=10)
Delete/Delete_from_1_000          44.98n ± 1%   107.85n ± 1%  +139.77% (p=0.000 n=10)
Delete/Delete_from_10_000         58.24n ± 1%   108.90n ± 1%   +86.98% (p=0.000 n=10)
Delete/Delete_from_100_000        59.94n ± 1%   119.60n ± 2%   +99.52% (p=0.000 n=10)
Delete/Delete_from_1_000_000      58.76n ± 2%   132.00n ± 1%  +124.64% (p=0.000 n=10)
geomean                           60.19n         115.1n        +91.22%

                             │ bart/update.bm │           lpmtrie/update.bm           │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100            67.99n ± 1%   301.35n ± 3%  +343.23% (p=0.000 n=10)
Insert/Insert_into_1_000          68.18n ± 1%   324.95n ± 1%  +376.64% (p=0.000 n=10)
Insert/Insert_into_10_000         68.48n ± 1%   373.85n ± 1%  +445.97% (p=0.000 n=10)
Insert/Insert_into_100_000        70.14n ± 1%   404.60n ± 1%  +476.81% (p=0.000 n=10)
Insert/Insert_into_1_000_000      67.59n ± 1%   488.30n ± 1%  +622.44% (p=0.000 n=10)
Delete/Delete_from_100            44.97n ± 1%    72.90n ± 0%   +62.13% (p=0.000 n=10)
Delete/Delete_from_1_000          44.98n ± 1%    92.51n ± 0%  +105.67% (p=0.000 n=10)
Delete/Delete_from_10_000         58.24n ± 1%   122.30n ± 1%  +109.99% (p=0.000 n=10)
Delete/Delete_from_100_000        59.94n ± 1%   183.80n ± 0%  +206.61% (p=0.000 n=10)
Delete/Delete_from_1_000_000      58.76n ± 2%   316.15n ± 0%  +438.04% (p=0.000 n=10)
geomean                           60.19n         225.9n       +275.36%

```

## lookup (longest-prefix-match)

In the lookup, `art` is the champion, closely followed by `bart`. 

```
                                    │ bart/lookup.bm │            art/lookup.bm             │
                                    │     sec/op     │    sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4             77.22n ± 26%   46.34n ± 16%  -39.99% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             79.11n ± 31%   46.18n ± 14%  -41.62% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              82.19n ± 33%   29.14n ±  7%  -64.55% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              28.70n ± 47%   28.82n ±  0%        ~ (p=0.382 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     90.89n ± 24%   46.23n ± 18%  -49.14% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     87.09n ± 28%   50.70n ± 10%  -41.79% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4     130.35n ± 28%   30.21n ± 11%  -76.82% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      94.00n ± 17%   30.02n ± 16%  -68.07% (p=0.000 n=10)
geomean                                 78.19n         37.39n        -52.18%

                                    │ bart/lookup.bm │            cidrtree/lookup.bm             │
                                    │     sec/op     │     sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4             77.22n ± 26%   1002.50n ±  23%  +1198.24% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             79.11n ± 31%   1383.00n ±   8%  +1648.20% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              82.19n ± 33%   1311.50n ±  13%  +1495.60% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              28.70n ± 47%    490.38n ± 262%  +1608.92% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     90.89n ± 24%   1312.50n ±  26%  +1343.97% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     87.09n ± 28%   1717.00n ±  29%  +1871.41% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      130.3n ± 28%    1368.5n ±  19%   +949.87% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      94.00n ± 17%   1442.00n ±   8%  +1433.96% (p=0.000 n=10)
geomean                                 78.19n           1.188µ         +1419.18%

                                    │ bart/lookup.bm │           critbitgo/lookup.bm           │
                                    │     sec/op     │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             77.22n ± 26%   315.80n ± 16%   +308.96% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             79.11n ± 31%   432.10n ± 13%   +446.20% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              82.19n ± 33%   593.35n ± 46%   +621.88% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              28.70n ± 47%   518.75n ± 27%  +1707.81% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     90.89n ± 24%   322.35n ± 20%   +254.64% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     87.09n ± 28%   400.85n ± 21%   +360.24% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      130.3n ± 28%    564.9n ± 16%   +333.33% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      94.00n ± 17%   479.95n ± 45%   +410.56% (p=0.000 n=10)
geomean                                 78.19n          442.6n         +465.96%

                                    │ bart/lookup.bm │           lpmtrie/lookup.bm            │
                                    │     sec/op     │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4             77.22n ± 26%   246.30n ± 17%  +218.96% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             79.11n ± 31%   226.85n ± 67%  +186.75% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              82.19n ± 33%   237.45n ± 69%  +188.89% (p=0.007 n=10)
LpmTier1Pfxs/RandomMissIP6              28.70n ± 47%    83.56n ± 20%  +191.20% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     90.89n ± 24%   274.40n ±  5%  +201.89% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     87.09n ± 28%   213.05n ± 12%  +144.62% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      130.3n ± 28%    294.3n ±  4%  +125.78% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      94.00n ± 17%   239.00n ± 23%  +154.24% (p=0.000 n=10)
geomean                                 78.19n          215.0n        +174.90%
```
