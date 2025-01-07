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
RandomPfx4Size/1_000       97.45 ± 2%   7478.00 ± 0%  +7573.68% (p=0.002 n=6)
RandomPfx4Size/10_000      68.19 ± 0%   6097.00 ± 0%  +8841.19% (p=0.002 n=6)
RandomPfx4Size/100_000     76.62 ± 0%   5471.00 ± 0%  +7040.43% (p=0.002 n=6)
RandomPfx4Size/200_000     70.78 ± 0%   4827.00 ± 0%  +6719.72% (p=0.002 n=6)
RandomPfx6Size/1_000       99.34 ± 2%   7877.00 ± 0%  +7829.33% (p=0.002 n=6)
RandomPfx6Size/10_000      75.99 ± 0%   6811.00 ± 0%  +8863.02% (p=0.002 n=6)
RandomPfx6Size/100_000     99.76 ± 0%   7829.00 ± 0%  +7747.83% (p=0.002 n=6)
RandomPfx6Size/200_000     96.09 ± 0%   7599.00 ± 0%  +7808.21% (p=0.002 n=6)
RandomPfxSize/1_000        100.4 ± 2%    7498.0 ± 0%  +7368.13% (p=0.002 n=6)
RandomPfxSize/10_000       70.53 ± 0%   6214.00 ± 0%  +8710.44% (p=0.002 n=6)
RandomPfxSize/100_000      76.30 ± 0%   5785.00 ± 0%  +7481.91% (p=0.002 n=6)
RandomPfxSize/200_000      75.96 ± 0%   5458.00 ± 0%  +7085.36% (p=0.002 n=6)
geomean                    72.42        5.173Ki       +7214.83%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidrtree/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         99.82 ± 2%    69.06 ± 3%   -30.82% (p=0.002 n=6)
Tier1PfxSize/10_000        73.46 ± 0%    64.35 ± 0%   -12.40% (p=0.002 n=6)
Tier1PfxSize/100_000       32.92 ± 0%    64.03 ± 0%   +94.50% (p=0.002 n=6)
Tier1PfxSize/200_000       22.15 ± 0%    64.02 ± 0%  +189.03% (p=0.002 n=6)
RandomPfx4Size/1_000       97.45 ± 2%    67.89 ± 3%   -30.33% (p=0.002 n=6)
RandomPfx4Size/10_000      68.19 ± 0%    64.35 ± 0%    -5.63% (p=0.002 n=6)
RandomPfx4Size/100_000     76.62 ± 0%    64.03 ± 0%   -16.43% (p=0.002 n=6)
RandomPfx4Size/200_000     70.78 ± 0%    64.02 ± 0%    -9.55% (p=0.002 n=6)
RandomPfx6Size/1_000       99.34 ± 2%    68.06 ± 2%   -31.49% (p=0.002 n=6)
RandomPfx6Size/10_000      75.99 ± 0%    64.35 ± 0%   -15.32% (p=0.002 n=6)
RandomPfx6Size/100_000     99.76 ± 0%    64.03 ± 0%   -35.82% (p=0.002 n=6)
RandomPfx6Size/200_000     96.09 ± 0%    64.02 ± 0%   -33.37% (p=0.002 n=6)
RandomPfxSize/1_000       100.40 ± 2%    67.89 ± 3%   -32.38% (p=0.002 n=6)
RandomPfxSize/10_000       70.53 ± 0%    64.39 ± 0%    -8.71% (p=0.002 n=6)
RandomPfxSize/100_000      76.30 ± 0%    64.03 ± 0%   -16.08% (p=0.002 n=6)
RandomPfxSize/200_000      75.96 ± 0%    64.02 ± 0%   -15.72% (p=0.002 n=6)
geomean                    72.42         65.13        -10.06%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          critbitgo/size.bm          │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         99.82 ± 2%   119.10 ± 2%   +19.31% (p=0.002 n=6)
Tier1PfxSize/10_000        73.46 ± 0%   114.70 ± 0%   +56.14% (p=0.002 n=6)
Tier1PfxSize/100_000       32.92 ± 0%   114.40 ± 0%  +247.51% (p=0.002 n=6)
Tier1PfxSize/200_000       22.15 ± 0%   114.40 ± 0%  +416.48% (p=0.002 n=6)
RandomPfx4Size/1_000       97.45 ± 2%   115.90 ± 2%   +18.93% (p=0.002 n=6)
RandomPfx4Size/10_000      68.19 ± 0%   112.30 ± 0%   +64.69% (p=0.002 n=6)
RandomPfx4Size/100_000     76.62 ± 0%   112.00 ± 0%   +46.18% (p=0.002 n=6)
RandomPfx4Size/200_000     70.78 ± 0%   112.00 ± 0%   +58.24% (p=0.002 n=6)
RandomPfx6Size/1_000       99.34 ± 2%   132.00 ± 1%   +32.88% (p=0.002 n=6)
RandomPfx6Size/10_000      75.99 ± 0%   128.30 ± 0%   +68.84% (p=0.002 n=6)
RandomPfx6Size/100_000     99.76 ± 0%   128.00 ± 0%   +28.31% (p=0.002 n=6)
RandomPfx6Size/200_000     96.09 ± 0%   128.00 ± 0%   +33.21% (p=0.002 n=6)
RandomPfxSize/1_000        100.4 ± 2%    118.8 ± 1%   +18.33% (p=0.002 n=6)
RandomPfxSize/10_000       70.53 ± 0%   115.50 ± 0%   +63.76% (p=0.002 n=6)
RandomPfxSize/100_000      76.30 ± 0%   115.20 ± 0%   +50.98% (p=0.002 n=6)
RandomPfxSize/200_000      75.96 ± 0%   115.20 ± 0%   +51.66% (p=0.002 n=6)
geomean                    72.42         118.3        +63.38%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │           lpmtrie/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         99.82 ± 2%   214.90 ± 5%  +115.29% (p=0.002 n=6)
Tier1PfxSize/10_000        73.46 ± 0%   210.50 ± 5%  +186.55% (p=0.002 n=6)
Tier1PfxSize/100_000       32.92 ± 0%   209.90 ± 5%  +537.61% (p=0.002 n=6)
Tier1PfxSize/200_000       22.15 ± 0%   209.20 ± 5%  +844.47% (p=0.002 n=6)
RandomPfx4Size/1_000       97.45 ± 2%   211.70 ± 5%  +117.24% (p=0.002 n=6)
RandomPfx4Size/10_000      68.19 ± 0%   206.80 ± 5%  +203.27% (p=0.002 n=6)
RandomPfx4Size/100_000     76.62 ± 0%   199.40 ± 6%  +160.25% (p=0.002 n=6)
RandomPfx4Size/200_000     70.78 ± 0%   194.60 ± 7%  +174.94% (p=0.002 n=6)
RandomPfx6Size/1_000       99.34 ± 2%   227.50 ± 8%  +129.01% (p=0.002 n=6)
RandomPfx6Size/10_000      75.99 ± 0%   223.40 ± 8%  +193.99% (p=0.002 n=6)
RandomPfx6Size/100_000     99.76 ± 0%   219.70 ± 8%  +120.23% (p=0.002 n=6)
RandomPfx6Size/200_000     96.09 ± 0%   218.00 ± 8%  +126.87% (p=0.002 n=6)
RandomPfxSize/1_000        100.4 ± 2%    214.9 ± 5%  +114.04% (p=0.002 n=6)
RandomPfxSize/10_000       70.53 ± 0%   209.60 ± 5%  +197.18% (p=0.002 n=6)
RandomPfxSize/100_000      76.30 ± 0%   203.30 ± 7%  +166.45% (p=0.002 n=6)
RandomPfxSize/200_000      75.96 ± 0%   199.30 ± 7%  +162.37% (p=0.002 n=6)
geomean                    72.42         210.6       +190.83%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidranger/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                 │
Tier1PfxSize/1_000         99.82 ± 2%   539.20 ± 3%   +440.17% (p=0.002 n=6)
Tier1PfxSize/10_000        73.46 ± 0%   533.80 ± 3%   +626.65% (p=0.002 n=6)
Tier1PfxSize/100_000       32.92 ± 0%   527.20 ± 2%  +1501.46% (p=0.002 n=6)
Tier1PfxSize/200_000       22.15 ± 0%   522.20 ± 2%  +2257.56% (p=0.002 n=6)
RandomPfx4Size/1_000       97.45 ± 2%   526.30 ± 3%   +440.07% (p=0.002 n=6)
RandomPfx4Size/10_000      68.19 ± 0%   514.30 ± 3%   +654.22% (p=0.002 n=6)
RandomPfx4Size/100_000     76.62 ± 0%   479.90 ± 3%   +526.34% (p=0.002 n=6)
RandomPfx4Size/200_000     70.78 ± 0%   463.20 ± 3%   +554.42% (p=0.002 n=6)
RandomPfx6Size/1_000       99.34 ± 2%   593.20 ± 0%   +497.14% (p=0.002 n=6)
RandomPfx6Size/10_000      75.99 ± 0%   585.90 ± 0%   +671.02% (p=0.002 n=6)
RandomPfx6Size/100_000     99.76 ± 0%   574.20 ± 0%   +475.58% (p=0.002 n=6)
RandomPfx6Size/200_000     96.09 ± 0%   570.00 ± 0%   +493.19% (p=0.002 n=6)
RandomPfxSize/1_000        100.4 ± 2%    537.7 ± 2%   +435.56% (p=0.002 n=6)
RandomPfxSize/10_000       70.53 ± 0%   526.40 ± 2%   +646.35% (p=0.002 n=6)
RandomPfxSize/100_000      76.30 ± 0%   498.50 ± 2%   +553.34% (p=0.002 n=6)
RandomPfxSize/200_000      75.96 ± 0%   484.50 ± 2%   +537.84% (p=0.002 n=6)
geomean                    72.42         528.5        +629.82%
```

## lpm (longest-prefix-match)

For longest-prefix-match, `bart` and `art` are the champions.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │               art/lpm.bm               │
                                     │    sec/op     │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            27.57n ±  71%    40.76n ± 20%   +47.80% (p=0.009 n=10)
LpmTier1Pfxs/RandomMatchIP6            24.44n ±  71%    43.23n ± 13%   +76.92% (p=0.001 n=10)
LpmTier1Pfxs/RandomMissIP4             35.48n ±  61%    25.55n ± 12%         ~ (p=0.072 n=10)
LpmTier1Pfxs/RandomMissIP6             6.427n ±   2%   25.515n ±  0%  +297.00% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     6.979n ±   1%   40.760n ±  4%  +484.04% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     8.111n ± 210%   41.000n ±  1%  +405.49% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      18.54n ±  15%    25.47n ±  1%   +37.42% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      23.14n ±   5%    25.49n ±  1%   +10.16% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    17.45n ±  49%    39.95n ±  2%  +128.94% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    22.20n ±  53%    42.32n ±  0%   +90.67% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     27.53n ±  10%    25.55n ± 12%         ~ (p=0.190 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     22.68n ±  29%    25.52n ±  1%         ~ (p=0.305 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   11.79n ± 169%    46.68n ± 16%  +296.14% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   22.31n ±  54%    42.27n ±  4%   +89.44% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    39.72n ±  27%    27.29n ±  7%   -31.30% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    27.46n ±  26%    25.44n ± 13%    -7.34% (p=0.014 n=10)
geomean                                18.87n           32.90n         +74.33%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │              cidrtree/lpm.bm               │
                                     │    sec/op     │      sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4            27.57n ±  71%    826.40n ±   32%  +2896.92% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            24.44n ±  71%    991.45n ±   35%  +3957.50% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             35.48n ±  61%   1059.00n ±   97%  +2884.78% (p=0.009 n=10)
LpmTier1Pfxs/RandomMissIP6             6.427n ±   2%    39.065n ± 1641%   +507.83% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     6.979n ±   1%   331.550n ±   51%  +4650.68% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     8.111n ± 210%   456.550n ±   20%  +5528.78% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      18.54n ±  15%    459.55n ±   27%  +2379.36% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      23.14n ±   5%    574.55n ±   43%  +2382.93% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    17.45n ±  49%   1106.50n ±   19%  +6240.97% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    22.20n ±  53%    733.70n ±   30%  +3205.70% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     27.53n ±  10%    695.85n ±   20%  +2428.07% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     22.68n ±  29%    852.55n ±   11%  +3659.04% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   11.79n ± 169%   1003.50n ±   28%  +8415.06% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   22.31n ±  54%   1391.00n ±   77%  +6134.87% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    39.72n ±  27%   1039.00n ±   29%  +2516.14% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    27.46n ±  26%   1395.50n ±   29%  +4982.86% (p=0.000 n=10)
geomean                                18.87n            660.8n          +3401.55%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │             critbitgo/lpm.bm             │
                                     │    sec/op     │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            27.57n ±  71%    122.75n ± 28%   +345.15% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            24.44n ±  71%    191.25n ± 23%   +682.69% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             35.48n ±  61%    634.90n ± 32%  +1689.46% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             6.427n ±   2%   426.100n ± 53%  +6529.84% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     6.979n ±   1%   110.200n ± 13%  +1479.02% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     8.111n ± 210%   119.450n ± 21%  +1372.69% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      18.54n ±  15%    176.10n ± 15%   +850.09% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      23.14n ±   5%    162.10n ± 14%   +600.52% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    17.45n ±  49%    121.60n ± 54%   +596.85% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    22.20n ±  53%    152.25n ±  9%   +585.97% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     27.53n ±  10%    256.55n ± 26%   +832.06% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     22.68n ±  29%    244.05n ± 18%   +976.06% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   11.79n ± 169%    175.80n ± 41%  +1391.73% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   22.31n ±  54%    191.25n ± 32%   +757.24% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    39.72n ±  27%    548.30n ± 34%  +1280.59% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    27.46n ±  26%    492.55n ± 17%  +1694.03% (p=0.000 n=10)
geomean                                18.87n            217.5n        +1052.69%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │              lpmtrie/lpm.bm               │
                                     │    sec/op     │     sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4            27.57n ±  71%    210.85n ±  11%   +664.64% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            24.44n ±  71%    214.05n ±  67%   +776.00% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             35.48n ±  61%    108.97n ± 130%          ~ (p=0.063 n=10)
LpmTier1Pfxs/RandomMissIP6             6.427n ±   2%    12.935n ±   4%   +101.26% (p=0.001 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     6.979n ±   1%   134.650n ±   6%  +1829.36% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     8.111n ± 210%   100.550n ±   8%  +1139.67% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      18.54n ±  15%    128.75n ±  28%   +594.63% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      23.14n ±   5%     83.77n ±  13%   +261.99% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    17.45n ±  49%    149.15n ±   6%   +754.73% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    22.20n ±  53%    132.10n ±   6%   +495.18% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     27.53n ±  10%    141.55n ±   7%   +414.26% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     22.68n ±  29%    127.35n ±  11%   +461.51% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   11.79n ± 169%    175.30n ±   8%  +1387.48% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   22.31n ±  54%    159.65n ±   7%   +615.60% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    39.72n ±  27%    190.40n ±   5%   +379.42% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    27.46n ±  26%    179.20n ±  15%   +552.70% (p=0.000 n=10)
geomean                                18.87n            124.1n          +557.45%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │             cidranger/lpm.bm             │
                                     │    sec/op     │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            27.57n ±  71%    168.95n ± 34%   +512.69% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            24.44n ±  71%    206.35n ± 38%   +744.49% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             35.48n ±  61%    195.75n ± 50%   +451.72% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             6.427n ±   2%    59.765n ± 14%   +829.91% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     6.979n ±   1%    81.050n ± 50%  +1061.34% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     8.111n ± 210%   130.900n ±  7%  +1513.86% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      18.54n ±  15%    112.75n ±  9%   +508.31% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      23.14n ±   5%    137.70n ± 31%   +495.07% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    17.45n ±  49%    146.75n ± 33%   +740.97% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    22.20n ±  53%    167.10n ± 26%   +652.87% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     27.53n ±  10%    170.25n ± 19%   +518.53% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     22.68n ±  29%    168.20n ±  9%   +641.62% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   11.79n ± 169%    145.60n ± 40%  +1135.47% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   22.31n ±  54%    166.15n ± 31%   +644.73% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    39.72n ±  27%    194.15n ± 11%   +388.86% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    27.46n ±  26%    246.75n ± 19%   +798.74% (p=0.000 n=10)
geomean                                18.87n            148.3n         +685.82%
```

