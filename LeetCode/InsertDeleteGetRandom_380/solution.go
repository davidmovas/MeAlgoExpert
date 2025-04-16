package InsertDeleteGetRandom_380

import "math/rand/v2"

func v1() []any {
	var output []any

	set := Constructor()
	output = append(output, set.Insert(1))
	output = append(output, set.Remove(2))
	output = append(output, set.Insert(2))
	output = append(output, set.GetRandom())
	output = append(output, set.Remove(1))
	output = append(output, set.Insert(2))
	output = append(output, set.GetRandom())

	return output
}

type RandomizedSet struct {
	set      map[int]int
	elements []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		set: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.set[val]; exists {
		return false
	}

	this.elements = append(this.elements, val)
	this.set[val] = len(this.elements) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, exists := this.set[val]
	if !exists {
		return false
	}

	lastIdx := len(this.elements) - 1
	this.elements[idx], this.elements[lastIdx] = this.elements[lastIdx], this.elements[idx]
	this.set[this.elements[lastIdx]] = idx

	this.elements = this.elements[:lastIdx]
	delete(this.set, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.elements[rand.IntN(len(this.elements))]
}
