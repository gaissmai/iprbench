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

In comparison, the prefix lengths for the random test sets are equally distributed between /2-32 for IPv4
and /2-128 bits for IPv6, the randomly generated _default-routes_ with prefix length 0 have been sorted out,
they distorts the lookup times and there is no lookup miss at all.

The **RandomPrefixes** without IP versions labeling are composed of a distribution of 4 parts IPv4
to 1 part IPv6 prefixes, which is approximately the current ratio in the Internet backbone routers.

## make your own benchmarks

```
  $ make dep  
  $ make -B all
```

## lpm (longest-prefix-match)

`bart` is the fastest software algorithm for IP address lookup in routing tables.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │               art/lpm.bm               │
                                     │    sec/op     │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            25.07n ±  17%    40.76n ± 21%   +62.58% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            29.80n ±  51%    43.19n ± 14%   +44.95% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             20.54n ±  91%    25.44n ±  4%         ~ (p=0.898 n=10)
LpmTier1Pfxs/RandomMissIP6             6.298n ±   0%   25.505n ±  0%  +304.94% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     21.36n ±  16%    39.83n ±  2%   +86.43% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     15.58n ±  10%    40.96n ±  0%  +162.78% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      22.83n ±  19%    25.45n ±  1%         ~ (p=0.137 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      16.96n ±  63%    25.46n ±  0%   +50.12% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    25.34n ±  52%    39.53n ±  3%   +55.95% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    10.25n ± 125%    41.16n ±  3%  +301.56% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     26.84n ±  38%    25.45n ±  0%    -5.18% (p=0.022 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     21.19n ±   9%    25.47n ±  0%   +20.20% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   17.36n ±  53%    42.85n ± 16%  +146.83% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   15.85n ±  67%    42.35n ± 18%  +167.19% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    38.39n ±   9%    26.56n ±  9%   -30.82% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    27.27n ±  45%    25.58n ± 25%         ~ (p=0.517 n=10)
geomean                                19.74n           32.53n         +64.80%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │              cidrtree/lpm.bm               │
                                     │    sec/op     │      sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4            25.07n ±  17%    826.40n ±   32%  +3196.37% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            29.80n ±  51%    991.45n ±   35%  +3227.01% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             20.54n ±  91%   1059.00n ±   97%  +5055.79% (p=0.002 n=10)
LpmTier1Pfxs/RandomMissIP6             6.298n ±   0%    39.065n ± 1641%   +520.23% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     21.36n ±  16%    331.55n ±   51%  +1451.84% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     15.58n ±  10%    456.55n ±   20%  +2829.42% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      22.83n ±  19%    459.55n ±   27%  +1912.92% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      16.96n ±  63%    574.55n ±   43%  +3287.68% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    25.34n ±  52%   1106.50n ±   19%  +4265.75% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    10.25n ± 125%    733.70n ±   30%  +7058.05% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     26.84n ±  38%    695.85n ±   20%  +2492.59% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     21.19n ±   9%    852.55n ±   11%  +3923.36% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   17.36n ±  53%   1003.50n ±   28%  +5680.53% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   15.85n ±  67%   1391.00n ±   77%  +8676.03% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    38.39n ±   9%   1039.00n ±   29%  +2606.08% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    27.27n ±  45%   1395.50n ±   29%  +5017.35% (p=0.000 n=10)
geomean                                19.74n            660.8n          +3247.78%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │             critbitgo/lpm.bm             │
                                     │    sec/op     │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            25.07n ±  17%    122.75n ± 28%   +389.63% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            29.80n ±  51%    191.25n ± 23%   +541.78% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             20.54n ±  91%    634.90n ± 32%  +2991.04% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             6.298n ±   0%   426.100n ± 53%  +6665.10% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     21.36n ±  16%    110.20n ± 13%   +415.80% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     15.58n ±  10%    119.45n ± 21%   +666.44% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      22.83n ±  19%    176.10n ± 15%   +671.35% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      16.96n ±  63%    162.10n ± 14%   +855.78% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    25.34n ±  52%    121.60n ± 54%   +379.78% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    10.25n ± 125%    152.25n ±  9%  +1385.37% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     26.84n ±  38%    256.55n ± 26%   +855.85% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     21.19n ±   9%    244.05n ± 18%  +1051.72% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   17.36n ±  53%    175.80n ± 41%   +912.67% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   15.85n ±  67%    191.25n ± 32%  +1106.62% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    38.39n ±   9%    548.30n ± 34%  +1328.05% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    27.27n ±  45%    492.55n ± 17%  +1706.20% (p=0.000 n=10)
geomean                                19.74n            217.5n        +1002.07%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │              lpmtrie/lpm.bm              │
                                     │    sec/op     │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            25.07n ±  17%   210.85n ±  11%   +741.05% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            29.80n ±  51%   214.05n ±  67%   +618.29% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             20.54n ±  91%   108.97n ± 130%   +430.55% (p=0.007 n=10)
LpmTier1Pfxs/RandomMissIP6             6.298n ±   0%   12.935n ±   4%   +105.37% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     21.36n ±  16%   134.65n ±   6%   +530.24% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     15.58n ±  10%   100.55n ±   8%   +545.17% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      22.83n ±  19%   128.75n ±  28%   +463.95% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      16.96n ±  63%    83.77n ±  13%   +393.90% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    25.34n ±  52%   149.15n ±   6%   +488.48% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    10.25n ± 125%   132.10n ±   6%  +1188.78% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     26.84n ±  38%   141.55n ±   7%   +427.38% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     21.19n ±   9%   127.35n ±  11%   +500.99% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   17.36n ±  53%   175.30n ±   8%   +909.79% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   15.85n ±  67%   159.65n ±   7%   +907.26% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    38.39n ±   9%   190.40n ±   5%   +395.90% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    27.27n ±  45%   179.20n ±  15%   +557.13% (p=0.000 n=10)
geomean                                19.74n           124.1n          +528.58%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │            cidranger/lpm.bm             │
                                     │    sec/op     │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            25.07n ±  17%   168.95n ± 34%   +573.91% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            29.80n ±  51%   206.35n ± 38%   +592.45% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             20.54n ±  91%   195.75n ± 50%   +853.02% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             6.298n ±   0%   59.765n ± 14%   +848.88% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     21.36n ±  16%    81.05n ± 50%   +279.36% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     15.58n ±  10%   130.90n ±  7%   +739.91% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      22.83n ±  19%   112.75n ±  9%   +393.87% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      16.96n ±  63%   137.70n ± 31%   +711.91% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    25.34n ±  52%   146.75n ± 33%   +479.01% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    10.25n ± 125%   167.10n ± 26%  +1530.24% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     26.84n ±  38%   170.25n ± 19%   +534.31% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     21.19n ±   9%   168.20n ±  9%   +693.77% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   17.36n ±  53%   145.60n ± 40%   +738.71% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   15.85n ±  67%   166.15n ± 31%   +948.26% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    38.39n ±   9%   194.15n ± 11%   +405.66% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    27.27n ±  45%   246.75n ± 19%   +804.84% (p=0.000 n=10)
geomean                                19.74n           148.3n         +651.32%
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
Tier1PfxSize/1_000         106.9 ± 2%    7590.0 ± 0%  +7000.09% (p=0.002 n=6)
Tier1PfxSize/10_000        79.54 ± 0%   4889.00 ± 0%  +6046.59% (p=0.002 n=6)
Tier1PfxSize/100_000       35.73 ± 0%   1669.00 ± 0%  +4571.14% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%   1098.00 ± 0%  +4448.47% (p=0.002 n=6)
RandomPfx4Size/1_000       102.5 ± 2%    7381.0 ± 0%  +7100.98% (p=0.002 n=6)
RandomPfx4Size/10_000      77.37 ± 0%   6050.00 ± 0%  +7719.57% (p=0.002 n=6)
RandomPfx4Size/100_000     81.87 ± 0%   5488.00 ± 0%  +6603.31% (p=0.002 n=6)
RandomPfx4Size/200_000     72.96 ± 0%   4833.00 ± 0%  +6524.18% (p=0.002 n=6)
RandomPfx6Size/1_000       105.4 ± 2%    7864.0 ± 0%  +7361.10% (p=0.002 n=6)
RandomPfx6Size/10_000      85.15 ± 0%   6827.00 ± 0%  +7917.62% (p=0.002 n=6)
RandomPfx6Size/100_000     105.6 ± 0%    7832.0 ± 0%  +7316.67% (p=0.002 n=6)
RandomPfx6Size/200_000     100.7 ± 0%    7609.0 ± 0%  +7456.11% (p=0.002 n=6)
RandomPfxSize/1_000        105.5 ± 2%    7505.0 ± 0%  +7013.74% (p=0.002 n=6)
RandomPfxSize/10_000       75.46 ± 0%   6194.00 ± 0%  +8108.32% (p=0.002 n=6)
RandomPfxSize/100_000      82.40 ± 0%   5797.00 ± 0%  +6935.19% (p=0.002 n=6)
RandomPfxSize/200_000      78.43 ± 0%   5446.00 ± 0%  +6843.77% (p=0.002 n=6)
geomean                    77.53        5.168Ki       +6725.12%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidrtree/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000        106.90 ± 2%    69.06 ± 3%   -35.40% (p=0.002 n=6)
Tier1PfxSize/10_000        79.54 ± 0%    64.35 ± 0%   -19.10% (p=0.002 n=6)
Tier1PfxSize/100_000       35.73 ± 0%    64.03 ± 0%   +79.21% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%    64.02 ± 0%  +165.20% (p=0.002 n=6)
RandomPfx4Size/1_000      102.50 ± 2%    67.89 ± 3%   -33.77% (p=0.002 n=6)
RandomPfx4Size/10_000      77.37 ± 0%    64.35 ± 0%   -16.83% (p=0.002 n=6)
RandomPfx4Size/100_000     81.87 ± 0%    64.03 ± 0%   -21.79% (p=0.002 n=6)
RandomPfx4Size/200_000     72.96 ± 0%    64.02 ± 0%   -12.25% (p=0.002 n=6)
RandomPfx6Size/1_000      105.40 ± 2%    68.06 ± 2%   -35.43% (p=0.002 n=6)
RandomPfx6Size/10_000      85.15 ± 0%    64.35 ± 0%   -24.43% (p=0.002 n=6)
RandomPfx6Size/100_000    105.60 ± 0%    64.03 ± 0%   -39.37% (p=0.002 n=6)
RandomPfx6Size/200_000    100.70 ± 0%    64.02 ± 0%   -36.43% (p=0.002 n=6)
RandomPfxSize/1_000       105.50 ± 2%    67.89 ± 3%   -35.65% (p=0.002 n=6)
RandomPfxSize/10_000       75.46 ± 0%    64.39 ± 0%   -14.67% (p=0.002 n=6)
RandomPfxSize/100_000      82.40 ± 0%    64.03 ± 0%   -22.29% (p=0.002 n=6)
RandomPfxSize/200_000      78.43 ± 0%    64.02 ± 0%   -18.37% (p=0.002 n=6)
geomean                    77.53         65.13        -15.99%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          critbitgo/size.bm          │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         106.9 ± 2%    119.1 ± 2%   +11.41% (p=0.002 n=6)
Tier1PfxSize/10_000        79.54 ± 0%   114.70 ± 0%   +44.20% (p=0.002 n=6)
Tier1PfxSize/100_000       35.73 ± 0%   114.40 ± 0%  +220.18% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%   114.40 ± 0%  +373.90% (p=0.002 n=6)
RandomPfx4Size/1_000       102.5 ± 2%    115.9 ± 2%   +13.07% (p=0.002 n=6)
RandomPfx4Size/10_000      77.37 ± 0%   112.30 ± 0%   +45.15% (p=0.002 n=6)
RandomPfx4Size/100_000     81.87 ± 0%   112.00 ± 0%   +36.80% (p=0.002 n=6)
RandomPfx4Size/200_000     72.96 ± 0%   112.00 ± 0%   +53.51% (p=0.002 n=6)
RandomPfx6Size/1_000       105.4 ± 2%    132.0 ± 1%   +25.24% (p=0.002 n=6)
RandomPfx6Size/10_000      85.15 ± 0%   128.30 ± 0%   +50.68% (p=0.002 n=6)
RandomPfx6Size/100_000     105.6 ± 0%    128.0 ± 0%   +21.21% (p=0.002 n=6)
RandomPfx6Size/200_000     100.7 ± 0%    128.0 ± 0%   +27.11% (p=0.002 n=6)
RandomPfxSize/1_000        105.5 ± 2%    118.8 ± 1%   +12.61% (p=0.002 n=6)
RandomPfxSize/10_000       75.46 ± 0%   115.50 ± 0%   +53.06% (p=0.002 n=6)
RandomPfxSize/100_000      82.40 ± 0%   115.20 ± 0%   +39.81% (p=0.002 n=6)
RandomPfxSize/200_000      78.43 ± 0%   115.20 ± 0%   +46.88% (p=0.002 n=6)
geomean                    77.53         118.3        +52.60%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │           lpmtrie/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         106.9 ± 2%    214.9 ± 5%  +101.03% (p=0.002 n=6)
Tier1PfxSize/10_000        79.54 ± 0%   210.50 ± 5%  +164.65% (p=0.002 n=6)
Tier1PfxSize/100_000       35.73 ± 0%   209.90 ± 5%  +487.46% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%   209.20 ± 5%  +766.61% (p=0.002 n=6)
RandomPfx4Size/1_000       102.5 ± 2%    211.7 ± 5%  +106.54% (p=0.002 n=6)
RandomPfx4Size/10_000      77.37 ± 0%   206.80 ± 5%  +167.29% (p=0.002 n=6)
RandomPfx4Size/100_000     81.87 ± 0%   199.40 ± 6%  +143.56% (p=0.002 n=6)
RandomPfx4Size/200_000     72.96 ± 0%   194.60 ± 7%  +166.72% (p=0.002 n=6)
RandomPfx6Size/1_000       105.4 ± 2%    227.5 ± 8%  +115.84% (p=0.002 n=6)
RandomPfx6Size/10_000      85.15 ± 0%   223.40 ± 8%  +162.36% (p=0.002 n=6)
RandomPfx6Size/100_000     105.6 ± 0%    219.7 ± 8%  +108.05% (p=0.002 n=6)
RandomPfx6Size/200_000     100.7 ± 0%    218.0 ± 8%  +116.48% (p=0.002 n=6)
RandomPfxSize/1_000        105.5 ± 2%    214.9 ± 5%  +103.70% (p=0.002 n=6)
RandomPfxSize/10_000       75.46 ± 0%   209.60 ± 5%  +177.76% (p=0.002 n=6)
RandomPfxSize/100_000      82.40 ± 0%   203.30 ± 7%  +146.72% (p=0.002 n=6)
RandomPfxSize/200_000      78.43 ± 0%   199.30 ± 7%  +154.11% (p=0.002 n=6)
geomean                    77.53         210.6       +171.64%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidranger/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                 │
Tier1PfxSize/1_000         106.9 ± 2%    539.2 ± 3%   +404.40% (p=0.002 n=6)
Tier1PfxSize/10_000        79.54 ± 0%   533.80 ± 3%   +571.11% (p=0.002 n=6)
Tier1PfxSize/100_000       35.73 ± 0%   527.20 ± 2%  +1375.51% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%   522.20 ± 2%  +2063.21% (p=0.002 n=6)
RandomPfx4Size/1_000       102.5 ± 2%    526.3 ± 3%   +413.46% (p=0.002 n=6)
RandomPfx4Size/10_000      77.37 ± 0%   514.30 ± 3%   +564.73% (p=0.002 n=6)
RandomPfx4Size/100_000     81.87 ± 0%   479.90 ± 3%   +486.17% (p=0.002 n=6)
RandomPfx4Size/200_000     72.96 ± 0%   463.20 ± 3%   +534.87% (p=0.002 n=6)
RandomPfx6Size/1_000       105.4 ± 2%    593.2 ± 0%   +462.81% (p=0.002 n=6)
RandomPfx6Size/10_000      85.15 ± 0%   585.90 ± 0%   +588.08% (p=0.002 n=6)
RandomPfx6Size/100_000     105.6 ± 0%    574.2 ± 0%   +443.75% (p=0.002 n=6)
RandomPfx6Size/200_000     100.7 ± 0%    570.0 ± 0%   +466.04% (p=0.002 n=6)
RandomPfxSize/1_000        105.5 ± 2%    537.7 ± 2%   +409.67% (p=0.002 n=6)
RandomPfxSize/10_000       75.46 ± 0%   526.40 ± 2%   +597.59% (p=0.002 n=6)
RandomPfxSize/100_000      82.40 ± 0%   498.50 ± 2%   +504.98% (p=0.002 n=6)
RandomPfxSize/200_000      78.43 ± 0%   484.50 ± 2%   +517.75% (p=0.002 n=6)
geomean                    77.53         528.5        +581.67%
```

## update, insert/delete

`bart` is by far the fastest algorithm for inserts and one of the fastest for delete.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │              art/update.bm              │
                         │   sec/route    │   sec/route     vs base                 │
InsertRandomPfxs/1_000        42.44n ± 4%   281.00n ±   2%   +562.19% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.45n ± 1%   580.85n ±   6%   +966.86% (p=0.002 n=6)
InsertRandomPfxs/100_000      133.5n ± 9%   1743.5n ±  49%  +1206.48% (p=0.002 n=6)
InsertRandomPfxs/200_000      234.2n ± 1%   1846.0n ± 190%   +688.22% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.76n ± 1%    17.61n ±   0%     -0.87% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.51n ± 1%    16.94n ±   0%     -3.26% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.74n ± 1%    17.62n ±   0%     -0.65% (p=0.017 n=6)
DeleteRandomPfxs/200_000      19.08n ± 1%    19.52n ±   2%     +2.28% (p=0.002 n=6)
geomean                       40.75n         123.4n          +202.92%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │            cidrtree/update.bm            │
                         │   sec/route    │   sec/route     vs base                  │
InsertRandomPfxs/1_000        42.44n ± 4%   1320.50n ± 44%   +3011.82% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.45n ± 1%   1971.50n ±  4%   +3521.09% (p=0.002 n=6)
InsertRandomPfxs/100_000      133.5n ± 9%    3406.5n ±  4%   +2452.64% (p=0.002 n=6)
InsertRandomPfxs/200_000      234.2n ± 1%    4218.5n ±  3%   +1701.24% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.76n ± 1%     14.76n ±  1%     -16.89% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.51n ± 1%     15.15n ±  0%     -13.48% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.74n ± 1%     22.90n ±  4%     +29.09% (p=0.002 n=6)
DeleteRandomPfxs/200_000      19.08n ± 1%   3381.50n ± 63%  +17622.75% (p=0.002 n=6)
geomean                       40.75n          399.4n          +880.22%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │         critbitgo/update.bm          │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000        42.44n ± 4%   176.25n ± 1%  +315.34% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.45n ± 1%   245.95n ± 1%  +351.74% (p=0.002 n=6)
InsertRandomPfxs/100_000      133.5n ± 9%    614.1n ± 1%  +360.13% (p=0.002 n=6)
InsertRandomPfxs/200_000      234.2n ± 1%    797.0n ± 2%  +240.29% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.76n ± 1%    71.63n ± 4%  +303.32% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.51n ± 1%    74.20n ± 2%  +323.76% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.74n ± 1%    81.35n ± 1%  +358.57% (p=0.002 n=6)
DeleteRandomPfxs/200_000      19.08n ± 1%    86.52n ± 0%  +353.46% (p=0.002 n=6)
geomean                       40.75n         172.8n       +323.98%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           lpmtrie/update.bm           │
                         │   sec/route    │  sec/route    vs base                 │
InsertRandomPfxs/1_000        42.44n ± 4%   400.10n ± 3%   +842.85% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.45n ± 1%   492.45n ± 6%   +804.49% (p=0.002 n=6)
InsertRandomPfxs/100_000      133.5n ± 9%   1255.0n ± 2%   +840.43% (p=0.002 n=6)
InsertRandomPfxs/200_000      234.2n ± 1%   1317.5n ± 3%   +462.55% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.76n ± 1%    75.38n ± 2%   +324.44% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.51n ± 1%   149.15n ± 3%   +751.80% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.74n ± 1%   324.65n ± 2%  +1730.05% (p=0.002 n=6)
DeleteRandomPfxs/200_000      19.08n ± 1%   479.85n ± 8%  +2414.94% (p=0.002 n=6)
geomean                       40.75n         393.1n        +864.79%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           cidranger/update.bm            │
                         │   sec/route    │   sec/route     vs base                  │
InsertRandomPfxs/1_000        42.44n ± 4%   4596.00n ±  4%  +10730.68% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.45n ± 1%   6859.00n ±  3%  +12498.03% (p=0.002 n=6)
InsertRandomPfxs/100_000      133.5n ± 9%   11253.0n ±  4%   +8332.37% (p=0.002 n=6)
InsertRandomPfxs/200_000      234.2n ± 1%   13186.0n ±  4%   +5530.23% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.76n ± 1%     90.89n ±  3%    +411.77% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.51n ± 1%     92.27n ±  1%    +426.98% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.74n ± 1%    145.05n ±  2%    +717.64% (p=0.002 n=6)
DeleteRandomPfxs/200_000      19.08n ± 1%   1280.50n ± 32%   +6611.22% (p=0.002 n=6)
geomean                       40.75n          1.282µ         +3045.64%
```
