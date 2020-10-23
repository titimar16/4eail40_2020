// Package piece will handle game pieces.
package piece

import (
	"fmt"
	"github.com/chess/model/coord"
	"github.com/chess/model/player"
)

// Piece represents a game piece.
type Piece interface {
	fmt.Stringer
	//
	Color() player.Color
	//
	Moves() map[coord.ChessCoordinates]bool
}

func test() {
	//var p Piece

	//var move = destination - origin
	// Check if move is contained in a map
	//if p.Moves[move]

	// HashSet
}
