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
Tier1PfxSize/100           31.52Ki ± 5%    802.98Ki ± 0%  +2447.89% (p=0.002 n=6)
Tier1PfxSize/1_000         204.2Ki ± 1%    7420.4Ki ± 0%  +3533.98% (p=0.002 n=6)
Tier1PfxSize/10_000        1.288Mi ± 0%    47.172Mi ± 0%  +3562.71% (p=0.002 n=6)
Tier1PfxSize/100_000       5.397Mi ± 0%   160.300Mi ± 0%  +2870.31% (p=0.002 n=6)
Tier1PfxSize/1_000_000     24.83Mi ± 0%    378.23Mi ± 0%  +1423.34% (p=0.002 n=6)
RandomPfx4Size/100         31.82Ki ± 5%    706.61Ki ± 0%  +2120.62% (p=0.002 n=6)
RandomPfx4Size/1_000       218.9Ki ± 1%    7299.3Ki ± 0%  +3234.91% (p=0.002 n=6)
RandomPfx4Size/10_000      1.725Mi ± 0%    58.097Mi ± 0%  +3268.15% (p=0.002 n=6)
RandomPfx4Size/100_000     13.17Mi ± 0%    524.14Mi ± 0%  +3880.21% (p=0.002 n=6)
RandomPfx4Size/1_000_000   79.85Mi ± 0%
RandomPfx6Size/100         99.71Ki ± 2%    725.73Ki ± 0%   +627.84% (p=0.002 n=6)
RandomPfx6Size/1_000       921.1Ki ± 0%    7656.3Ki ± 0%   +731.19% (p=0.002 n=6)
RandomPfx6Size/10_000      8.615Mi ± 0%    65.549Mi ± 0%   +660.92% (p=0.002 n=6)
RandomPfx6Size/100_000     81.53Mi ± 0%    747.63Mi ± 0%   +816.97% (p=0.002 n=6)
RandomPfx6Size/1_000_000   761.6Mi ± 0%
RandomPfxSize/100          43.80Ki ± 4%    694.27Ki ± 0%  +1485.19% (p=0.002 n=6)
RandomPfxSize/1_000        372.3Ki ± 0%    7420.4Ki ± 0%  +1893.06% (p=0.002 n=6)
RandomPfxSize/10_000       3.122Mi ± 0%    59.355Mi ± 0%  +1801.07% (p=0.002 n=6)
RandomPfxSize/100_000      27.66Mi ± 0%    553.96Mi ± 0%  +1902.52% (p=0.002 n=6)
RandomPfxSize/1_000_000    221.0Mi ± 0%
geomean                    2.585Mi          22.57Mi       +1838.90%               ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm  │           cidrtree/size.bm            │
                         │     bytes     │     bytes      vs base                │
Tier1PfxSize/100            31.52Ki ± 5%   12.48Ki ± 14%   -60.41% (p=0.002 n=6)
Tier1PfxSize/1_000         204.20Ki ± 1%   81.51Ki ±  2%   -60.08% (p=0.002 n=6)
Tier1PfxSize/10_000        1318.8Ki ± 0%   784.6Ki ±  0%   -40.50% (p=0.002 n=6)
Tier1PfxSize/100_000        5.397Mi ± 0%   7.633Mi ±  0%   +41.43% (p=0.002 n=6)
Tier1PfxSize/1_000_000      24.83Mi ± 0%   76.30Mi ±  0%  +207.29% (p=0.002 n=6)
RandomPfx4Size/100          31.82Ki ± 5%   11.58Ki ± 14%   -63.61% (p=0.002 n=6)
RandomPfx4Size/1_000       218.88Ki ± 1%   81.51Ki ±  2%   -62.76% (p=0.002 n=6)
RandomPfx4Size/10_000      1766.3Ki ± 0%   784.6Ki ±  0%   -55.58% (p=0.002 n=6)
RandomPfx4Size/100_000     13.169Mi ± 0%   7.633Mi ±  0%   -42.04% (p=0.002 n=6)
RandomPfx4Size/1_000_000    79.85Mi ± 0%   76.30Mi ±  0%    -4.45% (p=0.002 n=6)
RandomPfx6Size/100          99.71Ki ± 2%   11.58Ki ± 14%   -88.39% (p=0.002 n=6)
RandomPfx6Size/1_000       921.12Ki ± 0%   81.51Ki ±  2%   -91.15% (p=0.002 n=6)
RandomPfx6Size/10_000      8821.3Ki ± 0%   784.6Ki ±  0%   -91.11% (p=0.002 n=6)
RandomPfx6Size/100_000     81.532Mi ± 0%   7.633Mi ±  0%   -90.64% (p=0.002 n=6)
RandomPfx6Size/1_000_000   761.65Mi ± 0%   76.30Mi ±  0%   -89.98% (p=0.002 n=6)
RandomPfxSize/100           43.80Ki ± 4%   11.58Ki ± 14%   -73.56% (p=0.002 n=6)
RandomPfxSize/1_000        372.31Ki ± 0%   81.51Ki ±  2%   -78.11% (p=0.002 n=6)
RandomPfxSize/10_000       3197.1Ki ± 0%   784.6Ki ±  0%   -75.46% (p=0.002 n=6)
RandomPfxSize/100_000      27.663Mi ± 0%   7.633Mi ±  0%   -72.41% (p=0.002 n=6)
RandomPfxSize/1_000_000    220.96Mi ± 0%   76.30Mi ±  0%   -65.47% (p=0.002 n=6)
geomean                     2.585Mi        856.4Ki         -67.65%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │           critbitgo/size.bm            │
                         │    bytes     │     bytes       vs base                │
