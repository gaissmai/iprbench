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
LpmTier1Pfxs/RandomMatchIP4            26.56n ±   6%    42.69n ± 12%   +60.71% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            25.54n ±  25%    44.90n ±  4%   +75.82% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             27.61n ±  51%    25.99n ±  7%         ~ (p=0.507 n=20)
LpmTier1Pfxs/RandomMissIP6             6.525n ±   2%   26.725n ±  3%  +309.58% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     21.46n ±  66%    41.33n ±  3%   +92.59% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     16.68n ±   5%    44.02n ±  3%  +163.94% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      23.58n ±   9%    26.35n ±  7%   +11.75% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      17.55n ±  14%    25.96n ±  5%   +47.91% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    24.35n ±  55%    40.52n ±  2%   +66.44% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    11.34n ±  96%    42.74n ±  1%  +277.11% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     27.98n ±   4%    26.66n ±  7%         ~ (p=0.101 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     23.63n ±   6%    26.35n ±  5%   +11.49% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   10.98n ± 107%    50.00n ±  9%  +355.37% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   24.97n ±  57%    44.02n ±  6%   +76.29% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    38.63n ±   8%    26.26n ±  8%   -32.02% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    27.89n ±  25%    25.47n ± 12%    -8.69% (p=0.031 n=20)
geomean                                20.50n           33.85n         +65.08%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │              cidrtree/lpm.bm              │
                                     │    sec/op     │     sec/op      vs base                   │
LpmTier1Pfxs/RandomMatchIP4            26.56n ±   6%    961.70n ± 15%   +3520.86% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            25.54n ±  25%    961.70n ± 18%   +3666.20% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             27.61n ±  51%   1027.50n ± 21%   +3621.48% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             6.525n ±   2%    39.015n ±  0%    +497.93% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     21.46n ±  66%    326.15n ± 11%   +1419.80% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     16.68n ±   5%    868.55n ± 15%   +5107.13% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      23.58n ±   9%    491.15n ± 16%   +1982.91% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      17.55n ±  14%    549.40n ± 32%   +3029.59% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    24.35n ±  55%   1081.50n ± 27%   +4342.39% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    11.34n ±  96%   1428.00n ± 13%  +12498.15% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     27.98n ±   4%    647.35n ± 13%   +2213.20% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     23.63n ±   6%    908.80n ± 12%   +3745.14% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   10.98n ± 107%    948.35n ± 29%   +8537.07% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   24.97n ±  57%   1451.50n ± 20%   +5712.98% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    38.63n ±   8%    980.50n ± 23%   +2438.18% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    27.89n ±  25%   1160.50n ± 14%   +4060.24% (p=0.000 n=20)
geomean                                20.50n            708.5n         +3355.45%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │             critbitgo/lpm.bm             │
                                     │    sec/op     │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            26.56n ±   6%    116.35n ± 22%   +338.06% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            25.54n ±  25%    170.65n ± 24%   +568.30% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             27.61n ±  51%    649.85n ± 18%  +2253.68% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             6.525n ±   2%   361.250n ± 32%  +5436.40% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     21.46n ±  66%     96.07n ± 10%   +347.67% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     16.68n ±   5%    103.50n ±  5%   +520.50% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      23.58n ±   9%    154.70n ± 14%   +556.06% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      17.55n ±  14%    152.80n ±  6%   +770.41% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    24.35n ±  55%    106.05n ± 11%   +335.61% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    11.34n ±  96%    159.25n ± 17%  +1304.94% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     27.98n ±   4%    252.75n ± 16%   +803.16% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     23.63n ±   6%    201.90n ± 13%   +754.24% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   10.98n ± 107%    158.50n ± 36%  +1343.53% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   24.97n ±  57%    170.10n ± 10%   +581.22% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    38.63n ±   8%    523.70n ± 17%  +1255.68% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    27.89n ±  25%    379.70n ± 17%  +1261.18% (p=0.000 n=20)
geomean                                20.50n            196.8n         +859.78%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │             lpmtrie/lpm.bm              │
                                     │    sec/op     │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            26.56n ±   6%   197.85n ±  8%   +644.92% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            25.54n ±  25%   150.20n ± 34%   +488.21% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             27.61n ±  51%   101.71n ± 97%   +268.38% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             6.525n ±   2%   13.245n ±  8%   +102.99% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     21.46n ±  66%   132.70n ±  2%   +518.36% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     16.68n ±   5%   126.75n ± 18%   +659.89% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      23.58n ±   9%   103.95n ±  2%   +340.84% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      17.55n ±  14%    91.00n ±  8%   +418.37% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    24.35n ±  55%   155.25n ±  8%   +537.71% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    11.34n ±  96%   122.00n ±  9%   +976.31% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     27.98n ±   4%   143.50n ±  6%   +412.77% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     23.63n ±   6%   128.50n ±  8%   +443.69% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   10.98n ± 107%   179.25n ±  5%  +1532.51% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   24.97n ±  57%   155.30n ± 12%   +521.95% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    38.63n ±   8%   196.80n ±  5%   +409.45% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    27.89n ±  25%   178.45n ± 16%   +539.72% (p=0.000 n=20)
geomean                                20.50n           121.2n         +491.21%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │  bart/lpm.bm  │            cidranger/lpm.bm             │
                                     │    sec/op     │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            26.56n ±   6%   176.45n ± 10%   +564.34% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            25.54n ±  25%   226.05n ± 36%   +785.26% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             27.61n ±  51%   166.00n ± 70%   +501.23% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             6.525n ±   2%   60.160n ±  3%   +821.99% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     21.46n ±  66%   125.35n ±  7%   +484.11% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     16.68n ±   5%   148.80n ±  4%   +792.09% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      23.58n ±   9%   118.95n ± 10%   +404.45% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      17.55n ±  14%   142.00n ±  8%   +708.89% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    24.35n ±  55%   123.05n ± 33%   +405.44% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    11.34n ±  96%   144.70n ±  6%  +1176.58% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     27.98n ±   4%   153.75n ± 11%   +449.40% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     23.63n ±   6%   174.75n ± 10%   +639.37% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   10.98n ± 107%   120.65n ± 19%   +998.82% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   24.97n ±  57%   154.75n ± 17%   +519.74% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    38.63n ±   8%   210.35n ±  6%   +444.52% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    27.89n ±  25%   235.40n ±  7%   +743.88% (p=0.000 n=20)
geomean                                20.50n           148.4n         +623.89%
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
Tier1PfxSize/1_000         107.0 ± 2%    7590.0 ± 0%  +6993.46% (p=0.002 n=6)
Tier1PfxSize/10_000        79.54 ± 0%   4889.00 ± 0%  +6046.59% (p=0.002 n=6)
Tier1PfxSize/100_000       35.73 ± 0%   1669.00 ± 0%  +4571.14% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%   1098.00 ± 0%  +4448.47% (p=0.002 n=6)
RandomPfx4Size/1_000       103.3 ± 2%    7563.0 ± 0%  +7221.39% (p=0.002 n=6)
RandomPfx4Size/10_000      76.16 ± 0%   6068.00 ± 0%  +7867.44% (p=0.002 n=6)
RandomPfx4Size/100_000     81.77 ± 0%   5489.00 ± 0%  +6612.73% (p=0.002 n=6)
RandomPfx4Size/200_000     72.84 ± 0%   4841.00 ± 0%  +6546.07% (p=0.002 n=6)
RandomPfx6Size/1_000       105.0 ± 2%    7812.0 ± 0%  +7340.00% (p=0.002 n=6)
RandomPfx6Size/10_000      84.77 ± 0%   6838.00 ± 0%  +7966.53% (p=0.002 n=6)
RandomPfx6Size/100_000     105.8 ± 0%    7832.0 ± 0%  +7302.65% (p=0.002 n=6)
RandomPfx6Size/200_000     100.9 ± 0%    7605.0 ± 0%  +7437.17% (p=0.002 n=6)
RandomPfxSize/1_000        106.0 ± 2%    7524.0 ± 0%  +6998.11% (p=0.002 n=6)
RandomPfxSize/10_000       75.21 ± 0%   6210.00 ± 0%  +8156.88% (p=0.002 n=6)
RandomPfxSize/100_000      82.42 ± 0%   5807.00 ± 0%  +6945.62% (p=0.002 n=6)
RandomPfxSize/200_000      78.50 ± 0%   5459.00 ± 0%  +6854.14% (p=0.002 n=6)
geomean                    77.48        5.178Ki       +6744.17%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidrtree/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000        107.00 ± 2%    69.06 ± 3%   -35.46% (p=0.002 n=6)
Tier1PfxSize/10_000        79.54 ± 0%    64.35 ± 0%   -19.10% (p=0.002 n=6)
Tier1PfxSize/100_000       35.73 ± 0%    64.03 ± 0%   +79.21% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%    64.02 ± 0%  +165.20% (p=0.002 n=6)
RandomPfx4Size/1_000      103.30 ± 2%    67.89 ± 3%   -34.28% (p=0.002 n=6)
RandomPfx4Size/10_000      76.16 ± 0%    64.35 ± 0%   -15.51% (p=0.002 n=6)
RandomPfx4Size/100_000     81.77 ± 0%    64.03 ± 0%   -21.69% (p=0.002 n=6)
RandomPfx4Size/200_000     72.84 ± 0%    64.02 ± 0%   -12.11% (p=0.002 n=6)
RandomPfx6Size/1_000      105.00 ± 2%    68.06 ± 2%   -35.18% (p=0.002 n=6)
RandomPfx6Size/10_000      84.77 ± 0%    64.35 ± 0%   -24.09% (p=0.002 n=6)
RandomPfx6Size/100_000    105.80 ± 0%    64.03 ± 0%   -39.48% (p=0.002 n=6)
RandomPfx6Size/200_000    100.90 ± 0%    64.02 ± 0%   -36.55% (p=0.002 n=6)
RandomPfxSize/1_000       106.00 ± 2%    67.89 ± 3%   -35.95% (p=0.002 n=6)
RandomPfxSize/10_000       75.21 ± 0%    64.37 ± 0%   -14.41% (p=0.002 n=6)
RandomPfxSize/100_000      82.42 ± 0%    64.03 ± 0%   -22.31% (p=0.002 n=6)
RandomPfxSize/200_000      78.50 ± 0%    64.02 ± 0%   -18.45% (p=0.002 n=6)
geomean                    77.48         65.13        -15.93%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          critbitgo/size.bm          │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         107.0 ± 2%    119.1 ± 2%   +11.31% (p=0.002 n=6)
Tier1PfxSize/10_000        79.54 ± 0%   114.70 ± 0%   +44.20% (p=0.002 n=6)
Tier1PfxSize/100_000       35.73 ± 0%   114.40 ± 0%  +220.18% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%   114.40 ± 0%  +373.90% (p=0.002 n=6)
RandomPfx4Size/1_000       103.3 ± 2%    115.9 ± 2%   +12.20% (p=0.002 n=6)
RandomPfx4Size/10_000      76.16 ± 0%   112.30 ± 0%   +47.45% (p=0.002 n=6)
RandomPfx4Size/100_000     81.77 ± 0%   112.00 ± 0%   +36.97% (p=0.002 n=6)
RandomPfx4Size/200_000     72.84 ± 0%   112.00 ± 0%   +53.76% (p=0.002 n=6)
RandomPfx6Size/1_000       105.0 ± 2%    132.0 ± 1%   +25.71% (p=0.002 n=6)
RandomPfx6Size/10_000      84.77 ± 0%   128.30 ± 0%   +51.35% (p=0.002 n=6)
RandomPfx6Size/100_000     105.8 ± 0%    128.0 ± 0%   +20.98% (p=0.002 n=6)
RandomPfx6Size/200_000     100.9 ± 0%    128.0 ± 0%   +26.86% (p=0.002 n=6)
RandomPfxSize/1_000        106.0 ± 2%    118.8 ± 1%   +12.08% (p=0.002 n=6)
RandomPfxSize/10_000       75.21 ± 0%   115.50 ± 0%   +53.57% (p=0.002 n=6)
RandomPfxSize/100_000      82.42 ± 0%   115.20 ± 0%   +39.77% (p=0.002 n=6)
RandomPfxSize/200_000      78.50 ± 0%   115.20 ± 0%   +46.75% (p=0.002 n=6)
geomean                    77.48         118.3        +52.71%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │           lpmtrie/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         107.0 ± 2%    214.9 ± 5%  +100.84% (p=0.002 n=6)
Tier1PfxSize/10_000        79.54 ± 0%   210.50 ± 5%  +164.65% (p=0.002 n=6)
Tier1PfxSize/100_000       35.73 ± 0%   209.90 ± 5%  +487.46% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%   209.20 ± 5%  +766.61% (p=0.002 n=6)
RandomPfx4Size/1_000       103.3 ± 2%    211.3 ± 5%  +104.55% (p=0.002 n=6)
RandomPfx4Size/10_000      76.16 ± 0%   206.80 ± 5%  +171.53% (p=0.002 n=6)
RandomPfx4Size/100_000     81.77 ± 0%   199.50 ± 7%  +143.98% (p=0.002 n=6)
RandomPfx4Size/200_000     72.84 ± 0%   194.60 ± 7%  +167.16% (p=0.002 n=6)
RandomPfx6Size/1_000       105.0 ± 2%    227.6 ± 8%  +116.76% (p=0.002 n=6)
RandomPfx6Size/10_000      84.77 ± 0%   223.10 ± 7%  +163.18% (p=0.002 n=6)
RandomPfx6Size/100_000     105.8 ± 0%    219.3 ± 8%  +107.28% (p=0.002 n=6)
RandomPfx6Size/200_000     100.9 ± 0%    217.8 ± 8%  +115.86% (p=0.002 n=6)
RandomPfxSize/1_000        106.0 ± 2%    214.4 ± 5%  +102.26% (p=0.002 n=6)
RandomPfxSize/10_000       75.21 ± 0%   209.70 ± 5%  +178.82% (p=0.002 n=6)
RandomPfxSize/100_000      82.42 ± 0%   203.20 ± 7%  +146.54% (p=0.002 n=6)
RandomPfxSize/200_000      78.50 ± 0%   199.20 ± 7%  +153.76% (p=0.002 n=6)
geomean                    77.48         210.5       +171.71%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidranger/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                 │
Tier1PfxSize/1_000         107.0 ± 2%    539.2 ± 3%   +403.93% (p=0.002 n=6)
Tier1PfxSize/10_000        79.54 ± 0%   533.80 ± 3%   +571.11% (p=0.002 n=6)
Tier1PfxSize/100_000       35.73 ± 0%   527.20 ± 2%  +1375.51% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%   522.20 ± 2%  +2063.21% (p=0.002 n=6)
RandomPfx4Size/1_000       103.3 ± 2%    526.6 ± 3%   +409.78% (p=0.002 n=6)
RandomPfx4Size/10_000      76.16 ± 0%   514.00 ± 3%   +574.89% (p=0.002 n=6)
RandomPfx4Size/100_000     81.77 ± 0%   479.90 ± 3%   +486.89% (p=0.002 n=6)
RandomPfx4Size/200_000     72.84 ± 0%   463.50 ± 3%   +536.33% (p=0.002 n=6)
RandomPfx6Size/1_000       105.0 ± 2%    594.5 ± 0%   +466.19% (p=0.002 n=6)
RandomPfx6Size/10_000      84.77 ± 0%   585.50 ± 0%   +590.69% (p=0.002 n=6)
RandomPfx6Size/100_000     105.8 ± 0%    574.6 ± 0%   +443.10% (p=0.002 n=6)
RandomPfx6Size/200_000     100.9 ± 0%    570.4 ± 0%   +465.31% (p=0.002 n=6)
RandomPfxSize/1_000        106.0 ± 2%    537.3 ± 2%   +406.89% (p=0.002 n=6)
RandomPfxSize/10_000       75.21 ± 0%   527.20 ± 2%   +600.97% (p=0.002 n=6)
RandomPfxSize/100_000      82.42 ± 0%   499.20 ± 2%   +505.68% (p=0.002 n=6)
RandomPfxSize/200_000      78.50 ± 0%   484.50 ± 2%   +517.20% (p=0.002 n=6)
geomean                    77.48         528.7        +582.41%
```

## update, insert/delete

`bart` is by far the fastest algorithm for inserts and one of the fastest for delete.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │              art/update.bm              │
                         │   sec/route    │   sec/route     vs base                 │
InsertRandomPfxs/1_000        42.59n ± 1%   278.35n ±   0%   +553.63% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.10n ± 1%   548.30n ±   1%   +913.59% (p=0.002 n=6)
InsertRandomPfxs/100_000      137.0n ± 1%   1850.0n ±  29%  +1250.86% (p=0.002 n=6)
InsertRandomPfxs/200_000      232.1n ± 1%   1728.5n ± 193%   +644.72% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.79n ± 0%    17.49n ±   0%     -1.66% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.41n ± 2%    16.90n ±   1%     -2.93% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.67n ± 1%    17.41n ±   1%     -1.47% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.97n ± 1%    18.98n ±   1%          ~ (p=0.779 n=6)
geomean                       40.75n         121.6n          +198.31%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │            cidrtree/update.bm            │
                         │   sec/route    │    sec/route     vs base                 │
InsertRandomPfxs/1_000        42.59n ± 1%   1312.00n ±   2%  +2980.90% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.10n ± 1%   1943.00n ±   1%  +3491.83% (p=0.002 n=6)
InsertRandomPfxs/100_000      137.0n ± 1%    3175.5n ±   7%  +2218.73% (p=0.002 n=6)
InsertRandomPfxs/200_000      232.1n ± 1%    3871.5n ±   1%  +1568.03% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.79n ± 0%     14.81n ±   7%    -16.73% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.41n ± 2%     15.19n ±   1%    -12.72% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.67n ± 1%     23.79n ±   6%    +34.61% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.97n ± 1%    432.50n ± 135%  +2179.92% (p=0.002 n=6)
geomean                       40.75n          303.8n          +645.55%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │          critbitgo/update.bm          │
                         │   sec/route    │   sec/route    vs base                │
InsertRandomPfxs/1_000        42.59n ± 1%   172.60n ±  2%  +305.31% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.10n ± 1%   241.05n ±  1%  +345.60% (p=0.002 n=6)
InsertRandomPfxs/100_000      137.0n ± 1%    564.4n ±  1%  +312.08% (p=0.002 n=6)
InsertRandomPfxs/200_000      232.1n ± 1%    767.9n ±  3%  +230.85% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.79n ± 0%    70.89n ±  3%  +298.57% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.41n ± 2%    73.51n ±  1%  +322.23% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.67n ± 1%    78.01n ± 14%  +341.48% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.97n ± 1%    82.72n ±  2%  +336.03% (p=0.002 n=6)
geomean                       40.75n         167.0n        +309.95%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           lpmtrie/update.bm           │
                         │   sec/route    │  sec/route    vs base                 │
InsertRandomPfxs/1_000        42.59n ± 1%   385.15n ± 2%   +804.43% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.10n ± 1%   480.85n ± 6%   +788.90% (p=0.002 n=6)
InsertRandomPfxs/100_000      137.0n ± 1%   1148.0n ± 2%   +738.26% (p=0.002 n=6)
InsertRandomPfxs/200_000      232.1n ± 1%   1252.0n ± 4%   +439.42% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.79n ± 0%    85.49n ± 0%   +380.69% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.41n ± 2%   146.30n ± 0%   +740.32% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.67n ± 1%   320.20n ± 2%  +1712.11% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.97n ± 1%   465.30n ± 2%  +2352.82% (p=0.002 n=6)
geomean                       40.75n         386.3n        +848.06%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           cidranger/update.bm            │
                         │   sec/route    │   sec/route     vs base                  │
InsertRandomPfxs/1_000        42.59n ± 1%   4495.00n ±  1%  +10455.36% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.10n ± 1%   6452.00n ±  3%  +11827.17% (p=0.002 n=6)
InsertRandomPfxs/100_000      137.0n ± 1%   10631.5n ±  5%   +7663.05% (p=0.002 n=6)
InsertRandomPfxs/200_000      232.1n ± 1%   11937.0n ±  3%   +5043.04% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.79n ± 0%     95.47n ±  2%    +436.83% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.41n ± 2%     92.30n ±  5%    +430.18% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.67n ± 1%    142.85n ±  2%    +708.43% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.97n ± 1%    722.80n ± 20%   +3710.23% (p=0.002 n=6)
geomean                       40.75n          1.163µ         +2754.40%
```
