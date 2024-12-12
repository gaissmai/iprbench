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
Tier1PfxSize/1_000         228.2Ki ± 1%    7420.4Ki ± 0%  +3151.00% (p=0.002 n=6)
Tier1PfxSize/10_000        1.445Mi ± 0%    47.172Mi ± 0%  +3164.47% (p=0.002 n=6)
Tier1PfxSize/100_000       6.339Mi ± 0%   160.300Mi ± 0%  +2428.91% (p=0.002 n=6)
Tier1PfxSize/1_000_000     31.67Mi ± 0%    378.23Mi ± 0%  +1094.45% (p=0.002 n=6)
RandomPfx4Size/100         37.73Ki ± 4%    706.59Ki ± 0%  +1772.93% (p=0.002 n=6)
RandomPfx4Size/1_000       247.2Ki ± 1%    7286.5Ki ± 0%  +2848.14% (p=0.002 n=6)
RandomPfx4Size/10_000      1.980Mi ± 0%    57.886Mi ± 0%  +2824.23% (p=0.002 n=6)
RandomPfx4Size/100_000     14.98Mi ± 0%    523.23Mi ± 0%  +3393.58% (p=0.002 n=6)
RandomPfx4Size/1_000_000   91.85Mi ± 0%
RandomPfx6Size/100         117.4Ki ± 1%     700.2Ki ± 0%   +496.63% (p=0.002 n=6)
RandomPfx6Size/1_000       1.046Mi ± 0%     7.477Mi ± 0%   +614.60% (p=0.002 n=6)
RandomPfx6Size/10_000      9.807Mi ± 0%    65.512Mi ± 0%   +568.00% (p=0.002 n=6)
RandomPfx6Size/100_000     92.24Mi ± 0%    748.96Mi ± 0%   +711.95% (p=0.002 n=6)
RandomPfx6Size/1_000_000   860.4Mi ± 0%
RandomPfxSize/100          51.11Ki ± 3%    694.25Ki ± 0%  +1258.36% (p=0.002 n=6)
RandomPfxSize/1_000        405.4Ki ± 0%    7439.5Ki ± 0%  +1735.33% (p=0.002 n=6)
RandomPfxSize/10_000       3.540Mi ± 0%    59.641Mi ± 0%  +1584.85% (p=0.002 n=6)
RandomPfxSize/100_000      31.48Mi ± 0%    553.83Mi ± 0%  +1659.07% (p=0.002 n=6)
RandomPfxSize/1_000_000    250.6Mi ± 0%
geomean                    2.964Mi          22.52Mi       +1585.25%               ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │  bart/size.bm  │           cidrtree/size.bm            │
                         │     bytes      │     bytes      vs base                │
Tier1PfxSize/100             34.78Ki ± 5%   12.48Ki ± 14%   -64.13% (p=0.002 n=6)
Tier1PfxSize/1_000          228.25Ki ± 1%   81.51Ki ±  2%   -64.29% (p=0.002 n=6)
Tier1PfxSize/10_000         1479.7Ki ± 0%   784.6Ki ±  0%   -46.97% (p=0.002 n=6)
Tier1PfxSize/100_000         6.339Mi ± 0%   7.633Mi ±  0%   +20.41% (p=0.002 n=6)
Tier1PfxSize/1_000_000       31.67Mi ± 0%   76.30Mi ±  0%  +140.95% (p=0.002 n=6)
RandomPfx4Size/100           37.73Ki ± 4%   11.58Ki ± 14%   -69.31% (p=0.002 n=6)
RandomPfx4Size/1_000        247.16Ki ± 1%   81.51Ki ±  2%   -67.02% (p=0.002 n=6)
RandomPfx4Size/10_000       2027.0Ki ± 0%   784.6Ki ±  0%   -61.29% (p=0.002 n=6)
RandomPfx4Size/100_000      14.977Mi ± 0%   7.633Mi ±  0%   -49.04% (p=0.002 n=6)
RandomPfx4Size/1_000_000     91.85Mi ± 0%   76.30Mi ±  0%   -16.94% (p=0.002 n=6)
RandomPfx6Size/100          117.36Ki ± 1%   11.58Ki ± 14%   -90.13% (p=0.002 n=6)
RandomPfx6Size/1_000       1071.40Ki ± 0%   81.51Ki ±  2%   -92.39% (p=0.002 n=6)
RandomPfx6Size/10_000      10042.5Ki ± 0%   784.6Ki ±  0%   -92.19% (p=0.002 n=6)
RandomPfx6Size/100_000      92.242Mi ± 0%   7.633Mi ±  0%   -91.72% (p=0.002 n=6)
RandomPfx6Size/1_000_000    860.38Mi ± 0%   76.30Mi ±  0%   -91.13% (p=0.002 n=6)
RandomPfxSize/100            51.11Ki ± 3%   11.58Ki ± 14%   -77.35% (p=0.002 n=6)
RandomPfxSize/1_000         405.35Ki ± 0%   81.51Ki ±  2%   -79.89% (p=0.002 n=6)
RandomPfxSize/10_000        3624.8Ki ± 0%   784.6Ki ±  0%   -78.35% (p=0.002 n=6)
RandomPfxSize/100_000       31.484Mi ± 0%   7.633Mi ±  0%   -75.76% (p=0.002 n=6)
RandomPfxSize/1_000_000     250.61Mi ± 0%   76.30Mi ±  0%   -69.56% (p=0.002 n=6)
geomean                      2.964Mi        856.4Ki         -71.78%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm  │           critbitgo/size.bm            │
                         │     bytes     │     bytes       vs base                │
