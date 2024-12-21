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

`bart` has a factor of ~20 lower memory consumption compared to `art`, but is by
a factor of ~1,25 slower in lookup times.

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
RandomPfx4Size/100         31.91Ki ± 5%    751.23Ki ± 0%  +2253.93% (p=0.002 n=6)
RandomPfx4Size/1_000       217.3Ki ± 1%    7241.9Ki ± 0%  +3232.85% (p=0.002 n=6)
RandomPfx4Size/10_000      1.738Mi ± 0%    57.774Mi ± 0%  +3224.76% (p=0.002 n=6)
RandomPfx4Size/100_000     13.19Mi ± 0%    522.14Mi ± 0%  +3860.08% (p=0.002 n=6)
RandomPfx4Size/1_000_000   79.85Mi ± 0%
RandomPfx6Size/100         94.30Ki ± 2%    732.11Ki ± 0%   +676.32% (p=0.002 n=6)
RandomPfx6Size/1_000       961.0Ki ± 0%    7720.0Ki ± 0%   +703.33% (p=0.002 n=6)
RandomPfx6Size/10_000      8.588Mi ± 0%    65.319Mi ± 0%   +660.61% (p=0.002 n=6)
RandomPfx6Size/100_000     81.63Mi ± 0%    748.08Mi ± 0%   +816.45% (p=0.002 n=6)
RandomPfx6Size/1_000_000   761.6Mi ± 0%
RandomPfxSize/100          43.28Ki ± 4%    700.64Ki ± 0%  +1518.81% (p=0.002 n=6)
RandomPfxSize/1_000        358.8Ki ± 0%    7445.9Ki ± 0%  +1975.11% (p=0.002 n=6)
RandomPfxSize/10_000       3.079Mi ± 0%    59.978Mi ± 0%  +1847.65% (p=0.002 n=6)
RandomPfxSize/100_000      27.65Mi ± 0%    554.55Mi ± 0%  +1905.64% (p=0.002 n=6)
RandomPfxSize/1_000_000    221.0Mi ± 0%
geomean                    2.575Mi          22.68Mi       +1856.79%               ¹
¹ benchmark set differs from baseline; geomeans may not be comparable

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm  │           cidrtree/size.bm            │
                         │     bytes     │     bytes      vs base                │
Tier1PfxSize/100            31.52Ki ± 5%   12.36Ki ± 14%   -60.78% (p=0.002 n=6)
Tier1PfxSize/1_000         204.20Ki ± 1%   81.54Ki ±  2%   -60.07% (p=0.002 n=6)
Tier1PfxSize/10_000        1318.8Ki ± 0%   784.7Ki ±  0%   -40.50% (p=0.002 n=6)
Tier1PfxSize/100_000        5.397Mi ± 0%   7.633Mi ±  0%   +41.43% (p=0.002 n=6)
Tier1PfxSize/1_000_000      24.83Mi ± 0%   76.30Mi ±  0%  +207.29% (p=0.002 n=6)
RandomPfx4Size/100          31.91Ki ± 5%   11.61Ki ± 14%   -63.62% (p=0.002 n=6)
RandomPfx4Size/1_000       217.29Ki ± 1%   81.54Ki ±  2%   -62.47% (p=0.002 n=6)
RandomPfx4Size/10_000      1779.4Ki ± 0%   784.7Ki ±  0%   -55.90% (p=0.002 n=6)
RandomPfx4Size/100_000     13.185Mi ± 0%   7.633Mi ±  0%   -42.11% (p=0.002 n=6)
RandomPfx4Size/1_000_000    79.85Mi ± 0%   76.30Mi ±  0%    -4.45% (p=0.002 n=6)
RandomPfx6Size/100          94.30Ki ± 2%   11.61Ki ± 14%   -87.69% (p=0.002 n=6)
RandomPfx6Size/1_000       961.01Ki ± 0%   81.54Ki ±  2%   -91.52% (p=0.002 n=6)
RandomPfx6Size/10_000      8793.8Ki ± 0%   784.7Ki ±  0%   -91.08% (p=0.002 n=6)
RandomPfx6Size/100_000     81.629Mi ± 0%   7.633Mi ±  0%   -90.65% (p=0.002 n=6)
RandomPfx6Size/1_000_000   761.65Mi ± 0%   76.30Mi ±  0%   -89.98% (p=0.002 n=6)
RandomPfxSize/100           43.28Ki ± 4%   11.61Ki ± 14%   -73.18% (p=0.002 n=6)
RandomPfxSize/1_000        358.82Ki ± 0%   81.54Ki ±  2%   -77.28% (p=0.002 n=6)
RandomPfxSize/10_000       3153.4Ki ± 0%   784.7Ki ±  0%   -75.12% (p=0.002 n=6)
RandomPfxSize/100_000      27.649Mi ± 0%   7.633Mi ±  0%   -72.39% (p=0.002 n=6)
RandomPfxSize/1_000_000    220.96Mi ± 0%   76.30Mi ±  0%   -65.47% (p=0.002 n=6)
geomean                     2.575Mi        856.4Ki         -67.53%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │           critbitgo/size.bm            │
                         │    bytes     │     bytes       vs base                │
