build: vendor
	go build

test: vendor
	go test -v github.com/athoune/yangtze/index
	go test -v github.com/athoune/yangtze/pattern
	go test -v github.com/athoune/yangtze/serialize
	go test -v github.com/athoune/yangtze/set
	go test -v github.com/athoune/yangtze/store

vendor:
	dep ensure

clean:
	rm -rf vendor yangtze
