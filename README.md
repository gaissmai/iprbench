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
Tier1PfxSize/100           45.88Ki ± 0%    802.94Ki ± 0%  +1650.27% (p=0.000 n=10)
Tier1PfxSize/1_000         309.9Ki ± 0%    7420.4Ki ± 0%  +2294.40% (p=0.000 n=10)
Tier1PfxSize/10_000        1.946Mi ± 0%    47.172Mi ± 0%  +2323.50% (p=0.000 n=10)
Tier1PfxSize/100_000       7.965Mi ± 0%   160.300Mi ± 0%  +1912.53% (p=0.000 n=10)
Tier1PfxSize/1_000_000     34.68Mi ± 0%    378.23Mi ± 0%   +990.77% (p=0.000 n=10)
RandomPfx4Size/100         50.00Ki ± 0%    751.22Ki ± 0%  +1402.44% (p=0.000 n=10)
RandomPfx4Size/1_000       327.3Ki ± 0%    7178.1Ki ± 0%  +2093.32% (p=0.000 n=10)
RandomPfx4Size/10_000      2.610Mi ± 0%    57.674Mi ± 0%  +2109.42% (p=0.000 n=10)
RandomPfx4Size/100_000     19.67Mi ± 0%    522.81Mi ± 0%  +2557.61% (p=0.000 n=10)
RandomPfx4Size/1_000_000   117.4Mi ± 0%
RandomPfx6Size/100         143.4Ki ± 0%     757.6Ki ± 0%   +428.25% (p=0.000 n=10)
RandomPfx6Size/1_000       1.378Mi ± 0%     7.527Mi ± 0%   +446.24% (p=0.000 n=10)
RandomPfx6Size/10_000      13.37Mi ± 0%     65.41Mi ± 0%   +389.35% (p=0.000 n=10)
RandomPfx6Size/100_000     125.6Mi ± 0%     748.1Mi ± 0%   +495.43% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1.144Gi ± 0%
RandomPfxSize/100          63.49Ki ± 0%    700.22Ki ± 0%  +1002.84% (p=0.000 n=10)
RandomPfxSize/1_000        552.0Ki ± 0%    7420.4Ki ± 0%  +1244.24% (p=0.000 n=10)
RandomPfxSize/10_000       4.747Mi ± 0%    59.729Mi ± 0%  +1158.15% (p=0.000 n=10)
RandomPfxSize/100_000      42.05Mi ± 0%    554.86Mi ± 0%  +1219.38% (p=0.000 n=10)
RandomPfxSize/1_000_000    334.8Mi ± 0%
geomean                    3.875Mi          22.70Mi       +1202.40%                ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

                         │  bart/size.bm  │           cidrtree/size.bm            │
                         │     bytes      │    bytes      vs base                 │