Tier1PfxSize/100           31.52Ki ± 5%    15.67Ki ± 11%   -50.27% (p=0.002 n=6)
Tier1PfxSize/1_000         204.2Ki ± 1%    115.3Ki ±  1%   -43.56% (p=0.002 n=6)
Tier1PfxSize/10_000        1.288Mi ± 0%    1.094Mi ±  0%   -15.09% (p=0.002 n=6)
Tier1PfxSize/100_000       5.397Mi ± 0%   10.913Mi ±  0%  +102.21% (p=0.002 n=6)
Tier1PfxSize/1_000_000     24.83Mi ± 0%   109.12Mi ±  0%  +339.47% (p=0.002 n=6)
RandomPfx4Size/100         31.91Ki ± 5%    14.70Ki ± 11%   -53.93% (p=0.002 n=6)
RandomPfx4Size/1_000       217.3Ki ± 1%    112.8Ki ±  1%   -48.11% (p=0.002 n=6)
RandomPfx4Size/10_000      1.738Mi ± 0%    1.071Mi ±  0%   -38.34% (p=0.002 n=6)
RandomPfx4Size/100_000     13.19Mi ± 0%    10.68Mi ±  0%   -18.96% (p=0.002 n=6)
RandomPfx4Size/1_000_000   79.85Mi ± 0%   106.81Mi ±  0%   +33.77% (p=0.002 n=6)
RandomPfx6Size/100         94.30Ki ± 2%    16.25Ki ± 10%   -82.77% (p=0.002 n=6)
RandomPfx6Size/1_000       961.0Ki ± 0%    128.4Ki ±  1%   -86.64% (p=0.002 n=6)
RandomPfx6Size/10_000      8.588Mi ± 0%    1.224Mi ±  0%   -85.75% (p=0.002 n=6)
RandomPfx6Size/100_000     81.63Mi ± 0%    12.21Mi ±  0%   -85.04% (p=0.002 n=6)
RandomPfx6Size/1_000_000   761.6Mi ± 0%    122.1Mi ±  0%   -83.97% (p=0.002 n=6)
RandomPfxSize/100          43.28Ki ± 4%    14.97Ki ± 11%   -65.42% (p=0.002 n=6)
RandomPfxSize/1_000        358.8Ki ± 0%    115.8Ki ±  1%   -67.73% (p=0.002 n=6)
RandomPfxSize/10_000       3.079Mi ± 0%    1.102Mi ±  0%   -64.21% (p=0.002 n=6)
RandomPfxSize/100_000      27.65Mi ± 0%    10.99Mi ±  0%   -60.25% (p=0.002 n=6)
RandomPfxSize/1_000_000    221.0Mi ± 0%    109.9Mi ±  0%   -50.28% (p=0.002 n=6)
geomean                    2.575Mi         1.194Mi         -53.64%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │            lpmtrie/size.bm             │
                         │    bytes     │     bytes       vs base                │
