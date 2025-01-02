package tic_tac_toe

type Player struct {
	Name string
	PieceType
}

func (p *Player) SetValues(name, pieceType string) {
	p.Name = name
	p.PieceType.Type = pieceType
}

func (p *Player) GetName() string {
	return p.Name
}

func (p *Player) GetPieceType() string {
	return p.PieceType.Type
}
