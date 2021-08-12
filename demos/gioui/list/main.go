// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"fmt"
	"math/rand"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"gioui.org/font/gofont"
)

func main() {
	// fmt.Println(rand)
	go func() {
		w := app.NewWindow()
		loop(w)
	}()
	app.Main()
}

func loop(w *app.Window) {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops

	programState.list.Axis = layout.Vertical
	programState.factor = 1
	for e := range w.Events() {
		fmt.Println(e)
		if e, ok := e.(system.FrameEvent); ok {
			gtx := layout.NewContext(&ops, e)
			programState.layout(gtx, th)

			e.Frame(gtx.Ops)
		}
	}
}

// START OMIT
type state struct {
	btn1   widget.Clickable
	btn2   widget.Clickable
	btn3   widget.Clickable
	btn4   widget.Clickable
	btn5   widget.Clickable
	btn6   widget.Clickable
	factor int
	shift int
	list   layout.List
}

var programState state

func (s *state) layout(gtx layout.Context, th *material.Theme) {
	// btn := Choice(s.btn2, s.btn1)
	// for s.btn.Clicked() {
	// for s.btn.Clicked() {
		// s.factor *= 10
	// 	s.factor = Choice(s.factor*10, s.factor/10)
	// }
	for s.btn1.Clicked() {
		s.factor *= 10
	}
	for s.btn2.Clicked() {
		s.factor /= 10
	}
	for s.btn3.Clicked() {
		s.factor = Choice(s.factor*10, s.factor/10)
	}
	for s.btn4.Clicked() {
		s.shift += 1
	}
	for s.btn5.Clicked() {
		s.shift -= 1
	}
	for s.btn6.Clicked() {
		// s.shift += Choice(s.center+1, s.center-1)
		s.shift += Choice(1, -1)
	}
	

	layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			// const numRows = 1e6 // A million rows.
			const numRows = 26 
			return s.list.Layout(gtx, numRows, func(gtx layout.Context, i int) layout.Dimensions {
				// txt := fmt.Sprintf("Row #%d", i*s.factor)
				txt := fmt.Sprintf("Row #%d", (i+1)*s.factor+s.shift)
				// txt := fmt.Sprintf("Row #%d", (i+1)*s.factor)
				// txt := fmt.Sprintf("Row #%d", s.center+(i+1)*s.factor)
				return material.Body1(th, txt).Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.S.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Button(th, &s.btn1, "Multiply by 10").Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.S.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Button(th, &s.btn2, "Divide by 10").Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.S.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Button(th, &s.btn3, "Wild Card").Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.S.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Button(th, &s.btn4, "Plus 1").Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.S.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Button(th, &s.btn5, "Minus 1").Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.S.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Button(th, &s.btn6, "Wild Card").Layout(gtx)
			})
		}),
		
	)
}

// END OMIT

func Choice(options...int) int {
	index := rand.Int() % len(options)
	return options[index]
}