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

## size of the routing tables


The memory consumption of the multibit trie `art` with **1_000_000** randomly distributed IPv6
prefixes brings the OOM killer in action.

`bart` with path compression has two orders of magnitude lower memory consumption compared to `art`
and is similar in low memory consumption to a binary search tree, like the `cidrtree` but
much faster in lookup times.


```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │              art/size.bm              │
                       │ bytes/route  │ bytes/route   vs base                 │
Tier1PfxSize/1_000         99.82 ± 2%   7590.00 ± 0%  +7503.69% (p=0.002 n=6)
Tier1PfxSize/10_000        73.46 ± 0%   4889.00 ± 0%  +6555.32% (p=0.002 n=6)
Tier1PfxSize/100_000       32.92 ± 0%   1669.00 ± 0%  +4969.87% (p=0.002 n=6)
Tier1PfxSize/200_000       22.15 ± 0%   1098.00 ± 0%  +4857.11% (p=0.002 n=6)
RandomPfx4Size/1_000       97.81 ± 2%   7478.00 ± 0%  +7545.44% (p=0.002 n=6)
RandomPfx4Size/10_000      68.30 ± 0%   6097.00 ± 0%  +8826.79% (p=0.002 n=6)
RandomPfx4Size/100_000     76.53 ± 0%   5471.00 ± 0%  +7048.83% (p=0.002 n=6)
RandomPfx4Size/200_000     70.84 ± 0%   4827.00 ± 0%  +6713.95% (p=0.002 n=6)
RandomPfx6Size/1_000       100.2 ± 2%    7877.0 ± 0%  +7761.28% (p=0.002 n=6)
RandomPfx6Size/10_000      75.62 ± 0%   6811.00 ± 0%  +8906.88% (p=0.002 n=6)
RandomPfx6Size/100_000     99.65 ± 0%   7829.00 ± 0%  +7756.50% (p=0.002 n=6)
RandomPfx6Size/200_000     96.21 ± 0%   7599.00 ± 0%  +7798.35% (p=0.002 n=6)
RandomPfxSize/1_000        103.0 ± 2%    7498.0 ± 0%  +7179.61% (p=0.002 n=6)
RandomPfxSize/10_000       70.52 ± 0%   6214.00 ± 0%  +8711.68% (p=0.002 n=6)
RandomPfxSize/100_000      76.37 ± 0%   5785.00 ± 0%  +7474.96% (p=0.002 n=6)
RandomPfxSize/200_000      76.06 ± 0%   5458.00 ± 0%  +7075.91% (p=0.002 n=6)
geomean                    72.58        5.173Ki       +7198.15%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidrtree/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         99.82 ± 2%    69.06 ± 3%   -30.82% (p=0.002 n=6)
Tier1PfxSize/10_000        73.46 ± 0%    64.35 ± 0%   -12.40% (p=0.002 n=6)
Tier1PfxSize/100_000       32.92 ± 0%    64.03 ± 0%   +94.50% (p=0.002 n=6)
Tier1PfxSize/200_000       22.15 ± 0%    64.02 ± 0%  +189.03% (p=0.002 n=6)
RandomPfx4Size/1_000       97.81 ± 2%    67.89 ± 3%   -30.59% (p=0.002 n=6)
RandomPfx4Size/10_000      68.30 ± 0%    64.35 ± 0%    -5.78% (p=0.002 n=6)
RandomPfx4Size/100_000     76.53 ± 0%    64.03 ± 0%   -16.33% (p=0.002 n=6)
RandomPfx4Size/200_000     70.84 ± 0%    64.02 ± 0%    -9.63% (p=0.002 n=6)
RandomPfx6Size/1_000      100.20 ± 2%    68.06 ± 2%   -32.08% (p=0.002 n=6)
RandomPfx6Size/10_000      75.62 ± 0%    64.35 ± 0%   -14.90% (p=0.002 n=6)
RandomPfx6Size/100_000     99.65 ± 0%    64.03 ± 0%   -35.75% (p=0.002 n=6)
RandomPfx6Size/200_000     96.21 ± 0%    64.02 ± 0%   -33.46% (p=0.002 n=6)
RandomPfxSize/1_000       103.00 ± 2%    67.89 ± 3%   -34.09% (p=0.002 n=6)
RandomPfxSize/10_000       70.52 ± 0%    64.39 ± 0%    -8.69% (p=0.002 n=6)
RandomPfxSize/100_000      76.37 ± 0%    64.03 ± 0%   -16.16% (p=0.002 n=6)
RandomPfxSize/200_000      76.06 ± 0%    64.02 ± 0%   -15.83% (p=0.002 n=6)
geomean                    72.58         65.13        -10.26%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          critbitgo/size.bm          │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         99.82 ± 2%   119.10 ± 2%   +19.31% (p=0.002 n=6)
Tier1PfxSize/10_000        73.46 ± 0%   114.70 ± 0%   +56.14% (p=0.002 n=6)
Tier1PfxSize/100_000       32.92 ± 0%   114.40 ± 0%  +247.51% (p=0.002 n=6)
Tier1PfxSize/200_000       22.15 ± 0%   114.40 ± 0%  +416.48% (p=0.002 n=6)
RandomPfx4Size/1_000       97.81 ± 2%   115.90 ± 2%   +18.50% (p=0.002 n=6)
RandomPfx4Size/10_000      68.30 ± 0%   112.30 ± 0%   +64.42% (p=0.002 n=6)
RandomPfx4Size/100_000     76.53 ± 0%   112.00 ± 0%   +46.35% (p=0.002 n=6)
RandomPfx4Size/200_000     70.84 ± 0%   112.00 ± 0%   +58.10% (p=0.002 n=6)
RandomPfx6Size/1_000       100.2 ± 2%    132.0 ± 1%   +31.74% (p=0.002 n=6)
RandomPfx6Size/10_000      75.62 ± 0%   128.30 ± 0%   +69.66% (p=0.002 n=6)
RandomPfx6Size/100_000     99.65 ± 0%   128.00 ± 0%   +28.45% (p=0.002 n=6)
RandomPfx6Size/200_000     96.21 ± 0%   128.00 ± 0%   +33.04% (p=0.002 n=6)
RandomPfxSize/1_000        103.0 ± 2%    118.8 ± 1%   +15.34% (p=0.002 n=6)
RandomPfxSize/10_000       70.52 ± 0%   115.50 ± 0%   +63.78% (p=0.002 n=6)
RandomPfxSize/100_000      76.37 ± 0%   115.20 ± 0%   +50.84% (p=0.002 n=6)
RandomPfxSize/200_000      76.06 ± 0%   115.20 ± 0%   +51.46% (p=0.002 n=6)
geomean                    72.58         118.3        +63.01%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │           lpmtrie/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         99.82 ± 2%   214.90 ± 5%  +115.29% (p=0.002 n=6)
Tier1PfxSize/10_000        73.46 ± 0%   210.50 ± 5%  +186.55% (p=0.002 n=6)
Tier1PfxSize/100_000       32.92 ± 0%   209.90 ± 5%  +537.61% (p=0.002 n=6)
Tier1PfxSize/200_000       22.15 ± 0%   209.20 ± 5%  +844.47% (p=0.002 n=6)
RandomPfx4Size/1_000       97.81 ± 2%   211.70 ± 5%  +116.44% (p=0.002 n=6)
RandomPfx4Size/10_000      68.30 ± 0%   206.80 ± 5%  +202.78% (p=0.002 n=6)
RandomPfx4Size/100_000     76.53 ± 0%   199.40 ± 6%  +160.55% (p=0.002 n=6)
RandomPfx4Size/200_000     70.84 ± 0%   194.60 ± 7%  +174.70% (p=0.002 n=6)
RandomPfx6Size/1_000       100.2 ± 2%    227.5 ± 8%  +127.05% (p=0.002 n=6)
RandomPfx6Size/10_000      75.62 ± 0%   223.40 ± 8%  +195.42% (p=0.002 n=6)
RandomPfx6Size/100_000     99.65 ± 0%   219.70 ± 8%  +120.47% (p=0.002 n=6)
RandomPfx6Size/200_000     96.21 ± 0%   218.00 ± 8%  +126.59% (p=0.002 n=6)
RandomPfxSize/1_000        103.0 ± 2%    214.9 ± 5%  +108.64% (p=0.002 n=6)
RandomPfxSize/10_000       70.52 ± 0%   209.60 ± 5%  +197.22% (p=0.002 n=6)
RandomPfxSize/100_000      76.37 ± 0%   203.30 ± 7%  +166.20% (p=0.002 n=6)
RandomPfxSize/200_000      76.06 ± 0%   199.30 ± 7%  +162.03% (p=0.002 n=6)
geomean                    72.58         210.6       +190.17%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidranger/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                 │
Tier1PfxSize/1_000         99.82 ± 2%   539.20 ± 3%   +440.17% (p=0.002 n=6)
Tier1PfxSize/10_000        73.46 ± 0%   533.80 ± 3%   +626.65% (p=0.002 n=6)
Tier1PfxSize/100_000       32.92 ± 0%   527.20 ± 2%  +1501.46% (p=0.002 n=6)
Tier1PfxSize/200_000       22.15 ± 0%   522.20 ± 2%  +2257.56% (p=0.002 n=6)
RandomPfx4Size/1_000       97.81 ± 2%   526.30 ± 3%   +438.08% (p=0.002 n=6)
RandomPfx4Size/10_000      68.30 ± 0%   514.30 ± 3%   +653.00% (p=0.002 n=6)
RandomPfx4Size/100_000     76.53 ± 0%   479.90 ± 3%   +527.07% (p=0.002 n=6)
RandomPfx4Size/200_000     70.84 ± 0%   463.20 ± 3%   +553.87% (p=0.002 n=6)
RandomPfx6Size/1_000       100.2 ± 2%    593.2 ± 0%   +492.02% (p=0.002 n=6)
RandomPfx6Size/10_000      75.62 ± 0%   585.90 ± 0%   +674.80% (p=0.002 n=6)
RandomPfx6Size/100_000     99.65 ± 0%   574.20 ± 0%   +476.22% (p=0.002 n=6)
RandomPfx6Size/200_000     96.21 ± 0%   570.00 ± 0%   +492.45% (p=0.002 n=6)
RandomPfxSize/1_000        103.0 ± 2%    537.7 ± 2%   +422.04% (p=0.002 n=6)
RandomPfxSize/10_000       70.52 ± 0%   526.40 ± 2%   +646.45% (p=0.002 n=6)
RandomPfxSize/100_000      76.37 ± 0%   498.50 ± 2%   +552.74% (p=0.002 n=6)
RandomPfxSize/200_000      76.06 ± 0%   484.50 ± 2%   +537.00% (p=0.002 n=6)
geomean                    72.58         528.5        +628.16%
```

