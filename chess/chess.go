package chess

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

type Config struct {
	Cell int // 棋盘间距
}

type Chess struct {
	*Config
	App fyne.App

	MainWindow fyne.Window

	// 聊天面板
	ChatContainer *fyne.Container
	ChatBoard     *canvas.Rectangle

	// 棋盘面板
	ChessContainer *fyne.Container
	ChessBoard     *canvas.Rectangle
}

func NewChess() *Chess {
	app := app.New()
	MainWindow := app.NewWindow("MainWindow")
	MainWindow.Resize(fyne.NewSize(800, 600))

	// 聊天面板
	ChatBoard := canvas.NewRectangle(color.Black)
	ChatBoard.Resize(fyne.NewSize(300, 600))
	ChatContainer := container.NewWithoutLayout(ChatBoard)
	ChatContainer.Resize(fyne.NewSize(300, 600))

	// 棋盘面板
	ChessBoard := canvas.NewRectangle(Yellow)
	ChessBoard.Resize(fyne.NewSize(700, 600))
	// 画线

	ChessContainer := container.NewWithoutLayout(ChessBoard)
	ChessContainer.Resize(fyne.NewSize(700, 600))
	ChessContainer.Move(fyne.NewPos(300, 0))

	MainWindow.SetContent(container.NewWithoutLayout(ChatContainer, ChessContainer))

	chess := &Chess{
		Config: &Config{
			Cell: 20,
		},
		App:        app,
		MainWindow: MainWindow,

		ChatContainer: ChatContainer,
		ChatBoard:     ChatBoard,

		ChessContainer: ChessContainer,
		ChessBoard:     ChessBoard,
	}
	return chess
}

func (chess *Chess) Run() {
	chess.MainWindow.Show()
	chess.App.Run()
}