Tier1PfxSize/100             45.88Ki ± 0%   12.45Ki ± 0%   -72.87% (p=0.000 n=10)
Tier1PfxSize/1_000          309.91Ki ± 0%   81.51Ki ± 0%   -73.70% (p=0.000 n=10)
Tier1PfxSize/10_000         1993.1Ki ± 0%   784.6Ki ± 0%   -60.63% (p=0.000 n=10)
Tier1PfxSize/100_000         7.965Mi ± 0%   7.633Mi ± 0%    -4.17% (p=0.000 n=10)
Tier1PfxSize/1_000_000       34.68Mi ± 0%   76.30Mi ± 0%  +120.03% (p=0.000 n=10)
RandomPfx4Size/100           50.00Ki ± 0%   11.58Ki ± 0%   -76.84% (p=0.000 n=10)
RandomPfx4Size/1_000        327.27Ki ± 0%   81.51Ki ± 0%   -75.09% (p=0.000 n=10)
RandomPfx4Size/10_000       2673.0Ki ± 0%   784.6Ki ± 0%   -70.65% (p=0.000 n=10)
RandomPfx4Size/100_000      19.672Mi ± 0%   7.633Mi ± 0%   -61.20% (p=0.000 n=10)
RandomPfx4Size/1_000_000    117.36Mi ± 0%   76.30Mi ± 0%   -34.99% (p=0.000 n=10)
RandomPfx6Size/100          143.41Ki ± 0%   11.58Ki ± 0%   -91.93% (p=0.000 n=10)
RandomPfx6Size/1_000       1410.95Ki ± 0%   81.51Ki ± 0%   -94.22% (p=0.000 n=10)
RandomPfx6Size/10_000      13688.0Ki ± 0%   785.0Ki ± 0%   -94.27% (p=0.000 n=10)
RandomPfx6Size/100_000     125.631Mi ± 0%   7.633Mi ± 0%   -93.92% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1170.97Mi ± 0%   76.30Mi ± 0%   -93.48% (p=0.000 n=10)
RandomPfxSize/100            63.49Ki ± 0%   11.58Ki ± 0%   -81.76% (p=0.000 n=10)
RandomPfxSize/1_000         552.02Ki ± 0%   81.51Ki ± 0%   -85.23% (p=0.000 n=10)
RandomPfxSize/10_000        4861.3Ki ± 0%   784.6Ki ± 0%   -83.86% (p=0.000 n=10)
RandomPfxSize/100_000       42.054Mi ± 0%   7.633Mi ± 0%   -81.85% (p=0.000 n=10)
RandomPfxSize/1_000_000     334.81Mi ± 0%   76.30Mi ± 0%   -77.21% (p=0.000 n=10)
geomean                      3.875Mi        856.3Ki        -78.42%

                         │ bart/size.bm  │           critbitgo/size.bm            │
                         │     bytes     │     bytes      vs base                 │
Tier1PfxSize/100            45.88Ki ± 0%    15.61Ki ± 0%   -65.97% (p=0.000 n=10)
Tier1PfxSize/1_000          309.9Ki ± 0%    115.2Ki ± 0%   -62.82% (p=0.000 n=10)
Tier1PfxSize/10_000         1.946Mi ± 0%    1.094Mi ± 0%   -43.82% (p=0.000 n=10)
Tier1PfxSize/100_000        7.965Mi ± 0%   10.913Mi ± 0%   +37.01% (p=0.000 n=10)
Tier1PfxSize/1_000_000      34.68Mi ± 0%   109.12Mi ± 0%  +214.68% (p=0.000 n=10)
RandomPfx4Size/100          50.00Ki ± 0%    14.67Ki ± 0%   -70.66% (p=0.000 n=10)
RandomPfx4Size/1_000        327.3Ki ± 0%    112.7Ki ± 0%   -65.56% (p=0.000 n=10)
RandomPfx4Size/10_000       2.610Mi ± 0%    1.071Mi ± 0%   -58.96% (p=0.000 n=10)
RandomPfx4Size/100_000      19.67Mi ± 0%    10.68Mi ± 0%   -45.69% (p=0.000 n=10)
RandomPfx4Size/1_000_000    117.4Mi ± 0%    106.8Mi ± 0%    -8.99% (p=0.000 n=10)
RandomPfx6Size/100         143.41Ki ± 0%    16.22Ki ± 0%   -88.69% (p=0.000 n=10)
RandomPfx6Size/1_000       1411.0Ki ± 0%    128.3Ki ± 0%   -90.90% (p=0.000 n=10)
RandomPfx6Size/10_000      13.367Mi ± 0%    1.224Mi ± 0%   -90.84% (p=0.000 n=10)
RandomPfx6Size/100_000     125.63Mi ± 0%    12.21Mi ± 0%   -90.28% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1171.0Mi ± 0%    122.1Mi ± 0%   -89.57% (p=0.000 n=10)
RandomPfxSize/100           63.49Ki ± 0%    14.94Ki ± 0%   -76.47% (p=0.000 n=10)
RandomPfxSize/1_000         552.0Ki ± 0%    115.8Ki ± 0%   -79.03% (p=0.000 n=10)
RandomPfxSize/10_000        4.747Mi ± 0%    1.102Mi ± 0%   -76.79% (p=0.000 n=10)
RandomPfxSize/100_000       42.05Mi ± 0%    10.99Mi ± 0%   -73.87% (p=0.000 n=10)
RandomPfxSize/1_000_000     334.8Mi ± 0%    109.9Mi ± 0%   -67.19% (p=0.000 n=10)
geomean                     3.875Mi         1.193Mi        -69.20%

                         │ bart/size.bm  │            lpmtrie/size.bm             │
                         │     bytes     │     bytes      vs base                 │
