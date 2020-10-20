init:
	export PATH=$PATH:/usr/local/go/bin && go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && go mod download

build:
	export PATH=$PATH:/usr/local/go/bin && go build .

run:
	nohup /home/HarpBlogGo/HarpBlog >std.log 2>&1 &
