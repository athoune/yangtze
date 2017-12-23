build: vendor
	go build

test: vendor
	go test -v github.com/athoune/yangtze/index

vendor:
	dep ensure

clean:
	rm -rf vendor yangtze