## update, insert/delete

`bart` is by far the fastest algorithm for inserts and one of the fastest for delete.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │              art/update.bm              │
                         │   sec/route    │   sec/route     vs base                 │
InsertRandomPfxs/1_000        43.72n ± 3%   280.35n ±   3%   +541.31% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.92n ± 0%   577.05n ±   9%   +950.71% (p=0.002 n=6)
InsertRandomPfxs/100_000      148.4n ± 3%   1723.0n ±  24%  +1061.05% (p=0.002 n=6)
InsertRandomPfxs/200_000      298.8n ± 1%   1817.0n ± 198%   +508.00% (p=0.002 n=6)
DeleteRandomPfxs/1_000        19.02n ± 5%    17.62n ±   1%     -7.34% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.45n ± 1%    16.90n ±   0%     -8.37% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.47n ± 0%    17.52n ±   3%     -5.20% (p=0.002 n=6)
DeleteRandomPfxs/200_000      20.23n ± 1%    19.46n ±   4%     -3.76% (p=0.002 n=6)
geomean                       43.96n         122.7n          +179.13%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │            cidrtree/update.bm            │
                         │   sec/route    │   sec/route     vs base                  │
InsertRandomPfxs/1_000        43.72n ± 3%   1320.50n ± 44%   +2920.70% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.92n ± 0%   1971.50n ±  4%   +3489.77% (p=0.002 n=6)
InsertRandomPfxs/100_000      148.4n ± 3%    3406.5n ±  4%   +2195.49% (p=0.002 n=6)
InsertRandomPfxs/200_000      298.8n ± 1%    4218.5n ±  3%   +1311.58% (p=0.002 n=6)
DeleteRandomPfxs/1_000        19.02n ± 5%     14.76n ±  1%     -22.38% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.45n ± 1%     15.15n ±  0%     -17.89% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.47n ± 0%     22.90n ±  4%     +23.95% (p=0.002 n=6)
DeleteRandomPfxs/200_000      20.23n ± 1%   3381.50n ± 63%  +16619.41% (p=0.002 n=6)
geomean                       43.96n          399.4n          +808.51%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │         critbitgo/update.bm          │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000        43.72n ± 3%   176.25n ± 1%  +303.18% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.92n ± 0%   245.95n ± 1%  +347.83% (p=0.002 n=6)
InsertRandomPfxs/100_000      148.4n ± 3%    614.1n ± 1%  +313.78% (p=0.002 n=6)
InsertRandomPfxs/200_000      298.8n ± 1%    797.0n ± 2%  +166.67% (p=0.002 n=6)
DeleteRandomPfxs/1_000        19.02n ± 5%    71.63n ± 4%  +276.70% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.45n ± 1%    74.20n ± 2%  +302.17% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.47n ± 0%    81.35n ± 1%  +340.32% (p=0.002 n=6)
DeleteRandomPfxs/200_000      20.23n ± 1%    86.52n ± 0%  +327.79% (p=0.002 n=6)
geomean                       43.96n         172.8n       +292.96%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           lpmtrie/update.bm           │
                         │   sec/route    │  sec/route    vs base                 │
