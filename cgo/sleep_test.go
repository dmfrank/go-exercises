package main

import "testing"

func BenchmarkCGO(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CGO()
	}
}

func BenchmarkGO(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GO()
	}
}

/*bash-3.2$ go test -bench=.
PASS
BenchmarkCGO-4	       1	2007446430 ns/op
BenchmarkGO-4 	       1	2006214842 ns/op
ok  	_/Users/dmitryfrank/Go/go-exercises/cgo	4.039s
*/
