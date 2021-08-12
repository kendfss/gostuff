// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"strconv"
	"time"

	"golang.org/x/exp/rand"
	"golang.org/x/exp/shiny/materialdesign/icons"

	"gonum.org/v1/gonum/spatial/r2"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/pointer"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	
	
	
	
	
	// "golang.org/x/exp/rand"

	"gonum.org/v1/gonum/spatial/barneshut"
	// "gonum.org/v1/gonum/spatial/r2"
)

// distribution tracks useful minimum and maximum information about
// the stars.
type distribution struct {
	min, max         r2.Vec
	maxSpeed         float64
	meanSpeed        float64
	minMass, maxMass float64

	speedSum     float64
	speedSamples int
}

// Update ensures that the distribution contains accurate min/max
// data for the slice of stars provided.
func (d *distribution) Update(stars []*mass) {
	var (
		speedSum     float64
		speedSamples int
	)
	for i, s := range stars {
		speed := distance(s.v, s.d)
		if i == 0 {
			d.minMass = s.m
		}
		if s.d.X < d.min.X {
			d.min.X = s.d.X
		}
		if s.d.Y < d.min.Y {
			d.min.Y = s.d.Y
		}
		if s.d.X > d.max.X {
			d.max.X = s.d.X
		}
		if s.d.Y > d.max.Y {
			d.max.Y = s.d.Y
		}
		if s.m > d.maxMass {
			d.maxMass = s.m
		}
		if s.m < d.minMass {
			d.minMass = s.m
		}
		if speed > d.maxSpeed {
			d.maxSpeed = speed
		}
		speedSamples++
		speedSum += speed
	}
	d.meanSpeed = speedSum / float64(speedSamples)
}

// EnsureSquare adjusts the distribution so that the min and max
// coordinates are the corners of a square (by padding one axis
// equally across the top and bottom). This helps to prevent visual
// distortion during the visualization, though it does not stop it
// completely.
func (d *distribution) EnsureSquare() {
	diff := d.max.Sub(d.min)
	if diff.X > diff.Y {
		padding := (diff.X - diff.Y) / 2
		d.max.Y += padding
		d.min.Y -= padding
	} else if diff.Y > diff.X {
		padding := (diff.Y - diff.X) / 2
		d.max.X += padding
		d.min.X -= padding
	}
}

// String describes the distribution in text form.
func (d distribution) String() string {
	return fmt.Sprintf("distance: (min: %v max: %v), mass: (min: %v, max: %v)", d.min, d.max, d.minMass, d.maxMass)
}

// Scale uses the min/max data within the distribution to compute the
// position, speed, and size of the star.
func (d distribution) Scale(star *mass) Star {
	s := Star{}
	s.X = float32((star.d.X - d.min.X) / (d.max.X - d.min.X))
	s.Y = float32((star.d.Y - d.min.Y) / (d.max.Y - d.min.Y))
	speed := math.Log(distance(star.v, star.d)) / math.Log(d.maxSpeed)
	s.Speed = float32(speed)
	s.Size = unit.Dp(float32(1 + ((star.m / (d.maxMass - d.minMass)) * 10)))
	return s
}

// distance implements the simple two-dimensional euclidean distance function.
func distance(a, b r2.Vec) float64 {
	return math.Sqrt((b.X-a.X)*(b.X-a.X) + (b.Y-a.Y)*(b.Y-a.Y))
}

var PlayIcon = func() *widget.Icon {
	ic, _ := widget.NewIcon(icons.AVPlayArrow)
	return ic
}()
var PauseIcon = func() *widget.Icon {
	ic, _ := widget.NewIcon(icons.AVPause)
	return ic
}()
var ClearIcon = func() *widget.Icon {
	ic, _ := widget.NewIcon(icons.ContentClear)
	return ic
}()

// viewport models a region of a larger space. Offset is the location
// of the upper-left corner of the view within the larger space. size
// is the dimensions of the viewport within the larger space.
type viewport struct {
	offset f32.Point
	size   f32.Point
}

// subview modifies v to describe a smaller region by zooming into the
// space described by v using other.
func (v *viewport) subview(other *viewport) {
	v.offset.X += other.offset.X * v.size.X
	v.offset.Y += other.offset.Y * v.size.Y
	v.size.X *= other.size.X
	v.size.Y *= other.size.Y
}

