package tic_tac_toe

import "fmt"

type Game struct {
	Players []Player
	Board
}

// Add a player to the end of the deque
func (g *Game) AddPlayerToEnd(player Player) {
	g.Players = append(g.Players, player)
}

// Add a player to the front of the deque
func (g *Game) AddPlayerToFront(player Player) {
	g.Players = append([]Player{player}, g.Players...)
}

// Remove a player from the front of the deque
func (g *Game) RemovePlayerFromFront() Player {
	if len(g.Players) == 0 {
		return Player{}
	}
	frontPlayer := g.Players[0]
	g.Players = g.Players[1:]
	g.Players = append(g.Players, frontPlayer)
	return frontPlayer
}

func (b *Game) RemovePlayerFromLast() {
	if len(b.Players) == 0 {
		// return Player{} // return an empty player if no players are left
	}
	// Get the player at the last (last player in the slice)
	// lastPlayer := b.Players[len(b.Players)-1]
	// Remove the player from the end of the slice (pop from last)
	b.Players = b.Players[:len(b.Players)-1]
	// return lastPlayer
}

func (b *Board) CheckWinner() bool {
	// Check Rows
	for i := 0; i < len(b.PlayingPiece); i++ {
		// Ensure the first cell is not empty
		if b.PlayingPiece[i][0] != "" {
			// Check if all cells in the row are the same
			winner := true
			for j := 1; j < len(b.PlayingPiece[i]); j++ {
				if b.PlayingPiece[i][j] != b.PlayingPiece[i][0] {
					winner = false
					break
				}
			}
			if winner {
				return true // Winner found
			}
		}
	}

	// Check Columns
	for j := 0; j < len(b.PlayingPiece[0]); j++ {
		// Ensure the first cell is not empty
		if b.PlayingPiece[0][j] != "" {
			// Check if all cells in the column are the same
			winner := true
			for i := 1; i < len(b.PlayingPiece); i++ {
				if b.PlayingPiece[i][j] != b.PlayingPiece[0][j] {
					winner = false
					break
				}
			}
			if winner {
				return true // Winner found
			}
		}
	}

	// Check Diagonal (top-left to bottom-right)
	if b.PlayingPiece[0][0] != "" {
		winner := true
		for i := 1; i < len(b.PlayingPiece); i++ {
			if b.PlayingPiece[i][i] != b.PlayingPiece[0][0] {
				winner = false
				break
			}
		}
		if winner {
			return true // Winner found
		}
	}

	// Check Anti-Diagonal (top-right to bottom-left)
	if b.PlayingPiece[0][len(b.PlayingPiece)-1] != "" {
		winner := true
		for i := 1; i < len(b.PlayingPiece); i++ {
			if b.PlayingPiece[i][len(b.PlayingPiece)-i-1] != b.PlayingPiece[0][len(b.PlayingPiece)-1] {
				winner = false
				break
			}
		}
		if winner {
			return true // Winner found
		}
	}

	// No winner
	return false
}

func NewGame(boardSize int, playerNames []string) *Game {
	// Initialize the board
	board := Board{
		Size:         boardSize,
		PlayingPiece: make([][]string, boardSize),
	}

	for i := range board.PlayingPiece {
		board.PlayingPiece[i] = make([]string, boardSize)
	}

	// Create and assign players with their respective signs
	signs := []rune{'X', 'O'} // Possible signs
	players := make([]Player, len(playerNames))

	for i, name := range playerNames {
		players[i] = Player{
			Name: name,
			PieceType: PieceType{
				Type: PeiceTypeConstant[i%len(signs)], // Cycle through the signs,
			},
		}
	}

	return &Game{
		Players: players,
		Board:   board,
	}
}

func (g *Game) StartGame() {
	winner := false
	for !winner {
		freeSpaceExist := g.Board.CheckFreeSpaceExist()
		if !freeSpaceExist {
			winner = true
			continue
		}

		currentPlayer := g.RemovePlayerFromFront()
		fmt.Printf("Player %s: Enter row, column\n", currentPlayer.Name)
		var row, col int

		fmt.Scan(&row)
		fmt.Scan(&col)
		if !g.Board.FillPiece(currentPlayer.PieceType.Type, row, col) {
			fmt.Printf("incorrect position try again")
			g.RemovePlayerFromLast()
			g.AddPlayerToFront(currentPlayer)

			continue
		}

		isWinner := g.CheckWinner()
		if isWinner {
			fmt.Printf("winner %s", currentPlayer.Name)
			break
		}
		g.Board.PrintBoard()

	}
}
