package main_test

import (
	"local/iprbench/common"
)

const pfxFile = "../testdata/prefixes.txt.gz"

var (
	prng = common.Prng

	tier1Routes   = common.ReadFullTableShuffled(pfxFile)
	randomRoutes  = common.RandomPrefixes(1_000_000)
	randomRoutes4 = common.RandomPrefixes4(1_000_000)
	randomRoutes6 = common.RandomPrefixes6(1_000_000)

	probe = tier1Routes[prng.IntN(len(tier1Routes))]
	sink  any
)