Tier1PfxSize/100           31.52Ki ± 5%    15.63Ki ± 11%   -50.40% (p=0.002 n=6)
Tier1PfxSize/1_000         204.2Ki ± 1%    115.2Ki ±  1%   -43.57% (p=0.002 n=6)
Tier1PfxSize/10_000        1.288Mi ± 0%    1.094Mi ±  0%   -15.09% (p=0.002 n=6)
Tier1PfxSize/100_000       5.397Mi ± 0%   10.913Mi ±  0%  +102.21% (p=0.002 n=6)
Tier1PfxSize/1_000_000     24.83Mi ± 0%   109.12Mi ±  0%  +339.47% (p=0.002 n=6)
RandomPfx4Size/100         31.82Ki ± 5%    14.67Ki ± 11%   -53.89% (p=0.002 n=6)
RandomPfx4Size/1_000       218.9Ki ± 1%    112.7Ki ±  1%   -48.50% (p=0.002 n=6)
RandomPfx4Size/10_000      1.725Mi ± 0%    1.071Mi ±  0%   -37.89% (p=0.002 n=6)
RandomPfx4Size/100_000     13.17Mi ± 0%    10.68Mi ±  0%   -18.86% (p=0.002 n=6)
RandomPfx4Size/1_000_000   79.85Mi ± 0%   106.81Mi ±  0%   +33.77% (p=0.002 n=6)
RandomPfx6Size/100         99.71Ki ± 2%    16.22Ki ± 10%   -83.73% (p=0.002 n=6)
RandomPfx6Size/1_000       921.1Ki ± 0%    128.3Ki ±  1%   -86.07% (p=0.002 n=6)
RandomPfx6Size/10_000      8.615Mi ± 0%    1.224Mi ±  0%   -85.79% (p=0.002 n=6)
RandomPfx6Size/100_000     81.53Mi ± 0%    12.21Mi ±  0%   -85.02% (p=0.002 n=6)
RandomPfx6Size/1_000_000   761.6Mi ± 0%    122.1Mi ±  0%   -83.97% (p=0.002 n=6)
RandomPfxSize/100          43.80Ki ± 4%    14.94Ki ± 11%   -65.89% (p=0.002 n=6)
RandomPfxSize/1_000        372.3Ki ± 0%    115.8Ki ±  1%   -68.90% (p=0.002 n=6)
RandomPfxSize/10_000       3.122Mi ± 0%    1.102Mi ±  0%   -64.70% (p=0.002 n=6)
RandomPfxSize/100_000      27.66Mi ± 0%    10.99Mi ±  0%   -60.27% (p=0.002 n=6)
RandomPfxSize/1_000_000    221.0Mi ± 0%    109.9Mi ±  0%   -50.28% (p=0.002 n=6)
geomean                    2.585Mi         1.193Mi         -53.84%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │            lpmtrie/size.bm             │
                         │    bytes     │     bytes       vs base                │