Tier1PfxSize/100            45.88Ki ± 0%    24.22Ki ± 0%   -47.21% (p=0.000 n=10)
Tier1PfxSize/1_000          309.9Ki ± 0%    202.3Ki ± 0%   -34.71% (p=0.000 n=10)
Tier1PfxSize/10_000         1.946Mi ± 0%    1.942Mi ± 0%    -0.22% (p=0.000 n=10)
Tier1PfxSize/100_000        7.965Mi ± 0%   19.330Mi ± 0%  +142.69% (p=0.000 n=10)
Tier1PfxSize/1_000_000      34.68Mi ± 0%   189.92Mi ± 0%  +447.70% (p=0.000 n=10)
RandomPfx4Size/100          50.00Ki ± 0%    23.20Ki ± 0%   -53.59% (p=0.000 n=10)
RandomPfx4Size/1_000        327.3Ki ± 0%    198.6Ki ± 0%   -39.31% (p=0.000 n=10)
RandomPfx4Size/10_000       2.610Mi ± 0%    1.897Mi ± 0%   -27.34% (p=0.000 n=10)
RandomPfx4Size/100_000      19.67Mi ± 0%    18.29Mi ± 0%    -7.03% (p=0.000 n=10)
RandomPfx4Size/1_000_000    117.4Mi ± 0%    163.6Mi ± 0%   +39.36% (p=0.000 n=10)
RandomPfx6Size/100         143.41Ki ± 0%    25.53Ki ± 0%   -82.20% (p=0.000 n=10)
RandomPfx6Size/1_000       1411.0Ki ± 0%    221.9Ki ± 0%   -84.27% (p=0.000 n=10)
RandomPfx6Size/10_000      13.367Mi ± 0%    2.131Mi ± 0%   -84.05% (p=0.000 n=10)
RandomPfx6Size/100_000     125.63Mi ± 0%    20.92Mi ± 0%   -83.35% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1171.0Mi ± 0%    204.0Mi ± 0%   -82.58% (p=0.000 n=10)
RandomPfxSize/100           63.49Ki ± 0%    23.56Ki ± 0%   -62.89% (p=0.000 n=10)
RandomPfxSize/1_000         552.0Ki ± 0%    203.0Ki ± 0%   -63.22% (p=0.000 n=10)
RandomPfxSize/10_000        4.747Mi ± 0%    1.942Mi ± 0%   -59.10% (p=0.000 n=10)
RandomPfxSize/100_000       42.05Mi ± 0%    18.80Mi ± 0%   -55.29% (p=0.000 n=10)
RandomPfxSize/1_000_000     334.8Mi ± 0%    171.9Mi ± 0%   -48.67% (p=0.000 n=10)
geomean                     3.875Mi         2.011Mi        -48.10%
```

## update, insert/delete

When it comes to updates, `art` and `bart` are far superior
to the other algorithms, only `critbitgo` comes close to playing in the same league .

```
                             │ bart/update.bm │            art/update.bm             │
                             │     sec/op     │    sec/op     vs base                │
Insert/Insert_into_100            71.30n ± 2%   107.10n ± 1%  +50.20% (p=0.000 n=10)
Insert/Insert_into_1_000          77.64n ± 2%   128.30n ± 1%  +65.25% (p=0.000 n=10)
Insert/Insert_into_10_000         82.63n ± 1%   133.80n ± 3%  +61.94% (p=0.000 n=10)
Insert/Insert_into_100_000        76.39n ± 5%   122.95n ± 2%  +60.96% (p=0.000 n=10)
Insert/Insert_into_1_000_000      72.13n ± 3%   123.55n ± 2%  +71.29% (p=0.000 n=10)
Delete/Delete_from_100            34.70n ± 2%    37.99n ± 7%   +9.47% (p=0.000 n=10)
Delete/Delete_from_1_000          60.80n ± 2%    59.82n ± 2%   -1.62% (p=0.012 n=10)
Delete/Delete_from_10_000         59.86n ± 2%    60.30n ± 1%        ~ (p=0.684 n=10)
Delete/Delete_from_100_000        58.29n ± 2%    60.62n ± 1%   +4.00% (p=0.000 n=10)
Delete/Delete_from_1_000_000      58.97n ± 2%    60.19n ± 1%   +2.07% (p=0.001 n=10)
geomean                           63.66n         82.13n       +29.00%

                             │ bart/update.bm │            cidrtree/update.bm            │
                             │     sec/op     │    sec/op      vs base                   │
