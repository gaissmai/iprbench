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
the randomly generated _default-routes_ with prefix length 0 have been sorted out, they distorts the lookup times and there is no lookup miss at all.

The **RandomPrefixes** without IP versions labeling are composed of a distribution of 4 parts IPv4 to 1 part IPv6 prefixes, which is approximately the current ratio in the Internet backbone routers.

## size of the routing tables


For the multibit tries `art` and `bart` the memory consumption explodes with more than **100_000 randomly distributed IPv6** prefixes in contrast to the other algorithms, but these two algorithms are much faster than the others.

`bart` has about a factor of 10 lower memory consumption compared to `art`.

`cidrtree` is the most economical in terms of memory consumption, but this is also not a trie but a binary search tree and slower by a magnitude than the other algorithms.

```
                         │  art/size.bm   │             bart/size.bm             │
                         │     bytes      │    bytes      vs base                │
Tier1PfxSize/100            802.94Ki ± 0%   45.88Ki ± 0%  -94.29% (p=0.000 n=10)
Tier1PfxSize/1_000          7420.4Ki ± 0%   309.9Ki ± 0%  -95.82% (p=0.000 n=10)
Tier1PfxSize/10_000         47.172Mi ± 0%   1.946Mi ± 0%  -95.87% (p=0.000 n=10)
Tier1PfxSize/100_000       160.300Mi ± 0%   7.965Mi ± 0%  -95.03% (p=0.000 n=10)
Tier1PfxSize/1_000_000      378.23Mi ± 0%   34.68Mi ± 0%  -90.83% (p=0.000 n=10)
RandomPfx4Size/100          738.47Ki ± 0%   48.23Ki ± 0%  -93.47% (p=0.000 n=10)
RandomPfx4Size/1_000        7299.3Ki ± 0%   335.2Ki ± 0%  -95.41% (p=0.000 n=10)
RandomPfx4Size/10_000       58.147Mi ± 0%   2.615Mi ± 0%  -95.50% (p=0.000 n=10)
RandomPfx4Size/100_000      523.82Mi ± 0%   19.63Mi ± 0%  -96.25% (p=0.000 n=10)
RandomPfx6Size/100           732.1Ki ± 0%   164.3Ki ± 0%  -77.55% (p=0.000 n=10)
RandomPfx6Size/1_000         7.483Mi ± 0%   1.424Mi ± 0%  -80.97% (p=0.000 n=10)
RandomPfx6Size/10_000        65.25Mi ± 0%   13.48Mi ± 0%  -79.34% (p=0.000 n=10)
RandomPfx6Size/100_000       747.8Mi ± 0%   125.6Mi ± 0%  -83.20% (p=0.000 n=10)
RandomPfxSize/100           547.20Ki ± 0%   72.19Ki ± 0%  -86.81% (p=0.000 n=10)
RandomPfxSize/1_000         6087.0Ki ± 0%   745.3Ki ± 0%  -87.76% (p=0.000 n=10)
RandomPfxSize/10_000        46.487Mi ± 0%   6.947Mi ± 0%  -85.06% (p=0.000 n=10)
RandomPfxSize/100_000       473.40Mi ± 0%   64.52Mi ± 0%  -86.37% (p=0.000 n=10)
RandomPfx4Size/1_000_000                    117.5Mi ± 0%
RandomPfx6Size/1_000_000                    1.145Gi ± 0%
RandomPfxSize/1_000_000                     576.3Mi ± 0%
geomean                      21.54Mi        4.271Mi       -91.21%

                         │ cidrtree/size.bm │               bart/size.bm               │
                         │      bytes       │     bytes       vs base                  │
Tier1PfxSize/100               12.45Ki ± 0%     45.88Ki ± 0%   +268.61% (p=0.000 n=10)
Tier1PfxSize/1_000             81.51Ki ± 0%    309.91Ki ± 0%   +280.22% (p=0.000 n=10)
Tier1PfxSize/10_000            784.6Ki ± 0%    1993.1Ki ± 0%   +154.02% (p=0.000 n=10)
Tier1PfxSize/100_000           7.633Mi ± 0%     7.965Mi ± 0%     +4.36% (p=0.000 n=10)
Tier1PfxSize/1_000_000         76.30Mi ± 0%     34.68Mi ± 0%    -54.55% (p=0.000 n=10)
RandomPfx4Size/100             11.58Ki ± 0%     48.23Ki ± 0%   +316.60% (p=0.000 n=10)
RandomPfx4Size/1_000           81.51Ki ± 0%    335.23Ki ± 0%   +311.29% (p=0.000 n=10)
RandomPfx4Size/10_000          784.6Ki ± 0%    2677.8Ki ± 0%   +241.28% (p=0.000 n=10)
RandomPfx4Size/100_000         7.633Mi ± 0%    19.633Mi ± 0%   +157.21% (p=0.000 n=10)
RandomPfx4Size/1_000_000       76.30Mi ± 0%    117.50Mi ± 0%    +54.00% (p=0.000 n=10)
RandomPfx6Size/100             11.58Ki ± 0%    164.34Ki ± 0%  +1319.43% (p=0.000 n=10)
RandomPfx6Size/1_000           81.51Ki ± 0%   1458.51Ki ± 0%  +1689.41% (p=0.000 n=10)
RandomPfx6Size/10_000          785.0Ki ± 0%   13805.4Ki ± 0%  +1658.71% (p=0.000 n=10)
RandomPfx6Size/100_000         7.633Mi ± 0%   125.622Mi ± 0%  +1545.83% (p=0.000 n=10)
RandomPfx6Size/1_000_000       76.30Mi ± 0%   1172.13Mi ± 0%  +1436.26% (p=0.000 n=10)
RandomPfxSize/100              11.50Ki ± 0%     72.19Ki ± 0%   +527.72% (p=0.000 n=10)
RandomPfxSize/1_000            76.35Ki ± 0%    745.27Ki ± 0%   +876.10% (p=0.000 n=10)
RandomPfxSize/10_000           688.8Ki ± 0%    7113.3Ki ± 0%   +932.74% (p=0.000 n=10)
RandomPfxSize/100_000          6.190Mi ± 0%    64.525Mi ± 0%   +942.42% (p=0.000 n=10)
RandomPfxSize/1_000_000        56.85Mi ± 0%    576.25Mi ± 0%   +913.63% (p=0.000 n=10)
geomean                        826.6Ki          4.271Mi        +429.08%

                         │ critbitgo/size.bm │              bart/size.bm               │
                         │       bytes       │     bytes      vs base                  │
Tier1PfxSize/100                15.52Ki ± 0%    45.88Ki ± 0%   +195.67% (p=0.000 n=10)
Tier1PfxSize/1_000              115.2Ki ± 0%    309.9Ki ± 0%   +168.95% (p=0.000 n=10)
Tier1PfxSize/10_000             1.094Mi ± 0%    1.946Mi ± 0%    +77.99% (p=0.000 n=10)
Tier1PfxSize/100_000           10.913Mi ± 0%    7.965Mi ± 0%    -27.01% (p=0.000 n=10)
Tier1PfxSize/1_000_000         109.12Mi ± 0%    34.68Mi ± 0%    -68.22% (p=0.000 n=10)
RandomPfx4Size/100              14.67Ki ± 0%    48.23Ki ± 0%   +228.75% (p=0.000 n=10)
RandomPfx4Size/1_000            112.7Ki ± 0%    335.2Ki ± 0%   +197.39% (p=0.000 n=10)
RandomPfx4Size/10_000           1.071Mi ± 0%    2.615Mi ± 0%   +144.08% (p=0.000 n=10)
RandomPfx4Size/100_000          10.68Mi ± 0%    19.63Mi ± 0%    +83.75% (p=0.000 n=10)
RandomPfx4Size/1_000_000        106.8Mi ± 0%    117.5Mi ± 0%    +10.00% (p=0.000 n=10)
RandomPfx6Size/100              16.22Ki ± 0%   164.34Ki ± 0%   +913.29% (p=0.000 n=10)
RandomPfx6Size/1_000            128.3Ki ± 0%   1458.5Ki ± 0%  +1036.48% (p=0.000 n=10)
RandomPfx6Size/10_000           1.224Mi ± 0%   13.482Mi ± 0%  +1001.19% (p=0.000 n=10)
RandomPfx6Size/100_000          12.21Mi ± 0%   125.62Mi ± 0%   +928.82% (p=0.000 n=10)
RandomPfx6Size/1_000_000        122.1Mi ± 0%   1172.1Mi ± 0%   +860.18% (p=0.000 n=10)
RandomPfxSize/100               15.28Ki ± 0%    72.19Ki ± 0%   +372.39% (p=0.000 n=10)
RandomPfxSize/1_000             113.2Ki ± 0%    745.3Ki ± 0%   +558.48% (p=0.000 n=10)
RandomPfxSize/10_000            1.014Mi ± 0%    6.947Mi ± 0%   +584.82% (p=0.000 n=10)
RandomPfxSize/100_000           9.377Mi ± 0%   64.525Mi ± 0%   +588.14% (p=0.000 n=10)
RandomPfxSize/1_000_000         86.53Mi ± 0%   576.25Mi ± 0%   +565.99% (p=0.000 n=10)
geomean                         1.165Mi         4.271Mi        +266.74%

                         │ lpmtrie/size.bm │              bart/size.bm              │
                         │      bytes      │     bytes      vs base                 │
Tier1PfxSize/100              24.22Ki ± 0%    45.88Ki ± 0%   +89.42% (p=0.000 n=10)
Tier1PfxSize/1_000            202.3Ki ± 0%    309.9Ki ± 0%   +53.16% (p=0.000 n=10)
Tier1PfxSize/10_000           1.942Mi ± 0%    1.946Mi ± 0%    +0.23% (p=0.000 n=10)
Tier1PfxSize/100_000         19.330Mi ± 0%    7.965Mi ± 0%   -58.79% (p=0.000 n=10)
Tier1PfxSize/1_000_000       189.92Mi ± 0%    34.68Mi ± 0%   -81.74% (p=0.000 n=10)
RandomPfx4Size/100            23.20Ki ± 0%    48.23Ki ± 0%  +107.88% (p=0.000 n=10)
RandomPfx4Size/1_000          198.4Ki ± 0%    335.2Ki ± 0%   +68.96% (p=0.000 n=10)
RandomPfx4Size/10_000         1.895Mi ± 0%    2.615Mi ± 0%   +37.99% (p=0.000 n=10)
RandomPfx4Size/100_000        18.26Mi ± 0%    19.63Mi ± 0%    +7.52% (p=0.000 n=10)
RandomPfx4Size/1_000_000      163.5Mi ± 0%    117.5Mi ± 0%   -28.16% (p=0.000 n=10)
RandomPfx6Size/100            25.53Ki ± 0%   164.34Ki ± 0%  +543.70% (p=0.000 n=10)
RandomPfx6Size/1_000          221.5Ki ± 0%   1458.5Ki ± 0%  +558.58% (p=0.000 n=10)
RandomPfx6Size/10_000         2.127Mi ± 0%   13.482Mi ± 0%  +533.75% (p=0.000 n=10)
RandomPfx6Size/100_000        20.92Mi ± 0%   125.62Mi ± 0%  +500.59% (p=0.000 n=10)
RandomPfx6Size/1_000_000      204.0Mi ± 0%   1172.1Mi ± 0%  +474.51% (p=0.000 n=10)
RandomPfxSize/100             23.28Ki ± 0%    72.19Ki ± 0%  +210.07% (p=0.000 n=10)
RandomPfxSize/1_000           184.9Ki ± 0%    745.3Ki ± 0%  +303.14% (p=0.000 n=10)
RandomPfxSize/10_000          1.672Mi ± 0%    6.947Mi ± 0%  +315.46% (p=0.000 n=10)
RandomPfxSize/100_000         15.45Mi ± 0%    64.52Mi ± 0%  +317.73% (p=0.000 n=10)
RandomPfxSize/1_000_000       142.2Mi ± 0%    576.3Mi ± 0%  +305.21% (p=0.000 n=10)
geomean                       1.947Mi         4.271Mi       +119.39%
```

