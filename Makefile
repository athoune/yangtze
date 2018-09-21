build: vendor
	go build

test: vendor
	go test -v -cover github.com/athoune/yangtze/index
	go test -v -cover github.com/athoune/yangtze/pattern
	go test -v -cover github.com/athoune/yangtze/serialize
	go test -v -cover github.com/athoune/yangtze/set
	go test -v -cover github.com/athoune/yangtze/store
	go test -v -cover github.com/athoune/yangtze/token

bench: vendor
	cd index && go test -v -bench . -benchmem

bench_index: vendor
	cd index && go test -v -bench BenchmarkIndex -cpuprofile=cpu.out && go tool pprof —http cpu.out

clean-index:
	rm -f index/cpu* index/profile*.pdf

vendor:
	dep ensure

clean:
	rm -rf vendor yangtze
