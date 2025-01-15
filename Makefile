.PHONY: all dep

all: size lpm update

dep:
	go mod tidy
	go install golang.org/x/perf/cmd/benchstat@latest

size: art/size.bm bart/size.bm cidrtree/size.bm critbitgo/size.bm lpmtrie/size.bm cidranger/size.bm netipds/size.bm
	@echo
	@benchstat -ignore=pkg bart/size.bm   art/size.bm
	@echo
	@benchstat -ignore=pkg bart/size.bm   cidrtree/size.bm
	@echo
	@benchstat -ignore=pkg bart/size.bm   critbitgo/size.bm
	@echo
	@benchstat -ignore=pkg bart/size.bm   lpmtrie/size.bm
	@echo
	@benchstat -ignore=pkg bart/size.bm   cidranger/size.bm
	@echo
	@benchstat -ignore=pkg bart/size.bm   netipds/size.bm

lpm: art/lpm.bm bart/lpm.bm cidrtree/lpm.bm critbitgo/lpm.bm lpmtrie/lpm.bm cidranger/lpm.bm netipds/lpm.bm
	@echo
	@benchstat -ignore=pkg bart/lpm.bm   art/lpm.bm
	@echo
	@benchstat -ignore=pkg bart/lpm.bm   cidrtree/lpm.bm
	@echo
	@benchstat -ignore=pkg bart/lpm.bm   critbitgo/lpm.bm
	@echo
	@benchstat -ignore=pkg bart/lpm.bm   lpmtrie/lpm.bm
	@echo
	@benchstat -ignore=pkg bart/lpm.bm   cidranger/lpm.bm
	@echo
	@benchstat -ignore=pkg bart/lpm.bm   netipds/lpm.bm

update: art/update.bm bart/update.bm cidrtree/update.bm critbitgo/update.bm lpmtrie/update.bm cidranger/update.bm netipds/update.bm
	@echo
	@benchstat -ignore=pkg bart/update.bm    art/update.bm
	@echo
	@benchstat -ignore=pkg bart/update.bm    cidrtree/update.bm
	@echo
	@benchstat -ignore=pkg bart/update.bm    critbitgo/update.bm
	@echo
	@benchstat -ignore=pkg bart/update.bm    lpmtrie/update.bm
	@echo
	@benchstat -ignore=pkg bart/update.bm    cidranger/update.bm
	@echo
	@benchstat -ignore=pkg bart/update.bm    netipds/update.bm

#
# benchmarks for lpm
#
art/lpm.bm:
	cd art &&       go test -run=XXX  -cpu=1 -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

bart/lpm.bm:
	cd bart &&      go test -run=XXX  -cpu=1 -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

cidrtree/lpm.bm:
	cd cidrtree &&  go test -run=XXX  -cpu=1 -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

critbitgo/lpm.bm:
	cd critbitgo && go test -run=XXX  -cpu=1 -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

lpmtrie/lpm.bm:
	cd lpmtrie &&   go test -run=XXX  -cpu=1 -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

cidranger/lpm.bm:
	cd cidranger && go test -run=XXX  -cpu=1 -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

netipds/lpm.bm:
	cd netipds && go test -run=XXX  -cpu=1 -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

#
# benchmarks for tree/trie sizes, deterministic -> -benchtime=1x
#
art/size.bm:
	cd art &&       go test -run=XXX  -cpu=1 -count=6 -bench=Size -benchtime=1x -timeout=25m   | tee size.bm

bart/size.bm:
	cd bart &&      go test -run=XXX  -cpu=1 -count=6 -bench=Size -benchtime=1x -timeout=25m   | tee size.bm

cidrtree/size.bm:
	cd cidrtree &&  go test -run=XXX  -cpu=1 -count=6 -bench=Size -benchtime=1x -timeout=25m | tee size.bm

critbitgo/size.bm:
	cd critbitgo && go test -run=XXX  -cpu=1 -count=6 -bench=Size -benchtime=1x -timeout=25m | tee size.bm

lpmtrie/size.bm:
	cd lpmtrie &&   go test -run=XXX  -cpu=1 -count=6 -bench=Size -benchtime=1x -timeout=25m | tee size.bm

cidranger/size.bm:
	cd cidranger && go test -run=XXX  -cpu=1 -count=6 -bench=Size -benchtime=1x -timeout=25m | tee size.bm

netipds/size.bm:
	cd netipds && go test -run=XXX  -cpu=1 -count=6 -bench=Size -benchtime=1x -timeout=25m | tee size.bm

#
# benchmarks for insert/delete
#

art/update.bm:
	cd art &&       go test -run=XXX  -cpu=1 -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

bart/update.bm:
	cd bart &&      go test -run=XXX  -cpu=1 -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

cidrtree/update.bm:
	cd cidrtree &&  go test -run=XXX  -cpu=1 -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

critbitgo/update.bm:
	cd critbitgo && go test -run=XXX  -cpu=1 -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

lpmtrie/update.bm:
	cd lpmtrie &&   go test -run=XXX  -cpu=1 -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

cidranger/update.bm:
	cd cidranger && go test -run=XXX  -cpu=1 -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

netipds/update.bm:
	cd netipds && go test -run=XXX  -cpu=1 -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm
