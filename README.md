# iprbench
comparing benchmarks for some golang IP routing table implementations:

```
	github.com/gaissmai/bart
	github.com/tailscale/art
	github.com/aromatt/netipds  <<< new kid around the block
	github.com/k-sone/critbitgo
	github.com/Asphaltt/lpmtrie
	github.com/yl2chen/cidranger
	github.com/gaissmai/cidrtree
```

The ~1_000_000 **Tier1** prefix test records (IPv4 and IPv6 routes) are from a full routing table with typical
ISP prefix distribution.

In comparison, the prefix lengths for the random test sets are equally distributed between /2-32 for IPv4
and /2-128 bits for IPv6, the randomly generated _default-routes_ with prefix length 0 have been sorted out,
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
                                     │  bart/lpm.bm  │              art/lpm.bm               │
                                     │    sec/op     │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4            37.48n ±   4%   44.94n ±  3%   +19.91% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            51.84n ±  17%   60.08n ±  5%   +15.92% (p=0.039 n=20)
LpmTier1Pfxs/RandomMissIP4             47.82n ±   7%   28.80n ±  1%   -39.77% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             50.61n ±  21%   36.17n ± 24%   -28.54% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     31.31n ±   9%   45.06n ±  1%   +43.92% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     24.51n ±  25%   45.69n ±  1%   +86.43% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      31.96n ±  12%   28.77n ±  0%         ~ (p=0.118 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      22.15n ±  19%   28.73n ±  0%   +29.73% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    36.86n ±   8%   45.04n ±  1%   +22.19% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    36.05n ±   9%   45.88n ±  1%   +27.24% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     37.98n ±   4%   28.76n ±  0%   -24.27% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     33.73n ±   5%   28.77n ±  0%   -14.73% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   24.26n ±  40%   53.19n ±  9%  +119.25% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   15.39n ± 135%   47.21n ± 20%  +206.66% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    42.67n ±  11%   28.98n ± 12%   -32.08% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    39.05n ±   9%   28.84n ±  6%   -26.16% (p=0.000 n=20)
geomean                                33.65n          37.77n         +12.24%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │            netipds/lpm.bm             │
                                     │    sec/op     │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4            37.48n ±   4%   62.60n ± 13%   +67.06% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            51.84n ±  17%   50.84n ± 11%         ~ (p=0.693 n=20)
LpmTier1Pfxs/RandomMissIP4             47.82n ±   7%   64.12n ±  3%   +34.08% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             50.61n ±  21%   54.68n ±  7%         ~ (p=0.174 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     31.31n ±   9%   50.27n ±  0%   +60.54% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     24.51n ±  25%   27.91n ±  6%   +13.92% (p=0.006 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      31.96n ±  12%   51.80n ±  4%   +62.09% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      22.15n ±  19%   28.86n ±  5%   +30.29% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    36.86n ±   8%   62.20n ±  3%   +68.75% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    36.05n ±   9%   32.66n ±  4%    -9.42% (p=0.001 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     37.98n ±   4%   63.62n ±  3%   +67.52% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     33.73n ±   5%   35.27n ±  3%    +4.56% (p=0.017 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   24.26n ±  40%   23.95n ±  0%         ~ (p=0.633 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   15.39n ± 135%   31.79n ± 33%  +106.53% (p=0.036 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    42.67n ±  11%   23.94n ±  0%   -43.91% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    39.05n ±   9%   40.79n ±  3%         ~ (p=0.365 n=20)
geomean                                33.65n          41.56n         +23.51%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │            critbitgo/lpm.bm             │
                                     │    sec/op     │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            37.48n ±   4%   159.90n ±  4%   +326.68% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            51.84n ±  17%   196.50n ± 14%   +279.09% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             47.82n ±   7%   658.60n ±  8%  +1277.10% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             50.61n ±  21%   508.40n ±  9%   +904.54% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     31.31n ±   9%    86.51n ±  3%   +176.30% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     24.51n ±  25%   105.55n ±  2%   +330.73% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      31.96n ±  12%   148.60n ± 20%   +364.96% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      22.15n ±  19%   157.85n ±  4%   +612.64% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    36.86n ±   8%    98.84n ±  2%   +168.16% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    36.05n ±   9%   118.60n ±  3%   +228.94% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     37.98n ±   4%   255.55n ± 13%   +572.94% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     33.73n ±   5%   210.70n ± 15%   +524.57% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   24.26n ±  40%   114.35n ±  9%   +371.35% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   15.39n ± 135%   139.10n ±  5%   +803.54% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    42.67n ±  11%   468.90n ± 21%   +998.90% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    39.05n ±   9%   357.35n ± 17%   +814.99% (p=0.000 n=20)
geomean                                33.65n           193.3n         +474.31%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │             lpmtrie/lpm.bm              │
                                     │    sec/op     │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            37.48n ±   4%   269.35n ±  8%   +618.75% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            51.84n ±  17%   325.20n ± 18%   +527.38% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             47.82n ±   7%   257.25n ±  9%   +437.90% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             50.61n ±  21%   254.45n ± 16%   +402.77% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     31.31n ±   9%   136.15n ±  3%   +334.85% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     24.51n ±  25%   103.65n ± 23%   +322.97% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      31.96n ±  12%   129.00n ±  2%   +303.63% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      22.15n ±  19%    96.52n ± 10%   +335.76% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    36.86n ±   8%   154.55n ± 13%   +319.29% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    36.05n ±   9%   142.40n ±  4%   +294.95% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     37.98n ±   4%   145.95n ± 19%   +284.33% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     33.73n ±   5%   136.35n ±  6%   +304.18% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   24.26n ±  40%   194.65n ±  6%   +702.35% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   15.39n ± 135%   174.65n ±  6%  +1034.46% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    42.67n ±  11%   205.75n ±  5%   +382.19% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    39.05n ±   9%   190.70n ±  5%   +388.29% (p=0.000 n=20)
geomean                                33.65n           172.0n         +411.20%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │            cidranger/lpm.bm            │
                                     │    sec/op     │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            37.48n ±   4%   276.45n ± 22%  +637.69% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            51.84n ±  17%   331.55n ± 25%  +539.63% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             47.82n ±   7%   272.10n ± 14%  +468.95% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             50.61n ±  21%   334.95n ± 18%  +561.83% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     31.31n ±   9%   113.90n ±  8%  +263.78% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     24.51n ±  25%   135.45n ±  8%  +452.74% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      31.96n ±  12%   115.75n ±  4%  +262.17% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      22.15n ±  19%   133.35n ±  7%  +502.03% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    36.86n ±   8%   151.65n ± 12%  +311.42% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    36.05n ±   9%   175.40n ±  7%  +386.48% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     37.98n ±   4%   149.35n ±  7%  +293.29% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     33.73n ±   5%   165.35n ±  5%  +390.14% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   24.26n ±  40%   108.75n ± 31%  +348.27% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   15.39n ± 135%   143.55n ± 36%  +832.45% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    42.67n ±  11%   190.05n ±  4%  +345.39% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    39.05n ±   9%   211.65n ±  6%  +441.93% (p=0.000 n=20)
geomean                                33.65n           175.6n        +421.93%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │             cidrtree/lpm.bm              │
                                     │    sec/op     │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            37.48n ±   4%    630.15n ±  6%  +1581.52% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            51.84n ±  17%    810.60n ± 18%  +1463.81% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             47.82n ±   7%   1148.00n ± 15%  +2300.42% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             50.61n ±  21%   1287.00n ± 19%  +2442.98% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     31.31n ±   9%    299.35n ± 15%   +856.08% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     24.51n ±  25%    352.70n ± 32%  +1339.30% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      31.96n ±  12%    474.85n ± 20%  +1385.76% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      22.15n ±  19%    487.20n ± 23%  +2099.55% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    36.86n ±   8%    483.35n ± 13%  +1211.31% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    36.05n ±   9%    536.55n ± 13%  +1388.14% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     37.98n ±   4%    683.70n ± 17%  +1700.39% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     33.73n ±   5%    728.70n ± 29%  +2060.07% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   24.26n ±  40%    594.50n ± 14%  +2350.54% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   15.39n ± 135%    690.10n ± 18%  +4382.62% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    42.67n ±  11%    819.65n ± 12%  +1820.90% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    39.05n ±   9%   1002.15n ± 27%  +2466.00% (p=0.000 n=20)
geomean                                33.65n            640.5n        +1803.32%
```

## size of the routing tables


`bart` has two orders of magnitude lower memory consumption compared to `art`
and is similar in low memory consumption to a binary search tree, like the `cidrtree` but
much faster in lookup times.


```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │              art/size.bm              │
                       │ bytes/route  │ bytes/route   vs base                 │
Tier1PfxSize/1_000         107.5 ± 2%    7591.0 ± 0%  +6961.40% (p=0.002 n=6)
Tier1PfxSize/10_000        79.57 ± 0%   4889.00 ± 0%  +6044.28% (p=0.002 n=6)
Tier1PfxSize/100_000       35.74 ± 0%   1669.00 ± 0%  +4569.84% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%   1098.00 ± 0%  +4448.47% (p=0.002 n=6)
RandomPfx4Size/1_000       101.8 ± 2%    7394.0 ± 0%  +7163.26% (p=0.002 n=6)
RandomPfx4Size/10_000      76.33 ± 0%   6043.00 ± 0%  +7816.94% (p=0.002 n=6)
RandomPfx4Size/100_000     81.79 ± 0%   5455.00 ± 0%  +6569.52% (p=0.002 n=6)
RandomPfx4Size/200_000     72.85 ± 0%   4832.00 ± 0%  +6532.81% (p=0.002 n=6)
RandomPfx6Size/1_000       106.6 ± 2%    7845.0 ± 0%  +7259.29% (p=0.002 n=6)
RandomPfx6Size/10_000      84.96 ± 0%   6829.00 ± 0%  +7937.90% (p=0.002 n=6)
RandomPfx6Size/100_000     105.9 ± 0%    7837.0 ± 0%  +7300.38% (p=0.002 n=6)
RandomPfx6Size/200_000     101.0 ± 0%    7602.0 ± 0%  +7426.73% (p=0.002 n=6)
RandomPfxSize/1_000        104.6 ± 2%    7564.0 ± 0%  +7131.36% (p=0.002 n=6)
RandomPfxSize/10_000       76.15 ± 0%   6226.00 ± 0%  +8075.97% (p=0.002 n=6)
RandomPfxSize/100_000      82.13 ± 0%   5792.00 ± 0%  +6952.23% (p=0.002 n=6)
RandomPfxSize/200_000      78.33 ± 0%   5454.00 ± 0%  +6862.85% (p=0.002 n=6)
geomean                    77.51        5.170Ki       +6729.99%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │           netipds/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         107.5 ± 2%    101.0 ± 2%    -6.05% (p=0.002 n=6)
Tier1PfxSize/10_000        79.57 ± 0%    96.08 ± 0%   +20.75% (p=0.002 n=6)
Tier1PfxSize/100_000       35.74 ± 0%    94.34 ± 0%  +163.96% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%    93.27 ± 0%  +286.37% (p=0.002 n=6)
RandomPfx4Size/1_000      101.80 ± 2%    99.73 ± 2%    -2.03% (p=0.002 n=6)
RandomPfx4Size/10_000      76.33 ± 0%    94.19 ± 0%   +23.40% (p=0.002 n=6)
RandomPfx4Size/100_000     81.79 ± 0%    86.51 ± 0%    +5.77% (p=0.002 n=6)
RandomPfx4Size/200_000     72.85 ± 0%    82.89 ± 0%   +13.78% (p=0.002 n=6)
RandomPfx6Size/1_000      106.60 ± 2%    99.93 ± 2%    -6.26% (p=0.002 n=6)
RandomPfx6Size/10_000      84.96 ± 0%    95.07 ± 0%   +11.90% (p=0.002 n=6)
RandomPfx6Size/100_000    105.90 ± 0%    92.64 ± 0%   -12.52% (p=0.002 n=6)
RandomPfx6Size/200_000    101.00 ± 0%    91.94 ± 0%    -8.97% (p=0.002 n=6)
RandomPfxSize/1_000       104.60 ± 2%    99.82 ± 2%    -4.57% (p=0.002 n=6)
RandomPfxSize/10_000       76.15 ± 0%    94.22 ± 0%   +23.73% (p=0.002 n=6)
RandomPfxSize/100_000      82.13 ± 0%    87.84 ± 0%    +6.95% (p=0.002 n=6)
RandomPfxSize/200_000      78.33 ± 0%    84.76 ± 0%    +8.21% (p=0.002 n=6)
geomean                    77.51         93.23        +20.29%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          critbitgo/size.bm          │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         107.5 ± 2%    119.5 ± 2%   +11.16% (p=0.002 n=6)
Tier1PfxSize/10_000        79.57 ± 0%   114.70 ± 0%   +44.15% (p=0.002 n=6)
Tier1PfxSize/100_000       35.74 ± 0%   114.40 ± 0%  +220.09% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%   114.40 ± 0%  +373.90% (p=0.002 n=6)
RandomPfx4Size/1_000       101.8 ± 2%    116.1 ± 2%   +14.05% (p=0.002 n=6)
RandomPfx4Size/10_000      76.33 ± 0%   112.40 ± 0%   +47.26% (p=0.002 n=6)
RandomPfx4Size/100_000     81.79 ± 0%   112.00 ± 0%   +36.94% (p=0.002 n=6)
RandomPfx4Size/200_000     72.85 ± 0%   112.00 ± 0%   +53.74% (p=0.002 n=6)
RandomPfx6Size/1_000       106.6 ± 2%    132.4 ± 1%   +24.20% (p=0.002 n=6)
RandomPfx6Size/10_000      84.96 ± 0%   128.40 ± 0%   +51.13% (p=0.002 n=6)
RandomPfx6Size/100_000     105.9 ± 0%    128.0 ± 0%   +20.87% (p=0.002 n=6)
RandomPfx6Size/200_000     101.0 ± 0%    128.0 ± 0%   +26.73% (p=0.002 n=6)
RandomPfxSize/1_000        104.6 ± 2%    119.1 ± 2%   +13.86% (p=0.002 n=6)
RandomPfxSize/10_000       76.15 ± 0%   115.50 ± 0%   +51.67% (p=0.002 n=6)
RandomPfxSize/100_000      82.13 ± 0%   115.20 ± 0%   +40.27% (p=0.002 n=6)
RandomPfxSize/200_000      78.33 ± 0%   115.20 ± 0%   +47.07% (p=0.002 n=6)
geomean                    77.51         118.4        +52.77%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │           lpmtrie/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         107.5 ± 2%    215.4 ± 5%  +100.37% (p=0.002 n=6)
Tier1PfxSize/10_000        79.57 ± 0%   210.50 ± 5%  +164.55% (p=0.002 n=6)
Tier1PfxSize/100_000       35.74 ± 0%   209.90 ± 5%  +487.30% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%   209.20 ± 5%  +766.61% (p=0.002 n=6)
RandomPfx4Size/1_000       101.8 ± 2%    212.0 ± 5%  +108.25% (p=0.002 n=6)
RandomPfx4Size/10_000      76.33 ± 0%   206.70 ± 5%  +170.80% (p=0.002 n=6)
RandomPfx4Size/100_000     81.79 ± 0%   199.40 ± 6%  +143.80% (p=0.002 n=6)
RandomPfx4Size/200_000     72.85 ± 0%   194.70 ± 7%  +167.26% (p=0.002 n=6)
RandomPfx6Size/1_000       106.6 ± 2%    228.3 ± 8%  +114.17% (p=0.002 n=6)
RandomPfx6Size/10_000      84.96 ± 0%   223.00 ± 7%  +162.48% (p=0.002 n=6)
RandomPfx6Size/100_000     105.9 ± 0%    219.3 ± 8%  +107.08% (p=0.002 n=6)
RandomPfx6Size/200_000     101.0 ± 0%    217.8 ± 8%  +115.64% (p=0.002 n=6)
RandomPfxSize/1_000        104.6 ± 2%    215.3 ± 6%  +105.83% (p=0.002 n=6)
RandomPfxSize/10_000       76.15 ± 0%   209.90 ± 5%  +175.64% (p=0.002 n=6)
RandomPfxSize/100_000      82.13 ± 0%   203.50 ± 7%  +147.78% (p=0.002 n=6)
RandomPfxSize/200_000      78.33 ± 0%   199.50 ± 7%  +154.69% (p=0.002 n=6)
geomean                    77.51         210.7       +171.87%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidranger/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                 │
Tier1PfxSize/1_000         107.5 ± 2%    539.6 ± 3%   +401.95% (p=0.002 n=6)
Tier1PfxSize/10_000        79.57 ± 0%   533.90 ± 3%   +570.98% (p=0.002 n=6)
Tier1PfxSize/100_000       35.74 ± 0%   527.20 ± 2%  +1375.10% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%   522.20 ± 2%  +2063.21% (p=0.002 n=6)
RandomPfx4Size/1_000       101.8 ± 2%    527.2 ± 3%   +417.88% (p=0.002 n=6)
RandomPfx4Size/10_000      76.33 ± 0%   514.60 ± 3%   +574.18% (p=0.002 n=6)
RandomPfx4Size/100_000     81.79 ± 0%   479.80 ± 3%   +486.62% (p=0.002 n=6)
RandomPfx4Size/200_000     72.85 ± 0%   463.30 ± 3%   +535.96% (p=0.002 n=6)
RandomPfx6Size/1_000       106.6 ± 2%    595.1 ± 0%   +458.26% (p=0.002 n=6)
RandomPfx6Size/10_000      84.96 ± 0%   584.20 ± 0%   +587.62% (p=0.002 n=6)
RandomPfx6Size/100_000     105.9 ± 0%    573.8 ± 0%   +441.83% (p=0.002 n=6)
RandomPfx6Size/200_000     101.0 ± 0%    569.6 ± 0%   +463.96% (p=0.002 n=6)
RandomPfxSize/1_000        104.6 ± 2%    536.9 ± 2%   +413.29% (p=0.002 n=6)
RandomPfxSize/10_000       76.15 ± 0%   526.80 ± 2%   +591.79% (p=0.002 n=6)
RandomPfxSize/100_000      82.13 ± 0%   498.00 ± 2%   +506.36% (p=0.002 n=6)
RandomPfxSize/200_000      78.33 ± 0%   484.30 ± 2%   +518.28% (p=0.002 n=6)
geomean                    77.51         528.5        +581.91%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidrtree/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000        107.50 ± 2%    69.26 ± 3%   -35.57% (p=0.002 n=6)
Tier1PfxSize/10_000        79.57 ± 0%    64.37 ± 0%   -19.10% (p=0.002 n=6)
Tier1PfxSize/100_000       35.74 ± 0%    64.04 ± 0%   +79.18% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%    64.02 ± 0%  +165.20% (p=0.002 n=6)
RandomPfx4Size/1_000      101.80 ± 2%    68.16 ± 3%   -33.05% (p=0.002 n=6)
RandomPfx4Size/10_000      76.33 ± 0%    64.37 ± 0%   -15.67% (p=0.002 n=6)
RandomPfx4Size/100_000     81.79 ± 0%    64.04 ± 0%   -21.70% (p=0.002 n=6)
RandomPfx4Size/200_000     72.85 ± 0%    64.02 ± 0%   -12.12% (p=0.002 n=6)
RandomPfx6Size/1_000      106.60 ± 2%    68.41 ± 3%   -35.83% (p=0.002 n=6)
RandomPfx6Size/10_000      84.96 ± 0%    64.37 ± 0%   -24.23% (p=0.002 n=6)
RandomPfx6Size/100_000    105.90 ± 0%    64.04 ± 0%   -39.53% (p=0.002 n=6)
RandomPfx6Size/200_000    101.00 ± 0%    64.02 ± 0%   -36.61% (p=0.002 n=6)
RandomPfxSize/1_000       104.60 ± 2%    68.16 ± 3%   -34.84% (p=0.002 n=6)
RandomPfxSize/10_000       76.15 ± 0%    64.37 ± 0%   -15.47% (p=0.002 n=6)
RandomPfxSize/100_000      82.13 ± 0%    64.04 ± 0%   -22.03% (p=0.002 n=6)
RandomPfxSize/200_000      78.33 ± 0%    64.02 ± 0%   -18.27% (p=0.002 n=6)
geomean                    77.51         65.20        -15.87%
```

## update, insert/delete

`bart` is by far the fastest algorithm for inserts and one of the fastest for delete.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │              art/update.bm              │
                         │   sec/route    │   sec/route     vs base                 │
InsertRandomPfxs/1_000        44.37n ± 3%   269.95n ±   9%   +508.34% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.83n ± 3%   663.40n ±  14%  +1110.03% (p=0.002 n=6)
InsertRandomPfxs/100_000      132.1n ± 3%   1750.0n ±  93%  +1224.75% (p=0.002 n=6)
InsertRandomPfxs/200_000      228.0n ± 1%   1865.5n ± 198%   +718.02% (p=0.002 n=6)
DeleteRandomPfxs/1_000        18.99n ± 1%    18.05n ±   5%     -4.92% (p=0.011 n=6)
DeleteRandomPfxs/10_000       18.69n ± 1%    16.77n ±   3%    -10.27% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.91n ± 1%    17.35n ±  12%          ~ (p=0.054 n=6)
DeleteRandomPfxs/200_000      20.38n ± 0%    19.99n ±  10%          ~ (p=0.372 n=6)
geomean                       42.18n         125.5n          +197.41%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           netipds/update.bm            │
                         │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000        44.37n ± 3%   164.95n ±  3%   +271.72% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.83n ± 3%   263.55n ±  9%   +380.71% (p=0.002 n=6)
InsertRandomPfxs/100_000      132.1n ± 3%    581.2n ± 15%   +340.01% (p=0.002 n=6)
InsertRandomPfxs/200_000      228.0n ± 1%    702.8n ± 21%   +208.16% (p=0.002 n=6)
DeleteRandomPfxs/1_000        18.99n ± 1%   142.80n ± 22%   +651.97% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.69n ± 1%   209.95n ±  3%  +1023.33% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.91n ± 1%   494.45n ± 12%  +2514.06% (p=0.002 n=6)
DeleteRandomPfxs/200_000      20.38n ± 0%   702.10n ± 23%  +3345.04% (p=0.002 n=6)
geomean                       42.18n         341.5n         +709.46%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │          critbitgo/update.bm          │
                         │   sec/route    │   sec/route    vs base                │
InsertRandomPfxs/1_000        44.37n ± 3%   173.05n ±  5%  +289.97% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.83n ± 3%   264.80n ±  7%  +382.99% (p=0.002 n=6)
InsertRandomPfxs/100_000      132.1n ± 3%    668.0n ± 12%  +405.64% (p=0.002 n=6)
InsertRandomPfxs/200_000      228.0n ± 1%    858.6n ±  4%  +276.50% (p=0.002 n=6)
DeleteRandomPfxs/1_000        18.99n ± 1%    75.28n ±  9%  +296.42% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.69n ± 1%    73.94n ± 27%  +295.59% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.91n ± 1%    79.25n ±  9%  +318.98% (p=0.002 n=6)
DeleteRandomPfxs/200_000      20.38n ± 0%    83.95n ±  4%  +311.95% (p=0.002 n=6)
geomean                       42.18n         177.2n        +320.12%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           lpmtrie/update.bm            │
                         │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000        44.37n ± 3%   369.95n ±  6%   +733.69% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.83n ± 3%   532.85n ± 17%   +871.91% (p=0.002 n=6)
InsertRandomPfxs/100_000      132.1n ± 3%   1262.0n ±  6%   +855.34% (p=0.002 n=6)
InsertRandomPfxs/200_000      228.0n ± 1%   1324.0n ±  8%   +480.57% (p=0.002 n=6)
DeleteRandomPfxs/1_000        18.99n ± 1%    78.36n ±  9%   +312.64% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.69n ± 1%   164.70n ±  9%   +781.22% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.91n ± 1%   382.25n ± 14%  +1920.88% (p=0.002 n=6)
DeleteRandomPfxs/200_000      20.38n ± 0%   492.85n ±  4%  +2318.30% (p=0.002 n=6)
geomean                       42.18n         410.2n         +872.29%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           cidranger/update.bm            │
                         │   sec/route    │   sec/route     vs base                  │
InsertRandomPfxs/1_000        44.37n ± 3%   4237.00n ±  3%   +9448.17% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.83n ± 3%   6404.50n ± 48%  +11581.71% (p=0.002 n=6)
InsertRandomPfxs/100_000      132.1n ± 3%   11578.0n ± 26%   +8664.57% (p=0.002 n=6)
InsertRandomPfxs/200_000      228.0n ± 1%   12429.5n ±  3%   +5350.34% (p=0.002 n=6)
DeleteRandomPfxs/1_000        18.99n ± 1%     87.94n ± 30%    +363.09% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.69n ± 1%    109.70n ± 15%    +486.94% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.91n ± 1%    161.95n ± 49%    +756.20% (p=0.002 n=6)
DeleteRandomPfxs/200_000      20.38n ± 0%    837.60n ± 15%   +4009.91% (p=0.002 n=6)
geomean                       42.18n          1.226µ         +2806.74%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │            cidrtree/update.bm             │
                         │   sec/route    │    sec/route     vs base                  │
InsertRandomPfxs/1_000        44.37n ± 3%   1360.00n ±   9%   +2964.79% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.83n ± 3%   1900.50n ±  23%   +3366.48% (p=0.002 n=6)
InsertRandomPfxs/100_000      132.1n ± 3%    3537.0n ±  23%   +2577.52% (p=0.002 n=6)
InsertRandomPfxs/200_000      228.0n ± 1%    4494.0n ±  11%   +1870.62% (p=0.002 n=6)
DeleteRandomPfxs/1_000        18.99n ± 1%     14.72n ±  10%     -22.46% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.69n ± 1%     14.99n ±   8%     -19.82% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.91n ± 1%     27.52n ± 162%     +45.47% (p=0.002 n=6)
DeleteRandomPfxs/200_000      20.38n ± 0%   6181.50n ±  11%  +30231.21% (p=0.002 n=6)
geomean                       42.18n          445.1n           +955.25%
```
