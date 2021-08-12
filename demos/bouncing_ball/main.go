// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image/color"
	"math"
	"time"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"gioui.org/font/gofont"
)

func main() {
	go func() {
		w := app.NewWindow()
		loop(w)
	}()
	app.Main()
}

func loop(w *app.Window) {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops

	for e := range w.Events() {
		if e, ok := e.(system.FrameEvent); ok {
			gtx := layout.NewContext(&ops, e)

			programState.layout(gtx, th)

			e.Frame(gtx.Ops)
		}
	}
}

// START OMIT
type state struct {
	btn         widget.Clickable
	ballVisible bool
	time        time.Time
}

var programState state

func (s *state) layout(gtx layout.Context, th *material.Theme) {
	for s.btn.Clicked() {
		s.ballVisible = !s.ballVisible
		s.time = gtx.Now
	}

	if s.ballVisible {
		s.drawBall(gtx)
	}
	layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		label := "Throw ball"
		if s.ballVisible {
			label = "Hide ball"
		}
		return material.Button(th, &s.btn, label).Layout(gtx)
	})
}

func (s *state) drawBall(gtx layout.Context) {
	defer op.Push(gtx.Ops).Pop()
	const size = 50

	// Animate bounce.
	const (
		v0 = 50.0
		g  = -9.81
	)
	now := gtx.Now
	t := now.Sub(s.time).Seconds()
	t *= 10
	t1 := -v0 / (.5 * g)
	t = math.Mod(t, t1)
	y := v0*t + .5*g*t*t

	// Position.
	bottom := float32(gtx.Constraints.Max.Y)
	op.Offset(f32.Point{X: 20, Y: bottom - size - float32(y)}).Add(gtx.Ops)

	// Draw.
	paint.ColorOp{Color: color.RGBA{A: 0xff, G: 0xff}}.Add(gtx.Ops)
	clip.Rect{
		Rect: f32.Rectangle{
			Max: f32.Point{
				X: size, 
				Y: size,
			},
		},
		NE: size * .5, 
		NW: size * .5, 
		SE: size * .5, 
		SW: size * .5,
	}.Op(gtx.Ops).Add(gtx.Ops)
	paint.PaintOp{
		Rect: f32.Rectangle{Max: f32.Point{X: size, Y: size}},
	}.Add(gtx.Ops)

	// Request immediate redraw.
	op.InvalidateOp{}.Add(gtx.Ops)
}

// END OMIT
