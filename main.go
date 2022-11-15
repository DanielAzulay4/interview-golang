package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	obj := Constructor()

	obj.Insert(1)
	obj.Insert(213123)
	obj.Insert(324)
	obj.Insert(5)
	obj.Insert(3)
	obj.Insert(4)

	fmt.Println(obj.GetRandom())
}

type RandomizedSet struct {
	mapping map[int]int
	counter []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		mapping: make(map[int]int, 0),
		counter: make([]int, 0),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.mapping[val]; ok {
		return false
	}

	rs.mapping[val] = len(rs.counter)
	rs.counter = append(rs.counter, val)

	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	if _, ok := rs.mapping[val]; !ok {
		return false
	}

	lastItem := rs.counter[len(rs.counter)-1]
	valIndex := rs.mapping[val]
	rs.mapping[valIndex] = lastItem
	rs.counter = append(rs.counter[:valIndex], rs.counter[valIndex+1:]...)

	return true
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.counter[rand.Intn(len(rs.counter))]
}
