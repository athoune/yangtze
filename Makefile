build: vendor
	go build

test: vendor
	go test -v github.com/athoune/yangtze/index
	go test -v github.com/athoune/yangtze/serialize

vendor:
	dep ensure

clean:
	rm -rf vendor yangtze
