# Golang Wasm Practice

<kbd><img width="818" alt="image" src="https://user-images.githubusercontent.com/17779386/170900973-ec5213f4-b799-4d12-9200-5515b1f8d296.png"></kbd>



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
GOOS=js GOARCH=wasm go build -o mine.wasm -trimpath
```

[Tinygo](https://tinygo.org/)

```bash
tinygo build -o mine.wasm -target wasm ./mine.go
```

## Run

```bash
goexec 'http.ListenAndServe(`:9000`, http.FileServer(http.Dir(`.`)))'
open http://localhost:9000
```

## Thanks

https://text.baldanders.info/golang/webassembly-with-tinygo/
