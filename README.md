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
                                     │ bart/lpm.bm  │               art/lpm.bm               │
                                     │    sec/op    │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            25.41n ±  9%    42.69n ± 12%   +67.99% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            25.96n ± 23%    44.90n ±  4%   +72.97% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             25.38n ± 11%    25.99n ±  7%         ~ (p=0.337 n=20)
LpmTier1Pfxs/RandomMissIP6             6.296n ±  0%   26.725n ±  3%  +324.44% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     21.71n ± 16%    41.33n ±  3%   +90.37% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     18.11n ±  4%    44.02n ±  3%  +143.10% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      20.91n ±  9%    26.35n ±  7%   +25.99% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      17.54n ± 64%    25.96n ±  5%   +48.08% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    23.16n ±  6%    40.52n ±  2%   +74.92% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    21.63n ± 65%    42.74n ±  1%   +97.66% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     26.47n ±  3%    26.66n ±  7%         ~ (p=0.615 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     21.47n ± 13%    26.35n ±  5%   +22.73% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   25.41n ± 59%    50.00n ±  9%   +96.77% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   11.07n ±  9%    44.02n ±  6%  +297.65% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    38.91n ±  4%    26.26n ±  8%   -32.50% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    29.81n ± 23%    25.47n ± 12%   -14.57% (p=0.030 n=20)
geomean                                20.98n          33.85n         +61.32%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │              cidrtree/lpm.bm              │
                                     │    sec/op    │     sec/op      vs base                   │
LpmTier1Pfxs/RandomMatchIP4            25.41n ±  9%    961.70n ± 15%   +3684.73% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            25.96n ± 23%    961.70n ± 18%   +3605.26% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             25.38n ± 11%   1027.50n ± 21%   +3948.46% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             6.296n ±  0%    39.015n ±  0%    +519.63% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     21.71n ± 16%    326.15n ± 11%   +1402.30% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     18.11n ±  4%    868.55n ± 15%   +4695.97% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      20.91n ±  9%    491.15n ± 16%   +2248.31% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      17.54n ± 64%    549.40n ± 32%   +3033.16% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    23.16n ±  6%   1081.50n ± 27%   +4568.68% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    21.63n ± 65%   1428.00n ± 13%   +6503.47% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     26.47n ±  3%    647.35n ± 13%   +2345.60% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     21.47n ± 13%    908.80n ± 12%   +4132.88% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   25.41n ± 59%    948.35n ± 29%   +3632.19% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   11.07n ±  9%   1451.50n ± 20%  +13012.01% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    38.91n ±  4%    980.50n ± 23%   +2420.24% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    29.81n ± 23%   1160.50n ± 14%   +3792.34% (p=0.000 n=20)
geomean                                20.98n           708.5n         +3276.72%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             critbitgo/lpm.bm             │
                                     │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            25.41n ±  9%    116.35n ± 22%   +357.89% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            25.96n ± 23%    170.65n ± 24%   +557.48% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             25.38n ± 11%    649.85n ± 18%  +2460.48% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             6.296n ±  0%   361.250n ± 32%  +5637.31% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     21.71n ± 16%     96.07n ± 10%   +342.51% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     18.11n ±  4%    103.50n ±  5%   +471.51% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      20.91n ±  9%    154.70n ± 14%   +639.66% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      17.54n ± 64%    152.80n ±  6%   +771.40% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    23.16n ±  6%    106.05n ± 11%   +357.80% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    21.63n ± 65%    159.25n ± 17%   +636.42% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     26.47n ±  3%    252.75n ± 16%   +854.85% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     21.47n ± 13%    201.90n ± 13%   +840.38% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   25.41n ± 59%    158.50n ± 36%   +523.77% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   11.07n ±  9%    170.10n ± 10%  +1436.59% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    38.91n ±  4%    523.70n ± 17%  +1246.10% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    29.81n ± 23%    379.70n ± 17%  +1173.52% (p=0.000 n=20)
geomean                                20.98n           196.8n         +837.91%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             lpmtrie/lpm.bm              │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            25.41n ±  9%   197.85n ±  8%   +678.63% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            25.96n ± 23%   150.20n ± 34%   +478.69% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             25.38n ± 11%   101.71n ± 97%   +300.75% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             6.296n ±  0%   13.245n ±  8%   +110.35% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     21.71n ± 16%   132.70n ±  2%   +511.24% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     18.11n ±  4%   126.75n ± 18%   +599.89% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      20.91n ±  9%   103.95n ±  2%   +397.01% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      17.54n ± 64%    91.00n ±  8%   +418.96% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    23.16n ±  6%   155.25n ±  8%   +570.19% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    21.63n ± 65%   122.00n ±  9%   +464.16% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     26.47n ±  3%   143.50n ±  6%   +442.12% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     21.47n ± 13%   128.50n ±  8%   +498.51% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   25.41n ± 59%   179.25n ±  5%   +605.43% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   11.07n ±  9%   155.30n ± 12%  +1302.89% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    38.91n ±  4%   196.80n ±  5%   +405.85% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    29.81n ± 23%   178.45n ± 16%   +498.52% (p=0.000 n=20)
geomean                                20.98n          121.2n         +477.74%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            cidranger/lpm.bm             │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            25.41n ±  9%   176.45n ± 10%   +594.41% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            25.96n ± 23%   226.05n ± 36%   +770.93% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             25.38n ± 11%   166.00n ± 70%   +554.06% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             6.296n ±  0%   60.160n ±  3%   +855.45% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     21.71n ± 16%   125.35n ±  7%   +477.38% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     18.11n ±  4%   148.80n ±  4%   +721.65% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      20.91n ±  9%   118.95n ± 10%   +468.73% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      17.54n ± 64%   142.00n ±  8%   +709.81% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    23.16n ±  6%   123.05n ± 33%   +431.19% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    21.63n ± 65%   144.70n ±  6%   +569.13% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     26.47n ±  3%   153.75n ± 11%   +480.85% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     21.47n ± 13%   174.75n ± 10%   +713.93% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   25.41n ± 59%   120.65n ± 19%   +374.81% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   11.07n ±  9%   154.75n ± 17%  +1297.92% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    38.91n ±  4%   210.35n ±  6%   +440.68% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    29.81n ± 23%   235.40n ±  7%   +689.54% (p=0.000 n=20)
geomean                                20.98n          148.4n         +607.39%
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
RandomPfx4Size/1_000       103.6 ± 2%    7563.0 ± 0%  +7200.19% (p=0.002 n=6)
RandomPfx4Size/10_000      76.14 ± 0%   6068.00 ± 0%  +7869.53% (p=0.002 n=6)
RandomPfx4Size/100_000     81.93 ± 0%   5489.00 ± 0%  +6599.62% (p=0.002 n=6)
RandomPfx4Size/200_000     73.04 ± 0%   4841.00 ± 0%  +6527.88% (p=0.002 n=6)
RandomPfx6Size/1_000       104.5 ± 2%    7812.0 ± 0%  +7375.60% (p=0.002 n=6)
RandomPfx6Size/10_000      84.89 ± 0%   6838.00 ± 0%  +7955.13% (p=0.002 n=6)
RandomPfx6Size/100_000     106.0 ± 0%    7832.0 ± 0%  +7288.68% (p=0.002 n=6)
RandomPfx6Size/200_000     100.9 ± 0%    7605.0 ± 0%  +7437.17% (p=0.002 n=6)
RandomPfxSize/1_000        101.4 ± 2%    7524.0 ± 0%  +7320.12% (p=0.002 n=6)
RandomPfxSize/10_000       75.09 ± 0%   6210.00 ± 0%  +8170.08% (p=0.002 n=6)
RandomPfxSize/100_000      82.39 ± 0%   5807.00 ± 0%  +6948.19% (p=0.002 n=6)
RandomPfxSize/200_000      78.49 ± 0%   5459.00 ± 0%  +6855.03% (p=0.002 n=6)
geomean                    77.28        5.178Ki       +6761.56%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidrtree/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000        107.00 ± 2%    69.06 ± 3%   -35.46% (p=0.002 n=6)
Tier1PfxSize/10_000        79.54 ± 0%    64.35 ± 0%   -19.10% (p=0.002 n=6)
Tier1PfxSize/100_000       35.73 ± 0%    64.03 ± 0%   +79.21% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%    64.02 ± 0%  +165.20% (p=0.002 n=6)
RandomPfx4Size/1_000      103.60 ± 2%    67.89 ± 3%   -34.47% (p=0.002 n=6)
RandomPfx4Size/10_000      76.14 ± 0%    64.35 ± 0%   -15.48% (p=0.002 n=6)
RandomPfx4Size/100_000     81.93 ± 0%    64.03 ± 0%   -21.85% (p=0.002 n=6)
RandomPfx4Size/200_000     73.04 ± 0%    64.02 ± 0%   -12.35% (p=0.002 n=6)
RandomPfx6Size/1_000      104.50 ± 2%    68.06 ± 2%   -34.87% (p=0.002 n=6)
RandomPfx6Size/10_000      84.89 ± 0%    64.35 ± 0%   -24.20% (p=0.002 n=6)
RandomPfx6Size/100_000    106.00 ± 0%    64.03 ± 0%   -39.59% (p=0.002 n=6)
RandomPfx6Size/200_000    100.90 ± 0%    64.02 ± 0%   -36.55% (p=0.002 n=6)
RandomPfxSize/1_000       101.40 ± 2%    67.89 ± 3%   -33.05% (p=0.002 n=6)
RandomPfxSize/10_000       75.09 ± 0%    64.37 ± 0%   -14.28% (p=0.002 n=6)
RandomPfxSize/100_000      82.39 ± 0%    64.03 ± 0%   -22.28% (p=0.002 n=6)
RandomPfxSize/200_000      78.49 ± 0%    64.02 ± 0%   -18.44% (p=0.002 n=6)
geomean                    77.28         65.13        -15.72%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          critbitgo/size.bm          │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         107.0 ± 2%    119.1 ± 2%   +11.31% (p=0.002 n=6)
Tier1PfxSize/10_000        79.54 ± 0%   114.70 ± 0%   +44.20% (p=0.002 n=6)
Tier1PfxSize/100_000       35.73 ± 0%   114.40 ± 0%  +220.18% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%   114.40 ± 0%  +373.90% (p=0.002 n=6)
RandomPfx4Size/1_000       103.6 ± 2%    115.9 ± 2%   +11.87% (p=0.002 n=6)
RandomPfx4Size/10_000      76.14 ± 0%   112.30 ± 0%   +47.49% (p=0.002 n=6)
RandomPfx4Size/100_000     81.93 ± 0%   112.00 ± 0%   +36.70% (p=0.002 n=6)
RandomPfx4Size/200_000     73.04 ± 0%   112.00 ± 0%   +53.34% (p=0.002 n=6)
RandomPfx6Size/1_000       104.5 ± 2%    132.0 ± 1%   +26.32% (p=0.002 n=6)
RandomPfx6Size/10_000      84.89 ± 0%   128.30 ± 0%   +51.14% (p=0.002 n=6)
RandomPfx6Size/100_000     106.0 ± 0%    128.0 ± 0%   +20.75% (p=0.002 n=6)
RandomPfx6Size/200_000     100.9 ± 0%    128.0 ± 0%   +26.86% (p=0.002 n=6)
RandomPfxSize/1_000        101.4 ± 2%    118.8 ± 1%   +17.16% (p=0.002 n=6)
RandomPfxSize/10_000       75.09 ± 0%   115.50 ± 0%   +53.82% (p=0.002 n=6)
RandomPfxSize/100_000      82.39 ± 0%   115.20 ± 0%   +39.82% (p=0.002 n=6)
RandomPfxSize/200_000      78.49 ± 0%   115.20 ± 0%   +46.77% (p=0.002 n=6)
geomean                    77.28         118.3        +53.10%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │           lpmtrie/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         107.0 ± 2%    214.9 ± 5%  +100.84% (p=0.002 n=6)
Tier1PfxSize/10_000        79.54 ± 0%   210.50 ± 5%  +164.65% (p=0.002 n=6)
Tier1PfxSize/100_000       35.73 ± 0%   209.90 ± 5%  +487.46% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%   209.20 ± 5%  +766.61% (p=0.002 n=6)
RandomPfx4Size/1_000       103.6 ± 2%    211.3 ± 5%  +103.96% (p=0.002 n=6)
RandomPfx4Size/10_000      76.14 ± 0%   206.80 ± 5%  +171.60% (p=0.002 n=6)
RandomPfx4Size/100_000     81.93 ± 0%   199.50 ± 7%  +143.50% (p=0.002 n=6)
RandomPfx4Size/200_000     73.04 ± 0%   194.60 ± 7%  +166.43% (p=0.002 n=6)
RandomPfx6Size/1_000       104.5 ± 2%    227.6 ± 8%  +117.80% (p=0.002 n=6)
RandomPfx6Size/10_000      84.89 ± 0%   223.10 ± 7%  +162.81% (p=0.002 n=6)
RandomPfx6Size/100_000     106.0 ± 0%    219.3 ± 8%  +106.89% (p=0.002 n=6)
RandomPfx6Size/200_000     100.9 ± 0%    217.8 ± 8%  +115.86% (p=0.002 n=6)
RandomPfxSize/1_000        101.4 ± 2%    214.4 ± 5%  +111.44% (p=0.002 n=6)
RandomPfxSize/10_000       75.09 ± 0%   209.70 ± 5%  +179.26% (p=0.002 n=6)
RandomPfxSize/100_000      82.39 ± 0%   203.20 ± 7%  +146.63% (p=0.002 n=6)
RandomPfxSize/200_000      78.49 ± 0%   199.20 ± 7%  +153.79% (p=0.002 n=6)
geomean                    77.28         210.5       +172.40%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidranger/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                 │
Tier1PfxSize/1_000         107.0 ± 2%    539.2 ± 3%   +403.93% (p=0.002 n=6)
Tier1PfxSize/10_000        79.54 ± 0%   533.80 ± 3%   +571.11% (p=0.002 n=6)
Tier1PfxSize/100_000       35.73 ± 0%   527.20 ± 2%  +1375.51% (p=0.002 n=6)
Tier1PfxSize/200_000       24.14 ± 0%   522.20 ± 2%  +2063.21% (p=0.002 n=6)
RandomPfx4Size/1_000       103.6 ± 2%    526.6 ± 3%   +408.30% (p=0.002 n=6)
RandomPfx4Size/10_000      76.14 ± 0%   514.00 ± 3%   +575.07% (p=0.002 n=6)
RandomPfx4Size/100_000     81.93 ± 0%   479.90 ± 3%   +485.74% (p=0.002 n=6)
RandomPfx4Size/200_000     73.04 ± 0%   463.50 ± 3%   +534.58% (p=0.002 n=6)
RandomPfx6Size/1_000       104.5 ± 2%    594.5 ± 0%   +468.90% (p=0.002 n=6)
RandomPfx6Size/10_000      84.89 ± 0%   585.50 ± 0%   +589.72% (p=0.002 n=6)
RandomPfx6Size/100_000     106.0 ± 0%    574.6 ± 0%   +442.08% (p=0.002 n=6)
RandomPfx6Size/200_000     100.9 ± 0%    570.4 ± 0%   +465.31% (p=0.002 n=6)
RandomPfxSize/1_000        101.4 ± 2%    537.3 ± 2%   +429.88% (p=0.002 n=6)
RandomPfxSize/10_000       75.09 ± 0%   527.20 ± 2%   +602.09% (p=0.002 n=6)
RandomPfxSize/100_000      82.39 ± 0%   499.20 ± 2%   +505.90% (p=0.002 n=6)
RandomPfxSize/200_000      78.49 ± 0%   484.50 ± 2%   +517.28% (p=0.002 n=6)
geomean                    77.28         528.7        +584.15%
```

## update, insert/delete

`bart` is by far the fastest algorithm for inserts and one of the fastest for delete.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │              art/update.bm              │
                         │   sec/route    │   sec/route     vs base                 │
InsertRandomPfxs/1_000        43.46n ± 2%   278.35n ±   0%   +540.47% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.53n ± 1%   548.30n ±   1%   +905.59% (p=0.002 n=6)
InsertRandomPfxs/100_000      125.1n ± 3%   1850.0n ±  29%  +1378.82% (p=0.002 n=6)
InsertRandomPfxs/200_000      219.2n ± 1%   1728.5n ± 193%   +688.55% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.63n ± 6%    17.49n ±   0%          ~ (p=0.052 n=6)
DeleteRandomPfxs/10_000       17.38n ± 1%    16.90n ±   1%     -2.76% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.48n ± 1%    17.41n ±   1%          ~ (p=0.078 n=6)
DeleteRandomPfxs/200_000      18.30n ± 3%    18.98n ±   1%     +3.66% (p=0.002 n=6)
geomean                       39.86n         121.6n          +204.95%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │            cidrtree/update.bm            │
                         │   sec/route    │    sec/route     vs base                 │
InsertRandomPfxs/1_000        43.46n ± 2%   1312.00n ±   2%  +2918.87% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.53n ± 1%   1943.00n ±   1%  +3463.50% (p=0.002 n=6)
InsertRandomPfxs/100_000      125.1n ± 3%    3175.5n ±   7%  +2438.37% (p=0.002 n=6)
InsertRandomPfxs/200_000      219.2n ± 1%    3871.5n ±   1%  +1666.20% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.63n ± 6%     14.81n ±   7%    -16.02% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.38n ± 1%     15.19n ±   1%    -12.57% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.48n ± 1%     23.79n ±   6%    +36.07% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.30n ± 3%    432.50n ± 135%  +2262.74% (p=0.002 n=6)
geomean                       39.86n          303.8n          +662.15%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │          critbitgo/update.bm          │
                         │   sec/route    │   sec/route    vs base                │
InsertRandomPfxs/1_000        43.46n ± 2%   172.60n ±  2%  +297.15% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.53n ± 1%   241.05n ±  1%  +342.09% (p=0.002 n=6)
InsertRandomPfxs/100_000      125.1n ± 3%    564.4n ±  1%  +351.12% (p=0.002 n=6)
InsertRandomPfxs/200_000      219.2n ± 1%    767.9n ±  3%  +250.32% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.63n ± 6%    70.89n ±  3%  +301.96% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.38n ± 1%    73.51n ±  1%  +322.96% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.48n ± 1%    78.01n ± 14%  +346.28% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.30n ± 3%    82.72n ±  2%  +351.87% (p=0.002 n=6)
geomean                       39.86n         167.0n        +319.07%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           lpmtrie/update.bm           │
                         │   sec/route    │  sec/route    vs base                 │
InsertRandomPfxs/1_000        43.46n ± 2%   385.15n ± 2%   +786.22% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.53n ± 1%   480.85n ± 6%   +781.89% (p=0.002 n=6)
InsertRandomPfxs/100_000      125.1n ± 3%   1148.0n ± 2%   +817.67% (p=0.002 n=6)
InsertRandomPfxs/200_000      219.2n ± 1%   1252.0n ± 4%   +471.17% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.63n ± 6%    85.49n ± 0%   +384.77% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.38n ± 1%   146.30n ± 0%   +741.77% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.48n ± 1%   320.20n ± 2%  +1731.81% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.30n ± 3%   465.30n ± 2%  +2441.93% (p=0.002 n=6)
geomean                       39.86n         386.3n        +869.16%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           cidranger/update.bm            │
                         │   sec/route    │   sec/route     vs base                  │
InsertRandomPfxs/1_000        43.46n ± 2%   4495.00n ±  1%  +10242.84% (p=0.002 n=6)
InsertRandomPfxs/10_000       54.53n ± 1%   6452.00n ±  3%  +11733.10% (p=0.002 n=6)
InsertRandomPfxs/100_000      125.1n ± 3%   10631.5n ±  5%   +8398.40% (p=0.002 n=6)
InsertRandomPfxs/200_000      219.2n ± 1%   11937.0n ±  3%   +5345.71% (p=0.002 n=6)
DeleteRandomPfxs/1_000        17.63n ± 6%     95.47n ±  2%    +441.39% (p=0.002 n=6)
DeleteRandomPfxs/10_000       17.38n ± 1%     92.30n ±  5%    +431.10% (p=0.002 n=6)
DeleteRandomPfxs/100_000      17.48n ± 1%    142.85n ±  2%    +717.22% (p=0.002 n=6)
DeleteRandomPfxs/200_000      18.30n ± 3%    722.80n ± 20%   +3848.65% (p=0.002 n=6)
geomean                       39.86n          1.163µ         +2817.94%
```
