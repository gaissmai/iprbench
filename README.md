# iprbench
comparing benchmarks for some golang IP routing table implementations:

```
	github.com/tailscale/art
	github.com/gaissmai/bart
	github.com/k-sone/critbitgo
	github.com/Asphaltt/lpmtrie
	github.com/gaissmai/cidrtree
	github.com/yl2chen/cidranger
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

`bart` has a factor of ~15 lower memory consumption compared to `art`, but is by
a factor of ~2 slower in lookup times.

`cidrtree` is the most economical in terms of memory consumption,
but this is also not a trie but a binary search tree and
slower by a magnitude than the other algorithms.


```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │                art/size.bm                │
                         │    bytes     │     bytes       vs base                   │
Tier1PfxSize/100           34.78Ki ± 5%    802.97Ki ± 0%  +2208.63% (p=0.002 n=6)
Tier1PfxSize/1_000         225.6Ki ± 1%    7420.4Ki ± 0%  +3189.39% (p=0.002 n=6)
Tier1PfxSize/10_000        1.415Mi ± 0%    47.172Mi ± 0%  +3234.59% (p=0.002 n=6)
Tier1PfxSize/100_000       5.825Mi ± 0%   160.300Mi ± 0%  +2652.03% (p=0.002 n=6)
Tier1PfxSize/1_000_000     25.79Mi ± 0%    378.23Mi ± 0%  +1366.36% (p=0.002 n=6)
RandomPfx4Size/100         35.45Ki ± 5%    706.59Ki ± 0%  +1893.04% (p=0.002 n=6)
RandomPfx4Size/1_000       243.0Ki ± 1%    7286.5Ki ± 0%  +2898.76% (p=0.002 n=6)
RandomPfx4Size/10_000      1.923Mi ± 0%    57.886Mi ± 0%  +2909.68% (p=0.002 n=6)
RandomPfx4Size/100_000     14.62Mi ± 0%    523.23Mi ± 0%  +3478.70% (p=0.002 n=6)
RandomPfx4Size/1_000_000   87.81Mi ± 0%
RandomPfx6Size/100         114.7Ki ± 1%     700.2Ki ± 0%   +510.45% (p=0.002 n=6)
RandomPfx6Size/1_000       1.018Mi ± 0%     7.477Mi ± 0%   +634.36% (p=0.002 n=6)
RandomPfx6Size/10_000      9.720Mi ± 0%    65.512Mi ± 0%   +574.01% (p=0.002 n=6)
RandomPfx6Size/100_000     92.11Mi ± 0%    748.96Mi ± 0%   +713.08% (p=0.002 n=6)
RandomPfx6Size/1_000_000   857.9Mi ± 0%
RandomPfxSize/100          49.43Ki ± 3%    694.25Ki ± 0%  +1304.52% (p=0.002 n=6)
RandomPfxSize/1_000        422.2Ki ± 0%    7439.5Ki ± 0%  +1662.01% (p=0.002 n=6)
RandomPfxSize/10_000       3.490Mi ± 0%    59.641Mi ± 0%  +1608.78% (p=0.002 n=6)
RandomPfxSize/100_000      31.14Mi ± 0%    553.83Mi ± 0%  +1678.63% (p=0.002 n=6)
RandomPfxSize/1_000_000    247.1Mi ± 0%
geomean                    2.877Mi          22.52Mi       +1638.93%               ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │  bart/size.bm  │           cidrtree/size.bm            │
                         │     bytes      │     bytes      vs base                │
