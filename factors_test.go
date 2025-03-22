package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPfactor(t *testing.T) {
	results := []int{}
	factor(12, &results)
	require.Equal(t, []int{2, 2, 3}, results)
	t.Logf("12 = %v", results)
}

func TestPrime(t *testing.T) {
	four := prime(4)
	require.False(t, four)

	seven := prime(7)
	require.True(t, seven)

}

func TestListPrime(t *testing.T) {
	for i := 0; i < 50; i++ {
		if prime(i) {
			t.Logf("%v\n", i)
		}
	}
}
