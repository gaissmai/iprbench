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


The memory consumption of the multibit trie `art` explodes with more
than **100_000 randomly distributed IPv6** prefixes in contrast to the other algorithms.

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

## lookup (longest-prefix-match)

In the lookup, `art` is the champion, closely followed by `bart`. 

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │            art/lookup.bm             │
                                    │     sec/op     │    sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4             62.66n ± 23%   46.12n ± 12%  -26.39% (p=0.005 n=10)
LpmTier1Pfxs/RandomMatchIP6             59.30n ± 31%   45.26n ± 13%  -23.68% (p=0.023 n=10)
LpmTier1Pfxs/RandomMissIP4              77.69n ± 48%   28.18n ±  0%  -63.73% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              16.77n ±  0%   27.81n ±  0%  +65.88% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     65.63n ± 33%   46.06n ± 12%  -29.82% (p=0.002 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     66.71n ± 21%   45.02n ± 17%  -32.51% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      99.32n ± 22%   29.87n ±  7%  -69.93% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      76.59n ± 25%   29.80n ±  7%  -61.10% (p=0.000 n=10)
geomean                                 59.69n         36.31n        -39.18%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │            cidrtree/lookup.bm             │
                                    │     sec/op     │     sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4             62.66n ± 23%   1177.50n ±  24%  +1779.04% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             59.30n ± 31%   1174.00n ±  19%  +1879.76% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              77.69n ± 48%   1321.50n ±  23%  +1601.10% (p=0.002 n=10)
LpmTier1Pfxs/RandomMissIP6              16.77n ±  0%     67.13n ± 579%   +300.39% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     65.63n ± 33%   1083.50n ±  52%  +1550.92% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     66.71n ± 21%   1519.50n ±  27%  +2177.77% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      99.32n ± 22%   1126.50n ±  17%  +1034.21% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      76.59n ± 25%   1401.50n ±  31%  +1729.87% (p=0.000 n=10)
geomean                                 59.69n           866.8n         +1352.16%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │           critbitgo/lookup.bm           │
                                    │     sec/op     │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             62.66n ± 23%   352.65n ± 11%   +462.75% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             59.30n ± 31%   509.85n ± 18%   +759.78% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              77.69n ± 48%   696.15n ± 40%   +796.12% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              16.77n ±  0%   695.45n ± 14%  +4048.23% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     65.63n ± 33%   390.05n ± 39%   +494.32% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     66.71n ± 21%   408.45n ± 18%   +512.28% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      99.32n ± 22%   466.30n ± 39%   +369.49% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      76.59n ± 25%   454.15n ± 20%   +492.96% (p=0.000 n=10)
geomean                                 59.69n          482.5n         +708.29%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │           lpmtrie/lookup.bm            │
                                    │     sec/op     │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4             62.66n ± 23%   236.50n ± 16%  +277.40% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             59.30n ± 31%   264.50n ± 58%  +346.04% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              77.69n ± 48%   238.90n ± 66%  +207.52% (p=0.004 n=10)
LpmTier1Pfxs/RandomMissIP6              16.77n ±  0%   106.03n ± 21%  +532.48% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     65.63n ± 33%   259.30n ±  8%  +295.09% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     66.71n ± 21%   206.45n ±  9%  +209.47% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      99.32n ± 22%   291.45n ±  7%  +193.45% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      76.59n ± 25%   230.60n ±  9%  +201.08% (p=0.000 n=10)
geomean                                 59.69n          221.1n        +270.32%

```

## update, insert/delete

`bart` is by far the fastest algorithm for updates under all competitors.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │             art/update.bm             │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100            55.17n ± 1%   105.60n ± 1%   +91.39% (p=0.000 n=10)
Insert/Insert_into_1_000          54.61n ± 1%   125.40n ± 1%  +129.63% (p=0.000 n=10)
Insert/Insert_into_10_000         54.70n ± 1%   129.55n ± 2%  +136.86% (p=0.000 n=10)
Insert/Insert_into_100_000        54.95n ± 5%   121.75n ± 2%  +121.54% (p=0.000 n=10)
Insert/Insert_into_1_000_000      59.52n ± 2%   120.70n ± 1%  +102.79% (p=0.000 n=10)
Delete/Delete_from_100            17.59n ± 1%    34.32n ± 4%   +95.17% (p=0.000 n=10)
Delete/Delete_from_1_000          41.99n ± 1%    59.12n ± 0%   +40.80% (p=0.000 n=10)
Delete/Delete_from_10_000         43.55n ± 1%    59.95n ± 0%   +37.66% (p=0.000 n=10)
Delete/Delete_from_100_000        42.92n ± 1%    58.71n ± 0%   +36.76% (p=0.000 n=10)
Delete/Delete_from_1_000_000      40.76n ± 0%    58.99n ± 0%   +44.73% (p=0.000 n=10)
geomean                           44.48n         79.91n        +79.65%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │            cidrtree/update.bm            │
                             │     sec/op     │    sec/op      vs base                   │
Insert/Insert_into_100            55.17n ± 1%    949.75n ± 0%   +1621.34% (p=0.000 n=10)
Insert/Insert_into_1_000          54.61n ± 1%   1673.00n ± 0%   +2963.54% (p=0.000 n=10)
Insert/Insert_into_10_000         54.70n ± 1%   2025.50n ± 0%   +3603.26% (p=0.000 n=10)
Insert/Insert_into_100_000        54.95n ± 5%   2981.00n ± 0%   +5324.44% (p=0.000 n=10)
Insert/Insert_into_1_000_000      59.52n ± 2%   3189.00n ± 0%   +5257.86% (p=0.000 n=10)
Delete/Delete_from_100            17.59n ± 1%   1189.00n ± 0%   +6661.44% (p=0.000 n=10)
Delete/Delete_from_1_000          41.99n ± 1%   2112.50n ± 0%   +4930.96% (p=0.000 n=10)
Delete/Delete_from_10_000         43.55n ± 1%   4306.50n ± 0%   +9788.63% (p=0.000 n=10)
Delete/Delete_from_100_000        42.92n ± 1%   4466.50n ± 0%  +10305.36% (p=0.000 n=10)
Delete/Delete_from_1_000_000      40.76n ± 0%   4220.00n ± 0%  +10254.56% (p=0.000 n=10)
geomean                           44.48n          2.396µ        +5286.71%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │          critbitgo/update.bm          │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100            55.17n ± 1%   109.55n ± 0%   +98.55% (p=0.000 n=10)
Insert/Insert_into_1_000          54.61n ± 1%   122.00n ± 0%  +123.40% (p=0.000 n=10)
Insert/Insert_into_10_000         54.70n ± 1%   125.80n ± 0%  +130.00% (p=0.000 n=10)
Insert/Insert_into_100_000        54.95n ± 5%   154.60n ± 1%  +181.32% (p=0.000 n=10)
Insert/Insert_into_1_000_000      59.52n ± 2%   179.80n ± 0%  +202.08% (p=0.000 n=10)
Delete/Delete_from_100            17.59n ± 1%    99.64n ± 0%  +466.62% (p=0.000 n=10)
Delete/Delete_from_1_000          41.99n ± 1%   110.65n ± 0%  +163.52% (p=0.000 n=10)
Delete/Delete_from_10_000         43.55n ± 1%   112.40n ± 0%  +158.09% (p=0.000 n=10)
Delete/Delete_from_100_000        42.92n ± 1%   129.80n ± 1%  +202.39% (p=0.000 n=10)
Delete/Delete_from_1_000_000      40.76n ± 0%   126.65n ± 1%  +210.76% (p=0.000 n=10)
geomean                           44.48n         125.3n       +181.64%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │           lpmtrie/update.bm            │
                             │     sec/op     │    sec/op     vs base                  │
Insert/Insert_into_100            55.17n ± 1%   318.45n ± 1%   +477.16% (p=0.000 n=10)
Insert/Insert_into_1_000          54.61n ± 1%   348.75n ± 1%   +538.62% (p=0.000 n=10)
Insert/Insert_into_10_000         54.70n ± 1%   361.75n ± 1%   +561.40% (p=0.000 n=10)
Insert/Insert_into_100_000        54.95n ± 5%   502.35n ± 1%   +814.11% (p=0.000 n=10)
Insert/Insert_into_1_000_000      59.52n ± 2%   682.75n ± 1%  +1047.09% (p=0.000 n=10)
Delete/Delete_from_100            17.59n ± 1%    76.59n ± 0%   +335.51% (p=0.000 n=10)
Delete/Delete_from_1_000          41.99n ± 1%   125.65n ± 0%   +199.24% (p=0.000 n=10)
Delete/Delete_from_10_000         43.55n ± 1%   145.25n ± 0%   +233.52% (p=0.000 n=10)
Delete/Delete_from_100_000        42.92n ± 1%   261.60n ± 2%   +509.44% (p=0.000 n=10)
Delete/Delete_from_1_000_000      40.76n ± 0%   298.00n ± 0%   +631.20% (p=0.000 n=10)
geomean                           44.48n         261.6n        +488.13%
```
