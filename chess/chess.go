package chess

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

var Cell float32 = 50            // 棋盘格子大小
var Border float32 = 40          // 棋盘边框大小
var LineStrokeWidth float32 = 1  // 棋盘线条大小
var PieceSize float32 = 43       // 棋子大小
var PieceStrokeWidth float32 = 3 // 棋子边框大小

type Chess struct {
	//*Config
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

func NewChess() *Chess {
	app := app.New()
	MainWindow := app.NewWindow("MainWindow")
	MainWindow.Resize(fyne.NewSize(800, 600))
	chess := &Chess{App: app, MainWindow: MainWindow}

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
	PieceContainer := chess.DrawPiece()

	// 容器
	ChessContainer := container.NewWithoutLayout(
		ChessBoard, LineContainer, PieceContainer,
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

// DrawPiece 画棋子
func (chess *Chess) DrawPiece() *fyne.Container {
	chess.Piece[0][0] = NewPiece("車", color.Black, 0, 0, PieceSize, PieceSize)
	chess.Piece[1][0] = NewPiece("馬", color.Black, 1, 0, PieceSize, PieceSize)
	chess.Piece[2][0] = NewPiece("象", color.Black, 2, 0, PieceSize, PieceSize)
	chess.Piece[3][0] = NewPiece("士", color.Black, 3, 0, PieceSize, PieceSize)
	chess.Piece[4][0] = NewPiece("將", color.Black, 4, 0, PieceSize, PieceSize)
	chess.Piece[5][0] = NewPiece("士", color.Black, 5, 0, PieceSize, PieceSize)
	chess.Piece[6][0] = NewPiece("象", color.Black, 6, 0, PieceSize, PieceSize)
	chess.Piece[7][0] = NewPiece("馬", color.Black, 7, 0, PieceSize, PieceSize)
	chess.Piece[8][0] = NewPiece("車", color.Black, 8, 0, PieceSize, PieceSize)
	chess.Piece[1][2] = NewPiece("炮", color.Black, 1, 2, PieceSize, PieceSize)
	chess.Piece[7][2] = NewPiece("炮", color.Black, 7, 2, PieceSize, PieceSize)
	chess.Piece[0][3] = NewPiece("卒", color.Black, 0, 3, PieceSize, PieceSize)
	chess.Piece[2][3] = NewPiece("卒", color.Black, 2, 3, PieceSize, PieceSize)
	chess.Piece[4][3] = NewPiece("卒", color.Black, 4, 3, PieceSize, PieceSize)
	chess.Piece[6][3] = NewPiece("卒", color.Black, 6, 3, PieceSize, PieceSize)
	chess.Piece[8][3] = NewPiece("卒", color.Black, 8, 3, PieceSize, PieceSize)

	chess.Piece[0][9] = NewPiece("車", Red, 0, 9, PieceSize, PieceSize)
	chess.Piece[1][9] = NewPiece("馬", Red, 1, 9, PieceSize, PieceSize)
	chess.Piece[2][9] = NewPiece("相", Red, 2, 9, PieceSize, PieceSize)
	chess.Piece[3][9] = NewPiece("仕", Red, 3, 9, PieceSize, PieceSize)
	chess.Piece[4][9] = NewPiece("帥", Red, 4, 9, PieceSize, PieceSize)
	chess.Piece[5][9] = NewPiece("仕", Red, 5, 9, PieceSize, PieceSize)
	chess.Piece[6][9] = NewPiece("相", Red, 6, 9, PieceSize, PieceSize)
	chess.Piece[7][9] = NewPiece("馬", Red, 7, 9, PieceSize, PieceSize)
	chess.Piece[8][9] = NewPiece("車", Red, 8, 9, PieceSize, PieceSize)
	chess.Piece[1][7] = NewPiece("炮", Red, 1, 7, PieceSize, PieceSize)
	chess.Piece[7][7] = NewPiece("炮", Red, 7, 7, PieceSize, PieceSize)
	chess.Piece[0][6] = NewPiece("兵", Red, 0, 6, PieceSize, PieceSize)
	chess.Piece[2][6] = NewPiece("兵", Red, 2, 6, PieceSize, PieceSize)
	chess.Piece[4][6] = NewPiece("兵", Red, 4, 6, PieceSize, PieceSize)
	chess.Piece[6][6] = NewPiece("兵", Red, 6, 6, PieceSize, PieceSize)
	chess.Piece[8][6] = NewPiece("兵", Red, 8, 6, PieceSize, PieceSize)

	var Pieces []fyne.CanvasObject
	for _, v := range chess.Piece {
		for _, vv := range v {
			if vv != nil {
				Pieces = append(Pieces, vv)
			}
		}
	}
	return container.NewWithoutLayout(Pieces...)
}

// DrawLine 画棋盘
func (chess *Chess) DrawLine() *fyne.Container {
	hor0 := NewLine(Border, Border+Cell*0, Border+Cell*8, Border+Cell*0, LineStrokeWidth*2, color.Black)
	hor1 := NewLine(Border, Border+Cell*1, Border+Cell*8, Border+Cell*1, LineStrokeWidth, color.Black)
	hor2 := NewLine(Border, Border+Cell*2, Border+Cell*8, Border+Cell*2, LineStrokeWidth, color.Black)
	hor3 := NewLine(Border, Border+Cell*3, Border+Cell*8, Border+Cell*3, LineStrokeWidth, color.Black)
	hor4 := NewLine(Border, Border+Cell*4, Border+Cell*8, Border+Cell*4, LineStrokeWidth, color.Black)
	hor5 := NewLine(Border, Border+Cell*5, Border+Cell*8, Border+Cell*5, LineStrokeWidth, color.Black)
	hor6 := NewLine(Border, Border+Cell*6, Border+Cell*8, Border+Cell*6, LineStrokeWidth, color.Black)
	hor7 := NewLine(Border, Border+Cell*7, Border+Cell*8, Border+Cell*7, LineStrokeWidth, color.Black)
	hor8 := NewLine(Border, Border+Cell*8, Border+Cell*8, Border+Cell*8, LineStrokeWidth, color.Black)
	hor9 := NewLine(Border, Border+Cell*9, Border+Cell*8, Border+Cell*9, LineStrokeWidth*2, color.Black)

	// 垂直线
	ver0 := NewLine(Border, Border, Border, Border+Cell*9, LineStrokeWidth*2, color.Black)
	ver8 := NewLine(Border+Cell*8, Border, Border+Cell*8, Border+Cell*9, LineStrokeWidth*2, color.Black)

	ver1 := NewLine(Border+Cell*1, Border, Border+Cell*1, Border+Cell*4, LineStrokeWidth, color.Black)
	ver2 := NewLine(Border+Cell*2, Border, Border+Cell*2, Border+Cell*4, LineStrokeWidth, color.Black)
	ver3 := NewLine(Border+Cell*3, Border, Border+Cell*3, Border+Cell*4, LineStrokeWidth, color.Black)
	ver4 := NewLine(Border+Cell*4, Border, Border+Cell*4, Border+Cell*4, LineStrokeWidth, color.Black)
	ver5 := NewLine(Border+Cell*5, Border, Border+Cell*5, Border+Cell*4, LineStrokeWidth, color.Black)
	ver6 := NewLine(Border+Cell*6, Border, Border+Cell*6, Border+Cell*4, LineStrokeWidth, color.Black)
	ver7 := NewLine(Border+Cell*7, Border, Border+Cell*7, Border+Cell*4, LineStrokeWidth, color.Black)

	ver11 := NewLine(Border+Cell*1, Border+Cell*5, Border+Cell*1, Border+Cell*9, LineStrokeWidth, color.Black)
	ver22 := NewLine(Border+Cell*2, Border+Cell*5, Border+Cell*2, Border+Cell*9, LineStrokeWidth, color.Black)
	ver33 := NewLine(Border+Cell*3, Border+Cell*5, Border+Cell*3, Border+Cell*9, LineStrokeWidth, color.Black)
	ver44 := NewLine(Border+Cell*4, Border+Cell*5, Border+Cell*4, Border+Cell*9, LineStrokeWidth, color.Black)
	ver55 := NewLine(Border+Cell*5, Border+Cell*5, Border+Cell*5, Border+Cell*9, LineStrokeWidth, color.Black)
	ver66 := NewLine(Border+Cell*6, Border+Cell*5, Border+Cell*6, Border+Cell*9, LineStrokeWidth, color.Black)
	ver77 := NewLine(Border+Cell*7, Border+Cell*5, Border+Cell*7, Border+Cell*9, LineStrokeWidth, color.Black)

	// 九宫线
	slash1 := NewLine(Border+Cell*3, Border, Border+Cell*5, Border+Cell*2, LineStrokeWidth*0.5, color.Black)
	slash2 := NewLine(Border+Cell*5, Border, Border+Cell*3, Border+Cell*2, LineStrokeWidth*0.5, color.Black)
	slash3 := NewLine(Border+Cell*3, Border+Cell*9, Border+Cell*5, Border+Cell*7, LineStrokeWidth*0.5, color.Black)
	slash4 := NewLine(Border+Cell*3, Border+Cell*7, Border+Cell*5, Border+Cell*9, LineStrokeWidth*0.5, color.Black)

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
