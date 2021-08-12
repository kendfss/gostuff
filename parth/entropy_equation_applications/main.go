// The ENTROPY EQUATION and its Applications _ Thermodynamics and Microstates EXPLAINED-3pl_pJQDW4k.mp4
package main


import (
	"log"
)

type Coordinates []int
type EnergyLevel int
type Particle struct {
	Coords Coordinates
	Level EnergyLevel
}

type State []Particle
func (self State) Energy() (val EnergyLevel) {
	for _, p := range self {
		val += p.Level
	}
	return
}

func Rack(length int) (rack State) {
	for i := 0; i < length; i++ {
		rack = append(rack, Particle{Coordinates{i}, 0})
	}
	return
}
func Racks(particles, levels int) []State {
	
}
func ways(particles, levels, level int) (ctr int) {
	// for l:=1; l<levels+1; l++ {
	// 	for p:=0; p<particles; p++ {
			
	// 	}
	// }
	for w:=0; w<particles*levels; w++ {
		way := 
	}
}


func main() {
	rack := Rack(4)
	rack[0].Level = 4
	rack[1].Level = 2
	rack[2].Level = 2
	rack[3].Level = 6
	log.Println(rack)
	log.Println(rack.Energy()) // 2.1 a
}