Insert/Insert_into_100            71.30n ± 2%    914.20n ± 1%   +1182.10% (p=0.000 n=10)
Insert/Insert_into_1_000          77.64n ± 2%   1442.50n ± 1%   +1757.93% (p=0.000 n=10)
Insert/Insert_into_10_000         82.63n ± 1%   1729.00n ± 1%   +1992.59% (p=0.000 n=10)
Insert/Insert_into_100_000        76.39n ± 5%   2627.00n ± 2%   +3339.16% (p=0.000 n=10)
Insert/Insert_into_1_000_000      72.13n ± 3%   2677.00n ± 4%   +3611.35% (p=0.000 n=10)
Delete/Delete_from_100            34.70n ± 2%   1049.00n ± 1%   +2923.05% (p=0.000 n=10)
Delete/Delete_from_1_000          60.80n ± 2%   2832.50n ± 1%   +4558.72% (p=0.000 n=10)
Delete/Delete_from_10_000         59.86n ± 2%   1377.00n ± 1%   +2200.37% (p=0.000 n=10)
Delete/Delete_from_100_000        58.29n ± 2%   5115.00n ± 1%   +8675.09% (p=0.000 n=10)
Delete/Delete_from_1_000_000      58.97n ± 2%   6363.00n ± 1%  +10691.15% (p=0.000 n=10)
geomean                           63.66n          2.153µ        +3281.00%

                             │ bart/update.bm │          critbitgo/update.bm          │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100            71.30n ± 2%   107.90n ± 2%   +51.32% (p=0.000 n=10)
Insert/Insert_into_1_000          77.64n ± 2%   114.85n ± 2%   +47.93% (p=0.000 n=10)
Insert/Insert_into_10_000         82.63n ± 1%   129.95n ± 1%   +57.28% (p=0.000 n=10)
Insert/Insert_into_100_000        76.39n ± 5%   151.90n ± 2%   +98.86% (p=0.000 n=10)
Insert/Insert_into_1_000_000      72.13n ± 3%   168.80n ± 1%  +134.02% (p=0.000 n=10)
Delete/Delete_from_100            34.70n ± 2%   101.90n ± 3%  +193.66% (p=0.000 n=10)
Delete/Delete_from_1_000          60.80n ± 2%   109.30n ± 2%   +79.77% (p=0.000 n=10)
Delete/Delete_from_10_000         59.86n ± 2%   112.75n ± 2%   +88.36% (p=0.000 n=10)
Delete/Delete_from_100_000        58.29n ± 2%   132.50n ± 3%  +127.31% (p=0.000 n=10)
Delete/Delete_from_1_000_000      58.97n ± 2%   127.30n ± 1%  +115.89% (p=0.000 n=10)
geomean                           63.66n         124.2n        +95.11%

                             │ bart/update.bm │           lpmtrie/update.bm           │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100            71.30n ± 2%   301.45n ± 4%  +322.76% (p=0.000 n=10)
