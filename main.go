package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	args := os.Args[1:]
	var num int
	if len(args) != 1 {
		fmt.Println("Error: exactly one argument required")
		os.Exit(1)
	} else {
		var err error
		num, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: argument must be a number")
			os.Exit(1)
		}
	}
	factors, err := factor(num)
	if err != nil {
		if err == ErrPrime {
			fmt.Println("Prime!")
			os.Exit(0)
		} else {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Factors:")
		for _, f := range factors {
			fmt.Printf("%d * %d\n", f.a, f.b)
		}
	}
}

type factors struct {
	a int
	b int
}

var ErrPrime = errors.New("number is prime")

func factor(num int) ([]factors, error) {
	var f []factors
	found := 0
	for a := 2; a < (num / 2); a++ {
		fmt.Printf("\r%d (%v%%, found: %d) ...", a, ((a * 100) / (num / 2)), found)
		b := float64(num) / float64(a)
		bR := int(b)
		if float64(bR) == b {
			found++
			f = append(f, factors{a, bR})
		}
	}
	fmt.Printf("complete.\n")
	if len(f) == 0 {
		return nil, ErrPrime
	}
	var deduplicated []factors
	for _, set := range f {
		if !slices.Contains(deduplicated, set) && !slices.Contains(deduplicated, factors{set.b, set.a}) {
			deduplicated = append(deduplicated, factors{set.b, set.a})
		}
	}
	return deduplicated, nil
}

func primes(max int) []int {
	var primes []int
	for i := 2; i < max; i++ {
		_, err := factor(i)
		if err == ErrPrime {
			fmt.Println(i)
			primes = append(primes, i)
		}
	}
	return primes
}
