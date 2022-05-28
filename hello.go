// +build js,wasm

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"syscall/js"
)

func mine(this js.Value, args []js.Value) interface{} {
	n, _ := strconv.Atoi(args[0].String())
	p := strings.Repeat("0", n)
	c := proofOfWork(100, p)
	js.Global().Get("document").Call("getElementById", "time").Set("innerHTML", "hoge")
	return js.ValueOf(strconv.Itoa(c))
}

func validProof(lastProof int, proof int, prefix string) bool {
	guess := strconv.Itoa(lastProof) + strconv.Itoa(proof)
	b := sha256.Sum256([]byte(guess))
	h := hex.EncodeToString(b[:])
	return strings.HasPrefix(h, prefix)
}

func proofOfWork(lastProof int, prefix string) int {
	p := 0
	for !validProof(lastProof, p, prefix) {
		p++
	}
	return p
}

func main() {
	ch := make(chan struct{})
	fmt.Println("wasm loaded!!")
	js.Global().Set("mineByWasm", js.FuncOf(mine))
	<-ch
}
