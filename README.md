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
Tier1PfxSize/1_000         101.0 ± 2%    7606.0 ± 0%  +7430.69% (p=0.002 n=6)
Tier1PfxSize/10_000        81.27 ± 0%   4905.00 ± 0%  +5935.44% (p=0.002 n=6)
Tier1PfxSize/100_000       47.41 ± 0%   1685.00 ± 0%  +3454.10% (p=0.002 n=6)
Tier1PfxSize/200_000       37.49 ± 0%   1114.00 ± 0%  +2871.46% (p=0.002 n=6)
RandomPfx4Size/1_000       100.3 ± 2%    7442.0 ± 0%  +7319.74% (p=0.002 n=6)
RandomPfx4Size/10_000      70.82 ± 0%   6028.00 ± 0%  +8411.72% (p=0.002 n=6)
RandomPfx4Size/100_000     83.74 ± 0%   5492.00 ± 0%  +6458.40% (p=0.002 n=6)
RandomPfx4Size/200_000     79.15 ± 0%   4849.00 ± 0%  +6026.34% (p=0.002 n=6)
RandomPfx6Size/1_000       100.7 ± 2%    7873.0 ± 0%  +7718.27% (p=0.002 n=6)
RandomPfx6Size/10_000      76.50 ± 0%   6850.00 ± 0%  +8854.25% (p=0.002 n=6)
RandomPfx6Size/100_000     101.1 ± 0%    7848.0 ± 0%  +7662.61% (p=0.002 n=6)
RandomPfx6Size/200_000     97.80 ± 0%   7619.00 ± 0%  +7690.39% (p=0.002 n=6)
RandomPfxSize/1_000        102.2 ± 2%    7645.0 ± 0%  +7380.43% (p=0.002 n=6)
RandomPfxSize/10_000       72.04 ± 0%   6255.00 ± 0%  +8582.68% (p=0.002 n=6)
RandomPfxSize/100_000      81.73 ± 0%   5824.00 ± 0%  +7025.90% (p=0.002 n=6)
RandomPfxSize/200_000      82.84 ± 0%   5471.00 ± 0%  +6504.30% (p=0.002 n=6)
geomean                    79.62        5.195Ki       +6581.13%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidrtree/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000        101.00 ± 2%    85.02 ± 2%   -15.82% (p=0.002 n=6)
Tier1PfxSize/10_000        81.27 ± 0%    80.35 ± 0%    -1.13% (p=0.002 n=6)
Tier1PfxSize/100_000       47.41 ± 0%    80.03 ± 0%   +68.80% (p=0.002 n=6)
Tier1PfxSize/200_000       37.49 ± 0%    80.02 ± 0%  +113.44% (p=0.002 n=6)
RandomPfx4Size/1_000      100.30 ± 2%    83.89 ± 2%   -16.36% (p=0.002 n=6)
RandomPfx4Size/10_000      70.82 ± 0%    80.35 ± 0%   +13.46% (p=0.002 n=6)
RandomPfx4Size/100_000     83.74 ± 0%    80.03 ± 0%    -4.43% (p=0.002 n=6)
RandomPfx4Size/200_000     79.15 ± 0%    80.02 ± 0%    +1.10% (p=0.002 n=6)
RandomPfx6Size/1_000      100.70 ± 2%    84.06 ± 2%   -16.52% (p=0.002 n=6)
RandomPfx6Size/10_000      76.50 ± 0%    80.35 ± 0%    +5.03% (p=0.002 n=6)
RandomPfx6Size/100_000    101.10 ± 0%    80.03 ± 0%   -20.84% (p=0.002 n=6)
RandomPfx6Size/200_000     97.80 ± 0%    80.02 ± 0%   -18.18% (p=0.002 n=6)
RandomPfxSize/1_000       102.20 ± 2%    83.89 ± 2%   -17.92% (p=0.002 n=6)
RandomPfxSize/10_000       72.04 ± 0%    80.39 ± 0%   +11.59% (p=0.002 n=6)
RandomPfxSize/100_000      81.73 ± 0%    80.03 ± 0%    -2.08% (p=0.002 n=6)
RandomPfxSize/200_000      82.84 ± 0%    80.02 ± 0%    -3.40% (p=0.002 n=6)
geomean                    79.62         81.14         +1.91%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          critbitgo/size.bm          │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         101.0 ± 2%    119.1 ± 2%   +17.92% (p=0.002 n=6)
Tier1PfxSize/10_000        81.27 ± 0%   114.70 ± 0%   +41.13% (p=0.002 n=6)
Tier1PfxSize/100_000       47.41 ± 0%   114.40 ± 0%  +141.30% (p=0.002 n=6)
Tier1PfxSize/200_000       37.49 ± 0%   114.40 ± 0%  +205.15% (p=0.002 n=6)
RandomPfx4Size/1_000       100.3 ± 2%    115.9 ± 2%   +15.55% (p=0.002 n=6)
RandomPfx4Size/10_000      70.82 ± 0%   112.30 ± 0%   +58.57% (p=0.002 n=6)
RandomPfx4Size/100_000     83.74 ± 0%   112.00 ± 0%   +33.75% (p=0.002 n=6)
RandomPfx4Size/200_000     79.15 ± 0%   112.00 ± 0%   +41.50% (p=0.002 n=6)
RandomPfx6Size/1_000       100.7 ± 2%    132.0 ± 1%   +31.08% (p=0.002 n=6)
RandomPfx6Size/10_000      76.50 ± 0%   128.30 ± 0%   +67.71% (p=0.002 n=6)
RandomPfx6Size/100_000     101.1 ± 0%    128.0 ± 0%   +26.61% (p=0.002 n=6)
RandomPfx6Size/200_000     97.80 ± 0%   128.00 ± 0%   +30.88% (p=0.002 n=6)
RandomPfxSize/1_000        102.2 ± 2%    119.0 ± 1%   +16.44% (p=0.002 n=6)
RandomPfxSize/10_000       72.04 ± 0%   115.50 ± 0%   +60.33% (p=0.002 n=6)
RandomPfxSize/100_000      81.73 ± 0%   115.20 ± 0%   +40.95% (p=0.002 n=6)
RandomPfxSize/200_000      82.84 ± 0%   115.20 ± 0%   +39.06% (p=0.002 n=6)
geomean                    79.62         118.3        +48.62%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │           lpmtrie/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         101.0 ± 2%    214.9 ± 5%  +112.77% (p=0.002 n=6)
Tier1PfxSize/10_000        81.27 ± 0%   210.50 ± 5%  +159.01% (p=0.002 n=6)
Tier1PfxSize/100_000       47.41 ± 0%   209.90 ± 5%  +342.73% (p=0.002 n=6)
Tier1PfxSize/200_000       37.49 ± 0%   209.20 ± 5%  +458.02% (p=0.002 n=6)
RandomPfx4Size/1_000       100.3 ± 2%    211.8 ± 5%  +111.17% (p=0.002 n=6)
RandomPfx4Size/10_000      70.82 ± 0%   206.60 ± 4%  +191.73% (p=0.002 n=6)
RandomPfx4Size/100_000     83.74 ± 0%   199.20 ± 6%  +137.88% (p=0.002 n=6)
RandomPfx4Size/200_000     79.15 ± 0%   194.70 ± 7%  +145.99% (p=0.002 n=6)
RandomPfx6Size/1_000       100.7 ± 2%    227.6 ± 8%  +126.02% (p=0.002 n=6)
RandomPfx6Size/10_000      76.50 ± 0%   222.60 ± 7%  +190.98% (p=0.002 n=6)
RandomPfx6Size/100_000     101.1 ± 0%    219.1 ± 8%  +116.72% (p=0.002 n=6)
RandomPfx6Size/200_000     97.80 ± 0%   217.80 ± 8%  +122.70% (p=0.002 n=6)
RandomPfxSize/1_000        102.2 ± 2%    214.4 ± 5%  +109.78% (p=0.002 n=6)
RandomPfxSize/10_000       72.04 ± 0%   210.10 ± 5%  +191.64% (p=0.002 n=6)
RandomPfxSize/100_000      81.73 ± 0%   203.50 ± 7%  +148.99% (p=0.002 n=6)
RandomPfxSize/200_000      82.84 ± 0%   199.40 ± 7%  +140.70% (p=0.002 n=6)
geomean                    79.62         210.5       +164.42%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidranger/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                 │
Tier1PfxSize/1_000         101.0 ± 2%    539.2 ± 3%   +433.86% (p=0.002 n=6)
Tier1PfxSize/10_000        81.27 ± 0%   533.80 ± 3%   +556.82% (p=0.002 n=6)
Tier1PfxSize/100_000       47.41 ± 0%   527.20 ± 2%  +1012.00% (p=0.002 n=6)
Tier1PfxSize/200_000       37.49 ± 0%   522.20 ± 2%  +1292.90% (p=0.002 n=6)
RandomPfx4Size/1_000       100.3 ± 2%    526.7 ± 3%   +425.12% (p=0.002 n=6)
RandomPfx4Size/10_000      70.82 ± 0%   514.20 ± 3%   +626.07% (p=0.002 n=6)
RandomPfx4Size/100_000     83.74 ± 0%   479.80 ± 3%   +472.96% (p=0.002 n=6)
RandomPfx4Size/200_000     79.15 ± 0%   463.50 ± 3%   +485.60% (p=0.002 n=6)
RandomPfx6Size/1_000       100.7 ± 2%    594.8 ± 0%   +490.67% (p=0.002 n=6)
RandomPfx6Size/10_000      76.50 ± 0%   585.80 ± 0%   +665.75% (p=0.002 n=6)
RandomPfx6Size/100_000     101.1 ± 0%    574.7 ± 0%   +468.45% (p=0.002 n=6)
RandomPfx6Size/200_000     97.80 ± 0%   570.30 ± 0%   +483.13% (p=0.002 n=6)
RandomPfxSize/1_000        102.2 ± 2%    536.7 ± 2%   +425.15% (p=0.002 n=6)
RandomPfxSize/10_000       72.04 ± 0%   527.00 ± 2%   +631.54% (p=0.002 n=6)
RandomPfxSize/100_000      81.73 ± 0%   498.80 ± 2%   +510.30% (p=0.002 n=6)
RandomPfxSize/200_000      82.84 ± 0%   484.60 ± 2%   +484.98% (p=0.002 n=6)
geomean                    79.62         528.7        +564.03%
```

## lpm (longest-prefix-match)

For longest-prefix-match, `bart` and `art` are the champions.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │              art/lpm.bm               │
                                     │    sec/op     │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4            36.81n ±  17%   41.16n ± 19%   +11.80% (p=0.043 n=10)
LpmTier1Pfxs/RandomMatchIP6            40.64n ±  40%   43.26n ± 14%         ~ (p=0.239 n=10)
LpmTier1Pfxs/RandomMissIP4             17.43n ± 218%   25.48n ± 12%         ~ (p=0.469 n=10)
LpmTier1Pfxs/RandomMissIP6             10.81n ±   5%   25.52n ±  1%  +136.08% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     32.67n ±  11%   39.48n ±  2%   +20.83% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     23.69n ±  43%   41.02n ±  1%   +73.15% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      21.75n ±  48%   25.48n ±  2%         ~ (p=0.085 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      19.62n ±  42%   25.48n ± 62%   +29.89% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    36.72n ±  45%   40.72n ±  2%         ~ (p=0.470 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    20.07n ± 119%   42.29n ±  3%  +110.69% (p=0.014 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     42.24n ±   7%   25.49n ±  4%   -39.64% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     31.03n ±  32%   25.57n ±  1%   -17.59% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   25.47n ±  66%   49.71n ± 14%   +95.21% (p=0.002 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   21.87n ±  87%   47.04n ± 21%  +115.14% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    55.63n ±  22%   26.07n ± 10%   -53.12% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    46.18n ±  19%   25.48n ±  3%   -44.82% (p=0.000 n=10)
geomean                                27.82n          33.15n         +19.15%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │              cidrtree/lpm.bm               │
                                     │    sec/op     │      sec/op       vs base                  │
LpmTier1Pfxs/RandomMatchIP4            36.81n ±  17%    828.80n ±   95%  +2151.56% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            40.64n ±  40%   1020.00n ±   26%  +2409.84% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             17.43n ± 218%   1119.00n ±   40%  +6318.12% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             10.81n ±   5%     41.40n ± 1248%   +283.02% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     32.67n ±  11%    402.85n ±   25%  +1132.90% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     23.69n ±  43%    388.45n ±   46%  +1539.72% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      21.75n ±  48%    546.45n ±   47%  +2411.84% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      19.62n ±  42%    580.35n ±   23%  +2857.95% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    36.72n ±  45%    585.50n ±   40%  +1494.72% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    20.07n ± 119%   1045.30n ±   36%  +5108.27% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     42.24n ±   7%    652.70n ±   41%  +1445.22% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     31.03n ±  32%   1021.00n ±   39%  +3189.83% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   25.47n ±  66%    872.85n ±   32%  +3327.65% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   21.87n ±  87%   1665.50n ±   40%  +7517.20% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    55.63n ±  22%   1094.50n ±   17%  +1867.64% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    46.18n ±  19%   1258.00n ±   28%  +2623.83% (p=0.000 n=10)
geomean                                27.82n            668.4n          +2302.60%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │            critbitgo/lpm.bm             │
                                     │    sec/op     │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            36.81n ±  17%   129.70n ± 32%   +252.35% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            40.64n ±  40%   167.20n ± 31%   +311.42% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             17.43n ± 218%   621.45n ± 31%  +3464.38% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP6             10.81n ±   5%   366.95n ± 27%  +3294.54% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     32.67n ±  11%    93.72n ± 11%   +186.82% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     23.69n ±  43%   117.65n ±  8%   +396.62% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      21.75n ±  48%   170.60n ± 14%   +684.19% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      19.62n ±  42%   163.10n ± 18%   +731.29% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    36.72n ±  45%   102.20n ± 17%   +178.36% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    20.07n ± 119%   136.80n ± 14%   +581.61% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     42.24n ±   7%   238.95n ± 43%   +465.70% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     31.03n ±  32%   239.45n ± 19%   +671.55% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   25.47n ±  66%   154.00n ± 14%   +504.75% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   21.87n ±  87%   175.85n ± 55%   +704.25% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    55.63n ±  22%   482.90n ± 30%   +768.13% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    46.18n ±  19%   432.40n ± 23%   +836.23% (p=0.000 n=10)
geomean                                27.82n           200.6n         +620.99%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │             lpmtrie/lpm.bm              │
                                     │    sec/op     │     sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            36.81n ±  17%   222.35n ±  21%  +504.05% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            40.64n ±  40%   205.15n ±  66%  +404.80% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             17.43n ± 218%    68.46n ± 219%  +292.66% (p=0.022 n=10)
LpmTier1Pfxs/RandomMissIP6             10.81n ±   5%    13.97n ±  10%   +29.28% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     32.67n ±  11%   139.00n ±  10%  +325.40% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     23.69n ±  43%    85.16n ±  14%  +259.48% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      21.75n ±  48%   116.70n ±  25%  +436.43% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      19.62n ±  42%    93.40n ±  15%  +376.04% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    36.72n ±  45%   174.25n ±  10%  +374.60% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    20.07n ± 119%   125.30n ±   4%  +524.31% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     42.24n ±   7%   170.55n ±  23%  +303.76% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     31.03n ±  32%   158.60n ±  19%  +411.04% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   25.47n ±  66%   178.30n ±  13%  +600.18% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   21.87n ±  87%   159.25n ±   8%  +628.33% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    55.63n ±  22%   189.15n ±  14%  +240.04% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    46.18n ±  19%   182.40n ±  18%  +294.93% (p=0.000 n=10)
geomean                                27.82n           124.3n         +346.89%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │            cidranger/lpm.bm            │
                                     │    sec/op     │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            36.81n ±  17%   177.95n ± 43%  +383.43% (p=0.000 n=10)
LpmTier1Pfxs/RandomMatchIP6            40.64n ±  40%   216.30n ± 48%  +432.23% (p=0.000 n=10)
LpmTier1Pfxs/RandomMissIP4             17.43n ± 218%   136.20n ± 66%  +681.19% (p=0.001 n=10)
LpmTier1Pfxs/RandomMissIP6             10.81n ±   5%    64.73n ± 35%  +498.84% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP4     32.67n ±  11%   129.75n ± 13%  +297.09% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMatchIP6     23.69n ±  43%   136.60n ± 15%  +476.61% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP4      21.75n ±  48%   116.05n ± 13%  +433.44% (p=0.000 n=10)
LpmRandomPfxs/1_000/RandomMissIP6      19.62n ±  42%   140.80n ± 12%  +617.64% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP4    36.72n ±  45%   127.75n ± 45%  +247.95% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMatchIP6    20.07n ± 119%   168.15n ± 10%  +737.82% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP4     42.24n ±   7%   152.25n ± 10%  +260.44% (p=0.000 n=10)
LpmRandomPfxs/10_000/RandomMissIP6     31.03n ±  32%   182.55n ± 17%  +488.21% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP4   25.47n ±  66%   140.30n ± 43%  +450.95% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMatchIP6   21.87n ±  87%   159.70n ± 42%  +630.39% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP4    55.63n ±  22%   209.90n ± 12%  +277.35% (p=0.000 n=10)
LpmRandomPfxs/100_000/RandomMissIP6    46.18n ±  19%   223.60n ± 14%  +384.14% (p=0.000 n=10)
geomean                                27.82n           149.4n        +437.08%
```

