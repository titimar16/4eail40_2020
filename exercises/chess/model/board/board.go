package board

import (
	"github.com/chess/model/coord"
	"github.com/chess/model/piece"
)

// Classic [8][8]piece.Piece
type Classic [8][8]piece.Piece

func (c *Classic) String() string {
	panic("implement me")
}

func (c *Classic) PieceAt(at coord.ChessCoordinates) piece.Piece {
	panic("implement me")
}

func (c *Classic) MovePiece(from, to coord.ChessCoordinates) error {
	panic("implement me")
}