Tier1PfxSize/100            34.78Ki ± 5%    15.63Ki ± 11%   -55.05% (p=0.002 n=6)
Tier1PfxSize/1_000          228.2Ki ± 1%    115.2Ki ±  1%   -49.52% (p=0.002 n=6)
Tier1PfxSize/10_000         1.445Mi ± 0%    1.094Mi ±  0%   -24.32% (p=0.002 n=6)
Tier1PfxSize/100_000        6.339Mi ± 0%   10.913Mi ±  0%   +72.16% (p=0.002 n=6)
Tier1PfxSize/1_000_000      31.67Mi ± 0%   109.12Mi ±  0%  +244.59% (p=0.002 n=6)
RandomPfx4Size/100          37.73Ki ± 4%    14.67Ki ± 11%   -61.11% (p=0.002 n=6)
RandomPfx4Size/1_000        247.2Ki ± 1%    112.7Ki ±  1%   -54.39% (p=0.002 n=6)
RandomPfx4Size/10_000       1.980Mi ± 0%    1.071Mi ±  0%   -45.88% (p=0.002 n=6)
RandomPfx4Size/100_000      14.98Mi ± 0%    10.68Mi ±  0%   -28.66% (p=0.002 n=6)
RandomPfx4Size/1_000_000    91.85Mi ± 0%   106.81Mi ±  0%   +16.29% (p=0.002 n=6)
RandomPfx6Size/100         117.36Ki ± 1%    16.22Ki ± 10%   -86.18% (p=0.002 n=6)
RandomPfx6Size/1_000       1071.4Ki ± 0%    128.3Ki ±  1%   -88.02% (p=0.002 n=6)
RandomPfx6Size/10_000       9.807Mi ± 0%    1.224Mi ±  0%   -87.52% (p=0.002 n=6)
RandomPfx6Size/100_000      92.24Mi ± 0%    12.21Mi ±  0%   -86.76% (p=0.002 n=6)
RandomPfx6Size/1_000_000    860.4Mi ± 0%    122.1Mi ±  0%   -85.81% (p=0.002 n=6)
RandomPfxSize/100           51.11Ki ± 3%    14.94Ki ± 11%   -70.77% (p=0.002 n=6)
RandomPfxSize/1_000         405.4Ki ± 0%    115.8Ki ±  1%   -71.44% (p=0.002 n=6)
RandomPfxSize/10_000        3.540Mi ± 0%    1.102Mi ±  0%   -68.87% (p=0.002 n=6)
RandomPfxSize/100_000       31.48Mi ± 0%    10.99Mi ±  0%   -65.09% (p=0.002 n=6)
RandomPfxSize/1_000_000     250.6Mi ± 0%    109.9Mi ±  0%   -56.16% (p=0.002 n=6)
geomean                     2.964Mi         1.193Mi         -59.74%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm  │            lpmtrie/size.bm             │
                         │     bytes     │     bytes       vs base                │
