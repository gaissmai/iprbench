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
RandomPfx4Size/1_000       97.63 ± 2%   7478.00 ± 0%  +7559.53% (p=0.002 n=6)
RandomPfx4Size/10_000      68.82 ± 0%   6097.00 ± 0%  +8759.34% (p=0.002 n=6)
RandomPfx4Size/100_000     76.91 ± 0%   5471.00 ± 0%  +7013.51% (p=0.002 n=6)
RandomPfx4Size/200_000     70.74 ± 0%   4827.00 ± 0%  +6723.58% (p=0.002 n=6)
RandomPfx6Size/1_000       100.1 ± 2%    7877.0 ± 0%  +7769.13% (p=0.002 n=6)
RandomPfx6Size/10_000      76.12 ± 0%   6811.00 ± 0%  +8847.71% (p=0.002 n=6)
RandomPfx6Size/100_000     99.64 ± 0%   7829.00 ± 0%  +7757.29% (p=0.002 n=6)
RandomPfx6Size/200_000     96.09 ± 0%   7599.00 ± 0%  +7808.21% (p=0.002 n=6)
RandomPfxSize/1_000        102.0 ± 2%    7498.0 ± 0%  +7250.98% (p=0.002 n=6)
RandomPfxSize/10_000       70.73 ± 0%   6214.00 ± 0%  +8685.52% (p=0.002 n=6)
RandomPfxSize/100_000      76.16 ± 0%   5785.00 ± 0%  +7495.85% (p=0.002 n=6)
RandomPfxSize/200_000      76.17 ± 0%   5458.00 ± 0%  +7065.55% (p=0.002 n=6)
geomean                    72.61        5.173Ki       +7195.67%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidrtree/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         99.82 ± 2%    69.06 ± 3%   -30.82% (p=0.002 n=6)
Tier1PfxSize/10_000        73.46 ± 0%    64.35 ± 0%   -12.40% (p=0.002 n=6)
Tier1PfxSize/100_000       32.92 ± 0%    64.03 ± 0%   +94.50% (p=0.002 n=6)
Tier1PfxSize/200_000       22.15 ± 0%    64.02 ± 0%  +189.03% (p=0.002 n=6)
RandomPfx4Size/1_000       97.63 ± 2%    67.89 ± 3%   -30.46% (p=0.002 n=6)
RandomPfx4Size/10_000      68.82 ± 0%    64.35 ± 0%    -6.50% (p=0.002 n=6)
RandomPfx4Size/100_000     76.91 ± 0%    64.03 ± 0%   -16.75% (p=0.002 n=6)
RandomPfx4Size/200_000     70.74 ± 0%    64.02 ± 0%    -9.50% (p=0.002 n=6)
RandomPfx6Size/1_000      100.10 ± 2%    68.06 ± 2%   -32.01% (p=0.002 n=6)
RandomPfx6Size/10_000      76.12 ± 0%    64.35 ± 0%   -15.46% (p=0.002 n=6)
RandomPfx6Size/100_000     99.64 ± 0%    64.03 ± 0%   -35.74% (p=0.002 n=6)
RandomPfx6Size/200_000     96.09 ± 0%    64.02 ± 0%   -33.37% (p=0.002 n=6)
RandomPfxSize/1_000       102.00 ± 2%    67.89 ± 3%   -33.44% (p=0.002 n=6)
RandomPfxSize/10_000       70.73 ± 0%    64.39 ± 0%    -8.96% (p=0.002 n=6)
RandomPfxSize/100_000      76.16 ± 0%    64.03 ± 0%   -15.93% (p=0.002 n=6)
RandomPfxSize/200_000      76.17 ± 0%    64.02 ± 0%   -15.95% (p=0.002 n=6)
geomean                    72.61         65.13        -10.29%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          critbitgo/size.bm          │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         99.82 ± 2%   119.10 ± 2%   +19.31% (p=0.002 n=6)
Tier1PfxSize/10_000        73.46 ± 0%   114.70 ± 0%   +56.14% (p=0.002 n=6)
Tier1PfxSize/100_000       32.92 ± 0%   114.40 ± 0%  +247.51% (p=0.002 n=6)
Tier1PfxSize/200_000       22.15 ± 0%   114.40 ± 0%  +416.48% (p=0.002 n=6)
RandomPfx4Size/1_000       97.63 ± 2%   115.90 ± 2%   +18.71% (p=0.002 n=6)
RandomPfx4Size/10_000      68.82 ± 0%   112.30 ± 0%   +63.18% (p=0.002 n=6)
RandomPfx4Size/100_000     76.91 ± 0%   112.00 ± 0%   +45.62% (p=0.002 n=6)
RandomPfx4Size/200_000     70.74 ± 0%   112.00 ± 0%   +58.33% (p=0.002 n=6)
RandomPfx6Size/1_000       100.1 ± 2%    132.0 ± 1%   +31.87% (p=0.002 n=6)
RandomPfx6Size/10_000      76.12 ± 0%   128.30 ± 0%   +68.55% (p=0.002 n=6)
RandomPfx6Size/100_000     99.64 ± 0%   128.00 ± 0%   +28.46% (p=0.002 n=6)
RandomPfx6Size/200_000     96.09 ± 0%   128.00 ± 0%   +33.21% (p=0.002 n=6)
RandomPfxSize/1_000        102.0 ± 2%    118.8 ± 1%   +16.47% (p=0.002 n=6)
RandomPfxSize/10_000       70.73 ± 0%   115.50 ± 0%   +63.30% (p=0.002 n=6)
RandomPfxSize/100_000      76.16 ± 0%   115.20 ± 0%   +51.26% (p=0.002 n=6)
RandomPfxSize/200_000      76.17 ± 0%   115.20 ± 0%   +51.24% (p=0.002 n=6)
geomean                    72.61         118.3        +62.95%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │           lpmtrie/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         99.82 ± 2%   214.90 ± 5%  +115.29% (p=0.002 n=6)
Tier1PfxSize/10_000        73.46 ± 0%   210.50 ± 5%  +186.55% (p=0.002 n=6)
Tier1PfxSize/100_000       32.92 ± 0%   209.90 ± 5%  +537.61% (p=0.002 n=6)
Tier1PfxSize/200_000       22.15 ± 0%   209.20 ± 5%  +844.47% (p=0.002 n=6)
RandomPfx4Size/1_000       97.63 ± 2%   211.70 ± 5%  +116.84% (p=0.002 n=6)
RandomPfx4Size/10_000      68.82 ± 0%   206.80 ± 5%  +200.49% (p=0.002 n=6)
RandomPfx4Size/100_000     76.91 ± 0%   199.40 ± 6%  +159.26% (p=0.002 n=6)
RandomPfx4Size/200_000     70.74 ± 0%   194.60 ± 7%  +175.09% (p=0.002 n=6)
RandomPfx6Size/1_000       100.1 ± 2%    227.5 ± 8%  +127.27% (p=0.002 n=6)
RandomPfx6Size/10_000      76.12 ± 0%   223.40 ± 8%  +193.48% (p=0.002 n=6)
RandomPfx6Size/100_000     99.64 ± 0%   219.70 ± 8%  +120.49% (p=0.002 n=6)
RandomPfx6Size/200_000     96.09 ± 0%   218.00 ± 8%  +126.87% (p=0.002 n=6)
RandomPfxSize/1_000        102.0 ± 2%    214.9 ± 5%  +110.69% (p=0.002 n=6)
RandomPfxSize/10_000       70.73 ± 0%   209.60 ± 5%  +196.34% (p=0.002 n=6)
RandomPfxSize/100_000      76.16 ± 0%   203.30 ± 7%  +166.94% (p=0.002 n=6)
RandomPfxSize/200_000      76.17 ± 0%   199.30 ± 7%  +161.65% (p=0.002 n=6)
geomean                    72.61         210.6       +190.07%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidranger/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                 │
Tier1PfxSize/1_000         99.82 ± 2%   539.20 ± 3%   +440.17% (p=0.002 n=6)
Tier1PfxSize/10_000        73.46 ± 0%   533.80 ± 3%   +626.65% (p=0.002 n=6)
Tier1PfxSize/100_000       32.92 ± 0%   527.20 ± 2%  +1501.46% (p=0.002 n=6)
Tier1PfxSize/200_000       22.15 ± 0%   522.20 ± 2%  +2257.56% (p=0.002 n=6)
RandomPfx4Size/1_000       97.63 ± 2%   526.30 ± 3%   +439.08% (p=0.002 n=6)
RandomPfx4Size/10_000      68.82 ± 0%   514.30 ± 3%   +647.31% (p=0.002 n=6)
RandomPfx4Size/100_000     76.91 ± 0%   479.90 ± 3%   +523.98% (p=0.002 n=6)
RandomPfx4Size/200_000     70.74 ± 0%   463.20 ± 3%   +554.79% (p=0.002 n=6)
RandomPfx6Size/1_000       100.1 ± 2%    593.2 ± 0%   +492.61% (p=0.002 n=6)
RandomPfx6Size/10_000      76.12 ± 0%   585.90 ± 0%   +669.71% (p=0.002 n=6)
RandomPfx6Size/100_000     99.64 ± 0%   574.20 ± 0%   +476.27% (p=0.002 n=6)
RandomPfx6Size/200_000     96.09 ± 0%   570.00 ± 0%   +493.19% (p=0.002 n=6)
RandomPfxSize/1_000        102.0 ± 2%    537.7 ± 2%   +427.16% (p=0.002 n=6)
RandomPfxSize/10_000       70.73 ± 0%   526.40 ± 2%   +644.24% (p=0.002 n=6)
RandomPfxSize/100_000      76.16 ± 0%   498.50 ± 2%   +554.54% (p=0.002 n=6)
RandomPfxSize/200_000      76.17 ± 0%   484.50 ± 2%   +536.08% (p=0.002 n=6)
geomean                    72.61         528.5        +627.91%
```

## lpm (longest-prefix-match)

For longest-prefix-match, `bart` and `art` are the champions.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │              art/lpm.bm               │
                                     │    sec/op    │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4            31.35n ±  8%   40.71n ± 20%   +29.86% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            33.87n ± 39%   43.23n ± 13%   +27.64% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             33.00n ± 41%   25.60n ± 12%         ~ (p=0.469 n=10)
LpmTier1Pfxs/RandomMissIP6             11.41n ±  5%   25.53n ±  1%  +123.85% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     21.06n ±  6%   39.89n ±  1%   +89.43% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     23.46n ±  2%   41.02n ±  1%   +74.79% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      21.46n ± 23%   25.46n ±  1%         ~ (p=0.108 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      23.33n ± 51%   25.45n ±  1%    +9.09% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    27.36n ± 50%   40.75n ±  6%   +48.94% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    28.42n ±  7%   42.26n ±  2%   +48.70% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     31.68n ±  7%   25.46n ±  0%   -19.61% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     26.80n ±  7%   25.45n ±  0%    -5.07% (p=0.005 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   29.74n ± 55%   49.08n ± 15%   +65.01% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   17.07n ± 81%   42.26n ±  2%  +147.64% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    41.74n ± 18%   25.96n ±  7%   -37.81% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    33.26n ± 22%   25.61n ±  3%   -22.97% (p=0.000 n=10)
geomean                                26.09n         32.89n         +26.09%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │              cidrtree/lpm.bm               │
                                     │    sec/op    │      sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4            31.35n ±  8%    826.40n ±   32%  +2536.04% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            33.87n ± 39%    991.45n ±   35%  +2827.22% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             33.00n ± 41%   1059.00n ±   97%  +3108.60% (p=0.003 n=10)
LpmTier1Pfxs/RandomMissIP6             11.41n ±  5%     39.07n ± 1641%   +242.53% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     21.06n ±  6%    331.55n ±   51%  +1474.31% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     23.46n ±  2%    456.55n ±   20%  +1845.66% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      21.46n ± 23%    459.55n ±   27%  +2041.92% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      23.33n ± 51%    574.55n ±   43%  +2362.71% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    27.36n ± 50%   1106.50n ±   19%  +3944.23% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    28.42n ±  7%    733.70n ±   30%  +2481.63% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     31.68n ±  7%    695.85n ±   20%  +2096.84% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     26.80n ±  7%    852.55n ±   11%  +3080.56% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   29.74n ± 55%   1003.50n ±   28%  +3274.24% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   17.07n ± 81%   1391.00n ±   77%  +8051.19% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    41.74n ± 18%   1039.00n ±   29%  +2388.92% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    33.26n ± 22%   1395.50n ±   29%  +4096.36% (p=0.000 n=10)
geomean                                26.09n           660.8n          +2432.98%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            critbitgo/lpm.bm             │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            31.35n ±  8%   122.75n ± 28%   +291.55% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            33.87n ± 39%   191.25n ± 23%   +464.66% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             33.00n ± 41%   634.90n ± 32%  +1823.65% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             11.41n ±  5%   426.10n ± 53%  +3636.08% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     21.06n ±  6%   110.20n ± 13%   +423.27% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     23.46n ±  2%   119.45n ± 21%   +409.06% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      21.46n ± 23%   176.10n ± 15%   +720.79% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      23.33n ± 51%   162.10n ± 14%   +594.81% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    27.36n ± 50%   121.60n ± 54%   +344.44% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    28.42n ±  7%   152.25n ±  9%   +435.71% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     31.68n ±  7%   256.55n ± 26%   +709.94% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     26.80n ±  7%   244.05n ± 18%   +810.46% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   29.74n ± 55%   175.80n ± 41%   +491.12% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   17.07n ± 81%   191.25n ± 32%  +1020.71% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    41.74n ± 18%   548.30n ± 34%  +1213.45% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    33.26n ± 22%   492.55n ± 17%  +1381.13% (p=0.000 n=10)
geomean                                26.09n          217.5n         +733.84%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             lpmtrie/lpm.bm              │
                                     │    sec/op    │     sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            31.35n ±  8%   210.85n ±  11%  +572.57% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            33.87n ± 39%   214.05n ±  67%  +531.98% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             33.00n ± 41%   108.97n ± 130%         ~ (p=0.102 n=10)
LpmTier1Pfxs/RandomMissIP6             11.41n ±  5%    12.94n ±   4%   +13.42% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     21.06n ±  6%   134.65n ±   6%  +539.36% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     23.46n ±  2%   100.55n ±   8%  +328.51% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      21.46n ± 23%   128.75n ±  28%  +500.09% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      23.33n ± 51%    83.77n ±  13%  +259.04% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    27.36n ± 50%   149.15n ±   6%  +445.14% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    28.42n ±  7%   132.10n ±   6%  +364.81% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     31.68n ±  7%   141.55n ±   7%  +346.88% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     26.80n ±  7%   127.35n ±  11%  +375.10% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   29.74n ± 55%   175.30n ±   8%  +489.44% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   17.07n ± 81%   159.65n ±   7%  +835.54% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    41.74n ± 18%   190.40n ±   5%  +356.10% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    33.26n ± 22%   179.20n ±  15%  +438.87% (p=0.000 n=10)
geomean                                26.09n          124.1n         +375.59%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            cidranger/lpm.bm            │
                                     │    sec/op    │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            31.35n ±  8%   168.95n ± 34%  +438.92% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            33.87n ± 39%   206.35n ± 38%  +509.24% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             33.00n ± 41%   195.75n ± 50%  +493.09% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             11.41n ±  5%    59.77n ± 14%  +424.02% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     21.06n ±  6%    81.05n ± 50%  +284.85% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     23.46n ±  2%   130.90n ±  7%  +457.85% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      21.46n ± 23%   112.75n ±  9%  +425.52% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      23.33n ± 51%   137.70n ± 31%  +490.23% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    27.36n ± 50%   146.75n ± 33%  +436.37% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    28.42n ±  7%   167.10n ± 26%  +487.97% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     31.68n ±  7%   170.25n ± 19%  +437.49% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     26.80n ±  7%   168.20n ±  9%  +527.49% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   29.74n ± 55%   145.60n ± 40%  +389.58% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   17.07n ± 81%   166.15n ± 31%  +873.63% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    41.74n ± 18%   194.15n ± 11%  +365.09% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    33.26n ± 22%   246.75n ± 19%  +641.99% (p=0.000 n=10)
geomean                                26.09n          148.3n        +468.46%
```

