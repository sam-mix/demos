package one

import "testing"

func BenchmarkSum1(b *testing.B) {
	var sum int
	for i := 0; i < b.N; i++ {
		CaseSumOne(&sum)
	}
}

func BenchmarkSum2(b *testing.B) {
	var sum int
	for i := 0; i < b.N; i++ {
		CaseSumTwo(&sum)
	}
}

func BenchmarkSum3(b *testing.B) {
	var sum int
	for i := 0; i < b.N; i++ {
		CaseSumThree(&sum)
	}
}
