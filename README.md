# iprbench
comparing benchmarks for some golang IP routing table implementations:

```
	github.com/gaissmai/bart
	github.com/tailscale/art
	github.com/aromatt/netipds
	github.com/k-sone/critbitgo
	github.com/Asphaltt/lpmtrie
	github.com/yl2chen/cidranger
	github.com/gaissmai/cidrtree
```

The ~1_000_000 **Tier1** prefix test records (IPv4 and IPv6 routes) are from a full Internet routing table with typical
ISP prefix distribution.

In comparison, the prefix lengths for the _real-world_ random test sets are equally distributed between /8-28 for IPv4
and /16-56 bits for IPv6 (limited to the 2000::/3 global unicast address space).

The _real-world_ **RandomPrefixes** without IP versions labeling are composed of a distribution of 4 parts IPv4
to 1 part IPv6 random prefixes, which is approximately the current ratio in the Internet backbone routers.

## make your own benchmarks

```
  $ # set the cpu feature flags, e.g.
  $ export GOAMD64='v2'
  $ make dep
  $ make -B all   # takes some time!
```

## lpm (longest-prefix-match)

`bart` is the fastest software algorithm for IP-address longest-prefix-match.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │              art/lpm.bm               │
                                     │    sec/op    │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4            25.51n ± 10%   46.79n ± 12%   +83.42% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            34.35n ± 29%   61.22n ±  5%   +78.21% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             29.78n ±  3%   28.75n ±  0%    -3.43% (p=0.009 n=20)
