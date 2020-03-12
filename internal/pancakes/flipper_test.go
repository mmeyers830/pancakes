package pancakes

import "testing"

type pancakeTest struct {
	name         string
	pancakeStack []rune
	flipCount    int
}

func TestCountFlip(t *testing.T) {
	tables := []pancakeTest{
		{
			name:         "-",
			pancakeStack: []rune{'-'},
			flipCount:    1,
		},
		{
			name:         "-+",
			pancakeStack: []rune{'-', '+'},
			flipCount:    1,
		},
		{
			name:         "+-",
			pancakeStack: []rune{'+', '-'},
			flipCount:    2,
		},
		{
			name:         "+++",
			pancakeStack: []rune{'+', '+', '+'},
			flipCount:    0,
		},
		{
			name:         "--+-",
			pancakeStack: []rune{'-', '-', '+', '-'},
			flipCount:    3,
		},
		{
			name:         "+-+-",
			pancakeStack: []rune{'+', '-', '+', '-'},
			flipCount:    4,
		},
		{
			name:         "----+-",
			pancakeStack: []rune{'-', '-', '-', '-', '+', '-'},
			flipCount:    3,
		},
		{
			name:         "+-+-+-+-+-+",
			pancakeStack: []rune{'+', '-', '+', '-', '+', '-', '+', '-', '+', '-', '+'},
			flipCount:    10,
		},
		{
			name:         "--+-+++-+--+-+++",
			pancakeStack: []rune{'-', '-', '+', '-', '+', '+', '+', '-', '+', '-', '-', '+', '-', '+', '+', '+'},
			flipCount:    9,
		},
	}

	flipper := Flipper{
		HappyChar: '+',
		PlainChar: '-',
	}
	for _, table := range tables {
		t.Run(table.name, func(t *testing.T) {
			flipCount, err := flipper.CountFlips(table.pancakeStack)
			if err != nil {
				t.Errorf("encountered error on flipping: %w", err)
			}
			if table.flipCount != flipCount {
				t.Errorf("expected %d, got %d", table.flipCount, flipCount)
			}
		})
	}
}
