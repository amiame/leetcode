package main

import (
	"testing"
	//"github.com/stretchr/testify/assert"
)

func TestIsNotUnderAttack(t *testing.T) {
	type testcase struct {
		placedSquares map[coord]interface{}
		square        coord
		want          bool
	}

	tests := []testcase{
		{
			placedSquares: map[coord]interface{}{
				coord{x: 0, y: 0}: true,
			},
			square: coord{x: 1, y: 2},
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run("TestIsNotUnderAttack", func(tc *testing.T) {
			placedSquares = test.placedSquares
			isNotUnderAttack(test.square.x, test.square.y)
		})
	}
}