Tier1PfxSize/100           31.52Ki ± 5%    24.25Ki ±  9%   -23.05% (p=0.002 n=6)
Tier1PfxSize/1_000         204.2Ki ± 1%    202.3Ki ±  4%    -0.91% (p=0.002 n=6)
Tier1PfxSize/10_000        1.288Mi ± 0%    1.942Mi ±  3%   +50.79% (p=0.002 n=6)
Tier1PfxSize/100_000       5.397Mi ± 0%   19.330Mi ±  3%  +258.19% (p=0.002 n=6)
Tier1PfxSize/1_000_000     24.83Mi ± 0%   189.92Mi ±  6%  +664.90% (p=0.002 n=6)
RandomPfx4Size/100         31.82Ki ± 5%    23.20Ki ±  9%   -27.08% (p=0.002 n=6)
RandomPfx4Size/1_000       218.9Ki ± 1%    198.5Ki ±  3%    -9.31% (p=0.002 n=6)
RandomPfx4Size/10_000      1.725Mi ± 0%    1.898Mi ±  3%   +10.01% (p=0.002 n=6)
RandomPfx4Size/100_000     13.17Mi ± 0%    18.28Mi ±  5%   +38.79% (p=0.002 n=6)
RandomPfx4Size/1_000_000   79.85Mi ± 0%   163.51Mi ±  8%  +104.77% (p=0.002 n=6)
RandomPfx6Size/100         99.71Ki ± 2%    25.53Ki ± 12%   -74.39% (p=0.002 n=6)
RandomPfx6Size/1_000       921.1Ki ± 0%    222.0Ki ±  8%   -75.90% (p=0.002 n=6)
RandomPfx6Size/10_000      8.615Mi ± 0%    2.129Mi ±  7%   -75.29% (p=0.002 n=6)
RandomPfx6Size/100_000     81.53Mi ± 0%    20.93Mi ±  8%   -74.33% (p=0.002 n=6)
RandomPfx6Size/1_000_000   761.6Mi ± 0%    204.0Mi ±  8%   -73.22% (p=0.002 n=6)
RandomPfxSize/100          43.80Ki ± 4%    23.56Ki ± 10%   -46.20% (p=0.002 n=6)
RandomPfxSize/1_000        372.3Ki ± 0%    202.8Ki ±  4%   -45.54% (p=0.002 n=6)
RandomPfxSize/10_000       3.122Mi ± 0%    1.943Mi ±  4%   -37.78% (p=0.002 n=6)
RandomPfxSize/100_000      27.66Mi ± 0%    18.78Mi ±  5%   -32.10% (p=0.002 n=6)
RandomPfxSize/1_000_000    221.0Mi ± 0%    171.9Mi ±  7%   -22.19% (p=0.002 n=6)
geomean                    2.585Mi         2.011Mi         -22.21%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │           cidranger/size.bm            │
                         │    bytes     │     bytes      vs base                 │
