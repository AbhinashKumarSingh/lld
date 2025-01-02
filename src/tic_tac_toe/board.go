package tic_tac_toe

import "fmt"

type Board struct {
	Size         int
	PlayingPiece [][]string
}

func (b *Board) SizeOfBoard(size int) {
	b.Size = size
	b.PlayingPiece = make([][]string, b.Size) // Create rows
	for i := range b.PlayingPiece {
		b.PlayingPiece[i] = make([]string, b.Size) // Create columns for each row
	}
}

func (b *Board) PrintBoard() {
	for i := 0; i < len(b.PlayingPiece); i++ {
		for j := 0; j < len(b.PlayingPiece[i]); j++ {
			fmt.Printf("| %s", b.PlayingPiece[i][j])
		}
		fmt.Println("\n")
	}
}

func (b *Board) FillPiece(piece string, row, column int) bool {
	if b.PlayingPiece[row][column] != "" {
		return false
	}
	b.PlayingPiece[row][column] = piece
	return true
}

func (b *Board) CheckFreeSpaceExist() bool {
	for i := 0; i < len(b.PlayingPiece); i++ {
		for j := 0; j < len(b.PlayingPiece[i]); j++ {
			if b.PlayingPiece[i][j] == "" {
				return true
			}
		}
	}
	return false
}