Tier1PfxSize/100             34.78Ki ± 5%   12.48Ki ± 14%   -64.13% (p=0.002 n=6)
Tier1PfxSize/1_000          225.59Ki ± 1%   81.51Ki ±  2%   -63.87% (p=0.002 n=6)
Tier1PfxSize/10_000         1448.6Ki ± 0%   784.6Ki ±  0%   -45.83% (p=0.002 n=6)
Tier1PfxSize/100_000         5.825Mi ± 0%   7.633Mi ±  0%   +31.04% (p=0.002 n=6)
Tier1PfxSize/1_000_000       25.79Mi ± 0%   76.30Mi ±  0%  +195.80% (p=0.002 n=6)
RandomPfx4Size/100           35.45Ki ± 5%   11.58Ki ± 14%   -67.34% (p=0.002 n=6)
RandomPfx4Size/1_000        242.98Ki ± 1%   81.51Ki ±  2%   -66.46% (p=0.002 n=6)
RandomPfx4Size/10_000       1969.5Ki ± 0%   784.6Ki ±  0%   -60.16% (p=0.002 n=6)
RandomPfx4Size/100_000      14.621Mi ± 0%   7.633Mi ±  0%   -47.79% (p=0.002 n=6)
RandomPfx4Size/1_000_000     87.81Mi ± 0%   76.30Mi ±  0%   -13.11% (p=0.002 n=6)
RandomPfx6Size/100          114.70Ki ± 1%   11.58Ki ± 14%   -89.91% (p=0.002 n=6)
RandomPfx6Size/1_000       1042.57Ki ± 0%   81.51Ki ±  2%   -92.18% (p=0.002 n=6)
RandomPfx6Size/10_000       9953.1Ki ± 0%   784.6Ki ±  0%   -92.12% (p=0.002 n=6)
RandomPfx6Size/100_000      92.114Mi ± 0%   7.633Mi ±  0%   -91.71% (p=0.002 n=6)
RandomPfx6Size/1_000_000    857.93Mi ± 0%   76.30Mi ±  0%   -91.11% (p=0.002 n=6)
RandomPfxSize/100            49.43Ki ± 3%   11.58Ki ± 14%   -76.58% (p=0.002 n=6)
RandomPfxSize/1_000         422.22Ki ± 0%   81.51Ki ±  2%   -80.70% (p=0.002 n=6)
RandomPfxSize/10_000        3574.1Ki ± 0%   784.6Ki ±  0%   -78.05% (p=0.002 n=6)
RandomPfxSize/100_000       31.138Mi ± 0%   7.633Mi ±  0%   -75.49% (p=0.002 n=6)
RandomPfxSize/1_000_000     247.10Mi ± 0%   76.30Mi ±  0%   -69.12% (p=0.002 n=6)
geomean                      2.877Mi        856.4Ki         -70.93%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm  │           critbitgo/size.bm            │
                         │     bytes     │     bytes       vs base                │
Tier1PfxSize/100            34.78Ki ± 5%    15.63Ki ± 11%   -55.05% (p=0.002 n=6)
Tier1PfxSize/1_000          225.6Ki ± 1%    115.2Ki ±  1%   -48.92% (p=0.002 n=6)
Tier1PfxSize/10_000         1.415Mi ± 0%    1.094Mi ±  0%   -22.70% (p=0.002 n=6)
Tier1PfxSize/100_000        5.825Mi ± 0%   10.913Mi ±  0%   +87.35% (p=0.002 n=6)
Tier1PfxSize/1_000_000      25.79Mi ± 0%   109.12Mi ±  0%  +323.03% (p=0.002 n=6)
RandomPfx4Size/100          35.45Ki ± 5%    14.67Ki ± 11%   -58.62% (p=0.002 n=6)
RandomPfx4Size/1_000        243.0Ki ± 1%    112.7Ki ±  1%   -53.61% (p=0.002 n=6)
RandomPfx4Size/10_000       1.923Mi ± 0%    1.071Mi ±  0%   -44.29% (p=0.002 n=6)
RandomPfx4Size/100_000      14.62Mi ± 0%    10.68Mi ±  0%   -26.92% (p=0.002 n=6)
RandomPfx4Size/1_000_000    87.81Mi ± 0%   106.81Mi ±  0%   +21.65% (p=0.002 n=6)
RandomPfx6Size/100         114.70Ki ± 1%    16.22Ki ± 10%   -85.86% (p=0.002 n=6)
RandomPfx6Size/1_000       1042.6Ki ± 0%    128.3Ki ±  1%   -87.69% (p=0.002 n=6)
RandomPfx6Size/10_000       9.720Mi ± 0%    1.224Mi ±  0%   -87.41% (p=0.002 n=6)
RandomPfx6Size/100_000      92.11Mi ± 0%    12.21Mi ±  0%   -86.74% (p=0.002 n=6)
RandomPfx6Size/1_000_000    857.9Mi ± 0%    122.1Mi ±  0%   -85.77% (p=0.002 n=6)
RandomPfxSize/100           49.43Ki ± 3%    14.94Ki ± 11%   -69.78% (p=0.002 n=6)
RandomPfxSize/1_000         422.2Ki ± 0%    115.8Ki ±  1%   -72.58% (p=0.002 n=6)
RandomPfxSize/10_000        3.490Mi ± 0%    1.102Mi ±  0%   -68.43% (p=0.002 n=6)
RandomPfxSize/100_000       31.14Mi ± 0%    10.99Mi ±  0%   -64.71% (p=0.002 n=6)
RandomPfxSize/1_000_000     247.1Mi ± 0%    109.9Mi ±  0%   -55.54% (p=0.002 n=6)
geomean                     2.877Mi         1.193Mi         -58.52%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm  │            lpmtrie/size.bm             │
                         │     bytes     │     bytes       vs base                │