// ensureSquare returns a copy of the rectangle that has been padded to
// be square by increasing the maximum coordinate.
func ensureSquare(r image.Rectangle) image.Rectangle {
	dx := r.Dx()
	dy := r.Dy()
	if dx > dy {
		r.Max.Y = r.Min.Y + dx
	} else if dy > dx {
		r.Max.X = r.Min.X + dy
	}
	return r
}

var (
	ops         op.Ops
	play, clear widget.Clickable
	playing     = false
	th          = material.NewTheme(gofont.Collection())
	selected    image.Rectangle
	selecting   = false
	view        *viewport
)

func main() {
	th.Palette.Fg, th.Palette.Bg = th.Palette.Bg, th.Palette.Fg
	dist := distribution{}

	seed := time.Now().UnixNano()
	rnd := rand.New(rand.NewSource(uint64(seed)))

	// Make 1000 stars in random locations.
	stars, plane := galaxy(1000, rnd)
	dist.Update(stars)

	desiredSize := unit.Dp(800)
	window := app.NewWindow(
		app.Size(desiredSize, desiredSize),
		app.Title("Seed: "+strconv.Itoa(int(seed))),
	)

	iterateSim := func() {
		if !playing {
			return
		}
		simulate(stars, plane, &dist)
		window.Invalidate()
	}
	for {
		select {
		case ev := <-window.Events():
			switch ev := ev.(type) {
			case system.DestroyEvent:
				if ev.Err != nil {
					log.Fatal(ev.Err)
				}
				return
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, ev)
				paint.Fill(gtx.Ops, th.Palette.Bg)

				layout.Center.Layout(gtx, func(gtx C) D {
					return widget.Border{
						Color: th.Fg,
						Width: unit.Dp(1),
					}.Layout(gtx, func(gtx C) D {
						if gtx.Constraints.Max.X > gtx.Constraints.Max.Y {
							gtx.Constraints.Max.X = gtx.Constraints.Max.Y
						} else {
							gtx.Constraints.Max.Y = gtx.Constraints.Max.X
						}
						gtx.Constraints.Min = gtx.Constraints.Max

						if clear.Clicked() {
							view = nil
						}
						if play.Clicked() {
							playing = !playing
						}

						layoutSelectionLayer(gtx)

						for _, s := range stars {
							dist.Scale(s).Layout(gtx, view)
						}
						layoutControls(gtx)
						return D{Size: gtx.Constraints.Max}
					})
				})

				ev.Frame(gtx.Ops)
				iterateSim()
			}
		}
	}
}

func layoutControls(gtx C) D {
	layout.N.Layout(gtx, func(gtx C) D {
		return material.Body1(th, "Click and drag to zoom in on a region").Layout(gtx)
	})
	layout.S.Layout(gtx, func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx C) D {
			return layout.Flex{
				Spacing: layout.SpaceEvenly,
			}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					var btn material.IconButtonStyle
					if playing {
						btn = material.IconButton(th, &play, PauseIcon)
					} else {
						btn = material.IconButton(th, &play, PlayIcon)
					}
					return btn.Layout(gtx)
				}),
				layout.Rigid(func(gtx C) D {
					if view == nil {
						gtx = gtx.Disabled()
					}
					return material.IconButton(th, &clear, ClearIcon).Layout(gtx)
				}),
			)
		})
	})
	return D{}
}

func layoutSelectionLayer(gtx C) D {
	for _, event := range gtx.Events(&selected) {
		switch event := event.(type) {
		case pointer.Event:
			var intPt image.Point
			intPt.X = int(event.Position.X)
			intPt.Y = int(event.Position.Y)
			switch event.Type {
			case pointer.Press:
				selecting = true
				selected.Min = intPt
				selected.Max = intPt
			case pointer.Drag:
				if intPt.X >= selected.Min.X && intPt.Y >= selected.Min.Y {
					selected.Max = intPt
				} else {
					selected.Min = intPt
				}
				selected = ensureSquare(selected)
			case pointer.Release:
				selecting = false
				newView := &viewport{
					offset: f32.Point{
						X: float32(selected.Min.X) / float32(gtx.Constraints.Max.X),
						Y: float32(selected.Min.Y) / float32(gtx.Constraints.Max.Y),
					},
					size: f32.Point{
						X: float32(selected.Dx()) / float32(gtx.Constraints.Max.X),
						Y: float32(selected.Dy()) / float32(gtx.Constraints.Max.Y),
					},
				}
				if view == nil {
					view = newView
				} else {
					view.subview(newView)
				}
			case pointer.Cancel:
				selecting = false
				selected = image.Rectangle{}
			}
		}
	}
	if selecting {
		paint.FillShape(gtx.Ops, color.NRGBA{R: 255, A: 100}, clip.Rect(selected).Op())
	}
	stack := op.Save(gtx.Ops)
	pointer.Rect(image.Rectangle{Max: gtx.Constraints.Max}).Add(gtx.Ops)
	pointer.CursorNameOp{Name: pointer.CursorCrossHair}.Add(gtx.Ops)
	pointer.InputOp{
		Tag:   &selected,
		Types: pointer.Press | pointer.Release | pointer.Drag,
	}.Add(gtx.Ops)
	stack.Load()

	return D{Size: gtx.Constraints.Max}
}

