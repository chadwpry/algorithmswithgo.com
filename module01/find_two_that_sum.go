package module01

import (
	"sort"
)

// FindTwoThatSum will look for two numbers in the provided
// numbers slice that sum up to the sum argument. It will then
// return the indices of each of these numbers.
//
// If there are multiple correct answers, any of them may be
// used. If there are no correct answers, (-1, -1) is returned.
//
// Eg:
//
//	FindTwoThatSum([1,2,3,4], 7) => (2, 3)
//	FindTwoThatSum([0, -1, 1], 0) => (1, 2)
//	FindTwoThatSum([0, 1, 1], 0) => (-1, -1)
//
// Or if we have duplicate answers any of them are okay. All
// of the following are correct.
//
//	FindTwoThatSum([10, 1, 12, 3, 7, 2, 2, 1], 4) => (5, 6)
//	FindTwoThatSum([10, 1, 12, 3, 7, 2, 2, 1], 4) => (1, 3)
//	FindTwoThatSum([10, 1, 12, 3, 7, 2, 2, 1], 4) => (1, 7)
//
// Each number will only be used once, and the original numbers
// list should not be rearranged or altered in any way.
//
// If you want to challenge
// yourself further, try adding different criteria for which
// numbers will be used when there are multiple correct answers.
// Eg:
//  1. Always return the lowest index possible for the first value.
//  2. Always return the index of lowest value possible for the
//     first return value.
//  3. Always return the index of the two values who have a minimal
//     difference. Eg prefer the values 2, 2 over 1, 3 over 0, 4
//     for the sum of 4.
type answer struct {
	x int
	y int
}

func FindTwoThatSum(numbers []int, sum int) (int, int) {
	answers := []answer{}
	numbersLen := len(numbers)

	for i := 0; i < numbersLen-1; i++ {
		for j := i + 1; j < numbersLen; j++ {
			if numbers[i]+numbers[j] == sum {
				// insert answer at the beginning of the slice
				answers = append(answers, answer{x: -1, y: -1})
				copy(answers[1:], answers)
				answers[0] = answer{x: i, y: j}
			}
		}
	}

	if len(answers) > 0 {
		sort.Slice(answers, func(i, j int) bool {
			return answers[i].x-answers[i].y > answers[j].x-answers[j].y
		})

		// fmt.Print(answers)
		// fmt.Printf(" %v, %v", answers[0].x, answers[0].y)
		// fmt.Println()

		return answers[0].x, answers[0].y
	}

	return -1, -1
}
