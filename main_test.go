package main

import "testing"

func TestEQ(t *testing.T) {
	a := "/some/path"
	b := "more/path"
	m := Manual(a, b)
	p := Path(a, b)
	if m != p {
		t.Error("%s did not match %s", m, p)
	}
}

var aSmall = "/some/path"
var bSmall = "more/path"

var aLong = "/company/service/sharding_info/1/node"
var bLong = "98983ERSO9274KOLSHADFOIHWEF"

var aVeryLong = "/company/service/sharding_info/1/node/company/service/sharding_info/1/node/company/service/sharding_info/1/node/company/service/sharding_info/1/node/company/service/sharding_info/1/node/company/service/sharding_info/1/node/company/service/sharding_info/1/node/company/service/sharding_info/1/node/company/service/sharding_info/1/node/company/service/sharding_info/1/node/company/service/sharding_info/1/node/company/service/sharding_info/1/node"
var bVeryLong = "/company/service/sharding_info/1/nod/company/service/sharding_info/1/nod/company/service/sharding_info/1/nod/company/service/sharding_info/1/node/eee98983ERSO9274KOLSHADFOIHWEF"

var result string

func BenchmarkManualSmall(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = Manual(aSmall, bSmall)
	}
	result = r
}

func BenchmarkPathSmall(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = Path(aSmall, bSmall)
	}
	result = r
}

func BenchmarkManualLarge(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = Manual(aLong, bLong)
	}
	result = r
}

func BenchmarkPathLarge(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = Path(aLong, bLong)
	}
	result = r
}

func BenchmarkManualVeryLarge(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = Manual(aVeryLong, bVeryLong)
	}
	result = r
}

func BenchmarkPathVeryLarge(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = Path(aVeryLong, bVeryLong)
	}
	result = r
}