## update, insert/delete

`bart` is by far the fastest algorithm for inserts and one of the fastest for delete.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │              art/update.bm              │
                         │   sec/route    │   sec/route     vs base                 │
InsertRandomPfxs/1_000        42.73n ± 2%   282.90n ±   1%   +562.14% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.75n ± 2%   583.05n ±   1%   +964.93% (p=0.002 n=6)
InsertRandomPfxs/100_000      147.8n ± 4%   1724.0n ±  39%  +1066.84% (p=0.002 n=6)
InsertRandomPfxs/200_000      295.4n ± 7%   1835.0n ± 196%   +521.19% (p=0.002 n=6)
DeleteRandomPfxs/1_000        18.85n ± 4%    17.63n ±   0%     -6.45% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.10n ± 2%    16.95n ±   2%     -6.33% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.05n ± 1%    17.70n ±  10%          ~ (p=0.065 n=6)
DeleteRandomPfxs/200_000      19.77n ± 1%    19.71n ±   4%          ~ (p=0.686 n=6)
geomean                       43.33n         123.6n          +185.20%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           cidrtree/update.bm            │
                         │   sec/route    │   sec/route    vs base                  │
InsertRandomPfxs/1_000        42.73n ± 2%   1311.50n ± 0%   +2969.63% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.75n ± 2%   1962.00n ± 2%   +3483.56% (p=0.002 n=6)
InsertRandomPfxs/100_000      147.8n ± 4%    3427.0n ± 3%   +2219.46% (p=0.002 n=6)
InsertRandomPfxs/200_000      295.4n ± 7%    4255.5n ± 3%   +1340.59% (p=0.002 n=6)
DeleteRandomPfxs/1_000        18.85n ± 4%     14.81n ± 1%     -21.46% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.10n ± 2%     15.17n ± 0%     -16.16% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.05n ± 1%     23.64n ± 4%     +31.00% (p=0.002 n=6)
DeleteRandomPfxs/200_000      19.77n ± 1%   5140.50n ± 2%  +25901.52% (p=0.002 n=6)
geomean                       43.33n          423.0n         +876.14%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │         critbitgo/update.bm          │
                         │   sec/route    │  sec/route    vs base                │
