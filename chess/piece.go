package chess

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

// Piece 自定义棋子组件
type Piece struct {
	widget.BaseWidget
	Name      string
	FontColor color.Color
	W, H      float32 // 大小
	X, Y      int     // 当前坐标
}

// NewPiece 初始化一个棋子
func NewPiece(name string, fontColor color.Color, x, y int, w, h float32) *Piece {
	piece := &Piece{
		Name:      name,
		FontColor: fontColor,
		X:         x,
		Y:         y,
		W:         w,
		H:         h,
	}

	piece.ExtendBaseWidget(piece)
	piece.Resize(fyne.NewSize(piece.W, piece.H))
	piece.SetPosition(x, y, false)
	return piece
}

// SetPosition 移动棋子
func (c *Piece) SetPosition(x int, y int, addAnimation bool) {
	if addAnimation {
		canvas.NewPositionAnimation(fyne.Position{X: float32(c.X) * Cell, Y: float32(c.Y) * Cell}, fyne.Position{X: float32(x) * Cell, Y: float32(y) * Cell}, time.Millisecond*500, c.Move).Start()
	} else {
		c.Move(fyne.Position{
			X: float32(x) * Cell,
			Y: float32(y) * Cell,
		})
	}
	c.X = x
	c.Y = y
}

func (c *Piece) CreateRenderer() fyne.WidgetRenderer {
	circle := canvas.NewCircle(PieceBorder)
	circle.StrokeColor = PiecePad
	circle.StrokeWidth = 3

	text := canvas.NewText(c.Name, c.FontColor)
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle = fyne.TextStyle{Bold: true}

	return &chessPieceRenderer{circle: circle, text: text, objects: []fyne.CanvasObject{circle, text}}
}

type chessPieceRenderer struct {
	circle  *canvas.Circle
	text    *canvas.Text
	objects []fyne.CanvasObject
}

func (r *chessPieceRenderer) Layout(size fyne.Size) {
	r.circle.Resize(size)
	r.text.Resize(size)
	r.text.TextSize = size.Width * 0.7 // 计算字体大小
}

func (r *chessPieceRenderer) MinSize() fyne.Size {
	return fyne.NewSize(100, 100)
}

func (r *chessPieceRenderer) Refresh() {
	r.circle.Refresh()
	r.text.Refresh()
}

func (r *chessPieceRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *chessPieceRenderer) Destroy() {}