Tier1PfxSize/100           31.52Ki ± 5%    24.84Ki ± 11%   -21.17% (p=0.002 n=6)
Tier1PfxSize/1_000         204.2Ki ± 1%    208.9Ki ±  5%    +2.30% (p=0.015 n=6)
Tier1PfxSize/10_000        1.288Mi ± 0%    2.007Mi ±  4%   +55.86% (p=0.002 n=6)
Tier1PfxSize/100_000       5.397Mi ± 0%   19.979Mi ±  5%  +270.20% (p=0.002 n=6)
Tier1PfxSize/1_000_000     24.83Mi ± 0%   196.30Mi ±  8%  +690.62% (p=0.002 n=6)
RandomPfx4Size/100         31.91Ki ± 5%    24.02Ki ± 10%   -24.75% (p=0.002 n=6)
RandomPfx4Size/1_000       217.3Ki ± 1%    205.5Ki ±  5%    -5.42% (p=0.002 n=6)
RandomPfx4Size/10_000      1.738Mi ± 0%    1.973Mi ±  5%   +13.56% (p=0.002 n=6)
RandomPfx4Size/100_000     13.19Mi ± 0%    19.04Mi ±  6%   +44.37% (p=0.002 n=6)
RandomPfx4Size/1_000_000   79.85Mi ± 0%   170.24Mi ±  9%  +113.20% (p=0.002 n=6)
RandomPfx6Size/100         94.30Ki ± 2%    25.56Ki ± 13%   -72.89% (p=0.002 n=6)
RandomPfx6Size/1_000       961.0Ki ± 0%    222.5Ki ±  8%   -76.85% (p=0.002 n=6)
RandomPfx6Size/10_000      8.588Mi ± 0%    2.128Mi ±  7%   -75.23% (p=0.002 n=6)
RandomPfx6Size/100_000     81.63Mi ± 0%    20.94Mi ±  8%   -74.35% (p=0.002 n=6)
RandomPfx6Size/1_000_000   761.6Mi ± 0%    204.0Mi ±  8%   -73.22% (p=0.002 n=6)
RandomPfxSize/100          43.28Ki ± 4%    24.22Ki ± 11%   -44.04% (p=0.002 n=6)
RandomPfxSize/1_000        358.8Ki ± 0%    209.2Ki ±  5%   -41.69% (p=0.002 n=6)
RandomPfxSize/10_000       3.079Mi ± 0%    2.003Mi ±  5%   -34.95% (p=0.002 n=6)
RandomPfxSize/100_000      27.65Mi ± 0%    19.41Mi ±  7%   -29.80% (p=0.002 n=6)
RandomPfxSize/1_000_000    221.0Mi ± 0%    177.3Mi ±  9%   -19.77% (p=0.002 n=6)
geomean                    2.575Mi         2.062Mi         -19.95%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/size.bm │           cidranger/size.bm            │
                         │    bytes     │     bytes      vs base                 │
