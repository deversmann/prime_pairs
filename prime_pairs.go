package main

import (
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	top := 9
	if len(os.Args) > 1 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil || i < 2 || i > 15 {
			fmt.Println("Single optional argument needs to be an integer from 2 to 15.  Defaulting to 9.")
		} else {
			top = i
		}
	}
	odds := make([]int, (top/2 + top%2))
	evens := make([]int, (top / 2))
	for i := range odds {
		odds[i] = 2*i + 1
	}
	for i := range evens {
		evens[i] = 2 * (i + 1)
	}

	oddlist := permutations(odds)
	evenlist := permutations(evens)
	solutions := [][]int{}

	for _, odd := range oddlist {
		for _, even := range evenlist {
			candidate := interleave(odd, even)
			if checkPrimePairs(candidate) {
				solutions = append(solutions, candidate)
				if len(odd) == len(even) {
					mirror := make([]int, len(candidate))
					for i, j := 0, len(candidate)-1; i < j; i, j = i+1, j-1 {
						mirror[i], mirror[j] = candidate[j], candidate[i]
					}
					solutions = append(solutions, mirror)
				}
			}
		}
	}

	endCompute := time.Now()

	sort.Slice(solutions, func(i, j int) bool {
		for k := 0; k < len(solutions[i]); k++ {
			if solutions[i][k] > solutions[j][k] {
				return false
			}
			if solutions[i][k] < solutions[j][k] {
				return true
			}
		}
		return true
	})

	fmt.Printf("There are %v solutions for the first %v positive integers:\n", len(solutions), top)
	for _, solution := range solutions {
		fmt.Println(fmt.Sprint(solution))
	}
	fmt.Printf("End of the %v solutions for the first %v positive integers\n", len(solutions), top)
	complete := time.Now()
	fmt.Printf("Compute took %v not including the additional %v to sort and display\n", endCompute.Sub(start), complete.Sub(endCompute))
}

func permutations(a []int) (retval [][]int) {
	var heapPermutation func([]int, int)
	heapPermutation = func(a []int, n int) {
		if n == 1 {
			permutation := make([]int, len(a))
			copy(permutation, a)
			retval = append(retval, permutation)
		} else {
			for i := 0; i < n; i++ {
				heapPermutation(a, n-1)
				if n%2 == 1 {
					temp := a[i]
					a[i] = a[n-1]
					a[n-1] = temp
				} else {
					temp := a[0]
					a[0] = a[n-1]
					a[n-1] = temp
				}
			}
		}
	}
	heapPermutation(a, len(a))
	return
}

func interleave(a []int, b []int) []int {
	retval := make([]int, len(a)+len(b))
	j := 0
	for i := 0; i < max(len(a), len(b)); i++ {
		if i < len(a) {
			retval[j] = a[i]
			j++
		}
		if i < len(b) {
			retval[j] = b[i]
			j++
		}
	}
	return retval
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func checkPrimePairs(a []int) bool {
	if len(a) < 2 {
		return false
	}
	for i := 1; i < len(a); i++ {
		x := big.NewInt(int64(a[i] + a[i-1]))
		if !x.ProbablyPrime(1) {
			return false
		}
	}
	return true
}
