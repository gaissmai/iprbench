# iprbench
comparing benchmarks for some golang IP routing table implementations:

```
	github.com/tailscale/art
	github.com/gaissmai/bart
	github.com/k-sone/critbitgo
	github.com/Asphaltt/lpmtrie
	github.com/gaissmai/cidrtree
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
                         │ bart/size.bm │                art/size.bm                 │
                         │    bytes     │     bytes       vs base                    │
Tier1PfxSize/100           42.82Ki ± 0%    802.96Ki ± 0%  +1775.19% (p=0.000 n=10)
Tier1PfxSize/1_000         288.9Ki ± 0%    7420.4Ki ± 0%  +2468.17% (p=0.000 n=10)
Tier1PfxSize/10_000        1.822Mi ± 0%    47.172Mi ± 0%  +2488.91% (p=0.000 n=10)
Tier1PfxSize/100_000       7.545Mi ± 0%   160.300Mi ± 0%  +2024.67% (p=0.000 n=10)
Tier1PfxSize/1_000_000     34.33Mi ± 0%    378.23Mi ± 0%  +1001.89% (p=0.000 n=10)
RandomPfx4Size/100         44.05Ki ± 0%    725.72Ki ± 0%  +1547.31% (p=0.000 n=10)
RandomPfx4Size/1_000       305.0Ki ± 0%    7171.8Ki ± 0%  +2251.70% (p=0.000 n=10)
RandomPfx4Size/10_000      2.416Mi ± 0%    57.886Mi ± 0%  +2296.27% (p=0.000 n=10)
RandomPfx4Size/100_000     18.28Mi ± 0%    522.54Mi ± 0%  +2758.31% (p=0.000 n=10)
RandomPfx4Size/1_000_000   109.8Mi ± 0%
RandomPfx6Size/100         148.8Ki ± 0%     738.5Ki ± 0%   +396.23% (p=0.000 n=10)
RandomPfx6Size/1_000       1.250Mi ± 0%     7.558Mi ± 0%   +504.50% (p=0.000 n=10)
RandomPfx6Size/10_000      12.26Mi ± 0%     65.44Mi ± 0%   +433.81% (p=0.000 n=10)
RandomPfx6Size/100_000     115.5Mi ± 0%     747.7Mi ± 0%   +547.12% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1.049Gi ± 0%
RandomPfxSize/100          63.97Ki ± 0%    656.00Ki ± 0%   +925.50% (p=0.000 n=10)
RandomPfxSize/1_000        515.1Ki ± 0%    7541.5Ki ± 0%  +1364.04% (p=0.000 n=10)
RandomPfxSize/10_000       4.399Mi ± 0%    59.834Mi ± 0%  +1260.22% (p=0.000 n=10)
RandomPfxSize/100_000      39.10Mi ± 0%    554.91Mi ± 0%  +1319.31% (p=0.000 n=10)
RandomPfxSize/1_000_000    309.2Mi ± 0%
geomean                    3.632Mi          22.56Mi       +1277.89%                ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │  bart/size.bm  │           cidrtree/size.bm            │
                         │     bytes      │    bytes      vs base                 │
Tier1PfxSize/100             42.82Ki ± 0%   12.48Ki ± 0%   -70.86% (p=0.000 n=10)
Tier1PfxSize/1_000          288.94Ki ± 0%   81.51Ki ± 0%   -71.79% (p=0.000 n=10)
Tier1PfxSize/10_000         1865.8Ki ± 0%   784.6Ki ± 0%   -57.95% (p=0.000 n=10)
Tier1PfxSize/100_000         7.545Mi ± 0%   7.633Mi ± 0%    +1.17% (p=0.000 n=10)
Tier1PfxSize/1_000_000       34.33Mi ± 0%   76.30Mi ± 0%  +122.27% (p=0.000 n=10)
RandomPfx4Size/100           44.05Ki ± 0%   11.58Ki ± 0%   -73.72% (p=0.000 n=10)
RandomPfx4Size/1_000        304.96Ki ± 0%   81.51Ki ± 0%   -73.27% (p=0.000 n=10)
RandomPfx4Size/10_000       2473.6Ki ± 0%   784.6Ki ± 0%   -68.28% (p=0.000 n=10)
RandomPfx4Size/100_000      18.282Mi ± 0%   7.633Mi ± 0%   -58.25% (p=0.000 n=10)
RandomPfx4Size/1_000_000    109.79Mi ± 0%   76.30Mi ± 0%   -30.50% (p=0.000 n=10)
RandomPfx6Size/100          148.81Ki ± 0%   11.58Ki ± 0%   -92.22% (p=0.000 n=10)
RandomPfx6Size/1_000       1280.25Ki ± 0%   81.51Ki ± 0%   -93.63% (p=0.000 n=10)
RandomPfx6Size/10_000      12552.8Ki ± 0%   784.6Ki ± 0%   -93.75% (p=0.000 n=10)
RandomPfx6Size/100_000     115.544Mi ± 0%   7.633Mi ± 0%   -93.39% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1074.54Mi ± 0%   76.30Mi ± 0%   -92.90% (p=0.000 n=10)
RandomPfxSize/100            63.97Ki ± 0%   11.58Ki ± 0%   -81.90% (p=0.000 n=10)
RandomPfxSize/1_000         515.12Ki ± 0%   81.51Ki ± 0%   -84.18% (p=0.000 n=10)
RandomPfxSize/10_000        4504.4Ki ± 0%   784.6Ki ± 0%   -82.58% (p=0.000 n=10)
RandomPfxSize/100_000       39.097Mi ± 0%   7.633Mi ± 0%   -80.48% (p=0.000 n=10)
RandomPfxSize/1_000_000     309.15Mi ± 0%   76.30Mi ± 0%   -75.32% (p=0.000 n=10)
geomean                      3.632Mi        856.4Ki        -76.97%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm  │           critbitgo/size.bm            │
                         │     bytes     │     bytes      vs base                 │
Tier1PfxSize/100            42.82Ki ± 0%    15.64Ki ± 0%   -63.47% (p=0.000 n=10)
Tier1PfxSize/1_000          288.9Ki ± 0%    115.2Ki ± 0%   -60.12% (p=0.000 n=10)
Tier1PfxSize/10_000         1.822Mi ± 0%    1.094Mi ± 0%   -39.98% (p=0.000 n=10)
Tier1PfxSize/100_000        7.545Mi ± 0%   10.913Mi ± 0%   +44.64% (p=0.000 n=10)
Tier1PfxSize/1_000_000      34.33Mi ± 0%   109.12Mi ± 0%  +217.88% (p=0.000 n=10)
RandomPfx4Size/100          44.05Ki ± 0%    14.67Ki ± 0%   -66.70% (p=0.000 n=10)
RandomPfx4Size/1_000        305.0Ki ± 0%    112.7Ki ± 0%   -63.04% (p=0.000 n=10)
RandomPfx4Size/10_000       2.416Mi ± 0%    1.071Mi ± 0%   -55.65% (p=0.000 n=10)
RandomPfx4Size/100_000      18.28Mi ± 0%    10.68Mi ± 0%   -41.56% (p=0.000 n=10)
RandomPfx4Size/1_000_000    109.8Mi ± 0%    106.8Mi ± 0%    -2.71% (p=0.000 n=10)
RandomPfx6Size/100         148.81Ki ± 0%    16.22Ki ± 0%   -89.10% (p=0.000 n=10)
RandomPfx6Size/1_000       1280.2Ki ± 0%    128.3Ki ± 0%   -89.98% (p=0.000 n=10)
RandomPfx6Size/10_000      12.259Mi ± 0%    1.224Mi ± 0%   -90.01% (p=0.000 n=10)
RandomPfx6Size/100_000     115.54Mi ± 0%    12.21Mi ± 0%   -89.43% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1074.5Mi ± 0%    122.1Mi ± 0%   -88.64% (p=0.000 n=10)
RandomPfxSize/100           63.97Ki ± 0%    14.94Ki ± 0%   -76.65% (p=0.000 n=10)
RandomPfxSize/1_000         515.1Ki ± 0%    115.8Ki ± 0%   -77.52% (p=0.000 n=10)
RandomPfxSize/10_000        4.399Mi ± 0%    1.102Mi ± 0%   -74.95% (p=0.000 n=10)
RandomPfxSize/100_000       39.10Mi ± 0%    10.99Mi ± 0%   -71.89% (p=0.000 n=10)
RandomPfxSize/1_000_000     309.2Mi ± 0%    109.9Mi ± 0%   -64.46% (p=0.000 n=10)
geomean                     3.632Mi         1.193Mi        -67.15%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm  │            lpmtrie/size.bm             │
                         │     bytes     │     bytes      vs base                 │
Tier1PfxSize/100            42.82Ki ± 0%    24.24Ki ± 0%   -43.39% (p=0.000 n=10)
Tier1PfxSize/1_000          288.9Ki ± 0%    202.3Ki ± 0%   -29.97% (p=0.000 n=10)
Tier1PfxSize/10_000         1.822Mi ± 0%    1.942Mi ± 0%    +6.58% (p=0.000 n=10)
Tier1PfxSize/100_000        7.545Mi ± 0%   19.330Mi ± 0%  +156.21% (p=0.000 n=10)
Tier1PfxSize/1_000_000      34.33Mi ± 0%   189.92Mi ± 0%  +453.28% (p=0.000 n=10)
RandomPfx4Size/100          44.05Ki ± 0%    23.20Ki ± 0%   -47.33% (p=0.000 n=10)
RandomPfx4Size/1_000        305.0Ki ± 0%    197.6Ki ± 0%   -35.21% (p=0.000 n=10)
RandomPfx4Size/10_000       2.416Mi ± 0%    1.894Mi ± 0%   -21.61% (p=0.000 n=10)
RandomPfx4Size/100_000      18.28Mi ± 0%    18.27Mi ± 0%    -0.04% (p=0.000 n=10)
RandomPfx4Size/1_000_000    109.8Mi ± 0%    163.5Mi ± 0%   +48.95% (p=0.000 n=10)
RandomPfx6Size/100         148.81Ki ± 0%    25.53Ki ± 0%   -82.84% (p=0.000 n=10)
RandomPfx6Size/1_000       1280.2Ki ± 0%    221.6Ki ± 0%   -82.69% (p=0.000 n=10)
RandomPfx6Size/10_000      12.259Mi ± 0%    2.129Mi ± 0%   -82.63% (p=0.000 n=10)
RandomPfx6Size/100_000     115.54Mi ± 0%    20.91Mi ± 0%   -81.91% (p=0.000 n=10)
RandomPfx6Size/1_000_000   1074.5Mi ± 0%    203.9Mi ± 0%   -81.02% (p=0.000 n=10)
RandomPfxSize/100           63.97Ki ± 0%    23.56Ki ± 0%   -63.17% (p=0.000 n=10)
RandomPfxSize/1_000         515.1Ki ± 0%    203.3Ki ± 0%   -60.54% (p=0.000 n=10)
RandomPfxSize/10_000        4.399Mi ± 0%    1.945Mi ± 0%   -55.79% (p=0.000 n=10)
RandomPfxSize/100_000       39.10Mi ± 0%    18.81Mi ± 0%   -51.90% (p=0.000 n=10)
RandomPfxSize/1_000_000     309.2Mi ± 0%    172.0Mi ± 0%   -44.37% (p=0.000 n=10)
geomean                     3.632Mi         2.010Mi        -44.66%
```