Tier1PfxSize/100           31.52Ki ± 5%    56.03Ki ± 5%    +77.79% (p=0.002 n=6)
Tier1PfxSize/1_000         204.2Ki ± 1%    525.9Ki ± 3%   +157.55% (p=0.002 n=6)
Tier1PfxSize/10_000        1.288Mi ± 0%    5.091Mi ± 2%   +295.29% (p=0.002 n=6)
Tier1PfxSize/100_000       5.397Mi ± 0%   50.258Mi ± 2%   +831.27% (p=0.002 n=6)
Tier1PfxSize/1_000_000     24.83Mi ± 0%   477.45Mi ± 2%  +1822.94% (p=0.002 n=6)
RandomPfx4Size/100         31.82Ki ± 5%    54.16Ki ± 6%    +70.19% (p=0.002 n=6)
RandomPfx4Size/1_000       218.9Ki ± 1%    514.0Ki ± 3%   +134.82% (p=0.002 n=6)
RandomPfx4Size/10_000      1.725Mi ± 0%    4.900Mi ± 3%   +184.07% (p=0.002 n=6)
RandomPfx4Size/100_000     13.17Mi ± 0%    45.76Mi ± 3%   +247.47% (p=0.002 n=6)
RandomPfx4Size/1_000_000   79.85Mi ± 0%   396.20Mi ± 3%   +396.19% (p=0.002 n=6)
RandomPfx6Size/100         99.71Ki ± 2%    61.08Ki ± 3%    -38.74% (p=0.002 n=6)
RandomPfx6Size/1_000       921.1Ki ± 0%    579.3Ki ± 0%    -37.11% (p=0.002 n=6)
RandomPfx6Size/10_000      8.615Mi ± 0%    5.582Mi ± 0%    -35.21% (p=0.002 n=6)
RandomPfx6Size/100_000     81.53Mi ± 0%    54.75Mi ± 0%    -32.85% (p=0.002 n=6)
RandomPfx6Size/1_000_000   761.6Mi ± 0%    534.1Mi ± 0%    -29.87% (p=0.002 n=6)
RandomPfxSize/100          43.80Ki ± 4%    55.07Ki ± 5%    +25.74% (p=0.002 n=6)
RandomPfxSize/1_000        372.3Ki ± 0%    524.9Ki ± 2%    +40.97% (p=0.002 n=6)
RandomPfxSize/10_000       3.122Mi ± 0%    5.024Mi ± 2%    +60.93% (p=0.002 n=6)
RandomPfxSize/100_000      27.66Mi ± 0%    47.53Mi ± 2%    +71.83% (p=0.002 n=6)
RandomPfxSize/1_000_000    221.0Mi ± 0%    424.9Mi ± 2%    +92.29% (p=0.002 n=6)
geomean                    2.585Mi         5.068Mi         +96.04%
```

## lookup (longest-prefix-match)

In the lookup, `art` is the champion, closely followed by `bart`. 

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │            art/lookup.bm             │
                                    │     sec/op     │    sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4             51.75n ± 51%   46.51n ± 18%        ~ (p=0.436 n=10)
LpmTier1Pfxs/RandomMatchIP6             63.55n ± 15%   46.62n ± 11%  -26.65% (p=0.003 n=10)
LpmTier1Pfxs/RandomMissIP4              64.92n ± 48%   29.93n ± 11%  -53.90% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              17.31n ±  7%   30.36n ±  6%  +75.39% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     71.84n ± 46%   50.87n ± 10%  -29.19% (p=0.003 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     73.26n ± 27%   50.37n ± 11%  -31.24% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4     115.95n ± 23%   30.04n ±  5%  -74.09% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      83.31n ± 10%   29.89n ±  9%  -64.13% (p=0.000 n=10)
geomean                                 60.85n         38.20n        -37.22%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │            cidrtree/lookup.bm            │
                                    │     sec/op     │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             51.75n ± 51%    965.50n ± 56%  +1765.88% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             63.55n ± 15%    916.20n ± 16%  +1341.70% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              64.92n ± 48%   1182.00n ± 30%  +1720.84% (p=0.002 n=10)
LpmTier1Pfxs/RandomMissIP6              17.31n ±  7%    419.05n ± 90%  +2320.85% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     71.84n ± 46%   1248.00n ± 33%  +1637.19% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     73.26n ± 27%   1314.00n ± 41%  +1693.61% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      115.9n ± 23%    1187.5n ± 13%   +924.15% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      83.31n ± 10%   1112.50n ± 30%  +1235.29% (p=0.000 n=10)
geomean                                 60.85n           993.5n        +1532.82%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │           critbitgo/lookup.bm           │
                                    │     sec/op     │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             51.75n ± 51%   400.05n ± 19%   +673.12% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             63.55n ± 15%   527.20n ± 12%   +729.58% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              64.92n ± 48%   573.70n ± 58%   +783.77% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              17.31n ±  7%   478.60n ± 46%  +2664.88% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     71.84n ± 46%   465.95n ± 39%   +548.59% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     73.26n ± 27%   640.80n ± 18%   +774.69% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      115.9n ± 23%    506.6n ± 30%   +336.91% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      83.31n ± 10%   488.95n ± 24%   +486.87% (p=0.000 n=10)
geomean                                 60.85n          505.8n         +731.26%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │             lpmtrie/lookup.bm             │
                                    │     sec/op     │     sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4             51.75n ± 51%    229.55n ±   9%   +343.62% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             63.55n ± 15%    234.90n ±  48%   +269.63% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              64.92n ± 48%     89.68n ± 147%          ~ (p=0.089 n=10)
LpmTier1Pfxs/RandomMissIP6              17.31n ±  7%     79.46n ±   1%   +359.04% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     71.84n ± 46%   1047.00n ±   2%  +1357.41% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     73.26n ± 27%    187.95n ±  23%   +156.55% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      115.9n ± 23%    1071.5n ±   7%   +824.11% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      83.31n ± 10%    223.60n ±  19%   +168.38% (p=0.000 n=10)
geomean                                 60.85n           255.4n          +319.78%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │            cidranger/lookup.bm            │
                                    │     sec/op     │     sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4             51.75n ± 51%    359.20n ±  34%   +594.17% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             63.55n ± 15%    291.45n ±  47%   +358.62% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              64.92n ± 48%    218.90n ±  50%   +237.21% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              17.31n ±  7%    269.40n ±   8%  +1456.33% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     71.84n ± 46%   1586.00n ±  66%  +2107.68% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     73.26n ± 27%    510.10n ± 125%   +596.29% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      115.9n ± 23%     451.8n ±  13%   +289.69% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      83.31n ± 10%    308.00n ±   7%   +269.68% (p=0.000 n=10)
geomean                                 60.85n           403.0n          +562.25%
```

