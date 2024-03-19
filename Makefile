.PHONY: all dep

all: size update lookup

dep:
	go mod tidy
	go install golang.org/x/perf/cmd/benchstat@latest

size: art/size.bm bart/size.bm cidrtree/size.bm critbitgo/size.bm lpmtrie/size.bm
	@echo
	@benchstat -ignore=cpu,pkg,goos,goarch bart/size.bm art/size.bm
	@echo
	@benchstat -ignore=cpu,pkg,goos,goarch bart/size.bm cidrtree/size.bm
	@echo
	@benchstat -ignore=cpu,pkg,goos,goarch bart/size.bm critbitgo/size.bm
	@echo
	@benchstat -ignore=cpu,pkg,goos,goarch bart/size.bm lpmtrie/size.bm

update: art/update.bm bart/update.bm cidrtree/update.bm critbitgo/update.bm lpmtrie/update.bm
	@echo
	@benchstat -ignore=cpu,pkg,goos,goarch bart/update.bm art/update.bm
	@echo
	@benchstat -ignore=cpu,pkg,goos,goarch bart/update.bm cidrtree/update.bm
	@echo
	@benchstat -ignore=cpu,pkg,goos,goarch bart/update.bm critbitgo/update.bm
	@echo
	@benchstat -ignore=cpu,pkg,goos,goarch bart/update.bm lpmtrie/update.bm

lookup: art/lookup.bm bart/lookup.bm cidrtree/lookup.bm critbitgo/lookup.bm lpmtrie/lookup.bm
	@echo
	@benchstat -ignore=cpu,pkg,goos,goarch bart/lookup.bm art/lookup.bm
	@echo
	@benchstat -ignore=cpu,pkg,goos,goarch bart/lookup.bm cidrtree/lookup.bm
	@echo
	@benchstat -ignore=cpu,pkg,goos,goarch bart/lookup.bm critbitgo/lookup.bm
	@echo
	@benchstat -ignore=cpu,pkg,goos,goarch bart/lookup.bm lpmtrie/lookup.bm

#
# benchmarks for lpm lookup
#
art/lookup.bm:
	cd art &&       go test -run=XXX  -cpu=1 -count=10 -bench=Lpm | tee lookup.bm

bart/lookup.bm:
	cd bart &&      go test -run=XXX  -cpu=1 -count=10 -bench=Lpm | tee lookup.bm

cidrtree/lookup.bm:
	cd cidrtree &&  go test -run=XXX  -cpu=1 -count=10 -bench=Lpm | tee lookup.bm

critbitgo/lookup.bm:
	cd critbitgo && go test -run=XXX  -cpu=1 -count=10 -bench=Lpm | tee lookup.bm

lpmtrie/lookup.bm:
	cd lpmtrie &&   go test -run=XXX  -cpu=1 -count=10 -bench=Lpm | tee lookup.bm

# TODO more lookup

#
# benchmarks for tree/trie sizes, deterministic -> -benchtime=1x
#
art/size.bm:
	cd art && go test -run=XXX  -cpu=1 -count=10 -bench=Size -benchtime=1x      | tee size.bm

bart/size.bm:
	cd bart && go test -run=XXX  -cpu=1 -count=10 -bench=Size -benchtime=1x     | tee size.bm

cidrtree/size.bm:
	cd cidrtree && go test -run=XXX  -cpu=1 -count=10 -bench=Size -benchtime=1x  | tee size.bm

critbitgo/size.bm:
	cd critbitgo && go test -run=XXX  -cpu=1 -count=10 -bench=Size -benchtime=1x | tee size.bm

lpmtrie/size.bm:
	cd lpmtrie && go test -run=XXX  -cpu=1 -count=10 -bench=Size -benchtime=1x   | tee size.bm

#
# benchmarks for insert/delete
#

art/update.bm:
	cd art && go test -run=XXX  -cpu=1 -count=10 -bench='Insert|Delete'       | tee update.bm

bart/update.bm:
	cd bart && go test -run=XXX  -cpu=1 -count=10 -bench='Insert|Delete'      | tee update.bm

cidrtree/update.bm:
	cd cidrtree && go test -run=XXX  -cpu=1 -count=10 -bench='Insert|Delete'  | tee update.bm

critbitgo/update.bm:
	cd critbitgo && go test -run=XXX  -cpu=1 -count=10 -bench='Insert|Delete' | tee update.bm

lpmtrie/update.bm:
	cd lpmtrie && go test -run=XXX  -cpu=1 -count=10 -bench='Insert|Delete'   | tee update.bm
