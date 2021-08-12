// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
    "image"
    "image/color"
    "log"
    "os"
    "time"
    "math/rand"

    "gioui.org/app"       // app contains Window handling.
    "gioui.org/io/key"    // key is used for keyboard events.
    "gioui.org/io/system" // system is used for system events (e.g. closing the window).
    "gioui.org/layout"    // layout is used for layouting widgets.
    "gioui.org/op"        // op is used for recording different operations.
    "gioui.org/unit"      // unit is used to define pixel-independent sizes
    "gioui.org/f32"        // f32 is used for shape calculations.
    "gioui.org/io/pointer" // system is used for system events (e.g. closing the window).
    // "gioui.org/layout"     // layout is used for layouting widgets.
    // "gioui.org/op"         // op is used for recording different operations.
    // "gioui.org/color"
    "gioui.org/op/clip"    // clip is used to draw the cell shape.
    "gioui.org/op/paint"   // paint is used to paint the cells.
)

// ###################################################
// #·······················MAIN······················#
// ###################################################

var (
    // cellSizePx is the cell size in pixels.
    cellSize = unit.Dp(2)
    // boardSize is the count of cells in a particular dimension.
    // boardSize = image.Pt(50, 50)
    boardSize = image.Pt(100, 100)
)

func main() {
    // The ui loop is separated from the application window creation
    // such that it can be used for testing.
    ui := NewUI()

    windowWidth := cellSize.Scale(float32(boardSize.X + 2))
    windowHeight := cellSize.Scale(float32(boardSize.Y + 2))
    // This creates a new application window and starts the UI.
    go func() {
        w := app.NewWindow(
            app.Title("Game of Life"),
            app.Size(windowWidth, windowHeight),
        )
        if err := ui.Run(w); err != nil {
            log.Println(err)
            os.Exit(1)
        }
        os.Exit(0)
    }()

    // This starts Gio main.
    app.Main()
}

// UI holds all of the application state.
type UI struct {
    // Board handles all game-of-life logic.
    Board *Board
}

// NewUI creates a new UI using the Go Fonts.
func NewUI() *UI {
    // We start with a new random board.
    board := NewBoard(boardSize)
    board.Randomize()

    return &UI{
        Board: board,
    }
}

// Run handles window events and renders the application.
func (ui *UI) Run(w *app.Window) error {
    var ops op.Ops

    // Update the board 3 times per second.
    advanceBoard := time.NewTicker(time.Second / 3)
    defer advanceBoard.Stop()

    // listen for events happening on the window.
    for {
        select {
        case e := <-w.Events():
            // detect the type of the event.
            switch e := e.(type) {
            // this is sent when the application should re-render.
            case system.FrameEvent:
                // gtx is used to pass around rendering and event information.
                gtx := layout.NewContext(&ops, e)
                // render and handle UI.
                ui.Layout(gtx)
                // render and handle the operations from the UI.
                e.Frame(gtx.Ops)

            // handle a global key press.
            case key.Event:
                switch e.Name {
                // when we click escape, let's close the window.
                case key.NameEscape:
                    return nil
                }

            // this is sent when the application is closed.
            case system.DestroyEvent:
                return e.Err
            }

        case <-advanceBoard.C:
            ui.Board.Advance()
            w.Invalidate()
        }
    }
}

// Layout displays the main program layout.
func (ui *UI) Layout(gtx layout.Context) layout.Dimensions {
    return layout.Center.Layout(gtx,
        BoardStyle{
            CellSizePx: gtx.Px(cellSize),
            Board:      ui.Board,
        }.Layout,
    )
}


// ###################################################
// #······················BOARD······················#
// ###################################################


// Board implements game of life logic.
type Board struct {
    // Size is the count of cells in a particular dimension.
    Size image.Point
    // Cells contains the alive or dead cells.
    Cells []byte

    // buffer is used to avoid reallocating a new cells
    // slice for every update.
    buffer []byte
}

// NewBoard returns a new game of life with the defined size.
func NewBoard(size image.Point) *Board {
    return &Board{
        Size:   size,
        Cells:  make([]byte, size.X*size.Y),
        buffer: make([]byte, size.X*size.Y),
    }
}

// Randomize randomizes each cell state.
func (b *Board) Randomize() {
    rand.Read(b.Cells)
    for i, v := range b.Cells {
        if v < 0x30 {
            b.Cells[i] = 1
        } else {
            b.Cells[i] = 0
        }
    }
}

