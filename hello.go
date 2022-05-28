// +build js,wasm

package main

import (
	"fmt"
	"strings"
	"syscall/js"
)

func say(this js.Value, args []js.Value) interface{} {
	fmt.Println(args)
	ss := []string{}
	for _, jss := range args {
		if s := jsString(jss); s != "" {
			ss = append(ss, s)
		}
	}
	return js.ValueOf("Hello, " + strings.Join(ss, ", "))
}

func jsString(j js.Value) string {
	if j.IsUndefined() || j.IsNull() {
		return ""
	}
	return j.String()
}

func main() {
	ch := make(chan struct{})
	js.Global().Set("say", js.FuncOf(say))
	<-ch // Code must not finish
}
