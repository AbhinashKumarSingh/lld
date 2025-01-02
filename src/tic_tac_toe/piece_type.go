package tic_tac_toe

type PieceType struct {
	Type string
}

func (p PieceType) PlayingPiece(pieceType string) {
	p.Type = pieceType
}