Tier1PfxSize/100            34.78Ki ± 5%    24.25Ki ±  9%   -30.28% (p=0.002 n=6)
Tier1PfxSize/1_000          225.6Ki ± 1%    202.3Ki ±  4%   -10.31% (p=0.002 n=6)
Tier1PfxSize/10_000         1.415Mi ± 0%    1.942Mi ±  3%   +37.28% (p=0.002 n=6)
Tier1PfxSize/100_000        5.825Mi ± 0%   19.330Mi ±  3%  +231.86% (p=0.002 n=6)
Tier1PfxSize/1_000_000      25.79Mi ± 0%   189.92Mi ±  6%  +636.29% (p=0.002 n=6)
RandomPfx4Size/100          35.45Ki ± 5%    23.20Ki ±  9%   -34.55% (p=0.002 n=6)
RandomPfx4Size/1_000        243.0Ki ± 1%    198.5Ki ±  3%   -18.30% (p=0.002 n=6)
RandomPfx4Size/10_000       1.923Mi ± 0%    1.898Mi ±  3%    -1.34% (p=0.002 n=6)
RandomPfx4Size/100_000      14.62Mi ± 0%    18.28Mi ±  5%   +25.00% (p=0.002 n=6)
RandomPfx4Size/1_000_000    87.81Mi ± 0%   163.51Mi ±  8%   +86.21% (p=0.002 n=6)
RandomPfx6Size/100         114.70Ki ± 1%    25.53Ki ± 12%   -77.74% (p=0.002 n=6)
RandomPfx6Size/1_000       1042.6Ki ± 0%    222.0Ki ±  8%   -78.71% (p=0.002 n=6)
RandomPfx6Size/10_000       9.720Mi ± 0%    2.129Mi ±  7%   -78.10% (p=0.002 n=6)
RandomPfx6Size/100_000      92.11Mi ± 0%    20.93Mi ±  8%   -77.28% (p=0.002 n=6)
RandomPfx6Size/1_000_000    857.9Mi ± 0%    204.0Mi ±  8%   -76.23% (p=0.002 n=6)
RandomPfxSize/100           49.43Ki ± 3%    23.56Ki ± 10%   -52.33% (p=0.002 n=6)
RandomPfxSize/1_000         422.2Ki ± 0%    202.8Ki ±  4%   -51.97% (p=0.002 n=6)
RandomPfxSize/10_000        3.490Mi ± 0%    1.943Mi ±  4%   -44.34% (p=0.002 n=6)
RandomPfxSize/100_000       31.14Mi ± 0%    18.78Mi ±  5%   -39.68% (p=0.002 n=6)
RandomPfxSize/1_000_000     247.1Mi ± 0%    171.9Mi ±  7%   -30.42% (p=0.002 n=6)
geomean                     2.877Mi         2.011Mi         -30.11%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm  │           cidranger/size.bm            │
                         │     bytes     │     bytes      vs base                 │
