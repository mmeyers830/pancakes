// Package pancakes handles the flipping and counting of pancakes flips
package pancakes

import (
	"errors"
	"fmt"
)

// errNotFound is returned when no pancake with the specified face is up
var errNotFound = errors.New("pancake not found")

// Flipper represents the spatula that flips the pancakes and keeps track of the flips
type Flipper struct {
	HappyChar rune
	PlainChar rune
}

// CountFlips optimally flips pancakes on the stack until they are all smiley face up (HappyChar)
// and returns the number of flips required
func (f *Flipper) CountFlips(stack []rune) (int, error) {
	flipCount := 0
	var err error
	// we remove the happy pancakes from the bottom of the stack as we go along, so stop when it's empty
	for !f.isStackAllHappy(stack) {
		// first we flip all the pancakes at the top that are HappyChar, so that when we flip the whole
		// stack, they will be happy again
		if stack[0] == f.HappyChar {
			stack, err = f.flipAllHappyOnTop(stack)
			if err != nil {
				return 0, fmt.Errorf("failed to flip happy top pancakes: %w", err)
			}
			flipCount++
		}

		// now we exclude all the happy pancakes at the bottom since they are already in their final state
		stack, err = f.removeAllHappyFromBottom(stack)
		if err != nil {
			return 0, fmt.Errorf("failed to remove happy bottom pancakes: %w", err)
		}

		// then we flip the whole stack that's left
		stack = f.flipStack(stack)
		flipCount++
	}

	return flipCount, nil
}

// flipAllHappyOnTop flips all the happy pancakes on top of the stack
func (f *Flipper) flipAllHappyOnTop(stack []rune) ([]rune, error) {
	firstPlainIndex, err := f.findFirstPlain(stack)
	if err == errNotFound {
		// if we didn't find any, then we flip the whole stack
		firstPlainIndex = len(stack)
	} else if err != nil {
		return nil, fmt.Errorf("failed to find first plain pancake: %w", err)
	}
	stack = f.flipTopStack(stack, firstPlainIndex)
	return stack, nil
}

// removeAllHappyFromBottom removes all happy pancakes from the bottom of the stack. If there are none then
// it will return the existing stack, and if they are all happy, it will return an empty stack
func (f *Flipper) removeAllHappyFromBottom(stack []rune) ([]rune, error) {
	lastPlainIndex, err := f.findLastPlain(stack)
	if err == errNotFound {
		// all pancakes are smiley up, so empty the whole array
		return []rune{}, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to find last plain pancake: %w", err)
	}
	stack = stack[:lastPlainIndex+1]
	return stack, nil
}

// findFirstPlain finds the first plain facing pancake in the stack
// if there are no pancakes with plain side up, it will return errNotFound
func (f *Flipper) findFirstPlain(stack []rune) (int, error) {
	for i, flapjack := range stack {
		if flapjack == f.PlainChar {
			return i, nil
		}
	}
	return 0, errNotFound
}

// findLastPlain finds the last plain facing pancake in the stack
// if there are no pancakes with plain side up, it will return errNotFound
func (f *Flipper) findLastPlain(stack []rune) (int, error) {
	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] == f.PlainChar {
			return i, nil
		}
	}
	return 0, errNotFound
}

// isStackAllHappy will return true if the whole pancake stack is smiley face up (HappyChar) or empty
func (f *Flipper) isStackAllHappy(stack []rune) bool {
	for _, r := range stack {
		if r == f.PlainChar {
			return false
		}
	}
	return true
}

// flipTopStack flips all pancakes with lower index than index and returns the stack with the flipped portion on top
func (f *Flipper) flipTopStack(stack []rune, index int) []rune {
	return append(f.flipStack(stack[:index]), stack[index:]...)
}

// flipStack flips the whole pancake stack that is passed in and returns it
func (f *Flipper) flipStack(stack []rune) []rune {
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = f.flipSingle(stack[j]), f.flipSingle(stack[i])
	}
	if len(stack)%2 != 0 {
		middleIndex := len(stack) / 2
		stack[middleIndex] = f.flipSingle(stack[middleIndex])
	}
	return stack
}

// flipSingle takes in a HappyChar or PlainChar and returns the opposite
func (f *Flipper) flipSingle(r rune) rune {
	if r == f.HappyChar {
		return f.PlainChar
	}
	return f.HappyChar
}
