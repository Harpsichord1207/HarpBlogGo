init:
    go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.cn,direct
	go mod download

build:
    go build .

run:
    nohup ./HarpBlogGo >go.log 2>&1 &