## update, insert/delete

`bart` is by far the fastest algorithm for inserts and one of the fastest for delete.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │              art/update.bm              │
                         │   sec/route    │   sec/route     vs base                 │
InsertRandomPfxs/1_000       43.43n ±  1%   280.35n ±   3%   +545.45% (p=0.002 n=6)
InsertRandomPfxs/10_000      54.54n ±  2%   577.05n ±   9%   +958.13% (p=0.002 n=6)
InsertRandomPfxs/100_000     143.6n ±  3%   1723.0n ±  24%  +1099.44% (p=0.002 n=6)
InsertRandomPfxs/200_000     298.1n ±  2%   1817.0n ± 198%   +509.42% (p=0.002 n=6)
DeleteRandomPfxs/1_000       18.82n ±  1%    17.62n ±   1%     -6.40% (p=0.002 n=6)
DeleteRandomPfxs/10_000      18.03n ±  1%    16.90n ±   0%     -6.21% (p=0.002 n=6)
DeleteRandomPfxs/100_000     18.01n ± 17%    17.52n ±   3%     -2.75% (p=0.035 n=6)
DeleteRandomPfxs/200_000     19.36n ± 15%    19.46n ±   4%          ~ (p=0.485 n=6)
geomean                      43.14n          122.7n          +184.45%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │            cidrtree/update.bm            │
                         │   sec/route    │   sec/route     vs base                  │
