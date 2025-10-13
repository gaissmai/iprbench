.PHONY: all dep

all: size lpm update

dep:
	go mod tidy
	go install tool

size: bart/size.bm lite/size.bm fast/size.bm netipds/size.bm critbitgo/size.bm lpmtrie/size.bm kentik-patricia/size.bm
	@echo
	@go tool benchstat -ignore=pkg bart/size.bm   lite/size.bm
	@echo
	@go tool benchstat -ignore=pkg bart/size.bm   fast/size.bm
	@echo
	@go tool benchstat -ignore=pkg bart/size.bm   netipds/size.bm
	@echo
	@go tool benchstat -ignore=pkg bart/size.bm   critbitgo/size.bm
	@echo
	@go tool benchstat -ignore=pkg bart/size.bm   lpmtrie/size.bm
	@echo
	@go tool benchstat -ignore=pkg bart/size.bm   kentik-patricia/size.bm

lpm: bart/lpm.bm lite/lpm.bm fast/lpm.bm netipds/lpm.bm critbitgo/lpm.bm lpmtrie/lpm.bm kentik-patricia/lpm.bm
	@echo
	@go tool benchstat -ignore=pkg bart/lpm.bm   lite/lpm.bm
	@echo
	@go tool benchstat -ignore=pkg bart/lpm.bm   fast/lpm.bm
	@echo
	@go tool benchstat -ignore=pkg bart/lpm.bm   netipds/lpm.bm
	@echo
	@go tool benchstat -ignore=pkg bart/lpm.bm   critbitgo/lpm.bm
	@echo
	@go tool benchstat -ignore=pkg bart/lpm.bm   lpmtrie/lpm.bm
	@echo
	@go tool benchstat -ignore=pkg bart/lpm.bm   kentik-patricia/lpm.bm


update: bart/update.bm lite/update.bm fast/update.bm netipds/update.bm critbitgo/update.bm lpmtrie/update.bm kentik-patricia/update.bm
	@echo
	@go tool benchstat -ignore=pkg bart/update.bm    lite/update.bm
	@echo
	@go tool benchstat -ignore=pkg bart/update.bm    fast/update.bm
	@echo
	@go tool benchstat -ignore=pkg bart/update.bm    netipds/update.bm
	@echo
	@go tool benchstat -ignore=pkg bart/update.bm    critbitgo/update.bm
	@echo
	@go tool benchstat -ignore=pkg bart/update.bm    lpmtrie/update.bm
	@echo
	@go tool benchstat -ignore=pkg bart/update.bm    kentik-patricia/update.bm

#
# benchmarks for lpm
#
bart/lpm.bm:
	cd bart &&      go test -run=XXX  -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

lite/lpm.bm:
	cd lite &&      go test -run=XXX  -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

fast/lpm.bm:
	cd fast &&      go test -run=XXX  -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

netipds/lpm.bm:
	cd netipds && go test -run=XXX  -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

critbitgo/lpm.bm:
	cd critbitgo && go test -run=XXX  -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

lpmtrie/lpm.bm:
	cd lpmtrie &&   go test -run=XXX  -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

kentik-patricia/lpm.bm:
	cd kentik-patricia &&  go test -run=XXX  -count=20  -bench=Lpm -timeout=25m | tee lpm.bm

#
# benchmarks for tree/trie sizes, deterministic -> -benchtime=1x
#
bart/size.bm:
	cd bart &&      go test -run=XXX  -count=6 -bench=Size -benchtime=1x -timeout=25m   | tee size.bm

lite/size.bm:
	cd lite &&      go test -run=XXX  -count=6 -bench=Size -benchtime=1x -timeout=25m   | tee size.bm

fast/size.bm:
	cd fast &&      go test -run=XXX  -count=6 -bench=Size -benchtime=1x -timeout=25m   | tee size.bm

netipds/size.bm:
	cd netipds && go test -run=XXX  -count=6 -bench=Size -benchtime=1x -timeout=25m | tee size.bm

critbitgo/size.bm:
	cd critbitgo && go test -run=XXX  -count=6 -bench=Size -benchtime=1x -timeout=25m | tee size.bm

lpmtrie/size.bm:
	cd lpmtrie &&   go test -run=XXX  -count=6 -bench=Size -benchtime=1x -timeout=25m | tee size.bm

kentik-patricia/size.bm:
	cd kentik-patricia &&  go test -run=XXX  -count=6 -bench=Size -benchtime=1x -timeout=25m | tee size.bm

#
# benchmarks for insert/delete
#

bart/update.bm:
	cd bart &&      go test -run=XXX  -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

lite/update.bm:
	cd lite &&      go test -run=XXX  -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

fast/update.bm:
	cd fast &&      go test -run=XXX  -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

netipds/update.bm:
	cd netipds && go test -run=XXX  -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

critbitgo/update.bm:
	cd critbitgo && go test -run=XXX  -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

lpmtrie/update.bm:
	cd lpmtrie &&   go test -run=XXX  -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

kentik-patricia/update.bm:
	cd kentik-patricia &&  go test -run=XXX  -count=6 -bench='Insert|Delete' -timeout=25m | tee update.bm

