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
                                     │ bart/lpm.bm  │               art/lpm.bm               │
                                     │    sec/op    │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            26.24n ±  8%    40.76n ± 20%   +55.32% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            31.61n ± 29%    43.23n ± 13%   +36.76% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             21.26n ± 89%    25.55n ± 12%         ~ (p=1.000 n=10)
LpmTier1Pfxs/RandomMissIP6             6.313n ±  1%   25.515n ±  0%  +304.17% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     16.37n ± 28%    40.76n ±  4%  +149.07% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     16.19n ±  0%    41.00n ±  1%  +153.24% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      17.04n ± 20%    25.47n ±  1%   +49.47% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      15.54n ± 59%    25.49n ±  1%   +64.03% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    16.76n ± 45%    39.95n ±  2%  +138.37% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    16.49n ± 43%    42.32n ±  0%  +156.64% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     27.07n ±  6%    25.55n ± 12%         ~ (p=0.247 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     23.06n ± 18%    25.52n ±  1%         ~ (p=0.128 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   25.41n ± 55%    46.68n ± 16%   +83.76% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   16.90n ± 48%    42.27n ±  4%  +150.09% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    40.41n ±  8%    27.29n ±  7%   -32.49% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    27.12n ±  8%    25.44n ± 13%         ~ (p=0.074 n=10)
geomean                                19.96n          32.90n         +64.83%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │              cidrtree/lpm.bm               │
                                     │    sec/op    │      sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4            26.24n ±  8%    826.40n ±   32%  +3049.39% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            31.61n ± 29%    991.45n ±   35%  +3036.51% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             21.26n ± 89%   1059.00n ±   97%  +4880.01% (p=0.001 n=10)
LpmTier1Pfxs/RandomMissIP6             6.313n ±  1%    39.065n ± 1641%   +518.80% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     16.37n ± 28%    331.55n ±   51%  +1925.97% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     16.19n ±  0%    456.55n ±   20%  +2719.95% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      17.04n ± 20%    459.55n ±   27%  +2596.89% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      15.54n ± 59%    574.55n ±   43%  +3597.23% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    16.76n ± 45%   1106.50n ±   19%  +6502.03% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    16.49n ± 43%    733.70n ±   30%  +4349.36% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     27.07n ±  6%    695.85n ±   20%  +2470.08% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     23.06n ± 18%    852.55n ±   11%  +3597.90% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   25.41n ± 55%   1003.50n ±   28%  +3850.01% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   16.90n ± 48%   1391.00n ±   77%  +8130.77% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    40.41n ±  8%   1039.00n ±   29%  +2470.83% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    27.12n ±  8%   1395.50n ±   29%  +5045.65% (p=0.000 n=10)
geomean                                19.96n           660.8n          +3210.60%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             critbitgo/lpm.bm             │
                                     │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            26.24n ±  8%    122.75n ± 28%   +367.80% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            31.61n ± 29%    191.25n ± 23%   +505.03% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             21.26n ± 89%    634.90n ± 32%  +2885.66% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             6.313n ±  1%   426.100n ± 53%  +6649.56% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     16.37n ± 28%    110.20n ± 13%   +573.39% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     16.19n ±  0%    119.45n ± 21%   +637.80% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      17.04n ± 20%    176.10n ± 15%   +933.45% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      15.54n ± 59%    162.10n ± 14%   +943.11% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    16.76n ± 45%    121.60n ± 54%   +625.54% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    16.49n ± 43%    152.25n ±  9%   +823.29% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     27.07n ±  6%    256.55n ± 26%   +847.55% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     23.06n ± 18%    244.05n ± 18%   +958.56% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   25.41n ± 55%    175.80n ± 41%   +591.99% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   16.90n ± 48%    191.25n ± 32%  +1031.66% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    40.41n ±  8%    548.30n ± 34%  +1256.67% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    27.12n ±  8%    492.55n ± 17%  +1716.19% (p=0.000 n=10)
geomean                                19.96n           217.5n         +989.83%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             lpmtrie/lpm.bm              │
                                     │    sec/op    │     sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            26.24n ±  8%   210.85n ±  11%  +703.54% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            31.61n ± 29%   214.05n ±  67%  +577.16% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             21.26n ± 89%   108.97n ± 130%  +412.46% (p=0.007 n=10)
LpmTier1Pfxs/RandomMissIP6             6.313n ±  1%   12.935n ±   4%  +104.89% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     16.37n ± 28%   134.65n ±   6%  +722.79% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     16.19n ±  0%   100.55n ±   8%  +521.06% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      17.04n ± 20%   128.75n ±  28%  +655.58% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      15.54n ± 59%    83.77n ±  13%  +439.03% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    16.76n ± 45%   149.15n ±   6%  +789.92% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    16.49n ± 43%   132.10n ±   6%  +701.09% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     27.07n ±  6%   141.55n ±   7%  +422.81% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     23.06n ± 18%   127.35n ±  11%  +452.37% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   25.41n ± 55%   175.30n ±   8%  +590.02% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   16.90n ± 48%   159.65n ±   7%  +844.67% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    40.41n ±  8%   190.40n ±   5%  +371.11% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    27.12n ±  8%   179.20n ±  15%  +560.77% (p=0.000 n=10)
geomean                                19.96n          124.1n         +521.60%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            cidranger/lpm.bm            │
                                     │    sec/op    │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            26.24n ±  8%   168.95n ± 34%  +543.86% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            31.61n ± 29%   206.35n ± 38%  +552.80% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             21.26n ± 89%   195.75n ± 50%  +820.53% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             6.313n ±  1%   59.765n ± 14%  +846.70% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     16.37n ± 28%    81.05n ± 50%  +395.26% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     16.19n ±  0%   130.90n ±  7%  +708.52% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      17.04n ± 20%   112.75n ±  9%  +561.68% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      15.54n ± 59%   137.70n ± 31%  +786.10% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    16.76n ± 45%   146.75n ± 33%  +775.60% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    16.49n ± 43%   167.10n ± 26%  +913.34% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     27.07n ±  6%   170.25n ± 19%  +528.81% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     23.06n ± 18%   168.20n ±  9%  +629.56% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   25.41n ± 55%   145.60n ± 40%  +473.12% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   16.90n ± 48%   166.15n ± 31%  +883.14% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    40.41n ±  8%   194.15n ± 11%  +380.39% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    27.12n ±  8%   246.75n ± 19%  +809.85% (p=0.000 n=10)
geomean                                19.96n          148.3n        +642.97%
```

## update, insert/delete

`bart` is by far the fastest algorithm for inserts and one of the fastest for delete.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │              art/update.bm              │
                         │   sec/route    │   sec/route     vs base                 │
InsertRandomPfxs/1_000        43.22n ± 1%   280.35n ±   3%   +548.66% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.56n ± 1%   577.05n ±   9%   +957.55% (p=0.002 n=6)
InsertRandomPfxs/100_000      150.2n ± 4%   1723.0n ±  24%  +1046.76% (p=0.002 n=6)
InsertRandomPfxs/200_000      296.4n ± 1%   1817.0n ± 198%   +513.02% (p=0.002 n=6)
DeleteRandomPfxs/1_000        19.04n ± 1%    17.62n ±   1%     -7.46% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.47n ± 0%    16.90n ±   0%     -8.47% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.44n ± 2%    17.52n ±   3%     -5.02% (p=0.002 n=6)
DeleteRandomPfxs/200_000      20.23n ± 2%    19.46n ±   4%     -3.81% (p=0.002 n=6)
geomean                       43.89n         122.7n          +179.57%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │            cidrtree/update.bm            │
                         │   sec/route    │   sec/route     vs base                  │
InsertRandomPfxs/1_000        43.22n ± 1%   1320.50n ± 44%   +2955.30% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.56n ± 1%   1971.50n ±  4%   +3513.12% (p=0.002 n=6)
InsertRandomPfxs/100_000      150.2n ± 4%    3406.5n ±  4%   +2167.22% (p=0.002 n=6)
InsertRandomPfxs/200_000      296.4n ± 1%    4218.5n ±  3%   +1323.25% (p=0.002 n=6)
DeleteRandomPfxs/1_000        19.04n ± 1%     14.76n ±  1%     -22.48% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.47n ± 0%     15.15n ±  0%     -17.98% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.44n ± 2%     22.90n ±  4%     +24.19% (p=0.002 n=6)
DeleteRandomPfxs/200_000      20.23n ± 2%   3381.50n ± 63%  +16611.14% (p=0.002 n=6)
geomean                       43.89n          399.4n          +809.95%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │         critbitgo/update.bm          │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000        43.22n ± 1%   176.25n ± 1%  +307.80% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.56n ± 1%   245.95n ± 1%  +350.75% (p=0.002 n=6)
InsertRandomPfxs/100_000      150.2n ± 4%    614.1n ± 1%  +308.69% (p=0.002 n=6)
InsertRandomPfxs/200_000      296.4n ± 1%    797.0n ± 2%  +168.88% (p=0.002 n=6)
DeleteRandomPfxs/1_000        19.04n ± 1%    71.63n ± 4%  +276.21% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.47n ± 0%    74.20n ± 2%  +301.73% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.44n ± 2%    81.35n ± 1%  +341.16% (p=0.002 n=6)
DeleteRandomPfxs/200_000      20.23n ± 2%    86.52n ± 0%  +327.58% (p=0.002 n=6)
geomean                       43.89n         172.8n       +293.58%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           lpmtrie/update.bm           │
                         │   sec/route    │  sec/route    vs base                 │
InsertRandomPfxs/1_000        43.22n ± 1%   400.10n ± 3%   +825.73% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.56n ± 1%   492.45n ± 6%   +802.50% (p=0.002 n=6)
InsertRandomPfxs/100_000      150.2n ± 4%   1255.0n ± 2%   +735.27% (p=0.002 n=6)
InsertRandomPfxs/200_000      296.4n ± 1%   1317.5n ± 3%   +344.50% (p=0.002 n=6)
DeleteRandomPfxs/1_000        19.04n ± 1%    75.38n ± 2%   +295.90% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.47n ± 0%   149.15n ± 3%   +707.53% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.44n ± 2%   324.65n ± 2%  +1660.57% (p=0.002 n=6)
DeleteRandomPfxs/200_000      20.23n ± 2%   479.85n ± 8%  +2271.39% (p=0.002 n=6)
geomean                       43.89n         393.1n        +795.63%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           cidranger/update.bm            │
                         │   sec/route    │   sec/route     vs base                  │
InsertRandomPfxs/1_000        43.22n ± 1%   4596.00n ±  4%  +10533.97% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.56n ± 1%   6859.00n ±  3%  +12470.33% (p=0.002 n=6)
InsertRandomPfxs/100_000      150.2n ± 4%   11253.0n ±  4%   +7389.52% (p=0.002 n=6)
InsertRandomPfxs/200_000      296.4n ± 1%   13186.0n ±  4%   +4348.72% (p=0.002 n=6)
DeleteRandomPfxs/1_000        19.04n ± 1%     90.89n ±  3%    +377.36% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.47n ± 0%     92.27n ±  1%    +399.59% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.44n ± 2%    145.05n ±  2%    +686.61% (p=0.002 n=6)
DeleteRandomPfxs/200_000      20.23n ± 2%   1280.50n ± 32%   +6228.14% (p=0.002 n=6)
geomean                       43.89n          1.282µ         +2820.14%
```