## update, insert/delete

When it comes to updates, `art` and `bart` are far superior to the other algorithms, only `critbitgo` comes close to playing in the same league 

```
                             │ art/update.bm │           bart/update.bm            │
                             │    sec/op     │   sec/op     vs base                │
Insert/Insert_into_100          129.30n ± 2%   67.99n ± 1%  -47.42% (p=0.000 n=10)
Insert/Insert_into_1_000        130.20n ± 1%   68.18n ± 1%  -47.64% (p=0.000 n=10)
Insert/Insert_into_10_000       122.90n ± 3%   68.48n ± 1%  -44.28% (p=0.000 n=10)
Insert/Insert_into_100_000      120.90n ± 1%   70.14n ± 1%  -41.98% (p=0.000 n=10)
Insert/Insert_into_1_000_000    120.40n ± 1%   67.59n ± 1%  -43.86% (p=0.000 n=10)
Delete/Delete_from_100           48.42n ± 3%   44.97n ± 1%   -7.14% (p=0.000 n=10)
Delete/Delete_from_1_000         35.87n ± 1%   44.98n ± 1%  +25.41% (p=0.000 n=10)
Delete/Delete_from_10_000        59.38n ± 0%   58.24n ± 1%   -1.91% (p=0.001 n=10)
Delete/Delete_from_100_000       59.40n ± 1%   59.94n ± 1%        ~ (p=0.079 n=10)
Delete/Delete_from_1_000_000     58.41n ± 1%   58.76n ± 2%        ~ (p=0.089 n=10)
geomean                          80.03n        60.19n       -24.78%

                             │ cidrtree/update.bm │           bart/update.bm            │
                             │       sec/op       │   sec/op     vs base                │
Insert/Insert_into_100               816.10n ± 0%   67.99n ± 1%  -91.67% (p=0.000 n=10)
Insert/Insert_into_1_000            1510.00n ± 1%   68.18n ± 1%  -95.49% (p=0.000 n=10)
Insert/Insert_into_10_000           2068.00n ± 1%   68.48n ± 1%  -96.69% (p=0.000 n=10)
Insert/Insert_into_100_000          2548.50n ± 1%   70.14n ± 1%  -97.25% (p=0.000 n=10)
Insert/Insert_into_1_000_000        2461.00n ± 0%   67.59n ± 1%  -97.25% (p=0.000 n=10)
Delete/Delete_from_100              1524.50n ± 0%   44.97n ± 1%  -97.05% (p=0.000 n=10)
Delete/Delete_from_1_000            1989.00n ± 1%   44.98n ± 1%  -97.74% (p=0.000 n=10)
Delete/Delete_from_10_000           3426.00n ± 0%   58.24n ± 1%  -98.30% (p=0.000 n=10)
Delete/Delete_from_100_000          4769.00n ± 0%   59.94n ± 1%  -98.74% (p=0.000 n=10)
Delete/Delete_from_1_000_000        4395.00n ± 0%   58.76n ± 2%  -98.66% (p=0.000 n=10)
geomean                               2.260µ        60.19n       -97.34%

                             │ critbitgo/update.bm │           bart/update.bm            │
                             │       sec/op        │   sec/op     vs base                │
Insert/Insert_into_100                107.10n ± 1%   67.99n ± 1%  -36.52% (p=0.000 n=10)
Insert/Insert_into_1_000              110.05n ± 1%   68.18n ± 1%  -38.05% (p=0.000 n=10)
Insert/Insert_into_10_000             116.70n ± 1%   68.48n ± 1%  -41.32% (p=0.000 n=10)
Insert/Insert_into_100_000            117.50n ± 5%   70.14n ± 1%  -40.30% (p=0.000 n=10)
Insert/Insert_into_1_000_000          136.10n ± 1%   67.59n ± 1%  -50.34% (p=0.000 n=10)
Delete/Delete_from_100                100.09n ± 2%   44.97n ± 1%  -55.07% (p=0.000 n=10)
Delete/Delete_from_1_000              107.85n ± 1%   44.98n ± 1%  -58.29% (p=0.000 n=10)
Delete/Delete_from_10_000             108.90n ± 1%   58.24n ± 1%  -46.52% (p=0.000 n=10)
Delete/Delete_from_100_000            119.60n ± 2%   59.94n ± 1%  -49.88% (p=0.000 n=10)
Delete/Delete_from_1_000_000          132.00n ± 1%   58.76n ± 2%  -55.48% (p=0.000 n=10)
geomean                                115.1n        60.19n       -47.71%

                             │ lpmtrie/update.bm │           bart/update.bm            │
                             │      sec/op       │   sec/op     vs base                │
Insert/Insert_into_100              301.35n ± 3%   67.99n ± 1%  -77.44% (p=0.000 n=10)
Insert/Insert_into_1_000            324.95n ± 1%   68.18n ± 1%  -79.02% (p=0.000 n=10)
Insert/Insert_into_10_000           373.85n ± 1%   68.48n ± 1%  -81.68% (p=0.000 n=10)
Insert/Insert_into_100_000          404.60n ± 1%   70.14n ± 1%  -82.66% (p=0.000 n=10)
Insert/Insert_into_1_000_000        488.30n ± 1%   67.59n ± 1%  -86.16% (p=0.000 n=10)
Delete/Delete_from_100               72.90n ± 0%   44.97n ± 1%  -38.32% (p=0.000 n=10)
Delete/Delete_from_1_000             92.51n ± 0%   44.98n ± 1%  -51.38% (p=0.000 n=10)
Delete/Delete_from_10_000           122.30n ± 1%   58.24n ± 1%  -52.38% (p=0.000 n=10)
Delete/Delete_from_100_000          183.80n ± 0%   59.94n ± 1%  -67.39% (p=0.000 n=10)
Delete/Delete_from_1_000_000        316.15n ± 0%   58.76n ± 2%  -81.41% (p=0.000 n=10)
geomean                              225.9n        60.19n       -73.36%
```

