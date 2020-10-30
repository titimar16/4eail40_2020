package coord

import (
	"testing"
)

func TestNewCartesian(t *testing.T) {
	c := NewCartesian(1, 2)
	if (c.x != 1) || (c.y != 2) {
		t.Errorf("expected 1 and 2 as coordinate")
	}
}

func TestCartesian_Coord(t *testing.T) {
	c := NewCartesian(1, 2)

	tests := map[int]int{
		0: 1,
		1: 2,
	}

	for n, want := range tests {
		t.Run(string(rune(n)), func(t *testing.T) {
			got, err := c.Coord(n)
			if err != nil {
				t.Error(err)
			}
			if got != want {
				t.Errorf("expected %d, but got %d", want, got)
			}
		})
	}

	// test for err
	t.Run("err", func(t *testing.T) {
		_, err := c.Coord(2)
		if err == nil {
			t.Errorf("expected and error for n == 2")
		}
	})
}

func TestCartesian_String(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"A1",
			fields{0, 0},
			"A1",
		},
		{
			"H8",
			fields{7, 7},
			"H8",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cartesian{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
