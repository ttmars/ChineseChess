package chess

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

// Pitch 选中框
type Pitch struct {
	widget.BaseWidget
	Color           color.Color
	LineStrokeWidth float32
	Length          float32 // 长度
	X, Y            int     // 当前坐标
}

func NewPitch(color color.Color, x, y int, length, lineStrokeWidth float32) *Pitch {
	pitch := &Pitch{
		Color:           color,
		X:               x,
		Y:               y,
		Length:          length,
		LineStrokeWidth: lineStrokeWidth,
	}

	pitch.ExtendBaseWidget(pitch)
	return pitch
}

func (c *Pitch) SetPosition(x int, y int) {
	c.Move(fyne.Position{
		X: BorderSize + float32(x)*CellSize - PieceSize/2,
		Y: BorderSize + float32(y)*CellSize - PieceSize/2,
	})
	c.X = x
	c.Y = y
	c.Show()
}

func (c *Pitch) CreateRenderer() fyne.WidgetRenderer {
	var lines [8]*canvas.Line
	blank := PieceSize - 2*c.Length
	lines[0] = NewLine(0, 0, c.Length, 0, c.LineStrokeWidth, Blue)
	lines[2] = NewLine(0, PieceSize, c.Length, PieceSize, c.LineStrokeWidth, Blue)

	lines[1] = NewLine(c.Length+blank, 0, PieceSize, 0, c.LineStrokeWidth, Blue)
	lines[3] = NewLine(c.Length+blank, PieceSize, PieceSize, PieceSize, c.LineStrokeWidth, Blue)

	lines[4] = NewLine(0, 0, 0, c.Length, c.LineStrokeWidth, Blue)
	lines[6] = NewLine(PieceSize, 0, PieceSize, c.Length, c.LineStrokeWidth, Blue)

	lines[5] = NewLine(0, c.Length+blank, 0, PieceSize, c.LineStrokeWidth, Blue)
	lines[7] = NewLine(PieceSize, c.Length+blank, PieceSize, PieceSize, c.LineStrokeWidth, Blue)

	return &pitchPieceRenderer{Lines: lines, objects: []fyne.CanvasObject{lines[0], lines[1], lines[2], lines[3], lines[4], lines[5], lines[6], lines[7]}}
}

type pitchPieceRenderer struct {
	Lines   [8]*canvas.Line
	objects []fyne.CanvasObject
}

func (r *pitchPieceRenderer) Layout(size fyne.Size) {
	//for _, line := range r.Lines {
	//	line.Resize(size)
	//}
	//r.Lines[1].Move(fyne.NewPos(50, 50))
}

func (r *pitchPieceRenderer) MinSize() fyne.Size {
	return fyne.NewSize(100, 100)
}

func (r *pitchPieceRenderer) Refresh() {
	for _, line := range r.Lines {
		line.Refresh()
	}
}

func (r *pitchPieceRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *pitchPieceRenderer) Destroy() {}
