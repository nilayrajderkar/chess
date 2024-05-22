package gamemanager

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/notnil/chess"
)

type GameStatus string

const (
	ACTIVE    GameStatus = "ACTIVE"
	STALEMATE GameStatus = "STALEMATE"
	CHECKMATE GameStatus = "CHECKMATE"
	CHECK     GameStatus = "CHECK"
)

type Board struct {
	player1     *websocket.Conn
	player2     *websocket.Conn
	status      GameStatus
	game  		*chess.Game
}

func createBoard(player1 *websocket.Conn, player2 *websocket.Conn) *Board {
	game := chess.NewGame()
	return &Board{player1: player1, player2: player2, status: ACTIVE, game: game}
}

func makeMove(game *chess.Game, player *websocket.Conn) {
	
	err := player.WriteMessage(websocket.TextMessage, []byte("Waiting for you to make a move"))
	if err != nil {
		fmt.Print(err)
	}

	_, message, err := player.ReadMessage()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Server has received move:", string(message))

	if game.Outcome() == chess.NoOutcome {
		err = game.MoveStr(string(message))
		if err != nil {
			err = player.WriteMessage(websocket.TextMessage, []byte("Invalid move, try again"))
			if err != nil {
				fmt.Print(err)
			}
		} else {
			err = player.WriteMessage(websocket.TextMessage, []byte("Your move is registered"))
			if err != nil {
				fmt.Print(err)
			}
		}
	} 
	// Have all conditions when game is over
}

func playGame(board *Board) {
	// TODO: Implement game logic
	game := board.game
	for {
		// Check if game is over
		if len(game.Moves()) %2 == 0 {
			makeMove(game, board.player1)
		} else {
			makeMove(game, board.player2)
		}	
	}
}