Tier1PfxSize/100           31.52Ki ± 5%    56.06Ki ± 5%    +77.89% (p=0.002 n=6)
Tier1PfxSize/1_000         204.2Ki ± 1%    525.9Ki ± 3%   +157.57% (p=0.002 n=6)
Tier1PfxSize/10_000        1.288Mi ± 0%    5.091Mi ± 2%   +295.30% (p=0.002 n=6)
Tier1PfxSize/100_000       5.397Mi ± 0%   50.258Mi ± 2%   +831.27% (p=0.002 n=6)
Tier1PfxSize/1_000_000     24.83Mi ± 0%   477.45Mi ± 2%  +1822.94% (p=0.002 n=6)
RandomPfx4Size/100         31.91Ki ± 5%    54.62Ki ± 6%    +71.16% (p=0.002 n=6)
RandomPfx4Size/1_000       217.3Ki ± 1%    513.5Ki ± 3%   +136.32% (p=0.002 n=6)
RandomPfx4Size/10_000      1.738Mi ± 0%    4.893Mi ± 3%   +181.59% (p=0.002 n=6)
RandomPfx4Size/100_000     13.19Mi ± 0%    45.74Mi ± 3%   +246.91% (p=0.002 n=6)
RandomPfx4Size/1_000_000   79.85Mi ± 0%   396.20Mi ± 3%   +396.17% (p=0.002 n=6)
RandomPfx6Size/100         94.30Ki ± 2%    61.11Ki ± 3%    -35.20% (p=0.002 n=6)
RandomPfx6Size/1_000       961.0Ki ± 0%    580.8Ki ± 0%    -39.56% (p=0.002 n=6)
RandomPfx6Size/10_000      8.588Mi ± 0%    5.599Mi ± 0%    -34.80% (p=0.002 n=6)
RandomPfx6Size/100_000     81.63Mi ± 0%    54.80Mi ± 0%    -32.86% (p=0.002 n=6)
RandomPfx6Size/1_000_000   761.6Mi ± 0%    534.1Mi ± 0%    -29.87% (p=0.002 n=6)
RandomPfxSize/100          43.28Ki ± 4%    55.12Ki ± 5%    +27.36% (p=0.002 n=6)
RandomPfxSize/1_000        358.8Ki ± 0%    524.1Ki ± 2%    +46.08% (p=0.002 n=6)
RandomPfxSize/10_000       3.079Mi ± 0%    5.024Mi ± 2%    +63.15% (p=0.002 n=6)
RandomPfxSize/100_000      27.65Mi ± 0%    47.54Mi ± 2%    +71.95% (p=0.002 n=6)
RandomPfxSize/1_000_000    221.0Mi ± 0%    424.9Mi ± 2%    +92.29% (p=0.002 n=6)
geomean                    2.575Mi         5.071Mi         +96.90%
```

## lookup (longest-prefix-match)

In the lookup, `art` is the champion, closely followed by `bart`. 

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │            art/lookup.bm             │
                                    │     sec/op     │    sec/op     vs base                │
LpmTier1Pfxs/RandomMatchIP4             47.73n ± 30%   46.31n ± 19%        ~ (p=0.393 n=10)
LpmTier1Pfxs/RandomMatchIP6             54.03n ± 16%   47.42n ± 13%  -12.23% (p=0.022 n=10)
LpmTier1Pfxs/RandomMissIP4              41.77n ± 48%   29.88n ±  0%        ~ (p=0.468 n=10)
LpmTier1Pfxs/RandomMissIP6              16.73n ±  2%   29.83n ±  0%  +78.27% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     65.70n ± 16%   51.15n ± 10%  -22.15% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     53.68n ± 22%   52.33n ± 10%        ~ (p=0.912 n=10)
LpmRandomPfxs100_000/RandomMissIP4      70.10n ± 18%   29.77n ±  0%  -57.52% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      57.05n ± 15%   29.71n ±  0%  -47.91% (p=0.000 n=10)
geomean                                 47.39n         38.30n        -19.16%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │            cidrtree/lookup.bm            │
                                    │     sec/op     │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             47.73n ± 30%    741.05n ± 94%  +1452.42% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             54.03n ± 16%   1302.00n ± 18%  +2309.77% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              41.77n ± 48%   1100.50n ± 97%  +2534.67% (p=0.003 n=10)
LpmTier1Pfxs/RandomMissIP6              16.73n ±  2%     64.55n ±  2%   +285.80% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     65.70n ± 16%   1002.15n ± 37%  +1425.34% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     53.68n ± 22%   1409.50n ± 36%  +2525.99% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      70.10n ± 18%   1301.00n ± 14%  +1755.92% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      57.05n ± 15%   1033.50n ± 19%  +1711.57% (p=0.000 n=10)
geomean                                 47.39n           775.0n        +1535.57%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │           critbitgo/lookup.bm           │
                                    │     sec/op     │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             47.73n ± 30%   382.20n ± 18%   +700.67% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             54.03n ± 16%   478.50n ± 23%   +785.62% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              41.77n ± 48%   600.70n ± 35%  +1338.11% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              16.73n ±  2%   451.80n ± 27%  +2600.54% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     65.70n ± 16%   310.85n ±  9%   +373.14% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     53.68n ± 22%   491.60n ± 15%   +815.88% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      70.10n ± 18%   478.15n ± 34%   +582.10% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      57.05n ± 15%   470.75n ± 25%   +725.15% (p=0.000 n=10)
geomean                                 47.39n          450.9n         +851.47%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │           lpmtrie/lookup.bm            │
                                    │     sec/op     │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4             47.73n ± 30%   203.95n ± 22%  +327.25% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             54.03n ± 16%   132.70n ± 59%  +145.60% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              41.77n ± 48%   144.90n ± 79%  +246.90% (p=0.050 n=10)
LpmTier1Pfxs/RandomMissIP6              16.73n ±  2%    13.02n ±  0%   -22.18% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     65.70n ± 16%   271.95n ±  6%  +313.93% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     53.68n ± 22%   155.90n ±  5%  +190.45% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      70.10n ± 18%   286.80n ± 10%  +309.13% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      57.05n ± 15%   182.70n ± 22%  +220.25% (p=0.000 n=10)
geomean                                 47.39n          135.5n        +185.89%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                    │ bart/lookup.bm │           cidranger/lookup.bm            │
                                    │     sec/op     │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4             47.73n ± 30%    356.95n ± 39%   +647.77% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6             54.03n ± 16%    288.95n ± 49%   +434.80% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4              41.77n ± 48%    209.65n ± 50%   +401.92% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6              16.73n ±  2%    252.05n ± 12%  +1406.58% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP4     65.70n ± 16%   1130.00n ± 71%  +1619.94% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMatchIP6     53.68n ± 22%    381.20n ± 33%   +610.20% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP4      70.10n ± 18%    383.95n ± 17%   +447.72% (p=0.000 n=10)
LpmRandomPfxs100_000/RandomMissIP6      57.05n ± 15%    273.90n ± 10%   +380.11% (p=0.000 n=10)
geomean                                 47.39n           354.0n         +647.17%
```

