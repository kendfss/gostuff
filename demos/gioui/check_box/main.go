// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"image/color"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

func main() {
	go func() {
		w := app.NewWindow()
		loop(w)
	}()
	app.Main()
}

func loop(w *app.Window) {
	var ops op.Ops

	for e := range w.Events() {
		if e, ok := e.(system.FrameEvent); ok {
			gtx := layout.NewContext(&ops, e)

			draw(gtx)

			e.Frame(gtx.Ops)
		}
	}
}

// START OMIT
type checkbox struct {
	checked bool
}

var boxes [10]checkbox

func draw(gtx layout.Context) {
	var children []layout.FlexChild
	for i := range boxes {
		box := &boxes[i]
		children = append(children,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(10)).Layout(gtx,
					box.layout,
				)
			}),
		)
	}
	layout.Flex{}.Layout(gtx, children...)
}

func (c *checkbox) layout(gtx layout.Context) layout.Dimensions {
	const size = 50

	// Process events using the key, c.
	for _, e := range gtx.Events(c) {
		if e, ok := e.(pointer.Event); ok {
			if e.Type == pointer.Press {
				c.checked = !c.checked
			}
		}
	}

	st := op.Push(gtx.Ops) // Save operation state.

	// Confine input to the area covered by the checkbox.
	pointer.Rect(image.Rectangle{Max: image.Point{
		X: size,
		Y: size,
	}}).Add(gtx.Ops)
	// Declare the filter with the key, c.
	pointer.InputOp{Tag: c, Types: pointer.Press}.Add(gtx.Ops)

	col := color.RGBA{A: 0xff, R: 0xff} // Red.
	if c.checked {
		col = color.RGBA{A: 0xff, G: 0xff} // Green.
	}

	// Draw checkbox. Red for unchecked, green for checked.
	paint.ColorOp{Color: col}.Add(gtx.Ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{
		X: size,
		Y: size,
	}}}.Add(gtx.Ops)

	st.Pop() // Restore operation state.

	// Specify layout dimensions.
	return layout.Dimensions{
		Size: image.Point{
			X: size, Y: size,
		},
	}
}

// END OMIT