LpmTier1Pfxs/RandomMissIP6             42.59n ± 31%   34.31n ± 25%         ~ (p=0.542 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     20.69n ±  4%   45.38n ±  1%  +119.39% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     20.69n ±  3%   47.25n ±  5%  +128.37% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      21.70n ±  7%   29.46n ±  5%   +35.78% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      21.42n ±  6%   30.27n ±  4%   +41.34% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    29.17n ±  5%   45.72n ±  1%   +56.72% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    23.70n ± 11%   46.74n ±  3%   +97.26% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     26.53n ±  7%   29.12n ±  6%    +9.74% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     23.20n ±  9%   30.32n ±  7%   +30.74% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   23.92n ± 41%   49.59n ± 12%  +107.34% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   31.21n ± 18%   50.32n ±  6%   +61.25% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    27.36n ±  5%   29.30n ±  6%    +7.07% (p=0.001 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    31.86n ±  2%   34.15n ±  4%    +7.20% (p=0.002 n=20)
geomean                                26.56n         38.71n         +45.78%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            netipds/lpm.bm             │
                                     │    sec/op    │    sec/op     vs base                 │
LpmTier1Pfxs/RandomMatchIP4            25.51n ± 10%   61.32n ± 15%  +140.40% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            34.35n ± 29%   52.59n ± 25%   +53.11% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             29.78n ±  3%   66.24n ±  6%  +122.49% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             42.59n ± 31%   55.12n ± 13%   +29.44% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     20.69n ±  4%   42.07n ±  7%  +103.41% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     20.69n ±  3%   32.99n ± 10%   +59.45% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      21.70n ±  7%   41.21n ±  6%   +89.91% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      21.42n ±  6%   33.11n ±  4%   +54.58% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    29.17n ±  5%   49.44n ± 11%   +69.49% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    23.70n ± 11%   37.98n ±  4%   +60.29% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     26.53n ±  7%   49.79n ±  4%   +87.66% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     23.20n ±  9%   38.73n ±  3%   +67.00% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   23.92n ± 41%   50.45n ± 13%  +110.96% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   31.21n ± 18%   42.61n ±  3%   +36.51% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    27.36n ±  5%   53.99n ±  3%   +97.30% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    31.86n ±  2%   44.87n ±  4%   +40.84% (p=0.000 n=20)
geomean                                26.56n         46.13n         +73.71%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            critbitgo/lpm.bm             │
                                     │    sec/op    │    sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            25.51n ± 10%   166.30n ±  6%   +551.90% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            34.35n ± 29%   252.75n ± 16%   +635.81% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             29.78n ±  3%   656.75n ± 20%  +2105.71% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             42.59n ± 31%   475.80n ± 12%  +1017.29% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     20.69n ±  4%    85.45n ±  7%   +313.08% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     20.69n ±  3%   112.30n ±  6%   +442.77% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      21.70n ±  7%   156.00n ±  9%   +618.89% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      21.42n ±  6%   172.65n ± 13%   +706.02% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    29.17n ±  5%    98.50n ±  6%   +237.68% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    23.70n ± 11%   125.85n ±  5%   +431.12% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     26.53n ±  7%   269.95n ± 13%   +917.53% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     23.20n ±  9%   230.75n ± 22%   +894.83% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   23.92n ± 41%   109.70n ±  7%   +358.71% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   31.21n ± 18%   133.90n ±  4%   +329.03% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    27.36n ±  5%   460.45n ± 14%  +1582.62% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    31.86n ±  2%   332.60n ± 15%   +944.11% (p=0.000 n=20)
geomean                                26.56n          198.7n         +648.18%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             lpmtrie/lpm.bm             │
                                     │    sec/op    │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            25.51n ± 10%   279.15n ±  7%  +994.28% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            34.35n ± 29%   327.60n ± 18%  +853.71% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             29.78n ±  3%   262.55n ±  4%  +781.78% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             42.59n ± 31%   228.75n ± 14%  +437.16% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     20.69n ±  4%   132.05n ±  6%  +538.39% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     20.69n ±  3%   120.90n ± 23%  +484.34% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      21.70n ±  7%   126.40n ± 17%  +482.49% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      21.42n ±  6%   106.90n ± 24%  +399.07% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    29.17n ±  5%   159.25n ± 10%  +445.94% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    23.70n ± 11%   143.80n ±  9%  +506.88% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     26.53n ±  7%   164.90n ± 11%  +521.56% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     23.20n ±  9%   138.25n ±  7%  +496.03% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   23.92n ± 41%   208.40n ±  6%  +771.42% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   31.21n ± 18%   188.35n ±  2%  +503.49% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    27.36n ±  5%   200.10n ±  5%  +631.23% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    31.86n ±  2%   195.85n ±  7%  +514.82% (p=0.000 n=20)
geomean                                26.56n          177.3n        +567.57%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │            cidranger/lpm.bm            │
                                     │    sec/op    │    sec/op      vs base                 │
LpmTier1Pfxs/RandomMatchIP4            25.51n ± 10%   265.15n ± 25%  +939.40% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            34.35n ± 29%   282.25n ± 27%  +721.69% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             29.78n ±  3%   282.45n ±  7%  +848.61% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             42.59n ± 31%   292.20n ± 22%  +586.16% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     20.69n ±  4%   115.70n ±  9%  +459.34% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     20.69n ±  3%   159.80n ± 12%  +672.35% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      21.70n ±  7%   113.20n ±  7%  +421.66% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      21.42n ±  6%   147.60n ±  8%  +589.08% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    29.17n ±  5%   152.20n ±  6%  +421.77% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    23.70n ± 11%   179.30n ± 11%  +656.70% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     26.53n ±  7%   156.55n ±  8%  +490.09% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     23.20n ±  9%   180.15n ±  7%  +676.68% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   23.92n ± 41%   169.25n ± 19%  +607.71% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   31.21n ± 18%   224.00n ±  8%  +617.72% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    27.36n ±  5%   184.70n ± 10%  +574.95% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    31.86n ±  2%   231.55n ±  5%  +626.89% (p=0.000 n=20)
geomean                                26.56n          187.7n        +606.95%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                                     │ bart/lpm.bm  │             cidrtree/lpm.bm              │
                                     │    sec/op    │     sec/op      vs base                  │
LpmTier1Pfxs/RandomMatchIP4            25.51n ± 10%    743.75n ± 17%  +2815.52% (p=0.000 n=20)
LpmTier1Pfxs/RandomMatchIP6            34.35n ± 29%    922.65n ± 16%  +2586.03% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP4             29.78n ±  3%   1131.00n ± 17%  +3698.49% (p=0.000 n=20)
LpmTier1Pfxs/RandomMissIP6             42.59n ± 31%   1415.00n ± 19%  +3222.77% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP4     20.69n ±  4%    316.15n ± 11%  +1428.40% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMatchIP6     20.69n ±  3%    333.35n ±  7%  +1511.16% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP4      21.70n ±  7%    434.80n ± 13%  +1903.69% (p=0.000 n=20)
LpmRandomPfxs/1_000/RandomMissIP6      21.42n ±  6%    507.40n ± 10%  +2268.81% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP4    29.17n ±  5%    474.05n ± 23%  +1525.13% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMatchIP6    23.70n ± 11%    569.80n ± 15%  +2304.73% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP4     26.53n ±  7%    672.80n ± 27%  +2436.00% (p=0.000 n=20)
LpmRandomPfxs/10_000/RandomMissIP6     23.20n ±  9%    715.95n ± 14%  +2986.66% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP4   23.92n ± 41%    602.35n ± 23%  +2418.71% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMatchIP6   31.21n ± 18%    826.40n ± 18%  +2547.87% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP4    27.36n ±  5%    894.05n ± 11%  +3167.13% (p=0.000 n=20)
LpmRandomPfxs/100_000/RandomMissIP6    31.86n ±  2%    995.15n ± 28%  +3024.00% (p=0.000 n=20)
geomean                                26.56n           665.2n        +2405.11%
```

## size of the routing tables


`bart` has two orders of magnitude lower memory consumption compared to `art`
and is even better in low memory consumption to a binary search tree, like the `cidrtree`.


```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │              art/size.bm               │
                       │ bytes/route  │ bytes/route   vs base                  │
Tier1PfxSize/1_000         88.13 ± 2%   7591.00 ± 0%   +8513.41% (p=0.002 n=6)
Tier1PfxSize/10_000        66.04 ± 0%   4889.00 ± 0%   +7303.09% (p=0.002 n=6)
Tier1PfxSize/100_000       30.24 ± 0%   1669.00 ± 0%   +5419.18% (p=0.002 n=6)
Tier1PfxSize/200_000       20.48 ± 0%   1098.00 ± 0%   +5261.33% (p=0.002 n=6)
RandomPfx4Size/1_000       66.52 ± 3%   5259.00 ± 0%   +7805.89% (p=0.002 n=6)
RandomPfx4Size/10_000      39.71 ± 0%   4059.00 ± 0%  +10121.61% (p=0.002 n=6)
RandomPfx4Size/100_000     51.53 ± 0%   3938.00 ± 0%   +7542.15% (p=0.002 n=6)
RandomPfx4Size/200_000     47.46 ± 0%   3476.00 ± 0%   +7224.06% (p=0.002 n=6)
RandomPfx6Size/1_000       68.38 ± 3%   6761.00 ± 0%   +9787.39% (p=0.002 n=6)
RandomPfx6Size/10_000      84.46 ± 0%   7333.00 ± 0%   +8582.22% (p=0.002 n=6)
RandomPfx6Size/100_000     56.63 ± 0%   5708.00 ± 0%   +9979.46% (p=0.002 n=6)
RandomPfx6Size/200_000     53.30 ± 0%   5526.00 ± 0%  +10267.73% (p=0.002 n=6)
RandomPfxSize/1_000        85.30 ± 2%   7537.00 ± 0%   +8735.87% (p=0.002 n=6)
RandomPfxSize/10_000       57.99 ± 0%   6058.00 ± 0%  +10346.63% (p=0.002 n=6)
RandomPfxSize/100_000      66.54 ± 0%   5300.00 ± 0%   +7865.13% (p=0.002 n=6)
RandomPfxSize/200_000      58.51 ± 0%   4586.00 ± 0%   +7737.98% (p=0.002 n=6)
geomean                    55.38        4.449Ki        +8127.91%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │           netipds/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         88.13 ± 2%   101.00 ± 2%   +14.60% (p=0.002 n=6)
Tier1PfxSize/10_000        66.04 ± 0%    96.08 ± 0%   +45.49% (p=0.002 n=6)
Tier1PfxSize/100_000       30.24 ± 0%    94.34 ± 0%  +211.97% (p=0.002 n=6)
Tier1PfxSize/200_000       20.48 ± 0%    93.27 ± 0%  +355.42% (p=0.002 n=6)
RandomPfx4Size/1_000       66.52 ± 3%    90.32 ± 2%   +35.78% (p=0.002 n=6)
RandomPfx4Size/10_000      39.71 ± 0%    76.82 ± 0%   +93.45% (p=0.002 n=6)
RandomPfx4Size/100_000     51.53 ± 0%    72.35 ± 0%   +40.40% (p=0.002 n=6)
RandomPfx4Size/200_000     47.46 ± 0%    71.35 ± 0%   +50.34% (p=0.002 n=6)
RandomPfx6Size/1_000       68.38 ± 3%   100.20 ± 2%   +46.53% (p=0.002 n=6)
RandomPfx6Size/10_000      84.46 ± 0%    94.25 ± 0%   +11.59% (p=0.002 n=6)
RandomPfx6Size/100_000     56.63 ± 0%    87.63 ± 0%   +54.74% (p=0.002 n=6)
RandomPfx6Size/200_000     53.30 ± 0%    85.90 ± 0%   +61.16% (p=0.002 n=6)
RandomPfxSize/1_000        85.30 ± 2%    99.78 ± 2%   +16.98% (p=0.002 n=6)
RandomPfxSize/10_000       57.99 ± 0%    94.34 ± 0%   +62.68% (p=0.002 n=6)
RandomPfxSize/100_000      66.54 ± 0%    87.05 ± 0%   +30.82% (p=0.002 n=6)
RandomPfxSize/200_000      58.51 ± 0%    83.21 ± 0%   +42.22% (p=0.002 n=6)
geomean                    55.38         88.75        +60.28%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          critbitgo/size.bm          │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         88.13 ± 2%   119.50 ± 2%   +35.60% (p=0.002 n=6)
Tier1PfxSize/10_000        66.04 ± 0%   114.70 ± 0%   +73.68% (p=0.002 n=6)
Tier1PfxSize/100_000       30.24 ± 0%   114.40 ± 0%  +278.31% (p=0.002 n=6)
Tier1PfxSize/200_000       20.48 ± 0%   114.40 ± 0%  +458.59% (p=0.002 n=6)
RandomPfx4Size/1_000       66.52 ± 3%   116.10 ± 2%   +74.53% (p=0.002 n=6)
RandomPfx4Size/10_000      39.71 ± 0%   112.40 ± 0%  +183.05% (p=0.002 n=6)
RandomPfx4Size/100_000     51.53 ± 0%   112.00 ± 0%  +117.35% (p=0.002 n=6)
RandomPfx4Size/200_000     47.46 ± 0%   112.00 ± 0%  +135.99% (p=0.002 n=6)
RandomPfx6Size/1_000       68.38 ± 3%   132.40 ± 1%   +93.62% (p=0.002 n=6)
RandomPfx6Size/10_000      84.46 ± 0%   128.40 ± 0%   +52.02% (p=0.002 n=6)
RandomPfx6Size/100_000     56.63 ± 0%   128.00 ± 0%  +126.03% (p=0.002 n=6)
RandomPfx6Size/200_000     53.30 ± 0%   128.00 ± 0%  +140.15% (p=0.002 n=6)
RandomPfxSize/1_000        85.30 ± 2%   119.50 ± 2%   +40.09% (p=0.002 n=6)
RandomPfxSize/10_000       57.99 ± 0%   115.60 ± 0%   +99.34% (p=0.002 n=6)
RandomPfxSize/100_000      66.54 ± 0%   115.30 ± 0%   +73.28% (p=0.002 n=6)
RandomPfxSize/200_000      58.51 ± 0%   115.20 ± 0%   +96.89% (p=0.002 n=6)
geomean                    55.38         118.4       +113.90%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │           lpmtrie/size.bm            │
                       │ bytes/route  │ bytes/route   vs base                │
Tier1PfxSize/1_000         88.13 ± 2%   215.40 ±  5%  +144.41% (p=0.002 n=6)
Tier1PfxSize/10_000        66.04 ± 0%   210.50 ±  5%  +218.75% (p=0.002 n=6)
Tier1PfxSize/100_000       30.24 ± 0%   209.90 ±  5%  +594.11% (p=0.002 n=6)
Tier1PfxSize/200_000       20.48 ± 0%   209.20 ±  5%  +921.48% (p=0.002 n=6)
RandomPfx4Size/1_000       66.52 ± 3%   205.20 ±  8%  +208.48% (p=0.002 n=6)
RandomPfx4Size/10_000      39.71 ± 0%   186.60 ±  9%  +369.91% (p=0.002 n=6)
RandomPfx4Size/100_000     51.53 ± 0%   179.60 ±  9%  +248.53% (p=0.002 n=6)
RandomPfx4Size/200_000     47.46 ± 0%   178.50 ± 10%  +276.11% (p=0.002 n=6)
RandomPfx6Size/1_000       68.38 ± 3%   228.00 ±  8%  +233.43% (p=0.002 n=6)
RandomPfx6Size/10_000      84.46 ± 0%   222.50 ±  8%  +163.44% (p=0.002 n=6)
RandomPfx6Size/100_000     56.63 ± 0%   213.90 ±  9%  +277.71% (p=0.002 n=6)
RandomPfx6Size/200_000     53.30 ± 0%   210.50 ±  9%  +294.93% (p=0.002 n=6)
RandomPfxSize/1_000        85.30 ± 2%   215.00 ±  5%  +152.05% (p=0.002 n=6)
RandomPfxSize/10_000       57.99 ± 0%   210.30 ±  5%  +262.65% (p=0.002 n=6)
RandomPfxSize/100_000      66.54 ± 0%   203.50 ±  7%  +205.83% (p=0.002 n=6)
RandomPfxSize/200_000      58.51 ± 0%   198.70 ±  8%  +239.60% (p=0.002 n=6)
geomean                    55.38         205.6        +271.31%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidranger/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                 │
Tier1PfxSize/1_000         88.13 ± 2%   539.60 ± 3%   +512.28% (p=0.002 n=6)
Tier1PfxSize/10_000        66.04 ± 0%   533.90 ± 3%   +708.45% (p=0.002 n=6)
Tier1PfxSize/100_000       30.24 ± 0%   527.20 ± 2%  +1643.39% (p=0.002 n=6)
Tier1PfxSize/200_000       20.48 ± 0%   522.20 ± 2%  +2449.80% (p=0.002 n=6)
RandomPfx4Size/1_000       66.52 ± 3%   481.50 ± 3%   +623.84% (p=0.002 n=6)
RandomPfx4Size/10_000      39.71 ± 0%   433.10 ± 3%   +990.66% (p=0.002 n=6)
RandomPfx4Size/100_000     51.53 ± 0%   413.70 ± 3%   +702.83% (p=0.002 n=6)
RandomPfx4Size/200_000     47.46 ± 0%   409.10 ± 3%   +761.99% (p=0.002 n=6)
RandomPfx6Size/1_000       68.38 ± 3%   595.10 ± 0%   +770.28% (p=0.002 n=6)
RandomPfx6Size/10_000      84.46 ± 0%   581.00 ± 0%   +587.90% (p=0.002 n=6)
RandomPfx6Size/100_000     56.63 ± 0%   547.20 ± 0%   +866.27% (p=0.002 n=6)
RandomPfx6Size/200_000     53.30 ± 0%   538.10 ± 0%   +909.57% (p=0.002 n=6)
RandomPfxSize/1_000        85.30 ± 2%   540.30 ± 2%   +533.41% (p=0.002 n=6)
RandomPfxSize/10_000       57.99 ± 0%   528.10 ± 2%   +810.67% (p=0.002 n=6)
RandomPfxSize/100_000      66.54 ± 0%   495.50 ± 2%   +644.66% (p=0.002 n=6)
RandomPfxSize/200_000      58.51 ± 0%   477.60 ± 2%   +716.27% (p=0.002 n=6)
geomean                    55.38         507.3        +816.12%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                       │ bart/size.bm │          cidrtree/size.bm           │
                       │ bytes/route  │ bytes/route  vs base                │
Tier1PfxSize/1_000         88.13 ± 2%    69.26 ± 3%   -21.41% (p=0.002 n=6)
Tier1PfxSize/10_000        66.04 ± 0%    64.37 ± 0%    -2.53% (p=0.002 n=6)
Tier1PfxSize/100_000       30.24 ± 0%    64.04 ± 0%  +111.77% (p=0.002 n=6)
Tier1PfxSize/200_000       20.48 ± 0%    64.02 ± 0%  +212.60% (p=0.002 n=6)
RandomPfx4Size/1_000       66.52 ± 3%    68.16 ± 3%    +2.47% (p=0.013 n=6)
RandomPfx4Size/10_000      39.71 ± 0%    64.37 ± 0%   +62.10% (p=0.002 n=6)
RandomPfx4Size/100_000     51.53 ± 0%    64.04 ± 0%   +24.28% (p=0.002 n=6)
RandomPfx4Size/200_000     47.46 ± 0%    64.02 ± 0%   +34.89% (p=0.002 n=6)
RandomPfx6Size/1_000       68.38 ± 3%    68.41 ± 3%    +0.04% (p=0.013 n=6)
RandomPfx6Size/10_000      84.46 ± 0%    64.37 ± 0%   -23.79% (p=0.002 n=6)
RandomPfx6Size/100_000     56.63 ± 0%    64.04 ± 0%   +13.08% (p=0.002 n=6)
RandomPfx6Size/200_000     53.30 ± 0%    64.02 ± 0%   +20.11% (p=0.002 n=6)
RandomPfxSize/1_000        85.30 ± 2%    68.16 ± 3%   -20.09% (p=0.002 n=6)
RandomPfxSize/10_000       57.99 ± 0%    64.37 ± 0%   +11.00% (p=0.002 n=6)
RandomPfxSize/100_000      66.54 ± 0%    64.04 ± 0%    -3.76% (p=0.002 n=6)
RandomPfxSize/200_000      58.51 ± 0%    64.02 ± 0%    +9.42% (p=0.002 n=6)
geomean                    55.38         65.20        +17.75%
```

## update, insert/delete

`bart` is by far the fastest algorithm for updates.

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │              art/update.bm              │
                         │   sec/route    │   sec/route     vs base                 │
InsertRandomPfxs/1_000       31.78n ±  3%   298.35n ±   6%   +838.80% (p=0.002 n=6)
InsertRandomPfxs/10_000      40.62n ± 22%   587.20n ±   7%  +1345.59% (p=0.002 n=6)
InsertRandomPfxs/100_000     275.7n ± 13%   2717.0n ±  24%   +885.49% (p=0.002 n=6)
InsertRandomPfxs/200_000     282.9n ± 40%   1892.5n ± 167%   +568.96% (p=0.002 n=6)
DeleteRandomPfxs/1_000       15.74n ±  3%    16.68n ±   2%     +5.94% (p=0.002 n=6)
DeleteRandomPfxs/10_000      15.44n ±  1%    16.88n ±  31%     +9.29% (p=0.002 n=6)
DeleteRandomPfxs/100_000     15.78n ±  1%    17.89n ±  54%    +13.37% (p=0.002 n=6)
DeleteRandomPfxs/200_000     17.42n ± 13%    20.25n ±  16%    +16.25% (p=0.009 n=6)
geomean                      40.14n          132.0n          +228.77%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           netipds/update.bm            │
                         │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000       31.78n ±  3%   135.35n ±  5%   +325.90% (p=0.002 n=6)
InsertRandomPfxs/10_000      40.62n ± 22%   195.60n ± 16%   +381.54% (p=0.002 n=6)
InsertRandomPfxs/100_000     275.7n ± 13%    431.6n ±  7%    +56.55% (p=0.002 n=6)
InsertRandomPfxs/200_000     282.9n ± 40%    606.6n ±  6%   +114.42% (p=0.002 n=6)
DeleteRandomPfxs/1_000       15.74n ±  3%   113.80n ±  3%   +622.77% (p=0.002 n=6)
DeleteRandomPfxs/10_000      15.44n ±  1%   170.00n ±  6%  +1000.68% (p=0.002 n=6)
DeleteRandomPfxs/100_000     15.78n ±  1%   411.15n ±  9%  +2504.69% (p=0.002 n=6)
DeleteRandomPfxs/200_000     17.42n ± 13%   620.20n ± 19%  +3460.28% (p=0.002 n=6)
geomean                      40.14n          276.5n         +589.00%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │          critbitgo/update.bm          │
                         │   sec/route    │   sec/route    vs base                │
InsertRandomPfxs/1_000       31.78n ±  3%   167.85n ±  4%  +428.16% (p=0.002 n=6)
InsertRandomPfxs/10_000      40.62n ± 22%   256.10n ± 21%  +530.48% (p=0.002 n=6)
InsertRandomPfxs/100_000     275.7n ± 13%    682.5n ± 28%  +147.55% (p=0.002 n=6)
InsertRandomPfxs/200_000     282.9n ± 40%    796.9n ± 11%  +181.67% (p=0.002 n=6)
DeleteRandomPfxs/1_000       15.74n ±  3%    70.56n ± 26%  +348.11% (p=0.002 n=6)
DeleteRandomPfxs/10_000      15.44n ±  1%    74.12n ±  3%  +379.93% (p=0.002 n=6)
DeleteRandomPfxs/100_000     15.78n ±  1%    78.38n ± 10%  +396.52% (p=0.002 n=6)
DeleteRandomPfxs/200_000     17.42n ± 13%    88.89n ±  3%  +410.30% (p=0.002 n=6)
geomean                      40.14n          174.3n        +334.28%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           lpmtrie/update.bm            │
                         │   sec/route    │   sec/route    vs base                 │
InsertRandomPfxs/1_000       31.78n ±  3%   390.00n ± 14%  +1127.19% (p=0.002 n=6)
InsertRandomPfxs/10_000      40.62n ± 22%   512.30n ± 37%  +1161.20% (p=0.002 n=6)
InsertRandomPfxs/100_000     275.7n ± 13%   1269.0n ± 21%   +360.28% (p=0.002 n=6)
InsertRandomPfxs/200_000     282.9n ± 40%   1285.5n ±  3%   +354.40% (p=0.002 n=6)
DeleteRandomPfxs/1_000       15.74n ±  3%    73.59n ±  3%   +367.35% (p=0.002 n=6)
DeleteRandomPfxs/10_000      15.44n ±  1%   134.05n ±  4%   +767.92% (p=0.002 n=6)
DeleteRandomPfxs/100_000     15.78n ±  1%   351.60n ±  8%  +2127.43% (p=0.002 n=6)
DeleteRandomPfxs/200_000     17.42n ± 13%   565.35n ± 12%  +3145.41% (p=0.002 n=6)
geomean                      40.14n          398.7n         +893.50%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │           cidranger/update.bm            │
                         │   sec/route    │   sec/route     vs base                  │
InsertRandomPfxs/1_000       31.78n ±  3%   4597.00n ±  3%  +14365.07% (p=0.002 n=6)
InsertRandomPfxs/10_000      40.62n ± 22%   6772.50n ±  8%  +16572.82% (p=0.002 n=6)
InsertRandomPfxs/100_000     275.7n ± 13%   10831.0n ±  2%   +3828.55% (p=0.002 n=6)
InsertRandomPfxs/200_000     282.9n ± 40%   11676.0n ±  8%   +4027.25% (p=0.002 n=6)
DeleteRandomPfxs/1_000       15.74n ±  3%    107.20n ± 42%    +580.85% (p=0.002 n=6)
DeleteRandomPfxs/10_000      15.44n ±  1%     93.51n ±  3%    +505.44% (p=0.002 n=6)
DeleteRandomPfxs/100_000     15.78n ±  1%    123.50n ± 11%    +682.39% (p=0.002 n=6)
DeleteRandomPfxs/200_000     17.42n ± 13%    251.35n ± 13%   +1342.88% (p=0.002 n=6)
geomean                      40.14n           1.026µ         +2455.64%

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
                         │ bart/update.bm │            cidrtree/update.bm             │
                         │   sec/route    │    sec/route      vs base                 │
InsertRandomPfxs/1_000       31.78n ±  3%   1256.50n ±    4%  +3853.74% (p=0.002 n=6)
InsertRandomPfxs/10_000      40.62n ± 22%   1845.50n ±    2%  +4443.33% (p=0.002 n=6)
InsertRandomPfxs/100_000     275.7n ± 13%    3209.5n ±   14%  +1064.13% (p=0.002 n=6)
InsertRandomPfxs/200_000     282.9n ± 40%    4354.0n ±    7%  +1439.06% (p=0.002 n=6)
DeleteRandomPfxs/1_000       15.74n ±  3%     15.04n ±    4%     -4.45% (p=0.009 n=6)
DeleteRandomPfxs/10_000      15.44n ±  1%     15.18n ±    3%          ~ (p=0.240 n=6)
DeleteRandomPfxs/100_000     15.78n ±  1%     23.77n ±   12%    +50.55% (p=0.002 n=6)
DeleteRandomPfxs/200_000     17.42n ± 13%    228.80n ± 2280%  +1213.43% (p=0.002 n=6)
geomean                      40.14n           282.2n           +603.17%
```
