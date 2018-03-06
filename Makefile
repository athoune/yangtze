build: vendor
	go build

test: vendor
	go test -v -cover github.com/athoune/yangtze/index
	go test -v -cover github.com/athoune/yangtze/pattern
	go test -v -cover github.com/athoune/yangtze/serialize
	go test -v -cover github.com/athoune/yangtze/set
	go test -v -cover github.com/athoune/yangtze/store
	go test -v -cover github.com/athoune/yangtze/token

vendor:
	dep ensure

clean:
	rm -rf vendor yangtze
