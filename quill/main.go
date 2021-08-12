package main

import (
	"image/color"
	"image"
)

func main() {
	
}





type Canvas struct {
	
}
type Location []int
type Dimensions []int
type Space interface { // controls the relocation behaviour of the a quill, like a topology would
	Move(Location, Dimensions) Location // given a new location and an "image"'s dimensions, return a new location
	At(Location) color.Color
	Set(Location, color.Color)
}
type Quill struct {
	topum Space 
}
type Morph interface { // a drawable object
	Coords() Location // the "centre" of a morph 
	Net() []Location // the set of points to be marked, relative to the center
	Colour() color.Color
}




func (self Quill) Draw(morph Morph, canvas image.Image) {
	
}
func (self *Quill) Recolour(colour color.Color) {
	
}

