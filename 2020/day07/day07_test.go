package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_inputToBags(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name    string
		args    args
		want    map[Bag][]Bag
		wantErr bool
	}{
		{
			name: "convert input to bags",
			args: args{
				input: []string{
					"light red bags contain 1 bright white bag, 2 muted yellow bags.",
					"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
					"bright white bags contain 1 shiny gold bag.",
					"faded blue bags contain no other bags.",
				},
			},
			want: map[Bag][]Bag{
				"light red": {
					"bright white",
					"muted yellow",
				},
				"dark orange": {
					"bright white",
					"muted yellow",
				},
				"bright white": {
					"shiny gold",
				},
				"faded blue": {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := inputToBags(tt.args.input)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getBags(t *testing.T) {
	type args struct {
		bags     map[Bag][]Bag
		wantBags []Bag
	}
	tests := []struct {
		name string
		args args
		want []Bag
	}{
		{
			name: "get bags depth 1",
			args: args{
				bags: map[Bag][]Bag{
					"light red": {
						"bright white",
						"muted yellow",
					},
					"dark orange": {
						"bright white",
						"muted yellow",
					},
					"bright white": {
						"shiny gold",
					},
					"faded blue": {},
				},
				wantBags: []Bag{"bright white"},
			},
			want: []Bag{"light red", "dark orange"},
		},
		{
			name: "get bags depth 2",
			args: args{
				bags: map[Bag][]Bag{
					"light red": {
						"bright white",
						"muted yellow",
					},
					"dark orange": {
						"bright white",
						"muted yellow",
					},
					"bright white": {
						"shiny gold",
					},
					"faded blue": {},
				},
				wantBags: []Bag{"dark orange", "light red"},
			},
			want: []Bag{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Finder{FoundBags: make(map[Bag]bool)}

			got := f.getBags(tt.args.bags, tt.args.wantBags)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDay07(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testinput",
			args: args{filename: "testinput.txt"},
			want: 4,
		},
		{
			name: "real input",
			args: args{filename: "input.txt"},
			want: 238,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Day07(tt.args.filename)
			assert.Equal(t, tt.want, got)
		})
	}
}
