package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		name     string
		position int
		want     int
	}{
		{
			name:     "0",
			position: 0,
			want:     0,
		},
		{
			name:     "1",
			position: 1,
			want:     1,
		},
		{
			name:     "2",
			position: 2,
			want:     1,
		},
		{
			name:     "3",
			position: 3,
			want:     2,
		},
		{
			name:     "5",
			position: 5,
			want:     5,
		}, {
			name:     "10",
			position: 10,
			want:     55,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, fibonacci(tc.position))
		})
	}
}
