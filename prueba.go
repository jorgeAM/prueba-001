package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

func main() {
	/*USAR LOS 4 PROCESADORES DE MI PC*/
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	var i uint64
	var sum, x float64
	i = 1
	//x = math.Pow10(15)
	x = 6
	wg.Add(int(x))
	for i <= uint64(x) {
		go func(i uint64) {
			defer wg.Done()
			sum += Sigma2(i)
		}(i)
		i++
	}
	wg.Wait()
	fmt.Println("SIGMA2(", x, ") = ", int(sum))
}

// GetDivisores -> conseguir divisores de un numero
func GetDivisores(a uint64) (divisores []uint64) {
	var i uint64
	var wg sync.WaitGroup
	wg.Add(int(a))
	for i = 1; i <= a; i++ {
		go func(i uint64) {
			defer wg.Done()
			if a%i == 0 {
				divisores = append(divisores, i)
			}
		}(i)
	}
	wg.Wait()
	return
}

// Sigma2 -> 1er función
func Sigma2(n uint64) (res float64) {
	a := GetDivisores(n)
	var wg sync.WaitGroup
	wg.Add(len(a))
	for i := range a {
		go func(i int) {
			defer wg.Done()
			f := float64(a[i])
			res += math.Pow(f, 2)
		}(i)
	}
	wg.Wait()
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