// Star represents a point of mass rendered within a specific region of a canvas.
type Star struct {
	X, Y  float32
	Speed float32
	Size  unit.Value
}

type (
	C = layout.Context
	D = layout.Dimensions
)

// Layout renders the star into the gtx assuming that it is visible within the
// provided viewport. Stars outside of the viewport will be skipped.
func (s Star) Layout(gtx layout.Context, view *viewport) layout.Dimensions {
	defer op.Save(gtx.Ops).Load()
	px := gtx.Px(s.Size)
	if view != nil {
		if s.X < view.offset.X || s.X > view.offset.X+view.size.X {
			return D{}
		}
		if s.Y < view.offset.Y || s.Y > view.offset.Y+view.size.Y {
			return D{}
		}
		s.X = (s.X - view.offset.X) / view.size.X
		s.Y = (s.Y - view.offset.Y) / view.size.Y
	}
	rr := float32(px / 2)
	x := s.X*float32(gtx.Constraints.Max.X) - rr
	y := s.Y*float32(gtx.Constraints.Max.Y) - rr

	op.Offset(f32.Pt(x, y)).Add(gtx.Ops)
	rect := f32.Rectangle{
		Max: f32.Pt(float32(px), float32(px)),
	}
	fill := color.NRGBA{R: 0xff, G: 128, B: 0xff, A: 50}
	fill.R = 255 - uint8(255*s.Speed)
	fill.B = uint8(255 * s.Speed)
	paint.FillShape(gtx.Ops, fill, clip.UniformRRect(rect, rr).Op(gtx.Ops))
	return D{}
}





























type mass struct {
	d r2.Vec  // position
	v r2.Vec  // velocity
	m float64 // mass
}

func (m *mass) Coord2() r2.Vec { return m.d }
func (m *mass) Mass() float64  { return m.m }
func (m *mass) move(f r2.Vec) {
	// F = ma
	f.X /= m.m
	f.Y /= m.m
	m.v = m.v.Add(f)

	// Update position.
	m.d = m.d.Add(m.v)
}

func galaxy(numStars int, rnd *rand.Rand) ([]*mass, barneshut.Plane) {

	// Make 50 stars in random locations and velocities.
	stars := make([]*mass, numStars)
	p := make([]barneshut.Particle2, len(stars))
	for i := range stars {
		s := &mass{
			d: r2.Vec{
				X: 100*rnd.Float64() - 50,
				Y: 100*rnd.Float64() - 50,
			},
			m: rnd.Float64(),
		}
		// Aim at the ground and miss.
		s.d = s.d.Scale(-1).Add(r2.Vec{
			X: 10 * rnd.NormFloat64(),
			Y: 10 * rnd.NormFloat64(),
		})

		stars[i] = s
		p[i] = s
	}
	// Make a plane to calculate approximate forces
	plane := barneshut.Plane{Particles: p}

	return stars, plane
}

func simulate(stars []*mass, plane barneshut.Plane, dist *distribution) {
	vectors := make([]r2.Vec, len(stars))
	// Build the data structure. For small systems
	// this step may be omitted and ForceOn will
	// perform the naive quadratic calculation
	// without building the data structure.
	err := plane.Reset()
	if err != nil {
		log.Fatal(err)
	}

	// Calculate the force vectors using the theta
	// parameter.
	const theta = 0.1
	// and an imaginary gravitational constant.
	const G = 10
	for j, s := range stars {
		vectors[j] = plane.ForceOn(s, theta, barneshut.Gravity2).Scale(G)
	}

	// Update positions.
	for j, s := range stars {
		s.move(vectors[j])
	}

	// Recompute the distribution of stars
	dist.Update(stars)
	dist.EnsureSquare()
}