## lpm (longest-prefix-match)

For longest-prefix-match, `bart` is the champion.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │               art/lpm.bm               │
                                     │    sec/op     │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            25.94n ±   9%    41.13n ± 19%   +58.58% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            30.23n ±  51%    43.12n ± 13%   +42.64% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             13.82n ± 195%    25.50n ±  2%         ~ (p=0.470 n=10)
LpmTier1Pfxs/RandomMissIP6             6.317n ±   0%   25.515n ±  0%  +303.94% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     16.27n ±  16%    40.64n ±  3%  +149.83% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     16.79n ±   2%    41.01n ±  0%  +144.22% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      14.97n ±  17%    25.46n ±  0%   +70.13% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      17.81n ±  19%    25.47n ± 62%   +43.01% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    9.404n ± 159%   40.000n ±  3%  +325.35% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    11.54n ± 126%    42.27n ±  3%  +266.18% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     24.96n ±  16%    25.51n ±  1%         ~ (p=0.470 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     22.16n ±  20%    25.61n ±  3%   +15.54% (p=0.021 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   11.83n ± 134%    42.57n ± 18%  +259.96% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   11.20n ± 137%    44.02n ± 30%  +293.17% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    39.28n ±  16%    25.84n ±  2%   -34.22% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    26.04n ±  46%    25.47n ±  1%         ~ (p=0.564 n=10)
geomean                                16.81n           32.69n         +94.49%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │               cidrtree/lpm.bm                │
                                     │    sec/op     │      sec/op        vs base                   │
LpmTier1Pfxs/RandomMatchIP4            25.94n ±   9%     826.40n ±   32%   +3085.81% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            30.23n ±  51%     991.45n ±   35%   +3179.69% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             13.82n ± 195%    1059.00n ±   97%   +7565.58% (p=0.001 n=10)
LpmTier1Pfxs/RandomMissIP6             6.317n ±   0%     39.065n ± 1641%    +518.46% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     16.27n ±  16%     331.55n ±   51%   +1938.43% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     16.79n ±   2%     456.55n ±   20%   +2619.18% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      14.97n ±  17%     459.55n ±   27%   +2970.83% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      17.81n ±  19%     574.55n ±   43%   +3126.00% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    9.404n ± 159%   1106.500n ±   19%  +11666.27% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    11.54n ± 126%     733.70n ±   30%   +6255.13% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     24.96n ±  16%     695.85n ±   20%   +2687.30% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     22.16n ±  20%     852.55n ±   11%   +3746.38% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   11.83n ± 134%    1003.50n ±   28%   +8386.26% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   11.20n ± 137%    1391.00n ±   77%  +12325.19% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    39.28n ±  16%    1039.00n ±   29%   +2545.11% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    26.04n ±  46%    1395.50n ±   29%   +5260.09% (p=0.000 n=10)
geomean                                16.81n             660.8n           +3831.65%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │             critbitgo/lpm.bm             │
                                     │    sec/op     │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            25.94n ±   9%    122.75n ± 28%   +373.21% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            30.23n ±  51%    191.25n ± 23%   +532.65% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             13.82n ± 195%    634.90n ± 32%  +4495.73% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             6.317n ±   0%   426.100n ± 53%  +6645.82% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     16.27n ±  16%    110.20n ± 13%   +577.53% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     16.79n ±   2%    119.45n ± 21%   +611.44% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      14.97n ±  17%    176.10n ± 15%  +1076.75% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      17.81n ±  19%    162.10n ± 14%   +810.16% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    9.404n ± 159%   121.600n ± 54%  +1193.07% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    11.54n ± 126%    152.25n ±  9%  +1218.75% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     24.96n ±  16%    256.55n ± 26%   +927.64% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     22.16n ±  20%    244.05n ± 18%  +1001.06% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   11.83n ± 134%    175.80n ± 41%  +1386.68% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   11.20n ± 137%    191.25n ± 32%  +1608.35% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    39.28n ±  16%    548.30n ± 34%  +1295.88% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    26.04n ±  46%    492.55n ± 17%  +1791.88% (p=0.000 n=10)
geomean                                16.81n            217.5n        +1194.28%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │              lpmtrie/lpm.bm               │
                                     │    sec/op     │     sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4            25.94n ±   9%    210.85n ±  11%   +712.84% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            30.23n ±  51%    214.05n ±  67%   +608.07% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             13.82n ± 195%    108.97n ± 130%   +688.82% (p=0.007 n=10)
LpmTier1Pfxs/RandomMissIP6             6.317n ±   0%    12.935n ±   4%   +104.78% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     16.27n ±  16%    134.65n ±   6%   +727.85% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     16.79n ±   2%    100.55n ±   8%   +498.87% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      14.97n ±  17%    128.75n ±  28%   +760.34% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      17.81n ±  19%     83.77n ±  13%   +370.33% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    9.404n ± 159%   149.150n ±   6%  +1486.03% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    11.54n ± 126%    132.10n ±   6%  +1044.22% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     24.96n ±  16%    141.55n ±   7%   +466.99% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     22.16n ±  20%    127.35n ±  11%   +474.55% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   11.83n ± 134%    175.30n ±   8%  +1382.45% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   11.20n ± 137%    159.65n ±   7%  +1326.08% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    39.28n ±  16%    190.40n ±   5%   +384.73% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    26.04n ±  46%    179.20n ±  15%   +588.30% (p=0.000 n=10)
geomean                                16.81n            124.1n          +638.20%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │             cidranger/lpm.bm             │
                                     │    sec/op     │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            25.94n ±   9%    168.95n ± 34%   +551.31% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            30.23n ±  51%    206.35n ± 38%   +582.60% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             13.82n ± 195%    195.75n ± 50%  +1316.94% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             6.317n ±   0%    59.765n ± 14%   +846.17% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     16.27n ±  16%     81.05n ± 50%   +398.31% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     16.79n ±   2%    130.90n ±  7%   +679.63% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      14.97n ±  17%    112.75n ±  9%   +653.42% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      17.81n ±  19%    137.70n ± 31%   +673.16% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    9.404n ± 159%   146.750n ± 33%  +1460.51% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    11.54n ± 126%    167.10n ± 26%  +1347.38% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     24.96n ±  16%    170.25n ± 19%   +581.95% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     22.16n ±  20%    168.20n ±  9%   +658.85% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   11.83n ± 134%    145.60n ± 40%  +1131.29% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   11.20n ± 137%    166.15n ± 31%  +1384.14% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    39.28n ±  16%    194.15n ± 11%   +394.27% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    26.04n ±  46%    246.75n ± 19%   +847.76% (p=0.000 n=10)
geomean                                16.81n            148.3n         +782.35%
```

## update, insert/delete

`bart` is by far the fastest algorithm for inserts and one of the fastest for delete.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │              art/update.bm              │
                         │   sec/route    │   sec/route     vs base                 │
InsertRandomPfxs/1_000        43.38n ± 1%   280.35n ±   3%   +546.19% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.94n ± 0%   577.05n ±   9%   +950.42% (p=0.002 n=6)
InsertRandomPfxs/100_000      144.1n ± 2%   1723.0n ±  24%  +1095.70% (p=0.002 n=6)
InsertRandomPfxs/200_000      293.4n ± 2%   1817.0n ± 198%   +519.19% (p=0.002 n=6)
DeleteRandomPfxs/1_000        19.07n ± 0%    17.62n ±   1%     -7.63% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.34n ± 1%    16.90n ±   0%     -7.82% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.29n ± 2%    17.52n ±   3%     -4.21% (p=0.002 n=6)
DeleteRandomPfxs/200_000      20.16n ± 1%    19.46n ±   4%     -3.47% (p=0.015 n=6)
geomean                       43.58n         122.7n          +181.62%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │            cidrtree/update.bm            │
                         │   sec/route    │   sec/route     vs base                  │
InsertRandomPfxs/1_000        43.38n ± 1%   1320.50n ± 44%   +2943.68% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.94n ± 0%   1971.50n ±  4%   +3488.79% (p=0.002 n=6)
InsertRandomPfxs/100_000      144.1n ± 2%    3406.5n ±  4%   +2263.98% (p=0.002 n=6)
InsertRandomPfxs/200_000      293.4n ± 2%    4218.5n ±  3%   +1337.55% (p=0.002 n=6)
DeleteRandomPfxs/1_000        19.07n ± 0%     14.76n ±  1%     -22.62% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.34n ± 1%     15.15n ±  0%     -17.39% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.29n ± 2%     22.90n ±  4%     +25.24% (p=0.002 n=6)
DeleteRandomPfxs/200_000      20.16n ± 1%   3381.50n ± 63%  +16669.15% (p=0.002 n=6)
geomean                       43.58n          399.4n          +816.62%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │         critbitgo/update.bm          │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000        43.38n ± 1%   176.25n ± 1%  +306.25% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.94n ± 0%   245.95n ± 1%  +347.71% (p=0.002 n=6)
InsertRandomPfxs/100_000      144.1n ± 2%    614.1n ± 1%  +326.13% (p=0.002 n=6)
InsertRandomPfxs/200_000      293.4n ± 2%    797.0n ± 2%  +171.58% (p=0.002 n=6)
DeleteRandomPfxs/1_000        19.07n ± 0%    71.63n ± 4%  +275.52% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.34n ± 1%    74.20n ± 2%  +304.58% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.29n ± 2%    81.35n ± 1%  +344.90% (p=0.002 n=6)
DeleteRandomPfxs/200_000      20.16n ± 1%    86.52n ± 0%  +329.06% (p=0.002 n=6)
geomean                       43.58n         172.8n       +296.47%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           lpmtrie/update.bm           │
                         │   sec/route    │  sec/route    vs base                 │
InsertRandomPfxs/1_000        43.38n ± 1%   400.10n ± 3%   +822.21% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.94n ± 0%   492.45n ± 6%   +796.42% (p=0.002 n=6)
InsertRandomPfxs/100_000      144.1n ± 2%   1255.0n ± 2%   +770.92% (p=0.002 n=6)
InsertRandomPfxs/200_000      293.4n ± 2%   1317.5n ± 3%   +348.97% (p=0.002 n=6)
DeleteRandomPfxs/1_000        19.07n ± 0%    75.38n ± 2%   +295.18% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.34n ± 1%   149.15n ± 3%   +713.25% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.29n ± 2%   324.65n ± 2%  +1675.50% (p=0.002 n=6)
DeleteRandomPfxs/200_000      20.16n ± 1%   479.85n ± 8%  +2279.62% (p=0.002 n=6)
geomean                       43.58n         393.1n        +802.19%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           cidranger/update.bm            │
                         │   sec/route    │   sec/route     vs base                  │
InsertRandomPfxs/1_000        43.38n ± 1%   4596.00n ±  4%  +10493.52% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.94n ± 0%   6859.00n ±  3%  +12385.66% (p=0.002 n=6)
InsertRandomPfxs/100_000      144.1n ± 2%   11253.0n ±  4%   +7709.16% (p=0.002 n=6)
InsertRandomPfxs/200_000      293.4n ± 2%   13186.0n ±  4%   +4393.44% (p=0.002 n=6)
DeleteRandomPfxs/1_000        19.07n ± 0%     90.89n ±  3%    +376.49% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.34n ± 1%     92.27n ±  1%    +403.14% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.29n ± 2%    145.05n ±  2%    +693.27% (p=0.002 n=6)
DeleteRandomPfxs/200_000      20.16n ± 1%   1280.50n ± 32%   +6250.11% (p=0.002 n=6)
geomean                       43.58n          1.282µ         +2841.52%
```
