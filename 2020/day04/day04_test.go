package day04

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_day4(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "case1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := day4()
			assert.NoError(t, err)
		})
	}
}

func Test_Run(t *testing.T) {
	res, _ := day4()
	fmt.Println("res:", res)
}
