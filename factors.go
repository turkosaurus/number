package main

func factor(num int, result *[]int) {
	for a := 2; a*a <= num; a++ {
		for num%a == 0 {
			*result = append(*result, a)
			num /= a
		}
	}
	if num > 1 {
		*result = append(*result, num)
	}
}

func prime(num int) bool {
	for a := 2; a*a <= num; a++ {
		if num%a == 0 {
			return false
		}
	}
	return true
}