InsertRandomPfxs/1_000        42.73n ± 2%   176.05n ± 2%  +312.05% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.75n ± 2%   248.90n ± 1%  +354.61% (p=0.002 n=6)
InsertRandomPfxs/100_000      147.8n ± 4%    626.9n ± 2%  +324.33% (p=0.002 n=6)
InsertRandomPfxs/200_000      295.4n ± 7%    802.5n ± 2%  +171.67% (p=0.002 n=6)
DeleteRandomPfxs/1_000        18.85n ± 4%    71.44n ± 2%  +278.97% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.10n ± 2%    74.09n ± 2%  +309.48% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.05n ± 1%    81.85n ± 2%  +353.49% (p=0.002 n=6)
DeleteRandomPfxs/200_000      19.77n ± 1%    86.20n ± 3%  +335.99% (p=0.002 n=6)
geomean                       43.33n         173.6n       +300.54%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           lpmtrie/update.bm            │
                         │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000        42.73n ± 2%   400.90n ±  5%   +838.33% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.75n ± 2%   500.10n ±  5%   +813.42% (p=0.002 n=6)
InsertRandomPfxs/100_000      147.8n ± 4%   1286.0n ±  3%   +770.39% (p=0.002 n=6)
InsertRandomPfxs/200_000      295.4n ± 7%   1381.5n ± 20%   +367.67% (p=0.002 n=6)
DeleteRandomPfxs/1_000        18.85n ± 4%    76.09n ± 47%   +303.66% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.10n ± 2%   145.40n ±  2%   +703.54% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.05n ± 1%   330.10n ±  5%  +1728.81% (p=0.002 n=6)
DeleteRandomPfxs/200_000      19.77n ± 1%   479.70n ±  6%  +2326.40% (p=0.002 n=6)
geomean                       43.33n         397.6n         +817.50%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           cidranger/update.bm            │
                         │   sec/route    │   sec/route     vs base                  │
InsertRandomPfxs/1_000        42.73n ± 2%   4604.50n ±  1%  +10677.06% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.75n ± 2%   7114.00n ±  7%  +12893.61% (p=0.002 n=6)
InsertRandomPfxs/100_000      147.8n ± 4%   11771.5n ±  6%   +7867.17% (p=0.002 n=6)
InsertRandomPfxs/200_000      295.4n ± 7%   13284.0n ±  4%   +4396.95% (p=0.002 n=6)
DeleteRandomPfxs/1_000        18.85n ± 4%     93.09n ±  3%    +393.87% (p=0.002 n=6)
DeleteRandomPfxs/10_000       18.10n ± 2%     93.42n ±  0%    +416.25% (p=0.002 n=6)
DeleteRandomPfxs/100_000      18.05n ± 1%    147.90n ±  2%    +719.39% (p=0.002 n=6)
DeleteRandomPfxs/200_000      19.77n ± 1%   1699.00n ± 51%   +8493.83% (p=0.002 n=6)
geomean                       43.33n          1.352µ         +3021.23%
```
