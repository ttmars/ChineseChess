package main

import (
	"ChineseChess/chess"
	"fyne.io/fyne/v2/container"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Circle")
	w.CenterOnScreen()

	circle := canvas.NewCircle(chess.PieceBorder)
	circle.StrokeColor = chess.PiecePad
	circle.StrokeWidth = 5
	circle.Resize(fyne.NewSize(60, 60))

	canvas.NewPositionAnimation(fyne.Position{0, 0}, fyne.Position{60, 60}, time.Millisecond*500, circle.Move).Start()

	c := container.NewWithoutLayout(circle)
	c.Resize(fyne.NewSize(300, 300))
	w.SetContent(c)
	w.Resize(fyne.NewSize(300, 300))
	w.ShowAndRun()
}