## lookup (longest-prefix-match)

In the lookup, `art` is the champion, closely followed by `bart`. 

```
                                    │ art/lookup.bm │             bart/lookup.bm             │
                                    │    sec/op     │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            45.23n ± 18%    71.15n ± 22%   +57.31% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            47.50n ± 12%    80.90n ± 33%   +70.30% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             29.02n ± 12%    73.08n ± 77%  +151.81% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             28.45n ±  1%    28.66n ±  0%    +0.76% (p=0.001 n=10)
LpmRandomPfxs100_000/RandomMatchIP4    45.23n ± 12%    89.70n ± 17%   +98.35% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6    46.99n ± 17%    90.22n ± 22%   +92.01% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4     32.22n ± 12%   112.20n ± 18%  +248.18% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6     28.61n ± 19%    94.99n ± 17%  +232.02% (p=0.000 n=10)
geomean                                36.95n          75.36n        +103.94%

                                    │ cidrtree/lookup.bm │            bart/lookup.bm            │
                                    │       sec/op       │    sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4               1026.00n ± 37%   71.15n ± 22%  -93.07% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6               1057.00n ±  9%   80.90n ± 33%  -92.35% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4                1214.00n ± 97%   73.08n ± 77%        ~ (p=0.143 n=10)
LpmTier1Pfxs/RandomMissIP6                  46.15n ±  0%   28.66n ±  0%  -37.89% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4       1110.50n ± 50%   89.70n ± 17%  -91.92% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6       1334.50n ± 35%   90.22n ± 22%  -93.24% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4         1124.0n ± 25%   112.2n ± 18%  -90.02% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6        1316.50n ± 24%   94.99n ± 17%  -92.78% (p=0.000 n=10)
geomean                                     777.3n         75.36n        -90.30%

                                    │ critbitgo/lookup.bm │            bart/lookup.bm            │
                                    │       sec/op        │    sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4                 289.90n ± 18%   71.15n ± 22%  -75.46% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6                 393.95n ± 17%   80.90n ± 33%  -79.46% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4                  543.05n ± 42%   73.08n ± 77%  -86.54% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6                  455.75n ± 43%   28.66n ±  0%  -93.71% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4         300.60n ± 25%   89.70n ± 17%  -70.16% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6         366.85n ±  7%   90.22n ± 22%  -75.41% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4           440.8n ± 27%   112.2n ± 18%  -74.55% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6          408.20n ± 12%   94.99n ± 17%  -76.73% (p=0.000 n=10)
geomean                                      392.3n         75.36n        -80.79%

                                    │ lpmtrie/lookup.bm │            bart/lookup.bm            │
                                    │      sec/op       │    sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4               230.20n ± 29%   71.15n ± 22%  -69.09% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6               199.45n ± 45%   80.90n ± 33%  -59.44% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4                188.35n ± 74%   73.08n ± 77%        ~ (p=0.052 n=10)
LpmTier1Pfxs/RandomMissIP6                 66.91n ± 42%   28.66n ±  0%  -57.16% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4       197.85n ± 11%   89.70n ± 17%  -54.66% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6       187.90n ± 10%   90.22n ± 22%  -51.98% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4         215.6n ±  7%   112.2n ± 18%  -47.97% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6        191.00n ± 16%   94.99n ± 17%  -50.27% (p=0.000 n=10)
geomean                                    175.2n         75.36n        -56.98%
```
