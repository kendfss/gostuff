package main

import (
	_"sync"
)

type State int

type System []State

func (self System) Contains(state State) bool {
	for _, elem := range self {
		if elem == state {
			return true
		}
	}
	return false
}
func (self System) Intersection(other System) System {
	out := System{}
	for _, elem := range other {
		if self.Contains(elem) {
			out = append(out, elem)
		}
	}
	return out
}
func (self System) Union(other System) System {
	out := System{}
	for _, elem := range other {
		if !self.Contains(elem) {
			out = append(out, elem)
		}
	}
	return out
}
func (self System) Sum(other System) System {
	out := System{}
	for _, elem := range other {
		out = append(out, elem)
	}
	return out	
}

type Transformation [2]State

type Task []Transformation
func (self Task) Args() []State {
	args := []State{}
	for _, t := range self {
		args = append(args, t[0])
	}
	return args
}
func (self Task) Outs() []State {
	outs := []State{}
	for _, t := range self {
		outs = append(outs, t[1])
	}
	return outs
}
func (self Task) Transitions() []Transition {
	args := self.Args()
	trans := []Transition{}
	for i := 0; i < _; i++ {
		
	}
}
func (self Task) Paths() [][]State {
	rack := [][]State{}
	return rack
}

type Transition {
	origin State
	outcome System
}

func main() {
	
}