## update, insert/delete

`bart` is by far the fastest algorithm for updates under all competitors.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │             art/update.bm              │
                             │     sec/op     │    sec/op      vs base                 │
Insert/Insert_into_100           46.45n ±  7%    88.89n ±  1%   +91.38% (p=0.000 n=10)
Insert/Insert_into_1_000         46.26n ±  1%   108.15n ±  4%  +133.76% (p=0.000 n=10)
Insert/Insert_into_10_000        46.77n ±  2%   107.60n ±  9%  +130.06% (p=0.000 n=10)
Insert/Insert_into_100_000       47.18n ±  4%   124.00n ± 14%  +162.85% (p=0.000 n=10)
Insert/Insert_into_1_000_000     46.35n ±  7%   105.55n ±  1%  +127.70% (p=0.000 n=10)
Delete/Delete_from_100           17.37n ± 18%    13.83n ±  2%   -20.38% (p=0.000 n=10)
Delete/Delete_from_1_000         38.56n ±  6%    40.23n ±  1%         ~ (p=0.060 n=10)
Delete/Delete_from_10_000        38.56n ±  0%    42.36n ±  2%    +9.83% (p=0.000 n=10)
Delete/Delete_from_100_000       38.54n ±  1%    40.38n ±  6%    +4.79% (p=0.000 n=10)
Delete/Delete_from_1_000_000     38.53n ±  0%    40.48n ±  2%    +5.06% (p=0.000 n=10)
geomean                          39.14n          59.12n         +51.06%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │            cidrtree/update.bm            │
                             │     sec/op     │     sec/op      vs base                  │
Insert/Insert_into_100           46.45n ±  7%    858.95n ± 11%  +1749.19% (p=0.000 n=10)
Insert/Insert_into_1_000         46.26n ±  1%   1061.00n ± 11%  +2193.31% (p=0.000 n=10)
Insert/Insert_into_10_000        46.77n ±  2%   1951.00n ±  5%  +4071.48% (p=0.000 n=10)
Insert/Insert_into_100_000       47.18n ±  4%   1423.50n ±  0%  +2917.49% (p=0.000 n=10)
Insert/Insert_into_1_000_000     46.35n ±  7%   2561.50n ±  0%  +5425.83% (p=0.000 n=10)
Delete/Delete_from_100           17.37n ± 18%    597.50n ±  0%  +3339.84% (p=0.000 n=10)
Delete/Delete_from_1_000         38.56n ±  6%   1441.00n ±  0%  +3637.03% (p=0.000 n=10)
Delete/Delete_from_10_000        38.56n ±  0%   1196.00n ±  0%  +3001.26% (p=0.000 n=10)
Delete/Delete_from_100_000       38.54n ±  1%   3183.00n ±  0%  +8158.95% (p=0.000 n=10)
Delete/Delete_from_1_000_000     38.53n ±  0%   3019.00n ±  1%  +7735.45% (p=0.000 n=10)
geomean                          39.14n           1.516µ        +3773.80%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │          critbitgo/update.bm          │
                             │     sec/op     │    sec/op     vs base                 │
