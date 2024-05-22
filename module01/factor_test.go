package module01

import (
	"fmt"
	"sort"
	"testing"
)

func TestFactor(t *testing.T) {
	tenPrimes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
	tests := []struct {
		primes []int
		number int
		want   []int
	}{
		{tenPrimes, 30, []int{2, 3, 5}},
		{tenPrimes, 28, []int{2, 2, 7}},
		{tenPrimes, 720, []int{2, 2, 2, 2, 3, 3, 5}},
		{[]int{2, 5, 7}, 30, []int{2, 5, 3}},
		{[]int{3, 5, 7}, 720, []int{3, 3, 5, 16}},
		{[]int{}, 4, []int{4}},
		{[]int{2, 4, 7}, 30, []int{2, 15}}, // NON PRIME (4) TEST
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v with primes %v", tc.number, tc.primes), func(t *testing.T) {
			got := Factor(tc.primes, tc.number)
			fmt.Println(got)
			sort.Ints(got)
			sort.Ints(tc.want)
			if err := intSlicesEqual(got, tc.want); err != nil {
				t.Fatalf("Factor() sorted = %v; want %v; err = %v", got, tc.want, err)
			}
		})
	}
}
