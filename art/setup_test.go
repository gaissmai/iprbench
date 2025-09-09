package main_test

import (
	"local/iprbench/common"
)

const pfxFile = "../testdata/prefixes.txt.gz"

var (
	tier1Routes   = common.ReadFullTableShuffled(pfxFile)
	randomRoutes  = common.RandomRealWorldPrefixes(1_000_000)
	randomRoutes4 = common.RandomRealWorldPrefixes4(1_000_000)
	randomRoutes6 = common.RandomRealWorldPrefixes6(1_000_000)
)