Insert/Insert_into_1_000          77.64n ± 2%   342.80n ± 2%  +341.52% (p=0.000 n=10)
Insert/Insert_into_10_000         82.63n ± 1%   359.25n ± 4%  +334.80% (p=0.000 n=10)
Insert/Insert_into_100_000        76.39n ± 5%   478.45n ± 2%  +526.37% (p=0.000 n=10)
Insert/Insert_into_1_000_000      72.13n ± 3%   605.40n ± 1%  +739.32% (p=0.000 n=10)
Delete/Delete_from_100            34.70n ± 2%    73.51n ± 1%  +111.84% (p=0.000 n=10)
Delete/Delete_from_1_000          60.80n ± 2%   122.15n ± 0%  +100.90% (p=0.000 n=10)
Delete/Delete_from_10_000         59.86n ± 2%   142.15n ± 0%  +137.47% (p=0.000 n=10)
Delete/Delete_from_100_000        58.29n ± 2%   271.10n ± 0%  +365.09% (p=0.000 n=10)
Delete/Delete_from_1_000_000      58.97n ± 2%   291.85n ± 0%  +394.95% (p=0.000 n=10)
geomean                           63.66n         253.3n       +297.81%
```

## lookup (longest-prefix-match)

In the lookup, `art` is the champion, closely followed by `bart`. 

```
                                    │ bart/lookup.bm │            art/lookup.bm             │
                                    │     sec/op     │    sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4             69.45n ± 31%   52.03n ± 18%        ~ (p=0.143 n=10)
LpmTier1Pfxs/RandomMatchIP6             65.16n ± 14%   52.32n ±  9%  -19.69% (p=0.001 n=10)
LpmTier1Pfxs/RandomMissIP4              87.31n ± 45%   36.28n ±  0%  -58.45% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              19.72n ±  0%   36.24n ±  1%  +83.70% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     80.14n ± 31%   60.94n ± 13%  -23.96% (p=0.007 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     71.68n ± 27%   53.04n ± 16%        ~ (p=0.060 n=10)
LpmRandomPfxs100_000/RandomMissIP4     109.27n ± 23%   36.18n ±  3%  -66.89% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      83.51n ± 17%   36.13n ±  0%  -56.74% (p=0.000 n=10)
geomean                                 67.05n         44.41n        -33.77%

                                    │ bart/lookup.bm │            cidrtree/lookup.bm            │
                                    │     sec/op     │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             69.45n ± 31%   1078.00n ± 22%  +1452.20% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             65.16n ± 14%   1213.00n ± 21%  +1761.71% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              87.31n ± 45%   1504.00n ± 32%  +1622.60% (p=0.002 n=10)
LpmTier1Pfxs/RandomMissIP6              19.72n ±  0%    820.90n ± 94%  +4061.72% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     80.14n ± 31%   1287.00n ± 22%  +1505.84% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     71.68n ± 27%   1352.50n ± 38%  +1786.86% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      109.3n ± 23%    1212.0n ± 12%  +1009.23% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      83.51n ± 17%   1423.50n ± 14%  +1604.59% (p=0.000 n=10)
geomean                                 67.05n           1.218µ        +1716.77%

                                    │ bart/lookup.bm │           critbitgo/lookup.bm           │
                                    │     sec/op     │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             69.45n ± 31%   547.85n ± 29%   +688.84% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             65.16n ± 14%   541.30n ± 14%   +730.79% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              87.31n ± 45%   700.05n ±  9%   +701.80% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              19.72n ±  0%   647.30n ± 17%  +3181.62% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     80.14n ± 31%   314.70n ± 19%   +292.66% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     71.68n ± 27%   451.30n ± 22%   +529.60% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      109.3n ± 23%    554.1n ± 16%   +407.12% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      83.51n ± 17%   435.35n ± 39%   +421.31% (p=0.000 n=10)
geomean                                 67.05n          510.4n         +661.19%

                                    │ bart/lookup.bm │           lpmtrie/lookup.bm            │
                                    │     sec/op     │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4             69.45n ± 31%   246.30n ± 17%  +254.64% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             65.16n ± 14%   226.85n ± 67%  +248.17% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              87.31n ± 45%   237.45n ± 69%         ~ (p=0.052 n=10)
LpmTier1Pfxs/RandomMissIP6              19.72n ±  0%    83.56n ± 20%  +323.62% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     80.14n ± 31%   274.40n ±  5%  +242.38% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     71.68n ± 27%   213.05n ± 12%  +197.22% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      109.3n ± 23%    294.3n ±  4%  +169.35% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      83.51n ± 17%   239.00n ± 23%  +186.19% (p=0.000 n=10)
geomean                                 67.05n          215.0n        +220.59%
```