## lookup (longest-prefix-match)

In the lookup, `art` is the champion, closely followed by `bart`. 

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │            art/lookup.bm             │
                                    │     sec/op     │    sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4             63.91n ± 25%   46.13n ± 12%  -27.83% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             58.82n ± 52%   45.42n ± 12%  -22.78% (p=0.001 n=10)
LpmTier1Pfxs/RandomMissIP4              78.15n ± 48%   28.20n ±  0%  -63.92% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              16.79n ±  3%   27.82n ±  0%  +65.69% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     62.53n ± 37%   52.07n ± 13%  -16.73% (p=0.003 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     76.17n ± 15%   45.21n ± 17%  -40.66% (p=0.001 n=10)
LpmRandomPfxs100_000/RandomMissIP4      94.27n ± 19%   29.88n ±  6%  -68.31% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      76.28n ± 21%   27.82n ± 13%  -63.52% (p=0.000 n=10)
geomean                                 60.04n         36.59n        -39.05%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │            cidrtree/lookup.bm             │
                                    │     sec/op     │     sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4             63.91n ± 25%   1094.00n ±  34%  +1611.65% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             58.82n ± 52%   1100.50n ±  23%  +1770.80% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              78.15n ± 48%    970.30n ±  97%          ~ (p=0.143 n=10)
LpmTier1Pfxs/RandomMissIP6              16.79n ±  3%     67.22n ± 835%   +300.33% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     62.53n ± 37%   1435.50n ±  36%  +2195.88% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     76.17n ± 15%   1452.50n ±  57%  +1806.79% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      94.27n ± 19%   1219.00n ±  20%  +1193.16% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      76.28n ± 21%   1173.50n ±  25%  +1438.51% (p=0.000 n=10)
geomean                                 60.04n           834.1n         +1289.10%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │           critbitgo/lookup.bm           │
                                    │     sec/op     │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             63.91n ± 25%   356.75n ± 14%   +458.16% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             58.82n ± 52%   448.70n ± 13%   +662.77% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              78.15n ± 48%   724.95n ± 41%   +827.70% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              16.79n ±  3%   603.00n ± 35%  +3491.42% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     62.53n ± 37%   423.90n ± 27%   +577.97% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     76.17n ± 15%   389.25n ± 13%   +410.99% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      94.27n ± 19%   527.25n ± 30%   +459.33% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      76.28n ± 21%   469.65n ± 17%   +515.73% (p=0.000 n=10)
geomean                                 60.04n          480.9n         +700.87%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │           lpmtrie/lookup.bm            │
                                    │     sec/op     │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4             63.91n ± 25%   245.50n ± 19%  +284.10% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             58.82n ± 52%   250.45n ± 63%  +325.75% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              78.15n ± 48%   233.20n ± 67%         ~ (p=0.075 n=10)
LpmTier1Pfxs/RandomMissIP6              16.79n ±  3%    88.63n ± 21%  +427.84% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     62.53n ± 37%   259.75n ±  9%  +315.43% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     76.17n ± 15%   199.95n ±  7%  +162.49% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      94.27n ± 19%   278.10n ±  5%  +195.02% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      76.28n ± 21%   235.70n ±  4%  +209.01% (p=0.000 n=10)
geomean                                 60.04n          213.6n        +255.67%
```

## update, insert/delete

`bart` is by far the fastest algorithm for updates under all competitors.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │             art/update.bm             │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100            54.14n ± 0%   105.85n ± 1%   +95.51% (p=0.000 n=10)
Insert/Insert_into_1_000          54.11n ± 0%   125.45n ± 1%  +131.84% (p=0.000 n=10)
Insert/Insert_into_10_000         54.09n ± 0%   129.05n ± 1%  +138.58% (p=0.000 n=10)
Insert/Insert_into_100_000        54.09n ± 0%   120.55n ± 2%  +122.87% (p=0.000 n=10)
Insert/Insert_into_1_000_000      55.93n ± 0%   121.00n ± 0%  +116.36% (p=0.000 n=10)
Delete/Delete_from_100            18.20n ± 0%    34.04n ± 3%   +87.03% (p=0.000 n=10)
Delete/Delete_from_1_000          41.94n ± 0%    59.06n ± 0%   +40.82% (p=0.000 n=10)
Delete/Delete_from_10_000         41.94n ± 0%    58.93n ± 0%   +40.50% (p=0.000 n=10)
Delete/Delete_from_100_000        41.99n ± 1%    58.71n ± 0%   +39.82% (p=0.000 n=10)
Delete/Delete_from_1_000_000      44.35n ± 0%    58.50n ± 0%   +31.92% (p=0.000 n=10)
geomean                           44.22n         79.57n        +79.94%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │            cidrtree/update.bm            │
                             │     sec/op     │    sec/op      vs base                   │
Insert/Insert_into_100            54.14n ± 0%    945.20n ± 0%   +1645.84% (p=0.000 n=10)
Insert/Insert_into_1_000          54.11n ± 0%   1452.00n ± 0%   +2583.42% (p=0.000 n=10)
Insert/Insert_into_10_000         54.09n ± 0%   2125.00n ± 0%   +3828.64% (p=0.000 n=10)
Insert/Insert_into_100_000        54.09n ± 0%   1595.00n ± 0%   +2848.79% (p=0.000 n=10)
Insert/Insert_into_1_000_000      55.93n ± 0%   2627.50n ± 0%   +4598.26% (p=0.000 n=10)
Delete/Delete_from_100            18.20n ± 0%   1502.00n ± 0%   +8152.75% (p=0.000 n=10)
Delete/Delete_from_1_000          41.94n ± 0%   1765.00n ± 0%   +4108.39% (p=0.000 n=10)
Delete/Delete_from_10_000         41.94n ± 0%   3444.00n ± 0%   +8111.73% (p=0.000 n=10)
Delete/Delete_from_100_000        41.99n ± 1%   3493.50n ± 0%   +8219.84% (p=0.000 n=10)
Delete/Delete_from_1_000_000      44.35n ± 0%   4819.00n ± 0%  +10765.84% (p=0.000 n=10)
geomean                           44.22n          2.125µ        +4706.01%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │          critbitgo/update.bm          │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100            54.14n ± 0%   109.95n ± 1%  +103.08% (p=0.000 n=10)
Insert/Insert_into_1_000          54.11n ± 0%   121.75n ± 1%  +125.00% (p=0.000 n=10)
Insert/Insert_into_10_000         54.09n ± 0%   125.50n ± 1%  +132.02% (p=0.000 n=10)
Insert/Insert_into_100_000        54.09n ± 0%   147.50n ± 1%  +172.69% (p=0.000 n=10)
Insert/Insert_into_1_000_000      55.93n ± 0%   152.20n ± 2%  +172.15% (p=0.000 n=10)
Delete/Delete_from_100            18.20n ± 0%    99.93n ± 1%  +449.07% (p=0.000 n=10)
Delete/Delete_from_1_000          41.94n ± 0%   108.40n ± 1%  +158.46% (p=0.000 n=10)
Delete/Delete_from_10_000         41.94n ± 0%   112.55n ± 1%  +168.36% (p=0.000 n=10)
Delete/Delete_from_100_000        41.99n ± 1%   131.40n ± 2%  +212.93% (p=0.000 n=10)
Delete/Delete_from_1_000_000      44.35n ± 0%   127.40n ± 2%  +187.26% (p=0.000 n=10)
geomean                           44.22n         122.6n       +177.35%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │           lpmtrie/update.bm            │
                             │     sec/op     │    sec/op     vs base                  │
Insert/Insert_into_100            54.14n ± 0%   314.55n ± 0%   +480.99% (p=0.000 n=10)
Insert/Insert_into_1_000          54.11n ± 0%   352.45n ± 2%   +551.36% (p=0.000 n=10)
Insert/Insert_into_10_000         54.09n ± 0%   362.60n ± 1%   +570.36% (p=0.000 n=10)
Insert/Insert_into_100_000        54.09n ± 0%   507.25n ± 1%   +837.79% (p=0.000 n=10)
Insert/Insert_into_1_000_000      55.93n ± 0%   656.80n ± 1%  +1074.43% (p=0.000 n=10)
Delete/Delete_from_100            18.20n ± 0%    77.03n ± 3%   +323.27% (p=0.000 n=10)
Delete/Delete_from_1_000          41.94n ± 0%   123.30n ± 0%   +193.99% (p=0.000 n=10)
Delete/Delete_from_10_000         41.94n ± 0%   149.50n ± 1%   +256.46% (p=0.000 n=10)
Delete/Delete_from_100_000        41.99n ± 1%   266.55n ± 0%   +534.79% (p=0.000 n=10)
Delete/Delete_from_1_000_000      44.35n ± 0%   309.70n ± 3%   +598.31% (p=0.000 n=10)
geomean                           44.22n         262.8n        +494.28%
```
