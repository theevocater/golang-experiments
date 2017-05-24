package main

import "path"
import _ "net/http/pprof"

func Manual(a string, b string) string {
	return a + "/" + b
}

func Path(a string, b string) string {
	return path.Join(a, b)
}