Tier1PfxSize/100            34.78Ki ± 5%    24.25Ki ±  9%   -30.28% (p=0.002 n=6)
Tier1PfxSize/1_000          228.2Ki ± 1%    202.3Ki ±  4%   -11.35% (p=0.002 n=6)
Tier1PfxSize/10_000         1.445Mi ± 0%    1.942Mi ±  3%   +34.40% (p=0.002 n=6)
Tier1PfxSize/100_000        6.339Mi ± 0%   19.330Mi ±  3%  +204.96% (p=0.002 n=6)
Tier1PfxSize/1_000_000      31.67Mi ± 0%   189.92Mi ±  6%  +499.76% (p=0.002 n=6)
RandomPfx4Size/100          37.73Ki ± 4%    23.20Ki ±  9%   -38.50% (p=0.002 n=6)
RandomPfx4Size/1_000        247.2Ki ± 1%    198.5Ki ±  3%   -19.68% (p=0.002 n=6)
RandomPfx4Size/10_000       1.980Mi ± 0%    1.898Mi ±  3%    -4.14% (p=0.002 n=6)
RandomPfx4Size/100_000      14.98Mi ± 0%    18.28Mi ±  5%   +22.03% (p=0.002 n=6)
RandomPfx4Size/1_000_000    91.85Mi ± 0%   163.51Mi ±  8%   +78.01% (p=0.002 n=6)
RandomPfx6Size/100         117.36Ki ± 1%    25.53Ki ± 12%   -78.25% (p=0.002 n=6)
RandomPfx6Size/1_000       1071.4Ki ± 0%    222.0Ki ±  8%   -79.28% (p=0.002 n=6)
RandomPfx6Size/10_000       9.807Mi ± 0%    2.129Mi ±  7%   -78.29% (p=0.002 n=6)
RandomPfx6Size/100_000      92.24Mi ± 0%    20.93Mi ±  8%   -77.31% (p=0.002 n=6)
RandomPfx6Size/1_000_000    860.4Mi ± 0%    204.0Mi ±  8%   -76.30% (p=0.002 n=6)
RandomPfxSize/100           51.11Ki ± 3%    23.56Ki ± 10%   -53.90% (p=0.002 n=6)
RandomPfxSize/1_000         405.4Ki ± 0%    202.8Ki ±  4%   -49.98% (p=0.002 n=6)
RandomPfxSize/10_000        3.540Mi ± 0%    1.943Mi ±  4%   -45.12% (p=0.002 n=6)
RandomPfxSize/100_000       31.48Mi ± 0%    18.78Mi ±  5%   -40.34% (p=0.002 n=6)
RandomPfxSize/1_000_000     250.6Mi ± 0%    171.9Mi ±  7%   -31.40% (p=0.002 n=6)
geomean                     2.964Mi         2.011Mi         -32.16%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm  │           cidranger/size.bm            │
                         │     bytes     │     bytes      vs base                 │
