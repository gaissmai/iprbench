# iprbench
comparing benchmarks for some golang IP routing table implementations:

```
	github.com/gaissmai/bart
	github.com/tailscale/art
	github.com/aromatt/netipds
	github.com/k-sone/critbitgo
	github.com/Asphaltt/lpmtrie
	github.com/yl2chen/cidranger
	github.com/gaissmai/cidrtree
```

The ~1_000_000 **Tier1** prefix test records (IPv4 and IPv6 routes) are from a full routing table with typical
ISP prefix distribution.

In comparison, the prefix lengths for the random test sets are equally distributed between /3-32 for IPv4
and /3-128 bits for IPv6, the randomly generated _default-routes_ with prefix length 0 have been sorted out,
they distorts the lookup times and there is no lookup miss at all.

The **RandomPrefixes** without IP versions labeling are composed of a distribution of 4 parts IPv4
to 1 part IPv6 prefixes, which is approximately the current ratio in the Internet backbone routers.

## make your own benchmarks

```
  $ make dep
  $ make -B all   # takes some time!
```

## lpm (longest-prefix-match)

`bart` is the fastest software algorithm for IP address lookup in routing tables.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │              art/lpm.bm               │
                                     │    sec/op    │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4            25.32n ±  7%   44.94n ±  3%   +77.43% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            32.36n ± 33%   60.08n ±  5%   +85.68% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             28.90n ±  6%   28.80n ±  1%         ~ (p=0.784 n=20)
