package main

import "crypto/rand"

type RandomizedSet struct {
	nums    []int
	indices map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indices[val]; ok {
		return false
	}

	this.indices[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	inx, ok := this.indices[val]
	if !ok {
		return false
	}

	last := len(this.nums) - 1
	this.nums[inx] = this.nums[last]
	this.indices[this.nums[inx]] = inx
	this.nums = this.nums[:last]
	delete(this.indices, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