Insert/Insert_into_100           46.45n ±  7%    64.88n ± 0%   +39.67% (p=0.000 n=10)
Insert/Insert_into_1_000         46.26n ±  1%    74.98n ± 0%   +62.08% (p=0.000 n=10)
Insert/Insert_into_10_000        46.77n ±  2%    78.31n ± 0%   +67.43% (p=0.000 n=10)
Insert/Insert_into_100_000       47.18n ±  4%   110.65n ± 1%  +134.55% (p=0.000 n=10)
Insert/Insert_into_1_000_000     46.35n ±  7%   143.65n ± 0%  +209.89% (p=0.000 n=10)
Delete/Delete_from_100           17.37n ± 18%    52.81n ± 0%  +204.03% (p=0.000 n=10)
Delete/Delete_from_1_000         38.56n ±  6%    60.19n ± 0%   +56.08% (p=0.000 n=10)
Delete/Delete_from_10_000        38.56n ±  0%    66.20n ± 0%   +71.65% (p=0.000 n=10)
Delete/Delete_from_100_000       38.54n ±  1%    91.83n ± 0%  +138.26% (p=0.000 n=10)
Delete/Delete_from_1_000_000     38.53n ±  0%   100.20n ± 3%  +160.06% (p=0.000 n=10)
geomean                          39.14n          80.70n       +106.21%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │           lpmtrie/update.bm            │
                             │     sec/op     │    sec/op     vs base                  │
Insert/Insert_into_100           46.45n ±  7%   490.75n ± 1%   +956.51% (p=0.000 n=10)
Insert/Insert_into_1_000         46.26n ±  1%   589.80n ± 1%  +1174.83% (p=0.000 n=10)
Insert/Insert_into_10_000        46.77n ±  2%   588.30n ± 6%  +1157.86% (p=0.000 n=10)
Insert/Insert_into_100_000       47.18n ±  4%   442.40n ± 0%   +837.78% (p=0.000 n=10)
Insert/Insert_into_1_000_000     46.35n ±  7%   471.15n ± 0%   +916.40% (p=0.000 n=10)
Delete/Delete_from_100           17.37n ± 18%    61.90n ± 0%   +256.39% (p=0.000 n=10)
Delete/Delete_from_1_000         38.56n ±  6%   124.45n ± 0%   +222.74% (p=0.000 n=10)
Delete/Delete_from_10_000        38.56n ±  0%   131.45n ± 0%   +240.85% (p=0.000 n=10)
Delete/Delete_from_100_000       38.54n ±  1%   262.90n ± 0%   +582.15% (p=0.000 n=10)
Delete/Delete_from_1_000_000     38.53n ±  0%   303.55n ± 0%   +687.83% (p=0.000 n=10)
geomean                          39.14n          279.1n        +613.15%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                             │ bart/update.bm │           cidranger/update.bm            │
                             │     sec/op     │    sec/op      vs base                   │
Insert/Insert_into_100           46.45n ±  7%   2217.50n ± 0%   +4673.95% (p=0.000 n=10)
Insert/Insert_into_1_000         46.26n ±  1%   3849.50n ± 0%   +8220.54% (p=0.000 n=10)
Insert/Insert_into_10_000        46.77n ±  2%   4488.50n ± 0%   +9496.96% (p=0.000 n=10)
Insert/Insert_into_100_000       47.18n ±  4%   7162.00n ± 0%  +15081.77% (p=0.000 n=10)
Insert/Insert_into_1_000_000     46.35n ±  7%   7877.00n ± 0%  +16892.77% (p=0.000 n=10)
Delete/Delete_from_100           17.37n ± 18%    367.35n ± 0%   +2014.85% (p=0.000 n=10)
Delete/Delete_from_1_000         38.56n ±  6%    400.75n ± 1%    +939.29% (p=0.000 n=10)
Delete/Delete_from_10_000        38.56n ±  0%    416.10n ± 0%    +978.96% (p=0.000 n=10)
Delete/Delete_from_100_000       38.54n ±  1%    531.55n ± 0%   +1279.22% (p=0.000 n=10)
Delete/Delete_from_1_000_000     38.53n ±  0%    575.15n ± 0%   +1392.73% (p=0.000 n=10)
geomean                          39.14n           1.448µ        +3599.46%
```