## update, insert/delete

`bart` is by far the fastest algorithm for updates under all competitors.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │             art/update.bm             │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100           47.09n ±  5%    88.30n ± 1%   +87.49% (p=0.000 n=10)
Insert/Insert_into_1_000         48.27n ± 26%   107.05n ± 3%  +121.77% (p=0.000 n=10)
Insert/Insert_into_10_000        47.06n ± 11%   107.35n ± 1%  +128.14% (p=0.000 n=10)
Insert/Insert_into_100_000       48.45n ± 14%   106.45n ± 1%  +119.73% (p=0.000 n=10)
Insert/Insert_into_1_000_000     48.88n ±  7%   105.30n ± 0%  +115.40% (p=0.000 n=10)
Delete/Delete_from_100           17.76n ±  6%    13.80n ± 0%   -22.30% (p=0.000 n=10)
Delete/Delete_from_1_000         41.25n ±  9%    40.15n ± 0%         ~ (p=0.470 n=10)
Delete/Delete_from_10_000        39.66n ±  5%    42.32n ± 0%    +6.71% (p=0.001 n=10)
Delete/Delete_from_100_000       41.02n ± 14%    40.17n ± 0%         ~ (p=0.138 n=10)
Delete/Delete_from_1_000_000     40.36n ±  5%    40.44n ± 0%         ~ (p=0.839 n=10)
geomean                          40.61n          58.03n        +42.92%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │            cidrtree/update.bm            │
                             │     sec/op     │    sec/op      vs base                   │
