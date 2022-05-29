# Golang Wasm Practice

## Build

Go
```bash
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./
GOOS=js GOARCH=wasm go build -o hello.wasm -trimpath
```

[Tinygo](https://tinygo.org/)
```bash
tinygo build -o hello.wasm -target wasm ./hello.go
```

## Run

```bash
go run main.go
```