Tier1PfxSize/100            34.78Ki ± 5%    56.03Ki ± 5%    +61.10% (p=0.002 n=6)
Tier1PfxSize/1_000          225.6Ki ± 1%    525.9Ki ± 3%   +133.13% (p=0.002 n=6)
Tier1PfxSize/10_000         1.415Mi ± 0%    5.091Mi ± 2%   +259.88% (p=0.002 n=6)
Tier1PfxSize/100_000        5.825Mi ± 0%   50.258Mi ± 2%   +762.84% (p=0.002 n=6)
Tier1PfxSize/1_000_000      25.79Mi ± 0%   477.45Mi ± 2%  +1751.01% (p=0.002 n=6)
RandomPfx4Size/100          35.45Ki ± 5%    54.16Ki ± 6%    +52.75% (p=0.002 n=6)
RandomPfx4Size/1_000        243.0Ki ± 1%    514.0Ki ± 3%   +111.52% (p=0.002 n=6)
RandomPfx4Size/10_000       1.923Mi ± 0%    4.900Mi ± 3%   +154.77% (p=0.002 n=6)
RandomPfx4Size/100_000      14.62Mi ± 0%    45.76Mi ± 3%   +212.96% (p=0.002 n=6)
RandomPfx4Size/1_000_000    87.81Mi ± 0%   396.20Mi ± 3%   +351.22% (p=0.002 n=6)
RandomPfx6Size/100         114.70Ki ± 1%    61.08Ki ± 3%    -46.75% (p=0.002 n=6)
RandomPfx6Size/1_000       1042.6Ki ± 0%    579.3Ki ± 0%    -44.44% (p=0.002 n=6)
RandomPfx6Size/10_000       9.720Mi ± 0%    5.582Mi ± 0%    -42.57% (p=0.002 n=6)
RandomPfx6Size/100_000      92.11Mi ± 0%    54.75Mi ± 0%    -40.56% (p=0.002 n=6)
RandomPfx6Size/1_000_000    857.9Mi ± 0%    534.1Mi ± 0%    -37.74% (p=0.002 n=6)
RandomPfxSize/100           49.43Ki ± 3%    55.07Ki ± 5%    +11.41% (p=0.002 n=6)
RandomPfxSize/1_000         422.2Ki ± 0%    524.9Ki ± 2%    +24.31% (p=0.002 n=6)
RandomPfxSize/10_000        3.490Mi ± 0%    5.024Mi ± 2%    +43.96% (p=0.002 n=6)
RandomPfxSize/100_000       31.14Mi ± 0%    47.53Mi ± 2%    +52.66% (p=0.002 n=6)
RandomPfxSize/1_000_000     247.1Mi ± 0%    424.9Mi ± 2%    +71.95% (p=0.002 n=6)
geomean                     2.877Mi         5.068Mi         +76.13%
```

## lookup (longest-prefix-match)

In the lookup, `art` is the champion, closely followed by `bart`. 

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │            art/lookup.bm             │
                                    │     sec/op     │    sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4             55.25n ± 49%   49.98n ± 10%        ~ (p=0.481 n=10)
LpmTier1Pfxs/RandomMatchIP6             67.78n ± 41%   46.56n ± 11%  -31.31% (p=0.002 n=10)
LpmTier1Pfxs/RandomMissIP4              80.12n ± 49%   29.93n ±  1%  -62.65% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              19.43n ±  3%   29.87n ±  1%  +53.77% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     76.14n ± 38%   50.49n ± 11%  -33.68% (p=0.023 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     69.18n ± 23%   47.53n ± 19%  -31.29% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4     106.29n ± 20%   30.67n ±  7%  -71.14% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      80.64n ± 14%   30.50n ±  6%  -62.18% (p=0.000 n=10)
geomean                                 63.46n         38.34n        -39.59%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │            cidrtree/lookup.bm            │
                                    │     sec/op     │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             55.25n ± 49%    965.50n ± 56%  +1647.51% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             67.78n ± 41%    916.20n ± 16%  +1251.73% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              80.12n ± 49%   1182.00n ± 30%  +1375.29% (p=0.002 n=10)
LpmTier1Pfxs/RandomMissIP6              19.43n ±  3%    419.05n ± 90%  +2057.27% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     76.14n ± 38%   1248.00n ± 33%  +1539.19% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     69.18n ± 23%   1314.00n ± 41%  +1799.39% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      106.3n ± 20%    1187.5n ± 13%  +1017.23% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      80.64n ± 14%   1112.50n ± 30%  +1279.50% (p=0.000 n=10)
geomean                                 63.46n           993.5n        +1465.49%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │           critbitgo/lookup.bm           │
                                    │     sec/op     │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             55.25n ± 49%   400.05n ± 19%   +624.07% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             67.78n ± 41%   527.20n ± 12%   +677.81% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              80.12n ± 49%   573.70n ± 58%   +616.05% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              19.43n ±  3%   478.60n ± 46%  +2363.84% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     76.14n ± 38%   465.95n ± 39%   +512.00% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     69.18n ± 23%   640.80n ± 18%   +826.28% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      106.3n ± 20%    506.6n ± 30%   +376.62% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      80.64n ± 14%   488.95n ± 24%   +506.30% (p=0.000 n=10)
geomean                                 63.46n          505.8n         +696.98%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │             lpmtrie/lookup.bm             │
                                    │     sec/op     │     sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4             55.25n ± 49%    229.55n ±   9%   +315.48% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             67.78n ± 41%    234.90n ±  48%   +246.56% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              80.12n ± 49%     89.68n ± 147%          ~ (p=0.247 n=10)
LpmTier1Pfxs/RandomMissIP6              19.43n ±  3%     79.46n ±   1%   +309.06% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     76.14n ± 38%   1047.00n ±   2%  +1275.19% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     69.18n ± 23%    187.95n ±  23%   +171.68% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      106.3n ± 20%    1071.5n ±   7%   +908.09% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      80.64n ± 14%    223.60n ±  19%   +177.26% (p=0.000 n=10)
geomean                                 63.46n           255.4n          +302.47%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │            cidranger/lookup.bm            │
                                    │     sec/op     │     sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4             55.25n ± 49%    359.20n ±  34%   +550.14% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             67.78n ± 41%    291.45n ±  47%   +329.99% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              80.12n ± 49%    218.90n ±  50%   +173.22% (p=0.001 n=10)
LpmTier1Pfxs/RandomMissIP6              19.43n ±  3%    269.40n ±   8%  +1286.87% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     76.14n ± 38%   1586.00n ±  66%  +1983.14% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     69.18n ± 23%    510.10n ± 125%   +637.35% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      106.3n ± 20%     451.8n ±  13%   +325.11% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      80.64n ± 14%    308.00n ±   7%   +281.92% (p=0.000 n=10)
geomean                                 63.46n           403.0n          +534.94%
```

