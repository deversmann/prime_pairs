package main

import (
	"fmt"
	"math/big"
	"sort"
)

func main() {
	top := 9
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

	sort.SliceStable(solutions, func(i, j int) bool {
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

	fmt.Printf("There are %v solutions:\n", len(solutions))
	for _, solution := range solutions {
		fmt.Println(fmt.Sprint(solution))
	}
}

func permutations(a []int) (retval [][]int) {
	storeperm := func(perm []int) {
		retval = append(retval, perm)
	}
	heapPermutation(a, len(a), len(a), storeperm)
	return
}

func heapPermutation(a []int, size int, n int, store func(a []int)) {
	if size == 1 {
		permutation := make([]int, n)
		copy(permutation, a)
		store(permutation)
	}
	for i := 0; i < size; i++ {
		heapPermutation(a, size-1, n, store)
		if size%2 == 1 {
			temp := a[0]
			a[0] = a[size-1]
			a[size-1] = temp
		} else {
			temp := a[i]
			a[i] = a[size-1]
			a[size-1] = temp
		}
	}
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
