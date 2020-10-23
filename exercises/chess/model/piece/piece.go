// Package piece will handle game pieces.
package piece

import "fmt"

// Piece represents a game piece.
type Piece interface {
	fmt.Stringer
}