Tier1PfxSize/100            34.78Ki ± 5%    56.03Ki ± 5%    +61.10% (p=0.002 n=6)
Tier1PfxSize/1_000          228.2Ki ± 1%    525.9Ki ± 3%   +130.41% (p=0.002 n=6)
Tier1PfxSize/10_000         1.445Mi ± 0%    5.091Mi ± 2%   +252.31% (p=0.002 n=6)
Tier1PfxSize/100_000        6.339Mi ± 0%   50.258Mi ± 2%   +692.88% (p=0.002 n=6)
Tier1PfxSize/1_000_000      31.67Mi ± 0%   477.45Mi ± 2%  +1407.78% (p=0.002 n=6)
RandomPfx4Size/100          37.73Ki ± 4%    54.16Ki ± 6%    +43.55% (p=0.002 n=6)
RandomPfx4Size/1_000        247.2Ki ± 1%    514.0Ki ± 3%   +107.95% (p=0.002 n=6)
RandomPfx4Size/10_000       1.980Mi ± 0%    4.900Mi ± 3%   +147.54% (p=0.002 n=6)
RandomPfx4Size/100_000      14.98Mi ± 0%    45.76Mi ± 3%   +205.51% (p=0.002 n=6)
RandomPfx4Size/1_000_000    91.85Mi ± 0%   396.20Mi ± 3%   +331.33% (p=0.002 n=6)
RandomPfx6Size/100         117.36Ki ± 1%    61.08Ki ± 3%    -47.96% (p=0.002 n=6)
RandomPfx6Size/1_000       1071.4Ki ± 0%    579.3Ki ± 0%    -45.93% (p=0.002 n=6)
RandomPfx6Size/10_000       9.807Mi ± 0%    5.582Mi ± 0%    -43.09% (p=0.002 n=6)
RandomPfx6Size/100_000      92.24Mi ± 0%    54.75Mi ± 0%    -40.65% (p=0.002 n=6)
RandomPfx6Size/1_000_000    860.4Mi ± 0%    534.1Mi ± 0%    -37.92% (p=0.002 n=6)
RandomPfxSize/100           51.11Ki ± 3%    55.07Ki ± 5%     +7.75% (p=0.002 n=6)
RandomPfxSize/1_000         405.4Ki ± 0%    524.9Ki ± 2%    +29.48% (p=0.002 n=6)
RandomPfxSize/10_000        3.540Mi ± 0%    5.024Mi ± 2%    +41.94% (p=0.002 n=6)
RandomPfxSize/100_000       31.48Mi ± 0%    47.53Mi ± 2%    +50.98% (p=0.002 n=6)
RandomPfxSize/1_000_000     250.6Mi ± 0%    424.9Mi ± 2%    +69.54% (p=0.002 n=6)
geomean                     2.964Mi         5.068Mi         +70.97%
```

## lookup (longest-prefix-match)

In the lookup, `art` is the champion, closely followed by `bart`. 

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │            art/lookup.bm             │
                                    │     sec/op     │    sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4             56.84n ± 44%   49.98n ± 10%        ~ (p=0.481 n=10)
LpmTier1Pfxs/RandomMatchIP6             63.20n ± 28%   46.56n ± 11%  -26.32% (p=0.023 n=10)
LpmTier1Pfxs/RandomMissIP4              81.31n ± 47%   29.93n ±  1%  -63.20% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              19.52n ± 52%   29.87n ±  1%  +52.98% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     78.32n ± 41%   50.49n ± 11%  -35.53% (p=0.011 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     70.23n ± 19%   47.53n ± 19%  -32.32% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4     118.00n ± 30%   30.67n ±  7%  -74.01% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      82.05n ± 11%   30.50n ±  6%  -62.83% (p=0.000 n=10)
geomean                                 64.61n         38.34n        -40.66%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │            cidrtree/lookup.bm            │
                                    │     sec/op     │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             56.84n ± 44%    965.50n ± 56%  +1598.78% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             63.20n ± 28%    916.20n ± 16%  +1349.80% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              81.31n ± 47%   1182.00n ± 30%  +1353.70% (p=0.002 n=10)
LpmTier1Pfxs/RandomMissIP6              19.52n ± 52%    419.05n ± 90%  +2046.22% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     78.32n ± 41%   1248.00n ± 33%  +1493.56% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     70.23n ± 19%   1314.00n ± 41%  +1770.86% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      118.0n ± 30%    1187.5n ± 13%   +906.36% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      82.05n ± 11%   1112.50n ± 30%  +1255.88% (p=0.000 n=10)
geomean                                 64.61n           993.5n        +1437.70%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │           critbitgo/lookup.bm           │
                                    │     sec/op     │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             56.84n ± 44%   400.05n ± 19%   +603.88% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             63.20n ± 28%   527.20n ± 12%   +734.24% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              81.31n ± 47%   573.70n ± 58%   +605.57% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              19.52n ± 52%   478.60n ± 46%  +2351.22% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     78.32n ± 41%   465.95n ± 39%   +494.97% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     70.23n ± 19%   640.80n ± 18%   +812.37% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      118.0n ± 30%    506.6n ± 30%   +329.32% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      82.05n ± 11%   488.95n ± 24%   +495.92% (p=0.000 n=10)
geomean                                 64.61n          505.8n         +682.83%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │             lpmtrie/lookup.bm             │
                                    │     sec/op     │     sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4             56.84n ± 44%    229.55n ±   9%   +303.89% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             63.20n ± 28%    234.90n ±  48%   +271.71% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              81.31n ± 47%     89.68n ± 147%          ~ (p=0.912 n=10)
LpmTier1Pfxs/RandomMissIP6              19.52n ± 52%     79.46n ±   1%   +306.97% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     78.32n ± 41%   1047.00n ±   2%  +1236.91% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     70.23n ± 19%    187.95n ±  23%   +167.60% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      118.0n ± 30%    1071.5n ±   7%   +808.05% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      82.05n ± 11%    223.60n ±  19%   +172.52% (p=0.000 n=10)
geomean                                 64.61n           255.4n          +295.33%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │            cidranger/lookup.bm            │
                                    │     sec/op     │     sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4             56.84n ± 44%    359.20n ±  34%   +532.00% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             63.20n ± 28%    291.45n ±  47%   +361.19% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              81.31n ± 47%    218.90n ±  50%   +169.22% (p=0.003 n=10)
LpmTier1Pfxs/RandomMissIP6              19.52n ± 52%    269.40n ±   8%  +1279.77% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     78.32n ± 41%   1586.00n ±  66%  +1925.15% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     70.23n ± 19%    510.10n ± 125%   +626.28% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      118.0n ± 30%     451.8n ±  13%   +282.92% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      82.05n ± 11%    308.00n ±   7%   +275.38% (p=0.000 n=10)
geomean                                 64.61n           403.0n          +523.67%
```

