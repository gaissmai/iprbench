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
                                     │  bart/lpm.bm  │               art/lpm.bm               │
                                     │    sec/op     │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            26.18n ±  22%    40.90n ±  5%   +56.26% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            24.18n ±  36%    46.69n ±  9%   +93.11% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             26.84n ±  54%    25.49n ±  3%         ~ (p=0.615 n=20)
LpmTier1Pfxs/RandomMissIP6             6.322n ±   4%   25.465n ±  0%  +302.80% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     15.92n ±   4%    40.15n ±  1%  +152.09% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     18.11n ±   9%    40.71n ±  0%  +124.70% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      17.62n ±  14%    25.45n ±  0%   +44.43% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      16.00n ±  58%    25.46n ±  0%   +59.16% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    24.19n ±   7%    40.52n ±  1%   +67.53% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    20.15n ±   3%    41.90n ±  2%  +107.94% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     26.48n ±   6%    25.48n ±  0%         ~ (p=0.053 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.86n ±  15%    25.46n ±  0%   +28.20% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   10.55n ± 134%    42.26n ± 15%  +300.52% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   21.38n ±  53%    43.39n ± 14%  +103.02% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    38.10n ±   2%    25.50n ±  3%   -33.09% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    27.50n ±  21%    25.58n ±  4%    -7.00% (p=0.011 n=20)
geomean                                19.73n           32.72n         +65.83%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │             netipds/lpm.bm             │
                                     │    sec/op     │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            26.18n ±  22%    47.72n ±  9%   +82.33% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            24.18n ±  36%    36.58n ± 20%   +51.31% (p=0.001 n=20)
LpmTier1Pfxs/RandomMissIP4             26.84n ±  54%    55.55n ± 10%  +106.95% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             6.322n ±   4%   30.390n ± 52%  +380.70% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     15.92n ±   4%    41.66n ±  0%  +161.63% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     18.11n ±   9%    27.88n ±  6%   +53.91% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      17.62n ±  14%    49.32n ±  4%  +179.83% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      16.00n ±  58%    27.36n ±  4%   +71.03% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    24.19n ±   7%    59.64n ±  4%  +146.57% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    20.15n ±   3%    25.92n ± 17%   +28.64% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     26.48n ±   6%    62.11n ±  2%  +134.55% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.86n ±  15%    33.56n ±  3%   +68.98% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   10.55n ± 134%    57.67n ± 13%  +446.64% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   21.38n ±  53%    21.02n ± 16%         ~ (p=0.172 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    38.10n ±   2%    74.70n ±  3%   +96.04% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    27.50n ±  21%    39.27n ±  2%   +42.76% (p=0.000 n=20)
geomean                                19.73n           40.53n        +105.36%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │             critbitgo/lpm.bm             │
                                     │    sec/op     │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            26.18n ±  22%    117.75n ± 22%   +349.86% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            24.18n ±  36%    152.45n ± 14%   +530.61% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             26.84n ±  54%    617.85n ± 16%  +2201.97% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             6.322n ±   4%   281.700n ± 29%  +4355.87% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     15.92n ±   4%     88.34n ±  3%   +454.73% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     18.11n ±   9%    107.10n ±  1%   +491.22% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      17.62n ±  14%    154.25n ± 16%   +775.18% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      16.00n ±  58%    147.30n ±  7%   +820.62% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    24.19n ±   7%    108.40n ± 20%   +348.12% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    20.15n ±   3%    129.70n ± 14%   +543.67% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     26.48n ±   6%    247.75n ± 14%   +835.61% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.86n ±  15%    231.05n ± 14%  +1063.39% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   10.55n ± 134%    154.15n ± 60%  +1361.14% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   21.38n ±  53%    175.25n ± 10%   +719.88% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    38.10n ±   2%    447.70n ± 17%  +1074.91% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    27.50n ±  21%    365.20n ± 14%  +1227.76% (p=0.000 n=20)
geomean                                19.73n            187.8n         +851.47%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │              lpmtrie/lpm.bm              │
                                     │    sec/op     │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            26.18n ±  22%   207.70n ±   5%   +693.51% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            24.18n ±  36%   138.05n ±  44%   +471.04% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             26.84n ±  54%    30.58n ± 361%    +13.95% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             6.322n ±   4%   12.380n ±   0%    +95.82% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     15.92n ±   4%   130.00n ±   4%   +716.33% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     18.11n ±   9%   109.10n ±   5%   +502.26% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      17.62n ±  14%   102.90n ±  23%   +483.83% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      16.00n ±  58%    93.55n ±  11%   +484.69% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    24.19n ±   7%   151.45n ±   4%   +526.09% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    20.15n ±   3%   125.15n ±   8%   +521.09% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     26.48n ±   6%   136.35n ±   6%   +414.92% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.86n ±  15%   123.70n ±   6%   +522.86% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   10.55n ± 134%   177.80n ±   6%  +1585.31% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   21.38n ±  53%   162.50n ±   6%   +660.23% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    38.10n ±   2%   195.95n ±   5%   +414.24% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    27.50n ±  21%   180.45n ±  15%   +556.06% (p=0.000 n=20)
geomean                                19.73n           110.4n          +459.22%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │            cidranger/lpm.bm             │
                                     │    sec/op     │     sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            26.18n ±  22%   175.55n ±  13%  +570.68% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            24.18n ±  36%   180.35n ±  24%  +646.02% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             26.84n ±  54%    46.30n ± 259%   +72.50% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             6.322n ±   4%   49.305n ±   9%  +679.90% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     15.92n ±   4%    84.13n ±  25%  +428.29% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     18.11n ±   9%   135.00n ±   2%  +645.24% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      17.62n ±  14%   108.65n ±   9%  +516.45% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      16.00n ±  58%   125.65n ±   9%  +685.31% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    24.19n ±   7%    81.77n ±  49%  +238.01% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    20.15n ±   3%   119.15n ±  38%  +491.32% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     26.48n ±   6%   143.40n ±   7%  +441.54% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.86n ±  15%   156.80n ±   3%  +689.53% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   10.55n ± 134%   114.65n ±  38%  +986.73% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   21.38n ±  53%   133.95n ±  17%  +526.67% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    38.10n ±   2%   185.20n ±   3%  +386.03% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    27.50n ±  21%   195.10n ±  10%  +609.33% (p=0.000 n=20)
geomean                                19.73n           117.7n         +496.57%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │             cidrtree/lpm.bm              │
                                     │    sec/op     │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            26.18n ±  22%    789.75n ± 12%  +2917.19% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            24.18n ±  36%    900.05n ±  7%  +3623.06% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             26.84n ±  54%    860.45n ± 97%  +3105.85% (p=0.001 n=20)
LpmTier1Pfxs/RandomMissIP6             6.322n ±   4%    39.045n ±  1%   +517.61% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     15.92n ±   4%    408.65n ±  5%  +2466.09% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     18.11n ±   9%    199.95n ± 76%  +1003.78% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      17.62n ±  14%    490.80n ± 18%  +2684.68% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      16.00n ±  58%    594.30n ± 12%  +3614.37% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    24.19n ±   7%    589.90n ± 26%  +2338.61% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    20.15n ±   3%    745.55n ± 15%  +3600.00% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     26.48n ±   6%    753.85n ± 17%  +2746.87% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     19.86n ±  15%    875.75n ± 22%  +4309.62% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   10.55n ± 134%    841.65n ± 32%  +7877.73% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   21.38n ±  53%   1275.00n ± 13%  +5864.91% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    38.10n ±   2%    972.05n ± 16%  +2450.98% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    27.50n ±  21%   1159.00n ±  9%  +4113.78% (p=0.000 n=20)
geomean                                19.73n            587.2n        +2875.49%
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
RandomPfx4Size/1_000       103.4 ± 2%    7413.0 ± 0%  +7069.25% (p=0.002 n=6)
RandomPfx4Size/10_000      75.57 ± 0%   6062.00 ± 0%  +7921.70% (p=0.002 n=6)
RandomPfx4Size/100_000     81.91 ± 0%   5480.00 ± 0%  +6590.27% (p=0.002 n=6)
RandomPfx4Size/200_000     72.92 ± 0%   4829.00 ± 0%  +6522.33% (p=0.002 n=6)
RandomPfx6Size/1_000       107.0 ± 2%    7799.0 ± 0%  +7188.79% (p=0.002 n=6)
RandomPfx6Size/10_000      84.79 ± 0%   6797.00 ± 0%  +7916.28% (p=0.002 n=6)
RandomPfx6Size/100_000     105.9 ± 0%    7829.0 ± 0%  +7292.82% (p=0.002 n=6)
RandomPfx6Size/200_000     101.0 ± 0%    7598.0 ± 0%  +7422.77% (p=0.002 n=6)
RandomPfxSize/1_000        105.2 ± 2%    7609.0 ± 0%  +7132.89% (p=0.002 n=6)
RandomPfxSize/10_000       76.12 ± 0%   6232.00 ± 0%  +8087.07% (p=0.002 n=6)
RandomPfxSize/100_000      82.42 ± 0%   5801.00 ± 0%  +6938.34% (p=0.002 n=6)
RandomPfxSize/200_000      78.45 ± 0%   5447.00 ± 0%  +6843.28% (p=0.002 n=6)
geomean                    77.60        5.171Ki       +6723.40%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │           netipds/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         107.5 ± 2%    101.0 ± 2%    -6.05% (p=0.002 n=6)
Tier1PfxSize/10_000        79.57 ± 0%    96.08 ± 0%   +20.75% (p=0.002 n=6)
Tier1PfxSize/100_000       35.74 ± 0%    94.34 ± 0%  +163.96% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%    93.27 ± 0%  +286.37% (p=0.002 n=6)
RandomPfx4Size/1_000      103.40 ± 2%    99.49 ± 2%    -3.78% (p=0.002 n=6)
RandomPfx4Size/10_000      75.57 ± 0%    94.21 ± 0%   +24.67% (p=0.002 n=6)
RandomPfx4Size/100_000     81.91 ± 0%    86.44 ± 0%    +5.53% (p=0.002 n=6)
RandomPfx4Size/200_000     72.92 ± 0%    82.81 ± 0%   +13.56% (p=0.002 n=6)
RandomPfx6Size/1_000       107.0 ± 2%    100.3 ± 2%    -6.26% (p=0.002 n=6)
RandomPfx6Size/10_000      84.79 ± 0%    95.19 ± 0%   +12.27% (p=0.002 n=6)
RandomPfx6Size/100_000    105.90 ± 0%    92.67 ± 0%   -12.49% (p=0.002 n=6)
RandomPfx6Size/200_000    101.00 ± 0%    91.84 ± 0%    -9.07% (p=0.002 n=6)
RandomPfxSize/1_000       105.20 ± 2%    99.49 ± 2%    -5.43% (p=0.002 n=6)
RandomPfxSize/10_000       76.12 ± 0%    94.29 ± 0%   +23.87% (p=0.002 n=6)
RandomPfxSize/100_000      82.42 ± 0%    87.87 ± 0%    +6.61% (p=0.002 n=6)
RandomPfxSize/200_000      78.45 ± 0%    84.78 ± 0%    +8.07% (p=0.002 n=6)
geomean                    77.60         93.22        +20.12%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          critbitgo/size.bm          │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         107.5 ± 2%    119.5 ± 2%   +11.16% (p=0.002 n=6)
Tier1PfxSize/10_000        79.57 ± 0%   114.70 ± 0%   +44.15% (p=0.002 n=6)
Tier1PfxSize/100_000       35.74 ± 0%   114.40 ± 0%  +220.09% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%   114.40 ± 0%  +373.90% (p=0.002 n=6)
RandomPfx4Size/1_000       103.4 ± 2%    116.1 ± 2%   +12.28% (p=0.002 n=6)
RandomPfx4Size/10_000      75.57 ± 0%   112.40 ± 0%   +48.74% (p=0.002 n=6)
RandomPfx4Size/100_000     81.91 ± 0%   112.00 ± 0%   +36.74% (p=0.002 n=6)
RandomPfx4Size/200_000     72.92 ± 0%   112.00 ± 0%   +53.59% (p=0.002 n=6)
RandomPfx6Size/1_000       107.0 ± 2%    132.4 ± 1%   +23.74% (p=0.002 n=6)
RandomPfx6Size/10_000      84.79 ± 0%   128.40 ± 0%   +51.43% (p=0.002 n=6)
RandomPfx6Size/100_000     105.9 ± 0%    128.0 ± 0%   +20.87% (p=0.002 n=6)
RandomPfx6Size/200_000     101.0 ± 0%    128.0 ± 0%   +26.73% (p=0.002 n=6)
RandomPfxSize/1_000        105.2 ± 2%    119.1 ± 2%   +13.21% (p=0.002 n=6)
RandomPfxSize/10_000       76.12 ± 0%   115.50 ± 0%   +51.73% (p=0.002 n=6)
RandomPfxSize/100_000      82.42 ± 0%   115.20 ± 0%   +39.77% (p=0.002 n=6)
RandomPfxSize/200_000      78.45 ± 0%   115.20 ± 0%   +46.85% (p=0.002 n=6)
geomean                    77.60         118.4        +52.58%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │           lpmtrie/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         107.5 ± 2%    215.4 ± 5%  +100.37% (p=0.002 n=6)
Tier1PfxSize/10_000        79.57 ± 0%   210.50 ± 5%  +164.55% (p=0.002 n=6)
Tier1PfxSize/100_000       35.74 ± 0%   209.90 ± 5%  +487.30% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%   209.20 ± 5%  +766.61% (p=0.002 n=6)
RandomPfx4Size/1_000       103.4 ± 2%    212.2 ± 5%  +105.22% (p=0.002 n=6)
RandomPfx4Size/10_000      75.57 ± 0%   207.00 ± 4%  +173.92% (p=0.002 n=6)
RandomPfx4Size/100_000     81.91 ± 0%   199.40 ± 6%  +143.44% (p=0.002 n=6)
RandomPfx4Size/200_000     72.92 ± 0%   194.70 ± 7%  +167.00% (p=0.002 n=6)
RandomPfx6Size/1_000       107.0 ± 2%    228.1 ± 8%  +113.18% (p=0.002 n=6)
RandomPfx6Size/10_000      84.79 ± 0%   223.20 ± 7%  +163.24% (p=0.002 n=6)
RandomPfx6Size/100_000     105.9 ± 0%    219.2 ± 8%  +106.99% (p=0.002 n=6)
RandomPfx6Size/200_000     101.0 ± 0%    217.9 ± 8%  +115.74% (p=0.002 n=6)
RandomPfxSize/1_000        105.2 ± 2%    215.0 ± 5%  +104.37% (p=0.002 n=6)
RandomPfxSize/10_000       76.12 ± 0%   210.00 ± 5%  +175.88% (p=0.002 n=6)
RandomPfxSize/100_000      82.42 ± 0%   203.30 ± 7%  +146.66% (p=0.002 n=6)
RandomPfxSize/200_000      78.45 ± 0%   199.50 ± 7%  +154.30% (p=0.002 n=6)
geomean                    77.60         210.7       +171.54%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidranger/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                 │
Tier1PfxSize/1_000         107.5 ± 2%    539.6 ± 3%   +401.95% (p=0.002 n=6)
Tier1PfxSize/10_000        79.57 ± 0%   533.90 ± 3%   +570.98% (p=0.002 n=6)
Tier1PfxSize/100_000       35.74 ± 0%   527.20 ± 2%  +1375.10% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%   522.20 ± 2%  +2063.21% (p=0.002 n=6)
RandomPfx4Size/1_000       103.4 ± 2%    525.8 ± 3%   +408.51% (p=0.002 n=6)
RandomPfx4Size/10_000      75.57 ± 0%   513.30 ± 3%   +579.24% (p=0.002 n=6)
RandomPfx4Size/100_000     81.91 ± 0%   479.90 ± 3%   +485.89% (p=0.002 n=6)
RandomPfx4Size/200_000     72.92 ± 0%   463.00 ± 3%   +534.94% (p=0.002 n=6)
RandomPfx6Size/1_000       107.0 ± 2%    594.1 ± 0%   +455.23% (p=0.002 n=6)
RandomPfx6Size/10_000      84.79 ± 0%   585.30 ± 0%   +590.29% (p=0.002 n=6)
RandomPfx6Size/100_000     105.9 ± 0%    574.0 ± 0%   +442.02% (p=0.002 n=6)
RandomPfx6Size/200_000     101.0 ± 0%    570.0 ± 0%   +464.36% (p=0.002 n=6)
RandomPfxSize/1_000        105.2 ± 2%    535.7 ± 2%   +409.22% (p=0.002 n=6)
RandomPfxSize/10_000       76.12 ± 0%   527.20 ± 2%   +592.59% (p=0.002 n=6)
RandomPfxSize/100_000      82.42 ± 0%   498.50 ± 2%   +504.83% (p=0.002 n=6)
RandomPfxSize/200_000      78.45 ± 0%   484.40 ± 2%   +517.46% (p=0.002 n=6)
geomean                    77.60         528.4        +580.85%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidrtree/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000        107.50 ± 2%    69.17 ± 3%   -35.66% (p=0.002 n=6)
Tier1PfxSize/10_000        79.57 ± 0%    64.37 ± 0%   -19.10% (p=0.002 n=6)
Tier1PfxSize/100_000       35.74 ± 0%    64.04 ± 0%   +79.18% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%    64.02 ± 0%  +165.20% (p=0.002 n=6)
RandomPfx4Size/1_000      103.40 ± 2%    68.16 ± 3%   -34.08% (p=0.002 n=6)
RandomPfx4Size/10_000      75.57 ± 0%    64.37 ± 0%   -14.82% (p=0.002 n=6)
RandomPfx4Size/100_000     81.91 ± 0%    64.04 ± 0%   -21.82% (p=0.002 n=6)
RandomPfx4Size/200_000     72.92 ± 0%    64.02 ± 0%   -12.21% (p=0.002 n=6)
RandomPfx6Size/1_000      107.00 ± 2%    68.41 ± 3%   -36.07% (p=0.002 n=6)
RandomPfx6Size/10_000      84.79 ± 0%    64.37 ± 0%   -24.08% (p=0.002 n=6)
RandomPfx6Size/100_000    105.90 ± 0%    64.04 ± 0%   -39.53% (p=0.002 n=6)
RandomPfx6Size/200_000    101.00 ± 0%    64.02 ± 0%   -36.61% (p=0.002 n=6)
RandomPfxSize/1_000       105.20 ± 2%    68.16 ± 3%   -35.21% (p=0.002 n=6)
RandomPfxSize/10_000       76.12 ± 0%    64.37 ± 0%   -15.44% (p=0.002 n=6)
RandomPfxSize/100_000      82.42 ± 0%    64.04 ± 0%   -22.30% (p=0.002 n=6)
RandomPfxSize/200_000      78.45 ± 0%    64.02 ± 0%   -18.39% (p=0.002 n=6)
geomean                    77.60         65.20        -15.98%
```

## update, insert/delete

`bart` is by far the fastest algorithm for inserts and one of the fastest for delete.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │             art/update.bm              │
                         │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000        43.75n ± 1%   265.15n ±  3%   +506.06% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.16n ± 1%   532.75n ±  1%   +883.66% (p=0.002 n=6)
InsertRandomPfxs/100_000      125.5n ± 1%   1503.0n ± 52%  +1097.61% (p=0.002 n=6)
InsertRandomPfxs/200_000      220.9n ± 1%   1717.5n ± 15%   +677.50% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.84n ± 0%    19.21n ±  0%     +7.71% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.55n ± 0%    18.51n ±  0%     +5.44% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.76n ± 2%    19.02n ±  1%     +7.15% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.64n ± 1%    20.68n ±  8%    +10.92% (p=0.002 n=6)
geomean                       40.19n         122.6n         +205.02%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           netipds/update.bm            │
                         │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000        43.75n ± 1%   156.60n ±  0%   +257.94% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.16n ± 1%   240.45n ±  0%   +343.96% (p=0.002 n=6)
InsertRandomPfxs/100_000      125.5n ± 1%    447.3n ±  4%   +256.41% (p=0.002 n=6)
InsertRandomPfxs/200_000      220.9n ± 1%    627.4n ±  4%   +184.00% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.84n ± 0%   135.30n ±  3%   +658.62% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.55n ± 0%   201.65n ±  0%  +1048.68% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.76n ± 2%   418.10n ± 15%  +2254.83% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.64n ± 1%   597.85n ±  8%  +3107.35% (p=0.002 n=6)
geomean                       40.19n         303.5n         +655.26%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │         critbitgo/update.bm          │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000        43.75n ± 1%   169.80n ± 4%  +288.11% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.16n ± 1%   240.15n ± 1%  +343.41% (p=0.002 n=6)
InsertRandomPfxs/100_000      125.5n ± 1%    563.0n ± 2%  +348.57% (p=0.002 n=6)
InsertRandomPfxs/200_000      220.9n ± 1%    760.2n ± 3%  +244.16% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.84n ± 0%    67.86n ± 0%  +280.49% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.55n ± 0%    71.17n ± 2%  +305.41% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.76n ± 2%    76.44n ± 1%  +330.53% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.64n ± 1%    80.69n ± 2%  +332.89% (p=0.002 n=6)
geomean                       40.19n         163.9n       +307.73%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           lpmtrie/update.bm            │
                         │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000        43.75n ± 1%   363.55n ±  4%   +730.97% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.16n ± 1%   474.95n ±  1%   +776.94% (p=0.002 n=6)
InsertRandomPfxs/100_000      125.5n ± 1%   1154.5n ±  2%   +819.92% (p=0.002 n=6)
InsertRandomPfxs/200_000      220.9n ± 1%   1227.5n ± 17%   +455.68% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.84n ± 0%    77.98n ±  0%   +337.20% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.55n ± 0%   145.75n ±  0%   +730.25% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.76n ± 2%   317.45n ±  1%  +1687.95% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.64n ± 1%   461.80n ±  3%  +2377.47% (p=0.002 n=6)
geomean                       40.19n         377.0n         +837.94%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           cidranger/update.bm            │
                         │   sec/route    │   sec/route     vs base                  │
InsertRandomPfxs/1_000        43.75n ± 1%   4229.50n ±  2%   +9567.43% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.16n ± 1%   6225.50n ±  3%  +11394.65% (p=0.002 n=6)
InsertRandomPfxs/100_000      125.5n ± 1%    9719.5n ± 14%   +7644.62% (p=0.002 n=6)
InsertRandomPfxs/200_000      220.9n ± 1%   11964.5n ±  6%   +5316.25% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.84n ± 0%     86.74n ±  0%    +386.35% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.55n ± 0%     87.12n ±  8%    +396.27% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.76n ± 2%    135.85n ±  1%    +665.14% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.64n ± 1%    563.30n ± 12%   +2922.00% (p=0.002 n=6)
geomean                       40.19n          1.074µ         +2572.39%
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           cidrtree/update.bm            │
                         │   sec/route    │   sec/route     vs base                 │
InsertRandomPfxs/1_000        43.75n ± 1%   1272.50n ±  3%  +2808.57% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.16n ± 1%   1894.00n ±  4%  +3397.05% (p=0.002 n=6)
InsertRandomPfxs/100_000      125.5n ± 1%    3173.5n ±  2%  +2428.69% (p=0.002 n=6)
InsertRandomPfxs/200_000      220.9n ± 1%    3848.0n ±  2%  +1641.96% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.84n ± 0%     14.71n ±  1%    -17.55% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.55n ± 0%     15.03n ±  1%    -14.41% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.76n ± 2%     23.68n ±  3%    +33.37% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.64n ± 1%    391.15n ± 60%  +1998.44% (p=0.002 n=6)
geomean                       40.19n          296.8n         +638.52%
```
