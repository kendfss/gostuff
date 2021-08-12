// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"fmt"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"gioui.org/font/gofont"
)

// START OMIT
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

	// UI state.
	var btn widget.Clickable
	var count int

	for e := range w.Events() {
		if e, ok := e.(system.FrameEvent); ok {
			gtx := layout.NewContext(&ops, e)

			for btn.Clicked() {
				count++
			}

			layout.W.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.H5(th, fmt.Sprintf("Number of clicks: %d", count)).Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.Button(th, &btn, "Click me!").Layout(gtx)
					}),
				)
			})

			e.Frame(gtx.Ops)
		}
	}
}

// END OMIT