InsertRandomPfxs/1_000       43.43n ±  1%   1320.50n ± 44%   +2940.17% (p=0.002 n=6)
InsertRandomPfxs/10_000      54.54n ±  2%   1971.50n ±  4%   +3515.11% (p=0.002 n=6)
InsertRandomPfxs/100_000     143.6n ±  3%    3406.5n ±  4%   +2271.39% (p=0.002 n=6)
InsertRandomPfxs/200_000     298.1n ±  2%    4218.5n ±  3%   +1314.89% (p=0.002 n=6)
DeleteRandomPfxs/1_000       18.82n ±  1%     14.76n ±  1%     -21.59% (p=0.002 n=6)
DeleteRandomPfxs/10_000      18.03n ±  1%     15.15n ±  0%     -15.95% (p=0.002 n=6)
DeleteRandomPfxs/100_000     18.01n ± 17%     22.90n ±  4%     +27.15% (p=0.002 n=6)
DeleteRandomPfxs/200_000     19.36n ± 15%   3381.50n ± 63%  +17370.94% (p=0.002 n=6)
geomean                      43.14n           399.4n          +825.84%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │         critbitgo/update.bm          │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000       43.43n ±  1%   176.25n ± 1%  +305.78% (p=0.002 n=6)
InsertRandomPfxs/10_000      54.54n ±  2%   245.95n ± 1%  +350.99% (p=0.002 n=6)
InsertRandomPfxs/100_000     143.6n ±  3%    614.1n ± 1%  +327.46% (p=0.002 n=6)
InsertRandomPfxs/200_000     298.1n ±  2%    797.0n ± 2%  +167.30% (p=0.002 n=6)
DeleteRandomPfxs/1_000       18.82n ±  1%    71.63n ± 4%  +280.50% (p=0.002 n=6)
DeleteRandomPfxs/10_000      18.03n ±  1%    74.20n ± 2%  +311.65% (p=0.002 n=6)
DeleteRandomPfxs/100_000     18.01n ± 17%    81.35n ± 1%  +351.69% (p=0.002 n=6)
DeleteRandomPfxs/200_000     19.36n ± 15%    86.52n ± 0%  +347.02% (p=0.002 n=6)
geomean                      43.14n          172.8n       +300.45%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           lpmtrie/update.bm           │
                         │   sec/route    │  sec/route    vs base                 │
