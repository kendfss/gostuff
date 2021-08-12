// https://www.golangprograms.com/golang-chart-package.html

import (
	"fmt"
	"image"
	"log"
	"os"
	"github.com/EdlinOrg/prominentcolor"
)



package main

import (
	"os"

	"github.com/wcharczuk/go-chart"
)

func main() {
	graph := chart.BarChart{
		Title: "Test Bar Chart",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		Bars: []chart.Value{
			{Value: 5.25, Label: "Blue"},
			{Value: 4.88, Label: "Green"},
			{Value: 4.74, Label: "Gray"},
			{Value: 3.22, Label: "Orange"},
			{Value: 3, Label: "Test"},
			{Value: 2.27, Label: "??"},
			{Value: 1, Label: "!!"},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}

// Example - Scatter Chart

package main

import (
	"os"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

func main() {
	viridisByY := func(xr, yr chart.Range, index int, x, y float64) drawing.Color {
		return chart.Viridis(y, yr.GetMin(), yr.GetMax())
	}

	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					StrokeWidth:      chart.Disabled,
					DotWidth:         5,
					DotColorProvider: viridisByY,
				},
				XValues: chart.Seq{Sequence: chart.NewLinearSequence().WithStart(0).WithEnd(127)}.Values(),
				YValues: chart.Seq{Sequence: chart.NewRandomSequence().WithLen(128).WithMin(0).WithMax(1024)}.Values(),
			},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}

