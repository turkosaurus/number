package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// func BenchmarkPrime(b *testing.B) {
// 	for i := 2; i < b.N; i++ {
// 		b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
// 			factors, err := factor(i)
// 			if err != ErrPrime {
// 				require.NoError(b, err)
// 			}
// 			b.Logf("Factors: %v", factors)
// 		})
// 	}
// }

func TestPrimes(t *testing.T) {
	results := primes(100000)
	require.NotEmpty(t, results)
	for _, r := range results {
		t.Logf("%d", r)
	}
}
