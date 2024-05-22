package gamemanager

import (
	"fmt"

	"github.com/gorilla/websocket"
)


var inQueue *websocket.Conn = nil
var boards []*Board

func GameManager(conn *websocket.Conn) {
    if inQueue == nil {
        err := conn.WriteMessage(websocket.TextMessage, []byte("finding an opponent..."))
        if err != nil {
            fmt.Print(err)
        }
        inQueue = conn
    } else {
        err := conn.WriteMessage(websocket.TextMessage, []byte("Pairing with opponent"))
        if err != nil {
            fmt.Print(err)
        }
        defer conn.Close()
        
        // Notify both players that they are paired
        pairingMessage := fmt.Sprintf("Pairing %s and %s", inQueue.RemoteAddr().String(), conn.RemoteAddr().String())
        err = conn.WriteMessage(websocket.TextMessage, []byte(pairingMessage))
        if err != nil {
            fmt.Print(err)
        }
        err = inQueue.WriteMessage(websocket.TextMessage, []byte(pairingMessage))
        if err != nil {
            fmt.Print(err)
        }
        
		startGame(inQueue, conn)
        inQueue = nil
    }
}

func startGame(player1 *websocket.Conn, player2 *websocket.Conn) {
    board := createBoard(player1, player2)
    // Add game to currently running games
    boards = append(boards, board)
    playGame(board)
}

// TODO: Implement end game function when socket connection is closed or game is over