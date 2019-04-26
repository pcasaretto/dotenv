all: bin/dotenv-Darwin-x86_64 bin/dotenv-Linux-x86_64 bin/dotenv-Linux-aarch64
clean:
	rm bin/*
bin/dotenv-Darwin-x86_64: main.go internal/parser/parser.go
	env GOARCH="amd64" GOOS="darwin" go build -o bin/dotenv-Darwin-x86_64
bin/dotenv-Linux-x86_64: main.go internal/parser/parser.go
	env GOARCH="amd64" GOOS="linux" go build -o bin/dotenv-Linux-x86_64
bin/dotenv-Linux-aarch64: main.go internal/parser/parser.go
	env GOARCH="arm64" GOOS="linux" go build -o bin/dotenv-Linux-aarch64