## update, insert/delete

`bart` is by far the fastest algorithm for updates under all competitors.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │             art/update.bm             │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100           48.12n ±  1%   105.85n ± 1%  +119.97% (p=0.000 n=10)
Insert/Insert_into_1_000         48.27n ± 13%   125.45n ± 1%  +159.89% (p=0.000 n=10)
Insert/Insert_into_10_000        47.95n ±  1%   129.05n ± 1%  +169.11% (p=0.000 n=10)
Insert/Insert_into_100_000       55.85n ±  0%   120.55n ± 2%  +115.87% (p=0.000 n=10)
Insert/Insert_into_1_000_000     49.80n ±  1%   121.00n ± 0%  +143.00% (p=0.000 n=10)
Delete/Delete_from_100           16.91n ±  0%    34.04n ± 3%  +101.30% (p=0.000 n=10)
Delete/Delete_from_1_000         39.03n ±  0%    59.06n ± 0%   +51.34% (p=0.000 n=10)
Delete/Delete_from_10_000        38.92n ±  0%    58.93n ± 0%   +51.42% (p=0.000 n=10)
Delete/Delete_from_100_000       40.14n ±  1%    58.71n ± 0%   +46.26% (p=0.000 n=10)
Delete/Delete_from_1_000_000     39.19n ±  1%    58.50n ± 0%   +49.30% (p=0.000 n=10)
geomean                          40.71n          79.57n        +95.44%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │            cidrtree/update.bm            │
                             │     sec/op     │    sec/op      vs base                   │
