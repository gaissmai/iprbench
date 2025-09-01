.PHONY: all dep

all: size lpm update

dep:
	go mod tidy
	go install tool

size: bart/size.bm fat/size.bm art/size.bm netipds/size.bm critbitgo/size.bm lpmtrie/size.bm \
	cidranger/size.bm cidrtree/size.bm go-iptrie/size.bm kentik-patricia/size.bm
	@echo
	@go tool benchstat -ignore=pkg bart/size.bm   fat/size.bm
	@echo
	@go tool benchstat -ignore=pkg bart/size.bm   art/size.bm
	@echo
	@go tool benchstat -ignore=pkg bart/size.bm   netipds/size.bm
	@echo
	@go tool benchstat -ignore=pkg bart/size.bm   critbitgo/size.bm
	@echo
	@go tool benchstat -ignore=pkg bart/size.bm   lpmtrie/size.bm
	@echo
	@go tool benchstat -ignore=pkg bart/size.bm   cidranger/size.bm
	@echo
	@go tool benchstat -ignore=pkg bart/size.bm   cidrtree/size.bm
	@echo
	@go tool benchstat -ignore=pkg bart/size.bm   go-iptrie/size.bm
	@echo
	@go tool benchstat -ignore=pkg bart/size.bm   kentik-patricia/size.bm

lpm: bart/lpm.bm fat/lpm.bm art/lpm.bm netipds/lpm.bm critbitgo/lpm.bm lpmtrie/lpm.bm \
	cidranger/lpm.bm cidrtree/lpm.bm go-iptrie/lpm.bm kentik-patricia/lpm.bm
	@echo
	@go tool benchstat -ignore=pkg bart/lpm.bm   fat/lpm.bm
	@echo
	@go tool benchstat -ignore=pkg bart/lpm.bm   art/lpm.bm
	@echo
	@go tool benchstat -ignore=pkg bart/lpm.bm   netipds/lpm.bm
	@echo
	@go tool benchstat -ignore=pkg bart/lpm.bm   critbitgo/lpm.bm
	@echo
	@go tool benchstat -ignore=pkg bart/lpm.bm   lpmtrie/lpm.bm
	@echo
	@go tool benchstat -ignore=pkg bart/lpm.bm   cidranger/lpm.bm
	@echo
	@go tool benchstat -ignore=pkg bart/lpm.bm   cidrtree/lpm.bm
	@echo
	@go tool benchstat -ignore=pkg bart/lpm.bm   go-iptrie/lpm.bm
	@echo
	@go tool benchstat -ignore=pkg bart/lpm.bm   kentik-patricia/lpm.bm


update: bart/update.bm fat/update.bm art/update.bm netipds/update.bm critbitgo/update.bm lpmtrie/update.bm \
	cidranger/update.bm cidrtree/update.bm go-iptrie/update.bm kentik-patricia/update.bm
	@echo
	@go tool benchstat -ignore=pkg bart/update.bm    fat/update.bm
	@echo
	@go tool benchstat -ignore=pkg bart/update.bm    art/update.bm
	@echo
	@go tool benchstat -ignore=pkg bart/update.bm    netipds/update.bm
	@echo
	@go tool benchstat -ignore=pkg bart/update.bm    critbitgo/update.bm
	@echo
	@go tool benchstat -ignore=pkg bart/update.bm    lpmtrie/update.bm
	@echo
	@go tool benchstat -ignore=pkg bart/update.bm    cidranger/update.bm
	@echo
	@go tool benchstat -ignore=pkg bart/update.bm    cidrtree/update.bm
	@echo
	@go tool benchstat -ignore=pkg bart/update.bm    go-iptrie/update.bm
	@echo
	@go tool benchstat -ignore=pkg bart/update.bm    kentik-patricia/update.bm

#
# benchmarks for lpm
#
bart/lpm.bm:
	cd bart &&      go test -run=XXX  -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

fat/lpm.bm:
	cd fat &&       go test -run=XXX  -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

art/lpm.bm:
	cd art &&       go test -run=XXX  -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

netipds/lpm.bm:
	cd netipds && go test -run=XXX  -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

critbitgo/lpm.bm:
	cd critbitgo && go test -run=XXX  -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

lpmtrie/lpm.bm:
	cd lpmtrie &&   go test -run=XXX  -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

cidranger/lpm.bm:
	cd cidranger && go test -run=XXX  -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

cidrtree/lpm.bm:
	cd cidrtree &&  go test -run=XXX  -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

go-iptrie/lpm.bm:
	cd go-iptrie &&  go test -run=XXX  -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

kentik-patricia/lpm.bm:
	cd kentik-patricia &&  go test -run=XXX  -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

#
# benchmarks for tree/trie sizes, deterministic -> -benchtime=1x
#
bart/size.bm:
	cd bart &&      go test -run=XXX  -count=6 -bench=Size -benchtime=1x -timeout=25m   | tee size.bm

fat/size.bm:
	cd fat &&       go test -run=XXX  -count=6 -bench=Size -benchtime=1x -timeout=25m   | tee size.bm

art/size.bm:
	cd art &&       go test -run=XXX  -count=6 -bench=Size -benchtime=1x -timeout=25m   | tee size.bm

netipds/size.bm:
	cd netipds && go test -run=XXX  -count=6 -bench=Size -benchtime=1x -timeout=25m | tee size.bm

critbitgo/size.bm:
	cd critbitgo && go test -run=XXX  -count=6 -bench=Size -benchtime=1x -timeout=25m | tee size.bm

lpmtrie/size.bm:
	cd lpmtrie &&   go test -run=XXX  -count=6 -bench=Size -benchtime=1x -timeout=25m | tee size.bm

cidranger/size.bm:
	cd cidranger && go test -run=XXX  -count=6 -bench=Size -benchtime=1x -timeout=25m | tee size.bm

cidrtree/size.bm:
	cd cidrtree &&  go test -run=XXX  -count=6 -bench=Size -benchtime=1x -timeout=25m | tee size.bm

go-iptrie/size.bm:
	cd go-iptrie &&  go test -run=XXX  -count=6 -bench=Size -benchtime=1x -timeout=25m | tee size.bm

kentik-patricia/size.bm:
	cd kentik-patricia &&  go test -run=XXX  -count=6 -bench=Size -benchtime=1x -timeout=25m | tee size.bm

#
# benchmarks for insert/delete
#

bart/update.bm:
	cd bart &&      go test -run=XXX  -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

fat/update.bm:
	cd fat &&       go test -run=XXX  -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

art/update.bm:
	cd art &&       go test -run=XXX  -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

netipds/update.bm:
	cd netipds && go test -run=XXX  -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

critbitgo/update.bm:
	cd critbitgo && go test -run=XXX  -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

lpmtrie/update.bm:
	cd lpmtrie &&   go test -run=XXX  -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

cidranger/update.bm:
	cd cidranger && go test -run=XXX  -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

cidrtree/update.bm:
	cd cidrtree &&  go test -run=XXX  -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

go-iptrie/update.bm:
	cd go-iptrie &&  go test -run=XXX  -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

kentik-patricia/update.bm:
	cd kentik-patricia &&  go test -run=XXX  -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