Insert/Insert_into_100           47.09n ±  5%    945.20n ± 0%   +1907.01% (p=0.000 n=10)
Insert/Insert_into_1_000         48.27n ± 26%   1452.00n ± 0%   +2908.08% (p=0.000 n=10)
Insert/Insert_into_10_000        47.06n ± 11%   2125.00n ± 0%   +4415.99% (p=0.000 n=10)
Insert/Insert_into_100_000       48.45n ± 14%   1595.00n ± 0%   +3192.39% (p=0.000 n=10)
Insert/Insert_into_1_000_000     48.88n ±  7%   2627.50n ± 0%   +5274.86% (p=0.000 n=10)
Delete/Delete_from_100           17.76n ±  6%   1502.00n ± 0%   +8357.21% (p=0.000 n=10)
Delete/Delete_from_1_000         41.25n ±  9%   1765.00n ± 0%   +4178.27% (p=0.000 n=10)
Delete/Delete_from_10_000        39.66n ±  5%   3444.00n ± 0%   +8584.91% (p=0.000 n=10)
Delete/Delete_from_100_000       41.02n ± 14%   3493.50n ± 0%   +8416.58% (p=0.000 n=10)
Delete/Delete_from_1_000_000     40.36n ±  5%   4819.00n ± 0%  +11840.04% (p=0.000 n=10)
geomean                          40.61n           2.125µ        +5133.60%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │          critbitgo/update.bm          │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100           47.09n ±  5%   109.95n ± 1%  +133.46% (p=0.000 n=10)
Insert/Insert_into_1_000         48.27n ± 26%   121.75n ± 1%  +152.23% (p=0.000 n=10)
Insert/Insert_into_10_000        47.06n ± 11%   125.50n ± 1%  +166.71% (p=0.000 n=10)
Insert/Insert_into_100_000       48.45n ± 14%   147.50n ± 1%  +204.47% (p=0.000 n=10)
Insert/Insert_into_1_000_000     48.88n ±  7%   152.20n ± 2%  +211.34% (p=0.000 n=10)
Delete/Delete_from_100           17.76n ±  6%    99.93n ± 1%  +462.67% (p=0.000 n=10)
Delete/Delete_from_1_000         41.25n ±  9%   108.40n ± 1%  +162.76% (p=0.000 n=10)
Delete/Delete_from_10_000        39.66n ±  5%   112.55n ± 1%  +183.82% (p=0.000 n=10)
Delete/Delete_from_100_000       41.02n ± 14%   131.40n ± 2%  +220.33% (p=0.000 n=10)
Delete/Delete_from_1_000_000     40.36n ±  5%   127.40n ± 2%  +215.66% (p=0.000 n=10)
geomean                          40.61n          122.6n       +202.03%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │           lpmtrie/update.bm            │
                             │     sec/op     │    sec/op     vs base                  │
Insert/Insert_into_100           47.09n ±  5%   314.55n ± 0%   +567.91% (p=0.000 n=10)
Insert/Insert_into_1_000         48.27n ± 26%   352.45n ± 2%   +630.16% (p=0.000 n=10)
Insert/Insert_into_10_000        47.06n ± 11%   362.60n ± 1%   +670.59% (p=0.000 n=10)
Insert/Insert_into_100_000       48.45n ± 14%   507.25n ± 1%   +947.06% (p=0.000 n=10)
Insert/Insert_into_1_000_000     48.88n ±  7%   656.80n ± 1%  +1243.56% (p=0.000 n=10)
Delete/Delete_from_100           17.76n ±  6%    77.03n ± 3%   +333.76% (p=0.000 n=10)
Delete/Delete_from_1_000         41.25n ±  9%   123.30n ± 0%   +198.87% (p=0.000 n=10)
Delete/Delete_from_10_000        39.66n ±  5%   149.50n ± 1%   +277.00% (p=0.000 n=10)
Delete/Delete_from_100_000       41.02n ± 14%   266.55n ± 0%   +549.80% (p=0.000 n=10)
Delete/Delete_from_1_000_000     40.36n ±  5%   309.70n ± 3%   +667.34% (p=0.000 n=10)
geomean                          40.61n          262.8n        +547.15%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │           cidranger/update.bm            │
                             │     sec/op     │    sec/op      vs base                   │
Insert/Insert_into_100           47.09n ±  5%   2442.50n ± 1%   +5086.33% (p=0.000 n=10)
Insert/Insert_into_1_000         48.27n ± 26%   4130.00n ± 2%   +8456.04% (p=0.000 n=10)
Insert/Insert_into_10_000        47.06n ± 11%   4849.00n ± 1%  +10204.96% (p=0.000 n=10)
Insert/Insert_into_100_000       48.45n ± 14%   8110.50n ± 1%  +16641.67% (p=0.000 n=10)
Insert/Insert_into_1_000_000     48.88n ±  7%   7579.00n ± 1%  +15403.73% (p=0.000 n=10)
Delete/Delete_from_100           17.76n ±  6%    385.05n ± 3%   +2068.07% (p=0.000 n=10)
Delete/Delete_from_1_000         41.25n ±  9%    421.45n ± 0%    +921.57% (p=0.000 n=10)
Delete/Delete_from_10_000        39.66n ±  5%    442.30n ± 1%   +1015.37% (p=0.000 n=10)
Delete/Delete_from_100_000       41.02n ± 14%    534.85n ± 8%   +1203.88% (p=0.000 n=10)
Delete/Delete_from_1_000_000     40.36n ±  5%    549.85n ± 0%   +1262.36% (p=0.000 n=10)
geomean                          40.61n           1.514µ        +3629.67%
```
