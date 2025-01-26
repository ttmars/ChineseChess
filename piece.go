package main

import (
	"ChineseChess/chess"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"image/color"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Chinese Chess Piece")
	myWindow.CenterOnScreen()

	chess.CellSize = 50
	chessPiece := chess.NewPiece("帥", color.Black, 3, 3, 60, 60)
	chessPiece.SetPosition(0, 0, true)

	c := container.NewWithoutLayout(chessPiece)
	c.Resize(fyne.NewSize(400, 400))

	myWindow.SetContent(c)
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.ShowAndRun()
}