## update, insert/delete

`bart` is by far the fastest algorithm for updates under all competitors.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │             art/update.bm             │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100           50.32n ±  3%   105.85n ± 1%  +110.35% (p=0.000 n=10)
Insert/Insert_into_1_000         49.93n ±  0%   125.45n ± 1%  +151.25% (p=0.000 n=10)
Insert/Insert_into_10_000        50.05n ±  0%   129.05n ± 1%  +157.84% (p=0.000 n=10)
Insert/Insert_into_100_000       50.00n ±  1%   120.55n ± 2%  +141.12% (p=0.000 n=10)
Insert/Insert_into_1_000_000     50.69n ±  0%   121.00n ± 0%  +138.71% (p=0.000 n=10)
Delete/Delete_from_100           16.82n ±  1%    34.04n ± 3%  +102.38% (p=0.000 n=10)
Delete/Delete_from_1_000         39.99n ±  1%    59.06n ± 0%   +47.69% (p=0.000 n=10)
Delete/Delete_from_10_000        40.15n ±  1%    58.93n ± 0%   +46.76% (p=0.000 n=10)
Delete/Delete_from_100_000       40.27n ±  2%    58.71n ± 0%   +45.79% (p=0.000 n=10)
Delete/Delete_from_1_000_000     40.58n ± 23%    58.50n ± 0%   +44.19% (p=0.000 n=10)
geomean                          41.19n          79.57n        +93.17%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │            cidrtree/update.bm            │
                             │     sec/op     │    sec/op      vs base                   │
