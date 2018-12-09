package brickcolor

import "testing"

func TestAll(t *testing.T) {
	if got, want := len(All), 208; got < want {
		t.Fatalf("Only found %d brick colors, want at least %d", got, want)
	}

	if got, want := All[0], White; got != want {
		t.Fatalf("All[0] = %v, want %v", got, want)
	}
}

func TestBrickColors(t *testing.T) {
	for _, tc := range []struct {
		bc     BrickColor
		name   string
		number int
		hex    string
	}{
		{White, "White", 1, "F2F3F3"},
		{MediumRoyalBlue, "Medium Royal blue", 213, "6C81B7"},
	} {
		if got, want := tc.bc.Name, tc.name; got != want {
			t.Fatalf("tc.bc.Name = %q, want %q", got, want)
		}

		if got, want := tc.bc.Number, tc.number; got != want {
			t.Fatalf("tc.bc.Number = %d, want %d", got, want)
		}

		if got, want := tc.bc.Hex, tc.hex; got != want {
			t.Fatalf("tc.bc.Hex = %q, want %q", got, want)
		}
	}
}