InsertRandomPfxs/1_000       43.43n ±  1%   400.10n ± 3%   +821.15% (p=0.002 n=6)
InsertRandomPfxs/10_000      54.54n ±  2%   492.45n ± 6%   +803.00% (p=0.002 n=6)
InsertRandomPfxs/100_000     143.6n ±  3%   1255.0n ± 2%   +773.65% (p=0.002 n=6)
InsertRandomPfxs/200_000     298.1n ±  2%   1317.5n ± 3%   +341.89% (p=0.002 n=6)
DeleteRandomPfxs/1_000       18.82n ±  1%    75.38n ± 2%   +300.42% (p=0.002 n=6)
DeleteRandomPfxs/10_000      18.03n ±  1%   149.15n ± 3%   +727.46% (p=0.002 n=6)
DeleteRandomPfxs/100_000     18.01n ± 17%   324.65n ± 2%  +1702.61% (p=0.002 n=6)
DeleteRandomPfxs/200_000     19.36n ± 15%   479.85n ± 8%  +2379.20% (p=0.002 n=6)
geomean                      43.14n          393.1n        +811.26%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           cidranger/update.bm            │
                         │   sec/route    │   sec/route     vs base                  │
InsertRandomPfxs/1_000       43.43n ±  1%   4596.00n ±  4%  +10481.33% (p=0.002 n=6)
InsertRandomPfxs/10_000      54.54n ±  2%   6859.00n ±  3%  +12477.24% (p=0.002 n=6)
InsertRandomPfxs/100_000     143.6n ±  3%   11253.0n ±  4%   +7733.62% (p=0.002 n=6)
InsertRandomPfxs/200_000     298.1n ±  2%   13186.0n ±  4%   +4322.61% (p=0.002 n=6)
DeleteRandomPfxs/1_000       18.82n ±  1%     90.89n ±  3%    +382.82% (p=0.002 n=6)
DeleteRandomPfxs/10_000      18.03n ±  1%     92.27n ±  1%    +411.93% (p=0.002 n=6)
DeleteRandomPfxs/100_000     18.01n ± 17%    145.05n ±  2%    +705.39% (p=0.002 n=6)
DeleteRandomPfxs/200_000     19.36n ± 15%   1280.50n ± 32%   +6515.86% (p=0.002 n=6)
geomean                      43.14n           1.282µ         +2871.11%
```
