# Golang Wasm Practice

## Requirement

```bash
$ go version
go version go1.16.2 darwin/arm64
```

```bash
go get -u github.com/shurcooL/goexec
```

## Build

Go
```bash
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./
GOOS=js GOARCH=wasm go build -o hello.wasm hello.go
```

[Tinygo](https://tinygo.org/)
```bash
tinygo build -o hello.wasm -target wasm ./hello.go
```

## Run

```bash
goexec 'http.ListenAndServe(`:9000`, http.FileServer(http.Dir(`.`)))'
open http://localhost:9000
```
