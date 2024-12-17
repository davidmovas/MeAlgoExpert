package DepthFirstSearch

import (
	"reflect"
	"testing"
)

func Test_DepthFirstSearch(t *testing.T) {
	testFunctions := map[string]func(root *Node, array []string) []string{
		"DepthFirstSearch_V1": DepthFirstSearch,
	}

	type args struct {
		node  *Node
		array []string
	}

	tests := []struct {
		args args
		want []string
	}{
		{
			args: args{
				node: &Node{
					name: "A",
					children: []*Node{
						{
							name: "B",
							children: []*Node{
								{
									name: "E",
								},
								{
									name: "F",
									children: []*Node{
										{
											name: "I",
										},
										{
											name: "J",
										},
									},
								},
							},
						},
						{
							name: "C",
						},
						{
							name: "D",
							children: []*Node{
								{
									name: "G",
									children: []*Node{
										{
											name: "K",
										},
									},
								},
								{
									name: "H",
								},
							},
						},
					},
				},
			},
			want: []string{"A", "B", "E", "F", "I", "J", "C", "D", "G", "K", "H"},
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.node, tt.args.array); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
