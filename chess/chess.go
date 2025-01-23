package chess

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

var Cell float32
var Border float32

type Config struct {
	Cell            float32 // 棋子间距
	Border          float32 // 边框间距
	LineStrokeWidth float32 // 线条宽度
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

	// 棋子
	Piece [9][10]*Piece
}

func InitConfig() *Config {
	Cell = 50
	Border = 40
	return &Config{
		Cell:            Cell,
		Border:          Border,
		LineStrokeWidth: 1,
	}
}

func NewChess() *Chess {
	app := app.New()
	MainWindow := app.NewWindow("MainWindow")
	MainWindow.Resize(fyne.NewSize(800, 600))
	chess := &Chess{Config: InitConfig(), App: app, MainWindow: MainWindow}

	// 聊天面板
	chess.CreateChatBoard()
	// 棋盘面板
	chess.CreateChessBoard()

	MainWindow.SetContent(container.NewWithoutLayout(chess.ChatContainer, chess.ChessContainer))
	return chess
}

// CreateChessBoard 创建棋盘面板
func (chess *Chess) CreateChessBoard() {
	ChessBoard := canvas.NewRectangle(Yellow)
	ChessBoard.Resize(fyne.NewSize(700, 600))

	// 棋盘
	LineContainer := chess.DrawLine()

	// 棋子
	chess.Piece[0][0] = NewPiece("車", color.Black, 0, 0, 30, 30)

	// 容器
	ChessContainer := container.NewWithoutLayout(
		ChessBoard, LineContainer,
		chess.Piece[0][0],
	)
	ChessContainer.Resize(fyne.NewSize(700, 600))
	ChessContainer.Move(fyne.NewPos(300, 0))

	chess.ChessBoard = ChessBoard
	chess.ChessContainer = ChessContainer
}

// CreateChatBoard 创建聊天面板
func (chess *Chess) CreateChatBoard() {
	ChatBoard := canvas.NewRectangle(color.Black)
	ChatBoard.Resize(fyne.NewSize(300, 600))
	ChatContainer := container.NewWithoutLayout(ChatBoard)
	ChatContainer.Resize(fyne.NewSize(300, 600))

	chess.ChatBoard = ChatBoard
	chess.ChatContainer = ChatContainer
}