Insert/Insert_into_100           48.12n ±  1%    945.20n ± 0%   +1864.26% (p=0.000 n=10)
Insert/Insert_into_1_000         48.27n ± 13%   1452.00n ± 0%   +2908.08% (p=0.000 n=10)
Insert/Insert_into_10_000        47.95n ±  1%   2125.00n ± 0%   +4331.24% (p=0.000 n=10)
Insert/Insert_into_100_000       55.85n ±  0%   1595.00n ± 0%   +2756.12% (p=0.000 n=10)
Insert/Insert_into_1_000_000     49.80n ±  1%   2627.50n ± 0%   +5176.63% (p=0.000 n=10)
Delete/Delete_from_100           16.91n ±  0%   1502.00n ± 0%   +8782.32% (p=0.000 n=10)
Delete/Delete_from_1_000         39.03n ±  0%   1765.00n ± 0%   +4422.74% (p=0.000 n=10)
Delete/Delete_from_10_000        38.92n ±  0%   3444.00n ± 0%   +8750.06% (p=0.000 n=10)
Delete/Delete_from_100_000       40.14n ±  1%   3493.50n ± 0%   +8603.29% (p=0.000 n=10)
Delete/Delete_from_1_000_000     39.19n ±  1%   4819.00n ± 0%  +12198.07% (p=0.000 n=10)
geomean                          40.71n           2.125µ        +5119.80%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │          critbitgo/update.bm          │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100           48.12n ±  1%   109.95n ± 1%  +128.49% (p=0.000 n=10)
Insert/Insert_into_1_000         48.27n ± 13%   121.75n ± 1%  +152.23% (p=0.000 n=10)
Insert/Insert_into_10_000        47.95n ±  1%   125.50n ± 1%  +161.70% (p=0.000 n=10)
Insert/Insert_into_100_000       55.85n ±  0%   147.50n ± 1%  +164.12% (p=0.000 n=10)
Insert/Insert_into_1_000_000     49.80n ±  1%   152.20n ± 2%  +205.65% (p=0.000 n=10)
Delete/Delete_from_100           16.91n ±  0%    99.93n ± 1%  +490.95% (p=0.000 n=10)
Delete/Delete_from_1_000         39.03n ±  0%   108.40n ± 1%  +177.77% (p=0.000 n=10)
Delete/Delete_from_10_000        38.92n ±  0%   112.55n ± 1%  +189.22% (p=0.000 n=10)
Delete/Delete_from_100_000       40.14n ±  1%   131.40n ± 2%  +227.35% (p=0.000 n=10)
Delete/Delete_from_1_000_000     39.19n ±  1%   127.40n ± 2%  +225.12% (p=0.000 n=10)
geomean                          40.71n          122.6n       +201.23%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │           lpmtrie/update.bm            │
                             │     sec/op     │    sec/op     vs base                  │
Insert/Insert_into_100           48.12n ±  1%   314.55n ± 0%   +553.68% (p=0.000 n=10)
Insert/Insert_into_1_000         48.27n ± 13%   352.45n ± 2%   +630.16% (p=0.000 n=10)
Insert/Insert_into_10_000        47.95n ±  1%   362.60n ± 1%   +656.13% (p=0.000 n=10)
Insert/Insert_into_100_000       55.85n ±  0%   507.25n ± 1%   +808.32% (p=0.000 n=10)
Insert/Insert_into_1_000_000     49.80n ±  1%   656.80n ± 1%  +1219.01% (p=0.000 n=10)
Delete/Delete_from_100           16.91n ±  0%    77.03n ± 3%   +355.56% (p=0.000 n=10)
Delete/Delete_from_1_000         39.03n ±  0%   123.30n ± 0%   +215.95% (p=0.000 n=10)
Delete/Delete_from_10_000        38.92n ±  0%   149.50n ± 1%   +284.17% (p=0.000 n=10)
Delete/Delete_from_100_000       40.14n ±  1%   266.55n ± 0%   +564.05% (p=0.000 n=10)
Delete/Delete_from_1_000_000     39.19n ±  1%   309.70n ± 3%   +690.35% (p=0.000 n=10)
geomean                          40.71n          262.8n        +545.44%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │           cidranger/update.bm            │
                             │     sec/op     │    sec/op      vs base                   │
Insert/Insert_into_100           48.12n ±  1%   2442.50n ± 1%   +4975.85% (p=0.000 n=10)
Insert/Insert_into_1_000         48.27n ± 13%   4130.00n ± 2%   +8456.04% (p=0.000 n=10)
Insert/Insert_into_10_000        47.95n ±  1%   4849.00n ± 1%  +10011.56% (p=0.000 n=10)
Insert/Insert_into_100_000       55.85n ±  0%   8110.50n ± 1%  +14423.23% (p=0.000 n=10)
Insert/Insert_into_1_000_000     49.80n ±  1%   7579.00n ± 1%  +15120.40% (p=0.000 n=10)
Delete/Delete_from_100           16.91n ±  0%    385.05n ± 3%   +2177.05% (p=0.000 n=10)
Delete/Delete_from_1_000         39.03n ±  0%    421.45n ± 0%    +979.95% (p=0.000 n=10)
Delete/Delete_from_10_000        38.92n ±  0%    442.30n ± 1%   +1036.58% (p=0.000 n=10)
Delete/Delete_from_100_000       40.14n ±  1%    534.85n ± 8%   +1232.46% (p=0.000 n=10)
Delete/Delete_from_1_000_000     39.19n ±  1%    549.85n ± 0%   +1303.22% (p=0.000 n=10)
geomean                          40.71n           1.514µ        +3619.84%
```