// Pt returns the coordinate given a index in b.Cells.
func (b *Board) Pt(i int) image.Point {
    x, y := i%b.Size.X, i/b.Size.Y
    return image.Point{X: x, Y: y}
}

// At returns the b.Cells index, given a wrapped coordinate.
func (b *Board) At(c image.Point) int {
    if c.X < 0 {
        c.X += b.Size.X
    }
    if c.X >= b.Size.X {
        c.X -= b.Size.X
    }
    if c.Y < 0 {
        c.Y += b.Size.Y
    }
    if c.Y >= b.Size.Y {
        c.Y -= b.Size.Y
    }
    return b.Size.Y*c.Y + c.X
}

// SetWithoutWrap sets a cell to alive.
func (b *Board) SetWithoutWrap(c image.Point) {
    if !c.In(image.Rectangle{Max: b.Size}) {
        return
    }

    b.Cells[b.At(c)] = 1
}

// Advance advances the board state by 1.
func (b *Board) Advance() {
    next, cur := b.buffer, b.Cells
    defer func() { b.Cells, b.buffer = next, cur }()

    for i := range next {
        next[i] = 0
    }

    for y := 0; y < b.Size.Y; y++ {
        for x := 0; x < b.Size.X; x++ {
            var t byte
            t += cur[b.At(image.Pt(x-1, y-1))]
            t += cur[b.At(image.Pt(x+0, y-1))]
            t += cur[b.At(image.Pt(x+1, y-1))]
            t += cur[b.At(image.Pt(x-1, y+0))]
            t += cur[b.At(image.Pt(x+1, y+0))]
            t += cur[b.At(image.Pt(x-1, y+1))]
            t += cur[b.At(image.Pt(x+0, y+1))]
            t += cur[b.At(image.Pt(x+1, y+1))]

            // Any live cell with fewer than two live neighbours dies, as if by underpopulation.
            // Any live cell with two or three live neighbours lives on to the next generation.
            // Any live cell with more than three live neighbours dies, as if by overpopulation.
            // Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

            p := b.At(image.Pt(x, y))
            switch {
            case t < 2:
                t = 0
            case t == 2:
                t = cur[p]
            case t == 3:
                t = 1
            case t > 3:
                t = 0
            }

            next[p] = t
        }
    }
}





// ###################################################
// #······················STYLE······················#
// ###################################################


// BoardStyle draws Board with rectangles.
type BoardStyle struct {
    CellSizePx int
    *Board
}

// Layout draws the Board and accepts input for adding alive cells.
func (board BoardStyle) Layout(gtx layout.Context) layout.Dimensions {
    defer op.Save(gtx.Ops).Load()

    // Calculate the board size based on the cell size in pixels.
    size := board.Size.Mul(board.CellSizePx)
    gtx.Constraints = layout.Exact(size)

    // Handle any input from a pointer.
    for _, ev := range gtx.Events(board.Board) {
        if ev, ok := ev.(pointer.Event); ok {
            p := image.Pt(int(ev.Position.X), int(ev.Position.Y))
            // Calculate the board coordinate given a cursor position.
            p = p.Div(board.CellSizePx)
            board.SetWithoutWrap(p)
        }
    }
    // Register to listen for pointer Drag events.
    pointer.Rect(image.Rectangle{Max: size}).Add(gtx.Ops)
    pointer.InputOp{Tag: board.Board, Types: pointer.Drag}.Add(gtx.Ops)

    cellSize := float32(board.CellSizePx)

    // Draw a shape for each alive cell.
    var p clip.Path
    p.Begin(gtx.Ops)
    for i, v := range board.Cells {
        if v == 0 {
            continue
        }

        c := layout.FPt(board.Pt(i).Mul(board.CellSizePx))
        p.MoveTo(f32.Pt(c.X, c.Y))
        p.LineTo(f32.Pt(c.X+cellSize, c.Y))
        p.LineTo(f32.Pt(c.X+cellSize, c.Y+cellSize))
        p.LineTo(f32.Pt(c.X, c.Y+cellSize))
        p.Close()
    }
    clip.Outline{Path: p.End()}.Op().Add(gtx.Ops)

    // Paint the shape with a black color.
    paint.ColorOp{Color: color.NRGBA{A: 0xFF}}.Add(gtx.Ops)
    paint.PaintOp{}.Add(gtx.Ops)

    return layout.Dimensions{Size: size}
}