// DrawLine 画棋盘
func (chess *Chess) DrawLine() *fyne.Container {
	// 水平线
	hor0 := NewLine(chess.Border, chess.Border+chess.Cell*0, chess.Border+chess.Cell*8, chess.Border+chess.Cell*0, chess.LineStrokeWidth*2, color.Black)
	hor1 := NewLine(chess.Border, chess.Border+chess.Cell*1, chess.Border+chess.Cell*8, chess.Border+chess.Cell*1, chess.LineStrokeWidth, color.Black)
	hor2 := NewLine(chess.Border, chess.Border+chess.Cell*2, chess.Border+chess.Cell*8, chess.Border+chess.Cell*2, chess.LineStrokeWidth, color.Black)
	hor3 := NewLine(chess.Border, chess.Border+chess.Cell*3, chess.Border+chess.Cell*8, chess.Border+chess.Cell*3, chess.LineStrokeWidth, color.Black)
	hor4 := NewLine(chess.Border, chess.Border+chess.Cell*4, chess.Border+chess.Cell*8, chess.Border+chess.Cell*4, chess.LineStrokeWidth, color.Black)
	hor5 := NewLine(chess.Border, chess.Border+chess.Cell*5, chess.Border+chess.Cell*8, chess.Border+chess.Cell*5, chess.LineStrokeWidth, color.Black)
	hor6 := NewLine(chess.Border, chess.Border+chess.Cell*6, chess.Border+chess.Cell*8, chess.Border+chess.Cell*6, chess.LineStrokeWidth, color.Black)
	hor7 := NewLine(chess.Border, chess.Border+chess.Cell*7, chess.Border+chess.Cell*8, chess.Border+chess.Cell*7, chess.LineStrokeWidth, color.Black)
	hor8 := NewLine(chess.Border, chess.Border+chess.Cell*8, chess.Border+chess.Cell*8, chess.Border+chess.Cell*8, chess.LineStrokeWidth, color.Black)
	hor9 := NewLine(chess.Border, chess.Border+chess.Cell*9, chess.Border+chess.Cell*8, chess.Border+chess.Cell*9, chess.LineStrokeWidth*2, color.Black)

	// 垂直线
	ver0 := NewLine(chess.Border, chess.Border, chess.Border, chess.Border+chess.Cell*9, chess.LineStrokeWidth*2, color.Black)
	ver8 := NewLine(chess.Border+chess.Cell*8, chess.Border, chess.Border+chess.Cell*8, chess.Border+chess.Cell*9, chess.LineStrokeWidth*2, color.Black)

	ver1 := NewLine(chess.Border+chess.Cell*1, chess.Border, chess.Border+chess.Cell*1, chess.Border+chess.Cell*4, chess.LineStrokeWidth, color.Black)
	ver2 := NewLine(chess.Border+chess.Cell*2, chess.Border, chess.Border+chess.Cell*2, chess.Border+chess.Cell*4, chess.LineStrokeWidth, color.Black)
	ver3 := NewLine(chess.Border+chess.Cell*3, chess.Border, chess.Border+chess.Cell*3, chess.Border+chess.Cell*4, chess.LineStrokeWidth, color.Black)
	ver4 := NewLine(chess.Border+chess.Cell*4, chess.Border, chess.Border+chess.Cell*4, chess.Border+chess.Cell*4, chess.LineStrokeWidth, color.Black)
	ver5 := NewLine(chess.Border+chess.Cell*5, chess.Border, chess.Border+chess.Cell*5, chess.Border+chess.Cell*4, chess.LineStrokeWidth, color.Black)
	ver6 := NewLine(chess.Border+chess.Cell*6, chess.Border, chess.Border+chess.Cell*6, chess.Border+chess.Cell*4, chess.LineStrokeWidth, color.Black)
	ver7 := NewLine(chess.Border+chess.Cell*7, chess.Border, chess.Border+chess.Cell*7, chess.Border+chess.Cell*4, chess.LineStrokeWidth, color.Black)

	ver11 := NewLine(chess.Border+chess.Cell*1, chess.Border+chess.Cell*5, chess.Border+chess.Cell*1, chess.Border+chess.Cell*9, chess.LineStrokeWidth, color.Black)
	ver22 := NewLine(chess.Border+chess.Cell*2, chess.Border+chess.Cell*5, chess.Border+chess.Cell*2, chess.Border+chess.Cell*9, chess.LineStrokeWidth, color.Black)
	ver33 := NewLine(chess.Border+chess.Cell*3, chess.Border+chess.Cell*5, chess.Border+chess.Cell*3, chess.Border+chess.Cell*9, chess.LineStrokeWidth, color.Black)
	ver44 := NewLine(chess.Border+chess.Cell*4, chess.Border+chess.Cell*5, chess.Border+chess.Cell*4, chess.Border+chess.Cell*9, chess.LineStrokeWidth, color.Black)
	ver55 := NewLine(chess.Border+chess.Cell*5, chess.Border+chess.Cell*5, chess.Border+chess.Cell*5, chess.Border+chess.Cell*9, chess.LineStrokeWidth, color.Black)
	ver66 := NewLine(chess.Border+chess.Cell*6, chess.Border+chess.Cell*5, chess.Border+chess.Cell*6, chess.Border+chess.Cell*9, chess.LineStrokeWidth, color.Black)
	ver77 := NewLine(chess.Border+chess.Cell*7, chess.Border+chess.Cell*5, chess.Border+chess.Cell*7, chess.Border+chess.Cell*9, chess.LineStrokeWidth, color.Black)

	// 九宫线
	slash1 := NewLine(chess.Border+chess.Cell*3, chess.Border, chess.Border+chess.Cell*5, chess.Border+chess.Cell*2, chess.LineStrokeWidth*0.5, color.Black)
	slash2 := NewLine(chess.Border+chess.Cell*5, chess.Border, chess.Border+chess.Cell*3, chess.Border+chess.Cell*2, chess.LineStrokeWidth*0.5, color.Black)
	slash3 := NewLine(chess.Border+chess.Cell*3, chess.Border+chess.Cell*9, chess.Border+chess.Cell*5, chess.Border+chess.Cell*7, chess.LineStrokeWidth*0.5, color.Black)
	slash4 := NewLine(chess.Border+chess.Cell*3, chess.Border+chess.Cell*7, chess.Border+chess.Cell*5, chess.Border+chess.Cell*9, chess.LineStrokeWidth*0.5, color.Black)

	return container.NewWithoutLayout(
		hor0, hor1, hor2, hor3, hor4, hor5, hor6, hor7, hor8, hor9,
		ver0, ver8,
		ver1, ver2, ver3, ver4, ver5, ver6, ver7,
		ver11, ver22, ver33, ver44, ver55, ver66, ver77,
		slash1, slash2, slash3, slash4,
	)
}

// NewLine 画一条直线
func NewLine(x1, y1, x2, y2, StrokeWidth float32, color color.Color) *canvas.Line {
	line := &canvas.Line{
		Position1:   fyne.Position{X: x1, Y: y1},
		Position2:   fyne.Position{X: x2, Y: y2},
		Hidden:      false,
		StrokeColor: color,
		StrokeWidth: StrokeWidth,
	}
	return line
}

func (chess *Chess) Run() {
	chess.MainWindow.Show()
	chess.App.Run()
}
