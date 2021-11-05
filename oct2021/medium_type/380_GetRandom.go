package medium_type

import (
	"math/rand"
)

type RandomizedSet struct {
	M map[int]int
	S []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{M: make(map[int]int)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.M[val]; ok {
		return false
	}

	this.S = append(this.S, val)
	this.M[val] = len(this.S) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	rmvIdx, ok := this.M[val]
	if !ok {
		return false
	}

	delete(this.M, val)
	if rmvIdx != len(this.S)-1 {
		this.M[this.S[len(this.S)-1]] = rmvIdx
		this.S[rmvIdx] = this.S[len(this.S)-1]
	}
	this.S = this.S[:len(this.S)-1]

	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	idx := rand.Int() % len(this.S)
	return this.S[idx]
}
