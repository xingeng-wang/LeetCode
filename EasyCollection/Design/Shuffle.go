package Design

import (
	"math/rand"
	"time"
)

// Shuffle a set of numbers without duplicates.
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/98/design/670/
type Solution struct {
	original []int
	current  []int
}

func Constructor(nums []int) Solution {
	o := make([]int, len(nums))
	c := make([]int, len(nums))
	copy(o, nums)
	copy(c, nums)
	return Solution{
		original: o,
		current:  c,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	copy(this.current, this.original)
	return this.current
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := len(this.current); i > 0; i-- {
		randIndex := r.Intn(i)
		this.current[i-1], this.current[randIndex] = this.current[randIndex], this.current[i-1]
	}
	return this.current
}
