package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(SIGMA2(6))
}

// GetDivisores -> conseguir divisores de un numero
func GetDivisores(a uint64) (divisores []uint64) {
	var i uint64
	for i = 1; i <= a; i++ {
		if a%i == 0 {
			divisores = append(divisores, i)
		}
	}
	return
}

// Sigma2 -> 1er función
func Sigma2(n uint64) (res float64) {
	a := GetDivisores(n)
	for i := range a {
		f := float64(a[i])
		res += math.Pow(f, 2)
	}
	return
}

// SIGMA2 -> 2da función
func SIGMA2(n uint64) (sum float64) {
	var i uint64
	i = 1
	for i <= n {
		sum += Sigma2(i)
		i++
	}
	return
}