LpmTier1Pfxs/RandomMissIP6             38.49n ± 24%   36.17n ± 24%         ~ (p=0.579 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     22.39n ±  5%   45.06n ±  1%  +101.30% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     20.00n ± 20%   45.69n ±  1%  +128.43% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      22.32n ±  5%   28.77n ±  0%   +28.87% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      20.80n ± 28%   28.73n ±  0%   +38.18% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    27.03n ±  5%   45.04n ±  1%   +66.66% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    25.16n ±  5%   45.88n ±  1%   +82.33% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     27.20n ±  6%   28.76n ±  0%    +5.74% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     25.32n ±  5%   28.77n ±  0%   +13.63% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   16.02n ± 54%   53.19n ±  9%  +231.92% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   21.30n ± 60%   47.21n ± 20%  +121.64% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    30.40n ± 16%   28.98n ± 12%         ~ (p=0.654 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    28.23n ±  8%   28.84n ±  6%         ~ (p=0.058 n=20)
geomean                                25.17n         37.77n         +50.06%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            netipds/lpm.bm             │
                                     │    sec/op    │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4            25.32n ±  7%   62.60n ± 13%  +147.21% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            32.36n ± 33%   50.84n ± 11%   +57.12% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             28.90n ±  6%   64.12n ±  3%  +121.89% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             38.49n ± 24%   54.68n ±  7%   +42.05% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     22.39n ±  5%   50.27n ±  0%  +124.55% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     20.00n ± 20%   27.91n ±  6%   +39.58% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      22.32n ±  5%   51.80n ±  4%  +132.05% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      20.80n ± 28%   28.86n ±  5%   +38.78% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    27.03n ±  5%   62.20n ±  3%  +130.16% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    25.16n ±  5%   32.66n ±  4%   +29.81% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     27.20n ±  6%   63.62n ±  3%  +133.88% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     25.32n ±  5%   35.27n ±  3%   +39.34% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   16.02n ± 54%   23.95n ±  0%   +49.42% (p=0.027 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   21.30n ± 60%   31.79n ± 33%   +49.27% (p=0.003 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    30.40n ± 16%   23.94n ±  0%   -21.27% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    28.23n ±  8%   40.79n ±  3%   +44.49% (p=0.000 n=20)
geomean                                25.17n         41.56n         +65.12%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            critbitgo/lpm.bm             │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            25.32n ±  7%   159.90n ±  4%   +531.39% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            32.36n ± 33%   196.50n ± 14%   +507.23% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             28.90n ±  6%   658.60n ±  8%  +2178.89% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             38.49n ± 24%   508.40n ±  9%  +1220.86% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     22.39n ±  5%    86.51n ±  3%   +286.46% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     20.00n ± 20%   105.55n ±  2%   +427.75% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      22.32n ±  5%   148.60n ± 20%   +565.62% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      20.80n ± 28%   157.85n ±  4%   +659.08% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    27.03n ±  5%    98.84n ±  2%   +265.75% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    25.16n ±  5%   118.60n ±  3%   +371.38% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     27.20n ±  6%   255.55n ± 13%   +839.52% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     25.32n ±  5%   210.70n ± 15%   +732.31% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   16.02n ± 54%   114.35n ±  9%   +613.57% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   21.30n ± 60%   139.10n ±  5%   +553.05% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    30.40n ± 16%   468.90n ± 21%  +1442.43% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    28.23n ±  8%   357.35n ± 17%  +1165.85% (p=0.000 n=20)
geomean                                25.17n          193.3n         +667.81%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             lpmtrie/lpm.bm              │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            25.32n ±  7%   269.35n ±  8%   +963.57% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            32.36n ± 33%   325.20n ± 18%   +904.94% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             28.90n ±  6%   257.25n ±  9%   +790.14% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             38.49n ± 24%   254.45n ± 16%   +561.08% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     22.39n ±  5%   136.15n ±  3%   +508.22% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     20.00n ± 20%   103.65n ± 23%   +418.25% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      22.32n ±  5%   129.00n ±  2%   +477.83% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      20.80n ± 28%    96.52n ± 10%   +364.15% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    27.03n ±  5%   154.55n ± 13%   +471.88% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    25.16n ±  5%   142.40n ±  4%   +465.98% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     27.20n ±  6%   145.95n ± 19%   +436.58% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     25.32n ±  5%   136.35n ±  6%   +438.61% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   16.02n ± 54%   194.65n ±  6%  +1114.66% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   21.30n ± 60%   174.65n ±  6%   +719.95% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    30.40n ± 16%   205.75n ±  5%   +576.81% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    28.23n ±  8%   190.70n ±  5%   +575.52% (p=0.000 n=20)
geomean                                25.17n          172.0n         +583.42%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            cidranger/lpm.bm            │
                                     │    sec/op    │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            25.32n ±  7%   276.45n ± 22%  +991.61% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            32.36n ± 33%   331.55n ± 25%  +924.57% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             28.90n ±  6%   272.10n ± 14%  +841.52% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             38.49n ± 24%   334.95n ± 18%  +770.23% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     22.39n ±  5%   113.90n ±  8%  +408.82% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     20.00n ± 20%   135.45n ±  8%  +577.25% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      22.32n ±  5%   115.75n ±  4%  +418.48% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      20.80n ± 28%   133.35n ±  7%  +541.26% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    27.03n ±  5%   151.65n ± 12%  +461.15% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    25.16n ±  5%   175.40n ±  7%  +597.14% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     27.20n ±  6%   149.35n ±  7%  +449.08% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     25.32n ±  5%   165.35n ±  5%  +553.17% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   16.02n ± 54%   108.75n ± 31%  +578.63% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   21.30n ± 60%   143.55n ± 36%  +573.94% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    30.40n ± 16%   190.05n ±  4%  +525.16% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    28.23n ±  8%   211.65n ±  6%  +649.73% (p=0.000 n=20)
geomean                                25.17n          175.6n        +597.77%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             cidrtree/lpm.bm              │
                                     │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            25.32n ±  7%    630.15n ±  6%  +2388.25% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            32.36n ± 33%    810.60n ± 18%  +2404.94% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             28.90n ±  6%   1148.00n ± 15%  +3872.32% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             38.49n ± 24%   1287.00n ± 19%  +3243.73% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     22.39n ±  5%    299.35n ± 15%  +1237.28% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     20.00n ± 20%    352.70n ± 32%  +1663.50% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      22.32n ±  5%    474.85n ± 20%  +2026.99% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      20.80n ± 28%    487.20n ± 23%  +2242.87% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    27.03n ±  5%    483.35n ± 13%  +1688.53% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    25.16n ±  5%    536.55n ± 13%  +2032.55% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     27.20n ±  6%    683.70n ± 17%  +2413.60% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     25.32n ±  5%    728.70n ± 29%  +2778.53% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   16.02n ± 54%    594.50n ± 14%  +3609.83% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   21.30n ± 60%    690.10n ± 18%  +3139.91% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    30.40n ± 16%    819.65n ± 12%  +2596.22% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    28.23n ±  8%   1002.15n ± 27%  +3449.95% (p=0.000 n=20)
geomean                                25.17n           640.5n        +2444.56%
```

## size of the routing tables


`bart` has two orders of magnitude lower memory consumption compared to `art`
and is similar in low memory consumption to a binary search tree, like the `cidrtree` but
much faster in lookup times.


```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │              art/size.bm               │
                       │ bytes/route  │ bytes/route   vs base                  │
Tier1PfxSize/1_000         88.13 ± 2%   7591.00 ± 0%   +8513.41% (p=0.002 n=6)
Tier1PfxSize/10_000        66.04 ± 0%   4889.00 ± 0%   +7303.09% (p=0.002 n=6)
Tier1PfxSize/100_000       30.24 ± 0%   1669.00 ± 0%   +5419.18% (p=0.002 n=6)
Tier1PfxSize/200_000       20.48 ± 0%   1098.00 ± 0%   +5261.33% (p=0.002 n=6)
RandomPfx4Size/1_000       83.88 ± 2%   7394.00 ± 0%   +8714.97% (p=0.002 n=6)
RandomPfx4Size/10_000      61.56 ± 0%   6043.00 ± 0%   +9716.44% (p=0.002 n=6)
RandomPfx4Size/100_000     68.02 ± 0%   5455.00 ± 0%   +7919.70% (p=0.002 n=6)
RandomPfx4Size/200_000     60.81 ± 0%   4832.00 ± 0%   +7846.06% (p=0.002 n=6)
RandomPfx6Size/1_000       87.42 ± 2%   7845.00 ± 0%   +8873.92% (p=0.002 n=6)
RandomPfx6Size/10_000      67.78 ± 0%   6829.00 ± 0%   +9975.24% (p=0.002 n=6)
RandomPfx6Size/100_000     86.26 ± 0%   7837.00 ± 0%   +8985.32% (p=0.002 n=6)
RandomPfx6Size/200_000     81.52 ± 0%   7602.00 ± 0%   +9225.32% (p=0.002 n=6)
RandomPfxSize/1_000        84.43 ± 2%   7564.00 ± 0%   +8858.90% (p=0.002 n=6)
RandomPfxSize/10_000       59.68 ± 0%   6226.00 ± 0%  +10332.31% (p=0.002 n=6)
RandomPfxSize/100_000      67.88 ± 0%   5792.00 ± 0%   +8432.70% (p=0.002 n=6)
RandomPfxSize/200_000      64.58 ± 0%   5454.00 ± 0%   +8345.34% (p=0.002 n=6)
geomean                    63.55        5.170Ki        +8229.38%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │           netipds/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         88.13 ± 2%   101.00 ± 2%   +14.60% (p=0.002 n=6)
Tier1PfxSize/10_000        66.04 ± 0%    96.08 ± 0%   +45.49% (p=0.002 n=6)
Tier1PfxSize/100_000       30.24 ± 0%    94.34 ± 0%  +211.97% (p=0.002 n=6)
Tier1PfxSize/200_000       20.48 ± 0%    93.27 ± 0%  +355.42% (p=0.002 n=6)
RandomPfx4Size/1_000       83.88 ± 2%    99.73 ± 2%   +18.90% (p=0.002 n=6)
RandomPfx4Size/10_000      61.56 ± 0%    94.19 ± 0%   +53.01% (p=0.002 n=6)
RandomPfx4Size/100_000     68.02 ± 0%    86.51 ± 0%   +27.18% (p=0.002 n=6)
RandomPfx4Size/200_000     60.81 ± 0%    82.89 ± 0%   +36.31% (p=0.002 n=6)
RandomPfx6Size/1_000       87.42 ± 2%    99.93 ± 2%   +14.31% (p=0.002 n=6)
RandomPfx6Size/10_000      67.78 ± 0%    95.07 ± 0%   +40.26% (p=0.002 n=6)
RandomPfx6Size/100_000     86.26 ± 0%    92.64 ± 0%    +7.40% (p=0.002 n=6)
RandomPfx6Size/200_000     81.52 ± 0%    91.94 ± 0%   +12.78% (p=0.002 n=6)
RandomPfxSize/1_000        84.43 ± 2%    99.82 ± 2%   +18.23% (p=0.002 n=6)
RandomPfxSize/10_000       59.68 ± 0%    94.22 ± 0%   +57.88% (p=0.002 n=6)
RandomPfxSize/100_000      67.88 ± 0%    87.84 ± 0%   +29.40% (p=0.002 n=6)
RandomPfxSize/200_000      64.58 ± 0%    84.76 ± 0%   +31.25% (p=0.002 n=6)
geomean                    63.55         93.23        +46.70%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          critbitgo/size.bm          │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         88.13 ± 2%   119.50 ± 2%   +35.60% (p=0.002 n=6)
Tier1PfxSize/10_000        66.04 ± 0%   114.70 ± 0%   +73.68% (p=0.002 n=6)
Tier1PfxSize/100_000       30.24 ± 0%   114.40 ± 0%  +278.31% (p=0.002 n=6)
Tier1PfxSize/200_000       20.48 ± 0%   114.40 ± 0%  +458.59% (p=0.002 n=6)
RandomPfx4Size/1_000       83.88 ± 2%   116.10 ± 2%   +38.41% (p=0.002 n=6)
RandomPfx4Size/10_000      61.56 ± 0%   112.40 ± 0%   +82.59% (p=0.002 n=6)
RandomPfx4Size/100_000     68.02 ± 0%   112.00 ± 0%   +64.66% (p=0.002 n=6)
RandomPfx4Size/200_000     60.81 ± 0%   112.00 ± 0%   +84.18% (p=0.002 n=6)
RandomPfx6Size/1_000       87.42 ± 2%   132.40 ± 1%   +51.45% (p=0.002 n=6)
RandomPfx6Size/10_000      67.78 ± 0%   128.40 ± 0%   +89.44% (p=0.002 n=6)
RandomPfx6Size/100_000     86.26 ± 0%   128.00 ± 0%   +48.39% (p=0.002 n=6)
RandomPfx6Size/200_000     81.52 ± 0%   128.00 ± 0%   +57.02% (p=0.002 n=6)
RandomPfxSize/1_000        84.43 ± 2%   119.10 ± 2%   +41.06% (p=0.002 n=6)
RandomPfxSize/10_000       59.68 ± 0%   115.50 ± 0%   +93.53% (p=0.002 n=6)
RandomPfxSize/100_000      67.88 ± 0%   115.20 ± 0%   +69.71% (p=0.002 n=6)
RandomPfxSize/200_000      64.58 ± 0%   115.20 ± 0%   +78.38% (p=0.002 n=6)
geomean                    63.55         118.4        +86.31%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │           lpmtrie/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         88.13 ± 2%   215.40 ± 5%  +144.41% (p=0.002 n=6)
Tier1PfxSize/10_000        66.04 ± 0%   210.50 ± 5%  +218.75% (p=0.002 n=6)
Tier1PfxSize/100_000       30.24 ± 0%   209.90 ± 5%  +594.11% (p=0.002 n=6)
Tier1PfxSize/200_000       20.48 ± 0%   209.20 ± 5%  +921.48% (p=0.002 n=6)
RandomPfx4Size/1_000       83.88 ± 2%   212.00 ± 5%  +152.74% (p=0.002 n=6)
RandomPfx4Size/10_000      61.56 ± 0%   206.70 ± 5%  +235.77% (p=0.002 n=6)
RandomPfx4Size/100_000     68.02 ± 0%   199.40 ± 6%  +193.15% (p=0.002 n=6)
RandomPfx4Size/200_000     60.81 ± 0%   194.70 ± 7%  +220.18% (p=0.002 n=6)
RandomPfx6Size/1_000       87.42 ± 2%   228.30 ± 8%  +161.15% (p=0.002 n=6)
RandomPfx6Size/10_000      67.78 ± 0%   223.00 ± 7%  +229.01% (p=0.002 n=6)
RandomPfx6Size/100_000     86.26 ± 0%   219.30 ± 8%  +154.23% (p=0.002 n=6)
RandomPfx6Size/200_000     81.52 ± 0%   217.80 ± 8%  +167.17% (p=0.002 n=6)
RandomPfxSize/1_000        84.43 ± 2%   215.30 ± 6%  +155.00% (p=0.002 n=6)
RandomPfxSize/10_000       59.68 ± 0%   209.90 ± 5%  +251.71% (p=0.002 n=6)
RandomPfxSize/100_000      67.88 ± 0%   203.50 ± 7%  +199.79% (p=0.002 n=6)
RandomPfxSize/200_000      64.58 ± 0%   199.50 ± 7%  +208.92% (p=0.002 n=6)
geomean                    63.55         210.7       +231.56%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidranger/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                 │
Tier1PfxSize/1_000         88.13 ± 2%   539.60 ± 3%   +512.28% (p=0.002 n=6)
Tier1PfxSize/10_000        66.04 ± 0%   533.90 ± 3%   +708.45% (p=0.002 n=6)
Tier1PfxSize/100_000       30.24 ± 0%   527.20 ± 2%  +1643.39% (p=0.002 n=6)
Tier1PfxSize/200_000       20.48 ± 0%   522.20 ± 2%  +2449.80% (p=0.002 n=6)
RandomPfx4Size/1_000       83.88 ± 2%   527.20 ± 3%   +528.52% (p=0.002 n=6)
RandomPfx4Size/10_000      61.56 ± 0%   514.60 ± 3%   +735.93% (p=0.002 n=6)
RandomPfx4Size/100_000     68.02 ± 0%   479.80 ± 3%   +605.38% (p=0.002 n=6)
RandomPfx4Size/200_000     60.81 ± 0%   463.30 ± 3%   +661.88% (p=0.002 n=6)
RandomPfx6Size/1_000       87.42 ± 2%   595.10 ± 0%   +580.74% (p=0.002 n=6)
RandomPfx6Size/10_000      67.78 ± 0%   584.20 ± 0%   +761.91% (p=0.002 n=6)
RandomPfx6Size/100_000     86.26 ± 0%   573.80 ± 0%   +565.20% (p=0.002 n=6)
RandomPfx6Size/200_000     81.52 ± 0%   569.60 ± 0%   +598.72% (p=0.002 n=6)
RandomPfxSize/1_000        84.43 ± 2%   536.90 ± 2%   +535.91% (p=0.002 n=6)
RandomPfxSize/10_000       59.68 ± 0%   526.80 ± 2%   +782.71% (p=0.002 n=6)
RandomPfxSize/100_000      67.88 ± 0%   498.00 ± 2%   +633.65% (p=0.002 n=6)
RandomPfxSize/200_000      64.58 ± 0%   484.30 ± 2%   +649.92% (p=0.002 n=6)
geomean                    63.55         528.5        +731.61%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidrtree/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         88.13 ± 2%    69.26 ± 3%   -21.41% (p=0.002 n=6)
Tier1PfxSize/10_000        66.04 ± 0%    64.37 ± 0%    -2.53% (p=0.002 n=6)
Tier1PfxSize/100_000       30.24 ± 0%    64.04 ± 0%  +111.77% (p=0.002 n=6)
Tier1PfxSize/200_000       20.48 ± 0%    64.02 ± 0%  +212.60% (p=0.002 n=6)
RandomPfx4Size/1_000       83.88 ± 2%    68.16 ± 3%   -18.74% (p=0.002 n=6)
RandomPfx4Size/10_000      61.56 ± 0%    64.37 ± 0%    +4.56% (p=0.002 n=6)
RandomPfx4Size/100_000     68.02 ± 0%    64.04 ± 0%    -5.85% (p=0.002 n=6)
RandomPfx4Size/200_000     60.81 ± 0%    64.02 ± 0%    +5.28% (p=0.002 n=6)
RandomPfx6Size/1_000       87.42 ± 2%    68.41 ± 3%   -21.75% (p=0.002 n=6)
RandomPfx6Size/10_000      67.78 ± 0%    64.37 ± 0%    -5.03% (p=0.002 n=6)
RandomPfx6Size/100_000     86.26 ± 0%    64.04 ± 0%   -25.76% (p=0.002 n=6)
RandomPfx6Size/200_000     81.52 ± 0%    64.02 ± 0%   -21.47% (p=0.002 n=6)
RandomPfxSize/1_000        84.43 ± 2%    68.16 ± 3%   -19.27% (p=0.002 n=6)
RandomPfxSize/10_000       59.68 ± 0%    64.37 ± 0%    +7.86% (p=0.002 n=6)
RandomPfxSize/100_000      67.88 ± 0%    64.04 ± 0%    -5.66% (p=0.002 n=6)
RandomPfxSize/200_000      64.58 ± 0%    64.02 ± 0%    -0.87% (p=0.002 n=6)
geomean                    63.55         65.20         +2.60%
```

## update, insert/delete

`bart` is by far the fastest algorithm for updates.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │              art/update.bm               │
                         │   sec/route    │    sec/route     vs base                 │
InsertRandomPfxs/1_000        30.43n ± 2%    269.95n ±   9%   +787.12% (p=0.002 n=6)
InsertRandomPfxs/10_000       41.37n ± 2%    663.40n ±  14%  +1503.38% (p=0.002 n=6)
InsertRandomPfxs/100_000      95.84n ± 3%   1750.00n ±  93%  +1725.86% (p=0.002 n=6)
InsertRandomPfxs/200_000      188.2n ± 1%    1865.5n ± 198%   +890.97% (p=0.002 n=6)
DeleteRandomPfxs/1_000        15.89n ± 1%     18.05n ±   5%    +13.59% (p=0.002 n=6)
DeleteRandomPfxs/10_000       15.58n ± 1%     16.77n ±   3%     +7.64% (p=0.002 n=6)
DeleteRandomPfxs/100_000      15.60n ± 1%     17.35n ±  12%    +11.18% (p=0.002 n=6)
DeleteRandomPfxs/200_000      16.75n ± 0%     19.99n ±  10%    +19.31% (p=0.002 n=6)
geomean                       33.18n          125.5n          +278.07%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           netipds/update.bm            │
                         │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000        30.43n ± 2%   164.95n ±  3%   +442.06% (p=0.002 n=6)
InsertRandomPfxs/10_000       41.37n ± 2%   263.55n ±  9%   +536.98% (p=0.002 n=6)
InsertRandomPfxs/100_000      95.84n ± 3%   581.25n ± 15%   +506.45% (p=0.002 n=6)
InsertRandomPfxs/200_000      188.2n ± 1%    702.8n ± 21%   +273.31% (p=0.002 n=6)
DeleteRandomPfxs/1_000        15.89n ± 1%   142.80n ± 22%   +798.40% (p=0.002 n=6)
DeleteRandomPfxs/10_000       15.58n ± 1%   209.95n ±  3%  +1247.56% (p=0.002 n=6)
DeleteRandomPfxs/100_000      15.60n ± 1%   494.45n ± 12%  +3068.54% (p=0.002 n=6)
DeleteRandomPfxs/200_000      16.75n ± 0%   702.10n ± 23%  +4091.64% (p=0.002 n=6)
geomean                       33.18n         341.5n         +928.99%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │          critbitgo/update.bm          │
                         │   sec/route    │   sec/route    vs base                │
InsertRandomPfxs/1_000        30.43n ± 2%   173.05n ±  5%  +468.68% (p=0.002 n=6)
InsertRandomPfxs/10_000       41.37n ± 2%   264.80n ±  7%  +540.00% (p=0.002 n=6)
InsertRandomPfxs/100_000      95.84n ± 3%   667.95n ± 12%  +596.91% (p=0.002 n=6)
InsertRandomPfxs/200_000      188.2n ± 1%    858.6n ±  4%  +356.10% (p=0.002 n=6)
DeleteRandomPfxs/1_000        15.89n ± 1%    75.28n ±  9%  +373.61% (p=0.002 n=6)
DeleteRandomPfxs/10_000       15.58n ± 1%    73.94n ± 27%  +374.55% (p=0.002 n=6)
DeleteRandomPfxs/100_000      15.60n ± 1%    79.25n ±  9%  +407.85% (p=0.002 n=6)
DeleteRandomPfxs/200_000      16.75n ± 0%    83.95n ±  4%  +401.22% (p=0.002 n=6)
geomean                       33.18n         177.2n        +434.06%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │            lpmtrie/update.bm            │
                         │   sec/route    │   sec/route     vs base                 │
InsertRandomPfxs/1_000        30.43n ± 2%    369.95n ±  6%  +1115.74% (p=0.002 n=6)
InsertRandomPfxs/10_000       41.37n ± 2%    532.85n ± 17%  +1187.85% (p=0.002 n=6)
InsertRandomPfxs/100_000      95.84n ± 3%   1262.00n ±  6%  +1216.71% (p=0.002 n=6)
InsertRandomPfxs/200_000      188.2n ± 1%    1324.0n ±  8%   +603.32% (p=0.002 n=6)
DeleteRandomPfxs/1_000        15.89n ± 1%     78.36n ±  9%   +392.99% (p=0.002 n=6)
DeleteRandomPfxs/10_000       15.58n ± 1%    164.70n ±  9%   +957.12% (p=0.002 n=6)
DeleteRandomPfxs/100_000      15.60n ± 1%    382.25n ± 14%  +2349.54% (p=0.002 n=6)
DeleteRandomPfxs/200_000      16.75n ± 0%    492.85n ±  4%  +2842.39% (p=0.002 n=6)
geomean                       33.18n          410.2n        +1135.98%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │            cidranger/update.bm            │
                         │   sec/route    │    sec/route     vs base                  │
InsertRandomPfxs/1_000        30.43n ± 2%    4237.00n ±  3%  +13823.76% (p=0.002 n=6)
InsertRandomPfxs/10_000       41.37n ± 2%    6404.50n ± 48%  +15379.15% (p=0.002 n=6)
InsertRandomPfxs/100_000      95.84n ± 3%   11578.00n ± 26%  +11979.92% (p=0.002 n=6)
InsertRandomPfxs/200_000      188.2n ± 1%    12429.5n ±  3%   +6502.66% (p=0.002 n=6)
DeleteRandomPfxs/1_000        15.89n ± 1%      87.94n ± 30%    +453.26% (p=0.002 n=6)
DeleteRandomPfxs/10_000       15.58n ± 1%     109.70n ± 15%    +604.11% (p=0.002 n=6)
DeleteRandomPfxs/100_000      15.60n ± 1%     161.95n ± 49%    +937.81% (p=0.002 n=6)
DeleteRandomPfxs/200_000      16.75n ± 0%     837.60n ± 15%   +4900.60% (p=0.002 n=6)
geomean                       33.18n           1.226µ         +3595.08%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │            cidrtree/update.bm             │
                         │   sec/route    │    sec/route     vs base                  │
InsertRandomPfxs/1_000        30.43n ± 2%   1360.00n ±   9%   +4369.27% (p=0.002 n=6)
InsertRandomPfxs/10_000       41.37n ± 2%   1900.50n ±  23%   +4493.35% (p=0.002 n=6)
InsertRandomPfxs/100_000      95.84n ± 3%   3537.00n ±  23%   +3590.33% (p=0.002 n=6)
InsertRandomPfxs/200_000      188.2n ± 1%    4494.0n ±  11%   +2287.25% (p=0.002 n=6)
DeleteRandomPfxs/1_000        15.89n ± 1%     14.72n ±  10%           ~ (p=0.310 n=6)
DeleteRandomPfxs/10_000       15.58n ± 1%     14.99n ±   8%           ~ (p=0.056 n=6)
DeleteRandomPfxs/100_000      15.60n ± 1%     27.52n ± 162%     +76.32% (p=0.002 n=6)
DeleteRandomPfxs/200_000      16.75n ± 0%   6181.50n ±  11%  +36804.48% (p=0.002 n=6)
geomean                       33.18n          445.1n          +1241.45%
```