InsertRandomPfxs/1_000        43.72n ± 3%   400.10n ± 3%   +815.25% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.92n ± 0%   492.45n ± 6%   +796.67% (p=0.002 n=6)
InsertRandomPfxs/100_000      148.4n ± 3%   1255.0n ± 2%   +745.69% (p=0.002 n=6)
InsertRandomPfxs/200_000      298.8n ± 1%   1317.5n ± 3%   +340.86% (p=0.002 n=6)
DeleteRandomPfxs/1_000        19.02n ± 5%    75.38n ± 2%   +296.42% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.45n ± 1%   149.15n ± 3%   +708.40% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.47n ± 0%   324.65n ± 2%  +1657.24% (p=0.002 n=6)
DeleteRandomPfxs/200_000      20.23n ± 1%   479.85n ± 8%  +2272.56% (p=0.002 n=6)
geomean                       43.96n         393.1n        +794.21%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           cidranger/update.bm            │
                         │   sec/route    │   sec/route     vs base                  │
InsertRandomPfxs/1_000        43.72n ± 3%   4596.00n ±  4%  +10413.55% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.92n ± 0%   6859.00n ±  3%  +12389.08% (p=0.002 n=6)
InsertRandomPfxs/100_000      148.4n ± 3%   11253.0n ±  4%   +7482.88% (p=0.002 n=6)
InsertRandomPfxs/200_000      298.8n ± 1%   13186.0n ±  4%   +4312.25% (p=0.002 n=6)
DeleteRandomPfxs/1_000        19.02n ± 5%     90.89n ±  3%    +377.99% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.45n ± 1%     92.27n ±  1%    +400.14% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.47n ± 0%    145.05n ±  2%    +685.12% (p=0.002 n=6)
DeleteRandomPfxs/200_000      20.23n ± 1%   1280.50n ± 32%   +6231.27% (p=0.002 n=6)
geomean                       43.96n          1.282µ         +2815.50%
```