Insert/Insert_into_100           50.32n ±  3%    945.20n ± 0%   +1778.38% (p=0.000 n=10)
Insert/Insert_into_1_000         49.93n ±  0%   1452.00n ± 0%   +2808.07% (p=0.000 n=10)
Insert/Insert_into_10_000        50.05n ±  0%   2125.00n ± 0%   +4145.75% (p=0.000 n=10)
Insert/Insert_into_100_000       50.00n ±  1%   1595.00n ± 0%   +3090.32% (p=0.000 n=10)
Insert/Insert_into_1_000_000     50.69n ±  0%   2627.50n ± 0%   +5083.47% (p=0.000 n=10)
Delete/Delete_from_100           16.82n ±  1%   1502.00n ± 0%   +8829.85% (p=0.000 n=10)
Delete/Delete_from_1_000         39.99n ±  1%   1765.00n ± 0%   +4313.60% (p=0.000 n=10)
Delete/Delete_from_10_000        40.15n ±  1%   3444.00n ± 0%   +8477.83% (p=0.000 n=10)
Delete/Delete_from_100_000       40.27n ±  2%   3493.50n ± 0%   +8575.19% (p=0.000 n=10)
Delete/Delete_from_1_000_000     40.58n ± 23%   4819.00n ± 0%  +11776.77% (p=0.000 n=10)
geomean                          41.19n           2.125µ        +5059.16%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │          critbitgo/update.bm          │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100           50.32n ±  3%   109.95n ± 1%  +118.50% (p=0.000 n=10)
Insert/Insert_into_1_000         49.93n ±  0%   121.75n ± 1%  +143.84% (p=0.000 n=10)
Insert/Insert_into_10_000        50.05n ±  0%   125.50n ± 1%  +150.75% (p=0.000 n=10)
Insert/Insert_into_100_000       50.00n ±  1%   147.50n ± 1%  +195.03% (p=0.000 n=10)
Insert/Insert_into_1_000_000     50.69n ±  0%   152.20n ± 2%  +200.26% (p=0.000 n=10)
Delete/Delete_from_100           16.82n ±  1%    99.93n ± 1%  +494.11% (p=0.000 n=10)
Delete/Delete_from_1_000         39.99n ±  1%   108.40n ± 1%  +171.07% (p=0.000 n=10)
Delete/Delete_from_10_000        40.15n ±  1%   112.55n ± 1%  +180.32% (p=0.000 n=10)
Delete/Delete_from_100_000       40.27n ±  2%   131.40n ± 2%  +226.30% (p=0.000 n=10)
Delete/Delete_from_1_000_000     40.58n ± 23%   127.40n ± 2%  +213.99% (p=0.000 n=10)
geomean                          41.19n          122.6n       +197.73%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │           lpmtrie/update.bm            │
                             │     sec/op     │    sec/op     vs base                  │
Insert/Insert_into_100           50.32n ±  3%   314.55n ± 0%   +525.10% (p=0.000 n=10)
Insert/Insert_into_1_000         49.93n ±  0%   352.45n ± 2%   +605.89% (p=0.000 n=10)
Insert/Insert_into_10_000        50.05n ±  0%   362.60n ± 1%   +624.48% (p=0.000 n=10)
Insert/Insert_into_100_000       50.00n ±  1%   507.25n ± 1%   +914.60% (p=0.000 n=10)
Insert/Insert_into_1_000_000     50.69n ±  0%   656.80n ± 1%  +1195.72% (p=0.000 n=10)
Delete/Delete_from_100           16.82n ±  1%    77.03n ± 3%   +358.00% (p=0.000 n=10)
Delete/Delete_from_1_000         39.99n ±  1%   123.30n ± 0%   +208.33% (p=0.000 n=10)
Delete/Delete_from_10_000        40.15n ±  1%   149.50n ± 1%   +272.35% (p=0.000 n=10)
Delete/Delete_from_100_000       40.27n ±  2%   266.55n ± 0%   +561.91% (p=0.000 n=10)
Delete/Delete_from_1_000_000     40.58n ± 23%   309.70n ± 3%   +663.28% (p=0.000 n=10)
geomean                          41.19n          262.8n        +537.94%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │           cidranger/update.bm            │
                             │     sec/op     │    sec/op      vs base                   │
Insert/Insert_into_100           50.32n ±  3%   2442.50n ± 1%   +4753.93% (p=0.000 n=10)
Insert/Insert_into_1_000         49.93n ±  0%   4130.00n ± 2%   +8171.58% (p=0.000 n=10)
Insert/Insert_into_10_000        50.05n ±  0%   4849.00n ± 1%   +9588.31% (p=0.000 n=10)
Insert/Insert_into_100_000       50.00n ±  1%   8110.50n ± 1%  +16122.62% (p=0.000 n=10)
Insert/Insert_into_1_000_000     50.69n ±  0%   7579.00n ± 1%  +14851.67% (p=0.000 n=10)
Delete/Delete_from_100           16.82n ±  1%    385.05n ± 3%   +2189.24% (p=0.000 n=10)
Delete/Delete_from_1_000         39.99n ±  1%    421.45n ± 0%    +953.89% (p=0.000 n=10)
Delete/Delete_from_10_000        40.15n ±  1%    442.30n ± 1%   +1001.62% (p=0.000 n=10)
Delete/Delete_from_100_000       40.27n ±  2%    534.85n ± 8%   +1228.16% (p=0.000 n=10)
Delete/Delete_from_1_000_000     40.58n ± 23%    549.85n ± 0%   +1255.14% (p=0.000 n=10)
geomean                          41.19n           1.514µ        +3576.63%
```
