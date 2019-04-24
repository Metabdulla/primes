package primes

import (
	"fmt"
	"math"
	"math/big"
)



func FindPrimes(N int) []int {
	var prime = make([]bool, N)
	var result []int
	var i, j int
	for i = 2; i < N; i++ {
		if i%2 != 0 {
			prime[i] = true
		}
	}
	for i = 3; float64(i) <= math.Sqrt(float64(N)); i++ {
		if prime[i] {
			for j = i + i; j < N; j += i {
				prime[j] = false
			}
		}
	}
	result = append(result, 2)
	for i, v := range prime {
		if v == true {
			result = append(result, i)
		}
	}
	return result
}

func IsPrimeFast(i int, primes []int) bool {
	// Yes, this is a dumb way to write this code,
	// but calling Sqrt repeatedly in this way demonstrates
	// the benefit of using a direct SQRT instruction on systems
	// that have one, whereas the obvious loop seems not to
	// demonstrate such a benefit.
	somePrimes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113}

	if len(primes) < len(somePrimes) {
		primes = somePrimes
	}

	k := 0
	for j := primes[k]; float64(j) <= math.Sqrt(float64(i)); {
		if i%j == 0 {
			//fmt.Println(j)
			return false
		}
		k++
		if k < len(primes) {
			j = primes[k]
		} else {
			j = j + 2
		}
	}
	return true
}

func IsPrime(i int) bool {
	// Yes, this is a dumb way to write this code,
	// but calling Sqrt repeatedly in this way demonstrates
	// the benefit of using a direct SQRT instruction on systems
	// that have one, whereas the obvious loop seems not to
	// demonstrate such a benefit.
	for j := 2; float64(j) <= math.Sqrt(float64(i)); j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}

func SecPrime(x *big.Int, n int) bool {
	for j := 2; j <= n; j++ {
		y := big.NewInt(int64(j))
		z := big.NewInt(1)
		if z.Mod(x, y).Cmp(big.NewInt(0)) == 0 {
			fmt.Println(j)
			return false
		} else if j%10 == 0 {
			fmt.Println(j)
		}
	}
	return true
}
