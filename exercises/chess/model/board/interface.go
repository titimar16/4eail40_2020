// Package board will contain types and logic for handling chess board(s).
package board

import (
	"fmt"
	"github.com/chess/model/coord"
	"github.com/chess/model/piece"
)

//Board is an interface to a chess board.
// It defines methods for handling pieces on it.

type Board interface {
	fmt.Stringer
	// PieceAt retrieves piece at give coordinates.
	// Return nil if not piece was found.
	PieceAt(at coord.ChessCoordinates) piece.Piece
	// MovePiece moves a piece from given coordinate to given coordinates.
	// Returns and error if destination was occupied
	MovePiece(from, to coord.ChessCoordinates) error
}

//  TODO exo : Implement a ClassicBuilder (don't forget the test(s)) !
type Builder interface {
	AddPiece(p piece.Piece, at coord.ChessCoordinates) Builder
	Build() Board
}
