package board

import (
	"github.com/chess/model/coord"
	"github.com/chess/model/piece"
	"reflect"
	"testing"
)

type mockCoord int

// Coord returns x if n==0, y if n==1
func (c mockCoord) Coord(n int) (int, error) {
	return int(c), nil
}

func (c mockCoord) String() string {
	return "1"
}

func TestClassic_MovePiece(t *testing.T) {
	class := Classic{}
	pi := piece.Rook{}
	pi1 := piece.Rook{}
	coordin := coord.NewCartesian(0, 0)
	x, _ := coordin.Coord(0)
	y, _ := coordin.Coord(1)
	class[x][y] = pi
	coordin1 := coord.NewCartesian(7, 0)
	x1, _ := coordin1.Coord(0)
	y1, _ := coordin1.Coord(1)
	class[x1][y1] = pi1

	type args struct {
		from coord.ChessCoordinates
		to   coord.ChessCoordinates
	}
	tests := []struct {
		name    string
		c       Classic
		args    args
		wantErr bool
	}{
		{
			"testmock",
			Classic{},
			args{mockCoord(0), mockCoord(1)},
			true,
		},
		{
			"no piece at from",
			class,
			args{from: coord.NewCartesian(3, 5), to: coordin},
			true,
		},
		{
			"occupied at to",
			class,
			args{from: coordin1, to: coordin},
			true,
		},
		{
			"move accepted",
			class,
			args{from: coordin, to: coord.NewCartesian(4, 4)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.MovePiece(tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("Classic.MovePiece() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClassic_PieceAt(t *testing.T) {
	class := Classic{}
	pi := piece.Rook{}
	coordin := coord.NewCartesian(0, 0)
	x, _ := coordin.Coord(0)
	y, _ := coordin.Coord(1)
	class[x][y] = pi
	type args struct {
		at coord.ChessCoordinates
	}
	tests := []struct {
		name string
		c    Classic
		args args
		want piece.Piece
	}{
		{
			"test A1",
			class,
			args{at: coordin},
			pi,
		},
		{
			"test F5",
			class,
			args{at: coord.NewCartesian(4, 4)},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.PieceAt(tt.args.at); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PieceAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClassic_PlacePieceAt(t *testing.T) {
	class := Classic{}
	pi := piece.Rook{}
	coordin := coord.NewCartesian(0, 0)
	x, _ := coordin.Coord(0)
	y, _ := coordin.Coord(1)
	class[x][y] = pi
	type args struct {
		p  piece.Piece
		at coord.ChessCoordinates
	}
	tests := []struct {
		name    string
		c       Classic
		args    args
		wantErr bool
	}{
		{
			"Test pass",
			class,
			args{p: pi, at: coord.NewCartesian(4, 4)},
			false,
		},
		{
			"Test failed",
			class,
			args{p: pi, at: coordin},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.PlacePieceAt(tt.args.p, tt.args.at); (err != nil) != tt.wantErr {
				t.Errorf("PlacePieceAt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
